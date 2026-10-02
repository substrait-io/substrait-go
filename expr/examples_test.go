// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
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
	extSet, err := wire.GetExtensionSet(&plan, collection)
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
	fromProto, err := wire.ExprFromProto(&exprProto, nil, reg)
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
	toProto := wire.ExprToProto(ex)

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

func sampleNestedExpr(reg expr.ExtensionRegistry, substraitExtURN string) expr.Expression {
	var (
		add = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitExtURN, Name: "add"})
		sub = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitExtURN, Name: "subtract"})
		mul = ext.NewScalarFuncVariant(ext.FunctionID{URN: substraitExtURN, Name: "multiply"})
	)

	baseSchema := types.NewRecordTypeFromTypes(
		[]types.Type{
			&types.BooleanType{},
			&types.Int32Type{},
			&types.Int64Type{},
			&types.Float32Type{},
		})

	// add(literal, sub(ref, mul(literal, ref)))
	exp := expr.MustExpr(expr.NewCustomScalarFunc(reg, add, &types.Float64Type{}, nil,
		expr.NewPrimitiveLiteral(float64(1.0), false),
		expr.MustExpr(expr.NewCustomScalarFunc(reg, sub, &types.Float32Type{}, nil,
			expr.MustExpr(expr.NewRootFieldRef(expr.NewStructFieldRef(3), baseSchema)),
			expr.MustExpr(expr.NewCustomScalarFunc(reg, mul, &types.Int64Type{}, nil,
				expr.NewPrimitiveLiteral(int64(2), false),
				expr.MustExpr(expr.NewFieldRef(expr.NewNestedLiteral(expr.StructLiteralValue{
					expr.NewByteSliceLiteral([]byte("baz"), true),
					expr.NewPrimitiveLiteral("foobar", false),
					expr.NewPrimitiveLiteral(int32(5), false),
				}, false), expr.NewStructFieldRef(2), nil)),
			)),
		)),
	))

	return exp
}

func ExampleExpression_Visit() {
	const substraitExtURN = "extension:io.substrait:functions_arithmetic"
	var (
		exp                 = sampleNestedExpr(expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError()), substraitExtURN)
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
	// subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32
	// .field(3) => fp32
	// multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64
	// i64(2)
	// [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32
	// add(fp64(1), subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32) => fp64
	//
	// PostOrder:
	// fp64(1)
	// .field(3) => fp32
	// i64(2)
	// [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32
	// multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64
	// subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32
	// add(fp64(1), subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32) => fp64
}
