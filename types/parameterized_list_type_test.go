// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

func TestParameterizedListType(t *testing.T) {
	decimalType := &ParameterizedDecimalType{
		Precision:   integer_parameters.NewVariableIntParam("P"),
		Scale:       integer_parameters.NewVariableIntParam("S"),
		Nullability: NullabilityRequired,
	}
	int8Type := &Int8Type{}
	for _, td := range []struct {
		name                           string
		param                          FuncDefArgType
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []interface{}
		expectedReturnType             Type
	}{
		{"parameterized param", decimalType, "list?<decimal<P,S>>", "list<decimal<P,S>>", true, []interface{}{decimalType}, nil},
		{"concrete param", int8Type, "list?<i8>", "list<i8>", false, nil, &ListType{Nullability: NullabilityRequired, Type: int8Type}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &ParameterizedListType{Type: td.param}
			assert.Equal(t, NullabilityUnspecified, pd.GetNullability())
			require.Equal(t, td.expectedNullableString, pd.SetNullability(NullabilityNullable).String())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(NullabilityRequired).String())
			assert.Equal(t, NullabilityRequired, pd.GetNullability())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
			assert.Equal(t, "list", pd.ShortString())
			retType, err := pd.ReturnType(nil, nil)
			if td.expectedReturnType == nil {
				assert.Error(t, err)
				require.True(t, pd.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
			}
		})
	}
}
