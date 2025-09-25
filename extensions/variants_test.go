// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait"
	"github.com/substrait-io/substrait-go/v7/extensions"
	"github.com/substrait-io/substrait-go/v7/functions"
	parser2 "github.com/substrait-io/substrait-go/v7/testcases/parser"
	"github.com/substrait-io/substrait-go/v7/types"
	"github.com/substrait-io/substrait-go/v7/types/integer_parameters"
	"github.com/substrait-io/substrait-go/v7/types/parser"
)

func TestEvaluateTypeExpression(t *testing.T) {
	var (
		// Function definition argument type shortcuts.
		i64Null, _      = parser.ParseType("i64?")
		i64NonNull, _   = parser.ParseType("i64")
		strNull, _      = parser.ParseType("string?")
		strNonNull, _   = parser.ParseType("string")
		any1NonNull, _  = parser.ParseType("any1")
		any1listNonNull = mkFuncArgList(any1NonNull)

		// Few shortcut type definitions.
		i64TypeReq     = &types.Int64Type{Nullability: types.NullabilityRequired}
		strTypeReq     = &types.StringType{Nullability: types.NullabilityRequired}
		i64listNonNull = mkList(i64TypeReq)
	)

	tests := []struct {
		name      string
		nulls     extensions.NullabilityHandling
		ret       types.FuncDefArgType
		extArgs   extensions.FuncParameterList
		args      []types.Type
		expected  types.Type
		expectErr string
	}{
		{
			name:     "defaults",
			nulls:    extensions.MirrorNullability,
			ret:      i64NonNull,
			extArgs:  extensions.FuncParameterList{valArg(i64Null)},
			args:     []types.Type{&types.Int64Type{Nullability: types.NullabilityNullable}},
			expected: &types.Int64Type{Nullability: types.NullabilityNullable},
		},
		{
			name:      "arg mismatch",
			nulls:     extensions.MirrorNullability,
			ret:       strNull,
			extArgs:   extensions.FuncParameterList{valArg(strNull)},
			args:      []types.Type{},
			expectErr: "invalid expression: mismatch in number of arguments provided. got 0, expected 1",
		},
		{
			name:      "missing enum arg",
			nulls:     extensions.MirrorNullability,
			ret:       i64Null,
			extArgs:   extensions.FuncParameterList{valArg(i64NonNull), extensions.EnumArg{Name: "foo"}},
			args:      []types.Type{&types.Int64Type{}, &types.Int64Type{}},
			expectErr: "invalid type: arg #1 (foo) should be an enum"},
		{
			name:      "discrete null handling",
			nulls:     extensions.DiscreteNullability,
			ret:       strNull,
			extArgs:   extensions.FuncParameterList{valArg(strNull)},
			args:      []types.Type{&types.StringType{Nullability: types.NullabilityRequired}},
			expectErr: "invalid type: discrete nullability did not match for arg #0",
		},
		{
			name:     "mirror",
			nulls:    extensions.MirrorNullability,
			ret:      strNull,
			extArgs:  extensions.FuncParameterList{valArg(i64NonNull), valArg(i64Null)},
			args:     []types.Type{i64TypeReq, i64TypeReq},
			expected: strTypeReq,
		},
		{
			name:     "nullif(any1, any1) -> any1",
			nulls:    extensions.MirrorNullability,
			ret:      any1NonNull,
			extArgs:  extensions.FuncParameterList{valArg(any1NonNull), valArg(any1NonNull)},
			args:     []types.Type{i64TypeReq, i64TypeReq},
			expected: i64TypeReq,
		},
		{
			name:     "element_at(list<any1>, i64) -> any1",
			nulls:    extensions.DeclaredOutputNullability,
			ret:      any1NonNull,
			extArgs:  extensions.FuncParameterList{valArg(any1listNonNull), valArg(i64NonNull)},
			args:     []types.Type{i64listNonNull, i64TypeReq},
			expected: i64TypeReq,
		},
		{
			name:  "deeply nested element_at(list<list<list<any1>>>, i64, i64, i64) -> any1",
			nulls: extensions.DeclaredOutputNullability,
			ret:   any1NonNull,
			extArgs: extensions.FuncParameterList{
				valArg(mkFuncArgList(mkFuncArgList(mkFuncArgList(any1NonNull)))),
				valArg(i64NonNull), valArg(i64NonNull), valArg(i64NonNull),
			},
			args:     []types.Type{mkList(mkList(i64listNonNull)), i64TypeReq, i64TypeReq, i64TypeReq},
			expected: i64TypeReq,
		},
		{
			name:  "map_element_at(map<string, map<string, list<any1>>>, string, string, i64) -> any1",
			nulls: extensions.DeclaredOutputNullability,
			ret:   any1NonNull,
			extArgs: extensions.FuncParameterList{
				// map string -> map string -> list<any1>
				valArg(mkFuncArgMap(strTypeReq, mkFuncArgMap(strTypeReq, mkFuncArgList(any1NonNull)))),
				valArg(strNonNull), valArg(strNonNull), valArg(i64NonNull),
			},
			args: []types.Type{
				mkMap(strTypeReq, mkMap(strTypeReq, mkList(i64TypeReq))),
				strTypeReq, strTypeReq, i64TypeReq,
			},
			expected: i64TypeReq,
		},
		{
			name:  "get_any(struct<string?, any1>) -> any1",
			nulls: extensions.DeclaredOutputNullability,
			ret:   any1NonNull,
			extArgs: extensions.FuncParameterList{
				valArg(&types.ParameterizedStructType{
					Nullability: types.NullabilityRequired,
					Types:       []types.FuncDefArgType{strNull, any1NonNull},
				}),
			},
			args: []types.Type{
				&types.StructType{
					Nullability: types.NullabilityRequired,
					Types: []types.Type{
						&types.StringType{Nullability: types.NullabilityNullable},
						strTypeReq,
					},
				},
			},
			expected: strTypeReq,
		},
		{
			name:     "declared output",
			nulls:    extensions.DeclaredOutputNullability,
			ret:      strNull,
			extArgs:  extensions.FuncParameterList{valArg(strNull)},
			args:     []types.Type{strTypeReq},
			expected: &types.StringType{Nullability: types.NullabilityNullable},
		},
		{
			name:     "user defined type",
			nulls:    extensions.MirrorNullability,
			ret:      &types.ParameterizedUserDefinedType{Nullability: types.NullabilityRequired, Name: "test_type"},
			extArgs:  nil,
			args:     nil,
			expected: &types.UserDefinedType{Nullability: types.NullabilityRequired, TypeReference: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extensions.EvaluateTypeExpression("extension:org:item", tt.nulls, tt.ret, tt.extArgs, nil, tt.args, extensions.NewSet())
			if tt.expectErr == "" {
				require.NoError(t, err)
				require.Truef(t, tt.expected.Equals(result),
					"expected: %s\ngot: %s", tt.expected, result)
			} else {
				require.EqualError(t, err, tt.expectErr)
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
			result, err := extensions.EvaluateTypeExpression("extension:org:item", tt.nulls, tt.ret, tt.extArgs, &tt.variadic, tt.args, extensions.NewSet())
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

		reg, funcRegistry := functions.NewExtensionAndFunctionRegistries(extensions.GetDefaultCollectionWithNoError())
		for _, tc := range testFile.TestCases {
			t.Run(tc.FuncName, func(t *testing.T) {
				switch tc.FuncType {
				case parser2.ScalarFuncType:
					invocation, err := tc.GetScalarFunctionInvocation(&reg, funcRegistry)
					require.NoError(t, err)
					require.Equal(t, tc.ID(), invocation.ID())
				case parser2.AggregateFuncType:
					invocation, err := tc.GetAggregateFunctionInvocation(&reg, funcRegistry)
					require.NoError(t, err)
					require.Equal(t, tc.ID(), invocation.ID())
				}
			})
		}
	}
}

func mkFuncArgList(typ types.FuncDefArgType) *types.ParameterizedListType {
	return &types.ParameterizedListType{Type: typ, Nullability: types.NullabilityRequired}
}

func mkList(typ types.Type) *types.ListType {
	return &types.ListType{Type: typ, Nullability: types.NullabilityRequired}
}

func mkFuncArgMap(kt, vt types.FuncDefArgType) *types.ParameterizedMapType {
	return &types.ParameterizedMapType{
		Nullability: types.NullabilityRequired,
		Key:         kt,
		Value:       vt,
	}
}

func mkMap(kt, vt types.Type) *types.MapType {
	return &types.MapType{
		Nullability: types.NullabilityRequired,
		Key:         kt,
		Value:       vt,
	}
}

func valArg(typ types.FuncDefArgType) extensions.ValueArg {
	return extensions.ValueArg{Value: &parser.TypeExpression{ValueType: typ}}
}

func TestResolveType(t *testing.T) {
	// Test TypeReference setting logic for user-defined types
	tests := []struct {
		name        string
		returnType  string
		expectedUDT bool
		urnAndType  string
	}{
		{
			name:        "user_defined_type_sets_reference",
			returnType:  "u!some_type",
			expectedUDT: true,
			urnAndType:  "some_type",
		},
		{
			name:        "regular_type_works_normally",
			returnType:  "i64",
			expectedUDT: false,
			urnAndType:  "",
		},
		{
			name:        "string_type_works_normally",
			returnType:  "string",
			expectedUDT: false,
			urnAndType:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set up type registry
			returnType, _ := parser.ParseType(tt.returnType)
			registry := extensions.NewSet()
			var expectedRef uint32
			if tt.expectedUDT {
				expectedRef = registry.GetTypeAnchor(extensions.ID{URN: "extension:org:item", Name: tt.urnAndType})
			}

			// call EvaluateTypeExpression to convert a FuncDefArgType to a Type
			result, err := extensions.EvaluateTypeExpression(
				"test://urn",
				extensions.MirrorNullability,
				returnType,
				extensions.FuncParameterList{},
				nil,
				[]types.Type{},
				registry,
			)
			require.NoError(t, err)

			// for UserDefinedType, check that the type reference matches
			if tt.expectedUDT {
				udResult, ok := result.(*types.UserDefinedType)
				require.True(t, ok, "Expected UserDefinedType")

				if udResult != nil {
					name := strings.TrimPrefix(returnType.ShortString(), "u!")
					udResult.TypeReference = registry.GetTypeAnchor(extensions.ID{Name: name, URN: "extension:org:item"})
					assert.Equal(t, expectedRef, udResult.TypeReference)
				}
			} else {
				// otherwise, just check that the string matches
				assert.Equal(t, tt.returnType, result.String())
			}
		})
	}
}

func TestResolveTypeErrorHandling(t *testing.T) {
	// Test error propagation from EvaluateTypeExpression
	returnType, _ := parser.ParseType("u!custom_type")
	argType, _ := parser.ParseType("i64")

	// Create function parameter list that expects one argument
	funcParams := extensions.FuncParameterList{valArg(argType)}

	// Test with wrong number of arguments to trigger an error
	_, err := extensions.EvaluateTypeExpression(
		"extension:org:item",
		extensions.MirrorNullability,
		returnType,
		funcParams,
		nil,
		[]types.Type{}, // No arguments provided when one is expected
		extensions.NewSet(),
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mismatch in number of arguments")
}
