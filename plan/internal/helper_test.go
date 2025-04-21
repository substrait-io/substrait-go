package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v4/expr"
	ext "github.com/substrait-io/substrait-go/v4/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestVirtualTableExpressionFromProto(t *testing.T) {
	// define extensions with no plan for now
	const planExt = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
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
	extSet := ext.GetExtensionSet(&plan)
	literal1 := expr.NewPrimitiveLiteral(int32(1), false)
	expr1 := literal1.ToProto()

	reg := expr.NewExtensionRegistry(extSet, ext.GetDefaultCollectionWithNoError())
	rows := &proto.Expression_Nested_Struct{Fields: []*proto.Expression{
		expr1,
	}}
	exprRows, err := VirtualTableExpressionFromProto(rows, reg)
	require.NoError(t, err)
	require.Len(t, exprRows, 1)
}
