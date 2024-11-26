// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
	"golang.org/x/exp/slices"
)

// Builder is the base object for constructing the various elements of a plan.
// The intent is to create a single builder and then utilize it for all
// necessary constructions while building a full plan.
//
// Any extensions that are referenced for functions or user defined types, etc.
// will be added to the internal extension set so that the final Plan when
// constructed will have the appropriate extension anchors and definitions.
// This will maintain consistency across the plan for the user without them
// having to manually do so.
type Builder interface {
	// GetFunctionRef retrieves the function anchor reference for a given
	// function identified by its namespace and function name. This also
	// ensures that any plans built from this builder will contain this
	// function anchor in its extensions section.
	GetFunctionRef(nameSpace, key string) types.FunctionRef

	// Construct a user-defined type from the extension namespace and typename,
	// along with optional type parameters. It will add the type to the internal
	// extension set if it doesn't already exist and assign it a type reference.
	UserDefinedType(nameSpace, typeName string, params ...types.TypeParam) types.UserDefinedType
	// RootFieldRef constructs a Root Field Reference to the column of the input
	// relation indicated by the passed in index. This will ensure the output
	// type is properly propagated based on the reference.
	//
	// Will return an error if the index is < 0 or > the number of output fields.
	RootFieldRef(input Rel, index int32) (*expr.FieldReference, error)
	// JoinedRecordFieldRef constructs a root field reference for the full tuple of
	// the inputs to a join, to construct an expression that is viable to use as
	// the condition or post join filter for a join relation.
	JoinedRecordFieldRef(left, right Rel, index int32) (*expr.FieldReference, error)
	// ScalarFn constructs a ScalarFunction from the passed in namespace and
	// function name key. This is equivalent to calling expr.NewScalarFunc using
	// the builder's extension registry. An error will be returned if the indicated
	// function was not already in the extension collection the builder was created
	// with or if the arguments of the function don't match the provided argument
	// types.
	ScalarFn(nameSpace, key string, opts []*types.FunctionOption, args ...types.FuncArg) (*expr.ScalarFunction, error)
	// AggregateFn constructs an AggregateFunction from the passed in namespace and
	// function name key. This is equivalent to calling expr.NewAggregateFunc using
	// the builder's extension registry. An error will be returned if the indicated
	// function was not already in the extension collection the builder was created
	// with or if the arguments of the function don't match the provided argument
	// types.
	AggregateFn(nameSpace, key string, opts []*types.FunctionOption, args ...types.FuncArg) (*expr.AggregateFunction, error)
	// SortFields is a convenience method to construct a list of sort fields
	// from the column indices of an existing relation. This will return an error
	// if any of the indices are < 0 or > the number of columns in the output
	// of the relation. This will use types.SortAscNullsLast as the sort kind
	// for each field in the returned slice.
	SortFields(input Rel, indices ...int32) ([]expr.SortField, error)
	// Measure is a convenience method to construct the input for an Aggregate Rel
	// Consisting of the provided aggregate function and optional filter expression.
	Measure(measure *expr.AggregateFunction, filter expr.Expression) AggRelMeasure

	// The Remap variant for each method produces that type of relation
	// with an optional output mapping to reorder or exclude specific columns
	// from the output.

	Project(input Rel, exprs ...expr.Expression) (*ProjectRel, error)
	ProjectRemap(input Rel, remap []int32, exprs ...expr.Expression) (*ProjectRel, error)
	AggregateColumnsRemap(input Rel, remap []int32, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error)
	AggregateColumns(input Rel, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error)
	AggregateExprsRemap(input Rel, remap []int32, measures []AggRelMeasure, groups ...[]expr.Expression) (*AggregateRel, error)
	AggregateExprs(input Rel, measures []AggRelMeasure, groups ...[]expr.Expression) (*AggregateRel, error)
	CreateTableAsSelectRemap(input Rel, remap []int32, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error)
	CreateTableAsSelect(input Rel, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error)
	CrossRemap(left, right Rel, remap []int32) (*CrossRel, error)
	Cross(left, right Rel) (*CrossRel, error)
	// FetchRemap constructs a fetch relation providing an offset (skipping some
	// number of rows) and a count (restricting output to a maximum number of
	// rows).  If count is FETCH_COUNT_ALL_RECORDS (-1) all records will be
	// returned.  Similar to Fetch but allows for reordering and restricting the
	// returned columns.
	FetchRemap(input Rel, offset, count int64, remap []int32) (*FetchRel, error)
	// Fetch constructs a fetch relation providing an offset (skipping some number of
	// rows) and a count (restricting output to a maximum number of rows).  If count
	// is FETCH_COUNT_ALL_RECORDS (-1) all records will be returned.
	Fetch(input Rel, offset, count int64) (*FetchRel, error)
	FilterRemap(input Rel, condition expr.Expression, remap []int32) (*FilterRel, error)
	Filter(input Rel, condition expr.Expression) (*FilterRel, error)
	JoinAndFilterRemap(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType, remap []int32) (*JoinRel, error)
	JoinAndFilter(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType) (*JoinRel, error)
	JoinRemap(left, right Rel, condition expr.Expression, joinType JoinType, remap []int32) (*JoinRel, error)
	Join(left, right Rel, condition expr.Expression, joinType JoinType) (*JoinRel, error)
	NamedScanRemap(tableName []string, schema types.NamedStruct, remap []int32) (*NamedTableReadRel, error)
	NamedScan(tableName []string, schema types.NamedStruct) *NamedTableReadRel
	NamedWriteRemap(input Rel, op WriteOp, tableName []string, schema types.NamedStruct, remap []int32) (*NamedTableWriteRel, error)
	// NamedInsert inserts data from the input relation into a named table.
	NamedInsert(input Rel, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error)
	// NamedDelete deletes rows from a specified named table based on the
	// provided input relation, which typically includes conditions that filter
	// the rows to delete.
	NamedDelete(input Rel, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error)
	VirtualTableRemap(fields []string, remap []int32, values ...expr.StructLiteralValue) (*VirtualTableReadRel, error)
	VirtualTable(fields []string, values ...expr.StructLiteralValue) (*VirtualTableReadRel, error)
	SortRemap(input Rel, remap []int32, sorts ...expr.SortField) (*SortRel, error)
	Sort(input Rel, sorts ...expr.SortField) (*SortRel, error)
	SetRemap(op SetOp, remap []int32, inputs ...Rel) (*SetRel, error)
	Set(op SetOp, inputs ...Rel) (*SetRel, error)

	// Plan constructs a new plan with the provided root relation and optionally
	// other relations. It will use the current substrait version of this
	// library as the plan substrait version.
	Plan(root Rel, rootNames []string, others ...Rel) (*Plan, error)
	// PlanWithTypes is the same as Plan, only it provides the ability to set
	// the list of expectedTypeURLs that indicate the different protobuf types
	// that may be in use with this plan for advanced extensions, optimizations,
	// and so on.
	PlanWithTypes(root Rel, rootNames []string, expectedTypeURLs []string, others ...Rel) (*Plan, error)
}

