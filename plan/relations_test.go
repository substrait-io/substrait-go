package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/expr"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// TestExtensionDefinition is a simple test implementation of ExtensionRelDefinition
type TestExtensionDefinition struct {
	schema types.RecordType
	detail []byte
	exprs  []expr.Expression
}

func (t *TestExtensionDefinition) Schema(inputs []Rel) types.RecordType {
	return t.schema
}

func (t *TestExtensionDefinition) Build(inputs []Rel) *anypb.Any {
	if t.detail == nil {
		return nil
	}
	message := &wrapperspb.StringValue{Value: string(t.detail)}
	any, _ := anypb.New(message)
	return any
}

func (t *TestExtensionDefinition) Expressions(inputs []Rel) []expr.Expression {
	return t.exprs
}

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
	extReg := expr.NewExtensionRegistry(extensions.NewSet(), extensions.GetDefaultCollectionWithNoError())
	aggregateFnID := extensions.ID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "avg",
	}
	aggregateFn, err := expr.NewAggregateFunc(extReg,
		aggregateFnID, nil, types.AggInvocationAll,
		types.AggPhaseInitialToResult, nil, createPrimitiveFloat(1.0))
	require.NoError(t, err)
	aggregateFnRevised, err := expr.NewAggregateFunc(extReg,
		aggregateFnID, nil, types.AggInvocationAll,
		types.AggPhaseInitialToResult, nil, createPrimitiveFloat(9.0))
	require.NoError(t, err)

	aggregateRel := &AggregateRel{input: createVirtualTableReadRel(1),
		groupingExpressions: []expr.Expression{createPrimitiveFloat(1.0)},
		groupingReferences:  [][]uint32{{0}},
		measures:            []AggRelMeasure{{measure: aggregateFn, filter: expr.NewPrimitiveLiteral(false, false)}}}
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
	icebergTableReadRel := &IcebergTableReadRel{
		baseReadRel: baseReadRel{filter: expr.NewPrimitiveLiteral(true, false)},
		tableType: &Direct{
			MetadataUri: "s3://bucket/path/to/metadata.json",
		},
	}

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
			name:      "AggregateRel Copy with new inputs",
			relation:  aggregateRel,
			newInputs: []Rel{createVirtualTableReadRel(6)},
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(6),
				groupingReferences:  aggregateRel.groupingReferences,
				groupingExpressions: aggregateRel.groupingExpressions,
				measures:            aggregateRel.measures},
		},
		{
			name:            "AggregateRel Copy with same inputs and noOpRewrite",
			relation:        aggregateRel,
			newInputs:       aggregateRel.GetInputs(),
			rewriteFunc:     noOpRewrite,
			expectedSameRel: true,
		},
		{
			name:      "AggregateRel Copy with new Inputs and noOpReWrite",
			relation:  aggregateRel,
			newInputs: []Rel{createVirtualTableReadRel(7)},
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(7),
				groupingExpressions: aggregateRel.groupingExpressions,
				groupingReferences:  aggregateRel.groupingReferences,
				measures:            aggregateRel.measures},
		},
		{
			name:      "AggregateRel Copy with new Inputs and rewriteFunc",
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
			expectedRel: &AggregateRel{input: createVirtualTableReadRel(8),
				groupingExpressions: []expr.Expression{createPrimitiveFloat(9.0)},
				groupingReferences:  [][]uint32{{0}},
				measures:            []AggRelMeasure{{measure: aggregateFnRevised, filter: expr.NewPrimitiveLiteral(true, false)}}},
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
			name:      "VirtualTableReadRel Copy with rewriteFunc",
			relation:  virtualTableReadRel,
			newInputs: []Rel{},
			rewriteFunc: func(expression expr.Expression) (expr.Expression, error) {
				return expr.NewPrimitiveLiteral(true, false), nil
			},
			expectedRel: &VirtualTableReadRel{
				baseReadRel: baseReadRel{
					filter:           &expr.PrimitiveLiteral[bool]{Value: true, Type: &types.BooleanType{Nullability: types.NullabilityRequired}},
					bestEffortFilter: &expr.PrimitiveLiteral[bool]{Value: true, Type: &types.BooleanType{Nullability: types.NullabilityRequired}}},
				values: []expr.VirtualTableExpressionValue{[]expr.Expression{expr.NewPrimitiveLiteral(true, false)}}},
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
		{
			name:            "IcebergTableReadRel Copy with new inputs",
			relation:        icebergTableReadRel,
			newInputs:       []Rel{},
			expectedSameRel: true,
		},
		{
			name:        "IcebergTableReadRel Copy with new inputs and rewriteFunc",
			relation:    icebergTableReadRel,
			newInputs:   []Rel{},
			rewriteFunc: func(e expr.Expression) (expr.Expression, error) { return createPrimitiveFloat(9.0), nil },
			expectedRel: &IcebergTableReadRel{
				baseReadRel: baseReadRel{
					filter:           createPrimitiveFloat(9.0),
					bestEffortFilter: createPrimitiveFloat(9.0),
				},
				tableType: &Direct{
					MetadataUri: "s3://bucket/path/to/metadata.json",
				},
			},
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

func TestRelations_AdvancedExtensions(t *testing.T) {
	extReg := expr.NewExtensionRegistry(extensions.NewSet(), extensions.GetDefaultCollectionWithNoError())
	aggregateFnID := extensions.ID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "avg",
	}
	aggregateFn, err := expr.NewAggregateFunc(extReg,
		aggregateFnID, nil, types.AggInvocationAll,
		types.AggPhaseInitialToResult, nil, createPrimitiveFloat(1.0))
	require.NoError(t, err)

	aggregateRel := &AggregateRel{input: createVirtualTableReadRel(1),
		groupingExpressions: []expr.Expression{createPrimitiveFloat(1.0)},
		groupingReferences:  [][]uint32{{0}},
		measures:            []AggRelMeasure{{measure: aggregateFn, filter: expr.NewPrimitiveLiteral(false, false)}}}
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
	icebergTableReadRel := &IcebergTableReadRel{
		baseReadRel: baseReadRel{filter: expr.NewPrimitiveLiteral(true, false)},
		tableType: &Direct{
			MetadataUri: "s3://bucket/path/to/metadata.json",
		},
	}

	relations := []Rel{
		aggregateRel,
		crossRel,
		extensionLeafRel,
		extensionMultiRel,
		fetchRel,
		filterRel,
		hashJoinRel,
		joinRel,
		localFileReadRel,
		mergeJoinRel,
		namedTableReadRel,
		projectRel,
		setRel,
		sortRel,
		virtualTableReadRel,
		namedTableWriteRel,
		icebergTableReadRel,
	}

	val1, err := anypb.New(expr.NewPrimitiveLiteral("foo", false).ToProto())
	assert.NoError(t, err)

	exampleAdvancedExtension1 := &extensions.AdvancedExtension{
		Optimization: []*anypb.Any{val1},
		Enhancement:  val1,
	}

	val2, err := anypb.New(expr.NewPrimitiveLiteral("bar", false).ToProto())
	assert.NoError(t, err)

	exampleAdvancedExtension2 := &extensions.AdvancedExtension{
		Optimization: []*anypb.Any{val2},
		Enhancement:  val2,
	}

	for _, relation := range relations {
		// setting an extension should return the old/existing extension
		// setting an extension for the first time means the old extension should be nil
		oldExtension := relation.SetAdvancedExtension(exampleAdvancedExtension1)
		assert.Nil(t, oldExtension)
		assert.Equal(t, exampleAdvancedExtension1, relation.GetAdvancedExtension())

		// setting it again
		oldExtension = relation.SetAdvancedExtension(exampleAdvancedExtension2)
		assert.Equal(t, exampleAdvancedExtension1, oldExtension)
		assert.Equal(t, exampleAdvancedExtension2, relation.GetAdvancedExtension())

		// setting it to nil
		oldExtension = relation.SetAdvancedExtension(nil)
		assert.Equal(t, exampleAdvancedExtension2, oldExtension)
		assert.Nil(t, relation.GetAdvancedExtension())

	}
}

func TestAggregateRelToBuilder(t *testing.T) {
	extReg := expr.NewExtensionRegistry(extensions.NewSet(), extensions.GetDefaultCollectionWithNoError())
	aggregateFnID := extensions.ID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "avg",
	}
	aggregateFn, err := expr.NewAggregateFunc(extReg,
		aggregateFnID, nil, types.AggInvocationAll,
		types.AggPhaseInitialToResult, nil, createPrimitiveFloat(1.0))
	require.NoError(t, err)

	aggregateRel := &AggregateRel{input: createVirtualTableReadRel(1),
		groupingExpressions: []expr.Expression{createPrimitiveFloat(1.0)},
		groupingReferences:  [][]uint32{{0}},
		measures:            []AggRelMeasure{{measure: aggregateFn, filter: expr.NewPrimitiveLiteral(false, false)}}}

	builder := aggregateRel.ToBuilder()
	got, err := builder.Build()
	assert.NoError(t, err)
	assert.Equal(t, aggregateRel, got)
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

func (f *fakeRel) Remap(mapping ...int32) (Rel, error) {
	panic("unused")
}

func TestVirtualTableReadRelRecordType(t *testing.T) {
	b := NewBuilderDefault()
	rel, err := b.VirtualTable([]string{"a", "b"},
		expr.StructLiteralValue{
			&expr.PrimitiveLiteral[int64]{Value: 11, Type: &types.Int64Type{}},
			&expr.PrimitiveLiteral[string]{Value: "12", Type: &types.StringType{}}},
		expr.StructLiteralValue{
			&expr.PrimitiveLiteral[int64]{Value: 21, Type: &types.Int64Type{}},
			&expr.PrimitiveLiteral[string]{Value: "22", Type: &types.StringType{}}})
	assert.NoError(t, err)

	expected := *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}, &types.StringType{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	newRel, err := rel.Remap(0)
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = newRel.RecordType()
	assert.Equal(t, expected, result)
}

func TestExtensionTableReadRelRecordType(t *testing.T) {
	// We don't have a way of setting the base schema yet so test with an empty schema.
	rel := &ExtensionTableReadRel{}

	expected := *types.NewRecordTypeFromTypes(nil)
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	_, err := rel.Remap(0)
	assert.ErrorContains(t, err, "output mapping index out of range")
}

func TestLocalFileReadRelRecordType(t *testing.T) {
	// We don't have a way of setting the base schema yet so test with an empty schema.
	rel := &LocalFileReadRel{}

	expected := *types.NewRecordTypeFromTypes(nil)
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	_, err := rel.Remap(0)
	assert.ErrorContains(t, err, "output mapping index out of range")
}

func TestProjectRecordType(t *testing.T) {
	var rel ProjectRel
	rel.input = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}

	expected := *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}, &types.Int64Type{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	newRel, err := rel.Remap(0)
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = newRel.RecordType()
	assert.Equal(t, expected, result)
}

