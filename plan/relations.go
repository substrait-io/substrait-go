// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/anypb"
)

// MultiRel is a convenience interface representing any relation
// that takes an arbitrary number of inputs.
type MultiRel interface {
	Rel

	Inputs() []Rel
}

// BiRel is a convenience interface representing any relation that
// takes exactly two input relations such as joins.
type BiRel interface {
	Rel

	Left() Rel
	Right() Rel
}

// SingleInputRel is a convenience interface representing any relation
// that consists of exactly one input relation, such as a filter.
type SingleInputRel interface {
	Rel

	Input() Rel
}

// ReadRel is a scan operator of base data (physical or virtual) and
// allows filtering and projection of that underlying data.
type ReadRel interface {
	Rel

	fromProtoReadRel(*proto.ReadRel, expr.ExtensionRegistry) error

	BaseSchema() types.NamedStruct
	Filter() expr.Expression
	BestEffortFilter() expr.Expression
	Projection() *expr.MaskExpression
}

type baseReadRel struct {
	RelCommon

	baseSchema       types.NamedStruct
	filter           expr.Expression
	bestEffortFilter expr.Expression
	projection       *expr.MaskExpression
	advExtension     *extensions.AdvancedExtension
}

func (b *baseReadRel) fromProtoReadRel(rel *proto.ReadRel, reg expr.ExtensionRegistry) error {
	b.RelCommon.fromProtoCommon(rel.Common)

	b.baseSchema = types.NewNamedStructFromProto(rel.BaseSchema)
	var err error
	b.filter, err = expr.ExprFromProto(rel.Filter, &b.baseSchema.Struct, reg)
	if err != nil {
		return err
	}

	b.bestEffortFilter, err = expr.ExprFromProto(rel.BestEffortFilter, &b.baseSchema.Struct, reg)
	if err != nil {
		return err
	}

	b.projection = expr.MaskExpressionFromProto(rel.Projection)
	b.advExtension = rel.AdvancedExtension
	return nil
}

func (b *baseReadRel) RecordType() types.StructType {
	return b.baseSchema.Struct
}

func (b *baseReadRel) BaseSchema() types.NamedStruct                       { return b.baseSchema }
func (b *baseReadRel) Filter() expr.Expression                             { return b.filter }
func (b *baseReadRel) BestEffortFilter() expr.Expression                   { return b.bestEffortFilter }
func (b *baseReadRel) Projection() *expr.MaskExpression                    { return b.projection }
func (b *baseReadRel) GetAdvancedExtension() *extensions.AdvancedExtension { return b.advExtension }

func (b *baseReadRel) toReadRelProto() *proto.ReadRel {
	return &proto.ReadRel{
		Common:            b.RelCommon.toProto(),
		BaseSchema:        b.baseSchema.ToProto(),
		Filter:            b.filter.ToProto(),
		BestEffortFilter:  b.bestEffortFilter.ToProto(),
		Projection:        b.projection.ToProto(),
		AdvancedExtension: b.advExtension,
	}
}

// NamedTableReadRel is a named scan of a base table. The list of strings
// that make up the names are to represent namespacing (e.g. mydb.mytable).
// This assumes a shared catalog between systems exchanging a message.
type NamedTableReadRel struct {
	baseReadRel

	names        []string
	advExtension *extensions.AdvancedExtension
}

func (n *NamedTableReadRel) Names() []string { return n.names }

func (n *NamedTableReadRel) NamedTableAdvancedExtension() *extensions.AdvancedExtension {
	return n.advExtension
}

func (n *NamedTableReadRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: n.ToProto(),
		},
	}
}

func (n *NamedTableReadRel) ToProto() *proto.Rel {
	readRel := n.toReadRelProto()
	readRel.ReadType = &proto.ReadRel_NamedTable_{
		NamedTable: &proto.ReadRel_NamedTable{
			Names:             n.names,
			AdvancedExtension: n.advExtension,
		},
	}
	return &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: readRel,
		},
	}
}

// VirtualTableReadRel represents a table composed of literals.
type VirtualTableReadRel struct {
	baseReadRel

	values []expr.StructLiteralValue
}

