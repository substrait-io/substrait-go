// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
)

var (
	extReg    = NewEmptyExtensionRegistry(&extensions.DefaultCollection)
	uPointRef = extReg.GetTypeAnchor(extensions.ID{
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
		ex           Expression
		initialBound bool
		outputType   types.Type
	}{
		{NewPrimitiveLiteral(int32(1), true), true,
			&types.Int32Type{Nullability: types.NullabilityNullable}},
		{MustExpr(NewRootFieldRef(NewStructFieldRef(10), &boringSchema.Struct)), false,
			&types.StringType{}},
		{MustExpr(NewScalarFunc(extReg, subID, nil,
			NewPrimitiveLiteral(int8(1), false),
			NewPrimitiveLiteral(int8(5), false))), false,
			&types.Int8Type{Nullability: types.NullabilityRequired}},
		{MustExpr(NewScalarFunc(extReg, addID, nil,
			NewPrimitiveLiteral(int8(1), false),
			MustExpr(NewRootFieldRef(NewStructFieldRef(1), &boringSchema.Struct)))), false,
			&types.Int8Type{Nullability: types.NullabilityNullable}},
		{MustExpr(NewScalarFunc(extReg, indexInID, nil, MustExpr(NewRootFieldRef(NewStructFieldRef(2), &boringSchema.Struct)),
			NewListExpr(false, MustExpr(NewRootFieldRef(NewStructFieldRef(3), &boringSchema.Struct)),
				NewPrimitiveLiteral(int32(10), true)))), false,
			&types.Int64Type{Nullability: types.NullabilityNullable}},
		{MustExpr(NewWindowFunc(extReg, rankID, nil, types.AggInvocationAll, types.AggPhaseInitialToResult)),
			false, &types.Int64Type{Nullability: types.NullabilityNullable}},
		{MustExpr(NewScalarFunc(extReg, extractID, nil, types.Enum("YEAR"),
			MustExpr(NewRootFieldRef(NewStructFieldRef(9), &boringSchema.Struct)))), false,
			&types.Int64Type{Nullability: types.NullabilityRequired}},
	}

	for _, tt := range tests {
		t.Run(tt.ex.String(), func(t *testing.T) {
			assert.Truef(t, tt.outputType.Equals(tt.ex.GetType()), "expected: %s\ngot: %s", tt.outputType, tt.ex.GetType())
		})
	}
}
