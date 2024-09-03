package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
)

func TestParameterizedVarCharType(t *testing.T) {
	for _, td := range []struct {
		name                string
		typ                 types.ParameterizedSingleIntegerType
		nullability         types.Nullability
		integerOption       types.IntegerParam
		expectedString      string
		expectedBaseString  string
		expectedShortString string
	}{
		{"nullable varchar", &types.ParameterizedVarCharType{}, types.NullabilityNullable, types.IntegerParam{Name: "L1"}, "varchar?<L1>", "varchar", "vchar"},
		{"non nullable varchar", &types.ParameterizedVarCharType{}, types.NullabilityRequired, types.IntegerParam{Name: "L1"}, "varchar<L1>", "varchar", "vchar"},
		{"nullable fixChar", &types.ParameterizedFixedCharType{}, types.NullabilityNullable, types.IntegerParam{Name: "L1"}, "char?<L1>", "char", "fchar"},
		{"non nullable fixChar", &types.ParameterizedFixedCharType{}, types.NullabilityRequired, types.IntegerParam{Name: "L1"}, "char<L1>", "char", "fchar"},
		{"nullable fixBinary", &types.ParameterizedFixedBinaryType{}, types.NullabilityNullable, types.IntegerParam{Name: "L1"}, "fixedbinary?<L1>", "fixedbinary", "fbin"},
		{"non nullable fixBinary", &types.ParameterizedFixedBinaryType{}, types.NullabilityRequired, types.IntegerParam{Name: "L1"}, "fixedbinary<L1>", "fixedbinary", "fbin"},
		{"nullable precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{}, types.NullabilityNullable, types.IntegerParam{Name: "L1"}, "precision_timestamp?<L1>", "precision_timestamp", "prets"},
		{"non nullable precisionTimeStamp", &types.ParameterizedPrecisionTimestampType{}, types.NullabilityRequired, types.IntegerParam{Name: "L1"}, "precision_timestamp<L1>", "precision_timestamp", "prets"},
		{"nullable precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{}, types.NullabilityNullable, types.IntegerParam{Name: "L1"}, "precision_timestamp_tz?<L1>", "precision_timestamp_tz", "pretstz"},
		{"non nullable precisionTimeStampTz", &types.ParameterizedPrecisionTimestampTzType{}, types.NullabilityRequired, types.IntegerParam{Name: "L1"}, "precision_timestamp_tz<L1>", "precision_timestamp_tz", "pretstz"},
	} {
		t.Run(td.name, func(t *testing.T) {
			pt := td.typ.WithIntegerOption(td.integerOption).WithNullability(td.nullability)
			require.Equal(t, td.expectedString, pt.String())
			require.Equal(t, td.expectedShortString, pt.ShortString())
			require.True(t, pt.Equals(pt))
		})
	}
}

func TestParameterizedDecimalType(t *testing.T) {
	for _, td := range []struct {
		name                string
		precision           string
		scale               string
		nullability         types.Nullability
		expectedString      string
		expectedBaseString  string
		expectedShortString string
	}{
		{"nullable decimal", "P", "S", types.NullabilityNullable, "decimal?<P,S>", "decimal", "dec"},
		{"non nullable decimal", "P", "S", types.NullabilityRequired, "decimal<P,S>", "decimal", "dec"},
	} {
		t.Run(td.name, func(t *testing.T) {
			precision := types.IntegerParam{Name: td.precision}
			scale := types.IntegerParam{Name: td.scale}
			pt := types.ParameterizedDecimalType{Precision: precision, Scale: scale, Nullability: td.nullability}
			require.Equal(t, td.expectedString, pt.String())
			//require.Equal(t, td.expectedBaseString, pt.BaseString())
			require.Equal(t, td.expectedShortString, pt.ShortString())
			require.True(t, pt.Equals(pt))
		})
	}
}