func (v *VirtualTableReadRel) Values() []expr.StructLiteralValue {
	return v.values
}

func (v *VirtualTableReadRel) ToProto() *proto.Rel {
	readRel := v.toReadRelProto()
	values := make([]*proto.Expression_Literal_Struct, len(v.values))
	for i, v := range v.values {
		values[i] = v.ToProto()
	}

	readRel.ReadType = &proto.ReadRel_VirtualTable_{
		VirtualTable: &proto.ReadRel_VirtualTable{
			Values: values,
		},
	}
	return &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: readRel,
		},
	}
}

func (v *VirtualTableReadRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: v.ToProto(),
		},
	}
}

// ExtensionTableReadRel is a stub type that can be used to extend
// and introduce new table types outside the specification by utilizing
// protobuf Any type.
type ExtensionTableReadRel struct {
	baseReadRel

	detail *anypb.Any
}

func (e *ExtensionTableReadRel) Detail() *anypb.Any { return e.detail }

func (e *ExtensionTableReadRel) ToProto() *proto.Rel {
	readRel := e.toReadRelProto()
	readRel.ReadType = &proto.ReadRel_ExtensionTable_{
		ExtensionTable: &proto.ReadRel_ExtensionTable{
			Detail: e.detail,
		},
	}
	return &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: readRel,
		},
	}
}

func (e *ExtensionTableReadRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: e.ToProto(),
		},
	}
}

// PathType is the type of a LocalFileReadRel's uris.
type PathType int8

const (
	// A uri that can refer to either a single folder or a single file
	URIPath PathType = iota
	// A URI where the path portion is a glob expression that can
	// identify zero or more paths. Consumers should support
	// POSIX syntax. The recursive globstar (**) may not be supported.
	URIPathGlob
	// A URI that refers to a single file.
	URIFile
	// A URI that refers to a single folder.
	URIFolder
)

type (
	ParquetReadOptions   proto.ReadRel_LocalFiles_FileOrFiles_ParquetReadOptions
	ArrowReadOptions     proto.ReadRel_LocalFiles_FileOrFiles_ArrowReadOptions
	OrcReadOptions       proto.ReadRel_LocalFiles_FileOrFiles_OrcReadOptions
	DwrfReadOptions      proto.ReadRel_LocalFiles_FileOrFiles_DwrfReadOptions
	ExtensionReadOptions anypb.Any

	FileFormat interface {
		isFileFormat()
	}
)

func (*ParquetReadOptions) isFileFormat()   {}
func (*ArrowReadOptions) isFileFormat()     {}
func (*OrcReadOptions) isFileFormat()       {}
func (*DwrfReadOptions) isFileFormat()      {}
func (*ExtensionReadOptions) isFileFormat() {}

// FileOrFiles represents the contents of a LocalFiles table. Many files
// consist of indivisible chunks (e.g. parquet row groups or CSV rows).
// If a slice partially selects an indivisible chunk then the consumer
// should employ some rule to decide which slice to include the chunk in.
// (e.g. include it in the slice that contains the midpoint of the chunk).
type FileOrFiles struct {
	PathType PathType
	Path     string
	// PartIndex is the index of the partition that this item belongs to
	PartIndex uint64
	// Start and Len are the start position and length of bytes to
	// read from this item.
	Start, Len uint64

	Format FileFormat
}

