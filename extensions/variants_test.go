// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/parser"
)

func TestEvaluateTypeExpression(t *testing.T) {
	var (
		p, _          = parser.New()
		i64Null, _    = p.ParseString("i64?")
		i64NonNull, _ = p.ParseString("i64")
		strNull, _    = p.ParseString("string?")
	)

	tests := []struct {
		name     string
		nulls    extensions.NullabilityHandling
		ret      parser.TypeExpression
		extArgs  extensions.ArgumentList
		args     []types.Type
		expected types.Type
		err      string
	}{
		{"defaults", "", *i64NonNull, extensions.ArgumentList{
			extensions.ValueArg{Value: i64Null}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable}, ""},
		{"arg mismatch", "", *strNull, extensions.ArgumentList{extensions.ValueArg{Value: strNull}},
			[]types.Type{}, nil, "invalid expression: mismatch in number of arguments provided. got 0, expected 1"},
		{"missing enum arg", "", *i64Null, extensions.ArgumentList{
			extensions.ValueArg{Value: i64NonNull}, extensions.EnumArg{Name: "foo"}},
			[]types.Type{&types.Int64Type{}, &types.Int64Type{}}, nil, "invalid type: arg #1 (foo) should be an enum"},
		{"discrete null handling", extensions.DiscreteNullability, *strNull, extensions.ArgumentList{
			extensions.ValueArg{Value: strNull}},
			[]types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			nil, "invalid type: discrete nullability did not match for arg #0"},
		{"mirror", extensions.MirrorNullability, *strNull, extensions.ArgumentList{
			extensions.ValueArg{Value: i64NonNull}, extensions.ValueArg{Value: i64Null}},
			[]types.Type{
				&types.Int64Type{Nullability: types.NullabilityRequired},
				&types.Int64Type{Nullability: types.NullabilityRequired}},
			&types.StringType{Nullability: types.NullabilityRequired}, ""},
		{"declared output", extensions.DeclaredOutputNullability, *strNull, extensions.ArgumentList{
			extensions.ValueArg{Value: strNull}},
			[]types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			&types.StringType{Nullability: types.NullabilityNullable}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extensions.EvaluateTypeExpression(tt.nulls, tt.ret, tt.extArgs, tt.args)
			if tt.err == "" {
				assert.NoError(t, err)
				assert.Truef(t, tt.expected.Equals(result), "expected: %s\ngot: %s", tt.expected, result)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
