package plan_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v4/expr"
	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/plan"
	"github.com/substrait-io/substrait-go/v4/types"
)

var (
	v1 = expr.PrimitiveLiteral[int32]{Value: 1, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
	v2 = expr.PrimitiveLiteral[int32]{Value: 2, Type: &types.Int32Type{Nullability: types.NullabilityRequired}}
)

// makeAddExpr constructs expression val1 + val2.
func makeAddExpr(t *testing.T, b plan.Builder, val1, val2 expr.Literal) expr.Expression {
	id := extensions.ID{
		URI:  "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml",
		Name: "add:i32_i32",
	}
	b.GetFunctionRef(id.URI, id.Name)
	scalarExpr, err := b.ScalarFn(id.URI, id.Name, nil, val1, val2)
	require.NoError(t, err)
	return scalarExpr
}

func buildLiteralExpressions(_ *testing.T, _ plan.Builder) []expr.VirtualTableExpressionValue {
	return []expr.VirtualTableExpressionValue{{&v1, &v2}}
}

// buildScalarAddExpression builds a scalar binary add expression
func buildScalarAddExpression(t *testing.T, b plan.Builder) []expr.VirtualTableExpressionValue {
	s1 := makeAddExpr(t, b, &v1, &v1)
	s2 := makeAddExpr(t, b, &v2, &v2)
	return []expr.VirtualTableExpressionValue{{s1, s2}}
}

// TestNamedTableInsertRoundTrip verifies that generated plans match the expected JSON.
func TestVirtualTableFromExprRoundTrip(t *testing.T) {
	for _, td := range []struct {
		name             string
		fieldNames       []string
		buildExprForTest func(t *testing.T, b plan.Builder) []expr.VirtualTableExpressionValue
	}{
		{"value_with_literal", []string{"col0", "col1"}, buildLiteralExpressions},
		{"value_with_scalar", []string{"col0", "col1"}, buildScalarAddExpression},
	} {
		t.Run(td.name, func(t *testing.T) {
			// Load the expected JSON. This will be our baseline for comparison.
			expectedJson, err := testdata.ReadFile(fmt.Sprintf("testdata/%s.json", td.name))
			require.NoError(t, err)

			// build plan for Project with virtual table
			b := plan.NewBuilderDefault()
			valueExpr := td.buildExprForTest(t, b)
			virtualTableExpr, err := b.VirtualTableFromExpr(td.fieldNames, valueExpr...)
			require.NoError(t, err)
			virtualTablePlan, err := b.Plan(virtualTableExpr, td.fieldNames)
			require.NoError(t, err)

			// Check that the generated plan matches the expected JSON.
			checkRoundTrip(t, string(expectedJson), virtualTablePlan)
		})
	}
}