func (f *FileOrFiles) fromProto(p *proto.ReadRel_LocalFiles_FileOrFiles) {
	f.PartIndex = p.PartitionIndex
	f.Start, f.Len = p.Start, p.Length

	switch path := p.PathType.(type) {
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriFile:
		f.PathType, f.Path = URIFile, path.UriFile
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriFolder:
		f.PathType, f.Path = URIFolder, path.UriFolder
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriPath:
		f.PathType, f.Path = URIPath, path.UriPath
	case *proto.ReadRel_LocalFiles_FileOrFiles_UriPathGlob:
		f.PathType, f.Path = URIPathGlob, path.UriPathGlob
	}

	switch format := p.FileFormat.(type) {
	case *proto.ReadRel_LocalFiles_FileOrFiles_Arrow:
		f.Format = (*ArrowReadOptions)(format.Arrow)
	case *proto.ReadRel_LocalFiles_FileOrFiles_Dwrf:
		f.Format = (*DwrfReadOptions)(format.Dwrf)
	case *proto.ReadRel_LocalFiles_FileOrFiles_Extension:
		f.Format = (*ExtensionReadOptions)(format.Extension)
	case *proto.ReadRel_LocalFiles_FileOrFiles_Orc:
		f.Format = (*OrcReadOptions)(format.Orc)
	case *proto.ReadRel_LocalFiles_FileOrFiles_Parquet:
		f.Format = (*ParquetReadOptions)(format.Parquet)
	}
}

func (f *FileOrFiles) ToProto() *proto.ReadRel_LocalFiles_FileOrFiles {
	ret := &proto.ReadRel_LocalFiles_FileOrFiles{
		PartitionIndex: f.PartIndex,
		Start:          f.Start,
		Length:         f.Len,
	}
	switch f.PathType {
	case URIPath:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriPath{UriPath: f.Path}
	case URIPathGlob:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriPathGlob{UriPathGlob: f.Path}
	case URIFile:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriFile{UriFile: f.Path}
	case URIFolder:
		ret.PathType = &proto.ReadRel_LocalFiles_FileOrFiles_UriFolder{UriFolder: f.Path}
	}

	switch fm := f.Format.(type) {
	case *ParquetReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Parquet{
			Parquet: (*proto.ReadRel_LocalFiles_FileOrFiles_ParquetReadOptions)(fm),
		}
	case *ArrowReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Arrow{
			Arrow: (*proto.ReadRel_LocalFiles_FileOrFiles_ArrowReadOptions)(fm),
		}
	case *OrcReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Orc{
			Orc: (*proto.ReadRel_LocalFiles_FileOrFiles_OrcReadOptions)(fm),
		}
	case *DwrfReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Dwrf{
			Dwrf: (*proto.ReadRel_LocalFiles_FileOrFiles_DwrfReadOptions)(fm),
		}
	case *ExtensionReadOptions:
		ret.FileFormat = &proto.ReadRel_LocalFiles_FileOrFiles_Extension{
			Extension: (*anypb.Any)(fm),
		}
	}
	return ret
}

// LocalFileReadRel represents a list of files in input of a scan operation.
type LocalFileReadRel struct {
	baseReadRel

	items        []FileOrFiles
	advExtension *extensions.AdvancedExtension
}

func (lf *LocalFileReadRel) Item(i int) FileOrFiles {
	return lf.items[i]
}

func (lf *LocalFileReadRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return lf.advExtension
}

func (lf *LocalFileReadRel) ToProto() *proto.Rel {
	items := make([]*proto.ReadRel_LocalFiles_FileOrFiles, len(lf.items))
	for i, f := range lf.items {
		items[i] = f.ToProto()
	}

	readRel := lf.toReadRelProto()
	readRel.ReadType = &proto.ReadRel_LocalFiles_{
		LocalFiles: &proto.ReadRel_LocalFiles{
			Items:             items,
			AdvancedExtension: lf.advExtension,
		},
	}
	return &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: readRel,
		},
	}
}

func (lf *LocalFileReadRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: lf.ToProto(),
		},
	}
}

// ProjectRel represents calculated expressions of fields (e.g. a+b),
// the OutputMapping will be used to represent classical relational
// projections.
type ProjectRel struct {
	RelCommon

	input        Rel
	exprs        []expr.Expression
	advExtension *extensions.AdvancedExtension
}

func (p *ProjectRel) RecordType() types.StructType {
	initial := p.input.RecordType()
	output := slices.Grow(slices.Clone(initial.Types), len(p.exprs))

	for _, e := range p.exprs {
		output = append(output, e.GetType())
	}

	return types.StructType{
		Nullability: initial.Nullability,
		Types:       output,
	}
}
func (p *ProjectRel) Input() Rel                     { return p.input }
func (p *ProjectRel) Expressions() []expr.Expression { return p.exprs }
func (p *ProjectRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return p.advExtension
}

