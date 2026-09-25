// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/types/known/anypb"
)

// RelToProto encodes a relation as its protobuf message.
func RelToProto(rel plan.Rel) *proto.Rel {
	switch r := rel.(type) {
	case *plan.NamedTableReadRel:
		return namedTableReadRelToProto(r)
	case *plan.VirtualTableReadRel:
		return virtualTableReadRelToProto(r)
	case *plan.ExtensionTableReadRel:
		return extensionTableReadRelToProto(r)
	case *plan.IcebergTableReadRel:
		return icebergTableReadRelToProto(r)
	case *plan.LocalFileReadRel:
		return localFileReadRelToProto(r)
	case *plan.FilterRel:
		return filterRelToProto(r)
	case *plan.FetchRel:
		return fetchRelToProto(r)
	case *plan.ProjectRel:
		return projectRelToProto(r)
	case *plan.AggregateRel:
		return aggregateRelToProto(r)
	case *plan.SortRel:
		return sortRelToProto(r)
	case *plan.SetRel:
		return setRelToProto(r)
	case *plan.CrossRel:
		return crossRelToProto(r)
	case *plan.JoinRel:
		return joinRelToProto(r)
	case *plan.HashJoinRel:
		return hashJoinRelToProto(r)
	case *plan.MergeJoinRel:
		return mergeJoinRelToProto(r)
	case *plan.NamedTableWriteRel:
		return namedTableWriteRelToProto(r)
	case *plan.ExtensionSingleRel:
		return extensionSingleRelToProto(r)
	default:
		panic(fmt.Sprintf("wire: unhandled relation %T", rel))
	}
}

type readRelReader interface {
	BaseSchema() types.NamedStruct
	Filter() expr.Expression
	BestEffortFilter() expr.Expression
	Projection() *expr.MaskExpression
}

func baseReadRelToProto(rc *plan.RelCommon, advExt *extensions.AdvancedExtension, r readRelReader) *proto.ReadRel {
	out := &proto.ReadRel{
		Common:            relCommonToProto(rc),
		BaseSchema:        NamedStructToProto(r.BaseSchema()),
		AdvancedExtension: advancedExtensionToProto(advExt),
	}
	if f := r.Filter(); f != nil {
		out.Filter = ExprToProto(f)
	}
	if f := r.BestEffortFilter(); f != nil {
		out.BestEffortFilter = ExprToProto(f)
	}
	if p := r.Projection(); p != nil {
		out.Projection = MaskExpressionToProto(p)
	}
	return out
}

