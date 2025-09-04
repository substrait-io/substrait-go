package plan_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/plan"
	"github.com/substrait-io/substrait-go/v6/types"
)

func TestVirtualTableFromGoTypes_BasicTypes(t *testing.T) {
	b := plan.NewBuilderDefault()

	fieldNames := []string{"bool_col", "int32_col", "int64_col", "float32_col", "float64_col", "string_col"}
	tuples := [][]any{
		{true, int32(42), int64(100), float32(3.14), float64(2.71), "hello"},
		{false, int32(-10), int64(200), float32(1.23), float64(4.56), "world"},
	}
	table, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nil)
	require.NoError(t, err)
	require.NotNil(t, table)
	schema := table.RecordType()
	require.Equal(t, int32(6), schema.FieldCount())

	expectedTypes := []types.Type{
		&types.BooleanType{Nullability: types.NullabilityRequired},
		&types.Int32Type{Nullability: types.NullabilityRequired},
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.Float32Type{Nullability: types.NullabilityRequired},
		&types.Float64Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityRequired},
	}

	for i, expectedType := range expectedTypes {
		actualType := schema.Types()[i]
		require.True(t, expectedType.Equals(actualType), "Type mismatch at index %d: expected %s, got %s", i, expectedType, actualType)
	}

	values := table.Values()
	require.Len(t, values, 2)

	row1 := values[0]
	require.Len(t, row1, 6)

	boolLit, ok := row1[0].(*expr.PrimitiveLiteral[bool])
	require.True(t, ok)
	require.Equal(t, true, boolLit.Value)

	stringLit, ok := row1[5].(*expr.PrimitiveLiteral[string])
	require.True(t, ok)
	require.Equal(t, "hello", stringLit.Value)
}

func TestVirtualTableFromGoTypes_WithNullability(t *testing.T) {
	b := plan.NewBuilderDefault()

	fieldNames := []string{"nullable_int", "required_int"}
	tuples := [][]any{
		{int32(42), int32(100)},
		{int32(84), int32(200)},
	}
	nullableColumns := []bool{true, false}

	table, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
	require.NoError(t, err)
	require.NotNil(t, table)

	schema := table.RecordType()
	require.Equal(t, int32(2), schema.FieldCount())

	nullableType := schema.Types()[0]
	require.Equal(t, types.NullabilityNullable, nullableType.GetNullability())

	requiredType := schema.Types()[1]
	require.Equal(t, types.NullabilityRequired, requiredType.GetNullability())
}

func TestVirtualTableFromGoTypes_WithNullValues(t *testing.T) {
	b := plan.NewBuilderDefault()

	fieldNames := []string{"name", "age", "active"}
	tuples := [][]any{
		{"Alice", int32(25), true},
		{"Bob", nil, false},
		{nil, int32(30), nil},
	}
	nullableColumns := []bool{true, true, true}

	table, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
	require.NoError(t, err)
	require.NotNil(t, table)

	values := table.Values()
	require.Len(t, values, 3)

	row2 := values[1]
	require.Len(t, row2, 3)

	nameLit, ok := row2[0].(*expr.PrimitiveLiteral[string])
	require.True(t, ok)
	require.Equal(t, "Bob", nameLit.Value)

	ageNull, ok := row2[1].(*expr.NullLiteral)
	require.True(t, ok)
	require.True(t, ageNull.GetType().Equals(&types.Int32Type{Nullability: types.NullabilityNullable}))

	row3 := values[2]
	nameNull, ok := row3[0].(*expr.NullLiteral)
	require.True(t, ok)
	require.True(t, nameNull.GetType().Equals(&types.StringType{Nullability: types.NullabilityNullable}))

	activeNull, ok := row3[2].(*expr.NullLiteral)
	require.True(t, ok)
	require.True(t, activeNull.GetType().Equals(&types.BooleanType{Nullability: types.NullabilityNullable}))
}

func TestVirtualTableFromGoTypes_ErrorCases(t *testing.T) {
	b := plan.NewBuilderDefault()

	t.Run("empty tuples", func(t *testing.T) {
		_, err := b.VirtualTableFromGoTypes([]string{"col1"}, [][]any{}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot infer column types from empty tuples")
	})

	t.Run("empty field names", func(t *testing.T) {
		_, err := b.VirtualTableFromGoTypes([]string{}, [][]any{{"val"}}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "must provide at least one field name")
	})

	t.Run("mismatched tuple length", func(t *testing.T) {
		fieldNames := []string{"col1", "col2"}
		tuples := [][]any{
			{"valid", "tuple"},
			{"invalid"},
		}
		_, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tuple 1 has 1 values, expected 2")
	})

	t.Run("mismatched nullability length", func(t *testing.T) {
		fieldNames := []string{"col1", "col2"}
		tuples := [][]any{{"val1", "val2"}}
		nullableColumns := []bool{true} // should be length 2

		_, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nullableColumns length (1) must match fieldNames length (2)")
	})

	t.Run("all null column", func(t *testing.T) {
		fieldNames := []string{"all_null_col"}
		tuples := [][]any{
			{nil},
			{nil},
		}
		_, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "contains only null values, cannot infer type")
	})

	t.Run("unsupported type", func(t *testing.T) {
		fieldNames := []string{"unsupported"}
		tuples := [][]any{
			{make(chan int)},
		}
		_, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "unsupported Go type")
	})

	t.Run("type mismatch across rows", func(t *testing.T) {
		fieldNames := []string{"mixed_type"}
		tuples := [][]any{
			{int32(42)},
			{"string"},
		}
		_, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "type mismatch")
	})

	t.Run("type mismatch with nulls", func(t *testing.T) {
		fieldNames := []string{"col1", "col2"}
		tuples := [][]any{
			{nil, int32(1)},
			{int32(42), int32(2)},
			{"string", int32(3)},
		}
		_, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "type mismatch")
		require.Contains(t, err.Error(), "found string in row 2")
	})
}

func TestVirtualTableFromGoTypes_HappyPath(t *testing.T) {
	b := plan.NewBuilderDefault()

	fieldNames := []string{"id", "name", "active", "score"}
	tuples := [][]any{
		{int64(1), "Alice", true, float64(95.5)},
		{int64(2), "Bob", false, float64(87.2)},
		{int64(3), nil, true, nil},
	}
	nullableColumns := []bool{false, true, false, true}

	table, err := b.VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
	require.NoError(t, err)

	planObj, err := b.Plan(table, fieldNames)
	require.NoError(t, err)
	require.NotNil(t, planObj)

	proto, err := planObj.ToProto()
	require.NoError(t, err)
	require.NotNil(t, proto)
	require.Len(t, proto.Relations, 1)
}
