package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
)

func TestAnyType(t *testing.T) {
	for _, td := range []struct {
		testName       string
		argName        string
		nullability    types.Nullability
		expectedString string
	}{
		{"any", "any", types.NullabilityNullable, "any?"},
		{"anyrequired", "any", types.NullabilityRequired, "any"},
		{"anyOtherName", "any1", types.NullabilityNullable, "any1?"},
		{"T name", "T", types.NullabilityNullable, "T?"},
	} {
		t.Run(td.testName, func(t *testing.T) {
			arg := &types.AnyType{
				Name:        td.argName,
				Nullability: td.nullability,
			}
			anyType := arg.WithNullability(td.nullability)
			require.Equal(t, td.expectedString, anyType.String())
			require.Equal(t, td.nullability, anyType.GetNullability())
			require.Equal(t, td.argName, anyType.ShortString())
			// any type should be equal to any other type including itself
			require.True(t, anyType.Equals(anyType))
			require.True(t, anyType.Equals(&types.Int8Type{}))
		})
	}
}
