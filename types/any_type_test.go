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
			require.Equal(t, td.expectedString, arg.String())
			require.Equal(t, td.nullability, arg.GetNullability())
			require.Equal(t, td.argName, arg.ShortString())
		})
	}
}
