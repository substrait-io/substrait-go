// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"bytes"
	"encoding/json"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

func TestScalarFunctionMissingOutputTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	functionReference := registry.GetFuncAnchor(ext.FunctionID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "add:i64_i64",
	})

	_, err := wire.ExprFromProto(&proto.Expression{
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
	registry := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())
	functionReference := registry.GetFuncAnchor(ext.FunctionID{
		URN:  "extension:io.substrait:functions_arithmetic",
		Name: "sum:i64",
	})

	_, err := wire.ExprFromProto(&proto.Expression{
		RexType: &proto.Expression_WindowFunction_{WindowFunction: &proto.Expression_WindowFunction{
			FunctionReference: functionReference,
			// OutputType intentionally omitted.
		}},
	}, nil, registry)

	require.Error(t, err)
}

func TestCastMissingTypeReturnsError(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError())

	_, err := wire.ExprFromProto(&proto.Expression{
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

func TestCastVisit(t *testing.T) {
	var builder = plan.NewBuilderDefault()
	castExpr := expr.MustExpr(builder.GetExprBuilder().Cast(builder.GetExprBuilder().Wrap(
		expr.NewLiteral[float64](12.0, true)),
		&types.Float64Type{Nullability: types.NullabilityRequired}).FailBehavior(
		types.CastFailBehaviorThrowException).BuildExpr())

	type relationTestCase struct {
		name            string
		rewriteFunction func(rex expr.Expression) expr.Expression
		want            float64
	}
	testCases := []relationTestCase{
		{"no change", func(ex expr.Expression) expr.Expression { return ex }, 12},
		{"changed", func(ex expr.Expression) expr.Expression {
			lit, err := expr.NewLiteral[float64](16.0, true)
			require.NoError(t, err)
			return lit
		}, 16},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			visitedCastExpr := castExpr.Visit(tc.rewriteFunction)
			visitedCastProto := wire.ExprToProto(visitedCastExpr)
			assert.IsType(t, &proto.Expression_Cast_{}, visitedCastProto.GetRexType())
			assert.Equal(t, tc.want, visitedCastProto.GetCast().GetInput().GetLiteral().GetFp64())
		})
	}
}
