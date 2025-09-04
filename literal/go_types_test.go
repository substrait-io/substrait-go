package literal

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/types"
)

func TestGoTypeToSubstraitType(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		nullable bool
		expected types.Type
	}{
		{"bool required", true, false, &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"bool nullable", false, true, &types.BooleanType{Nullability: types.NullabilityNullable}},
		{"int8 required", int8(42), false, &types.Int8Type{Nullability: types.NullabilityRequired}},
		{"int16 required", int16(42), false, &types.Int16Type{Nullability: types.NullabilityRequired}},
		{"int32 required", int32(42), false, &types.Int32Type{Nullability: types.NullabilityRequired}},
		{"int required", int(42), false, &types.Int64Type{Nullability: types.NullabilityRequired}},
		{"int64 required", int64(42), false, &types.Int64Type{Nullability: types.NullabilityRequired}},
		{"float32 required", float32(3.14), false, &types.Float32Type{Nullability: types.NullabilityRequired}},
		{"float64 required", float64(3.14), false, &types.Float64Type{Nullability: types.NullabilityRequired}},
		{"string required", "hello", false, &types.StringType{Nullability: types.NullabilityRequired}},
		{"string nullable", "hello", true, &types.StringType{Nullability: types.NullabilityNullable}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GoTypeToSubstraitType(tt.value, tt.nullable)
			require.NoError(t, err)
			require.True(t, tt.expected.Equals(result), "Expected %s, got %s", tt.expected, result)
		})
	}
}

func TestGoTypeToSubstraitType_UnsupportedType(t *testing.T) {
	_, err := GoTypeToSubstraitType(make(chan int), false)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported Go type")
}

func TestGoValueToExpression(t *testing.T) {
	result, err := GoValueToExpression(true, &types.BooleanType{Nullability: types.NullabilityRequired})
	require.NoError(t, err)
	require.Equal(t, types.NullabilityRequired, result.GetType().GetNullability())

	result, err = GoValueToExpression(int32(42), &types.Int32Type{Nullability: types.NullabilityNullable})
	require.NoError(t, err)
	require.Equal(t, types.NullabilityNullable, result.GetType().GetNullability())

	result, err = GoValueToExpression("hello", &types.StringType{Nullability: types.NullabilityRequired})
	require.NoError(t, err)
	require.Equal(t, types.NullabilityRequired, result.GetType().GetNullability())
}

func TestGoValueToExpression_TypeMismatch(t *testing.T) {
	_, err := GoValueToExpression("hello", &types.Int32Type{Nullability: types.NullabilityRequired})
	require.Error(t, err)
	require.Contains(t, err.Error(), "type mismatch")
}

func TestVirtualTableFromGoTypes_BasicTypes(t *testing.T) {
	fieldNames := []string{"bool_col", "int32_col", "int64_col", "float32_col", "float64_col", "string_col"}
	tuples := [][]any{
		{true, int32(42), int64(100), float32(3.14), float64(2.71), "hello"},
		{false, int32(-10), int64(200), float32(1.23), float64(4.56), "world"},
	}

	values, columnTypes, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
	require.NoError(t, err)
	require.Len(t, values, 2)
	require.Len(t, columnTypes, 6)

	expectedTypes := []types.Type{
		&types.BooleanType{Nullability: types.NullabilityRequired},
		&types.Int32Type{Nullability: types.NullabilityRequired},
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.Float32Type{Nullability: types.NullabilityRequired},
		&types.Float64Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityRequired},
	}

	for i, expectedType := range expectedTypes {
		require.True(t, expectedType.Equals(columnTypes[i]), "Type mismatch at index %d: expected %s, got %s", i, expectedType, columnTypes[i])
	}

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
	fieldNames := []string{"nullable_int", "required_int"}
	tuples := [][]any{
		{int32(42), int32(100)},
		{int32(84), int32(200)},
	}
	nullableColumns := []bool{true, false}

	values, columnTypes, err := VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
	require.NoError(t, err)
	require.Len(t, columnTypes, 2)

	nullableType := columnTypes[0]
	require.Equal(t, types.NullabilityNullable, nullableType.GetNullability())

	requiredType := columnTypes[1]
	require.Equal(t, types.NullabilityRequired, requiredType.GetNullability())

	require.Len(t, values, 2)
	row1 := values[0]

	require.Equal(t, types.NullabilityNullable, row1[0].GetType().GetNullability())
	require.Equal(t, types.NullabilityRequired, row1[1].GetType().GetNullability())
}

