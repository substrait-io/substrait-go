// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
)

type Builder interface {
	UserDefinedType(nameSpace, typeName string, params ...types.TypeParam) types.UserDefinedType
	RootFieldRef(input Rel, index int32) (*expr.FieldReference, error)
	ScalarFn(nameSpace, key string, opts []*types.FunctionOption, args ...types.FuncArg) (*expr.ScalarFunction, error)
	AggregateFn(nameSpace, key string, opts []*types.FunctionOption, args ...types.FuncArg) (*expr.AggregateFunction, error)
	SortFields(input Rel, indices ...int32) ([]expr.SortField, error)
	Measure(measure *expr.AggregateFunction, filter expr.Expression) AggRelMeasure

	Project(input Rel, exprs []expr.Expression) *ProjectRel
	ProjectRemap(input Rel, exprs []expr.Expression, remap []int32) *ProjectRel
	AggregateColumnsRemap(input Rel, remap []int32, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error)
	AggregateColumns(input Rel, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error)
	AggregateExprsRemap(input Rel, remap []int32, measures []AggRelMeasure, groups ...[]expr.Expression) *AggregateRel
	AggregateExprs(input Rel, measures []AggRelMeasure, groups ...[]expr.Expression) *AggregateRel
	CrossRemap(left, right Rel, remap []int32) *CrossRel
	Cross(left, right Rel) *CrossRel
	FetchRemap(input Rel, offset, count int64, remap []int32) *FetchRel
	Fetch(input Rel, offset, count int64) *FetchRel
	FilterRemap(input Rel, condition expr.Expression, remap []int32) *FilterRel
	Filter(input Rel, condition expr.Expression) *FilterRel
	JoinAndFilterRemap(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType, remap []int32) *JoinRel
	JoinAndFilter(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType) *JoinRel
	JoinRemap(left, right Rel, condition expr.Expression, joinType JoinType, remap []int32) *JoinRel
	Join(left, right Rel, condition expr.Expression, joinType JoinType) *JoinRel
	NamedScanRemap(tableName []string, schema types.NamedStruct, remap []int32) *NamedTableReadRel
	NamedScan(tableName []string, schema types.NamedStruct) *NamedTableReadRel
	SortRemap(input Rel, remap []int32, sorts ...expr.SortField) *SortRel
	Sort(input Rel, sorts ...expr.SortField) *SortRel
	SetRemap(op SetOp, remap []int32, inputs ...Rel) (*SetRel, error)
	Set(op SetOp, inputs ...Rel) (*SetRel, error)

	PlanWithTypes(v *types.Version, root Rel, rootNames []string, expectedTypeURLs []string, others ...Rel) *Plan
	Plan(v *types.Version, root Rel, rootNames []string, others ...Rel) *Plan
}

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

type builder struct {
	ext    *extensions.Collection
	extSet extensions.Set

	reg expr.ExtensionRegistry
}

func (b *builder) UserDefinedType(nameSpace, typeName string, params ...types.TypeParam) types.UserDefinedType {
	id := extensions.ID{URI: nameSpace, Name: typeName}
	return types.UserDefinedType{
		Nullability:    types.NullabilityNullable,
		TypeReference:  b.extSet.GetTypeAnchor(id),
		TypeParameters: params,
	}
}

