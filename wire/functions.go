// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func scalarFunctionToProto(s *expr.ScalarFunction) *proto.Expression {
	args := make([]*proto.FunctionArgument, s.NArgs())
	for i := range args {
		args[i] = FuncArgToProto(s.Arg(i))
	}

	return &proto.Expression{
		RexType: &proto.Expression_ScalarFunction_{
			ScalarFunction: &proto.Expression_ScalarFunction{
				FunctionReference: s.FuncRef(),
				Options:           FunctionOptionsToProto(s.GetOptions()),
				OutputType:        TypeToProto(s.GetType()),
				Arguments:         args,
			},
		},
	}
}
