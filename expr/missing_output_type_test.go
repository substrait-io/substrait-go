// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestScalarFunctionMissingOutputTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	functionReference := registry.GetFuncAnchor(extensions.ID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "add:i64_i64",
	})

	_, err := expr.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_ScalarFunction_{ScalarFunction: &proto.Expression_ScalarFunction{
			FunctionReference: functionReference,
			// OutputType intentionally omitted.
			Arguments: []*proto.FunctionArgument{
				literalI64Arg(1),
				literalI64Arg(2),
			},
		}},
	}, nil, registry)

	require.Error(t, err)
}

func TestWindowFunctionMissingOutputTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	functionReference := registry.GetFuncAnchor(extensions.ID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "sum:i64",
	})

	_, err := expr.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_WindowFunction_{WindowFunction: &proto.Expression_WindowFunction{
			FunctionReference: functionReference,
			// OutputType intentionally omitted.
		}},
	}, nil, registry)

	require.Error(t, err)
}

func TestCastMissingTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())

	_, err := expr.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_Cast_{Cast: &proto.Expression_Cast{
			// Type intentionally omitted.
			Input: &proto.Expression{
				RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
					LiteralType: &proto.Expression_Literal_I64{I64: 1},
				}},
			},
		}},
	}, nil, registry)

	require.Error(t, err)
}

func literalI64Arg(value int64) *proto.FunctionArgument {
	return &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Value{Value: &proto.Expression{
		RexType: &proto.Expression_Literal_{Literal: &proto.Expression_Literal{
			LiteralType: &proto.Expression_Literal_I64{I64: value},
		}},
	}}}
}
