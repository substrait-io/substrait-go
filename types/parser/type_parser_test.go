// SPDX-License-Identifier: Apache-2.0

package parser_test

import (
	"github.com/substrait-io/substrait-go/proto"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
	"github.com/substrait-io/substrait-go/types/parser"
)

func TestParser(t *testing.T) {
	parameterLeaf_L1 := integer_parameters.NewVariableIntParam("L1")
	parameterLeaf_P := integer_parameters.NewVariableIntParam("P")
	parameterLeaf_S := integer_parameters.NewVariableIntParam("S")
	concreteLeaf_5 := integer_parameters.NewConcreteIntParam(5)
	concreteLeaf_38 := integer_parameters.NewConcreteIntParam(38)
	concreteLeaf_10 := integer_parameters.NewConcreteIntParam(10)
	concreteLeaf_EMinus5 := integer_parameters.NewConcreteIntParam(int32(types.PrecisionEMinus5Seconds))
	tests := []struct {
		expr        string
		expected    string
		shortName   string
		expectedTyp types.FuncDefArgType
	}{
		{"2", "2", "", nil},
		{"-2", "-2", "", nil},
		{"i16?", "i16?", "i16", &types.Int16Type{Nullability: types.NullabilityNullable}},
		{"boolean", "boolean", "bool", &types.BooleanType{Nullability: types.NullabilityRequired}},
		{"fixedchar<5>", "fixedchar<5>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: concreteLeaf_5}},
		{"decimal<10,5>", "decimal<10,5>", "dec", &types.ParameterizedDecimalType{Precision: concreteLeaf_10, Scale: concreteLeaf_5, Nullability: types.NullabilityRequired}},
		{"list<decimal<10,5>>", "list<decimal<10,5>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: concreteLeaf_10, Scale: concreteLeaf_5, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"list?<decimal?<10,5>>", "list?<decimal?<10,5>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: concreteLeaf_10, Scale: concreteLeaf_5, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"struct<i16?,i32>", "struct<i16?, i32>", "struct", &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<boolean?,struct?<i16?,i32?,i64?>>", "map<boolean?, struct?<i16?, i32?, i64?>>", "map", &types.ParameterizedMapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityRequired}},
		{"map?<boolean?,struct?<i16?,i32?,i64?>>", "map?<boolean?, struct?<i16?, i32?, i64?>>", "map", &types.ParameterizedMapType{Key: &types.BooleanType{Nullability: types.NullabilityNullable}, Value: &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.Int16Type{Nullability: types.NullabilityNullable}, &types.Int32Type{Nullability: types.NullabilityNullable}, &types.Int64Type{Nullability: types.NullabilityNullable}}, Nullability: types.NullabilityNullable}, Nullability: types.NullabilityNullable}},
		{"precision_timestamp<5>", "precision_timestamp<5>", "prets", &types.ParameterizedPrecisionTimestampType{IntegerOption: concreteLeaf_EMinus5}},
		{"precision_timestamp_tz<5>", "precision_timestamp_tz<5>", "pretstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: concreteLeaf_EMinus5}},
		{"varchar<L1>", "varchar<L1>", "vchar", &types.ParameterizedVarCharType{IntegerOption: parameterLeaf_L1}},
		{"fixedchar<L1>", "fixedchar<L1>", "fchar", &types.ParameterizedFixedCharType{IntegerOption: parameterLeaf_L1}},
		{"fixedbinary<L1>", "fixedbinary<L1>", "fbin", &types.ParameterizedFixedBinaryType{IntegerOption: parameterLeaf_L1}},
		{"precision_timestamp<L1>", "precision_timestamp<L1>", "prets", &types.ParameterizedPrecisionTimestampType{IntegerOption: parameterLeaf_L1}},
		{"precision_timestamp_tz<L1>", "precision_timestamp_tz<L1>", "pretstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: parameterLeaf_L1}},
		{"decimal<P,S>", "decimal<P,S>", "dec", &types.ParameterizedDecimalType{Precision: parameterLeaf_S, Scale: parameterLeaf_S, Nullability: types.NullabilityRequired}},
		{"decimal<38,S>", "decimal<38,S>", "dec", &types.ParameterizedDecimalType{Precision: concreteLeaf_38, Scale: parameterLeaf_S, Nullability: types.NullabilityRequired}},
		{"any", "any", "any", types.AnyType{Nullability: types.NullabilityRequired}},
		{"any1?", "any1?", "any", types.AnyType{Nullability: types.NullabilityNullable}},
		{"list<decimal<P,S>>", "list<decimal<P,S>>", "list", &types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: parameterLeaf_P, Scale: parameterLeaf_S, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"struct<list?<decimal<P,S>>, i16>", "struct<list?<decimal<P,S>>, i16>", "struct", &types.ParameterizedStructType{Types: []types.FuncDefArgType{&types.ParameterizedListType{Type: &types.ParameterizedDecimalType{Precision: parameterLeaf_P, Scale: parameterLeaf_S, Nullability: types.NullabilityRequired}, Nullability: types.NullabilityNullable}, &types.Int16Type{Nullability: types.NullabilityRequired}}, Nullability: types.NullabilityRequired}},
		{"map<decimal<P,S>, i16>", "map<decimal<P,S>, i16>", "map", &types.ParameterizedMapType{Key: &types.ParameterizedDecimalType{Precision: parameterLeaf_P, Scale: parameterLeaf_S, Nullability: types.NullabilityRequired}, Value: &types.Int16Type{Nullability: types.NullabilityRequired}, Nullability: types.NullabilityRequired}},
		{"precision_timestamp_tz?<L1>", "precision_timestamp_tz?<L1>", "pretstz", &types.ParameterizedPrecisionTimestampTzType{IntegerOption: parameterLeaf_L1}},
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, td := range tests {
		t.Run(td.expr, func(t *testing.T) {
			d, err := p.ParseString(td.expr)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, d.Expr.String())
			if td.shortName != "" {
				assert.Equal(t, td.shortName, d.Expr.(*parser.Type).ShortType())
				typ, err := d.Expr.(*parser.Type).ArgType()
				assert.NoError(t, err)
				assert.Equal(t, reflect.TypeOf(td.expectedTyp), reflect.TypeOf(typ))
			}
		})
	}
}

func TestParserRetType(t *testing.T) {
	tests := []struct {
		expr        string
		expected    string
		shortName   string
		expectedTyp types.Type
	}{
		{"interval_day?<1>", "interval_day?<1>", "iday", &types.IntervalDayType{}},
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, td := range tests {
		t.Run(td.expr, func(t *testing.T) {
			d, err := p.ParseString(td.expr)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, d.Expr.String())
			if td.shortName != "" {
				assert.Equal(t, td.shortName, d.Expr.(*parser.Type).ShortType())
				retType, err := d.Expr.(*parser.Type).RetType()
				assert.NoError(t, err)
				assert.Equal(t, reflect.TypeOf(td.expectedTyp), reflect.TypeOf(retType))
			}
		})
	}
}

func TestParserListType(t *testing.T) {
	tests := []struct {
		expr        string
		expected    string
		expectedTyp types.Type
	}{
		{
			expr:     "list<i32>",
			expected: "list<i32>",
			expectedTyp: &types.ListType{
				Nullability: types.NullabilityRequired,
				Type:        &types.Int32Type{Nullability: types.NullabilityRequired},
			},
		},
		{
			expr:     "list?<i16?>",
			expected: "list?<i16?>",
			expectedTyp: &types.ListType{
				Nullability: types.NullabilityNullable,
				Type:        &types.Int16Type{Nullability: types.NullabilityNullable},
			},
		},
		{
			expr:     "list<i16>",
			expected: "list<i16>",
			expectedTyp: &types.ListType{
				Nullability: types.NullabilityRequired,
				Type:        &types.Int16Type{Nullability: types.NullabilityRequired},
			},
		},
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, td := range tests {
		t.Run(td.expr, func(t *testing.T) {
			d, err := p.ParseString(td.expr)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, d.Expr.String())

			if tExpr, ok := d.Expr.(*parser.Type); ok {
				retType, err := tExpr.RetType()
				assert.NoError(t, err)
				assert.Equal(t, reflect.TypeOf(td.expectedTyp), reflect.TypeOf(retType))
				assert.Equal(t, td.expectedTyp, retType)
			}
		})
	}
}

func TestParseUDT(t *testing.T) {
	tests := []struct {
		expr                string
		expected            string
		expectedTyp         types.Type
		expectedNullability proto.Type_Nullability
		expectedOptional    bool
	}{
		{"u!customtype1", "u!customtype1", &types.UserDefinedType{}, proto.Type_NULLABILITY_REQUIRED, false},
		{"u!customtype2?", "u!customtype2?", &types.UserDefinedType{}, proto.Type_NULLABILITY_NULLABLE, true},
	}

	p, err := parser.New()
	require.NoError(t, err)

	for _, td := range tests {
		t.Run(td.expr, func(t *testing.T) {
			d, err := p.ParseString(td.expr)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, d.Expr.String())
			retType, err := d.Expr.(*parser.Type).RetType()
			assert.NoError(t, err)
			assert.Equal(t, td.expectedNullability, retType.GetNullability())
			assert.Equal(t, td.expectedOptional, d.Expr.(*parser.Type).Optional())
			assert.Equal(t, reflect.TypeOf(td.expectedTyp), reflect.TypeOf(retType))
		})
	}
}
