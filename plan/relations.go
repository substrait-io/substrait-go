// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v3"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/extensions"
	"github.com/substrait-io/substrait-go/v3/proto"
	"github.com/substrait-io/substrait-go/v3/types"
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
	if rel.Common != nil {
		b.RelCommon.fromProtoCommon(rel.Common)
	}

	b.baseSchema = types.NewNamedStructFromProto(rel.BaseSchema)
	var err error
	if rel.Filter != nil {
		b.filter, err = expr.ExprFromProto(rel.Filter, types.NewRecordTypeFromStruct(b.baseSchema.Struct), reg)
		if err != nil {
			return err
		}
	}

	if rel.BestEffortFilter != nil {
		b.bestEffortFilter, err = expr.ExprFromProto(rel.BestEffortFilter, types.NewRecordTypeFromStruct(b.baseSchema.Struct), reg)
		if err != nil {
			return err
		}
	}

	if rel.Projection != nil {
		b.projection = expr.MaskExpressionFromProto(rel.Projection)
	}

	b.advExtension = rel.AdvancedExtension
	return nil
}

func (b *baseReadRel) directOutputSchema() types.RecordType {
	return *types.NewRecordTypeFromStruct(b.baseSchema.Struct)
}

func (b *baseReadRel) RecordType() types.RecordType {
	return b.remap(b.directOutputSchema())
}

func (b *baseReadRel) BaseSchema() types.NamedStruct                       { return b.baseSchema }
func (b *baseReadRel) Filter() expr.Expression                             { return b.filter }
func (b *baseReadRel) BestEffortFilter() expr.Expression                   { return b.bestEffortFilter }
func (b *baseReadRel) Projection() *expr.MaskExpression                    { return b.projection }
func (b *baseReadRel) GetAdvancedExtension() *extensions.AdvancedExtension { return b.advExtension }

func (b *baseReadRel) SetProjection(p *expr.MaskExpression) {
	b.projection = p
}

func (b *baseReadRel) toReadRelProto() *proto.ReadRel {
	out := &proto.ReadRel{
		Common:            b.RelCommon.toProto(),
		BaseSchema:        b.baseSchema.ToProto(),
		AdvancedExtension: b.advExtension,
	}
	if b.filter != nil {
		out.Filter = b.filter.ToProto()
	}
	if b.bestEffortFilter != nil {
		out.BestEffortFilter = b.bestEffortFilter.ToProto()
	}
	if b.projection != nil {
		out.Projection = b.projection.ToProto()
	}

	return out
}

func (b *baseReadRel) GetInputs() []Rel {
	return []Rel{}
}

func (b *baseReadRel) copyExpressions(rewriteFunc RewriteFunc) ([]expr.Expression, error) {
	filter, err := rewriteFunc(b.filter)
	if err != nil {
		return nil, err
	}
	bestEffortFilter, err := rewriteFunc(b.bestEffortFilter)
	if err != nil {
		return nil, err
	}
	return []expr.Expression{filter, bestEffortFilter}, nil
}

func (b *baseReadRel) getExpressions() []expr.Expression {
	return []expr.Expression{b.filter, b.bestEffortFilter}
}

func (b *baseReadRel) updateFilters(filters []expr.Expression) {
	b.filter, b.bestEffortFilter = filters[0], filters[1]
}

// isSequentialFromZero checks if the given slice of integers is a sequence starting from zero
// where each element in the slice is equal to its index.
func isSequentialFromZero(order []int32) bool {
	for x, i := range order {
		if i != int32(x) {
			return false
		}
	}
	return true
}

