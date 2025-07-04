// SPDX-License-Identifier: Apache-2.0

//lint:file-ignore SA1019 Using a deprecated function, variable, constant or field

package plan

import (
	"fmt"
	"runtime/debug"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v4"
	"github.com/substrait-io/substrait-go/v4/expr"
	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/plan/internal"
	"github.com/substrait-io/substrait-go/v4/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/anypb"
)

var CurrentVersion = types.Version{
	MajorNumber: 0,
	MinorNumber: 29,
	PatchNumber: 0,
	Producer:    "substrait-go",
}

func init() {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if strings.HasPrefix(dep.Path, "github.com/substrait-io/substrait-go/v4") {
				CurrentVersion.Producer += " " + dep.Version
				break
			}
		}

		var goarch, goos string
		for _, s := range info.Settings {
			if s.Key == "GOARCH" {
				goarch = s.Value
			} else if s.Key == "GOOS" {
				goos = s.Value
			}
		}

		if goos != "" && goarch != "" {
			CurrentVersion.Producer += " " + goos + "/" + goarch
		}
	}
}

// groupingExprs takes 2-dimensional slice of expressions and returns
// a single slice of unique expressions and a slice of references to
// the unique expressions for each group.
func groupingExprs(groups [][]expr.Expression) ([]expr.Expression, [][]uint32) {
	groupingExpressions := make([]expr.Expression, 0)
	groupingReferences := make([][]uint32, 0)
	for _, group := range groups {
		refs := make([]uint32, 0)
		for _, expr := range group {
			existingExpr := false
			for eIndex, existing := range groupingExpressions {
				if existing.Equals(expr) {
					existingExpr = true
					refs = append(refs, uint32(eIndex))
					break
				}
			}
			if !existingExpr {
				groupingExpressions = append(groupingExpressions, expr)
				refs = append(refs, uint32(len(groupingExpressions)-1))
			}
		}
		groupingReferences = append(groupingReferences, refs)
	}
	return groupingExpressions, groupingReferences
}

// Relation is either a Root relation (a relation + list of column names)
// or another relation (such as a CTE or other reference).
type Relation struct {
	root *Root
	rel  Rel
}

func (r *Relation) FromProto(p *proto.PlanRel, reg expr.ExtensionRegistry) error {
	r.root, r.rel = nil, nil

	switch rel := p.RelType.(type) {
	case *proto.PlanRel_Rel:
		input, err := RelFromProto(rel.Rel, reg)
		if err != nil {
			return err
		}

		r.rel = input
		return nil
	case *proto.PlanRel_Root:
		input, err := RelFromProto(rel.Root.Input, reg)
		if err != nil {
			return err
		}

		r.root = &Root{
			input: input,
			names: rel.Root.Names,
		}
		return nil
	}

	return fmt.Errorf("%w: no rel or root set", substraitgo.ErrInvalidRel)
}

// IsRoot returns true if this is the root of the plan Relation tree.
func (r *Relation) IsRoot() bool {
	return r.root != nil
}

func (r *Relation) Root() *Root { return r.root }
func (r *Relation) Rel() Rel    { return r.rel }

func (r *Relation) ToProto() *proto.PlanRel {
	if r.IsRoot() {
		return r.root.ToProtoPlanRel()
	}

	return r.rel.ToProtoPlanRel()
}

type Version interface {
	GetGitHash() string
	GetMajorNumber() uint32
	GetMinorNumber() uint32
	GetPatchNumber() uint32
	GetProducer() string
	fmt.Stringer
}

type AdvancedExtension interface {
	GetEnhancement() *anypb.Any
	GetOptimization() []*anypb.Any
}

// Plan describes a set of operations to complete. For
// compactness, identifiers are normalized at the plan level.
type Plan struct {
	version          *types.Version
	extensions       extensions.Set
	expectedTypeURLs []string
	advExtension     *extensions.AdvancedExtension
	relations        []Relation

	reg expr.ExtensionRegistry
}

// Version returns the substrait version of the plan.
func (p *Plan) Version() Version { return p.version }

