// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

// TestCreateTableAsSelectRoundTrip verifies that generated plans match the expected JSON.
func TestCreateTableAsSelectRoundTrip(t *testing.T) {
	for _, td := range []struct {
		name            string
		ctasTableName   []string
		ctasTableSchema types.NamedStruct
		getProjection   func(t *testing.T, b plan.Builder) plan.Rel
	}{
		{"ctas_basic", []string{"main", "employee_salaries"}, employeeSalariesSchema, getProjectionForTest1},
		{"ctas_with_filter", []string{"main", "filtered_employees"}, employeeSchemaNullable, getProjectionForTest2},
	} {
		t.Run(td.name, func(t *testing.T) {
			// Load the expected JSON. This will be our baseline for comparison.
			expectedJson, err := testdataFS.ReadFile(fmt.Sprintf("testdata/plan/%s.json", td.name))
			require.NoError(t, err)

			// build plan for CTAS
			b := plan.NewBuilderDefault()
			ctasRel, err := b.CreateTableAsSelect(td.getProjection(t, b), td.ctasTableName, td.ctasTableSchema)
			require.NoError(t, err)
			ctasPlan, err := b.Plan(ctasRel, td.ctasTableSchema.Names)
			require.NoError(t, err)

			// Check that the generated plan matches the expected JSON.
			checkRoundTrip(t, string(expectedJson), ctasPlan)
		})
	}
}
