// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v3/expr"
	ext "github.com/substrait-io/substrait-go/v3/extensions"
	"github.com/substrait-io/substrait-go/v3/proto"
	"github.com/substrait-io/substrait-go/v3/types"
	"github.com/substrait-io/substrait-go/v3/types/parser"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

const sampleYAML = `---
scalar_functions:
  -
    name: "add"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i8
          - name: y
            value: i8
        options:
          overflow:
            values: [ SILENT, SATURATE, ERROR ]
        return: i8
`

var collection ext.Collection

func init() {
	err := collection.Load("https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml", strings.NewReader(sampleYAML))
	if err != nil {
		panic(err)
	}
}

func ExampleExpression_scalarFunction() {
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

	// json proto to represent of add(field_ref(0), float64(10))
	const scalarFunction = `{
		"scalarFunction": {
		  "functionReference": 2,
		  "outputType": {"i32": {}},
		  "arguments": [
			{"value": {"selection": {
				"rootReference": {},
				"directReference": {"structField": {"field": 0}}}}},
			{"value": {"literal": {"fp64": 10}}}
		  ]
		}
	  }`

	var exprProto proto.Expression
	if err := protojson.Unmarshal([]byte(scalarFunction), &exprProto); err != nil {
		panic(err)
	}

	reg := expr.NewExtensionRegistry(extSet, &collection)
	// convert from protobuf to Expression!
	fromProto, err := expr.ExprFromProto(&exprProto, nil, reg)
	if err != nil {
		panic(err)
	}

	// manually define the entire expression instead of going through
	// having to construct the protobuf
	const substraitext = `https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml`

	var addVariant = ext.NewScalarFuncVariant(ext.ID{URI: substraitext, Name: "add:i32_i32"})

	var ex expr.Expression
	refArg, _ := expr.NewRootFieldRef(expr.NewStructFieldRef(0), types.NewRecordTypeFromTypes([]types.Type{&types.Int32Type{}}))
	ex, _ = expr.NewCustomScalarFunc(reg, addVariant, &types.Int32Type{}, nil,
		refArg, expr.NewPrimitiveLiteral(float64(10), false))

	// call ToProto to convert our manual expression to proto.Expression
	toProto := ex.ToProto()

	// output some info!

	// print string represention of the expression
	fmt.Println(fromProto)
	// print the string representation of our
	// manually constructed expression
	fmt.Println(ex)

	// verify that the Equals methods work recursively
	fmt.Println(ex.Equals(fromProto))
	// confirm our manually constructed expression is the same
	// as the one we got from protojson
	fmt.Println(pb.Equal(&exprProto, toProto))

	// Output:
	// add(.field(0), fp64(10)) => i32
	// add(.field(0) => i32, fp64(10)) => i32
	// true
	// true
}

func sampleNestedExpr(reg expr.ExtensionRegistry, substraitExtURI string) expr.Expression {
	var (
		add = ext.NewScalarFuncVariant(ext.ID{URI: substraitExtURI, Name: "add"})
		sub = ext.NewScalarFuncVariant(ext.ID{URI: substraitExtURI, Name: "subtract"})
		mul = ext.NewScalarFuncVariant(ext.ID{URI: substraitExtURI, Name: "multiply"})
	)

	baseSchema := types.NewRecordTypeFromTypes(
		[]types.Type{
			&types.BooleanType{},
			&types.Int32Type{},
			&types.Int64Type{},
			&types.Float32Type{},
		})

	// add(literal, sub(ref, mul(literal, ref)))
	exp := expr.MustExpr(expr.NewCustomScalarFunc(reg, add, &types.Float64Type{}, nil,
		expr.NewPrimitiveLiteral(float64(1.0), false),
		expr.MustExpr(expr.NewCustomScalarFunc(reg, sub, &types.Float32Type{}, nil,
			expr.MustExpr(expr.NewRootFieldRef(expr.NewStructFieldRef(3), baseSchema)),
			expr.MustExpr(expr.NewCustomScalarFunc(reg, mul, &types.Int64Type{}, nil,
				expr.NewPrimitiveLiteral(int64(2), false),
				expr.MustExpr(expr.NewFieldRef(expr.NewNestedLiteral(expr.StructLiteralValue{
					expr.NewByteSliceLiteral([]byte("baz"), true),
					expr.NewPrimitiveLiteral("foobar", false),
					expr.NewPrimitiveLiteral(int32(5), false),
				}, false), expr.NewStructFieldRef(2), nil)),
			)),
		)),
	))

	return exp
}