const FETCH_COUNT_ALL_RECORDS = -1

func NewBuilderDefault() Builder {
	return NewBuilder(&extensions.DefaultCollection)
}

func NewBuilder(c *extensions.Collection) Builder {
	set := extensions.NewSet()
	return &builder{
		ext:    c,
		extSet: set,
		reg:    expr.NewExtensionRegistry(set, c),
	}
}

var (
	errOutputMappingOutOfRange = fmt.Errorf("%w: output mapping index out of range", substraitgo.ErrInvalidRel)
	errNilInputRel             = fmt.Errorf("%w: input Relation must not be nil", substraitgo.ErrInvalidRel)
)

type builder struct {
	ext    *extensions.Collection
	extSet extensions.Set

	reg expr.ExtensionRegistry
}

func (b *builder) GetFunctionRef(nameSpace, key string) types.FunctionRef {
	return types.FunctionRef(b.extSet.GetFuncAnchor(extensions.ID{URI: nameSpace, Name: key}))
}

func (b *builder) UserDefinedType(nameSpace, typeName string, params ...types.TypeParam) types.UserDefinedType {
	id := extensions.ID{URI: nameSpace, Name: typeName}
	return types.UserDefinedType{
		Nullability:    types.NullabilityNullable,
		TypeReference:  b.extSet.GetTypeAnchor(id),
		TypeParameters: params,
	}
}

