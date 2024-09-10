package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parameter_types"
)

func TestParameterizedMapType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   parameter_types.LeafIntParamAbstractType("P"),
		Scale:       parameter_types.LeafIntParamAbstractType("S"),
		Nullability: types.NullabilityRequired,
	}
	int8Type := &types.Int8Type{Nullability: types.NullabilityNullable}
	listType := &types.ParameterizedListType{Type: decimalType, Nullability: types.NullabilityNullable}
	for _, td := range []struct {
		name                    string
		Key                     types.Type
		Value                   types.Type
		nullability             types.Nullability
		expectedString          string
		expectedShortString     string
		expectedNrAbstractParam int
	}{
		{"single abstract param", decimalType, int8Type, types.NullabilityNullable, "map?<decimal<P,S>, i8?>", "map", 1},
		{"both abstract param", decimalType, listType, types.NullabilityNullable, "map?<decimal<P,S>, list?<decimal<P,S>>>", "map", 2},
	} {
		t.Run(td.name, func(t *testing.T) {
			pm := &types.ParameterizedMapType{Key: td.Key, Value: td.Value}
			pmType := pm.WithNullability(td.nullability)
			require.Equal(t, td.expectedString, pmType.String())
			require.Equal(t, td.expectedShortString, pmType.ShortString())
			require.True(t, pmType.Equals(pmType))

			pmAbsParamType, ok := pmType.(parameter_types.AbstractParameterType)
			require.True(t, ok)
			require.Equal(t, td.expectedString, pmAbsParamType.GetAbstractParamName())
			pmAbstractType, ok := pmType.(types.ParameterizedAbstractType)
			require.True(t, ok)
			require.Len(t, pmAbstractType.GetAbstractParameters(), td.expectedNrAbstractParam)
		})
	}
}