func namedTableReadRelToProto(n *plan.NamedTableReadRel) *proto.Rel {
	readRel := baseReadRelToProto(&n.RelCommon, n.GetAdvancedExtension(), n)
	readRel.ReadType = &proto.ReadRel_NamedTable_{
		NamedTable: &proto.ReadRel_NamedTable{
			Names:             n.Names(),
			AdvancedExtension: advancedExtensionToProto(n.NamedTableAdvancedExtension()),
		},
	}
	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

func virtualTableReadRelToProto(v *plan.VirtualTableReadRel) *proto.Rel {
	readRel := baseReadRelToProto(&v.RelCommon, v.GetAdvancedExtension(), v)
	values := make([]*proto.Expression_Nested_Struct, len(v.Values()))
	for i, val := range v.Values() {
		values[i] = VirtualTableExpressionValueToProto(val)
	}
	readRel.ReadType = &proto.ReadRel_VirtualTable_{
		VirtualTable: &proto.ReadRel_VirtualTable{Expressions: values},
	}
	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

func extensionTableReadRelToProto(e *plan.ExtensionTableReadRel) *proto.Rel {
	readRel := baseReadRelToProto(&e.RelCommon, e.GetAdvancedExtension(), e)
	readRel.ReadType = &proto.ReadRel_ExtensionTable_{
		ExtensionTable: &proto.ReadRel_ExtensionTable{Detail: e.Detail()},
	}
	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

func icebergTableReadRelToProto(n *plan.IcebergTableReadRel) *proto.Rel {
	readRel := baseReadRelToProto(&n.RelCommon, n.GetAdvancedExtension(), n)

	if directTableType, ok := n.TableType().(*plan.Direct); ok {
		direct := &proto.ReadRel_IcebergTable_MetadataFileRead{
			MetadataUri: directTableType.MetadataUri,
		}
		if directTableType.SnapshotId != "" {
			direct.Snapshot = &proto.ReadRel_IcebergTable_MetadataFileRead_SnapshotId{
				SnapshotId: string(directTableType.SnapshotId),
			}
		} else if directTableType.SnapshotTimestamp != 0 {
			direct.Snapshot = &proto.ReadRel_IcebergTable_MetadataFileRead_SnapshotTimestamp{
				SnapshotTimestamp: int64(directTableType.SnapshotTimestamp),
			}
		}
		readRel.ReadType = &proto.ReadRel_IcebergTable_{
			IcebergTable: &proto.ReadRel_IcebergTable{
				TableType: &proto.ReadRel_IcebergTable_Direct{Direct: direct},
			},
		}
	}

	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

func localFileReadRelToProto(lf *plan.LocalFileReadRel) *proto.Rel {
	items := make([]*proto.ReadRel_LocalFiles_FileOrFiles, len(lf.Items()))
	for i := range lf.Items() {
		item := lf.Item(i)
		items[i] = fileOrFilesToProto(&item)
	}

	readRel := baseReadRelToProto(&lf.RelCommon, lf.ReadRelAdvancedExtension(), lf)
	readRel.ReadType = &proto.ReadRel_LocalFiles_{
		LocalFiles: &proto.ReadRel_LocalFiles{
			Items:             items,
			AdvancedExtension: advancedExtensionToProto(lf.GetAdvancedExtension()),
		},
	}
	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

func fileOrFilesToProto(f *plan.FileOrFiles) *proto.ReadRel_LocalFiles_FileOrFiles {
	ret := &proto.ReadRel_LocalFiles_FileOrFiles{
		PartitionIndex: f.PartIndex,
		Start:          f.Start,
		Length:         f.Len,
	}
	switch f.PathType {
	case plan.URIPath:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriPath{UriPath: f.Path}
	case plan.URIPathGlob:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriPathGlob{UriPathGlob: f.Path}
	case plan.URIFile:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriFile{UriFile: f.Path}
	case plan.URIFolder:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriFolder{UriFolder: f.Path}
	}

	switch fm := f.Format.(type) {
	case *plan.ParquetReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Parquet{
			Parquet: &proto.ReadRel_LocalFiles_FileOrFiles_ParquetReadOptions{},
		}
	case *plan.ArrowReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Arrow{
			Arrow: &proto.ReadRel_LocalFiles_FileOrFiles_ArrowReadOptions{},
		}
	case *plan.OrcReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Orc{
			Orc: &proto.ReadRel_LocalFiles_FileOrFiles_OrcReadOptions{},
		}
	case *plan.DwrfReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Dwrf{
			Dwrf: &proto.ReadRel_LocalFiles_FileOrFiles_DwrfReadOptions{},
		}
	case *plan.ExtensionReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Extension{
			Extension: (*anypb.Any)(fm),
		}
	}
	return ret
}

func filterRelToProto(fr *plan.FilterRel) *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Filter{
			Filter: &proto.FilterRel{
				Common:            relCommonToProto(&fr.RelCommon),
				Input:             RelToProto(fr.Input()),
				Condition:         ExprToProto(fr.Condition()),
				AdvancedExtension: advancedExtensionToProto(fr.GetAdvancedExtension()),
			},
		},
	}
}

func fetchRelToProto(f *plan.FetchRel) *proto.Rel {
	fetchRel := &proto.FetchRel{
		Common:            relCommonToProto(&f.RelCommon),
		Input:             RelToProto(f.Input()),
		AdvancedExtension: advancedExtensionToProto(f.GetAdvancedExtension()),
	}
	if f.Offset() != nil {
		fetchRel.OffsetMode = &proto.FetchRel_OffsetExpr{OffsetExpr: ExprToProto(f.Offset())}
	}
	if f.Count() != nil {
		fetchRel.CountMode = &proto.FetchRel_CountExpr{CountExpr: ExprToProto(f.Count())}
	}
	return &proto.Rel{RelType: &proto.Rel_Fetch{Fetch: fetchRel}}
}

func projectRelToProto(p *plan.ProjectRel) *proto.Rel {
	exprs := make([]*proto.Expression, len(p.Expressions()))
	for i, e := range p.Expressions() {
		exprs[i] = ExprToProto(e)
	}
	return &proto.Rel{
		RelType: &proto.Rel_Project{
			Project: &proto.ProjectRel{
				Common:            relCommonToProto(&p.RelCommon),
				Input:             RelToProto(p.Input()),
				Expressions:       exprs,
				AdvancedExtension: advancedExtensionToProto(p.GetAdvancedExtension()),
			},
		},
	}
}

func aggregateRelToProto(ar *plan.AggregateRel) *proto.Rel {
	groupingExprs := make([]*proto.Expression, len(ar.GroupingExpressions()))
	for i, e := range ar.GroupingExpressions() {
		groupingExprs[i] = ExprToProto(e)
	}

	refs := ar.GroupingReferences()
	groupings := make([]*proto.AggregateRel_Grouping, len(refs))
	for i := range refs {
		groupings[i] = &proto.AggregateRel_Grouping{ExpressionReferences: refs[i]}
	}

	measures := make([]*proto.AggregateRel_Measure, len(ar.Measures()))
	for i := range ar.Measures() {
		m := ar.Measures()[i]
		measures[i] = aggRelMeasureToProto(&m)
	}

	return &proto.Rel{
		RelType: &proto.Rel_Aggregate{
			Aggregate: &proto.AggregateRel{
				Common:              relCommonToProto(&ar.RelCommon),
				Input:               RelToProto(ar.Input()),
				GroupingExpressions: groupingExprs,
				Groupings:           groupings,
				Measures:            measures,
				AdvancedExtension:   advancedExtensionToProto(ar.GetAdvancedExtension()),
			},
		},
	}
}

func groupingExprs(groups [][]expr.Expression) ([]expr.Expression, [][]uint32) {
	groupingExpressions := make([]expr.Expression, 0)
	groupingReferences := make([][]uint32, 0)
	for _, group := range groups {
		refs := make([]uint32, 0)
		for _, e := range group {
			existingExpr := false
			for eIndex, existing := range groupingExpressions {
				if existing.Equals(e) {
					existingExpr = true
					refs = append(refs, uint32(eIndex))
					break
				}
			}
			if !existingExpr {
				groupingExpressions = append(groupingExpressions, e)
				refs = append(refs, uint32(len(groupingExpressions)-1))
			}
		}
		groupingReferences = append(groupingReferences, refs)
	}
	return groupingExpressions, groupingReferences
}

func aggRelMeasureToProto(am *plan.AggRelMeasure) *proto.AggregateRel_Measure {
	ret := &proto.AggregateRel_Measure{
		Measure: AggregateFunctionToProto(am.Measure()),
	}
	if f := am.RawFilter(); f != nil {
		ret.Filter = ExprToProto(f)
	}
	return ret
}

func sortRelToProto(sr *plan.SortRel) *proto.Rel {
	sorts := make([]*proto.SortField, len(sr.Sorts()))
	for i := range sr.Sorts() {
		s := sr.Sorts()[i]
		sorts[i] = SortFieldToProto(&s)
	}
	return &proto.Rel{
		RelType: &proto.Rel_Sort{
			Sort: &proto.SortRel{
				Common:            relCommonToProto(&sr.RelCommon),
				Input:             RelToProto(sr.Input()),
				Sorts:             sorts,
				AdvancedExtension: advancedExtensionToProto(sr.GetAdvancedExtension()),
			},
		},
	}
}

func setRelToProto(s *plan.SetRel) *proto.Rel {
	inputs := make([]*proto.Rel, len(s.Inputs()))
	for i, in := range s.Inputs() {
		inputs[i] = RelToProto(in)
	}
	return &proto.Rel{
		RelType: &proto.Rel_Set{
			Set: &proto.SetRel{
				Common:            relCommonToProto(&s.RelCommon),
				Inputs:            inputs,
				Op:                proto.SetRel_SetOp(s.Op()),
				AdvancedExtension: advancedExtensionToProto(s.GetAdvancedExtension()),
			},
		},
	}
}

func crossRelToProto(c *plan.CrossRel) *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Cross{
			Cross: &proto.CrossRel{
				Common:            relCommonToProto(&c.RelCommon),
				Left:              RelToProto(c.Left()),
				Right:             RelToProto(c.Right()),
				AdvancedExtension: advancedExtensionToProto(c.GetAdvancedExtension()),
			},
		},
	}
}

func joinRelToProto(j *plan.JoinRel) *proto.Rel {
	outRel := &proto.JoinRel{
		Common:            relCommonToProto(&j.RelCommon),
		Left:              RelToProto(j.Left()),
		Right:             RelToProto(j.Right()),
		Expression:        ExprToProto(j.Expr()),
		Type:              proto.JoinRel_JoinType(j.Type()),
		AdvancedExtension: advancedExtensionToProto(j.GetAdvancedExtension()),
	}
	if f := j.RawPostJoinFilter(); f != nil {
		outRel.PostJoinFilter = ExprToProto(f)
	}
	return &proto.Rel{RelType: &proto.Rel_Join{Join: outRel}}
}

func hashJoinRelToProto(hr *plan.HashJoinRel) *proto.Rel {
	ret := &proto.HashJoinRel{
		Common:            relCommonToProto(&hr.RelCommon),
		Left:              RelToProto(hr.Left()),
		Right:             RelToProto(hr.Right()),
		Keys:              comparisonJoinKeysToProto(hr.Keys()),
		Type:              proto.HashJoinRel_JoinType(hr.Type()),
		AdvancedExtension: advancedExtensionToProto(hr.GetAdvancedExtension()),
	}
	if leftKeys, rightKeys, ok := tryEqualityJoinKeysToLegacyProto(hr.Keys()); ok {
		ret.LeftKeys = leftKeys
		ret.RightKeys = rightKeys
	}
	if f := hr.RawPostJoinFilter(); f != nil {
		ret.PostJoinFilter = ExprToProto(f)
	}
	return &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: ret}}
}