func TestExtensionSingleRecordType(t *testing.T) {
	var rel ExtensionSingleRel
	rel.input = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}
	rel.definition = &UndecodedExtension{}

	expected := *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}, &types.Int64Type{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	newRel, err := rel.Remap(0)
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = newRel.RecordType()
	assert.Equal(t, expected, result)
}

func TestExtensionLeafRecordType(t *testing.T) {
	var rel ExtensionLeafRel
	rel.definition = &UndecodedExtension{}

	expected := *types.NewRecordTypeFromTypes(nil)
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	_, err := rel.Remap(0)
	assert.ErrorContains(t, err, "output mapping index out of range")
}

func TestExtensionMultiRecordType(t *testing.T) {
	var rel ExtensionMultiRel
	rel.definition = &UndecodedExtension{}

	expected := *types.NewRecordTypeFromTypes(nil)
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	_, err := rel.Remap(0)
	assert.ErrorContains(t, err, "output mapping index out of range")
}

func TestHashJoinRecordType(t *testing.T) {
	var rel HashJoinRel
	rel.left = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}
	rel.right = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.StringType{}, &types.StringType{}})}

	expected := *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.StringType{}, &types.StringType{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	newRel, err := rel.Remap(0)
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes([]types.Type{&types.Int64Type{}})
	result = newRel.RecordType()
	assert.Equal(t, expected, result)
}