func (p *ProjectRel) ToProto() *proto.Rel {
	exprs := make([]*proto.Expression, len(p.exprs))
	for i, e := range p.exprs {
		exprs[i] = e.ToProto()
	}

	return &proto.Rel{
		RelType: &proto.Rel_Project{
			Project: &proto.ProjectRel{
				Common:            p.toProto(),
				Input:             p.input.ToProto(),
				Expressions:       exprs,
				AdvancedExtension: p.advExtension,
			},
		},
	}
}

func (p *ProjectRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: p.ToProto(),
		},
	}
}

type JoinType = proto.JoinRel_JoinType

const (
	JoinTypeUnspecified = proto.JoinRel_JOIN_TYPE_UNSPECIFIED
	JoinTypeInner       = proto.JoinRel_JOIN_TYPE_INNER
	JoinTypeOuter       = proto.JoinRel_JOIN_TYPE_OUTER
	JoinTypeLeft        = proto.JoinRel_JOIN_TYPE_LEFT
	JoinTypeRight       = proto.JoinRel_JOIN_TYPE_RIGHT
	JoinTypeSemi        = proto.JoinRel_JOIN_TYPE_SEMI
	JoinTypeAnti        = proto.JoinRel_JOIN_TYPE_ANTI
	JoinTypeSingle      = proto.JoinRel_JOIN_TYPE_SINGLE
)

// JoinRel is a binary Join relational operator representing left-join-right,
// including various join types, a join condition and a post join filter expr.
type JoinRel struct {
	RelCommon

	left, right    Rel
	expr           expr.Expression
	postJoinFilter expr.Expression
	joinType       JoinType
	advExtension   *extensions.AdvancedExtension
}

func (j *JoinRel) RecordType() types.StructType {
	return types.StructType{
		Nullability: proto.Type_NULLABILITY_REQUIRED,
		Types:       append(j.left.RecordType().Types, j.right.RecordType().Types...),
	}
}

func (j *JoinRel) Left() Rel                       { return j.left }
func (j *JoinRel) Right() Rel                      { return j.right }
func (j *JoinRel) Expr() expr.Expression           { return j.expr }
func (j *JoinRel) PostJoinFilter() expr.Expression { return j.postJoinFilter }
func (j *JoinRel) Type() JoinType                  { return j.joinType }
func (j *JoinRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return j.advExtension
}

func (j *JoinRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Join{
			Join: &proto.JoinRel{
				Common:            j.toProto(),
				Left:              j.left.ToProto(),
				Right:             j.right.ToProto(),
				Expression:        j.expr.ToProto(),
				PostJoinFilter:    j.postJoinFilter.ToProto(),
				Type:              j.joinType,
				AdvancedExtension: j.advExtension,
			},
		},
	}
}

func (j *JoinRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: j.ToProto(),
		},
	}
}

// CrossRel is a cartesian product relational operator of two tables.
type CrossRel struct {
	RelCommon

	left, right  Rel
	advExtension *extensions.AdvancedExtension
}

func (c *CrossRel) RecordType() types.StructType {
	return types.StructType{
		Nullability: proto.Type_NULLABILITY_REQUIRED,
		Types:       append(c.left.RecordType().Types, c.right.RecordType().Types...),
	}
}

func (c *CrossRel) Left() Rel  { return c.left }
func (c *CrossRel) Right() Rel { return c.right }
func (c *CrossRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return c.advExtension
}

func (c *CrossRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Cross{
			Cross: &proto.CrossRel{
				Common:            c.toProto(),
				Left:              c.left.ToProto(),
				Right:             c.right.ToProto(),
				AdvancedExtension: c.advExtension,
			},
		},
	}
}

func (c *CrossRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: c.ToProto(),
		},
	}
}

// FetchRel is a relational operator representing LIMIT/OFFSET or
// TOP type semantics.
type FetchRel struct {
	RelCommon

	input         Rel
	offset, count int64
	advExtension  *extensions.AdvancedExtension
}

