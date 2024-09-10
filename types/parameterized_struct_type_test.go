package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parameter_types"
)

func TestParameterizedStructType(t *testing.T) {
	decimalType := &types.ParameterizedDecimalType{
		Precision:   parameter_types.LeafIntParamAbstractType("P"),
		Scale:       parameter_types.LeafIntParamAbstractType("S"),
		Nullability: types.NullabilityRequired,
	}
	int8Type := &types.Int8Type{Nullability: types.NullabilityNullable}
	listType := &types.ParameterizedListType{Type: decimalType, Nullability: types.NullabilityNullable}
	for _, td := range []struct {
		name                    string
		types                   []types.Type
		nullability             types.Nullability
		expectedString          string
		expectedShortString     string
		expectedNrAbstractParam int
	}{
		{"single abstract param", []types.Type{decimalType}, types.NullabilityNullable, "struct?<decimal<P,S>>", "struct", 1},
		{"multiple abstract param", []types.Type{decimalType, int8Type, listType}, types.NullabilityRequired, "struct<decimal<P,S>, i8?, list?<decimal<P,S>>>", "struct", 2},
	} {
		t.Run(td.name, func(t *testing.T) {
			ps := &types.ParameterizedStructType{Type: td.types}
			psType := ps.WithNullability(td.nullability)
			require.Equal(t, td.expectedString, psType.String())
			require.Equal(t, td.expectedShortString, psType.ShortString())
			require.True(t, psType.Equals(psType))

			psAbsParamType, ok := psType.(parameter_types.AbstractParameterType)
			require.True(t, ok)
			require.Equal(t, td.expectedString, psAbsParamType.GetAbstractParamName())
			psAbstractType, ok := psType.(types.ParameterizedAbstractType)
			require.True(t, ok)
			require.Len(t, psAbstractType.GetAbstractParameters(), td.expectedNrAbstractParam)
		})
	}
}
