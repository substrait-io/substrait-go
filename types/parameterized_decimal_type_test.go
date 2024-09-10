package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parameter_types"
)

func TestParameterizedDecimalType(t *testing.T) {
	for _, td := range []struct {
		name                string
		precision           string
		scale               string
		nullability         types.Nullability
		expectedString      string
		expectedShortString string
	}{
		{"nullable decimal", "P", "S", types.NullabilityNullable, "decimal?<P,S>", "dec"},
		{"non nullable decimal", "P", "S", types.NullabilityRequired, "decimal<P,S>", "dec"},
	} {
		t.Run(td.name, func(t *testing.T) {
			precision := parameter_types.LeafIntParamAbstractType(td.precision)
			scale := parameter_types.LeafIntParamAbstractType(td.scale)
			pd := &types.ParameterizedDecimalType{Precision: precision, Scale: scale}
			pdType := pd.WithNullability(td.nullability)
			require.Equal(t, td.expectedString, pdType.String())
			require.Equal(t, td.expectedShortString, pdType.ShortString())
			require.True(t, pdType.Equals(pdType))

			pdAbsParamType, ok := pdType.(parameter_types.AbstractParameterType)
			require.True(t, ok)
			require.Equal(t, td.expectedString, pdAbsParamType.GetAbstractParamName())
			pdAbstractType, ok := pdType.(types.ParameterizedAbstractType)
			require.True(t, ok)
			require.Len(t, pdAbstractType.GetAbstractParameters(), 2)
		})
	}
}

func TestParameterizedDecimalSingleAbstractParam(t *testing.T) {
	precision := parameter_types.LeafIntParamConcreteType(38)
	scale := parameter_types.LeafIntParamAbstractType("S")

	pd := &types.ParameterizedDecimalType{Precision: precision, Scale: scale}
	pdType := pd.WithNullability(types.NullabilityNullable)
	require.Equal(t, "decimal?<38,S>", pdType.String())
	require.Equal(t, "dec", pdType.ShortString())
	require.True(t, pdType.Equals(pdType))

	pdAbsParamType, ok := pdType.(parameter_types.AbstractParameterType)
	require.True(t, ok)
	require.Equal(t, "decimal?<38,S>", pdAbsParamType.GetAbstractParamName())
	pdAbstractType, ok := pdType.(types.ParameterizedAbstractType)
	require.True(t, ok)
	// only one abstract param
	require.Len(t, pdAbstractType.GetAbstractParameters(), 1)
}
