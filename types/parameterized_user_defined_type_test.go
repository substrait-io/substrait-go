package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/substrait-io/substrait-go/v5/types"
	"github.com/substrait-io/substrait-go/v5/types/integer_parameters"
)

func TestParameterizedUserDefinedType(t *testing.T) {
	decimalType := &ParameterizedDecimalType{
		Precision:   integer_parameters.NewVariableIntParam("P"),
		Scale:       integer_parameters.NewVariableIntParam("S"),
		Nullability: NullabilityRequired,
	}
	int8Type := &Int8Type{Nullability: NullabilityNullable}
	//userDefineType := &ParameterizedUserDefinedType{TypeParameters: []UDTParameter{}, Nullability: NullabilityNullable}
	for _, td := range []struct {
		name                           string
		Params                         []UDTParameter
		Args                           []interface{}
		expectedNullableString         string
		expectedNullableRequiredString string
		expectedHasParameterizedParam  bool
		expectedParameterizedParams    []any
		expectedReturnType             Type
	}{
		{"udt_noparam", []UDTParameter{}, []any{}, "u!udt_noparam?", "u!udt_noparam", false, nil, &UserDefinedType{Nullability: NullabilityRequired}},
		{"concrete_udt", []UDTParameter{&DataTypeUDTParam{int8Type}}, []any{&DataTypeParameter{int8Type}}, "u!concrete_udt?<i8?>", "u!concrete_udt<i8?>", false, nil, &UserDefinedType{Nullability: NullabilityRequired, TypeParameters: []TypeParam{&DataTypeParameter{Type: int8Type}}}},
		{"variable_udt", []UDTParameter{&DataTypeUDTParam{decimalType}}, []any{}, "u!variable_udt?<decimal<P,S>>", "u!variable_udt<decimal<P,S>>", true, []any{&DataTypeUDTParam{decimalType}}, nil},
		{"udt_with_int", []UDTParameter{&IntegerUDTParam{Integer: 10}}, []any{IntegerParameter(10)}, "u!udt_with_int?<10>", "u!udt_with_int<10>", false, nil, &UserDefinedType{Nullability: NullabilityRequired, TypeParameters: []TypeParam{IntegerParameter(10)}}},
		{"udt_with_str", []UDTParameter{&StringUDTParam{StringVal: "test"}}, []any{StringParameter("test")}, "u!udt_with_str?<test>", "u!udt_with_str<test>", false, nil, &UserDefinedType{Nullability: NullabilityRequired, TypeParameters: []TypeParam{StringParameter("test")}}},
		{"udt_with_int_and_str", []UDTParameter{&IntegerUDTParam{Integer: 10}, &StringUDTParam{StringVal: "test"}}, []any{IntegerParameter(10), StringParameter("test")}, "u!udt_with_int_and_str?<10, test>", "u!udt_with_int_and_str<10, test>", false, nil, &UserDefinedType{Nullability: NullabilityRequired, TypeParameters: []TypeParam{IntegerParameter(10), StringParameter("test")}}},
	} {
		t.Run(td.name, func(t *testing.T) {
			pd := &ParameterizedUserDefinedType{TypeParameters: td.Params, Name: td.name}
			assert.Equal(t, NullabilityUnspecified, pd.GetNullability())
			require.Equal(t, td.expectedNullableString, pd.SetNullability(NullabilityNullable).String())
			assert.Equal(t, NullabilityNullable, pd.GetNullability())
			require.Equal(t, td.expectedNullableRequiredString, pd.SetNullability(NullabilityRequired).String())
			assert.Equal(t, NullabilityRequired, pd.GetNullability())
			require.Equal(t, td.expectedHasParameterizedParam, pd.HasParameterizedParam())
			require.Equal(t, td.expectedParameterizedParams, pd.GetParameterizedParams())
			assert.Equal(t, fmt.Sprintf("u!%s", td.name), pd.ShortString())

			retType, err := pd.ReturnType(nil, nil)
			if td.expectedReturnType == nil {
				assert.Error(t, err)
				require.True(t, pd.HasParameterizedParam())
			} else {
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, retType)
				resultType, err := pd.WithParameters(td.Args)
				require.Nil(t, err)
				require.Equal(t, td.expectedReturnType, resultType)
			}
		})
	}

}