// ExtensionRegistry returns the set of registered extensions for this plan
// that it may depend on.
func (p *Plan) ExtensionRegistry() expr.ExtensionRegistry { return p.reg }

// ExpectedTypeURLs is a list of anypb.Any protobuf entities that this plan
// may use. This can be used to warn if some embedded message types are
// unknown. Note that this list may include message types which are ignorable
// (optimizations) or are unused. In many cases, a consumer may be able to
// work with a plan even if one or more message types defined here are unknown.
//
// This returns a clone of the slice, so that the Plan itself remains
// immutable.
func (p *Plan) ExpectedTypeURLs() []string {
	return slices.Clone(p.expectedTypeURLs)
}

// AdvancedExtension returns optional additional extensions associated with
// this plan such as optimizations or enhancements.
func (p *Plan) AdvancedExtension() AdvancedExtension { return p.advExtension }

// Relations returns the full slice of relation trees that are in this plan.
//
// This returns a clone of the internal slice so that the plan itself remains
// immutable.
func (p *Plan) Relations() []Relation {
	return slices.Clone(p.relations)
}

// GetRoots returns a slice containing *only* the relations which are
// considered Root relations from the list (as opposed to CTEs or references).
func (p *Plan) GetRoots() (roots []*Root) {
	roots = make([]*Root, 0, 1)
	for _, r := range p.relations {
		if r.IsRoot() {
			roots = append(roots, r.root)
		}
	}
	return roots
}

// GetNonRootRelations returns a slice containing only the relations from
// this plan which are not considered Roots.
func (p *Plan) GetNonRootRelations() (rels []Rel) {
	rels = make([]Rel, 0, 1)
	for _, r := range p.relations {
		if !r.IsRoot() {
			rels = append(rels, r.rel)
		}
	}
	return rels
}

func FromProto(plan *proto.Plan, c *extensions.Collection) (*Plan, error) {
	ret := &Plan{
		version:          plan.Version,
		extensions:       extensions.GetExtensionSet(plan),
		advExtension:     plan.AdvancedExtensions,
		expectedTypeURLs: plan.ExpectedTypeUrls,
		relations:        make([]Relation, len(plan.Relations)),
	}

	ret.reg = expr.NewExtensionRegistry(ret.extensions, c)
	for i, r := range plan.Relations {
		if err := ret.relations[i].FromProto(r, ret.reg); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (p *Plan) ToProto() (*proto.Plan, error) {
	uris, decls := p.extensions.ToProto()
	relations := make([]*proto.PlanRel, len(p.relations))
	for i, r := range p.relations {
		relations[i] = r.ToProto()
	}
	return &proto.Plan{
		Version:            p.version,
		ExpectedTypeUrls:   p.expectedTypeURLs,
		AdvancedExtensions: p.advExtension,
		Relations:          relations,
		Extensions:         decls,
		ExtensionUris:      uris,
	}, nil
}

// Root is a relation with output field names.
// This is used as the root of a Rel tree.
type Root struct {
	input Rel
	names []string
}

func (r *Root) Input() Rel { return r.input }

// Names are the field names in depth-first order.
func (r *Root) Names() []string { return r.names }

func (r *Root) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Root{
			Root: &proto.RelRoot{
				Input: r.input.ToProto(),
				Names: r.names,
			},
		},
	}
}

func (r *Root) RecordType() types.NamedStruct {
	return types.NamedStruct{
		Names:  r.names,
		Struct: *r.input.RecordType().AsStructType(),
	}
}

type RewriteFunc func(expr.Expression) (expr.Expression, error)