func (f *FetchRel) RecordType() types.StructType { return f.input.RecordType() }
func (f *FetchRel) Input() Rel                   { return f.input }
func (f *FetchRel) Offset() int64                { return f.offset }
func (f *FetchRel) Count() int64                 { return f.count }
func (f *FetchRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return f.advExtension
}

func (f *FetchRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Fetch{
			Fetch: &proto.FetchRel{
				Common:            f.toProto(),
				Input:             f.input.ToProto(),
				Offset:            f.offset,
				Count:             f.count,
				AdvancedExtension: f.advExtension,
			},
		},
	}
}

func (f *FetchRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: f.ToProto(),
		},
	}
}

type AggRelMeasure struct {
	measure *expr.AggregateFunction
	filter  expr.Expression
}

func (am *AggRelMeasure) Measure() *expr.AggregateFunction { return am.measure }
func (am *AggRelMeasure) Filter() expr.Expression          { return am.filter }

func (am *AggRelMeasure) ToProto() *proto.AggregateRel_Measure {
	return &proto.AggregateRel_Measure{
		Measure: am.measure.ToProto(),
		Filter:  am.filter.ToProto(),
	}
}

// AggregateRel is a relational operator representing a GROUP BY aggregate.
type AggregateRel struct {
	RelCommon

	input        Rel
	groups       [][]expr.Expression
	measures     []AggRelMeasure
	advExtension *extensions.AdvancedExtension
}

func (ar *AggregateRel) RecordType() types.StructType {
	groupTypes := make([]types.Type, 0, len(ar.groups)+len(ar.measures))
	for _, g := range ar.groups {
		for _, e := range g {
			groupTypes = append(groupTypes, e.GetType())
		}
	}

	for _, m := range ar.measures {
		groupTypes = append(groupTypes, m.measure.GetType())
	}

	return types.StructType{
		Nullability: proto.Type_NULLABILITY_REQUIRED,
		Types:       groupTypes,
	}
}

func (ar *AggregateRel) Input() Rel { return ar.input }

// Groupings is a list of expression groupings that the aggregation measures should
// be calculated for.
func (ar *AggregateRel) Groupings() [][]expr.Expression { return ar.groups }
func (ar *AggregateRel) Measures() []AggRelMeasure      { return ar.measures }
func (ar *AggregateRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return ar.advExtension
}

func (ar *AggregateRel) ToProto() *proto.Rel {
	groupings := make([]*proto.AggregateRel_Grouping, len(ar.groups))
	for i, ex := range ar.groups {
		groupExprs := make([]*proto.Expression, len(ex))
		for j, e := range ex {
			groupExprs[j] = e.ToProto()
		}
		groupings[i] = &proto.AggregateRel_Grouping{
			GroupingExpressions: groupExprs,
		}
	}

	measures := make([]*proto.AggregateRel_Measure, len(ar.measures))
	for i, m := range ar.measures {
		measures[i] = m.ToProto()
	}

	return &proto.Rel{
		RelType: &proto.Rel_Aggregate{
			Aggregate: &proto.AggregateRel{
				Common:            ar.toProto(),
				Input:             ar.input.ToProto(),
				Groupings:         groupings,
				Measures:          measures,
				AdvancedExtension: ar.advExtension,
			},
		},
	}
}

func (ar *AggregateRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: ar.ToProto(),
		},
	}
}

// SortRel is an ORDER BY relational operator, describing a base relation,
// it includes a list of fields to sort on.
type SortRel struct {
	RelCommon

	input        Rel
	sorts        []expr.SortField
	advExtension *extensions.AdvancedExtension
}

func (sr *SortRel) RecordType() types.StructType { return sr.input.RecordType() }
func (sr *SortRel) Input() Rel                   { return sr.input }
func (sr *SortRel) Sorts() []expr.SortField      { return sr.sorts }
func (sr *SortRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return sr.advExtension
}

