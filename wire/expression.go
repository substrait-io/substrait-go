// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// ExprToProto encodes an expression as its protobuf message.
func ExprToProto(e expr.Expression) *proto.Expression {
	switch e := e.(type) {
	case expr.Literal:
		return &proto.Expression{
			RexType: &proto.Expression_Literal_{Literal: LiteralToProto(e)},
		}
	default:
		panic(fmt.Sprintf("wire: unhandled expression %T", e))
	}
}

// VirtualTableExpressionValueToProto encodes a virtual-table row of expressions.
func VirtualTableExpressionValueToProto(s expr.VirtualTableExpressionValue) *proto.Expression_Nested_Struct {
	fields := make([]*proto.Expression, len(s))
	for i, f := range s {
		fields[i] = ExprToProto(f)
	}
	return &proto.Expression_Nested_Struct{Fields: fields}
}

// FuncArgFromProto decodes a function argument from its protobuf message.
func FuncArgFromProto(e *proto.FunctionArgument, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (types.FuncArg, error) {
	switch et := e.ArgType.(type) {
	case *proto.FunctionArgument_Enum:
		return types.Enum(et.Enum), nil
	case *proto.FunctionArgument_Type:
		return TypeFromProto(et.Type), nil
	case *proto.FunctionArgument_Value:
		return ExprFromProto(et.Value, baseSchema, reg)
	}
	return nil, substraitgo.ErrNotImplemented
}

// ExprFromProto decodes an expression from its protobuf message.
func ExprFromProto(e *proto.Expression, baseSchema *types.RecordType, reg expr.ExtensionRegistry) (expr.Expression, error) {
	if e == nil {
		return nil, fmt.Errorf("%w: protobuf Expression is nil", substraitgo.ErrInvalidExpr)
	}

	switch et := e.RexType.(type) {
	case *proto.Expression_Literal_:
		return LiteralFromProto(et.Literal), nil
	case *proto.Expression_Enum_:
		return nil, fmt.Errorf("%w: deprecated", substraitgo.ErrNotImplemented)
	}
	return nil, fmt.Errorf("%w: ExprFromProto: %s", substraitgo.ErrNotImplemented, e)
}
