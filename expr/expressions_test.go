// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/expr"
	ext "github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

const sampleYAML = `---
scalar_functions:
  -
    name: "add"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i8
          - name: y
            value: i8
        options:
          overflow:
            values: [ SILENT, SATURATE, ERROR ]
        return: i8
`

var collection ext.Collection

func init() {
	collection.Load("https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml", strings.NewReader(sampleYAML))
}

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

	var plan types.Plan
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}

	// get the extension set
	extSet := ext.GetExtensionSet(&plan)

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
	fromProto, err := expr.ExprFromProto(&exprProto, nil, extSet, &collection)
	if err != nil {
		panic(err)
	}

	// manually define the entire expression instead of going through
	// having to construct the protobuf
	const substraitext = `https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml`

	var expr expr.Expression = &expr.ScalarFunction{
		FuncRef: 2,
		ID:      ext.ID{URI: substraitext, Name: "add"},
		Args: []types.FuncArg{
			&expr.FieldReference{
				Root:      expr.RootReference,
				Reference: &expr.StructFieldRef{Field: 0},
			},
			expr.NewPrimitiveLiteral(float64(10), false),
		},
		OutputType: &types.Int32Type{},
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

func sampleNestedExpr(substraitExtURI string) expr.Expression {
	var (
		addID         = ext.ID{URI: substraitExtURI, Name: "add"}
		addRef uint32 = 2
		subID         = ext.ID{URI: substraitExtURI, Name: "subtract"}
		subRef uint32 = 3
		mulID         = ext.ID{URI: substraitExtURI, Name: "multiply"}
		mulRef uint32 = 4
	)

	// add(literal, sub(ref, mul(literal, ref)))
	exp := &expr.ScalarFunction{
		FuncRef:    addRef,
		ID:         addID,
		OutputType: &types.Float64Type{},
		Args: []types.FuncArg{
			expr.NewPrimitiveLiteral(float64(1.0), false),
			&expr.ScalarFunction{
				FuncRef:    subRef,
				ID:         subID,
				OutputType: &types.Float32Type{},
				Args: []types.FuncArg{
					&expr.FieldReference{
						Root: expr.RootReference,
						Reference: &expr.StructFieldRef{
							Field: 3,
						},
					},
					&expr.ScalarFunction{
						FuncRef:    mulRef,
						ID:         mulID,
						OutputType: &types.Int64Type{},
						Args: []types.FuncArg{
							expr.NewPrimitiveLiteral(int64(2), false),
							&expr.FieldReference{
								Root: expr.NewNestedLiteral(expr.StructLiteralValue{
									expr.NewByteSliceLiteral([]byte("baz"), true),
									expr.NewPrimitiveLiteral("foobar", false),
									expr.NewPrimitiveLiteral(int32(5), false),
								}, false),
								Reference: &expr.StructFieldRef{
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

	var (
		plan            types.Plan
		emptyCollection ext.Collection
	)
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}
	// get the extension set
	extSet := ext.GetExtensionSet(&plan)

	tests := []expr.Expression{
		sampleNestedExpr(substraitExtURI),
		// TODO: add more nested field tests after parsing is implemented
		//       which will make it easier to generate nested expressions
		//       to test with.
	}

	for _, exp := range tests {
		protoExpr := exp.ToProto()
		out, err := expr.ExprFromProto(protoExpr, nil, extSet, &emptyCollection)
		require.NoError(t, err)
		assert.Truef(t, exp.Equals(out), "expected: %s\ngot: %s", exp, out)
	}
}

func ExampleExpression_Visit() {
	const substraitExtURI = "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	var (
		exp                 = sampleNestedExpr(substraitExtURI)
		preVisit, postVisit expr.VisitFunc
	)

	preVisit = func(e expr.Expression) expr.Expression {
		fmt.Println(e)
		return e.Visit(preVisit)
	}
	postVisit = func(e expr.Expression) expr.Expression {
		out := e.Visit(postVisit)
		fmt.Println(e)
		return out
	}
	fmt.Println("PreOrder:")
	fmt.Println(exp.Visit(preVisit))
	fmt.Println()
	fmt.Println("PostOrder:")
	fmt.Println(exp.Visit(postVisit))

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
