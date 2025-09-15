package plan

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v6/expr"
	"github.com/substrait-io/substrait-go/v6/extensions"
	"github.com/substrait-io/substrait-go/v6/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

// This is a test where we will load in an extension with both a URI and URN present.
// We then make sure that a produced plan contains the appropriate data present for both
// URI and URN (e.g. correct anchors for both present) by comparing the produced plan
// with an expected output plan. This test can be dropped when the URI -> URN migration
// is complete.
func TestExtensionURNAndURIInPlanProtobuf(t *testing.T) {

	const uri = "http://localhost/test.yaml"
	const extensionYAML = `---
urn: "urn:example:test"
scalar_functions:
  - name: "add"
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	var collection extensions.Collection
	require.NoError(t, collection.LoadWithURI(uri, strings.NewReader(extensionYAML)))

	builder := NewBuilder(&collection)
	baseSchema := types.NamedStruct{
		Names:  []string{"a", "b"},
		Struct: types.StructType{Types: []types.Type{&types.Int32Type{}, &types.Int32Type{}}},
	}
	scan := builder.NamedScan([]string{"test_table"}, baseSchema)
	exprBuilder := builder.GetExprBuilder()
	exprBuilder.BaseSchema = types.NewRecordTypeFromStruct(baseSchema.Struct)

	scalarFunc := exprBuilder.ScalarFunc(extensions.ID{
		URI:  uri,
		Name: "add",
	}).Args(
		exprBuilder.RootRef(expr.NewStructFieldRef(0)),
		exprBuilder.RootRef(expr.NewStructFieldRef(1)),
	)

	addExpr, err := scalarFunc.BuildExpr()
	require.NoError(t, err)

	projected, err := builder.Project(scan, addExpr)
	require.NoError(t, err)

	planObj, err := builder.Plan(projected, []string{"a", "b", "result"})
	require.NoError(t, err)

	gotProto, err := planObj.ToProto()
	require.NoError(t, err)

	const expectedJSON = `{
		"version": {
			"minorNumber": 29,
			"producer": "substrait-go darwin/arm64"
		},
		"extensionUris": [{
			"extensionUriAnchor": 1,
			"uri": "http://localhost/test.yaml"
		}],
		"extensionUrns": [{
			"extensionUrnAnchor": 1,
			"urn": "urn:example:test"
		}],
		"extensions": [{
			"extensionFunction": {
				"extensionUriReference": 1,
				"extensionUrnReference": 1,
				"functionAnchor": 1,
				"name": "add:i32_i32"
			}
		}],
		"relations": [{
			"root": {
				"input": {
					"project": {
						"common": {"direct": {}},
						"input": {
							"read": {
								"common": {"direct": {}},
								"baseSchema": {
									"names": ["a", "b"],
									"struct": {
										"types": [{"i32": {}}, {"i32": {}}]
									}
								},
								"namedTable": {"names": ["test_table"]}
							}
						},
						"expressions": [{
							"scalarFunction": {
								"functionReference": 1,
								"arguments": [
									{
										"value": {
											"selection": {
												"directReference": {"structField": {}},
												"rootReference": {}
											}
										}
									},
									{
										"value": {
											"selection": {
												"directReference": {"structField": {"field": 1}},
												"rootReference": {}
											}
										}
									}
								],
								"outputType": {
									"i32": {
										"nullability": "NULLABILITY_NULLABLE"
									}
								}
							}
						}]
					}
				},
				"names": ["a", "b", "result"]
			}
		}]
	}`

	var expected proto.Plan
	require.NoError(t, protojson.Unmarshal([]byte(expectedJSON), &expected))

	if diff := cmp.Diff(&expected, gotProto, protocmp.Transform()); diff != "" {
		t.Errorf("Plan protobuf mismatch (-want +got):\n%s", diff)
	}

	assert.True(t, collection.URILoaded(uri))
	assert.True(t, collection.URNLoaded("urn:example:test"))
}

func TestRelFromProto(t *testing.T) {

	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	literal5 := &proto.Expression_Literal{LiteralType: &proto.Expression_Literal_I64{I64: 5}}
	exprLiteral5 := &proto.Expression{RexType: &proto.Expression_Literal_{Literal: literal5}}

	nestedStructExpr1 := &proto.Expression_Nested_Struct{Fields: []*proto.Expression{exprLiteral5}}
	virtualTableWithExpression := &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{Expressions: []*proto.Expression_Nested_Struct{nestedStructExpr1}}}
	readRelWithExpression := &proto.ReadRel{ReadType: virtualTableWithExpression}

	literalStruct := &proto.Expression_Literal_Struct{Fields: []*proto.Expression_Literal{literal5}}
	virtualTableWithLiteral := &proto.ReadRel_VirtualTable_{VirtualTable: &proto.ReadRel_VirtualTable{Values: []*proto.Expression_Literal_Struct{literalStruct}}}
	readRelWithLiteral := &proto.ReadRel{ReadType: virtualTableWithLiteral}

	for _, td := range []struct {
		name     string
		readType *proto.ReadRel
	}{
		{"virtual table with expression", readRelWithExpression},
		{"virtual table with deprecated literal", readRelWithLiteral},
	} {
		t.Run(td.name, func(t *testing.T) {
			rel := &proto.Rel{RelType: &proto.Rel_Read{Read: td.readType}}

			outRel, err := RelFromProto(rel, registry)
			require.NoError(t, err)
			gotRel := outRel.ToProto()
			gotReadRel, ok := gotRel.RelType.(*proto.Rel_Read)
			require.True(t, ok)
			gotVirtualTableReadRel, ok := gotReadRel.Read.ReadType.(*proto.ReadRel_VirtualTable_)
			require.True(t, ok)
			// in case of both deprecated or new expression, the output should be the same as the new expression
			if diff := cmp.Diff(gotVirtualTableReadRel, virtualTableWithExpression, protocmp.Transform()); diff != "" {
				t.Errorf("expression proto didn't match, diff:\n%v", diff)
			}
		})
	}

}
