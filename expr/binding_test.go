// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
)

var (
	extSet    = extensions.NewSet()
	uPointRef = extSet.GetTypeAnchor(extensions.ID{
		URI:  extensions.SubstraitDefaultURIPrefix + "extension_types.yaml",
		Name: "point",
	})

	subID = extensions.ID{
		URI:  extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml",
		Name: "subtract"}
	addID = extensions.ID{
		URI:  extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml",
		Name: "add"}
	indexInID = extensions.ID{
		URI:  extensions.SubstraitDefaultURIPrefix + "functions_set.yaml",
		Name: "index_in"}
	rankID = extensions.ID{
		URI:  extensions.SubstraitDefaultURIPrefix + "functions_arithmetic.yaml",
		Name: "rank"}
	extractID = extensions.ID{
		URI:  extensions.SubstraitDefaultURIPrefix + "functions_datetime.yaml",
		Name: "extract"}

	boringSchema = types.NamedStruct{
		Names: []string{
			"bool", "i8", "i32", "i32_req",
			"point", "i64", "f32", "f32_req",
			"f64", "date_req", "str", "bin"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types: []types.Type{
				&types.BooleanType{},
				&types.Int8Type{},
				&types.Int32Type{},
				&types.Int32Type{Nullability: types.NullabilityRequired},
				&types.UserDefinedType{
					TypeReference: uPointRef,
				},
				&types.Int64Type{},
				&types.Float32Type{},
				&types.Float32Type{Nullability: types.NullabilityRequired},
				&types.Float64Type{},
				&types.DateType{Nullability: types.NullabilityRequired},
				&types.StringType{},
				&types.BinaryType{},
			},
		},
	}
)

func TestBoundExpressions(t *testing.T) {
	tests := []struct {
		ex           expr.Expression
		initialBound bool
		outputType   types.Type
	}{
		{expr.NewPrimitiveLiteral(int32(1), true), true,
			&types.Int32Type{Nullability: types.NullabilityNullable}},
		{&expr.FieldReference{Reference: expr.NewStructFieldRef(10)}, false,
			&types.StringType{}},
		{expr.NewScalarFunc(subID, nil,
			expr.NewPrimitiveLiteral(int8(1), false),
			expr.NewPrimitiveLiteral(int8(5), false)), false,
			&types.Int8Type{Nullability: types.NullabilityRequired}},
		{expr.NewScalarFunc(addID, nil,
			expr.NewPrimitiveLiteral(int8(1), false),
			&expr.FieldReference{Reference: expr.NewStructFieldRef(1)}), false,
			&types.Int8Type{Nullability: types.NullabilityNullable}},
		{expr.NewScalarFunc(indexInID, nil, &expr.FieldReference{Reference: expr.NewStructFieldRef(2)},
			expr.NewListExpr(false, &expr.FieldReference{Reference: expr.NewStructFieldRef(3)},
				expr.NewPrimitiveLiteral(int32(10), true))), false,
			&types.Int64Type{Nullability: types.NullabilityNullable}},
		{expr.NewWindowFunc(rankID, types.AggPhaseInitialToResult, types.AggInvocationAll,
			nil), false, &types.Int64Type{Nullability: types.NullabilityNullable}},
		{expr.NewScalarFunc(extractID, nil, types.Enum("YEAR"),
			&expr.FieldReference{Reference: expr.NewStructFieldRef(9)}), false,
			&types.Int64Type{Nullability: types.NullabilityRequired}},
	}

	for _, tt := range tests {
		t.Run(tt.ex.String(), func(t *testing.T) {
			assert.Equal(t, tt.initialBound, tt.ex.IsBound())
			b, err := expr.BindExpression(tt.ex, boringSchema, extSet, &extensions.DefaultCollection)
			require.NoError(t, err)
			assert.Truef(t, tt.outputType.Equals(b.GetType()), "expected: %s\ngot: %s", tt.outputType, b.GetType())
		})
	}
}
