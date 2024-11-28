package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

func noOpRewrite(e expr.Expression) (expr.Expression, error) {
	return e, nil
}

func createVirtualTableReadRel(value int64) *VirtualTableReadRel {
	return &VirtualTableReadRel{values: []expr.VirtualTableExpressionValue{[]expr.Expression{&expr.PrimitiveLiteral[int64]{Value: value}}}}
}

func createPrimitiveFloat(value float64) expr.Expression {
	return expr.NewPrimitiveLiteral(value, false)
}

func createPrimitiveBool(value bool) expr.Expression {
	return expr.NewPrimitiveLiteral(value, false)
}

func TestRelations_Copy(t *testing.T) {
	aggregateRel := &AggregateRel{input: createVirtualTableReadRel(1), groups: [][]expr.Expression{{createPrimitiveFloat(1.0)}},
		measures: []AggRelMeasure{{filter: expr.NewPrimitiveLiteral(true, false)}}}
	crossRel := &CrossRel{left: createVirtualTableReadRel(1), right: createVirtualTableReadRel(2)}
	extensionLeafRel := &ExtensionLeafRel{}
	extensionMultiRel := &ExtensionMultiRel{inputs: []Rel{createVirtualTableReadRel(1), createVirtualTableReadRel(2)}}
	fetchRel := &FetchRel{input: createVirtualTableReadRel(1), offset: 1, count: 2}
	filterRel := &FilterRel{input: createVirtualTableReadRel(1), cond: expr.NewPrimitiveLiteral(true, false)}
	hashJoinRel := &HashJoinRel{left: createVirtualTableReadRel(1), right: createVirtualTableReadRel(2), joinType: HashMergeInner, leftKeys: []*expr.FieldReference{}, rightKeys: []*expr.FieldReference{}, postJoinFilter: expr.NewPrimitiveLiteral(true, false)}
	joinRel := &JoinRel{left: createVirtualTableReadRel(1), right: createVirtualTableReadRel(2), joinType: JoinTypeInner, expr: expr.NewPrimitiveLiteral(true, false), postJoinFilter: expr.NewPrimitiveLiteral(true, false)}
	localFileReadRel := &LocalFileReadRel{items: []FileOrFiles{{Path: "path"}}, baseReadRel: baseReadRel{filter: expr.NewPrimitiveLiteral(true, false), bestEffortFilter: expr.NewPrimitiveLiteral(true, false)}}
	mergeJoinRel := &MergeJoinRel{left: createVirtualTableReadRel(1), right: createVirtualTableReadRel(2), joinType: HashMergeInner, leftKeys: []*expr.FieldReference{}, rightKeys: []*expr.FieldReference{}, postJoinFilter: expr.NewPrimitiveLiteral(true, false)}
	namedTableReadRel := &NamedTableReadRel{names: []string{"mytest"}, baseReadRel: baseReadRel{filter: expr.NewPrimitiveLiteral(true, false), bestEffortFilter: expr.NewPrimitiveLiteral(true, false)}}
	projectRel := &ProjectRel{input: createVirtualTableReadRel(1), exprs: []expr.Expression{createPrimitiveFloat(1.0), createPrimitiveFloat(2.0)}}
	setRel := &SetRel{inputs: []Rel{createVirtualTableReadRel(1), createVirtualTableReadRel(2), createVirtualTableReadRel(3)}, op: SetOpUnionAll}
	sortRel := &SortRel{input: createVirtualTableReadRel(1), sorts: []expr.SortField{{Expr: createPrimitiveFloat(1.0), Kind: types.SortAscNullsFirst}}}
	virtualTableReadRel := &VirtualTableReadRel{values: []expr.VirtualTableExpressionValue{[]expr.Expression{&expr.PrimitiveLiteral[int64]{Value: 1}}}}
	namedTableWriteRel := &NamedTableWriteRel{input: namedTableReadRel}

	type relationTestCase struct {
		name            string
		relation        Rel
		newInputs       []Rel
		rewriteFunc     func(expr.Expression) (expr.Expression, error)
		expectedRel     Rel
		expectedSameRel bool
	}
	testCases := []relationTestCase{
		{
			name:        "AggregateRel Copy with new inputs",
			relation:    aggregateRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(6), groups: aggregateRel.groups, measures: aggregateRel.measures},
		},
		{
			name:            "AggregateRel Copy with same inputs and noOpRewrite",
			relation:        aggregateRel,
			newInputs:       aggregateRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "AggregateRel Copy with new Inputs and noOpReWrite",
			relation:    aggregateRel,
			newInputs:   []Rel{createVirtualTableReadRel(7)},
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(7), groups: aggregateRel.groups, measures: aggregateRel.measures},
		},
		{
			name:      "AggregateRel Copy with new Inputs and reWriteFunc",
			relation:  aggregateRel,
			newInputs: []Rel{createVirtualTableReadRel(8)},
			rewriteFunc: func(expression expr.Expression) (expr.Expression, error) {
				switch expression.(type) {
				case *expr.PrimitiveLiteral[float64]:
					return createPrimitiveFloat(9.0), nil
				case *expr.PrimitiveLiteral[bool]:
					return expr.NewPrimitiveLiteral(true, false), nil
				}
				panic("unexpected expression type")
			},
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(8), groups: [][]expr.Expression{{createPrimitiveFloat(9.0)}},
				measures: []AggRelMeasure{{filter: expr.NewPrimitiveLiteral(true, false)}}},
		},
		{
			name:            "ExtensionLeafRel Copy with new inputs",
			relation:        extensionLeafRel,
			newInputs:       []Rel{},
			expectedSameRel: true,
		},
		{
			name:        "ExtensionMultiRel Copy with new inputs",
			relation:    extensionMultiRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &ExtensionMultiRel{inputs: []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)}},
		},
		{
			name:            "ExtensionMultiRel Copy with same inputs and noOpRewrite",
			relation:        extensionMultiRel,
			newInputs:       extensionMultiRel.GetInputs(),
			expectedSameRel: true,
		},
		{
			name:        "FetchRel Copy with new inputs",
			relation:    fetchRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &FetchRel{input: createVirtualTableReadRel(6), offset: 1, count: 2},
		},
		{
			name:            "FetchRel Copy with same inputs and noOpRewrite",
			relation:        fetchRel,
			newInputs:       fetchRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "FilterRel Copy with new inputs",
			relation:    filterRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &FilterRel{input: createVirtualTableReadRel(6), cond: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:            "FilterRel Copy with same inputs and noOpRewrite",
			relation:        filterRel,
			newInputs:       filterRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "FilterRel Copy with new inputs and noOpRewrite",
			relation:    filterRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: noOpRewrite,
			expectedRel: &FilterRel{input: createVirtualTableReadRel(6), cond: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:        "FilterRel Copy with new inputs and rewriteFunc",
			relation:    filterRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveBool(true), nil },
			expectedRel: &FilterRel{input: createVirtualTableReadRel(6), cond: createPrimitiveBool(true)},
		},
		{
			name:        "HashJoinRel Copy with new inputs",
			relation:    hashJoinRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &HashJoinRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7), joinType: HashMergeInner, leftKeys: []*expr.FieldReference{}, rightKeys: []*expr.FieldReference{}, postJoinFilter: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:        "JoinRel Copy with new inputs",
			relation:    joinRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &JoinRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7), joinType: JoinTypeInner, expr: expr.NewPrimitiveLiteral(true, false), postJoinFilter: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:            "JoinRel Copy with same inputs and noOpRewrite",
			relation:        joinRel,
			newInputs:       joinRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "JoinRel Copy with new inputs and noOpRewrite",
			relation:    joinRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			rewriteFunc: noOpRewrite,
			expectedRel: &JoinRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7), joinType: JoinTypeInner, expr: expr.NewPrimitiveLiteral(true, false), postJoinFilter: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:        "JoinRel Copy with new inputs and rewriteFunc",
			relation:    joinRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveBool(false), nil },
			expectedRel: &JoinRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7), joinType: JoinTypeInner, expr: createPrimitiveBool(false), postJoinFilter: expr.NewPrimitiveLiteral(false, false)},
		},
		{
			name:            "LocalFileReadRel Copy with new inputs",
			relation:        localFileReadRel,
			newInputs:       []Rel{},
			expectedSameRel: true,
		},
		{
			name:        "MergeJoinRel Copy with new inputs",
			relation:    mergeJoinRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &MergeJoinRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7), joinType: HashMergeInner, leftKeys: []*expr.FieldReference{}, rightKeys: []*expr.FieldReference{}, postJoinFilter: expr.NewPrimitiveLiteral(true, false)},
		},
		{
			name:            "NamedTableReadRel Copy with new inputs",
			relation:        namedTableReadRel,
			newInputs:       []Rel{},
			expectedSameRel: true,
		},
		{
			name:        "ProjectRel Copy with new inputs",
			relation:    projectRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &ProjectRel{input: createVirtualTableReadRel(6), exprs: []expr.Expression{createPrimitiveFloat(1.0), createPrimitiveFloat(2.0)}},
		},
		{
			name:            "ProjectRel Copy with same inputs and noOpRewrite",
			relation:        projectRel,
			newInputs:       []Rel{projectRel.input},
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "ProjectRel Copy with new inputs and noOpRewrite",
			relation:    projectRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: noOpRewrite,
			expectedRel: &ProjectRel{input: createVirtualTableReadRel(6), exprs: []expr.Expression{createPrimitiveFloat(1.0), createPrimitiveFloat(2.0)}},
		},
		{
			name:        "ProjectRel Copy with new inputs and rewriteFunc",
			relation:    projectRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveFloat(9.0), nil },
			expectedRel: &ProjectRel{input: createVirtualTableReadRel(6), exprs: []expr.Expression{createPrimitiveFloat(9.0), createPrimitiveFloat(9.0)}},
		},
		{
			name:        "CrossRel Copy with new inputs",
			relation:    crossRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			expectedRel: &CrossRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7)},
		},
		{
			name:            "CrossRel Copy with same inputs and noOpRewrite",
			relation:        crossRel,
			newInputs:       crossRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "CrossRel Copy with new inputs and noOpRewrite",
			relation:    crossRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			rewriteFunc: noOpRewrite,
			expectedRel: &CrossRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7)},
		},
		{
			name:        "CrossRel Copy with new inputs and rewriteFunc",
			relation:    crossRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveFloat(9.0), nil },
			expectedRel: &CrossRel{left: createVirtualTableReadRel(6), right: createVirtualTableReadRel(7)},
		},
		{
			name:        "SetRel Copy with new inputs",
			relation:    setRel,
			newInputs:   []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7), createVirtualTableReadRel(8)},
			expectedRel: &SetRel{inputs: []Rel{createVirtualTableReadRel(6), createVirtualTableReadRel(7), createVirtualTableReadRel(8)}, op: SetOpUnionAll},
		},
		{
			name:            "SetRel Copy with same inputs and noOpRewrite",
			relation:        setRel,
			newInputs:       setRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "SortRel Copy with new inputs",
			relation:    sortRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &SortRel{input: createVirtualTableReadRel(6), sorts: []expr.SortField{{Expr: createPrimitiveFloat(1.0), Kind: types.SortAscNullsFirst}}},
		},
		{
			name:            "SortRel Copy with same inputs and noOpRewrite",
			relation:        sortRel,
			newInputs:       sortRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "SortRel Copy with new inputs and noOpRewrite",
			relation:    sortRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: noOpRewrite,
			expectedRel: &SortRel{input: createVirtualTableReadRel(6), sorts: []expr.SortField{{Expr: createPrimitiveFloat(1.0), Kind: types.SortAscNullsFirst}}},
		},
		{
			name:        "SortRel Copy with new inputs and rewriteFunc",
			relation:    sortRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveFloat(9.0), nil },
			expectedRel: &SortRel{input: createVirtualTableReadRel(6), sorts: []expr.SortField{{Expr: createPrimitiveFloat(9.0), Kind: types.SortAscNullsFirst}}},
		},
		{
			name:            "VirtualTableReadRel Copy with new inputs",
			relation:        virtualTableReadRel,
			newInputs:       []Rel{},
			expectedSameRel: true,
		},
		{
			name:        "NamedTableWriteRel Copy with new inputs",
			relation:    namedTableWriteRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			expectedRel: &NamedTableWriteRel{input: createVirtualTableReadRel(6)},
		},
		{
			name:            "NamedTableWriteRel Copy with same inputs and noOpRewrite",
			relation:        namedTableWriteRel,
			newInputs:       namedTableWriteRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:        "NamedTableWriteRel Copy with new inputs and noOpRewrite",
			relation:    namedTableWriteRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: noOpRewrite,
			expectedRel: &NamedTableWriteRel{input: createVirtualTableReadRel(6)},
		},
		{
			name:        "NamedTableWriteRel Copy with new inputs and rewriteFunc",
			relation:    namedTableWriteRel,
			newInputs:   []Rel{createVirtualTableReadRel(6)},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveFloat(9.0), nil },
			expectedRel: &NamedTableWriteRel{input: createVirtualTableReadRel(6)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got Rel
			var err error
			if tc.rewriteFunc == nil {
				got, err = tc.relation.Copy(tc.newInputs...)
			} else {
				got, err = tc.relation.CopyWithExpressionRewrite(tc.rewriteFunc, tc.newInputs...)
			}
			if !assert.NoError(t, err) {
				return
			}
			if tc.expectedSameRel {
				assert.Equal(t, tc.relation, got)
			} else {
				assert.Equal(t, tc.expectedRel, got)
			}
		})
	}
}

