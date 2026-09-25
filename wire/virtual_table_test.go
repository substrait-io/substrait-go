package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestVirtualTableExpressionFromProto(t *testing.T) {
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

	// get the extension set
	collection := ext.GetDefaultCollectionWithNoError()
	extSet, err := GetExtensionSet(&plan, collection)
	require.NoError(t, err)
	expr1 := &proto.Expression{RexType: &proto.Expression_Literal_{
		Literal: &proto.Expression_Literal{LiteralType: &proto.Expression_Literal_I32{I32: 1}}}}

	reg := expr.NewExtensionRegistry(extSet, collection)
	rows := &proto.Expression_Nested_Struct{Fields: []*proto.Expression{
		expr1,
	}}
	exprRows, err := virtualTableExpressionFromProto(rows, reg)
	require.NoError(t, err)
	require.Len(t, exprRows, 1)
}