// RemapHelper implements the core functionality of RemapHelper for relations.
func RemapHelper(r Rel, mapping []int32) (Rel, error) {
	newRel, err := r.Copy(r.GetInputs()...)
	if err != nil {
		return nil, err
	}
	if len(mapping) == 0 {
		newRel.setMapping([]int32{})
		return newRel, nil
	}
	nOutput := r.RecordType().FieldCount()
	oldMapping := r.OutputMapping()
	newMapping := make([]int32, 0, len(mapping))
	for _, idx := range mapping {
		if idx < 0 || idx >= nOutput {
			return nil, errOutputMappingOutOfRange
		}
		if len(oldMapping) > 0 {
			newMapping = append(newMapping, oldMapping[idx])
		} else {
			newMapping = append(newMapping, idx)
		}
	}
	if len(newMapping) == int(r.directOutputSchema().FieldCount()) && isSequentialFromZero(newMapping) {
		newRel.setMapping(nil)
	} else {
		newRel.setMapping(newMapping)
	}
	return newRel, nil
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

func (n *NamedTableReadRel) RecordType() types.RecordType {
	return n.remap(n.directOutputSchema())
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

func (n *NamedTableReadRel) Copy(_ ...Rel) (Rel, error) {
	return n, nil
}

func (n *NamedTableReadRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, _ ...Rel) (Rel, error) {
	newExprs, err := n.copyExpressions(rewriteFunc)
	if err != nil {
		return nil, err
	}
	if slices.Equal(newExprs, n.getExpressions()) {
		return n, nil
	}
	nt := *n
	nt.updateFilters(newExprs)
	return &nt, nil
}

func (n *NamedTableReadRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(n, mapping)
}

// VirtualTableReadRel represents a table composed of literals.
type VirtualTableReadRel struct {
	baseReadRel

	values []expr.VirtualTableExpressionValue
}

func (v *VirtualTableReadRel) Values() []expr.VirtualTableExpressionValue {
	return v.values
}

func (v *VirtualTableReadRel) ToProto() *proto.Rel {
	readRel := v.toReadRelProto()
	values := make([]*proto.Expression_Nested_Struct, len(v.values))
	for i, v := range v.values {
		values[i] = v.ToProto()
	}

	readRel.ReadType = &proto.ReadRel_VirtualTable_{
		VirtualTable: &proto.ReadRel_VirtualTable{
			Expressions: values,
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

func (v *VirtualTableReadRel) Copy(_ ...Rel) (Rel, error) {
	return v, nil
}

func (v *VirtualTableReadRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, _ ...Rel) (Rel, error) {
	newExprs, err := v.copyExpressions(rewriteFunc)
	if err != nil {
		return nil, err
	}
	if slices.Equal(newExprs, v.getExpressions()) {
		return v, nil
	}
	vtr := *v
	vtr.updateFilters(newExprs)
	return &vtr, nil
}

func (v *VirtualTableReadRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(v, mapping)
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

func (e *ExtensionTableReadRel) Copy(_ ...Rel) (Rel, error) {
	return e, nil
}

func (e *ExtensionTableReadRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, _ ...Rel) (Rel, error) {
	newExprs, err := e.copyExpressions(rewriteFunc)
	if err != nil {
		return nil, err
	}
	if slices.Equal(newExprs, e.getExpressions()) {
		return e, nil
	}
	etr := *e
	etr.updateFilters(newExprs)
	return &etr, nil
}

func (e *ExtensionTableReadRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(e, mapping)
}

type (
	SnapshotId        string
	SnapshotTimestamp int64

	IcebergSnapshot interface {
		isSnapshot()
	}
)

func (SnapshotId) isSnapshot()        {}
func (SnapshotTimestamp) isSnapshot() {}

type (
	Direct struct {
		MetadataUri string
		SnapshotTimestamp
		SnapshotId
	}

	IcebergTableType interface {
		isTableType()
	}
)

func (*Direct) isTableType() {}

// IcebergTableReadRel is a scan on iceberg table.
type IcebergTableReadRel struct {
	baseReadRel

	tableType    IcebergTableType
	advExtension *extensions.AdvancedExtension
}

func (n *IcebergTableReadRel) NamedTableAdvancedExtension() *extensions.AdvancedExtension {
	return n.advExtension
}

func (n *IcebergTableReadRel) RecordType() types.RecordType {
	return n.remap(n.directOutputSchema())
}

func (n *IcebergTableReadRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: n.ToProto(),
		},
	}
}

func (n *IcebergTableReadRel) ToProto() *proto.Rel {
	readRel := n.toReadRelProto()

	if directTableType, ok := n.tableType.(*Direct); ok {
		direct := &proto.ReadRel_IcebergTable_MetadataFileRead{
			MetadataUri: directTableType.MetadataUri,
		}

		// SnapshotId and SnapshotTimestamp are mutually exclusive
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
				TableType: &proto.ReadRel_IcebergTable_Direct{
					Direct: direct,
				},
			},
		}
	}

	return &proto.Rel{
		RelType: &proto.Rel_Read{
			Read: readRel,
		},
	}
}

func (n *IcebergTableReadRel) Copy(_ ...Rel) (Rel, error) {
	return n, nil
}

func (n *IcebergTableReadRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, _ ...Rel) (Rel, error) {
	newExprs, err := n.copyExpressions(rewriteFunc)
	if err != nil {
		return nil, err
	}
	if slices.Equal(newExprs, n.getExpressions()) {
		return n, nil
	}
	nt := *n
	nt.updateFilters(newExprs)
	return &nt, nil
}

func (n *IcebergTableReadRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(n, mapping)
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

func (lf *LocalFileReadRel) Copy(_ ...Rel) (Rel, error) {
	return lf, nil
}

func (lf *LocalFileReadRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, _ ...Rel) (Rel, error) {
	newExprs, err := lf.copyExpressions(rewriteFunc)
	if err != nil {
		return nil, err
	}
	if slices.Equal(newExprs, lf.getExpressions()) {
		return lf, nil
	}
	lfr := *lf
	lfr.updateFilters(newExprs)
	return &lfr, nil
}

func (lf *LocalFileReadRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(lf, mapping)
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

func (p *ProjectRel) directOutputSchema() types.RecordType {
	initial := p.input.RecordType()
	output := slices.Grow(slices.Clone(initial.Types()), len(p.exprs))

	for _, e := range p.exprs {
		output = append(output, e.GetType())
	}

	return *types.NewRecordTypeFromTypes(output)
}
func (p *ProjectRel) RecordType() types.RecordType {
	return p.remap(p.directOutputSchema())
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

func (p *ProjectRel) GetInputs() []Rel {
	return []Rel{p.input}
}

func (p *ProjectRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	proj := *p
	proj.input = newInputs[0]
	return &proj, nil
}

func (p *ProjectRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	var err error
	exprs := make([]expr.Expression, len(p.exprs))
	for i, e := range p.exprs {
		if exprs[i], err = rewriteFunc(e); err != nil {
			return nil, err
		}
	}
	if slices.Equal(exprs, p.exprs) && slices.Equal(newInputs, p.GetInputs()) {
		return p, nil
	}
	proj := *p
	proj.input = newInputs[0]
	proj.exprs = exprs
	return &proj, nil
}

func (p *ProjectRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(p, mapping)
}

var defFilter = expr.NewPrimitiveLiteral(true, false)

type JoinType = proto.JoinRel_JoinType

const (
	JoinTypeUnspecified = proto.JoinRel_JOIN_TYPE_UNSPECIFIED
	JoinTypeInner       = proto.JoinRel_JOIN_TYPE_INNER
	JoinTypeOuter       = proto.JoinRel_JOIN_TYPE_OUTER
	JoinTypeLeft        = proto.JoinRel_JOIN_TYPE_LEFT
	JoinTypeRight       = proto.JoinRel_JOIN_TYPE_RIGHT
	JoinTypeLeftSemi    = proto.JoinRel_JOIN_TYPE_LEFT_SEMI
	JoinTypeLeftAnti    = proto.JoinRel_JOIN_TYPE_LEFT_ANTI
	JoinTypeLeftSingle  = proto.JoinRel_JOIN_TYPE_LEFT_SINGLE
	JoinTypeRightSemi   = proto.JoinRel_JOIN_TYPE_RIGHT_SEMI
	JoinTypeRightAnti   = proto.JoinRel_JOIN_TYPE_RIGHT_ANTI
	JoinTypeRightSingle = proto.JoinRel_JOIN_TYPE_RIGHT_SINGLE
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

func (j *JoinRel) directOutputSchema() types.RecordType {
	var typeList []types.Type
	switch j.joinType {
	case JoinTypeInner:
		return j.JoinedRecordType()
	case JoinTypeLeftSemi:
		return j.left.RecordType()
	case JoinTypeOuter:
		typeList = j.JoinedRecordType().Types()
		for i, t := range typeList {
			typeList[i] = t.WithNullability(types.NullabilityNullable)
		}
	case JoinTypeLeft, JoinTypeLeftSingle:
		left := j.left.RecordType()
		right := j.right.RecordType()
		typeList = make([]types.Type, 0, left.FieldCount()+right.FieldCount())
		typeList = append(typeList, left.Types()...)
		for _, r := range right.Types() {
			typeList = append(typeList, r.WithNullability(types.NullabilityNullable))
		}
	case JoinTypeRight:
		left := j.left.RecordType()
		right := j.right.RecordType()
		typeList = make([]types.Type, 0, left.FieldCount()+right.FieldCount())
		for _, l := range left.Types() {
			typeList = append(typeList, l.WithNullability(types.NullabilityNullable))
		}
		typeList = append(typeList, right.Types()...)
	case JoinTypeLeftAnti:
		typeList = j.left.RecordType().Types()
	case JoinTypeRightSemi, JoinTypeRightAnti, JoinTypeRightSingle:
		panic(fmt.Sprintf("join type: %v not supported", j.joinType))
	}

	return *types.NewRecordTypeFromTypes(typeList)
}

func (j *JoinRel) RecordType() types.RecordType {
	return j.remap(j.directOutputSchema())
}

func (j *JoinRel) JoinedRecordType() types.RecordType {
	return j.left.RecordType().Concat(j.right.RecordType())
}

func (j *JoinRel) Left() Rel             { return j.left }
func (j *JoinRel) Right() Rel            { return j.right }
func (j *JoinRel) Expr() expr.Expression { return j.expr }
func (j *JoinRel) PostJoinFilter() expr.Expression {
	if j.postJoinFilter == nil {
		return defFilter
	}
	return j.postJoinFilter
}
func (j *JoinRel) Type() JoinType { return j.joinType }
func (j *JoinRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return j.advExtension
}

func (j *JoinRel) ToProto() *proto.Rel {
	outRel := &proto.JoinRel{
		Common:            j.toProto(),
		Left:              j.left.ToProto(),
		Right:             j.right.ToProto(),
		Expression:        j.expr.ToProto(),
		Type:              j.joinType,
		AdvancedExtension: j.advExtension,
	}

	if j.postJoinFilter != nil {
		outRel.PostJoinFilter = j.postJoinFilter.ToProto()
	}

	return &proto.Rel{
		RelType: &proto.Rel_Join{
			Join: outRel,
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

func (j *JoinRel) GetInputs() []Rel {
	return []Rel{j.left, j.right}
}

func (j *JoinRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	join := *j
	join.left, join.right = newInputs[0], newInputs[1]
	return &join, nil
}

func (j *JoinRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	var err error
	currExprs := []expr.Expression{j.expr, j.postJoinFilter}
	exprs := make([]expr.Expression, 2)
	for i, e := range currExprs {
		if exprs[i], err = rewriteFunc(e); err != nil {
			return nil, err
		}
	}
	if slices.Equal(exprs, currExprs) && slices.Equal(newInputs, j.GetInputs()) {
		return j, nil
	}
	join := *j
	join.left, join.right = newInputs[0], newInputs[1]
	join.expr, join.postJoinFilter = exprs[0], exprs[1]
	return &join, nil
}

func (j *JoinRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(j, mapping)
}

// CrossRel is a cartesian product relational operator of two tables.
type CrossRel struct {
	RelCommon

	left, right  Rel
	advExtension *extensions.AdvancedExtension
}

func (c *CrossRel) directOutputSchema() types.RecordType {
	return c.left.RecordType().Concat(c.right.RecordType())
}
func (c *CrossRel) RecordType() types.RecordType {
	return c.remap(c.directOutputSchema())
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

func (c *CrossRel) GetInputs() []Rel {
	return []Rel{c.left, c.right}
}

func (c *CrossRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	cross := *c
	cross.left, cross.right = newInputs[0], newInputs[1]
	return &cross, nil
}

func (c *CrossRel) CopyWithExpressionRewrite(_ RewriteFunc, newInputs ...Rel) (Rel, error) {
	if slices.Equal(newInputs, c.GetInputs()) {
		return c, nil
	}
	return c.Copy(newInputs...)
}

func (c *CrossRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(c, mapping)
}

// FetchRel is a relational operator representing LIMIT/OFFSET or
// TOP type semantics.
type FetchRel struct {
	RelCommon

	input         Rel
	offset, count int64
	advExtension  *extensions.AdvancedExtension
}

func (f *FetchRel) directOutputSchema() types.RecordType { return f.input.RecordType() }
func (f *FetchRel) RecordType() types.RecordType {
	return f.remap(f.directOutputSchema())
}
func (f *FetchRel) Input() Rel    { return f.input }
func (f *FetchRel) Offset() int64 { return f.offset }
func (f *FetchRel) Count() int64  { return f.count }
func (f *FetchRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return f.advExtension
}

func (f *FetchRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Fetch{
			Fetch: &proto.FetchRel{
				Common: f.toProto(),
				Input:  f.input.ToProto(),
				OffsetMode: &proto.FetchRel_Offset{
					Offset: f.offset,
				},
				CountMode: &proto.FetchRel_Count{
					Count: f.count,
				},
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

func (f *FetchRel) GetInputs() []Rel {
	return []Rel{f.input}
}

func (f *FetchRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	fetch := *f
	fetch.input = newInputs[0]
	return &fetch, nil
}

func (f *FetchRel) CopyWithExpressionRewrite(_ RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	if newInputs[0] == f.input {
		return f, nil
	}
	return f.Copy(newInputs...)
}

func (f *FetchRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(f, mapping)
}

type AggRelMeasure struct {
	measure *expr.AggregateFunction
	filter  expr.Expression
}

func (am *AggRelMeasure) Measure() *expr.AggregateFunction { return am.measure }
func (am *AggRelMeasure) Filter() expr.Expression {
	if am.filter == nil {
		return defFilter
	}
	return am.filter
}

func (am *AggRelMeasure) ToProto() *proto.AggregateRel_Measure {
	ret := &proto.AggregateRel_Measure{
		Measure: am.measure.ToProto(),
	}
	if am.filter != nil {
		ret.Filter = am.filter.ToProto()
	}
	return ret
}

// AggregateRel is a relational operator representing a GROUP BY aggregate.
type AggregateRel struct {
	RelCommon

	input               Rel
	measures            []AggRelMeasure
	groupingExpressions []expr.Expression
	groupingReferences  [][]uint32
	advExtension        *extensions.AdvancedExtension
}

func (ar *AggregateRel) directOutputSchema() types.RecordType {
	groupTypes := make([]types.Type, 0, len(ar.groupingReferences)+len(ar.measures))
	for _, g := range ar.groupingReferences {
		for _, e := range g {
			expression := ar.groupingExpressions[e]
			groupTypes = append(groupTypes, expression.GetType())
		}
	}

	for _, m := range ar.measures {
		groupTypes = append(groupTypes, m.measure.GetType())
	}

	return *types.NewRecordTypeFromTypes(groupTypes)
}

func (ar *AggregateRel) RecordType() types.RecordType {
	return ar.remap(ar.directOutputSchema())
}

func (ar *AggregateRel) Input() Rel { return ar.input }

func (ar *AggregateRel) GroupingExpressions() []expr.Expression { return ar.groupingExpressions }
func (ar *AggregateRel) GroupingReferences() [][]uint32         { return ar.groupingReferences }
func (ar *AggregateRel) Measures() []AggRelMeasure              { return ar.measures }
func (ar *AggregateRel) GetAdvancedExtension() *extensions.AdvancedExtension {
	return ar.advExtension
}

func (ar *AggregateRel) ToProto() *proto.Rel {
	groupingExpressionsProto := make([]*proto.Expression, len(ar.groupingExpressions))
	for i, e := range ar.groupingExpressions {
		groupingExpressionsProto[i] = e.ToProto()
	}

	groupings := make([]*proto.AggregateRel_Grouping, len(ar.groupingReferences))
	for i := range ar.groupingReferences {
		groupings[i] = &proto.AggregateRel_Grouping{
			ExpressionReferences: ar.groupingReferences[i],
		}
	}

	measures := make([]*proto.AggregateRel_Measure, len(ar.measures))
	for i, m := range ar.measures {
		measures[i] = m.ToProto()
	}

	return &proto.Rel{
		RelType: &proto.Rel_Aggregate{
			Aggregate: &proto.AggregateRel{
				Common:              ar.toProto(),
				Input:               ar.input.ToProto(),
				GroupingExpressions: groupingExpressionsProto,
				Groupings:           groupings,
				Measures:            measures,
				AdvancedExtension:   ar.advExtension,
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

func (ar *AggregateRel) GetInputs() []Rel {
	return []Rel{ar.input}
}

func (ar *AggregateRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	aggregate := *ar
	aggregate.input = newInputs[0]
	return &aggregate, nil
}

func (ar *AggregateRel) rewriteAggregateFunc(rewriteFunc RewriteFunc, f *expr.AggregateFunction) (*expr.AggregateFunction, error) {
	if f == nil {
		return f, nil
	}
	newF := f.Clone()
	argsAreEqual := true
	for i := 0; i < f.NArgs(); i++ {
		arg := f.Arg(i)
		if exp, ok := arg.(expr.Expression); ok {
			var newExp expr.Expression
			var err error
			if newExp, err = rewriteFunc(exp); err != nil {
				return nil, err
			}
			newF.SetArg(i, newExp)
			argsAreEqual = argsAreEqual && exp == newExp
		}
	}
	if argsAreEqual {
		return f, nil
	}
	return newF, nil
}

func (ar *AggregateRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	var err error
	groupsAreEqual := true
	newGroupingExpressions := make([]expr.Expression, len(ar.groupingExpressions))
	for i, e := range ar.groupingExpressions {
		if newGroupingExpressions[i], err = rewriteFunc(e); err != nil {
			return nil, err
		}
		groupsAreEqual = groupsAreEqual && newGroupingExpressions[i] == e
	}
	measuresAreEqual := true
	newMeasures := make([]AggRelMeasure, len(ar.measures))
	for i, m := range ar.measures {
		if newMeasures[i].filter, err = rewriteFunc(m.filter); err != nil {
			return nil, err
		}
		if newMeasures[i].measure, err = ar.rewriteAggregateFunc(rewriteFunc, m.measure); err != nil {
			return nil, err
		}
		measuresAreEqual = measuresAreEqual && newMeasures[i].filter == m.filter && newMeasures[i].measure == m.measure
	}
	if groupsAreEqual && measuresAreEqual && newInputs[0] == ar.input {
		return ar, nil
	}
	aggregate := *ar
	aggregate.input = newInputs[0]
	aggregate.groupingExpressions = newGroupingExpressions
	aggregate.measures = newMeasures
	return &aggregate, nil
}

func (ar *AggregateRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(ar, mapping)
}

func (ar *AggregateRel) SetInput(input Rel) {
	ar.input = input
}

func (ar *AggregateRel) SetMeasures(measures []AggRelMeasure) {
	ar.measures = measures
}

func (ar *AggregateRel) SetGroupingExpressions(groupingExpressions []expr.Expression) {
	ar.groupingExpressions = groupingExpressions
}

func (ar *AggregateRel) SetGroupingReferences(groupingReferences [][]uint32) {
	ar.groupingReferences = groupingReferences
}

// SortRel is an ORDER BY relational operator, describing a base relation,
// it includes a list of fields to sort on.
type SortRel struct {
	RelCommon

	input        Rel
	sorts        []expr.SortField
	advExtension *extensions.AdvancedExtension
}

func (sr *SortRel) directOutputSchema() types.RecordType { return sr.input.RecordType() }
func (sr *SortRel) RecordType() types.RecordType {
	return sr.remap(sr.directOutputSchema())
}
func (sr *SortRel) Input() Rel              { return sr.input }
func (sr *SortRel) Sorts() []expr.SortField { return sr.sorts }
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

func (sr *SortRel) GetInputs() []Rel {
	return []Rel{sr.input}
}

func (sr *SortRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	proj := *sr
	proj.input = newInputs[0]
	return &proj, nil
}

func (sr *SortRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	var err error
	sortExpressionsAreEqual := true
	sorts := make([]expr.SortField, len(sr.sorts))
	for i, s := range sr.sorts {
		if sorts[i].Expr, err = rewriteFunc(s.Expr); err != nil {
			return nil, err
		}
		sortExpressionsAreEqual = sortExpressionsAreEqual && sorts[i].Expr == s.Expr
		sorts[i].Kind = s.Kind
	}
	if sortExpressionsAreEqual && newInputs[0] == sr.input {
		return sr, nil
	}
	sort := *sr
	sort.input = newInputs[0]
	sort.sorts = sorts
	return &sort, nil
}

func (sr *SortRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(sr, mapping)
}

// FilterRel is a relational operator capturing simple filters (
// as in the WHERE clause of a SQL query).
type FilterRel struct {
	RelCommon

	input        Rel
	cond         expr.Expression
	advExtension *extensions.AdvancedExtension
}

func (fr *FilterRel) directOutputSchema() types.RecordType { return fr.input.RecordType() }
func (fr *FilterRel) RecordType() types.RecordType {
	return fr.remap(fr.directOutputSchema())
}
func (fr *FilterRel) Input() Rel                 { return fr.input }
func (fr *FilterRel) Condition() expr.Expression { return fr.cond }
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

func (fr *FilterRel) GetInputs() []Rel {
	return []Rel{fr.input}
}

func (fr *FilterRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	proj := *fr
	proj.input = newInputs[0]
	return &proj, nil
}

func (fr *FilterRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	cond, err := rewriteFunc(fr.cond)
	if err != nil {
		return nil, err
	}
	if newInputs[0] == fr.input && cond == fr.cond {
		return fr, nil
	}
	filter := *fr
	filter.input = newInputs[0]
	filter.cond = cond
	return &filter, nil
}

func (fr *FilterRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(fr, mapping)
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

func (s *SetRel) directOutputSchema() types.RecordType { return s.inputs[0].RecordType() }
func (s *SetRel) RecordType() types.RecordType {
	return s.remap(s.directOutputSchema())
}
func (s *SetRel) Inputs() []Rel { return s.inputs }
func (s *SetRel) Op() SetOp     { return s.op }
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

func (s *SetRel) GetInputs() []Rel {
	return s.inputs
}

func (s *SetRel) Copy(newInputs ...Rel) (Rel, error) {
	set := *s
	set.inputs = newInputs
	return &set, nil
}

func (s *SetRel) CopyWithExpressionRewrite(_ RewriteFunc, newInputs ...Rel) (Rel, error) {
	if slices.Equal(newInputs, s.GetInputs()) {
		return s, nil
	}
	return s.Copy(newInputs...)
}

func (s *SetRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(s, mapping)
}

// ExtensionSingleRel is a stub to support extensions with a single input.
type ExtensionSingleRel struct {
	RelCommon

	input  Rel
	detail *anypb.Any
}

func (es *ExtensionSingleRel) directOutputSchema() types.RecordType {
	return es.input.RecordType()
}
func (es *ExtensionSingleRel) RecordType() types.RecordType {
	return es.remap(es.directOutputSchema())
}
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

func (es *ExtensionSingleRel) GetInputs() []Rel {
	return []Rel{es.input}
}

func (es *ExtensionSingleRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	proj := *es
	proj.input = newInputs[0]
	return &proj, nil
}

func (es *ExtensionSingleRel) CopyWithExpressionRewrite(_ RewriteFunc, newInputs ...Rel) (Rel, error) {
	if slices.Equal(newInputs, es.GetInputs()) {
		return es, nil
	}
	return es.Copy(newInputs...)
}

func (es *ExtensionSingleRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(es, mapping)
}

// ExtensionLeafRel is a stub to support extensions with zero inputs.
type ExtensionLeafRel struct {
	RelCommon

	detail *anypb.Any
}

func (el *ExtensionLeafRel) directOutputSchema() types.RecordType { return types.RecordType{} }
func (el *ExtensionLeafRel) RecordType() types.RecordType {
	return el.remap(el.directOutputSchema())
}
func (el *ExtensionLeafRel) Detail() *anypb.Any { return el.detail }

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

func (el *ExtensionLeafRel) GetInputs() []Rel {
	return []Rel{}
}

func (el *ExtensionLeafRel) Copy(_ ...Rel) (Rel, error) {
	return el, nil
}

func (el *ExtensionLeafRel) CopyWithExpressionRewrite(_ RewriteFunc, _ ...Rel) (Rel, error) {
	return el, nil
}

func (el *ExtensionLeafRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(el, mapping)
}

// ExtensionMultiRel is a stub to support extensions with multiple inputs.
type ExtensionMultiRel struct {
	RelCommon

	inputs []Rel
	detail *anypb.Any
}

func (em *ExtensionMultiRel) directOutputSchema() types.RecordType { return types.RecordType{} }
func (em *ExtensionMultiRel) RecordType() types.RecordType {
	return em.remap(em.directOutputSchema())
}
func (em *ExtensionMultiRel) Inputs() []Rel      { return em.inputs }
func (em *ExtensionMultiRel) Detail() *anypb.Any { return em.detail }

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

func (em *ExtensionMultiRel) GetInputs() []Rel {
	return em.inputs
}

func (em *ExtensionMultiRel) Copy(newInputs ...Rel) (Rel, error) {
	proj := *em
	proj.inputs = newInputs
	return &proj, nil
}

func (em *ExtensionMultiRel) CopyWithExpressionRewrite(_ RewriteFunc, newInputs ...Rel) (Rel, error) {
	if slices.Equal(newInputs, em.GetInputs()) {
		return em, nil
	}
	return em.Copy(newInputs...)
}

func (em *ExtensionMultiRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(em, mapping)
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

func (hr *HashJoinRel) directOutputSchema() types.RecordType {
	return hr.left.RecordType().Concat(hr.right.RecordType())
}

func (hr *HashJoinRel) RecordType() types.RecordType {
	return hr.remap(hr.directOutputSchema())
}

func (hr *HashJoinRel) Left() Rel                         { return hr.left }
func (hr *HashJoinRel) Right() Rel                        { return hr.right }
func (hr *HashJoinRel) LeftKeys() []*expr.FieldReference  { return hr.leftKeys }
func (hr *HashJoinRel) RightKeys() []*expr.FieldReference { return hr.rightKeys }
func (hr *HashJoinRel) PostJoinFilter() expr.Expression {
	if hr.postJoinFilter == nil {
		return defFilter
	}
	return hr.postJoinFilter
}
func (hr *HashJoinRel) Type() HashMergeJoinType { return hr.joinType }
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

	ret := &proto.Rel_HashJoin{
		HashJoin: &proto.HashJoinRel{
			Common:            hr.toProto(),
			Left:              hr.left.ToProto(),
			Right:             hr.right.ToProto(),
			LeftKeys:          keysLeft,
			RightKeys:         keysRight,
			Type:              proto.HashJoinRel_JoinType(hr.joinType),
			AdvancedExtension: hr.advExtension,
		},
	}

	if hr.postJoinFilter != nil {
		ret.HashJoin.PostJoinFilter = hr.postJoinFilter.ToProto()
	}

	return &proto.Rel{
		RelType: ret}
}

func (hr *HashJoinRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: hr.ToProto(),
		},
	}
}

func (hr *HashJoinRel) GetInputs() []Rel {
	return []Rel{hr.left, hr.right}
}

func (hr *HashJoinRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	join := *hr
	join.left, join.right = newInputs[0], newInputs[1]
	return &join, nil
}

func (hr *HashJoinRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	newExpr, err := rewriteFunc(hr.postJoinFilter)
	if err != nil {
		return nil, err
	}
	if newExpr == hr.postJoinFilter && slices.Equal(newInputs, hr.GetInputs()) {
		return hr, nil
	}
	join := *hr
	join.left, join.right = newInputs[0], newInputs[1]
	join.postJoinFilter = newExpr
	return &join, nil
}

func (hr *HashJoinRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(hr, mapping)
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

func (mr *MergeJoinRel) directOutputSchema() types.RecordType {
	return mr.left.RecordType().Concat(mr.right.RecordType())
}

func (mr *MergeJoinRel) RecordType() types.RecordType {
	return mr.remap(mr.directOutputSchema())
}

func (mr *MergeJoinRel) Left() Rel                         { return mr.left }
func (mr *MergeJoinRel) Right() Rel                        { return mr.right }
func (mr *MergeJoinRel) LeftKeys() []*expr.FieldReference  { return mr.leftKeys }
func (mr *MergeJoinRel) RightKeys() []*expr.FieldReference { return mr.rightKeys }
func (mr *MergeJoinRel) PostJoinFilter() expr.Expression {
	if mr.postJoinFilter == nil {
		return defFilter
	}
	return mr.postJoinFilter
}
func (mr *MergeJoinRel) Type() HashMergeJoinType { return mr.joinType }
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

	ret := &proto.Rel_MergeJoin{
		MergeJoin: &proto.MergeJoinRel{
			Common:            mr.toProto(),
			Left:              mr.left.ToProto(),
			Right:             mr.right.ToProto(),
			LeftKeys:          keysLeft,
			RightKeys:         keysRight,
			Type:              proto.MergeJoinRel_JoinType(mr.joinType),
			AdvancedExtension: mr.advExtension,
		},
	}

	if mr.postJoinFilter != nil {
		ret.MergeJoin.PostJoinFilter = mr.postJoinFilter.ToProto()
	}

	return &proto.Rel{
		RelType: ret}
}

func (mr *MergeJoinRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: mr.ToProto(),
		},
	}
}

func (mr *MergeJoinRel) GetInputs() []Rel {
	return []Rel{mr.left, mr.right}
}

func (mr *MergeJoinRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	merge := *mr
	merge.left, merge.right = newInputs[0], newInputs[1]
	return &merge, nil
}

func (mr *MergeJoinRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 2 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	newExpr, err := rewriteFunc(mr.postJoinFilter)
	if err != nil {
		return nil, err
	}
	if newExpr == mr.postJoinFilter && slices.Equal(newInputs, mr.GetInputs()) {
		return mr, nil
	}
	merge := *mr
	merge.left, merge.right = newInputs[0], newInputs[1]
	merge.postJoinFilter = newExpr
	return &merge, nil
}

func (mr *MergeJoinRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(mr, mapping)
}

type WriteOp = proto.WriteRel_WriteOp

const (
	WriteOpUnspecified = proto.WriteRel_WRITE_OP_UNSPECIFIED
	WriteOpInsert      = proto.WriteRel_WRITE_OP_INSERT
	WriteOpDelete      = proto.WriteRel_WRITE_OP_DELETE
	WriteOpUpdate      = proto.WriteRel_WRITE_OP_UPDATE
	WriteOpCTAS        = proto.WriteRel_WRITE_OP_CTAS
)

type OutputMode = proto.WriteRel_OutputMode

const (
	OutputModeUnspecified     = proto.WriteRel_OUTPUT_MODE_UNSPECIFIED
	OutputModeNoOutput        = proto.WriteRel_OUTPUT_MODE_NO_OUTPUT
	OutputModeModifiedRecords = proto.WriteRel_OUTPUT_MODE_MODIFIED_RECORDS
)

// NamedTableWriteRel is a relational operator that writes data to a table. The list of strings
// that make up the names are to represent namespacing (e.g. mydb.mytable).
// This assumes a shared catalog between systems exchanging a message.
// op as WriteOpCTAS is a special case of write operation where the output is written to a new table.
type NamedTableWriteRel struct {
	RelCommon

	names        []string
	advExtension *extensions.AdvancedExtension

	tableSchema types.NamedStruct
	op          WriteOp
	input       Rel
	outputMode  OutputMode
}

func (wr *NamedTableWriteRel) directOutputSchema() types.RecordType {
	switch wr.outputMode {
	case OutputModeNoOutput:
		return types.RecordType{}
	case OutputModeModifiedRecords:
		return *types.NewRecordTypeFromStruct(wr.tableSchema.Struct)
	case OutputModeUnspecified:
		panic("output mode not specified")
	}
	return types.RecordType{}
}

func (wr *NamedTableWriteRel) RecordType() types.RecordType {
	return wr.remap(wr.directOutputSchema())
}

func (wr *NamedTableWriteRel) Names() []string { return wr.names }
func (wr *NamedTableWriteRel) NamedTableAdvancedExtension() *extensions.AdvancedExtension {
	return wr.advExtension
}

func (wr *NamedTableWriteRel) TableSchema() types.NamedStruct {
	return wr.tableSchema
}
func (wr *NamedTableWriteRel) Op() WriteOp { return wr.op }
func (wr *NamedTableWriteRel) Input() Rel  { return wr.input }
func (wr *NamedTableWriteRel) OutputMode() OutputMode {
	return wr.outputMode
}

func (wr *NamedTableWriteRel) ToProto() *proto.Rel {
	return &proto.Rel{
		RelType: &proto.Rel_Write{
			Write: &proto.WriteRel{
				Common: wr.toProto(),
				WriteType: &proto.WriteRel_NamedTable{
					NamedTable: &proto.NamedObjectWrite{
						Names:             wr.names,
						AdvancedExtension: wr.advExtension,
					},
				},
				TableSchema: wr.tableSchema.ToProto(),
				Op:          wr.op,
				Input:       wr.input.ToProto(),
			},
		},
	}
}

func (wr *NamedTableWriteRel) ToProtoPlanRel() *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{
			Rel: wr.ToProto(),
		},
	}
}

func (wr *NamedTableWriteRel) GetInputs() []Rel {
	return []Rel{wr.input}
}

func (wr *NamedTableWriteRel) Copy(newInputs ...Rel) (Rel, error) {
	if len(newInputs) != 1 {
		return nil, substraitgo.ErrInvalidInputCount
	}
	write := *wr
	write.input = newInputs[0]
	return &write, nil
}

func (wr *NamedTableWriteRel) CopyWithExpressionRewrite(_ RewriteFunc, newInputs ...Rel) (Rel, error) {
	if slices.Equal(newInputs, wr.GetInputs()) {
		return wr, nil
	}
	return wr.Copy(newInputs...)
}

func (wr *NamedTableWriteRel) Remap(mapping ...int32) (Rel, error) {
	return RemapHelper(wr, mapping)
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
	_ Rel = (*NamedTableWriteRel)(nil)

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
	_ SingleInputRel = (*NamedTableWriteRel)(nil)
)
