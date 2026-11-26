// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

func sampleNestedExpr(reg expr.ExtensionRegistry, substraitExtURN string) expr.Expression {
	function := func(name string) ext.FunctionID {
		return ext.FunctionID{URN: substraitExtURN, Name: name}
	}

	baseSchema := types.NewRecordTypeFromTypes(
		[]types.Type{
			&types.BooleanType{},
			&types.Int32Type{},
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.Int64Type{Nullability: types.NullabilityRequired},
		})

	// add(literal, sub(ref, mul(literal, ref)))
	exp := expr.MustExpr(expr.NewScalarFunc(reg, function("add:i64_i64"), nil,
		expr.NewPrimitiveLiteral(int64(1), false),
		expr.MustExpr(expr.NewScalarFunc(reg, function("subtract:i64_i64"), nil,
			expr.MustExpr(expr.NewRootFieldRef(expr.NewStructFieldRef(3), baseSchema)),
			expr.MustExpr(expr.NewScalarFunc(reg, function("multiply:i64_i64"), nil,
				expr.NewPrimitiveLiteral(int64(2), false),
				expr.MustExpr(expr.NewFieldRef(expr.NewNestedLiteral(expr.StructLiteralValue{
					expr.NewByteSliceLiteral([]byte("baz"), true),
					expr.NewPrimitiveLiteral("foobar", false),
					expr.NewPrimitiveLiteral(int64(5), false),
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
	// i64(1)
	// subtract(.field(3) => i64, multiply(i64(2), [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64) => i64) => i64
	// .field(3) => i64
	// multiply(i64(2), [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64) => i64
	// i64(2)
	// [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64
	// add(i64(1), subtract(.field(3) => i64, multiply(i64(2), [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64) => i64) => i64) => i64
	//
	// PostOrder:
	// i64(1)
	// .field(3) => i64
	// i64(2)
	// [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64
	// multiply(i64(2), [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64) => i64
	// subtract(.field(3) => i64, multiply(i64(2), [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64) => i64) => i64
	// add(i64(1), subtract(.field(3) => i64, multiply(i64(2), [root:(struct<binary?, string, i64>([binary?([98 97 122]) string(foobar) i64(5)]))].field(2) => i64) => i64) => i64) => i64
}