// fakeRel is a pretend relation that allows direct control of its direct output schema.
type fakeRel struct {
	RelCommon

	outputType types.RecordType
}

func (f *fakeRel) directOutputSchema() types.RecordType {
	return f.outputType
}

func (f *fakeRel) RecordType() types.RecordType {
	return f.remap(f.directOutputSchema())
}

func (f *fakeRel) ToProto() *proto.Rel {
	panic("unused")
}

func (f *fakeRel) ToProtoPlanRel() *proto.PlanRel {
	panic("unused")
}

func (f *fakeRel) Copy(newInputs ...Rel) (Rel, error) {
	panic("unused")
}

func (f *fakeRel) GetInputs() []Rel {
	panic("unused")
}

func (f *fakeRel) CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error) {
	panic("unused")
}

func (f *fakeRel) ChangeMapping(mapping []int32) error {
	newMapping, err := ChangeMapping(f, mapping)
	f.mapping = newMapping
	return err
}

func TestProjectRecordType(t *testing.T) {
	var rel ProjectRel
	rel.input = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}

	rel.ClearMapping()
	expected := *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}, &types.Int64Type{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	err := rel.ChangeMapping([]int32{0})
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = rel.RecordType()
	assert.Equal(t, expected, result)
}

func TestExtensionSingleRecordType(t *testing.T) {
	var rel ExtensionSingleRel
	rel.input = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}

	rel.ClearMapping()
	expected := *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}, &types.Int64Type{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	err := rel.ChangeMapping([]int32{0})
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = rel.RecordType()
	assert.Equal(t, expected, result)
}

func TestHashJoinRecordType(t *testing.T) {
	var rel HashJoinRel
	rel.left = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}
	rel.right = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.StringType{}, &types.StringType{}})}

	rel.ClearMapping()
	expected := *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.StringType{}, &types.StringType{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	err := rel.ChangeMapping([]int32{0})
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = rel.RecordType()
	assert.Equal(t, expected, result)
}

func TestMergeJoinRecordType(t *testing.T) {
	var rel MergeJoinRel
	rel.left = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}
	rel.right = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.StringType{}, &types.StringType{}})}

	rel.ClearMapping()
	expected := *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.StringType{}, &types.StringType{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	err := rel.ChangeMapping([]int32{0})
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}})
	result = rel.RecordType()
	assert.Equal(t, expected, result)
}