func TestMergeJoinRecordType(t *testing.T) {
	var rel MergeJoinRel
	rel.left = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}})}
	rel.right = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.StringType{}, &types.StringType{}})}

	expected := *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.StringType{}, &types.StringType{}})
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	newRel, err := rel.Remap(0)
	assert.NoError(t, err)
	expected = *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}})
	result = newRel.RecordType()
	assert.Equal(t, expected, result)
}

func TestNamedTableWriteRecordType(t *testing.T) {
	var rel NamedTableWriteRel
	rel.input = &fakeRel{outputType: *types.NewRecordTypeFromTypes(
		[]types.Type{&types.Int64Type{}, &types.StringType{}})}
	rel.outputMode = proto.WriteRel_OUTPUT_MODE_MODIFIED_RECORDS

	expected := *types.NewRecordTypeFromTypes(nil)
	result := rel.RecordType()
	assert.Equal(t, expected, result)

	_, err := rel.Remap(0)
	assert.ErrorContains(t, err, "output mapping index out of range")
}

func TestExtensionRelDefinitionInterface(t *testing.T) {
	b := NewBuilderDefault()

	// Test custom schema definition
	customSchema := *types.NewRecordTypeFromTypes([]types.Type{
		&types.Int64Type{},
		&types.StringType{},
	})

	testDef := &TestExtensionDefinition{
		schema: customSchema,
		detail: []byte("test-extension"),
		exprs:  []expr.Expression{expr.NewPrimitiveLiteral(int32(42), false)},
	}

	// Test ExtensionLeafRel with definition
	leafRel, err := b.ExtensionLeaf(testDef)
	require.NoError(t, err)
	assert.NotNil(t, leafRel.Definition())
	assert.Equal(t, customSchema, leafRel.RecordType())
	assert.NotNil(t, leafRel.Detail())

	// Test ExtensionSingleRel with definition
	input := b.NamedScan([]string{"test"}, types.NamedStruct{
		Names:  []string{"col1"},
		Struct: types.StructType{Types: []types.Type{&types.Int64Type{}}},
	})
	singleRel, err := b.ExtensionSingle(input, testDef)
	require.NoError(t, err)
	assert.NotNil(t, singleRel.Definition())
	assert.Equal(t, customSchema, singleRel.RecordType())
	assert.NotNil(t, singleRel.Detail())

	// Test ExtensionMultiRel with definition
	inputs := []Rel{input, input}
	multiRel, err := b.ExtensionMulti(inputs, testDef)
	require.NoError(t, err)
	assert.NotNil(t, multiRel.Definition())
	assert.Equal(t, customSchema, multiRel.RecordType())
	assert.NotNil(t, multiRel.Detail())

	// Test backward compatibility behavior by using wrapper
	oldStyleDef := &TestExtensionDefinition{
		schema: types.RecordType{}, // Empty schema like before
		detail: []byte("old-style"),
		exprs:  nil,
	}

	oldLeaf, err := b.ExtensionLeaf(oldStyleDef)
	require.NoError(t, err)
	assert.NotNil(t, oldLeaf.Definition())
	assert.Equal(t, types.RecordType{}, oldLeaf.RecordType()) // Empty schema as before

	oldSingleDef := &TestExtensionDefinition{
		schema: input.RecordType(), // Input schema like before
		detail: []byte("old-style"),
		exprs:  nil,
	}
	oldSingle, err := b.ExtensionSingle(input, oldSingleDef)
	require.NoError(t, err)
	assert.NotNil(t, oldSingle.Definition())
	assert.Equal(t, input.RecordType(), oldSingle.RecordType()) // Input schema as before

	oldMultiDef := &TestExtensionDefinition{
		schema: types.RecordType{}, // Empty schema like before
		detail: []byte("old-style"),
		exprs:  nil,
	}
	oldMulti, err := b.ExtensionMulti(inputs, oldMultiDef)
	require.NoError(t, err)
	assert.NotNil(t, oldMulti.Definition())
	// Should return empty schema as before for backward compatibility
	assert.Equal(t, types.RecordType{}, oldMulti.RecordType())
}