func TestExpressionsRoundtrip(t *testing.T) {
	const substraitExtURI = "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	// define extensions with no plan for now
	const planExt = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "` + substraitExtURI + `"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 2,
					"name": "add:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 3,
					"name": "subtract:fp32_fp32"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 4,
					"name": "multiply:i64_i64"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 5,
					"name": "ntile:"
				}
			}
		],
		"relations": []
	}`

	var (
		plan proto.Plan
	)
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}
	// get the extension set
	extSet := ext.GetExtensionSet(&plan)
	reg := expr.NewExtensionRegistry(extSet, ext.GetDefaultCollectionWithNoError())
	tests := []expr.Expression{
		sampleNestedExpr(reg, substraitExtURI),
	}

	for _, exp := range tests {
		protoExpr := exp.ToProto()
		out, err := expr.ExprFromProto(protoExpr, nil, reg)
		require.NoError(t, err)
		assert.Truef(t, exp.Equals(out), "expected: %s\ngot: %s", exp, out)
	}
}

func ExampleExpression_Visit() {
	const substraitExtURI = "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	var (
		exp                 = sampleNestedExpr(expr.NewEmptyExtensionRegistry(ext.GetDefaultCollectionWithNoError()), substraitExtURI)
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
	// fp64(1)
	// subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32
	// .field(3) => fp32
	// multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64
	// i64(2)
	// [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32
	// add(fp64(1), subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32) => fp64
	//
	// PostOrder:
	// fp64(1)
	// .field(3) => fp32
	// i64(2)
	// [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32
	// multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64
	// subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32
	// add(fp64(1), subtract(.field(3) => fp32, multiply(i64(2), [root:(struct<binary?, string, i32>([binary?([98 97 122]) string(foobar) i32(5)]))].field(2) => i32) => i64) => fp32) => fp64
}

func TestRoundTripUsingTestData(t *testing.T) {
	const substraitExtURI = "https://github.com/substrait-io/substrait/blob/main/extensions/functions_arithmetic.yaml"
	// define extensions with no plan for now
	const planExt = `{
		"extensionUris": [
			{
				"extensionUriAnchor": 1,
				"uri": "` + substraitExtURI + `"
			}
		],
		"extensions": [
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 2,
					"name": "add:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 3,
					"name": "subtract:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 4,
					"name": "multiply:fp64_fp64"
				}
			},
			{
				"extensionFunction": {
					"extensionUriReference": 1,
					"functionAnchor": 5,
					"name": "ntile:i32"
				}
			}
		],
		"relations": []
	}`

	var (
		plan proto.Plan
	)
	if err := protojson.Unmarshal([]byte(planExt), &plan); err != nil {
		panic(err)
	}
	// get the extension set
	extSet := ext.GetExtensionSet(&plan)

	f, err := os.Open("./testdata/expressions.yaml")
	require.NoError(t, err)
	defer f.Close()

	dec := yaml.NewDecoder(f)
	var tmp map[string]any
	require.NoError(t, dec.Decode(&tmp))

	var (
		protoSchema proto.NamedStruct
	)

	raw, err := json.Marshal(tmp["baseSchema"])
	require.NoError(t, err)
	require.NoError(t, protojson.Unmarshal(raw, &protoSchema))
	baseSchema := types.NewNamedStructFromProto(&protoSchema)
	reg := expr.NewExtensionRegistry(extSet, ext.GetDefaultCollectionWithNoError())
	for _, tc := range tmp["cases"].([]any) {
		tt := tc.(map[string]any)
		t.Run(tt["name"].(string), func(t *testing.T) {
			test := tt["__test"].(map[string]any)

			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			require.NoError(t, enc.Encode(tt["expression"]))
			var ex proto.Expression
			require.NoError(t, protojson.Unmarshal(buf.Bytes(), &ex))

			e, err := expr.ExprFromProto(&ex, types.NewRecordTypeFromStruct(baseSchema.Struct), reg)
			require.NoError(t, err)

			result := e.ToProto()
			assert.Truef(t, pb.Equal(&ex, result), "expected: %s\ngot: %s", &ex, result)

			assert.True(t, e.Equals(e))

			if typTest, ok := test["type"].(string); ok {
				exp, err := parser.ParseType(typTest)
				require.NoError(t, err)

				assert.Equal(t, exp.String(), e.GetType().String())
			}

			strvalue, ok := test["string"].(string)
			if ok {
				strvalue = strings.TrimSpace(strvalue)
				t.Run("tostring", func(t *testing.T) {
					assert.Equal(t, strvalue, e.String())
				})
			}
		})
	}
}

func TestRoundTripExtendedExpression(t *testing.T) {
	f, err := os.Open("./testdata/extended_exprs.yaml")
	require.NoError(t, err)
	defer f.Close()

	dec := yaml.NewDecoder(f)
	var tmp map[string]any
	require.NoError(t, dec.Decode(&tmp))

	for _, tc := range tmp["tests"].([]any) {
		tt := tc.(map[string]any)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		require.NoError(t, enc.Encode(tt))
		var ex proto.ExtendedExpression
		require.NoError(t, protojson.Unmarshal(buf.Bytes(), &ex))

		result, err := expr.ExtendedFromProto(&ex, ext.GetDefaultCollectionWithNoError())
		require.NoError(t, err)

		out := result.ToProto()
		// because we read the extensions into a map, we can't guarantee
		// the order of the extensions. But we also don't care about the
		// order, so we can just sort them by functionAnchor to ensure
		// they match for pb.Equal
		sort.Slice(out.Extensions, func(i, j int) bool {
			return out.Extensions[i].GetExtensionFunction().FunctionAnchor <
				out.Extensions[j].GetExtensionFunction().FunctionAnchor
		})
		assert.Truef(t, pb.Equal(&ex, out), "expected: %s\ngot: %s", &ex, out)
	}
}