func TestVirtualTableFromGoTypes_WithNullValues(t *testing.T) {
	fieldNames := []string{"name", "age", "active"}
	tuples := [][]any{
		{"Alice", int32(25), true},
		{"Bob", nil, false},
		{nil, int32(30), nil},
	}
	nullableColumns := []bool{true, true, true}

	values, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
	require.NoError(t, err)
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
	t.Run("empty tuples", func(t *testing.T) {
		_, _, err := VirtualTableFromGoTypes([]string{"col1"}, [][]any{}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "must provide at least one tuple")
	})

	t.Run("empty field names", func(t *testing.T) {
		_, _, err := VirtualTableFromGoTypes([]string{}, [][]any{{"val"}}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "must provide at least one field name")
	})

	t.Run("mismatched tuple length", func(t *testing.T) {
		fieldNames := []string{"col1", "col2"}
		tuples := [][]any{
			{"valid", "tuple"},
			{"invalid"},
		}
		_, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tuple 1 has 1 values, expected 2")
	})

	t.Run("mismatched nullability length", func(t *testing.T) {
		fieldNames := []string{"col1", "col2"}
		tuples := [][]any{{"val1", "val2"}}
		nullableColumns := []bool{true}

		_, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nullableColumns)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nullableColumns length (1) must match fieldNames length (2)")
	})

	t.Run("all null column", func(t *testing.T) {
		fieldNames := []string{"all_null_col"}
		tuples := [][]any{
			{nil},
			{nil},
		}
		_, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "contains only null values, cannot infer type")
	})

	t.Run("unsupported type", func(t *testing.T) {
		fieldNames := []string{"unsupported"}
		tuples := [][]any{
			{make(chan int)},
		}
		_, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "unsupported Go type")
	})

	t.Run("type mismatch across rows", func(t *testing.T) {
		fieldNames := []string{"mixed_type"}
		tuples := [][]any{
			{int32(42)},
			{"string"},
		}
		_, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
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
		_, _, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "type mismatch")
		require.Contains(t, err.Error(), "found string in row 2")
	})
}

func TestVirtualTableSchema(t *testing.T) {
	fieldNames := []string{"id", "name"}
	columnTypes := []types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityNullable},
	}

	schema := types.NamedStruct{
		Names: fieldNames,
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       columnTypes,
		},
	}

	require.Equal(t, fieldNames, schema.Names)
	require.Equal(t, types.NullabilityRequired, schema.Struct.Nullability)
	require.Len(t, schema.Struct.Types, 2)
	require.True(t, columnTypes[0].Equals(schema.Struct.Types[0]))
	require.True(t, columnTypes[1].Equals(schema.Struct.Types[1]))
}

func TestVirtualTableFromGoTypes_WithSchemaHelper(t *testing.T) {
	fieldNames := []string{"id", "active"}
	tuples := [][]any{
		{int64(1), true},
		{int64(2), false},
	}

	values, columnTypes, err := VirtualTableFromGoTypes(fieldNames, tuples, nil)
	require.NoError(t, err)
	require.Len(t, values, 2)
	require.Len(t, columnTypes, 2)

	schema := types.NamedStruct{
		Names: fieldNames,
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       columnTypes,
		},
	}
	require.Equal(t, fieldNames, schema.Names)
	require.Len(t, schema.Struct.Types, 2)

	expectedTypes := []types.Type{
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.BooleanType{Nullability: types.NullabilityRequired},
	}

	for i, expectedType := range expectedTypes {
		require.True(t, expectedType.Equals(schema.Struct.Types[i]), "Type mismatch at index %d", i)
	}
}