func (b *builder) JoinedRecordFieldRef(left, right Rel, index int32) (*expr.FieldReference, error) {
	baseTypes := append(left.RecordType().Types, right.RecordType().Types...)
	if index < 0 || index > int32(len(baseTypes)) {
		return nil, fmt.Errorf("%w: cannot create field ref index %d, only %d fields to reference",
			substraitgo.ErrInvalidArg, index, len(baseTypes))
	}

	return expr.NewRootFieldRef(expr.NewStructFieldRef(index), &types.RecordType{Types: baseTypes})
}

func (b *builder) RootFieldRef(input Rel, index int32) (*expr.FieldReference, error) {
	base := input.RecordType()
	if index < 0 || index > int32(len(base.Types)) {
		return nil, fmt.Errorf("%w: cannot create field ref index %d, only %d fields in rel",
			substraitgo.ErrInvalidArg, index, len(base.Types))
	}

	return expr.NewRootFieldRef(expr.NewStructFieldRef(index), &base)
}

func (b *builder) ScalarFn(nameSpace, key string, opts []*types.FunctionOption, args ...types.FuncArg) (*expr.ScalarFunction, error) {
	id := extensions.ID{URI: nameSpace, Name: key}
	return expr.NewScalarFunc(b.reg, id, opts, args...)
}

func (b *builder) AggregateFn(nameSpace, key string, opts []*types.FunctionOption, args ...types.FuncArg) (*expr.AggregateFunction, error) {
	id := extensions.ID{URI: nameSpace, Name: key}
	return expr.NewAggregateFunc(b.reg, id, opts,
		types.AggInvocationAll, types.AggPhaseInitialToResult, nil, args...)
}

func (b *builder) Project(input Rel, exprs ...expr.Expression) (*ProjectRel, error) {
	return b.ProjectRemap(input, nil, exprs...)
}

func (b *builder) ProjectRemap(input Rel, remap []int32, exprs ...expr.Expression) (*ProjectRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	if len(exprs) == 0 {
		return nil, fmt.Errorf("%w: must provide at least one expression for project relation", substraitgo.ErrInvalidRel)
	}

	noutput := int32(len(input.RecordType().Types) + len(exprs))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &ProjectRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		exprs:     exprs,
	}, nil
}

func (b *builder) Measure(measure *expr.AggregateFunction, filter expr.Expression) AggRelMeasure {
	return AggRelMeasure{
		measure: measure,
		filter:  filter,
	}
}

func (b *builder) AggregateColumnsRemap(input Rel, remap []int32, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	if (len(measures) + len(groupByCols)) == 0 {
		return nil, fmt.Errorf("%w: must have at least one grouping expression or measure for AggregateRel",
			substraitgo.ErrInvalidRel)
	}

	exprs := make([][]expr.Expression, len(groupByCols))
	for i, c := range groupByCols {
		ref, err := b.RootFieldRef(input, c)
		if err != nil {
			return nil, err
		}
		exprs[i] = []expr.Expression{ref}
	}

	noutput := int32(len(measures) + len(groupByCols))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &AggregateRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		groups:    exprs,
		measures:  measures,
	}, nil
}

func (b *builder) AggregateColumns(input Rel, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error) {
	return b.AggregateColumnsRemap(input, nil, measures, groupByCols...)
}

func (b *builder) AggregateExprsRemap(input Rel, remap []int32, measures []AggRelMeasure, groups ...[]expr.Expression) (*AggregateRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	if (len(measures) + len(groups)) == 0 {
		return nil, fmt.Errorf("%w: must have at least one grouping expression or measure for AggregateRel",
			substraitgo.ErrInvalidRel)
	}

	if slices.ContainsFunc(groups, func(exlist []expr.Expression) bool {
		return len(exlist) == 0 || slices.ContainsFunc(exlist, func(e expr.Expression) bool { return e == nil })
	}) {
		return nil, fmt.Errorf("%w: groupings cannot contain empty expression list or nil expression", substraitgo.ErrInvalidRel)
	}

	noutput := int32(len(measures) + len(groups))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &AggregateRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		groups:    groups,
		measures:  measures,
	}, nil
}