func TestUndecodedExtensionBackwardCompatibility(t *testing.T) {
	// Create a test detail
	detail := &anypb.Any{
		TypeUrl: "test.extension",
		Value:   []byte("test data"),
	}

	// Test UndecodedExtension with no inputs
	unknownExt := &UndecodedExtension{detail: detail}

	// Test Schema method with no inputs
	schema := unknownExt.Schema(nil)
	assert.Equal(t, types.RecordType{}, schema)

	// Test Schema method with inputs
	inputs := []Rel{createVirtualTableReadRel(2)}
	schemaWithInputs := unknownExt.Schema(inputs)
	assert.Equal(t, inputs[0].RecordType(), schemaWithInputs)

	// Test Build method
	built := unknownExt.Build(inputs)
	assert.Equal(t, detail, built)

	// Test Expressions method
	exprs := unknownExt.Expressions(inputs)
	assert.Nil(t, exprs)
}

func TestExtensionRelFromProtoBackwardCompatibility(t *testing.T) {
	// This test verifies that extension relations loaded from proto
	// now have an UndecodedExtension definition instead of nil
	// We'll test this by checking the existing behavior works

	// Test that UndecodedExtension behaves correctly
	detail := &anypb.Any{
		TypeUrl: "test.extension",
		Value:   []byte("test data"),
	}

	unknownExt := &UndecodedExtension{detail: detail}

	// Test that it implements ExtensionRelDefinition
	var _ ExtensionRelDefinition = unknownExt

	// Test that it returns the correct detail
	assert.Equal(t, detail, unknownExt.Build(nil))
	assert.Equal(t, detail, unknownExt.Build([]Rel{}))

	// Test that expressions return nil (unknown extensions don't have expressions)
	assert.Nil(t, unknownExt.Expressions(nil))
	assert.Nil(t, unknownExt.Expressions([]Rel{}))

	// Test schema behavior - empty for no inputs, first input's schema with inputs
	assert.Equal(t, types.RecordType{}, unknownExt.Schema(nil))
	assert.Equal(t, types.RecordType{}, unknownExt.Schema([]Rel{}))
}
