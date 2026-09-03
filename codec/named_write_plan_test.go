// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

// TestNamedTableInsertRoundTrip verifies that generated plans match the expected JSON.
func TestNamedTableInsertRoundTrip(t *testing.T) {
	for _, td := range []struct {
		name        string
		tableName   []string
		tableSchema types.NamedStruct
		getInputRel func(t *testing.T, b plan.Builder) plan.Rel
	}{
		{"insert_from_select", []string{"main", "employee_salaries"}, employeeSalariesSchema, getProjectionForTest1},
	} {
		t.Run(td.name, func(t *testing.T) {
			// Load the expected JSON. This will be our baseline for comparison.
			expectedJson, err := testdataFS.ReadFile(fmt.Sprintf("testdata/plan/%s.json", td.name))
			require.NoError(t, err)

			// build plan for Insert
			b := plan.NewBuilderDefault()
			namedInsertRel, err := b.NamedInsert(td.getInputRel(t, b), td.tableName, td.tableSchema)
			require.NoError(t, err)
			namedInsertPlan, err := b.Plan(namedInsertRel, nil)
			require.NoError(t, err)

			// Check that the generated plan matches the expected JSON.
			checkRoundTrip(t, string(expectedJson), namedInsertPlan)
		})
	}
}

// TestNamedTableDeleteRoundTrip verifies that generated plans match the expected JSON.
func TestNamedTableDeleteRoundTrip(t *testing.T) {
	for _, td := range []struct {
		name        string
		tableName   []string
		tableSchema types.NamedStruct
		getInputRel func(t *testing.T, b plan.Builder) plan.Rel
	}{
		{"delete_with_filter", []string{"main", "employee_salaries"}, employeeSalariesSchema, getFilterForTest1},
	} {
		t.Run(td.name, func(t *testing.T) {
			// Load the expected JSON. This will be our baseline for comparison.
			expectedJson, err := testdataFS.ReadFile(fmt.Sprintf("testdata/plan/%s.json", td.name))
			require.NoError(t, err)

			// build plan for Delete
			b := plan.NewBuilderDefault()
			namedDeleteRel, err := b.NamedDelete(td.getInputRel(t, b), td.tableName, td.tableSchema)
			require.NoError(t, err)
			namedDeletePlan, err := b.Plan(namedDeleteRel, nil)
			require.NoError(t, err)

			// Check that the generated plan matches the expected JSON.
			checkRoundTrip(t, string(expectedJson), namedDeletePlan)
		})
	}
}
