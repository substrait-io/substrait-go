// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	pb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestConditionalDecoderRejectsInvalidConstruction(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	for name, input := range map[string]string{
		"if without branches":            `{"ifThen":{"else":{"literal":{"i32":1}}}}`,
		"if mismatched result":           `{"ifThen":{"ifs":[{"if":{"literal":{"boolean":true}},"then":{"literal":{"i8":1}}}],"else":{"literal":{"i32":1}}}}`,
		"if mismatched later result":     `{"ifThen":{"ifs":[{"if":{"literal":{"boolean":true}},"then":{"literal":{"i32":1}}},{"if":{"literal":{"boolean":false}},"then":{"literal":{"i8":1}}}],"else":{"literal":{"i32":1}}}}`,
		"if mismatched nullability":      `{"ifThen":{"ifs":[{"if":{"literal":{"boolean":true}},"then":{"literal":{"i32":2,"nullable":true}}}],"else":{"literal":{"i32":3}}}}`,
		"if without else":                `{"ifThen":{"ifs":[{"if":{"literal":{"boolean":true}},"then":{"literal":{"i32":1}}}]}}`,
		"if without condition":           `{"ifThen":{"ifs":[{"then":{"literal":{"i32":1}}}],"else":{"literal":{"i32":1}}}}`,
		"switch without branches":        `{"switchExpression":{"match":{"literal":{"i32":1}},"else":{"literal":{"i32":1}}}}`,
		"switch mismatched label":        `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i8":1},"then":{"literal":{"i32":1}}}],"else":{"literal":{"i32":1}}}}`,
		"switch mismatched result":       `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i8":1}}}],"else":{"literal":{"i32":1}}}}`,
		"switch mismatched later result": `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":1}}},{"if":{"i32":2},"then":{"literal":{"i8":1}}}],"else":{"literal":{"i32":1}}}}`,
		"switch mismatched nullability":  `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":2,"nullable":true}}}],"else":{"literal":{"i32":3}}}}`,
		"switch later nullability":       `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":2}}},{"if":{"i32":2},"then":{"literal":{"i32":3,"nullable":true}}}],"else":{"literal":{"i32":4}}}}`,
		"switch without else":            `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":2}}}]}}`,
		"switch invalid else":            `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":2}}}],"else":{"literal":{"precisionTime":{"precision":6,"value":"-1"}}}}}`,
		"switch without match":           `{"switchExpression":{"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":1}}}],"else":{"literal":{"i32":1}}}}`,
		"switch without label":           `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"then":{"literal":{"i32":1}}}],"else":{"literal":{"i32":1}}}}`,
	} {
		t.Run(name, func(t *testing.T) {
			var inputProto pb.Expression
			require.NoError(t, protojson.Unmarshal([]byte(input), &inputProto))
			var decoded expr.Expression
			var err error
			require.NotPanics(t, func() {
				decoded, err = substraitproto.ExprFromProto(&inputProto, nil, registry)
			})
			require.Nil(t, decoded)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
		})
	}
}

func TestConditionalDecoderRoundTrips(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	for name, input := range map[string]string{
		"if":                    `{"ifThen":{"ifs":[{"if":{"literal":{"boolean":true}},"then":{"literal":{"i32":1}}},{"if":{"literal":{"boolean":false}},"then":{"literal":{"i32":2}}}],"else":{"literal":{"i32":3}}}}`,
		"switch":                `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":2}}},{"if":{"i32":2},"then":{"literal":{"i32":3}}}],"else":{"literal":{"i32":4}}}}`,
		"switch nullable types": `{"switchExpression":{"match":{"literal":{"i32":1}},"ifs":[{"if":{"i32":1},"then":{"literal":{"i32":2,"nullable":true}}}],"else":{"literal":{"i32":3,"nullable":true}}}}`,
	} {
		t.Run(name, func(t *testing.T) {
			var inputProto pb.Expression
			require.NoError(t, protojson.Unmarshal([]byte(input), &inputProto))
			decoded, err := substraitproto.ExprFromProto(&inputProto, nil, registry)
			require.NoError(t, err)
			require.True(t, proto.Equal(&inputProto, substraitproto.ExprToProto(decoded)))
			require.NotNil(t, decoded.GetType())
		})
	}
}

func TestConditionalDecoderRejectsMissingMessages(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	value := &pb.Expression{RexType: &pb.Expression_Literal_{Literal: &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_I32{I32: 1}}}}
	for name, input := range map[string]*pb.Expression{
		"nil if wrapper":     {RexType: (*pb.Expression_IfThen_)(nil)},
		"nil if payload":     {RexType: &pb.Expression_IfThen_{}},
		"nil if clause":      {RexType: &pb.Expression_IfThen_{IfThen: &pb.Expression_IfThen{Else: value, Ifs: []*pb.Expression_IfThen_IfClause{nil}}}},
		"nil switch wrapper": {RexType: (*pb.Expression_SwitchExpression_)(nil)},
		"nil switch payload": {RexType: &pb.Expression_SwitchExpression_{}},
		"nil switch clause":  {RexType: &pb.Expression_SwitchExpression_{SwitchExpression: &pb.Expression_SwitchExpression{Match: value, Else: value, Ifs: []*pb.Expression_SwitchExpression_IfValue{nil}}}},
	} {
		t.Run(name, func(t *testing.T) {
			var decoded expr.Expression
			var err error
			require.NotPanics(t, func() {
				decoded, err = substraitproto.ExprFromProto(input, nil, registry)
			})
			require.Nil(t, decoded)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
		})
	}
}
