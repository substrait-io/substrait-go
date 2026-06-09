package plan_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/plan"
	"github.com/substrait-io/substrait-go/v8/types"
)

func TestBuilderUserDefinedTypeUsesExtensionID(t *testing.T) {
	builder := plan.NewBuilder(extensions.GetDefaultCollectionWithNoError())

	udt := builder.UserDefinedType("extension:test:types", "point", &types.DataTypeParameter{Type: &types.Int32Type{}})

	require.Equal(t, types.NullabilityNullable, udt.Nullability)
	require.Equal(t, extensions.TypeID{URN: "extension:test:types", Name: "point"}, udt.ID)
	require.Len(t, udt.TypeParameters, 1)
}
