package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parameter_types"
)

// a type to indicate all single integer type.
// helpful in initializing different single type integer type to the same interface
type parameterizedSingleIntegerType interface {
	types.Type
	WithIntegerOption(param parameter_types.LeafIntParamAbstractType) types.Type
}

func TestParameterizedSingleIntegerType(t *testing.T) {
	for _, td := range []struct {
		name                string
		typ                 parameterizedSingleIntegerType
		nullability         types.Nullability
		integerOption       parameter_types.LeafIntParamAbstractType
		expectedString      string
		expectedBaseString  string
		expectedShortString string
	}{
		{"nullable varchar", &types.ParameterizedVarCharType{}, types.NullabilityNullable, parameter_types.LeafIntParamAbstractType("L1"), "varchar?<L1>", "varchar", "vchar"},
		{"non nullable varchar", &types.ParameterizedVarCharType{}, types.NullabilityRequired, parameter_types.LeafIntParamAbstractType("L1"), "varchar<L1>", "varchar", "vchar"},
		{"nullable fixChar", &types.ParameterizedFixedCharType{}, types.NullabilityNullable, parameter_types.LeafIntParamAbstractType("L1"), "char?<L1>", "char", "fchar"},
		{"non nullable fixChar", &types.ParameterizedFixedCharType{}, types.NullabilityRequired, parameter_types.LeafIntParamAbstractType("L1"), "char<L1>", "char", "fchar"},
		{"nullable fixBinary", &types.ParameterizedFixedBinaryType{}, types.NullabilityNullable, parameter_types.LeafIntParamAbstractType("L1"), "fixedbinary?<L1>", "fixedbinary", "fbin"},
		{"non nullable fixBinary", &types.ParameterizedFixedBinaryType{}, types.NullabilityRequired, parameter_types.LeafIntParamAbstractType("L1"), "fixedbinary<L1>", "fixedbinary", "fbin"},
		{"nullable precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{}, types.NullabilityNullable, parameter_types.LeafIntParamAbstractType("L1"), "precision_timestamp?<L1>", "precision_timestamp", "prets"},
		{"non nullable precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{}, types.NullabilityRequired, parameter_types.LeafIntParamAbstractType("L1"), "precision_timestamp<L1>", "precision_timestamp", "prets"},
		{"nullable precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{}, types.NullabilityNullable, parameter_types.LeafIntParamAbstractType("L1"), "precision_timestamp_tz?<L1>", "precision_timestamp_tz", "pretstz"},
		{"non nullable precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{}, types.NullabilityRequired, parameter_types.LeafIntParamAbstractType("L1"), "precision_timestamp_tz<L1>", "precision_timestamp_tz", "pretstz"},
	} {
		t.Run(td.name, func(t *testing.T) {
			pt := td.typ.WithIntegerOption(td.integerOption).WithNullability(td.nullability)
			require.Equal(t, td.expectedString, pt.String())
			require.Equal(t, td.expectedShortString, pt.ShortString())
			require.True(t, pt.Equals(pt))
		})
	}
}
