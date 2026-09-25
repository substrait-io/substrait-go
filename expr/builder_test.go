// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestCustomTypesInFunctionOutput(t *testing.T) {
	custom := `%YAML 1.2
---
urn: extension:test:custom
types:
  - name: custom_type1
  - name: custom_type2
  - name: custom_type3
  - name: custom_type4

scalar_functions:
  - name: custom_function
    description: "custom function that takes in and returns custom types"
    impls:
      - args:
          - name: arg1
            value: u!custom_type2
        return: u!custom_type1

aggregate_functions:
  - name: "custom_aggr"
    description: "custom aggregator that takes in and returns custom types"
    impls:
      - args:
          - name: arg1
            value: u!custom_type2
        return: u!custom_type3

window_functions:
  - name: "custom_window"
    description: "custom window function that takes in and returns custom types"
    impls:
      - args:
          - name: arg1
            value: u!custom_type2
        return: u!custom_type1
`

	customReader := strings.NewReader(custom)
	collection := extensions.Collection{}
	err := collection.Load(customReader)
	require.NoError(t, err)

	planBuilder := plan.NewBuilder(&collection)

	customType1 := planBuilder.UserDefinedType("extension:test:custom", "custom_type1")
	customType2 := planBuilder.UserDefinedType("extension:test:custom", "custom_type2")
	customType3 := planBuilder.UserDefinedType("extension:test:custom", "custom_type3")

	anyVal, err := anypb.New(expr.NewPrimitiveLiteral("foo", false).ToProto())
	require.NoError(t, err)

	customLiteral := planBuilder.GetExprBuilder().Literal(&expr.ProtoLiteral{
		Type:  &customType2,
		Value: expr.UserDefinedValueAny{Value: anyVal},
	})

	// check scalar function
	scalar, err := planBuilder.GetExprBuilder().ScalarFunc(extensions.FunctionID{
		URN:  "extension:test:custom",
		Name: "custom_function",
	}).Args(
		customLiteral,
	).BuildExpr()
	require.NoError(t, err)
	scalarProto := wire.ExprToProto(scalar)

	fnCall := scalarProto.GetScalarFunction()
	require.Len(t, fnCall.Arguments, 1)
	require.Equal(t, customType2.TypeReference, fnCall.Arguments[0].GetValue().GetLiteral().GetUserDefined().GetTypeReference())
	require.Equal(t, customType1.TypeReference, fnCall.OutputType.GetUserDefined().TypeReference)

	// check aggregate function
	aggr, err := planBuilder.GetExprBuilder().AggFunc(extensions.FunctionID{
		URN:  "extension:test:custom",
		Name: "custom_aggr",
	}).Args(
		customLiteral,
	).Build()
	require.NoError(t, err)
	aggrProto := wire.AggregateFunctionToProto(aggr)

	require.Len(t, aggrProto.Arguments, 1)
	require.Equal(t, customType2.TypeReference, aggrProto.Arguments[0].GetValue().GetLiteral().GetUserDefined().GetTypeReference())
	require.Equal(t, customType3.TypeReference, aggrProto.OutputType.GetUserDefined().TypeReference)

	// check window function
	window, err := planBuilder.GetExprBuilder().WindowFunc(extensions.FunctionID{
		URN:  "extension:test:custom",
		Name: "custom_window",
	}).Args(
		customLiteral,
	).Phase(types.AggregationPhaseInitialToResult).Build()
	require.NoError(t, err)
	windowProto := wire.ExprToProto(window)

	windowFnCall := windowProto.GetWindowFunction()
	require.Len(t, windowFnCall.Arguments, 1)
	require.Equal(t, customType2.TypeReference, windowFnCall.Arguments[0].GetValue().GetLiteral().GetUserDefined().GetTypeReference())
	require.Equal(t, customType1.TypeReference, windowFnCall.OutputType.GetUserDefined().TypeReference)

	// build a full plan and check that user defined types are registered in the extensions
	table, err := planBuilder.VirtualTable([]string{"col_a", "col_b"}, []expr.Literal{expr.NewPrimitiveLiteral(int64(2), false), expr.NewPrimitiveLiteral(int64(3), false)})
	require.NoError(t, err)

	aggregated, err := planBuilder.GetRelBuilder().AggregateRel(table, []plan.AggRelMeasure{planBuilder.Measure(aggr, window)}).Build()
	require.NoError(t, err)

	project, err := planBuilder.Project(aggregated, scalar)
	require.NoError(t, err)

	p, err := planBuilder.Plan(project, []string{"output1", "output2"})
	require.NoError(t, err)

	pp, err := p.ToProto()
	require.NoError(t, err)

	// custom_type1 is referenced as an argument and return type, so should be registered in the extensions
	// custom_type2 is referenced as an argument and return type, so should be registered in the extensions
	// custom_type3 is only referenced as a return type, but should still be registered in the extensions
	// custom_type4 is not referenced in the plan at all, so not be registerd in the extensions
	typeExtensionsFound := []string{}
	for _, ext := range pp.Extensions {
		typeExt := ext.GetExtensionType()
		if typeExt == nil {
			continue
		}
		typeExtensionsFound = append(typeExtensionsFound, typeExt.GetName())
	}
	require.Equal(t, 3, len(typeExtensionsFound))
	require.Contains(t, typeExtensionsFound, "custom_type1")
	require.Contains(t, typeExtensionsFound, "custom_type2")
	require.Contains(t, typeExtensionsFound, "custom_type3")
}

func TestAny1TypeParameterConsistency(t *testing.T) {
	// equal(any1, any1) -> boolean
	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())

	equalID := extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_comparison",
		Name: "equal",
	}

	t.Run("non-variadic any1 - invalid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral(int32(5), false),
			expr.NewPrimitiveLiteral(int64(10), false))
		require.Error(t, err, "equal should reject mixed types (i32, i64)")

		_, err = expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral("hello", false),
			expr.NewPrimitiveLiteral(int32(5), false))
		require.Error(t, err, "equal should reject mixed types (string, i32)")
	})

	t.Run("non-variadic any1 - valid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral(int32(5), false),
			expr.NewPrimitiveLiteral(int32(10), false))
		require.NoError(t, err, "equal should accept matching types (i32, i32)")

		_, err = expr.NewScalarFunc(reg, equalID, nil,
			expr.NewPrimitiveLiteral("hello", false),
			expr.NewPrimitiveLiteral("world", false))
		require.NoError(t, err, "equal should accept matching types (string, string)")
	})

	// coalesce(any1, any1, ...) -> any1 (min: 2 args)
	coalesceID := extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_comparison",
		Name: "coalesce",
	}

	t.Run("variadic any1 - invalid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral(int32(5), false),
			expr.NewPrimitiveLiteral(int64(10), false))
		require.Error(t, err, "coalesce should reject mixed types (i32, i64)")

		_, err = expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral("hello", false),
			expr.NewPrimitiveLiteral(int32(5), false))
		require.Error(t, err, "coalesce should reject mixed types (string, i32)")
	})

	t.Run("variadic any1 - valid", func(t *testing.T) {
		_, err := expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral(int32(1), false),
			expr.NewPrimitiveLiteral(int32(2), false),
			expr.NewPrimitiveLiteral(int32(3), false))
		require.NoError(t, err, "coalesce should accept all matching types")

		_, err = expr.NewScalarFunc(reg, coalesceID, nil,
			expr.NewPrimitiveLiteral("a", false),
			expr.NewPrimitiveLiteral("b", false),
			expr.NewPrimitiveLiteral("c", false))
		require.NoError(t, err, "coalesce should accept all matching types")
	})
}