// Rel is a relation tree, representing one of the expected Relation
// types such as Fetch, Sort, Filter, Join, etc.
//
// It contains the common functionality between the different relations
// and should be type switched to determine which relation type it actually
// is for evaluation.
//
// All the exported methods in this interface should be considered constant.
type Rel interface {
	// Hint returns a set of changes to the operation which can influence
	// efficiency and performance but should not impact correctness.
	//
	// This includes things such as Stats and Runtime constraints.
	Hint() *Hint
	// OutputMapping is optional and may be nil. If this is nil, then
	// the result of this relation is the direct output as is (with no
	// reordering or projection of columns). Otherwise this is a slice
	// of indices into the underlying relation's output to map columns
	// to the intended result column order.
	//
	// For example, an output map of [5, 2, 1] means that the expected
	// result should be 3 columns consisting of the 5th, 2nd and 1st
	// output columns from the underlying relation.
	OutputMapping() []int32

	// Remap modifies the current relation by applying the provided
	// mapping to the current relation.  Typically used to remove any
	// unneeded columns or provide them in a different order.  If there
	// already is a mapping on this relation, this provides mapping over
	// the current mapping.
	//
	// If any column numbers specified are outside the currently available
	// input range an error is returned and the mapping is left unchanged.
	Remap(mapping ...int32) (Rel, error)

	// setMapping sets the current mapping and is for internal use.
	// It performs no checks.  End users should call Remap() instead.
	setMapping(mapping []int32)

	// directOutputSchema returns the output record type of the underlying
	// relation as a struct type.  Mapping is not applied.
	directOutputSchema() types.RecordType
	// RecordType returns the types used by all columns returned by
	// this relation after applying any provided mapping.
	RecordType() types.RecordType

	GetAdvancedExtension() *extensions.AdvancedExtension
	// SetAdvancedExtension sets an AdvancedExtension on this Rel, returning any existing one on this Rel. Use `nil` to remove any existing AdvancedExtension.
	SetAdvancedExtension(extension *extensions.AdvancedExtension) (existing *extensions.AdvancedExtension)

	ToProto() *proto.Rel
	ToProtoPlanRel() *proto.PlanRel

	// Copy creates a copy of this relation with new inputs
	Copy(newInputs ...Rel) (Rel, error)

	// GetInputs returns a list of zero or more inputs for this relation
	GetInputs() []Rel

	// CopyWithExpressionRewrite rewrites all expression trees in this Rel. Returns original Rel
	// if no changes were made, otherwise a newly created rel that includes the given expressions
	CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error)
}

