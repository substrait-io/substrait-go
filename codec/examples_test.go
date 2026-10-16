// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"fmt"

	"github.com/substrait-io/substrait-go/codec/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

func ExampleExpression_scalarFunction() {
	// define extensions with no plan for now
	const planExt = `{
		"extensionUrns": [
			{
				"extensionUrnAnchor": 1,
				"urn": "extension:io.substrait:functions_arithmetic"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUrnReference": 1,
					"functionAnchor": 2,
					"name": "add:i32_i32"
				}
			}
		],
		"relations": []
	}`

	var plan proto.Plan
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}

	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := codec.GetExtensionSet(&plan, collection)
	if err != nil {
		panic(err)
	}

	// json proto to represent of add(field_ref(0), float64(10))
	const scalarFunction = `{
		"scalarFunction": {
		  "functionReference": 2,
		  "outputType": {"i32": {}},
		  "arguments": [
			{"value": {"selection": {
				"rootReference": {},
				"directReference": {"structField": {"field": 0}}}}},
			{"value": {"literal": {"fp64": 10}}}
		  ]
		}
	  }`

	var exprProto proto.Expression
	if err := protojson.Unmarshal([]byte(scalarFunction), &exprProto); err != nil {
		panic(err)
	}

	reg := expr.NewExtensionRegistry(extSet, collection)
	// convert from protobuf to Expression!
	fromProto, err := codec.ExprFromProto(&exprProto, nil, reg)
	if err != nil {
		panic(err)
	}

	// manually define the entire expression instead of going through
	// having to construct the protobuf
	const substraitext = `extension:io.substrait:functions_arithmetic`

	var addVariant = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitext, Name: "add:i32_i32"})

	var ex expr.Expression
	refArg, _ := expr.NewRootFieldRef(expr.NewStructFieldRef(0), types.NewRecordTypeFromTypes([]types.Type{&types.Int32Type{}}))
	ex, _ = expr.NewCustomScalarFunc(reg, addVariant, &types.Int32Type{}, nil,
		refArg, expr.NewPrimitiveLiteral(float64(10), false))

	// convert our manual expression to proto.Expression
	toProto := codec.ExprToProto(ex)

	// output some info!

	// print string represention of the expression
	fmt.Println(fromProto)
	// print the string representation of our
	// manually constructed expression
	fmt.Println(ex)

	// verify that the Equals methods work recursively
	fmt.Println(ex.Equals(fromProto))
	// confirm our manually constructed expression is the same
	// as the one we got from protojson
	fmt.Println(pb.Equal(&exprProto, toProto))

	// Output:
	// add(.field(0), fp64(10)) => i32
	// add(.field(0) => i32, fp64(10)) => i32
	// true
	// true
}
