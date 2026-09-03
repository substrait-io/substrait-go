package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
)

var employeeSchema = types.NamedStruct{Names: []string{"employee_id", "name", "department_id", "salary", "role"},
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

func TestCreateTableAsSelect(t *testing.T) {
	b := plan.NewBuilderDefault()

	tableNames := []string{"main", "employee_salaries"}
	input := b.NamedScan(tableNames, employeeSchema)

	ctas, err := b.CreateTableAsSelect(input, tableNames, employeeSchema)
	require.NoError(t, err)
	assert.Equal(t, plan.OutputModeModifiedRecords, ctas.OutputMode())
	assert.Equal(t, ctas.TableSchema(), employeeSchema)
	assert.Equal(t, tableNames, ctas.Names())
	assert.Equal(t, "struct<i32, string?, i32?, decimal?<10,2>, string?>", ctas.RecordType().String())

	ctasRemap, err := ctas.Remap(0, 2)
	require.NoError(t, err)
	assert.Equal(t, "struct<i32, i32?>", ctasRemap.RecordType().String())
}
