// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestDynamicParameterEquals(t *testing.T) {
	i64Req := &types.Int64Type{Nullability: types.NullabilityRequired}
	fp64Req := &types.Float64Type{Nullability: types.NullabilityRequired}

	base := &expr.DynamicParameter{OutputType: i64Req, ParameterReference: 0}

	tests := []struct {
		name  string
		other expr.Expression
		want  bool
	}{
		{"same type and ref", &expr.DynamicParameter{OutputType: i64Req, ParameterReference: 0}, true},
		{"different ref", &expr.DynamicParameter{OutputType: i64Req, ParameterReference: 1}, false},
		{"different type", &expr.DynamicParameter{OutputType: fp64Req, ParameterReference: 0}, false},
		{"different expression kind", expr.NewPrimitiveLiteral(int64(42), false), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, base.Equals(tt.other))
		})
	}
}

func TestDynamicParameterVisit(t *testing.T) {
	dp := &expr.DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 5,
	}

	visited := dp.Visit(func(e expr.Expression) expr.Expression { return e })
	assert.Same(t, dp, visited, "Visit should return same pointer for leaf expression")
}

func TestDynamicParameterBuilderNilType(t *testing.T) {
	b := expr.ExprBuilder{
		Reg: expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
	}

	_, err := b.DynamicParam(nil, 0).BuildExpr()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "dynamic parameter must have an output type")
}

func TestDynamicParameterBuilderAsFuncArg(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
		BaseSchema: types.NewRecordTypeFromStruct(boringSchema.Struct),
	}

	dpBuilder := b.DynamicParam(&types.Int8Type{Nullability: types.NullabilityRequired}, 0)

	e, err := b.ScalarFunc(addID).Args(
		dpBuilder,
		b.Wrap(expr.NewLiteral(int8(5), false)),
	).BuildExpr()
	require.NoError(t, err)
	assert.Contains(t, e.String(), "$0:i8")
}

func TestDynamicParameterTypeMismatchInFunction(t *testing.T) {
	b := expr.ExprBuilder{
		Reg:        expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError()),
		BaseSchema: types.NewRecordTypeFromStruct(boringSchema.Struct),
	}

	tests := []struct {
		name   string
		funcID extensions.FunctionID
		dpType types.Type
		lit    func() (expr.Literal, error)
	}{
		{
			name:   "i32 where i8 expected",
			funcID: extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "add:i8_i8"},
			dpType: &types.Int32Type{Nullability: types.NullabilityRequired},
			lit:    func() (expr.Literal, error) { return expr.NewLiteral(int8(5), false) },
		},
		{
			name:   "string where numeric expected",
			funcID: addID,
			dpType: &types.StringType{Nullability: types.NullabilityRequired},
			lit:    func() (expr.Literal, error) { return expr.NewLiteral(int32(5), false) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := b.ScalarFunc(tt.funcID).Args(
				b.DynamicParam(tt.dpType, 0),
				b.Wrap(tt.lit()),
			).BuildExpr()
			require.Error(t, err)
		})
	}
}
