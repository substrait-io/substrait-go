// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"embed"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait"
	"github.com/substrait-io/substrait-go/v3/expr"
	"github.com/substrait-io/substrait-go/v3/extensions"
	"github.com/substrait-io/substrait-go/v3/functions"
	parser2 "github.com/substrait-io/substrait-go/v3/testcases/parser"
	"github.com/substrait-io/substrait-go/v3/types"
	"github.com/substrait-io/substrait-go/v3/types/integer_parameters"
	"github.com/substrait-io/substrait-go/v3/types/parser"
)

func TestEvaluateTypeExpression(t *testing.T) {
	var (
		i64Null, _    = parser.ParseType("i64?")
		i64NonNull, _ = parser.ParseType("i64")
		strNull, _    = parser.ParseType("string?")
	)

	tests := []struct {
		name     string
		nulls    extensions.NullabilityHandling
		ret      types.FuncDefArgType
		extArgs  extensions.FuncParameterList
		args     []types.Type
		expected types.Type
		err      string
	}{
		{"defaults", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable}, ""},
		{"arg mismatch", extensions.MirrorNullability, strNull, extensions.FuncParameterList{extensions.ValueArg{Value: &parser.TypeExpression{ValueType: strNull}}},
			[]types.Type{}, nil, "invalid expression: mismatch in number of arguments provided. got 0, expected 1"},
		{"missing enum arg", extensions.MirrorNullability, i64Null, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64NonNull}}, extensions.EnumArg{Name: "foo"}},
			[]types.Type{&types.Int64Type{}, &types.Int64Type{}}, nil, "invalid type: arg #1 (foo) should be an enum"},
		{"discrete null handling", extensions.DiscreteNullability, strNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: strNull}}},
			[]types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			nil, "invalid type: discrete nullability did not match for arg #0"},
		{"mirror", extensions.MirrorNullability, strNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64NonNull}}, extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{
				&types.Int64Type{Nullability: types.NullabilityRequired},
				&types.Int64Type{Nullability: types.NullabilityRequired}},
			&types.StringType{Nullability: types.NullabilityRequired}, ""},
		{"declared output", extensions.DeclaredOutputNullability, strNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: strNull}}},
			[]types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			&types.StringType{Nullability: types.NullabilityNullable}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extensions.EvaluateTypeExpression(tt.nulls, tt.ret, tt.extArgs, nil, tt.args)
			if tt.err == "" {
				assert.NoError(t, err)
				assert.Truef(t, tt.expected.Equals(result), "expected: %s\ngot: %s", tt.expected, result)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}

