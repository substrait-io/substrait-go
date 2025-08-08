package plan_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v5/expr"
	"github.com/substrait-io/substrait-go/v5/extensions"
	"github.com/substrait-io/substrait-go/v5/literal"
	"github.com/substrait-io/substrait-go/v5/plan"
	"github.com/substrait-io/substrait-go/v5/types"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// Embed test JSON files for expected output comparison.
//
//go:embed testdata/*.json
var testdata embed.FS

// schema structures for testing purposes.
var (
	employeeSchema = types.NamedStruct{Names: []string{"employee_id", "name", "department_id", "salary", "role"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityRequired},
				&types.StringType{Nullability: types.NullabilityNullable},
				&types.Int32Type{Nullability: types.NullabilityNullable},
				&types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable},
				&types.StringType{Nullability: types.NullabilityNullable},
			},
		}}

	employeeSalariesSchema = types.NamedStruct{Names: []string{"name", "salary"},
		Struct: types.StructType{
			Types: []types.Type{
				&types.StringType{Nullability: types.NullabilityNullable},
				&types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable},
			},
		}}

	employeeSchemaNullable = types.NamedStruct{Names: []string{"employee_id", "name", "department_id", "salary", "role"},
		Struct: types.StructType{
			Types: []types.Type{
				&types.Int32Type{Nullability: types.NullabilityNullable},
				&types.StringType{Nullability: types.NullabilityNullable},
				&types.Int32Type{Nullability: types.NullabilityNullable},
				&types.DecimalType{Precision: 10, Scale: 2, Nullability: types.NullabilityNullable},
				&types.StringType{Nullability: types.NullabilityNullable},
			},
		}}
)

// makeProjectionMaskExpr generates a MaskExpression to project or reorder columns by the given IDs.
func makeProjectionMaskExpr(columnIds []int) *expr.MaskExpression {
	structItems := make([]*substraitproto.Expression_MaskExpression_StructItem, len(columnIds))

	for index, columnId := range columnIds {
		structItems[index] = &substraitproto.Expression_MaskExpression_StructItem{
			Field: int32(columnId),
		}
	}

	return expr.MaskExpressionFromProto(
		&substraitproto.Expression_MaskExpression{
			Select: &substraitproto.Expression_MaskExpression_StructSelect{
				StructItems: structItems,
			},
			MaintainSingularStruct: true,
		},
	)
}

// makeNamedTableReadRel creates a named table read relation with the selected column IDs.
func makeNamedTableReadRel(b plan.Builder, tableNames []string, tableSchema types.NamedStruct, columnIds []int) plan.Rel {
	namedTableReadRel := b.NamedScan(tableNames, tableSchema)
	namedTableReadRel.SetProjection(makeProjectionMaskExpr(columnIds))
	return namedTableReadRel
}

// makeConditionExprForLike constructs a LIKE condition expression for the specified column and value.
func makeConditionExprForLike(t *testing.T, b plan.Builder, scan plan.Rel, colId int, valueLiteral expr.Literal) expr.Expression {
	id := extensions.ID{
		URI:  "https://github.com/substrait-io/substrait/blob/main/extensions/functions_string.yaml",
		Name: "contains:str_str",
	}
	b.GetFunctionRef(id.URI, id.Name)
	colIdRef, err := b.RootFieldRef(scan, int32(colId))
	require.NoError(t, err)
	scalarExpr, err := b.ScalarFn(id.URI, id.Name, nil, colIdRef, valueLiteral)
	require.NoError(t, err)
	return scalarExpr
}

func makeFilterRel(t *testing.T, b plan.Builder, input plan.Rel, condition expr.Expression) plan.Rel {
	filterRel, err := b.Filter(input, condition)
	require.NoError(t, err)
	return filterRel
}

func makeProjectRel(t *testing.T, b plan.Builder, input plan.Rel, columnIds []int) plan.Rel {
	refs := make([]expr.Expression, len(columnIds))
	for i, c := range columnIds {
		ref, err := b.RootFieldRef(input, int32(c))
		require.NoError(t, err)
		refs[i] = ref
	}
	project, err := b.Project(input, refs...)
	require.NoError(t, err)
	return project
}

// getProjectionForTest1 returns project rel for "Select name, salary from employees"
func getProjectionForTest1(t *testing.T, b plan.Builder) plan.Rel {
	namedScanRel := makeNamedTableReadRel(b, []string{"employees"}, employeeSchema, []int{1, 3})
	return makeProjectRel(t, b, namedScanRel, []int{0, 1})
}

// getProjectionForTest2 returns project rel for "Select * from employees where role LIKE 'Engineer'"
func getProjectionForTest2(t *testing.T, b plan.Builder) plan.Rel {
	// scanRel outputs role, employee_id, name, department_id, salary
	namedScanRel := makeNamedTableReadRel(b, []string{"employees"}, employeeSchema, []int{4, 0, 1, 2, 3})

	// column 0 from the output of namedScanRel is role
	// Build the filter with condition `role LIKE 'Engineer'`
	l := literal.NewString("Engineer", false)
	roleLikeEngineer := makeConditionExprForLike(t, b, namedScanRel, 1, l)
	filterRel := makeFilterRel(t, b, namedScanRel, roleLikeEngineer)

	// projectRel output employee_id, name, department_id, salary, role
	return makeProjectRel(t, b, filterRel, []int{1, 2, 3, 4, 0})
}

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
			expectedJson, err := testdata.ReadFile(fmt.Sprintf("testdata/%s.json", td.name))
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