func mergeJoinRelToProto(mr *plan.MergeJoinRel) *proto.Rel {
	ret := &proto.MergeJoinRel{
		Common:            relCommonToProto(&mr.RelCommon),
		Left:              RelToProto(mr.Left()),
		Right:             RelToProto(mr.Right()),
		Keys:              comparisonJoinKeysToProto(mr.Keys()),
		Type:              proto.MergeJoinRel_JoinType(mr.Type()),
		AdvancedExtension: advancedExtensionToProto(mr.GetAdvancedExtension()),
	}
	if leftKeys, rightKeys, ok := tryEqualityJoinKeysToLegacyProto(mr.Keys()); ok {
		ret.LeftKeys = leftKeys
		ret.RightKeys = rightKeys
	}
	if f := mr.RawPostJoinFilter(); f != nil {
		ret.PostJoinFilter = ExprToProto(f)
	}
	return &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: ret}}
}

func comparisonJoinKeysToProto(keys []*plan.ComparisonJoinKey) []*proto.ComparisonJoinKey {
	out := make([]*proto.ComparisonJoinKey, len(keys))
	for i, k := range keys {
		out[i] = comparisonJoinKeyToProto(k)
	}
	return out
}

func comparisonJoinKeyToProto(k *plan.ComparisonJoinKey) *proto.ComparisonJoinKey {
	return &proto.ComparisonJoinKey{
		Left:       fieldReferenceRefToProto(k.Left()),
		Right:      fieldReferenceRefToProto(k.Right()),
		Comparison: joinKeyComparisonToProto(k.Comparison()),
	}
}