func (sr *SortRel) ToProto() *proto.Rel {
	sorts := make([]*proto.SortField, len(sr.sorts))
	for i, s := range sr.sorts {
		sorts[i] = s.ToProto()
	}
	return &proto.Rel{
		RelType: &proto.Rel_Sort{
			Sort: &proto.SortRel{
				Common:            sr.toProto(),
				Input:             sr.input.ToProto(),
				Sorts:             sorts,
				AdvancedExtension: sr.advExtension,
			},
		},
	}
}

func (sr *SortRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: sr.ToProto(),
		},
	}
}

// FilterRel is a relational operator capturing simple filters (
// as in the WHERE clause of a SQL query).
type FilterRel struct {
	RelCommon

	input        Rel
	cond         expr.Expression
	advExtension *extensions.AdvancedExtension
}

func (fr *FilterRel) RecordType() types.StructType { return fr.input.RecordType() }
func (fr *FilterRel) Input() Rel                   { return fr.input }
func (fr *FilterRel) Condition() expr.Expression   { return fr.cond }
func (fr *FilterRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return fr.advExtension
}

func (fr *FilterRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Filter{
			Filter: &proto.FilterRel{
				Common:            fr.toProto(),
				Input:             fr.input.ToProto(),
				Condition:         fr.cond.ToProto(),
				AdvancedExtension: fr.advExtension,
			},
		},
	}
}

func (fr *FilterRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: fr.ToProto(),
		},
	}
}

type SetOp = proto.SetRel_SetOp

const (
	SetOpUnspecified          = proto.SetRel_SET_OP_UNSPECIFIED
	SetOpMinusPrimary         = proto.SetRel_SET_OP_MINUS_PRIMARY
	SetOpMinusMultiset        = proto.SetRel_SET_OP_MINUS_MULTISET
	SetOpIntersectionPrimary  = proto.SetRel_SET_OP_INTERSECTION_PRIMARY
	SetOpIntersectionMultiset = proto.SetRel_SET_OP_INTERSECTION_MULTISET
	SetOpUnionDistinct        = proto.SetRel_SET_OP_UNION_DISTINCT
	SetOpUnionAll             = proto.SetRel_SET_OP_UNION_ALL
)

// SetRel represents the relational set operators (intersection, union, etc.)
type SetRel struct {
	RelCommon

	inputs       []Rel
	op           SetOp
	advExtension *extensions.AdvancedExtension
}

func (s *SetRel) RecordType() types.StructType { return s.inputs[0].RecordType() }
func (s *SetRel) Inputs() []Rel                { return s.inputs }
func (s *SetRel) Op() SetOp                    { return s.op }
func (s *SetRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return s.advExtension
}

func (s *SetRel) ToProto() *proto.Rel {
	inputs := make([]*proto.Rel, len(s.inputs))
	for i, in := range s.inputs {
		inputs[i] = in.ToProto()
	}
	return &proto.Rel{
		RelType: &proto.Rel_Set{
			Set: &proto.SetRel{
				Common:            s.toProto(),
				Inputs:            inputs,
				Op:                s.op,
				AdvancedExtension: s.advExtension,
			},
		},
	}
}

func (s *SetRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: s.ToProto(),
		},
	}
}

// ExtensionSingleRel is a stub to support extensions with a single input.
type ExtensionSingleRel struct {
	RelCommon

	input  Rel
	detail *anypb.Any
}

func (es *ExtensionSingleRel) RecordType() types.StructType { return es.input.RecordType() }

func (es *ExtensionSingleRel) Input() Rel         { return es.input }
func (es *ExtensionSingleRel) Detail() *anypb.Any { return es.detail }

func (es *ExtensionSingleRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_ExtensionSingle{
			ExtensionSingle: &proto.ExtensionSingleRel{
				Common: es.toProto(),
				Input:  es.input.ToProto(),
				Detail: es.detail,
			},
		},
	}
}

func (es *ExtensionSingleRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: es.ToProto(),
		},
	}
}

// ExtensionLeafRel is a stub to support extensions with zero inputs.
type ExtensionLeafRel struct {
	RelCommon

	detail *anypb.Any
}

func (el *ExtensionLeafRel) RecordType() types.StructType { return types.StructType{} }
func (el *ExtensionLeafRel) Detail() *anypb.Any           { return el.detail }

