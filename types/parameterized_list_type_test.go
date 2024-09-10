package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parameter_types"
)

func TestParameterizedListType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   parameter_types.LeafIntParamAbstractType("P"),
		Scale:       parameter_types.LeafIntParamAbstractType("S"),
		Nullability: types.NullabilityRequired,
	}
	for _, td := range []struct {
		name                string
		typ                 types.ParameterizedAbstractType
		nullability         types.Nullability
		expectedString      string
		expectedShortString string
	}{
		{"list", decimalType, types.NullabilityNullable, "list?<decimal<P,S>>", "list"},
	} {
		t.Run(td.name, func(t *testing.T) {
			pl := &types.ParameterizedListType{Type: td.typ}
			plType := pl.WithNullability(td.nullability)
			require.Equal(t, td.expectedString, plType.String())
			require.Equal(t, td.expectedShortString, plType.ShortString())
			require.True(t, plType.Equals(plType))

			plAbsParamType, ok := plType.(parameter_types.AbstractParameterType)
			require.True(t, ok)
			require.Equal(t, td.expectedString, plAbsParamType.GetAbstractParamName())
			plAbstractType, ok := plType.(types.ParameterizedAbstractType)
			require.True(t, ok)
			require.Len(t, plAbstractType.GetAbstractParameters(), 1)
		})
	}
}