func RelFromProto(rel *proto.Rel, reg expr.ExtensionRegistry) (Rel, error) {
	switch rel := rel.RelType.(type) {
	case *proto.Rel_Read:
		var out ReadRel
		switch readType := rel.Read.ReadType.(type) {
		case *proto.ReadRel_ExtensionTable_:
			out = &ExtensionTableReadRel{detail: readType.ExtensionTable.Detail}
		case *proto.ReadRel_LocalFiles_:
			items := make([]FileOrFiles, len(readType.LocalFiles.Items))
			for i, item := range readType.LocalFiles.Items {
				items[i].fromProto(item)
			}
			out = &LocalFileReadRel{
				items:        items,
				advExtension: readType.LocalFiles.AdvancedExtension,
			}
		case *proto.ReadRel_NamedTable_:
			out = &NamedTableReadRel{
				names:        readType.NamedTable.Names,
				advExtension: readType.NamedTable.AdvancedExtension,
			}
		case *proto.ReadRel_VirtualTable_:
			if len(readType.VirtualTable.Values) > 0 && len(readType.VirtualTable.Expressions) > 0 {
				return nil, fmt.Errorf("VirtualTable Value can't have both liternal and expression")
			}
			var values []expr.VirtualTableExpressionValue
			for _, v := range readType.VirtualTable.Values {
				values = append(values, internal.VirtualTableExprFromLiteralProto(v))
			}
			for _, v := range readType.VirtualTable.Expressions {
				row, err := internal.VirtualTableExpressionFromProto(v, reg)
				if err != nil {
					return nil, err
				}
				values = append(values, row)
			}

			out = &VirtualTableReadRel{
				values: values,
			}
		case *proto.ReadRel_IcebergTable_:
			icebergTableType := readType.IcebergTable.TableType
			if icebergTableType == nil {
				return nil, fmt.Errorf("%w: IcebergTableType is required for IcebergTableReadRel", substraitgo.ErrInvalidRel)
			}
			if _, ok := icebergTableType.(IcebergTableType); ok {
				return nil, fmt.Errorf("%w: IcebergTableType must be a string", substraitgo.ErrInvalidRel)
			}
			if direct, ok := icebergTableType.(*proto.ReadRel_IcebergTable_Direct); ok {
				tableType := &Direct{
					MetadataUri: direct.Direct.MetadataUri,
				}
				if snapshotId, ok := direct.Direct.Snapshot.(*proto.ReadRel_IcebergTable_MetadataFileRead_SnapshotId); ok {
					tableType.SnapshotId = SnapshotId(snapshotId.SnapshotId)
				} else if snapshotTimestamp, ok := direct.Direct.Snapshot.(*proto.ReadRel_IcebergTable_MetadataFileRead_SnapshotTimestamp); ok {
					tableType.SnapshotTimestamp = SnapshotTimestamp(snapshotTimestamp.SnapshotTimestamp)
				}
				out = &IcebergTableReadRel{
					tableType: tableType,
				}
			} else {
				return nil, fmt.Errorf("%w: only IcebergTableType Direct is supported", substraitgo.ErrInvalidRel)
			}
		default:
			return nil, fmt.Errorf("%w: unknown ReadRel type", substraitgo.ErrInvalidRel)
		}

		if err := out.fromProtoReadRel(rel.Read, reg); err != nil {
			return nil, err
		}

		return out, nil
	case *proto.Rel_Filter:
		input, err := RelFromProto(rel.Filter.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to FilterRel: %w", err)
		}

		base := input.RecordType()
		cond, err := expr.ExprFromProto(rel.Filter.Condition, &base, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting condition for FilterRel: %w", err)
		}

		out := &FilterRel{
			input:        input,
			cond:         cond,
			advExtension: rel.Filter.AdvancedExtension,
		}
		if rel.Filter.Common != nil {
			out.fromProtoCommon(rel.Filter.Common)
		}
		return out, nil
	case *proto.Rel_Fetch:
		input, err := RelFromProto(rel.Fetch.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to FetchRel: %w", err)
		}

		var offset int64
		if off, ok := rel.Fetch.OffsetMode.(*proto.FetchRel_Offset); ok {
			offset = off.Offset
		} else {
			return nil, fmt.Errorf("%w: missing required Offset field for Fetch Relation", substraitgo.ErrInvalidRel)
		}

		var count int64
		if cnt, ok := rel.Fetch.CountMode.(*proto.FetchRel_Count); ok {
			count = cnt.Count
		} else {
			return nil, fmt.Errorf("%w: missing required Count field for Fetch Relation", substraitgo.ErrInvalidRel)
		}

		out := &FetchRel{
			input:        input,
			offset:       offset,
			count:        count,
			advExtension: rel.Fetch.AdvancedExtension,
		}
		if rel.Fetch.Common != nil {
			out.fromProtoCommon(rel.Fetch.Common)
		}
		return out, nil
	case *proto.Rel_Aggregate:
		input, err := RelFromProto(rel.Aggregate.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to AggregateRel: %w", err)
		}

		base := input.RecordType()
		var groupingExpressions []expr.Expression
		var groupingReferences [][]uint32
		if len(rel.Aggregate.GroupingExpressions) > 0 {
			for _, e := range rel.Aggregate.GroupingExpressions {
				expr, err := expr.ExprFromProto(e, &base, reg)
				if err != nil {
					return nil, fmt.Errorf("error getting grouping expr for AggregateRel: %w", err)
				}
				groupingExpressions = append(groupingExpressions, expr)
			}
			for _, g := range rel.Aggregate.Groupings {
				groupingReferences = append(groupingReferences, g.ExpressionReferences)
			}
		} else { // support old style grouping for backward compatibility
			groups := make([][]expr.Expression, len(rel.Aggregate.Groupings))
			for i, g := range rel.Aggregate.Groupings {
				groups[i] = make([]expr.Expression, len(g.GroupingExpressions))
				for j, e := range g.GroupingExpressions {
					groups[i][j], err = expr.ExprFromProto(e, &base, reg)
					if err != nil {
						return nil, fmt.Errorf("error getting grouping expr [%d][%d] for AggregateRel: %w",
							i, j, err)
					}
				}
			}
			groupingExpressions, groupingReferences = groupingExprs(groups)
		}

		measures := make([]AggRelMeasure, len(rel.Aggregate.Measures))
		for i, m := range rel.Aggregate.Measures {
			measures[i].measure, err = expr.NewAggregateFunctionFromProto(m.Measure, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting AggregateFunction for measure %d: %w", i, err)
			}

			if m.Filter != nil {
				measures[i].filter, err = expr.ExprFromProto(m.Filter, &base, reg)
				if err != nil {
					return nil, fmt.Errorf("error getting filter for Aggregate Measure %d: %w", i, err)
				}
			}
		}

		out := &AggregateRel{
			input:               input,
			measures:            measures,
			groupingReferences:  groupingReferences,
			groupingExpressions: groupingExpressions,
			advExtension:        rel.Aggregate.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Aggregate.Common)
		return out, nil
	case *proto.Rel_Sort:
		input, err := RelFromProto(rel.Sort.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to SortRel: %w", err)
		}

		base := input.RecordType()
		sorts := make([]expr.SortField, len(rel.Sort.Sorts))
		for i, s := range rel.Sort.Sorts {
			sorts[i], err = expr.SortFieldFromProto(s, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting SortField %d for SortRel: %w", i, err)
			}
		}

		if len(sorts) == 0 {
			return nil, fmt.Errorf("%w: missing required field Sorts for Sort Relation", substraitgo.ErrInvalidRel)
		}

		out := &SortRel{
			input:        input,
			sorts:        sorts,
			advExtension: rel.Sort.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Sort.Common)
		return out, nil
	case *proto.Rel_Join:
		if rel.Join.Type == JoinTypeUnspecified {
			return nil, fmt.Errorf("%w: JoinRel must not have unspecified join type", substraitgo.ErrInvalidRel)
		}

		left, err := RelFromProto(rel.Join.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to JoinRel: %w", err)
		}

		right, err := RelFromProto(rel.Join.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to JoinRel: %w", err)
		}

		out := &JoinRel{
			left:         left,
			right:        right,
			joinType:     rel.Join.Type,
			advExtension: rel.Join.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Join.Common)

		base := out.JoinedRecordType()
		out.expr, err = expr.ExprFromProto(rel.Join.Expression, &base, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting expr for JoinRel: %w", err)
		}

		if rel.Join.PostJoinFilter != nil {
			out.postJoinFilter, err = expr.ExprFromProto(rel.Join.PostJoinFilter, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error parsing PostJoinFilter for JoinRel: %w", err)
			}
		}

		return out, nil
	case *proto.Rel_Project:
		input, err := RelFromProto(rel.Project.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to ProjectRel: %w", err)
		}

		baseSchema := input.RecordType()

		exprs := make([]expr.Expression, len(rel.Project.Expressions))
		for i, e := range rel.Project.Expressions {
			exprs[i], err = expr.ExprFromProto(e, &baseSchema, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting expr %d for ProjectRel: %w", i, err)
			}
		}

		if len(exprs) == 0 {
			return nil, fmt.Errorf("%w: missing required Expressions field for Project relation", substraitgo.ErrInvalidRel)
		}

		out := &ProjectRel{
			input:        input,
			exprs:        exprs,
			advExtension: rel.Project.AdvancedExtension,
		}
		if rel.Project.Common != nil {
			out.fromProtoCommon(rel.Project.Common)
		}
		return out, nil
	case *proto.Rel_Set:
		inputs := make([]Rel, len(rel.Set.Inputs))
		if len(inputs) < 2 {
			return nil, fmt.Errorf("%w: SetRel must have at least 2 inputs, only found %d",
				substraitgo.ErrInvalidRel, len(inputs))
		}

		var err error
		for i, r := range rel.Set.Inputs {
			inputs[i], err = RelFromProto(r, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting input %d for SetRel: %w", i, err)
			}
		}

		if rel.Set.Op == SetOpUnspecified {
			return nil, fmt.Errorf("%w: set operation must not be unspecified", substraitgo.ErrInvalidRel)
		}

		primary := inputs[0].RecordType()
		for i, in := range inputs[1:] {
			t := in.RecordType()
			if !t.Equals(&primary) {
				return nil, fmt.Errorf("%w: set operation field mismatch found in input #%d, expected %s, got %s",
					substraitgo.ErrInvalidRel, i+1, &primary, &t)
			}
		}

		out := &SetRel{
			inputs:       inputs,
			op:           rel.Set.Op,
			advExtension: rel.Set.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Set.Common)

		return out, nil
	case *proto.Rel_ExtensionSingle:
		input, err := RelFromProto(rel.ExtensionSingle.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to ExtensionSingle: %w", err)
		}

		out := &ExtensionSingleRel{
			input:  input,
			detail: rel.ExtensionSingle.Detail,
		}
		out.fromProtoCommon(rel.ExtensionSingle.Common)

		return out, nil
	case *proto.Rel_ExtensionMulti:
		inputs := make([]Rel, len(rel.ExtensionMulti.Inputs))
		var err error
		for i, r := range rel.ExtensionMulti.Inputs {
			inputs[i], err = RelFromProto(r, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting input %d for ExtensionMultiRel: %w", i, err)
			}
		}

		out := &ExtensionMultiRel{
			inputs: inputs,
			detail: rel.ExtensionMulti.Detail,
		}
		out.fromProtoCommon(rel.ExtensionMulti.Common)

		return out, nil
	case *proto.Rel_ExtensionLeaf:
		out := &ExtensionLeafRel{
			detail: rel.ExtensionLeaf.Detail,
		}
		out.fromProtoCommon(rel.ExtensionLeaf.Common)

		return out, nil
	case *proto.Rel_Cross:
		left, err := RelFromProto(rel.Cross.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to CrossRel: %w", err)
		}

		right, err := RelFromProto(rel.Cross.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to CrossRel: %w", err)
		}

		out := &CrossRel{
			left:         left,
			right:        right,
			advExtension: rel.Cross.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Cross.Common)
		return out, nil
	case *proto.Rel_HashJoin:
		left, err := RelFromProto(rel.HashJoin.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to HashJoinRel: %w", err)
		}

		right, err := RelFromProto(rel.HashJoin.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to HashJoin: %w", err)
		}

		if len(rel.HashJoin.LeftKeys) != len(rel.HashJoin.RightKeys) {
			return nil, fmt.Errorf("%w: mismatched number of keys for hash join. Left: %d, Right: %d",
				substraitgo.ErrInvalidRel, len(rel.HashJoin.LeftKeys), len(rel.HashJoin.RightKeys))
		}

		leftBase, rightBase := left.RecordType(), right.RecordType()

		leftKeys := make([]*expr.FieldReference, len(rel.HashJoin.LeftKeys))
		for i, k := range rel.HashJoin.LeftKeys {
			leftKeys[i], err = expr.FieldReferenceFromProto(k, &leftBase, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting left key %d for HashJoinRel: %w", i, err)
			}
		}

		rightKeys := make([]*expr.FieldReference, len(rel.HashJoin.RightKeys))
		for i, k := range rel.HashJoin.RightKeys {
			rightKeys[i], err = expr.FieldReferenceFromProto(k, &rightBase, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting right key %d for HashJoinRel: %w", i, err)
			}
		}

		if len(leftKeys) != len(rightKeys) {
			return nil, fmt.Errorf("%w: must have same number of keys in left and right keys for hash join", substraitgo.ErrInvalidRel)
		}

		out := &HashJoinRel{
			left:         left,
			right:        right,
			leftKeys:     leftKeys,
			rightKeys:    rightKeys,
			joinType:     HashMergeJoinType(rel.HashJoin.Type),
			advExtension: rel.HashJoin.AdvancedExtension,
		}
		out.fromProtoCommon(rel.HashJoin.Common)

		if rel.HashJoin.PostJoinFilter != nil {
			base := out.RecordType()
			out.postJoinFilter, err = expr.ExprFromProto(rel.HashJoin.PostJoinFilter, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting post join filter for HashJoinRel: %w", err)
			}
		}

		return out, nil
	case *proto.Rel_MergeJoin:
		left, err := RelFromProto(rel.MergeJoin.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to MergeJoinRel: %w", err)
		}

		right, err := RelFromProto(rel.MergeJoin.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to HashJoin: %w", err)
		}

		if len(rel.MergeJoin.LeftKeys) != len(rel.MergeJoin.RightKeys) {
			return nil, fmt.Errorf("%w: mismatched number of keys for merge join. Left: %d, Right: %d",
				substraitgo.ErrInvalidRel, len(rel.MergeJoin.LeftKeys), len(rel.MergeJoin.RightKeys))
		}

		leftBase, rightBase := left.RecordType(), right.RecordType()

		leftKeys := make([]*expr.FieldReference, len(rel.MergeJoin.LeftKeys))
		for i, k := range rel.MergeJoin.LeftKeys {
			leftKeys[i], err = expr.FieldReferenceFromProto(k, &leftBase, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting left key %d for MergeJoin: %w", i, err)
			}
		}

		rightKeys := make([]*expr.FieldReference, len(rel.MergeJoin.RightKeys))
		for i, k := range rel.MergeJoin.RightKeys {
			rightKeys[i], err = expr.FieldReferenceFromProto(k, &rightBase, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting right key %d for MergeJoin: %w", i, err)
			}
		}

		if len(leftKeys) != len(rightKeys) {
			return nil, fmt.Errorf("%w: must have same number of keys in left and right keys for merge join", substraitgo.ErrInvalidRel)
		}

		out := &HashJoinRel{
			left:         left,
			right:        right,
			leftKeys:     leftKeys,
			rightKeys:    rightKeys,
			joinType:     HashMergeJoinType(rel.MergeJoin.Type),
			advExtension: rel.MergeJoin.AdvancedExtension,
		}
		out.fromProtoCommon(rel.MergeJoin.Common)

		if rel.MergeJoin.PostJoinFilter != nil {
			base := out.RecordType()
			out.postJoinFilter, err = expr.ExprFromProto(rel.MergeJoin.PostJoinFilter, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting post join filter for MergeJoin: %w", err)
			}
		}

		return out, nil
	case *proto.Rel_Write:
		input, err := RelFromProto(rel.Write.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to WriteRel: %w", err)
		}
		tableSchema := types.NewNamedStructFromProto(rel.Write.TableSchema)
		out := &NamedTableWriteRel{
			tableSchema: tableSchema,
			op:          rel.Write.Op,
			input:       input,
			outputMode:  rel.Write.Output,
		}
		if rel.Write.Common != nil {
			out.fromProtoCommon(rel.Write.Common)
		}
		switch rel.Write.Op {
		case proto.WriteRel_WRITE_OP_CTAS, proto.WriteRel_WRITE_OP_INSERT, proto.WriteRel_WRITE_OP_DELETE:
			switch writeType := rel.Write.WriteType.(type) {
			case *proto.WriteRel_NamedTable:
				out.names = writeType.NamedTable.Names
				out.advExtension = writeType.NamedTable.AdvancedExtension
			case *proto.WriteRel_ExtensionTable:
				return nil, fmt.Errorf("%w: ExtensionTable not supported for WriteRel", substraitgo.ErrInvalidRel)
			}
		default:
			return nil, fmt.Errorf("%w: WriteRel not supported for optype %v", substraitgo.ErrInvalidRel, rel.Write.Op)
		}
		return out, nil
	case nil:
		return nil, fmt.Errorf("%w: got nil", substraitgo.ErrInvalidRel)
	}

	return nil, substraitgo.ErrNotImplemented
}