func (el *ExtensionLeafRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_ExtensionLeaf{
			ExtensionLeaf: &proto.ExtensionLeafRel{
				Common: el.toProto(),
				Detail: el.detail,
			},
		},
	}
}

func (el *ExtensionLeafRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: el.ToProto(),
		},
	}
}

// ExtensionMultiRel is a stub to support extensions with multiple inputs.
type ExtensionMultiRel struct {
	RelCommon

	inputs []Rel
	detail *anypb.Any
}

func (em *ExtensionMultiRel) RecordType() types.StructType { return types.StructType{} }
func (em *ExtensionMultiRel) Inputs() []Rel                { return em.inputs }
func (em *ExtensionMultiRel) Detail() *anypb.Any           { return em.detail }

func (em *ExtensionMultiRel) ToProto() *proto.Rel {
	inputs := make([]*proto.Rel, len(em.inputs))
	for i, in := range em.inputs {
		inputs[i] = in.ToProto()
	}
	return &proto.Rel{
		RelType: &proto.Rel_ExtensionMulti{
			ExtensionMulti: &proto.ExtensionMultiRel{
				Common: em.toProto(),
				Inputs: inputs,
				Detail: em.detail,
			},
		},
	}
}

func (em *ExtensionMultiRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: em.ToProto(),
		},
	}
}

type HashMergeJoinType int8

const (
	HashMergeUnspecified HashMergeJoinType = iota
	HashMergeInner
	HashMergeOuter
	HashMergeLeft
	HashMergeRight
	HashMergeLeftSemi
	HashMergeRightSemi
	HashMergeLeftAnti
	HashMergeRightAnti
)

// HashJoinRel represents a relational operator to build a hash table out
// of the right input based on a set of join keys. It will then probe
// the hash table for incoming inputs, finding matches.
type HashJoinRel struct {
	RelCommon

	left, right         Rel
	leftKeys, rightKeys []*expr.FieldReference
	postJoinFilter      expr.Expression
	joinType            HashMergeJoinType
	advExtension        *extensions.AdvancedExtension
}

func (hr *HashJoinRel) RecordType() types.StructType {
	return types.StructType{
		Nullability: proto.Type_NULLABILITY_REQUIRED,
		Types:       append(hr.left.RecordType().Types, hr.right.RecordType().Types...),
	}
}

func (hr *HashJoinRel) Left() Rel                         { return hr.left }
func (hr *HashJoinRel) Right() Rel                        { return hr.right }
func (hr *HashJoinRel) LeftKeys() []*expr.FieldReference  { return hr.leftKeys }
func (hr *HashJoinRel) RightKeys() []*expr.FieldReference { return hr.rightKeys }
func (hr *HashJoinRel) PostJoinFilter() expr.Expression   { return hr.postJoinFilter }
func (hr *HashJoinRel) Type() HashMergeJoinType           { return hr.joinType }
func (hr *HashJoinRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return hr.advExtension
}

func (hr *HashJoinRel) ToProto() *proto.Rel {
	keysLeft := make([]*proto.Expression_FieldReference, len(hr.leftKeys))
	for i, k := range hr.leftKeys {
		keysLeft[i] = k.ToProtoFieldRef()
	}
	keysRight := make([]*proto.Expression_FieldReference, len(hr.rightKeys))
	for i, k := range hr.rightKeys {
		keysRight[i] = k.ToProtoFieldRef()
	}

	return &proto.Rel{
		RelType: &proto.Rel_HashJoin{
			HashJoin: &proto.HashJoinRel{
				Common:            hr.toProto(),
				Left:              hr.left.ToProto(),
				Right:             hr.right.ToProto(),
				LeftKeys:          keysLeft,
				RightKeys:         keysRight,
				PostJoinFilter:    hr.postJoinFilter.ToProto(),
				Type:              proto.HashJoinRel_JoinType(hr.joinType),
				AdvancedExtension: hr.advExtension,
			},
		},
	}
}

func (hr *HashJoinRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: hr.ToProto(),
		},
	}
}