func joinKeyComparisonToProto(c plan.JoinKeyComparison) *proto.ComparisonJoinKey_ComparisonType {
	switch c := c.(type) {
	case plan.SimpleComparison:
		return simpleComparisonToProto(c)
	case *plan.SimpleComparison:
		return simpleComparisonToProto(*c)
	case plan.CustomComparison:
		return customComparisonToProto(c)
	case *plan.CustomComparison:
		return customComparisonToProto(*c)
	}
	return nil
}

func simpleComparisonToProto(c plan.SimpleComparison) *proto.ComparisonJoinKey_ComparisonType {
	return &proto.ComparisonJoinKey_ComparisonType{
		InnerType: &proto.ComparisonJoinKey_ComparisonType_Simple{
			Simple: proto.ComparisonJoinKey_SimpleComparisonType(c.Type)},
	}
}

func customComparisonToProto(c plan.CustomComparison) *proto.ComparisonJoinKey_ComparisonType {
	return &proto.ComparisonJoinKey_ComparisonType{
		InnerType: &proto.ComparisonJoinKey_ComparisonType_CustomFunctionReference{
			CustomFunctionReference: c.FunctionReference},
	}
}

// tryEqualityJoinKeysToLegacyProto returns the deprecated left_keys/right_keys
// representation of the given join keys with ok=true, but only when every key
// is a plain SIMPLE_COMPARISON_TYPE_EQ comparison. Those are the only joins the
// deprecated fields can express; IS_NOT_DISTINCT_FROM, MIGHT_EQUAL and custom
// comparisons have no legacy encoding and an old consumer would silently treat
// them as equality, so for those it returns ok=false and the caller should emit
// only the modern keys field.
func tryEqualityJoinKeysToLegacyProto(keys []*plan.ComparisonJoinKey) (leftKeys, rightKeys []*proto.Expression_FieldReference, ok bool) {
	for _, k := range keys {
		switch simple := k.Comparison().(type) {
		case plan.SimpleComparison:
			if simple.Type != plan.SimpleComparisonTypeEq {
				return nil, nil, false
			}
		case *plan.SimpleComparison:
			if simple == nil || simple.Type != plan.SimpleComparisonTypeEq {
				return nil, nil, false
			}
		default:
			return nil, nil, false
		}
	}
	leftKeys = make([]*proto.Expression_FieldReference, len(keys))
	rightKeys = make([]*proto.Expression_FieldReference, len(keys))
	for i, k := range keys {
		leftKeys[i] = fieldReferenceRefToProto(k.Left())
		rightKeys[i] = fieldReferenceRefToProto(k.Right())
	}
	return leftKeys, rightKeys, true
}

func namedTableWriteRelToProto(wr *plan.NamedTableWriteRel) *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Write{
			Write: &proto.WriteRel{
				Common: relCommonToProto(&wr.RelCommon),
				WriteType: &proto.WriteRel_NamedTable{
					NamedTable: &proto.NamedObjectWrite{
						Names:             wr.Names(),
						AdvancedExtension: advancedExtensionToProto(wr.NamedTableAdvancedExtension()),
					},
				},
				TableSchema: NamedStructToProto(wr.TableSchema()),
				Op:          proto.WriteRel_WriteOp(wr.Op()),
				Input:       RelToProto(wr.Input()),
			},
		},
	}
}

func extensionSingleRelToProto(es *plan.ExtensionSingleRel) *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_ExtensionSingle{
			ExtensionSingle: &proto.ExtensionSingleRel{
				Common: relCommonToProto(&es.RelCommon),
				Input:  RelToProto(es.Input()),
				Detail: es.Detail(),
			},
		},
	}
}

// relCommonFromProto decodes the common fields shared by every relation.
func relCommonFromProto(c *proto.RelCommon) plan.RelCommon {
	if c == nil {
		return plan.NewRelCommon(nil, nil, nil)
	}
	var mapping []int32
	if emit, ok := c.GetEmitKind().(*proto.RelCommon_Emit_); ok {
		mapping = emit.Emit.OutputMapping
	}
	return plan.NewRelCommon(hintFromProto(c.Hint), mapping, advancedExtensionFromProto(c.AdvancedExtension))
}

// decodedReadRelBase carries the fields common to all read relations while a
// specific read relation is being decoded.
type decodedReadRelBase struct {
	common           plan.RelCommon
	baseSchema       types.NamedStruct
	filter           expr.Expression
	bestEffortFilter expr.Expression
	projection       *expr.MaskExpression
	advExtension     *extensions.AdvancedExtension
}

func readRelBaseFromProto(rel *proto.ReadRel, reg expr.ExtensionRegistry) (decodedReadRelBase, error) {
	var b decodedReadRelBase
	if rel.Common != nil {
		b.common = relCommonFromProto(rel.Common)
	}

	b.baseSchema = NamedStructFromProto(rel.BaseSchema)
	var err error
	if rel.Filter != nil {
		b.filter, err = ExprFromProto(rel.Filter, types.NewRecordTypeFromStruct(b.baseSchema.Struct), reg)
		if err != nil {
			return b, err
		}
	}

	if rel.BestEffortFilter != nil {
		b.bestEffortFilter, err = ExprFromProto(rel.BestEffortFilter, types.NewRecordTypeFromStruct(b.baseSchema.Struct), reg)
		if err != nil {
			return b, err
		}
	}

	if rel.Projection != nil {
		b.projection = MaskExpressionFromProto(rel.Projection)
	}

	b.advExtension = advancedExtensionFromProto(rel.AdvancedExtension)
	return b, nil
}

