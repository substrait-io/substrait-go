// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

var (
	extReg    = expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	uPointRef = extReg.GetTypeAnchor(extensions.TypeID{
		URN:  extensions.SubstraitDefaultURNPrefix + "extension_types",
		Name: "point",
	})

	subID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "subtract"}
	addID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "add"}
	indexInID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_set",
		Name: "index_in"}
	rankID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "rank"}
	firstValueID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "first_value"}
	extractID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_datetime",
		Name: "extract"}
	ntileID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "ntile"}
	sumID = extensions.FunctionID{
		URN:  extensions.SubstraitDefaultURNPrefix + "functions_arithmetic",
		Name: "sum"}

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
