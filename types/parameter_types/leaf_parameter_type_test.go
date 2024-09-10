package parameter_types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types/parameter_types"
)

func TestConcreteParameterType(t *testing.T) {
	concreteType1 := parameter_types.LeafIntParamConcreteType(1)
	require.Equal(t, "1", concreteType1.String())
}

func TestLeafParameterType(t *testing.T) {
	var concreteType1, concreteType2, abstractType1 parameter_types.LeafParameter

	concreteType1 = parameter_types.LeafIntParamConcreteType(1)
	concreteType2 = parameter_types.LeafIntParamConcreteType(2)

	abstractType1 = parameter_types.LeafIntParamAbstractType("P")

	// verify string val
	require.Equal(t, "1", concreteType1.String())
	require.Equal(t, "P", abstractType1.String())

	// concrete type is only compatible with same type
	require.True(t, concreteType1.IsCompatible(concreteType1))
	require.False(t, concreteType1.IsCompatible(concreteType2))

	// abstract type is compatible with both abstract and concrete type
	require.True(t, abstractType1.IsCompatible(abstractType1))
	require.True(t, abstractType1.IsCompatible(concreteType2))
}