// MergeJoinRel represents a join done by taking advantage of two sets
// that are sorted on the join keys. This allows the join operation to
// be done in a streaming fashion.
type MergeJoinRel struct {
	RelCommon

	left, right         Rel
	leftKeys, rightKeys []*expr.FieldReference
	postJoinFilter      expr.Expression
	joinType            HashMergeJoinType
	advExtension        *extensions.AdvancedExtension
}

func (mr *MergeJoinRel) RecordType() types.StructType {
	return types.StructType{
		Nullability: proto.Type_NULLABILITY_REQUIRED,
		Types:       append(mr.left.RecordType().Types, mr.right.RecordType().Types...),
	}
}

func (mr *MergeJoinRel) Left() Rel                         { return mr.left }
func (mr *MergeJoinRel) Right() Rel                        { return mr.right }
func (mr *MergeJoinRel) LeftKeys() []*expr.FieldReference  { return mr.leftKeys }
func (mr *MergeJoinRel) RightKeys() []*expr.FieldReference { return mr.rightKeys }
func (mr *MergeJoinRel) PostJoinFilter() expr.Expression   { return mr.postJoinFilter }
func (mr *MergeJoinRel) Type() HashMergeJoinType           { return mr.joinType }
func (mr *MergeJoinRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return mr.advExtension
}

func (mr *MergeJoinRel) ToProto() *proto.Rel {
	keysLeft := make([]*proto.Expression_FieldReference, len(mr.leftKeys))
	for i, k := range mr.leftKeys {
		keysLeft[i] = k.ToProtoFieldRef()
	}
	keysRight := make([]*proto.Expression_FieldReference, len(mr.rightKeys))
	for i, k := range mr.rightKeys {
		keysRight[i] = k.ToProtoFieldRef()
	}

	return &proto.Rel{
		RelType: &proto.Rel_MergeJoin{
			MergeJoin: &proto.MergeJoinRel{
				Common:            mr.toProto(),
				Left:              mr.left.ToProto(),
				Right:             mr.right.ToProto(),
				LeftKeys:          keysLeft,
				RightKeys:         keysRight,
				PostJoinFilter:    mr.postJoinFilter.ToProto(),
				Type:              proto.MergeJoinRel_JoinType(mr.joinType),
				AdvancedExtension: mr.advExtension,
			},
		},
	}
}

func (mr *MergeJoinRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: mr.ToProto(),
		},
	}
}

var (
	_ Rel = (*NamedTableReadRel)(nil)
	_ Rel = (*VirtualTableReadRel)(nil)
	_ Rel = (*LocalFileReadRel)(nil)
	_ Rel = (*ExtensionTableReadRel)(nil)
	_ Rel = (*ProjectRel)(nil)
	_ Rel = (*JoinRel)(nil)
	_ Rel = (*CrossRel)(nil)
	_ Rel = (*FetchRel)(nil)
	_ Rel = (*AggregateRel)(nil)
	_ Rel = (*FilterRel)(nil)
	_ Rel = (*SortRel)(nil)
	_ Rel = (*SetRel)(nil)
	_ Rel = (*ExtensionSingleRel)(nil)
	_ Rel = (*ExtensionLeafRel)(nil)
	_ Rel = (*ExtensionMultiRel)(nil)
	_ Rel = (*HashJoinRel)(nil)
	_ Rel = (*MergeJoinRel)(nil)

	_ MultiRel = (*SetRel)(nil)
	_ MultiRel = (*ExtensionMultiRel)(nil)

	_ BiRel = (*JoinRel)(nil)
	_ BiRel = (*CrossRel)(nil)
	_ BiRel = (*HashJoinRel)(nil)
	_ BiRel = (*MergeJoinRel)(nil)

	_ SingleInputRel = (*ProjectRel)(nil)
	_ SingleInputRel = (*FetchRel)(nil)
	_ SingleInputRel = (*AggregateRel)(nil)
	_ SingleInputRel = (*FilterRel)(nil)
	_ SingleInputRel = (*SortRel)(nil)
	_ SingleInputRel = (*ExtensionSingleRel)(nil)
)
