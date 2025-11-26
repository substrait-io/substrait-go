package literal

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v6"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/types"
)

// VirtualTableFromGoTypes converts Go values to VirtualTable expressions and infers their types.
// It accepts field names, tuples of polymorphic Go values, and optional nullability configuration.
// Returns the converted expressions, inferred column types, and any error encountered.
// If nullableColumns is nil, all columns default to non-nullable (required).
func VirtualTableFromGoTypes(fieldNames []string, tuples [][]any, nullableColumns []bool) ([]expr.VirtualTableExpressionValue, []types.Type, error) {
	// Need at least one tuple to infer column types from Go values.
	// Empty virtual tables are valid, but require explicit type specification via VirtualTableFromExpr.
	if len(tuples) == 0 {
		return nil, nil, fmt.Errorf("%w: must provide at least one tuple for virtual table", substraitgo.ErrInvalidRel)
	}

	nfields := len(fieldNames)
	if nfields == 0 {
		return nil, nil, fmt.Errorf("%w: must provide at least one field name", substraitgo.ErrInvalidRel)
	}

	for i, tuple := range tuples {
		if len(tuple) != nfields {
			return nil, nil, fmt.Errorf("%w: tuple %d has %d values, expected %d", substraitgo.ErrInvalidRel, i, len(tuple), nfields)
		}
	}

	if nullableColumns == nil {
		// default behavior is that none of the columns are nullable
		nullableColumns = make([]bool, nfields)
	} else if len(nullableColumns) != nfields {
		return nil, nil, fmt.Errorf("%w: nullableColumns length (%d) must match fieldNames length (%d) or be nil", substraitgo.ErrInvalidRel, len(nullableColumns), nfields)
	}

	columnTypes, err := inferColumnTypesFromGoTypes(tuples, fieldNames, nullableColumns)
	if err != nil {
		return nil, nil, err
	}

	if err := validateColumnTypesFromGoTypes(tuples, fieldNames, nullableColumns, columnTypes); err != nil {
		return nil, nil, err
	}

	values, err := convertGoTuplesToExpressions(tuples, fieldNames, columnTypes)
	if err != nil {
		return nil, nil, err
	}

	return values, columnTypes, nil
}

func GoTypeToSubstraitType(val any, nullable bool) (types.Type, error) {
	nullability := types.NullabilityRequired
	if nullable {
		nullability = types.NullabilityNullable
	}

	switch val.(type) {
	case bool:
		return &types.BooleanType{Nullability: nullability}, nil
	case int8:
		return &types.Int8Type{Nullability: nullability}, nil
	case int16:
		return &types.Int16Type{Nullability: nullability}, nil
	case int32:
		return &types.Int32Type{Nullability: nullability}, nil
	case int:
		return &types.Int64Type{Nullability: nullability}, nil
	case int64:
		return &types.Int64Type{Nullability: nullability}, nil
	case float32:
		return &types.Float32Type{Nullability: nullability}, nil
	case float64:
		return &types.Float64Type{Nullability: nullability}, nil
	case string:
		return &types.StringType{Nullability: nullability}, nil
	default:
		return nil, fmt.Errorf("unsupported Go type: %T", val)
	}
}

func GoValueToExpression(val any, expectedType types.Type) (expr.Expression, error) {
	actualType, err := GoTypeToSubstraitType(val, false)
	if err != nil {
		return nil, err
	}

	// Compare base types (ignore nullability for this check)
	actualBase := actualType.WithNullability(types.NullabilityRequired)
	expectedBase := expectedType.WithNullability(types.NullabilityRequired)

	if !actualBase.Equals(expectedBase) {
		return nil, fmt.Errorf("type mismatch: got %T, expected type compatible with %s", val, expectedType)
	}

	switch v := val.(type) {
	case bool:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case int8:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case int16:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case int32:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case int:
		return expr.NewPrimitiveLiteral(int64(v), expectedType.GetNullability() == types.NullabilityNullable), nil
	case int64:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case float32:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case float64:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	case string:
		return expr.NewPrimitiveLiteral(v, expectedType.GetNullability() == types.NullabilityNullable), nil
	default:
		return nil, fmt.Errorf("unsupported value type: %T", val)
	}
}

// inferColumnTypesFromGoTypes infers Substrait types from the first non-null value in each column
func inferColumnTypesFromGoTypes(tuples [][]any, fieldNames []string, nullableColumns []bool) ([]types.Type, error) {
	nfields := len(fieldNames)
	columnTypes := make([]types.Type, nfields)

	for colIdx := range nfields {
		var foundType types.Type
		for rowIdx := range len(tuples) {
			val := tuples[rowIdx][colIdx]
			if val != nil {
				var err error
				foundType, err = GoTypeToSubstraitType(val, nullableColumns[colIdx])
				if err != nil {
					return nil, fmt.Errorf("failed to infer type for column %d (%s): %w", colIdx, fieldNames[colIdx], err)
				}
				break
			}
		}

		if foundType == nil {
			return nil, fmt.Errorf("%w: column %d (%s) contains only null values, cannot infer type", substraitgo.ErrInvalidRel, colIdx, fieldNames[colIdx])
		}

		columnTypes[colIdx] = foundType
	}
	return columnTypes, nil
}

// validateColumnTypesFromGoTypes validates that the values in each column of every row conform to the type specified in columnTypes
func validateColumnTypesFromGoTypes(tuples [][]any, fieldNames []string, nullableColumns []bool, columnTypes []types.Type) error {
	nfields := len(fieldNames)

	for colIdx := range nfields {
		expectedType := columnTypes[colIdx]

		for rowIdx := range len(tuples) {
			val := tuples[rowIdx][colIdx]
			if val != nil {
				currentType, err := GoTypeToSubstraitType(val, nullableColumns[colIdx])
				if err != nil {
					return fmt.Errorf("invalid type in row %d, col %d (%s): %w", rowIdx, colIdx, fieldNames[colIdx], err)
				}

				// Compare base types (ignore nullability for this check)
				expectedBase := expectedType.WithNullability(types.NullabilityRequired)
				currentBase := currentType.WithNullability(types.NullabilityRequired)

				if !expectedBase.Equals(currentBase) {
					return fmt.Errorf("%w: type mismatch in column %d (%s): found %T in row %d, expected type compatible with %s",
						substraitgo.ErrInvalidRel, colIdx, fieldNames[colIdx], val, rowIdx, expectedType)
				}
			}
		}
	}
	return nil
}

func convertGoTuplesToExpressions(tuples [][]any, fieldNames []string, columnTypes []types.Type) ([]expr.VirtualTableExpressionValue, error) {
	nfields := len(fieldNames)
	values := make([]expr.VirtualTableExpressionValue, len(tuples))

	for rowIdx, tuple := range tuples {
		row := make(expr.VirtualTableExpressionValue, nfields)

		for colIdx, val := range tuple {
			expectedType := columnTypes[colIdx]

			if val == nil {
				row[colIdx] = expr.NewNullLiteral(expectedType)
			} else {
				exprVal, err := GoValueToExpression(val, expectedType)
				if err != nil {
					return nil, fmt.Errorf("failed to convert value at row %d, col %d (%s): %w", rowIdx, colIdx, fieldNames[colIdx], err)
				}
				row[colIdx] = exprVal
			}
		}

		values[rowIdx] = row
	}

	return values, nil
}