func TestVariantWithVariadic(t *testing.T) {
	var (
		i64Null, _     = parser.ParseType("i64?")
		i64NonNull, _  = parser.ParseType("i64")
		varcharNull, _ = parser.ParseType("varchar?<20>")
	)

	tests := []struct {
		name     string
		nulls    extensions.NullabilityHandling
		ret      types.FuncDefArgType
		extArgs  extensions.FuncParameterList
		args     []types.Type
		expected types.Type
		variadic extensions.VariadicBehavior
		err      string
	}{
		{"1param2args", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable},
				&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, ""},
		{"1param1arg", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, ""},
		{"1param0args", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{},
			&types.Int64Type{Nullability: types.NullabilityRequired},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, ""},
		{"2params3args", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: varcharNull}},
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.VarCharType{Nullability: types.NullabilityNullable, Length: 20},
				&types.Int64Type{Nullability: types.NullabilityNullable},
				&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, ""},
		{"2params2args", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: varcharNull}},
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.VarCharType{Nullability: types.NullabilityNullable, Length: 20},
				&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, ""},
		{"2params1arg", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: varcharNull}},
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.VarCharType{Nullability: types.NullabilityNullable, Length: 20}},
			&types.Int64Type{Nullability: types.NullabilityNullable},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, ""},
		{"2params1argBad", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: varcharNull}},
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			&types.Int64Type{Nullability: types.NullabilityNullable},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, "invalid type: argument types did not match"},
		{"2params0argsBad", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: varcharNull}},
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{},
			&types.Int64Type{Nullability: types.NullabilityRequired},
			extensions.VariadicBehavior{
				Min: 0, ParameterConsistency: extensions.ConsistentParams}, "invalid expression: mismatch in number of arguments provided. got 0, expected at least 1"},
		{"min2Variadic1ArgBad", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			nil, extensions.VariadicBehavior{
				Min: 2, ParameterConsistency: extensions.ConsistentParams},
			"invalid expression: mismatch in number of arguments provided, invalid number of variadic params. got 1 total"},
		{"min2Variadic1ArgBad", extensions.MirrorNullability, i64NonNull, extensions.FuncParameterList{
			extensions.ValueArg{Value: &parser.TypeExpression{ValueType: i64Null}}},
			[]types.Type{},
			nil, extensions.VariadicBehavior{
				Min: 2, ParameterConsistency: extensions.ConsistentParams},
			"invalid expression: mismatch in number of arguments provided, invalid number of variadic params. got 0 total"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extensions.EvaluateTypeExpression(tt.nulls, tt.ret, tt.extArgs, &tt.variadic, tt.args)
			if tt.err == "" {
				require.NoError(t, err)
				assert.Truef(t, tt.expected.Equals(result), "expected: %s\ngot: %s", tt.expected, result)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}

func TestHasSyncParams(t *testing.T) {

	apt_P := integer_parameters.NewVariableIntParam("P")
	apt_Q := integer_parameters.NewVariableIntParam("Q")
	cpt_38 := integer_parameters.NewConcreteIntParam(38)

	fct_P := &types.ParameterizedFixedCharType{IntegerOption: apt_P}
	fct_Q := &types.ParameterizedFixedCharType{IntegerOption: apt_Q}
	decimal_PQ := &types.ParameterizedDecimalType{Precision: apt_P, Scale: apt_Q}
	decimal_38_Q := &types.ParameterizedDecimalType{Precision: cpt_38, Scale: apt_Q}
	list_decimal_38_Q := &types.ParameterizedListType{Type: decimal_38_Q}
	map_fctQ_decimal38Q := &types.ParameterizedMapType{Key: fct_Q, Value: decimal_38_Q}
	struct_fctQ_ListDecimal38Q := &types.ParameterizedStructType{Types: []types.FuncDefArgType{fct_Q, list_decimal_38_Q}}
	for _, td := range []struct {
		name                  string
		params                []types.FuncDefArgType
		expectedHasSyncParams bool
	}{
		{"No Abstract Type", []types.FuncDefArgType{&types.Int64Type{}}, false},
		{"No Sync Param P, Q", []types.FuncDefArgType{fct_P, fct_Q}, false},
		{"Sync Params P, P", []types.FuncDefArgType{fct_P, fct_P}, true},
		{"Sync Params P, <P, Q>", []types.FuncDefArgType{fct_P, decimal_PQ}, true},
		{"No Sync Params P, <38, Q>", []types.FuncDefArgType{fct_P, decimal_38_Q}, false},
		{"Sync Params P, List<Decimal<P, Q>>", []types.FuncDefArgType{fct_P, list_decimal_38_Q}, false},
		{"No Sync Params fct<P>, Map<fct<Q>, decimal<38,Q>>", []types.FuncDefArgType{fct_P, map_fctQ_decimal38Q}, false},
		{"Sync Params fct<Q>, Map<fct<Q>, decimal<38,Q>>", []types.FuncDefArgType{fct_Q, map_fctQ_decimal38Q}, true},
		{"No Sync Params fct<P>, struct<fct<Q>, list<38,Q>>", []types.FuncDefArgType{fct_P, struct_fctQ_ListDecimal38Q}, false},
		{"Sync Params fct<Q>, struct<fct<Q>, list<38,Q>>", []types.FuncDefArgType{fct_Q, struct_fctQ_ListDecimal38Q}, true},
	} {
		t.Run(td.name, func(t *testing.T) {
			if td.expectedHasSyncParams {
				require.True(t, extensions.HasSyncParams(td.params))
			} else {
				require.False(t, extensions.HasSyncParams(td.params))
			}
		})
	}
}

func TestMatchWithSyncParams(t *testing.T) {
	testFileInfos := []struct {
		path     string
		funcType parser2.TestFuncType
		numTests int
	}{
		{"tests/cases/arithmetic_decimal/bitwise_or.test", parser2.ScalarFuncType, 14},
		{"tests/cases/arithmetic_decimal/bitwise_xor.test", parser2.ScalarFuncType, 14},
		{"tests/cases/arithmetic_decimal/bitwise_and.test", parser2.ScalarFuncType, 14},
		{"tests/cases/arithmetic_decimal/sqrt_decimal.test", parser2.ScalarFuncType, 14},
		{"tests/cases/arithmetic_decimal/sum_decimal.test", parser2.ScalarFuncType, 8},
	}
	for _, testFileInfo := range testFileInfos {
		fs := substrait.GetSubstraitTestsFS()
		testFile, err := parser2.ParseTestCaseFileFromFS(fs, testFileInfo.path)
		require.NoError(t, err)
		require.NotNil(t, testFile)
		assert.Len(t, testFile.TestCases, testFileInfo.numTests)

		reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(&extensions.DefaultCollection)
		for _, tc := range testFile.TestCases {
			t.Run(tc.FuncName, func(t *testing.T) {
				testGetFunctionInvocation(t, tc, &reg, funcRegistry)
			})
		}
	}
}

func TestLoadAllSubstraitTestFiles(t *testing.T) {
	got := substrait.GetSubstraitTestsFS()
	filePaths, err := listFiles(got, ".")
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(filePaths), 107)

	for _, filePath := range filePaths {
		t.Run(filePath, func(t *testing.T) {
			switch filePath {
			case "tests/cases/boolean/bool_and.test":
				t.Skip("Skipping bool_and.test")
			case "tests/cases/datetime/extract.test":
				// TODO deal with enum arguments in testcase
				t.Skip("Skipping extract.test")
			}

			testFile, err := parser2.ParseTestCaseFileFromFS(got, filePath)
			require.NoError(t, err)
			require.NotNil(t, testFile)
			reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(&extensions.DefaultCollection)
			for _, tc := range testFile.TestCases {
				testGetFunctionInvocation(t, tc, &reg, funcRegistry)
			}
		})
	}
}

func testGetFunctionInvocation(t *testing.T, tc *parser2.TestCase, reg *expr.ExtensionRegistry, registry functions.FunctionRegistry) {
	switch tc.FuncType {
	case parser2.ScalarFuncType:
		invocation, err := tc.GetScalarFunctionInvocation(reg, registry)
		require.NoError(t, err, "GetScalarFunctionInvocation failed with error in test case: %s", tc.CompoundFunctionName())
		require.Equal(t, tc.ID().URI, invocation.ID().URI)
	case parser2.AggregateFuncType:
		invocation, err := tc.GetAggregateFunctionInvocation(reg, registry)
		require.NoError(t, err)
		require.Equal(t, tc.ID().URI, invocation.ID().URI)
	}
}

func listFiles(embedFs embed.FS, root string) ([]string, error) {
	var files []string
	err := fs.WalkDir(embedFs, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