// fileOrFilesFromProto decodes a single local-file item.
func fileOrFilesFromProto(p *proto.ReadRel_LocalFiles_FileOrFiles) plan.FileOrFiles {
	var f plan.FileOrFiles
	f.PartIndex = p.PartitionIndex
	f.Start, f.Len = p.Start, p.Length

	switch path := p.PathType.(type) {
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriFile:
		f.PathType, f.Path = plan.URIFile, path.UriFile
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriFolder:
		f.PathType, f.Path = plan.URIFolder, path.UriFolder
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriPath:
		f.PathType, f.Path = plan.URIPath, path.UriPath
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriPathGlob:
		f.PathType, f.Path = plan.URIPathGlob, path.UriPathGlob
	}

	switch format := p.FileFormat.(type) {
	case *proto.ReadRel_LocalFiles_FileOrFiles_Arrow:
		f.Format = &plan.ArrowReadOptions{}
	case *proto.ReadRel_LocalFiles_FileOrFiles_Dwrf:
		f.Format = &plan.DwrfReadOptions{}
	case *proto.ReadRel_LocalFiles_FileOrFiles_Extension:
		f.Format = (*plan.ExtensionReadOptions)(format.Extension)
	case *proto.ReadRel_LocalFiles_FileOrFiles_Orc:
		f.Format = &plan.OrcReadOptions{}
	case *proto.ReadRel_LocalFiles_FileOrFiles_Parquet:
		f.Format = &plan.ParquetReadOptions{}
	}
	return f
}

// virtualTableExpressionFromProto decodes an expression-valued virtual table row.
func virtualTableExpressionFromProto(s *proto.Expression_Nested_Struct, reg expr.ExtensionRegistry) (expr.VirtualTableExpressionValue, error) {
	fields := make(expr.VirtualTableExpressionValue, len(s.Fields))
	for i, f := range s.Fields {
		val, err := ExprFromProto(f, nil, reg)
		if err != nil {
			return nil, err
		}
		fields[i] = val
	}
	return fields, nil
}

// virtualTableExprFromLiteralProto decodes a literal-valued virtual table row.
func virtualTableExprFromLiteralProto(s *proto.Expression_Literal_Struct) expr.VirtualTableExpressionValue {
	fields := make(expr.VirtualTableExpressionValue, len(s.Fields))
	for i, f := range s.Fields {
		fields[i] = LiteralFromProto(f)
	}
	return fields
}

func joinKeyComparisonFromProto(c *proto.ComparisonJoinKey_ComparisonType) (plan.JoinKeyComparison, error) {
	switch it := c.GetInnerType().(type) {
	case *proto.ComparisonJoinKey_ComparisonType_Simple:
		return plan.SimpleComparison{Type: plan.SimpleComparisonType(it.Simple)}, nil
	case *proto.ComparisonJoinKey_ComparisonType_CustomFunctionReference:
		return plan.CustomComparison{FunctionReference: it.CustomFunctionReference}, nil
	default:
		return nil, fmt.Errorf("%w: unsupported join key comparison type %T", substraitgo.ErrInvalidRel, it)
	}
}

// comparisonJoinKeysFromProto builds the join keys for a hash/merge join,
// preferring the keys field. The deprecated leftKeys/rightKeys are only used
// when keys is empty, in which case they are paired with an EQ comparison.
func comparisonJoinKeysFromProto(
	keys []*proto.ComparisonJoinKey,
	leftKeys, rightKeys []*proto.Expression_FieldReference,
	leftSchema, rightSchema *types.RecordType,
	reg expr.ExtensionRegistry,
) ([]*plan.ComparisonJoinKey, error) {
	if len(keys) > 0 {
		out := make([]*plan.ComparisonJoinKey, len(keys))
		for i, k := range keys {
			left, err := FieldReferenceFromProto(k.GetLeft(), leftSchema, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting left key %d for join: %w", i, err)
			}
			right, err := FieldReferenceFromProto(k.GetRight(), rightSchema, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting right key %d for join: %w", i, err)
			}
			comparison, err := joinKeyComparisonFromProto(k.GetComparison())
			if err != nil {
				return nil, err
			}
			out[i] = plan.NewComparisonJoinKey(left, right, comparison)
		}
		return out, nil
	}

	if len(leftKeys) != len(rightKeys) {
		return nil, fmt.Errorf("%w: mismatched number of keys for join. Left: %d, Right: %d",
			substraitgo.ErrInvalidRel, len(leftKeys), len(rightKeys))
	}
	out := make([]*plan.ComparisonJoinKey, len(leftKeys))
	for i := range leftKeys {
		left, err := FieldReferenceFromProto(leftKeys[i], leftSchema, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left key %d for join: %w", i, err)
		}
		right, err := FieldReferenceFromProto(rightKeys[i], rightSchema, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right key %d for join: %w", i, err)
		}
		out[i] = plan.NewEqualityJoinKey(left, right)
	}
	return out, nil
}

func decodeExtensionDef(reg expr.ExtensionRegistry, detail *anypb.Any) (plan.ExtensionRelDefinition, error) {
	if dec := reg.ExtensionRelDecoderFor(detail.GetTypeUrl()); dec != nil {
		raw, err := dec.DecodeExtensionRel(detail)
		if err != nil {
			return nil, err
		}
		def, ok := raw.(plan.ExtensionRelDefinition)
		if !ok {
			return nil, fmt.Errorf("ExtensionRelDecoder returned %T which does not implement ExtensionRelDefinition", raw)
		}
		return def, nil
	}
	return plan.NewUndecodedExtension(detail), nil
}

