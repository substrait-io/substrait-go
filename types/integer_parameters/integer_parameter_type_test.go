// SPDX-License-Identifier: Apache-2.0

package integer_parameters_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v4/types/integer_parameters"
)

func TestConcreteParameterType(t *testing.T) {
	concreteType1 := integer_parameters.ConcreteIntParam(1)
	require.Equal(t, "1", concreteType1.String())
}

func TestLeafParameterType(t *testing.T) {
	var concreteType1, concreteType2, abstractType1 integer_parameters.IntegerParameter

	concreteType1 = integer_parameters.NewConcreteIntParam(1)
	concreteType2 = integer_parameters.NewConcreteIntParam(2)

	abstractType1 = integer_parameters.NewVariableIntParam("P")

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
