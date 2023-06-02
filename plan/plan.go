// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

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

type Plan struct {
	version          *types.Version
	extensions       extensions.Set
	expectedTypeURLs []string
	advExtension     *extensions.AdvancedExtension
	relations        []Relation

	reg expr.ExtensionRegistry
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

type Root struct {
	input Rel
	names []string
}

func (r *Root) Input() Rel      { return r.input }
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

type Rel interface {
	Hint() *Hint
	Remap(types.StructType) types.StructType
	OutputMapping() []int32
	GetAdvancedExtension() *extensions.AdvancedExtension

	ToProto() *proto.Rel
	ToProtoPlanRel() *proto.PlanRel

	RecordType() types.StructType
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
			values := make([]expr.StructLiteralValue, len(readType.VirtualTable.Values))
			for i, v := range readType.VirtualTable.Values {
				values[i] = expr.StructLiteralFromProto(v)
			}

			out = &VirtualTableReadRel{
				values: values,
			}
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
		out.fromProtoCommon(rel.Filter.Common)

		return out, nil
	case *proto.Rel_Fetch:
		input, err := RelFromProto(rel.Fetch.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to FetchRel: %w", err)
		}

		out := &FetchRel{
			input:        input,
			offset:       rel.Fetch.Offset,
			count:        rel.Fetch.Count,
			advExtension: rel.Fetch.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Fetch.Common)
		return out, nil
	case *proto.Rel_Aggregate:
		input, err := RelFromProto(rel.Aggregate.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to AggregateRel: %w", err)
		}

		base := input.RecordType()
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
			input:        input,
			groups:       groups,
			measures:     measures,
			advExtension: rel.Aggregate.AdvancedExtension,
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

		out := &SortRel{
			input:        input,
			sorts:        sorts,
			advExtension: rel.Sort.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Sort.Common)
		return out, nil
	case *proto.Rel_Join:
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

		base := out.RecordType()
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
		out := &ProjectRel{
			input:        input,
			exprs:        exprs,
			advExtension: rel.Project.AdvancedExtension,
		}
		out.fromProtoCommon(rel.Project.Common)
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
	case nil:
		return nil, fmt.Errorf("%w: got nil", substraitgo.ErrInvalidRel)
	}

	return nil, substraitgo.ErrNotImplemented
}
