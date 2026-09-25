// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// FuncArgToProto encodes a function argument as its protobuf message.
func FuncArgToProto(a types.FuncArg) *proto.FunctionArgument {
	switch a := a.(type) {
	case types.Enum:
		return &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Enum{Enum: string(a)}}
	case types.Type:
		return &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{Type: TypeToProto(a)}}
	case expr.Expression:
		return &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Value{Value: ExprToProto(a)}}
	default:
		panic(fmt.Sprintf("wire: unhandled function argument %T", a))
	}
}