// RelFromProto decodes a relation and all of its inputs from protobuf.
func RelFromProto(rel *proto.Rel, reg expr.ExtensionRegistry) (plan.Rel, error) {
	switch rel := rel.RelType.(type) {
	case *proto.Rel_Read:
		// Decode the read type first, then the shared base, so a malformed read
		// type is reported before a base field (Filter/BestEffortFilter) that
		// also fails to decode. build assembles the relation once the base is ready.
		var build func(base decodedReadRelBase) plan.Rel
		switch readType := rel.Read.ReadType.(type) {
		case *proto.ReadRel_ExtensionTable_:
			detail := readType.ExtensionTable.Detail
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewExtensionTableReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), detail)
			}
		case *proto.ReadRel_LocalFiles_:
			items := make([]plan.FileOrFiles, len(readType.LocalFiles.Items))
			for i, item := range readType.LocalFiles.Items {
				items[i] = fileOrFilesFromProto(item)
			}
			advExtension := advancedExtensionFromProto(readType.LocalFiles.AdvancedExtension)
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewLocalFileReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), items, advExtension)
			}
		case *proto.ReadRel_NamedTable_:
			names := readType.NamedTable.Names
			advExtension := advancedExtensionFromProto(readType.NamedTable.AdvancedExtension)
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewNamedTableReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), names, advExtension)
			}
		case *proto.ReadRel_VirtualTable_:
			if len(readType.VirtualTable.Values) > 0 && len(readType.VirtualTable.Expressions) > 0 {
				return nil, fmt.Errorf("VirtualTable cannot declare both Values and Expressions")
			}
			var values []expr.VirtualTableExpressionValue
			for _, v := range readType.VirtualTable.Values {
				values = append(values, virtualTableExprFromLiteralProto(v))
			}
			for _, v := range readType.VirtualTable.Expressions {
				row, err := virtualTableExpressionFromProto(v, reg)
				if err != nil {
					return nil, err
				}
				values = append(values, row)
			}
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewVirtualTableReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), values)
			}
		case *proto.ReadRel_IcebergTable_:
			icebergTableType := readType.IcebergTable.TableType
			if icebergTableType == nil {
				return nil, fmt.Errorf("%w: IcebergTableType is required for IcebergTableReadRel", substraitgo.ErrInvalidRel)
			}
			if _, ok := icebergTableType.(plan.IcebergTableType); ok {
				return nil, fmt.Errorf("%w: IcebergTableType must be a string", substraitgo.ErrInvalidRel)
			}
			direct, ok := icebergTableType.(*proto.ReadRel_IcebergTable_Direct)
			if !ok {
				return nil, fmt.Errorf("%w: only IcebergTableType Direct is supported", substraitgo.ErrInvalidRel)
			}
			tableType := &plan.Direct{
				MetadataUri: direct.Direct.MetadataUri,
			}
			if snapshotId, ok := direct.Direct.Snapshot.(*proto.ReadRel_IcebergTable_MetadataFileRead_SnapshotId); ok {
				tableType.SnapshotId = plan.SnapshotId(snapshotId.SnapshotId)
			} else if snapshotTimestamp, ok := direct.Direct.Snapshot.(*proto.ReadRel_IcebergTable_MetadataFileRead_SnapshotTimestamp); ok {
				tableType.SnapshotTimestamp = plan.SnapshotTimestamp(snapshotTimestamp.SnapshotTimestamp)
			}
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewIcebergTableReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), tableType)
			}
		default:
			return nil, fmt.Errorf("%w: unknown ReadRel type", substraitgo.ErrInvalidRel)
		}

		base, err := readRelBaseFromProto(rel.Read, reg)
		if err != nil {
			return nil, err
		}
		return build(base), nil
	case *proto.Rel_Filter:
		input, err := RelFromProto(rel.Filter.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to FilterRel: %w", err)
		}

		base := input.RecordType()
		cond, err := ExprFromProto(rel.Filter.Condition, &base, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting condition for FilterRel: %w", err)
		}

		var common plan.RelCommon
		if rel.Filter.Common != nil {
			common = relCommonFromProto(rel.Filter.Common)
		}
		return plan.NewFilterRel(input, cond, common, advancedExtensionFromProto(rel.Filter.AdvancedExtension)), nil
	case *proto.Rel_Fetch:
		input, err := RelFromProto(rel.Fetch.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to FetchRel: %w", err)
		}

		base := input.RecordType()

		var offset expr.Expression
		switch om := rel.Fetch.OffsetMode.(type) {
		case *proto.FetchRel_Offset:
			offset = expr.NewPrimitiveLiteral(om.Offset, false)
		case *proto.FetchRel_OffsetExpr:
			e, exprErr := ExprFromProto(om.OffsetExpr, &base, reg)
			if exprErr != nil {
				return nil, fmt.Errorf("error getting offset expression for FetchRel: %w", exprErr)
			}
			offset = e
		}

		var count expr.Expression
		switch cm := rel.Fetch.CountMode.(type) {
		case *proto.FetchRel_Count:
			if cm.Count != plan.FETCH_COUNT_ALL_RECORDS {
				count = expr.NewPrimitiveLiteral(cm.Count, false)
			}
		case *proto.FetchRel_CountExpr:
			e, exprErr := ExprFromProto(cm.CountExpr, &base, reg)
			if exprErr != nil {
				return nil, fmt.Errorf("error getting count expression for FetchRel: %w", exprErr)
			}
			count = e
		}

		var common plan.RelCommon
		if rel.Fetch.Common != nil {
			common = relCommonFromProto(rel.Fetch.Common)
		}
		return plan.NewFetchRel(input, offset, count, common, advancedExtensionFromProto(rel.Fetch.AdvancedExtension)), nil
	case *proto.Rel_Project:
		input, err := RelFromProto(rel.Project.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to ProjectRel: %w", err)
		}

		baseSchema := input.RecordType()

		exprs := make([]expr.Expression, len(rel.Project.Expressions))
		for i, e := range rel.Project.Expressions {
			exprs[i], err = ExprFromProto(e, &baseSchema, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting expr %d for ProjectRel: %w", i, err)
			}
		}

		if len(exprs) == 0 {
			return nil, fmt.Errorf("%w: missing required Expressions field for Project relation", substraitgo.ErrInvalidRel)
		}

		var common plan.RelCommon
		if rel.Project.Common != nil {
			common = relCommonFromProto(rel.Project.Common)
		}
		return plan.NewProjectRel(input, exprs, common, advancedExtensionFromProto(rel.Project.AdvancedExtension)), nil
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
				ge, err := ExprFromProto(e, &base, reg)
				if err != nil {
					return nil, fmt.Errorf("error getting grouping expr for AggregateRel: %w", err)
				}
				groupingExpressions = append(groupingExpressions, ge)
			}
			for _, g := range rel.Aggregate.Groupings {
				groupingReferences = append(groupingReferences, g.ExpressionReferences)
			}
		} else { // support old style grouping for backward compatibility
			groups := make([][]expr.Expression, len(rel.Aggregate.Groupings))
			for i, g := range rel.Aggregate.Groupings {
				groups[i] = make([]expr.Expression, len(g.GroupingExpressions))
				for j, e := range g.GroupingExpressions {
					groups[i][j], err = ExprFromProto(e, &base, reg)
					if err != nil {
						return nil, fmt.Errorf("error getting grouping expr [%d][%d] for AggregateRel: %w",
							i, j, err)
					}
				}
			}
			groupingExpressions, groupingReferences = groupingExprs(groups)
		}

		measures := make([]plan.AggRelMeasure, len(rel.Aggregate.Measures))
		for i, m := range rel.Aggregate.Measures {
			measure, err := AggregateFunctionFromProto(m.Measure, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting AggregateFunction for measure %d: %w", i, err)
			}

			var filter expr.Expression
			if m.Filter != nil {
				filter, err = ExprFromProto(m.Filter, &base, reg)
				if err != nil {
					return nil, fmt.Errorf("error getting filter for Aggregate Measure %d: %w", i, err)
				}
			}
			measures[i] = plan.NewAggRelMeasure(measure, filter)
		}

		common := relCommonFromProto(rel.Aggregate.Common)
		return plan.NewAggregateRel(input, measures, groupingExpressions, groupingReferences, common, advancedExtensionFromProto(rel.Aggregate.AdvancedExtension)), nil
	case *proto.Rel_Sort:
		input, err := RelFromProto(rel.Sort.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to SortRel: %w", err)
		}

		base := input.RecordType()
		sorts := make([]expr.SortField, len(rel.Sort.Sorts))
		for i, s := range rel.Sort.Sorts {
			sorts[i], err = SortFieldFromProto(s, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting SortField %d for SortRel: %w", i, err)
			}
		}

		if len(sorts) == 0 {
			return nil, fmt.Errorf("%w: missing required field Sorts for Sort Relation", substraitgo.ErrInvalidRel)
		}

		common := relCommonFromProto(rel.Sort.Common)
		return plan.NewSortRel(input, sorts, common, advancedExtensionFromProto(rel.Sort.AdvancedExtension)), nil
	case *proto.Rel_Set:
		inputs := make([]plan.Rel, len(rel.Set.Inputs))
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

		if plan.SetOp(rel.Set.Op) == plan.SetOpUnspecified {
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

		common := relCommonFromProto(rel.Set.Common)
		return plan.NewSetRel(inputs, plan.SetOp(rel.Set.Op), common, advancedExtensionFromProto(rel.Set.AdvancedExtension)), nil
	case *proto.Rel_Cross:
		left, err := RelFromProto(rel.Cross.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to CrossRel: %w", err)
		}

		right, err := RelFromProto(rel.Cross.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to CrossRel: %w", err)
		}

		common := relCommonFromProto(rel.Cross.Common)
		return plan.NewCrossRel(left, right, common, advancedExtensionFromProto(rel.Cross.AdvancedExtension)), nil
	case *proto.Rel_Join:
		if plan.JoinType(rel.Join.Type) == plan.JoinTypeUnspecified {
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

		common := relCommonFromProto(rel.Join.Common)
		// The joined record type depends only on the inputs and join type, so a
		// key-less relation gives the base schema for decoding the expressions.
		base := plan.NewJoinRel(left, right, plan.JoinType(rel.Join.Type), nil, nil, common, advancedExtensionFromProto(rel.Join.AdvancedExtension)).JoinedRecordType()
		cond, err := ExprFromProto(rel.Join.Expression, &base, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting expr for JoinRel: %w", err)
		}

		var postJoinFilter expr.Expression
		if rel.Join.PostJoinFilter != nil {
			postJoinFilter, err = ExprFromProto(rel.Join.PostJoinFilter, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error parsing PostJoinFilter for JoinRel: %w", err)
			}
		}

		return plan.NewJoinRel(left, right, plan.JoinType(rel.Join.Type), cond, postJoinFilter, common, advancedExtensionFromProto(rel.Join.AdvancedExtension)), nil
	case *proto.Rel_HashJoin:
		left, err := RelFromProto(rel.HashJoin.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to HashJoinRel: %w", err)
		}

		right, err := RelFromProto(rel.HashJoin.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to HashJoin: %w", err)
		}

		leftBase, rightBase := left.RecordType(), right.RecordType()

		keys, err := comparisonJoinKeysFromProto(
			rel.HashJoin.Keys, rel.HashJoin.LeftKeys, rel.HashJoin.RightKeys, &leftBase, &rightBase, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting keys for HashJoinRel: %w", err)
		}

		common := relCommonFromProto(rel.HashJoin.Common)
		out := plan.NewHashJoinRel(left, right, keys, plan.HashMergeJoinType(rel.HashJoin.Type), nil, common, advancedExtensionFromProto(rel.HashJoin.AdvancedExtension))

		if rel.HashJoin.PostJoinFilter != nil {
			base := out.RecordType()
			postJoinFilter, err := ExprFromProto(rel.HashJoin.PostJoinFilter, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting post join filter for HashJoinRel: %w", err)
			}
			out = plan.NewHashJoinRel(left, right, keys, plan.HashMergeJoinType(rel.HashJoin.Type), postJoinFilter, common, advancedExtensionFromProto(rel.HashJoin.AdvancedExtension))
		}

		return out, nil
	case *proto.Rel_MergeJoin:
		left, err := RelFromProto(rel.MergeJoin.Left, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting left input to MergeJoinRel: %w", err)
		}

		right, err := RelFromProto(rel.MergeJoin.Right, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting right input to MergeJoinRel: %w", err)
		}

		leftBase, rightBase := left.RecordType(), right.RecordType()

		keys, err := comparisonJoinKeysFromProto(
			rel.MergeJoin.Keys, rel.MergeJoin.LeftKeys, rel.MergeJoin.RightKeys, &leftBase, &rightBase, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting keys for MergeJoinRel: %w", err)
		}

		common := relCommonFromProto(rel.MergeJoin.Common)
		out := plan.NewMergeJoinRel(left, right, keys, plan.HashMergeJoinType(rel.MergeJoin.Type), nil, common, advancedExtensionFromProto(rel.MergeJoin.AdvancedExtension))

		if rel.MergeJoin.PostJoinFilter != nil {
			base := out.RecordType()
			postJoinFilter, err := ExprFromProto(rel.MergeJoin.PostJoinFilter, &base, reg)
			if err != nil {
				return nil, fmt.Errorf("error getting post join filter for MergeJoin: %w", err)
			}
			out = plan.NewMergeJoinRel(left, right, keys, plan.HashMergeJoinType(rel.MergeJoin.Type), postJoinFilter, common, advancedExtensionFromProto(rel.MergeJoin.AdvancedExtension))
		}

		return out, nil
	case *proto.Rel_Write:
		input, err := RelFromProto(rel.Write.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to WriteRel: %w", err)
		}
		tableSchema := NamedStructFromProto(rel.Write.TableSchema)

		var common plan.RelCommon
		if rel.Write.Common != nil {
			common = relCommonFromProto(rel.Write.Common)
		}

		var names []string
		var advExtension *extensions.AdvancedExtension
		switch rel.Write.Op {
		case proto.WriteRel_WRITE_OP_CTAS, proto.WriteRel_WRITE_OP_INSERT, proto.WriteRel_WRITE_OP_DELETE:
			switch writeType := rel.Write.WriteType.(type) {
			case *proto.WriteRel_NamedTable:
				names = writeType.NamedTable.Names
				advExtension = advancedExtensionFromProto(writeType.NamedTable.AdvancedExtension)
			case *proto.WriteRel_ExtensionTable:
				return nil, fmt.Errorf("%w: ExtensionTable not supported for WriteRel", substraitgo.ErrInvalidRel)
			default:
				return nil, fmt.Errorf("%w: WriteRel requires a NamedTable write type", substraitgo.ErrInvalidRel)
			}
		default:
			return nil, fmt.Errorf("%w: WriteRel not supported for optype %v", substraitgo.ErrInvalidRel, rel.Write.Op)
		}
		return plan.NewNamedTableWriteRel(tableSchema, plan.WriteOp(rel.Write.Op), input, plan.OutputMode(rel.Write.Output), common, names, advExtension), nil
	case *proto.Rel_ExtensionSingle:
		input, err := RelFromProto(rel.ExtensionSingle.Input, reg)
		if err != nil {
			return nil, fmt.Errorf("error getting input to ExtensionSingle: %w", err)
		}

		definition, err := decodeExtensionDef(reg, rel.ExtensionSingle.Detail)
		if err != nil {
			return nil, fmt.Errorf("error decoding ExtensionSingle detail: %w", err)
		}
		common := relCommonFromProto(rel.ExtensionSingle.Common)
		return plan.NewExtensionSingleRel(input, definition, common), nil
	case nil:
		return nil, fmt.Errorf("%w: got nil", substraitgo.ErrInvalidRel)
	}

	return nil, substraitgo.ErrNotImplemented
}