func (b *builder) AggregateExprs(input Rel, measures []AggRelMeasure, groups ...[]expr.Expression) (*AggregateRel, error) {
	return b.AggregateExprsRemap(input, nil, measures, groups...)
}

func (b *builder) CreateTableAsSelectRemap(input Rel, remap []int32, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	nOutput := int32(len(input.RecordType().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= nOutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &NamedTableWriteRel{
		RelCommon:   RelCommon{mapping: remap},
		names:       tableName,
		tableSchema: schema,
		op:          WriteOpCTAS,
		input:       input,
		outputMode:  OutputModeModifiedRecords,
	}, nil
}

func (b *builder) CreateTableAsSelect(input Rel, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error) {
	return b.CreateTableAsSelectRemap(input, nil, tableName, schema)
}

func (b *builder) CrossRemap(left, right Rel, remap []int32) (*CrossRel, error) {
	if left == nil || right == nil {
		return nil, errNilInputRel
	}

	noutput := int32(len(left.RecordType().Types) + len(right.RecordType().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &CrossRel{
		RelCommon: RelCommon{mapping: remap},
		left:      left, right: right,
	}, nil
}

func (b *builder) Cross(left, right Rel) (*CrossRel, error) {
	return b.CrossRemap(left, right, nil)
}

func (b *builder) FetchRemap(input Rel, offset, count int64, remap []int32) (*FetchRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	noutput := int32(len(input.RecordType().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &FetchRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		offset:    offset, count: count,
	}, nil
}

func (b *builder) Fetch(input Rel, offset, count int64) (*FetchRel, error) {
	return b.FetchRemap(input, offset, count, nil)
}

func (b *builder) FilterRemap(input Rel, condition expr.Expression, remap []int32) (*FilterRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	if condition == nil {
		return nil, fmt.Errorf("%w: cannot use nil condition in filter relation",
			substraitgo.ErrInvalidRel)
	}

	if !condition.GetType().WithNullability(types.NullabilityUnspecified).Equals(&types.BooleanType{}) {
		return nil, fmt.Errorf("%w: condition for Filter Relation must yield boolean, not %s",
			substraitgo.ErrInvalidArg, condition.GetType())
	}

	noutput := int32(len(input.directOutputSchema().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &FilterRel{
		RelCommon: RelCommon{
			mapping: remap,
		},
		input: input,
		cond:  condition,
	}, nil
}

func (b *builder) Filter(input Rel, condition expr.Expression) (*FilterRel, error) {
	return b.FilterRemap(input, condition, nil)
}

func (b *builder) JoinAndFilterRemap(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType, remap []int32) (*JoinRel, error) {
	if left == nil || right == nil {
		return nil, errNilInputRel
	}

	if condition == nil {
		return nil, fmt.Errorf("%w: cannot use nil condition in filter relation",
			substraitgo.ErrInvalidRel)
	}

	if !condition.GetType().WithNullability(types.NullabilityUnspecified).Equals(&types.BooleanType{}) {
		return nil, fmt.Errorf("%w: condition for Join Relation must yield boolean, not %s",
			substraitgo.ErrInvalidArg, condition.GetType())
	}

	if joinType == JoinTypeUnspecified {
		return nil, fmt.Errorf("%w: join type must not be unspecified for Join relations",
			substraitgo.ErrInvalidArg)
	}

	if postJoinFilter != nil {
		if !postJoinFilter.GetType().WithNullability(types.NullabilityUnspecified).Equals(&types.BooleanType{}) {
			return nil, fmt.Errorf("%w: post join filter must be either nil or yield a boolean, not %s",
				substraitgo.ErrInvalidArg, postJoinFilter.GetType())
		}
	}

	out := &JoinRel{
		RelCommon: RelCommon{mapping: remap},
		left:      left, right: right,
		expr: condition, postJoinFilter: postJoinFilter,
		joinType: joinType,
	}

	noutput := int32(len(out.directOutputSchema().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return out, nil
}

func (b *builder) JoinAndFilter(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType) (*JoinRel, error) {
	return b.JoinAndFilterRemap(left, right, condition, postJoinFilter, joinType, nil)
}

func (b *builder) JoinRemap(left, right Rel, condition expr.Expression, joinType JoinType, remap []int32) (*JoinRel, error) {
	return b.JoinAndFilterRemap(left, right, condition, nil, joinType, remap)
}

func (b *builder) Join(left, right Rel, condition expr.Expression, joinType JoinType) (*JoinRel, error) {
	return b.JoinAndFilterRemap(left, right, condition, nil, joinType, nil)
}

func (b *builder) NamedWriteRemap(input Rel, op WriteOp, tableName []string, schema types.NamedStruct, remap []int32) (*NamedTableWriteRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	nOutput := int32(len(input.RecordType().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= nOutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	return &NamedTableWriteRel{
		RelCommon:   RelCommon{mapping: remap},
		names:       tableName,
		tableSchema: schema,
		op:          op,
		input:       input,
		outputMode:  OutputModeNoOutput,
	}, nil
}

func (b *builder) NamedInsert(input Rel, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error) {
	return b.NamedWriteRemap(input, WriteOpInsert, tableName, schema, nil)
}

func (b *builder) NamedDelete(input Rel, tableName []string, schema types.NamedStruct) (*NamedTableWriteRel, error) {
	return b.NamedWriteRemap(input, WriteOpDelete, tableName, schema, nil)
}

func (b *builder) NamedScanRemap(tableName []string, schema types.NamedStruct, remap []int32) (*NamedTableReadRel, error) {
	noutput := int32(len(schema.Struct.Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, fmt.Errorf("%w: output mapping index out of range",
				substraitgo.ErrInvalidRel)
		}
	}

	return &NamedTableReadRel{
		baseReadRel: baseReadRel{
			RelCommon: RelCommon{
				mapping: remap,
			},
			baseSchema: schema,
		},
		names: tableName,
	}, nil
}

func (b *builder) NamedScan(tableName []string, schema types.NamedStruct) *NamedTableReadRel {
	n, _ := b.NamedScanRemap(tableName, schema, nil)
	return n
}

func (b *builder) VirtualTableRemap(fieldNames []string, remap []int32, values ...expr.StructLiteralValue) (*VirtualTableReadRel, error) {
	// convert Literal to Expression
	exprs := make([]expr.VirtualTableExpressionValue, 0)
	for _, row := range values {
		rowExpr := make(expr.VirtualTableExpressionValue, 0)
		for _, col := range row {
			rowExpr = append(rowExpr, col)
		}
		exprs = append(exprs, rowExpr)
	}
	return b.VirtualTableFromExprRemap(fieldNames, remap, exprs...)
}

func (b *builder) VirtualTableFromExprRemap(fieldNames []string, remap []int32, values ...expr.VirtualTableExpressionValue) (*VirtualTableReadRel, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("%w: must provide at least one set of values for virtual table", substraitgo.ErrInvalidRel)
	}

	nfields := len(fieldNames)
	if len(values[0]) != nfields {
		return nil, fmt.Errorf("%w: mismatched number of fields (%d) and literal values (%d) in virtual table",
			substraitgo.ErrInvalidRel, nfields, len(values[0]))
	}

	for _, idx := range remap {
		if idx < 0 || idx >= int32(nfields) {
			return nil, fmt.Errorf("%w: output mapping index out of range",
				substraitgo.ErrInvalidRel)
		}
	}

	typeList := make([]types.Type, nfields)
	for i, v := range values[0] {
		typeList[i] = v.GetType()
	}

	for _, row := range values[1:] {
		for j, v := range row {
			if !typeList[j].Equals(v.GetType()) {
				return nil, fmt.Errorf("%w: inconsistent literal types for virtual table, found %s in col %d, expected %s",
					substraitgo.ErrInvalidRel, v.GetType(), j, typeList[j].GetType())
			}
		}
	}

	baseSchema := types.NamedStruct{
		Names: fieldNames,
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       typeList,
		},
	}

	return &VirtualTableReadRel{
		baseReadRel: baseReadRel{
			RelCommon:  RelCommon{mapping: remap},
			baseSchema: baseSchema,
		},
		values: values,
	}, nil
}

func (b *builder) VirtualTable(fields []string, values ...expr.StructLiteralValue) (*VirtualTableReadRel, error) {
	return b.VirtualTableRemap(fields, nil, values...)
}

func (b *builder) SortRemap(input Rel, remap []int32, sorts ...expr.SortField) (*SortRel, error) {
	if input == nil {
		return nil, errNilInputRel
	}

	noutput := int32(len(input.RecordType().Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	if len(sorts) == 0 {
		return nil, fmt.Errorf("%w: must provide at least one SortField for sort relation", substraitgo.ErrInvalidRel)
	}

	return &SortRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		sorts:     sorts,
	}, nil
}

func (b *builder) Sort(input Rel, sorts ...expr.SortField) (*SortRel, error) {
	return b.SortRemap(input, nil, sorts...)
}

func (b *builder) SortFields(input Rel, indices ...int32) ([]expr.SortField, error) {
	out := make([]expr.SortField, len(indices))
	for i, idx := range indices {
		ref, err := b.RootFieldRef(input, idx)
		if err != nil {
			return nil, err
		}
		out[i] = expr.SortField{Expr: ref, Kind: types.SortAscNullsLast}
	}
	return out, nil
}

func (b *builder) SetRemap(op SetOp, remap []int32, inputs ...Rel) (*SetRel, error) {
	if op == SetOpUnspecified {
		return nil, fmt.Errorf("%w: operation for set relation must not be unspecified", substraitgo.ErrInvalidArg)
	}

	if len(inputs) < 2 {
		return nil, fmt.Errorf("%w: must have at least 2 relations for a set relation, got %d",
			substraitgo.ErrInvalidRel, len(inputs))
	}

	for _, in := range inputs {
		if in == nil {
			return nil, errNilInputRel
		}
	}

	output := inputs[0].RecordType()

	noutput := int32(len(output.Types))
	for _, idx := range remap {
		if idx < 0 || idx >= noutput {
			return nil, errOutputMappingOutOfRange
		}
	}

	for _, in := range inputs[1:] {
		t := in.RecordType()
		if !output.Equals(&t) {
			return nil, fmt.Errorf("%w: mismatched column types in set relation, %s vs %s",
				substraitgo.ErrInvalidRel, &output, &t)
		}
	}

	return &SetRel{
		RelCommon: RelCommon{mapping: remap},
		op:        op,
		inputs:    inputs,
	}, nil
}

func (b *builder) Set(op SetOp, inputs ...Rel) (*SetRel, error) {
	return b.SetRemap(op, nil, inputs...)
}

func (b *builder) PlanWithTypes(root Rel, rootNames []string, expectedTypeURLs []string, others ...Rel) (*Plan, error) {
	if root == nil {
		return nil, fmt.Errorf("%w: must provide non-nil root relation for plan",
			substraitgo.ErrInvalidRel)
	}

	rec := len(root.RecordType().Types)
	if rec != len(rootNames) {
		return nil, fmt.Errorf("%w: mismatched number of names and result record columns, got %d expected %d",
			substraitgo.ErrInvalidRel, len(rootNames), rec)
	}

	relations := make([]Relation, len(others)+1)
	relations[0].root = &Root{
		input: root, names: rootNames,
	}

	for i, o := range others {
		relations[i].rel = o
	}

	return &Plan{
		version:          &CurrentVersion,
		extensions:       b.extSet,
		reg:              b.reg,
		expectedTypeURLs: expectedTypeURLs,
		relations:        relations,
	}, nil
}

func (b *builder) Plan(root Rel, rootNames []string, others ...Rel) (*Plan, error) {
	return b.PlanWithTypes(root, rootNames, nil, others...)
}

var (
	_ Builder = (*builder)(nil)
)
