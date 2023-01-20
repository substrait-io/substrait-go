// SPDX-License-Identifier: Apache-2.0

package substraitgo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func sampleNestedExpr(substraitExtURI string) substraitgo.Expression {
	var (
		addID         = substraitgo.ExtID{URI: substraitExtURI, Name: "add"}
		addRef uint32 = 2
		subID         = substraitgo.ExtID{URI: substraitExtURI, Name: "subtract"}
		subRef uint32 = 3
		mulID         = substraitgo.ExtID{URI: substraitExtURI, Name: "multiply"}
		mulRef uint32 = 4
	)

	// add(literal, sub(ref, mul(literal, ref)))
	exp := &substraitgo.ScalarFunction{
		FuncRef:    addRef,
		ID:         addID,
		OutputType: &substraitgo.Float64Type{},
		Args: []substraitgo.FuncArg{
			substraitgo.NewPrimitiveLiteral(float64(1.0), false),
			&substraitgo.ScalarFunction{
				FuncRef:    subRef,
				ID:         subID,
				OutputType: &substraitgo.Float32Type{},
				Args: []substraitgo.FuncArg{
					&substraitgo.FieldReference{
						Root: substraitgo.RootReference,
						Reference: &substraitgo.StructFieldRef{
							Field: 3,
						},
					},
					&substraitgo.ScalarFunction{
						FuncRef:    mulRef,
						ID:         mulID,
						OutputType: &substraitgo.Int64Type{},
						Args: []substraitgo.FuncArg{
							substraitgo.NewPrimitiveLiteral(int64(2), false),
							&substraitgo.FieldReference{
								Root: substraitgo.NewNestedLiteral(substraitgo.StructLiteralValue{
									substraitgo.NewByteSliceLiteral([]byte("baz"), true),
									substraitgo.NewPrimitiveLiteral("foobar", false),
									substraitgo.NewPrimitiveLiteral(int32(5), false),
								}, false),
								Reference: &substraitgo.StructFieldRef{
									Field: 2,
								},
							},
						},
					},
				},
			},
		},
	}

	return exp
}

func TestExpressionsRoundtrip(t *testing.T) {
	const substraitExtURI = "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	// define extensions with no plan for now
	const planExt = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "` + substraitExtURI + `"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 2,
					"name": "add"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 3,
					"name": "subtract"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 4,
					"name": "multiply"
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

	tests := []substraitgo.Expression{
		sampleNestedExpr(substraitExtURI),
		// TODO: add more nested field tests after parsing is implemented
		//       which will make it easier to generate nested expressions
		//       to test with.
	}

	for _, expr := range tests {
		protoExpr := expr.ToProto()
		out, err := substraitgo.ExprFromProto(protoExpr, nil, extSet)
		require.NoError(t, err)
		assert.Truef(t, expr.Equals(out), "expected: %s\ngot: %s", expr, out)
	}
}

func ExampleExpression_Visit() {
	const substraitExtURI = "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	var (
		expr                = sampleNestedExpr(substraitExtURI)
		preVisit, postVisit substraitgo.VisitFunc
	)

	preVisit = func(e substraitgo.Expression) substraitgo.Expression {
		fmt.Println(e)
		return e.Visit(preVisit)
	}
	postVisit = func(e substraitgo.Expression) substraitgo.Expression {
		out := e.Visit(postVisit)
		fmt.Println(e)
		return out
	}
	fmt.Println("PreOrder:")
	fmt.Println(expr.Visit(preVisit))
	fmt.Println()
	fmt.Println("PostOrder:")
	fmt.Println(expr.Visit(postVisit))

	// Output:
	// PreOrder:
	// fp64(1)
	// subtract(.field(3), multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)) => i64) => fp32
	// .field(3)
	// multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)) => i64
	// i64(2)
	// [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)
	// add(fp64(1), subtract(.field(3), multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)) => i64) => fp32) => fp64
	//
	// PostOrder:
	// fp64(1)
	// .field(3)
	// i64(2)
	// [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)
	// multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)) => i64
	// subtract(.field(3), multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)) => i64) => fp32
	// add(fp64(1), subtract(.field(3), multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2)) => i64) => fp32) => fp64
}
