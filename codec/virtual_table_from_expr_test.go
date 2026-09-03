// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
)

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
			expectedJson, err := testdataFS.ReadFile(fmt.Sprintf("testdata/plan/%s.json", td.name))
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