func (b *builder) RootFieldRef(input Rel, index int32) (*expr.FieldReference, error) {
	base := input.RecordType()
	if index > int32(len(base.Types)) {
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

func (b *builder) Project(input Rel, exprs []expr.Expression) *ProjectRel {
	return b.ProjectRemap(input, exprs, nil)
}

func (b *builder) ProjectRemap(input Rel, exprs []expr.Expression, remap []int32) *ProjectRel {
	return &ProjectRel{
		input: input,
		exprs: exprs,
	}
}

func (b *builder) Measure(measure *expr.AggregateFunction, filter expr.Expression) AggRelMeasure {
	return AggRelMeasure{
		measure: measure,
		filter:  filter,
	}
}

func (b *builder) AggregateColumnsRemap(input Rel, remap []int32, measures []AggRelMeasure, groupByCols ...int32) (*AggregateRel, error) {
	exprs := make([][]expr.Expression, len(groupByCols))
	for i, c := range groupByCols {
		ref, err := b.RootFieldRef(input, c)
		if err != nil {
			return nil, err
		}
		exprs[i] = []expr.Expression{ref}
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

func (b *builder) AggregateExprsRemap(input Rel, remap []int32, measures []AggRelMeasure, groups ...[]expr.Expression) *AggregateRel {
	return &AggregateRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		groups:    groups,
		measures:  measures,
	}
}

func (b *builder) AggregateExprs(input Rel, measures []AggRelMeasure, groups ...[]expr.Expression) *AggregateRel {
	return b.AggregateExprsRemap(input, nil, measures, groups...)
}

func (b *builder) CrossRemap(left, right Rel, remap []int32) *CrossRel {
	return &CrossRel{
		RelCommon: RelCommon{mapping: remap},
		left:      left, right: right,
	}
}

func (b *builder) Cross(left, right Rel) *CrossRel {
	return b.CrossRemap(left, right, nil)
}

func (b *builder) FetchRemap(input Rel, offset, count int64, remap []int32) *FetchRel {
	return &FetchRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		offset:    offset, count: count,
	}
}

func (b *builder) Fetch(input Rel, offset, count int64) *FetchRel {
	return b.FetchRemap(input, offset, count, nil)
}

func (b *builder) FilterRemap(input Rel, condition expr.Expression, remap []int32) *FilterRel {
	return &FilterRel{
		RelCommon: RelCommon{
			mapping: remap,
		},
		input: input,
		cond:  condition,
	}
}

func (b *builder) Filter(input Rel, condition expr.Expression) *FilterRel {
	return b.FilterRemap(input, condition, nil)
}

func (b *builder) JoinAndFilterRemap(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType, remap []int32) *JoinRel {
	return &JoinRel{
		RelCommon: RelCommon{mapping: remap},
		left:      left, right: right,
		expr: condition, postJoinFilter: postJoinFilter,
		joinType: joinType,
	}
}

func (b *builder) JoinAndFilter(left, right Rel, condition, postJoinFilter expr.Expression, joinType JoinType) *JoinRel {
	return b.JoinAndFilterRemap(left, right, condition, postJoinFilter, joinType, nil)
}

func (b *builder) JoinRemap(left, right Rel, condition expr.Expression, joinType JoinType, remap []int32) *JoinRel {
	return b.JoinAndFilterRemap(left, right, condition, nil, joinType, remap)
}

func (b *builder) Join(left, right Rel, condition expr.Expression, joinType JoinType) *JoinRel {
	return b.JoinAndFilterRemap(left, right, condition, nil, joinType, nil)
}

func (b *builder) NamedScanRemap(tableName []string, schema types.NamedStruct, remap []int32) *NamedTableReadRel {
	return &NamedTableReadRel{
		baseReadRel: baseReadRel{
			RelCommon: RelCommon{
				mapping: remap,
			},
			baseSchema: schema,
		},
		names: tableName,
	}
}

func (b *builder) NamedScan(tableName []string, schema types.NamedStruct) *NamedTableReadRel {
	return b.NamedScanRemap(tableName, schema, nil)
}

func (b *builder) SortRemap(input Rel, remap []int32, sorts ...expr.SortField) *SortRel {
	return &SortRel{
		RelCommon: RelCommon{mapping: remap},
		input:     input,
		sorts:     sorts,
	}
}

func (b *builder) Sort(input Rel, sorts ...expr.SortField) *SortRel {
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
	if len(inputs) < 2 {
		return nil, fmt.Errorf("%w: must have at least 2 relations for a SetRel, got %d",
			substraitgo.ErrInvalidRel, len(inputs))
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

func (b *builder) PlanWithTypes(v *types.Version, root Rel, rootNames []string, expectedTypeURLs []string, others ...Rel) *Plan {
	relations := make([]Relation, len(others)+1)
	relations[0].root = &Root{
		input: root, names: rootNames,
	}
	for i, o := range others {
		relations[i].rel = o
	}

	return &Plan{
		version:          v,
		extensions:       b.extSet,
		reg:              b.reg,
		expectedTypeURLs: expectedTypeURLs,
		relations:        relations,
	}
}

func (b *builder) Plan(v *types.Version, root Rel, rootNames []string, others ...Rel) *Plan {
	return b.PlanWithTypes(v, root, rootNames, nil, others...)
}
