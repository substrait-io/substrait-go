// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

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
