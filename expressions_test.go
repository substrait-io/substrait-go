// SPDX-License-Identifier: Apache-2.0

package substraitgo_test

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/proto"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

func ExampleExpression_scalarFunction() {
	// define extensions with no plan for now
	const planExt = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 2,
					"name": "add"
				}
			}
		],
		"relations": []
	}`

	var plan substraitgo.Plan
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}

	// get the extension set
	extSet := substraitgo.GetExtensionSet(&plan)

	// json proto to represent of add(field_ref(0), float64(10))
	const scalarFunction = `{
		"scalarFunction": {
		  "functionReference": 2,
		  "outputType": {"i32": {}},
		  "arguments": [
			{"value": {"selection": {
				"directReference": {"structField": {"field": 0}}}}},
			{"value": {"literal": {"fp64": 10}}}
		  ]
		}
	  }`

	var exprProto proto.Expression
	if err := protojson.Unmarshal([]byte(scalarFunction), &exprProto); err != nil {
		panic(err)
	}

	// convert from protobuf to Expression!
	fromProto, err := substraitgo.ExprFromProto(&exprProto, nil, extSet)
	if err != nil {
		panic(err)
	}

	// manually define the entire expression instead of going through
	// having to construct the protobuf
	const substraitext = `https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml`

	var expr substraitgo.Expression = &substraitgo.ScalarFunction{
		FuncRef: 2,
		ID:      substraitgo.ExtID{URI: substraitext, Name: "add"},
		Args: []substraitgo.FuncArg{
			&substraitgo.FieldReference{
				Root:      substraitgo.RootReference,
				Reference: &substraitgo.StructFieldRef{Field: 0},
			},
			substraitgo.NewPrimitiveLiteral(float64(10), false),
		},
		OutputType: &substraitgo.Int32Type{},
	}
	// call ToProto to convert our manual expression to proto.Expression
	toProto := expr.ToProto()

	// output some info!

	// print string represention of the expression
	fmt.Println(fromProto)
	// print the string representation of our
	// manually constructed expression
	fmt.Println(expr)

	// verify that the Equals methods work recursively
	fmt.Println(expr.Equals(fromProto))
	// confirm our manually constructed expression is the same
	// as the one we got from protojson
	fmt.Println(pb.Equal(&exprProto, toProto))

	// Output:
	// add(.field(0), fp64(10)) => i32
	// add(.field(0), fp64(10)) => i32
	// true
	// true
}
