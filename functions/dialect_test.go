// SPDX-License-Identifier: Apache-2.0

package functions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v7/extensions"
	. "github.com/substrait-io/substrait-go/v7/functions"
	"github.com/substrait-io/substrait-go/v7/types"
)

var gFunctionRegistry FunctionRegistry

func TestMain(m *testing.M) {
	gFunctionRegistry = NewFunctionRegistry(extensions.GetDefaultCollectionWithNoError())
	m.Run()
}

func TestDialectApis(t *testing.T) {
	dialectYaml := `
name: testdb
type: sql
dependencies:
  arithmetic:
    extension:io.substrait:functions_arithmetic
  comparison:
    extension:io.substrait:functions_comparison
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
  i64:
    sql_type_name: TINYINT,
    supported_as_column: true
scalar_functions:
- name: comparison.is_null
  local_name: IS NULL
  postfix: true
  supported_kernels:
  - any1
aggregate_functions:
- name: arithmetic.max
  aggregate: true
  supported_kernels:
  - i32
  - i64
- name: arithmetic.min
  aggregate: true
  supported_kernels:
  - i32
  - i64
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
- name: arithmetic.rank
  supported_kernels:
  - ""
`
	dialect, err := LoadDialect(t.Name(), strings.NewReader(dialectYaml))
	assert.Nil(t, err)
	assert.NotNil(t, dialect)
	assert.Equal(t, t.Name(), dialect.Name())
	localRegistry, err := dialect.LocalizeFunctionRegistry(gFunctionRegistry)
	assert.NoError(t, err)
	assert.Equal(t, t.Name(), localRegistry.GetDialect().Name())
	assert.Equal(t, gFunctionRegistry, localRegistry.GetFunctionRegistry())
}

func TestBadDialects(t *testing.T) {
	type testcase struct {
		error       string
		dialectYaml string
	}
	tests := []testcase{
		{`no supported types`, `name: testdb`},
		{`no supported types`, `name: testdb
type: sql
dependencies:
supported_types:
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
`},
		{`no functions`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
   sql_type_name: TINYINT,
   supported_as_column: true
`,
		},
		{`unknown dependency`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
`},
		{`invalid function name`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
aggregate_functions:
- name: max
  supported_kernels:
  - i32
`},
		{`invalid function`, `name: testdb
type: sql
dependencies:
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  supported_kernels:
`},
		{`unsupported type`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
window_functions:
- name: arithmetic.unknown_function
  supported_kernels:
  - i99
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.error, func(t *testing.T) {
			_, err := LoadDialect(t.Name(), strings.NewReader(tt.dialectYaml))
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}

	localizeTestcases := []testcase{
		{`no scalar function variant found`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
  fp32:
    sql_type_name: FLOAT,
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  supported_kernels:
    - i32_i32
    - i32_fp32
`},
		{`no aggregate function variant found`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
aggregate_functions:
- name: arithmetic.max
  supported_kernels:
    - i32
    - str
`},
		{`no window function variant found`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
window_functions:
- name: arithmetic.ntile
  supported_kernels:
    - str
    - i64
`},
	}
	for _, tt := range localizeTestcases {
		t.Run(tt.error, func(t *testing.T) {
			dialect, err := LoadDialect(t.Name(), strings.NewReader(tt.dialectYaml))
			assert.NoError(t, err)

			_, err = dialect.LocalizeFunctionRegistry(gFunctionRegistry)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}

	badTypeTestcases := []testcase{
		{`unknown type`, `name: testdb
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i32:
    sql_type_name: TINYINT,
    supported_as_column: true
  myth:
    sql_type_name: TINYINT,
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  supported_kernels:
  - i32_i32
`,
		},
	}
	for _, tt := range badTypeTestcases {
		t.Run(tt.error, func(t *testing.T) {
			dialect, err := LoadDialect(t.Name(), strings.NewReader(tt.dialectYaml))
			assert.NoError(t, err)

			typeRegistry := NewTypeRegistry()
			_, err = dialect.LocalizeTypeRegistry(typeRegistry)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}
}

func getLocalFunctionRegistry(t *testing.T, dialectYaml string, substraitFuncRegistry FunctionRegistry) LocalFunctionRegistry {
	dialect, err := LoadDialect(t.Name(), strings.NewReader(dialectYaml))
	require.NoError(t, err)
	localRegistry, err := dialect.LocalizeFunctionRegistry(substraitFuncRegistry)
	require.NoError(t, err)
	return localRegistry
}

func TestScalarFunctionsLookup(t *testing.T) {
	baseUrn := "extension:io.substrait:"
	arithmeticUrn := baseUrn + "functions_arithmetic"
	booleanUrn := baseUrn + "functions_boolean"
	comparisonUrn := baseUrn + "functions_comparison"
	stringUrn := baseUrn + "functions_string"
	allFunctions := gFunctionRegistry.GetAllFunctions()

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
  boolean: 
    extension:io.substrait:functions_boolean
  comparison: 
    extension:io.substrait:functions_comparison
  string: 
    extension:io.substrait:functions_string
supported_types:
  i32:
    sql_type_name: INTEGER
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
  fp32:
    sql_type_name: REAL
    supported_as_column: true
  fp64:
    sql_type_name: DOUBLE
    supported_as_column: true
  bool:
    sql_type_name: BOOLEAN
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  local_name: +
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - i32_i32
  - i64_i64
  - fp32_fp32
  - fp64_fp64
- name: arithmetic.subtract
  local_name: '-'
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - fp32_fp32
  - fp64_fp64
- name: boolean.and
  local_name: and
  infix: true
  supported_kernels:
  - bool
- name: comparison.is_null
  local_name: IS NULL
  postfix: true
  supported_kernels:
  - any1
- name: string.concat
  local_name: '||'
  required_options:
    null_handling: ACCEPT_NULLS
  infix: true
  supported_kernels:
  - vchar
  - str
  variadic: 1
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, gFunctionRegistry)
	tests := []struct {
		numArgs          int
		localName        string
		substraitName    string
		expectedUrn      string
		expectedNames    []string
		expectedNotation FunctionNotation
		isOverflowError  bool
	}{
		{2, "+", "add", arithmeticUrn, []string{"add:i32_i32", "add:i64_i64", "add:fp32_fp32", "add:fp64_fp64"}, INFIX, true},
		{2, "-", "subtract", arithmeticUrn, []string{"subtract:fp32_fp32", "subtract:fp64_fp64"}, INFIX, true},
		{1, "IS NULL", "is_null", comparisonUrn, []string{"is_null:any"}, POSTFIX, false},
		{3, "and", "and", booleanUrn, []string{"and:bool"}, INFIX, false},
		{2, "and", "and", booleanUrn, []string{"and:bool"}, INFIX, false},
		{1, "and", "and", booleanUrn, []string{"and:bool"}, INFIX, false},
		{0, "and", "and", booleanUrn, []string{"and:bool"}, INFIX, false},
		{2, "||", "concat", stringUrn, []string{"concat:vchar", "concat:str"}, INFIX, false},
		{3, "||", "concat", stringUrn, []string{"concat:vchar", "concat:str"}, INFIX, false},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			var fv []*LocalScalarFunctionVariant
			fv = localRegistry.GetScalarFunctions(LocalFunctionName(tt.localName), tt.numArgs)

			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.expectedNotation, fv[0].Notation())
			assert.Equal(t, tt.isOverflowError, fv[0].IsOptionSupported("overflow", "ERROR"))
			assert.False(t, fv[0].IsOptionSupported("overflow", "SILENT"))
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			fv = localRegistry.GetScalarFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			scalarFunctions := gFunctionRegistry.GetScalarFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
			scalarFunctions = gFunctionRegistry.GetScalarFunctionsByName(tt.substraitName)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestAggregateFunctionsLookup(t *testing.T) {
	expectedUrn := "extension:io.substrait:functions_arithmetic"
	allFunctions := gFunctionRegistry.GetAllFunctions()
	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i8:
    sql_type_name: INTEGER
    supported_as_column: true
  i16:
    sql_type_name: BIGINT
    supported_as_column: true
  fp32:
    sql_type_name: REAL
    supported_as_column: true
  fp64:
    sql_type_name: DOUBLE
    supported_as_column: true
aggregate_functions:
- name: arithmetic.min
  aggregate: true
  supported_kernels:
  - i8
  - i16
  - fp32
  - fp64
- name: arithmetic.max
  aggregate: true
  supported_kernels:
  - i8
  - i16
  - fp32
  - fp64
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, gFunctionRegistry)

	tests := []struct {
		numArgs          int
		localName        string
		substraitName    string
		expectedUrn      string
		expectedNames    []string
		expectedNotation FunctionNotation
	}{
		{1, "max", "max", expectedUrn, []string{"max:i8", "max:i16", "max:fp32", "max:fp64"}, PREFIX},
		{1, "min", "min", expectedUrn, []string{"min:i8", "min:i16", "min:fp32", "min:fp64"}, PREFIX},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			av := localRegistry.GetAggregateFunctions(LocalFunctionName(tt.localName), 1)

			assert.Greater(t, len(av), 0)
			assert.Equal(t, tt.expectedUrn, av[0].URN())
			assert.Equal(t, tt.localName, av[0].LocalName())
			assert.Equal(t, tt.expectedNotation, av[0].Notation())
			assert.False(t, av[0].IsOptionSupported("overflow", "ERROR"))
			checkCompoundNames(t, getAggregateCompoundNames(av), tt.expectedNames)

			av = localRegistry.GetAggregateFunctions(SubstraitFunctionName(tt.substraitName), 1)
			assert.Greater(t, len(av), 0)
			assert.Equal(t, tt.expectedUrn, av[0].URN())
			assert.Equal(t, tt.substraitName, av[0].LocalName())
			checkCompoundNames(t, getAggregateCompoundNames(av), tt.expectedNames)

			winFunctions := gFunctionRegistry.GetAggregateFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
			winFunctions = gFunctionRegistry.GetAggregateFunctionsByName(tt.substraitName)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestWindowFunctionsLookup(t *testing.T) {
	expectedUrn := "extension:io.substrait:functions_arithmetic"
	allFunctions := gFunctionRegistry.GetAllFunctions()
	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic
supported_types:
  i32:
    sql_type_name: INTEGER
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
window_functions:
- name: arithmetic.ntile
  supported_kernels:
  - i32
  - i64
- name: arithmetic.rank
  supported_kernels:
  - ""
`
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, gFunctionRegistry)
	tests := []struct {
		numArgs          int
		localName        string
		substraitName    string
		expectedUrn      string
		expectedNames    []string
		expectedNotation FunctionNotation
	}{
		{1, "ntile", "ntile", expectedUrn, []string{"ntile:i32", "ntile:i64"}, PREFIX},
		{0, "rank", "rank", expectedUrn, []string{"rank:", "rank:"}, PREFIX},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			wf := localRegistry.GetWindowFunctions(LocalFunctionName(tt.localName), tt.numArgs)
			assert.Greater(t, len(wf), 0)
			assert.Equal(t, tt.expectedUrn, wf[0].URN())
			assert.Equal(t, tt.localName, wf[0].LocalName())
			assert.Equal(t, tt.expectedNotation, wf[0].Notation())
			assert.False(t, wf[0].IsOptionSupported("overflow", "ERROR"))
			checkCompoundNames(t, getWindowCompoundNames(wf), tt.expectedNames)

			wf = localRegistry.GetWindowFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			assert.Greater(t, len(wf), 0)
			assert.Equal(t, tt.expectedUrn, wf[0].URN())
			assert.Equal(t, tt.substraitName, wf[0].LocalName())
			checkCompoundNames(t, getWindowCompoundNames(wf), tt.expectedNames)

			winFunctions := gFunctionRegistry.GetWindowFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
			winFunctions = gFunctionRegistry.GetWindowFunctionsByName(tt.substraitName)
			assert.Greater(t, len(winFunctions), 0)
			for _, f := range winFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func getScalarCompoundNames(fv []*LocalScalarFunctionVariant) []string {
	var compoundNames []string
	for _, f := range fv {
		compoundNames = append(compoundNames, f.CompoundName())
	}
	return compoundNames
}

func getAggregateCompoundNames(av []*LocalAggregateFunctionVariant) []string {
	var compoundNames []string
	for _, f := range av {
		compoundNames = append(compoundNames, f.CompoundName())
	}
	return compoundNames
}

func getWindowCompoundNames(wv []*LocalWindowFunctionVariant) []string {
	var compoundNames []string
	for _, f := range wv {
		compoundNames = append(compoundNames, f.CompoundName())
	}
	return compoundNames
}

func checkCompoundNames(t *testing.T, compoundNames []string, expectedNames []string) {
	for _, name := range expectedNames {
		assert.Contains(t, compoundNames, name)
	}
}

// test match functionality fails if it has sync param
func TestScalarFunctionsSyncParamsError(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_testsync"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: decimal<P,S>
          - name: y
            value: decimal<P,S>
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:test:def
supported_types:
  dec:
    sql_type_name: NUMBER
scalar_functions:
- name: arithmetic.func_testsync
  supported_kernels:
  - dec_dec
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	argTypes := []types.Type{
		&types.DecimalType{Nullability: types.NullabilityNullable, Precision: 10, Scale: 2},
		&types.DecimalType{Nullability: types.NullabilityNullable, Precision: 10, Scale: 2},
	}
	argTypesWithMismatchedParams := []types.Type{
		&types.DecimalType{Nullability: types.NullabilityNullable, Precision: 20, Scale: 2},
		&types.DecimalType{Nullability: types.NullabilityNullable, Precision: 10, Scale: 2},
	}

	fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_testsync"), 2)

	require.Len(t, fv, 1)
	isMatch, err := fv[0].Match(argTypes)
	require.NoError(t, err)
	require.True(t, isMatch)
	isMatch, err = fv[0].Match(argTypesWithMismatchedParams)
	require.NoError(t, err)
	require.False(t, isMatch)

	// test MatchAt
	for pos, typ := range argTypes {
		_, err = fv[0].MatchAt(typ, pos)
		require.NoError(t, err)
	}
}

// test match functionality with MIRROR nullability
func TestScalarFunctionsMirrorNullabilityMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_mirror"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_mirror
  supported_kernels:
  - i32_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}
	int32Required := &types.Int32Type{Nullability: types.NullabilityRequired}

	tests := []struct {
		name     string
		argTypes []types.Type
	}{
		{"All Arguments Nullable", []types.Type{int32Nullable, int32Nullable}},
		{"All Arguments Required", []types.Type{int32Required, int32Required}},
		{"Arguments Nullable Required Mix", []types.Type{int32Nullable, int32Required}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_mirror"), 2)

			require.Len(t, fv, 1)
			match, err := fv[0].Match(tt.argTypes)
			require.NoError(t, err)
			require.True(t, match)
			// non-variadic function, min/max argument count should be 2
			require.Equal(t, 2, fv[0].MinArgumentCount())
			require.Equal(t, 2, fv[0].MaxArgumentCount())

			// test MatchAt
			for pos, typ := range tt.argTypes {
				match, err = fv[0].MatchAt(typ, pos)
				require.NoError(t, err)
				require.True(t, match)
			}

		})
	}
}

// test match functionality DeclaredOutput nullability
func TestScalarFunctionsDeclaredOutputNullabilityMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_declared_output"
    description: "Subtract two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
        nullability: DECLARED_OUTPUT
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_declared_output
  supported_kernels:
  - i32_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}
	int32Required := &types.Int32Type{Nullability: types.NullabilityRequired}

	tests := []struct {
		name     string
		argTypes []types.Type
	}{
		{"All Arguments Nullable", []types.Type{int32Nullable, int32Nullable}},
		{"All Arguments Required", []types.Type{int32Required, int32Required}},
		{"Arguments Nullable Required Mix", []types.Type{int32Nullable, int32Required}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_declared_output"), 2)

			require.Len(t, fv, 1)
			match, err := fv[0].Match(tt.argTypes)
			require.NoError(t, err)
			require.True(t, match)

			for pos, typ := range tt.argTypes {
				match, err = fv[0].MatchAt(typ, pos)
				require.NoError(t, err)
				require.True(t, match)
			}

		})
	}
}

// test match functionality with DISCRETE nullability
func TestScalarFunctionsDiscreteNullabilityMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_discrete_required"
    description: "multiply two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
        nullability: DISCRETE
  -
    name: "func_discrete_nullable"
    description: "multiply two values."
    impls:
      - args:
          - name: x
            value: i32?
          - name: y
            value: i32?
        return: i32
        nullability: DISCRETE
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_discrete_required
  supported_kernels:
  - i32_i32
- name: arithmetic.func_discrete_nullable
  supported_kernels:
  - i32_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}
	int32Required := &types.Int32Type{Nullability: types.NullabilityRequired}

	tests := []struct {
		name        string
		localName   string
		argTypes    []types.Type
		shouldMatch bool
	}{
		{"param nullable, arg nullable, should match", "func_discrete_nullable", []types.Type{int32Nullable, int32Nullable}, true},
		{"param required, arg required, should match", "func_discrete_required", []types.Type{int32Required, int32Required}, true},
		{"param nullable, arg required, shouldn't match", "func_discrete_nullable", []types.Type{int32Required, int32Required}, false},
		{"param required, arg nullable, should match", "func_discrete_required", []types.Type{int32Nullable, int32Nullable}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := localRegistry.GetScalarFunctions(LocalFunctionName(tt.localName), 2)

			require.Len(t, fv, 1)
			match, err := fv[0].Match(tt.argTypes)
			require.NoError(t, err)
			require.Equal(t, tt.shouldMatch, match)

			for pos, typ := range tt.argTypes {
				match, err = fv[0].MatchAt(typ, pos)
				require.NoError(t, err)
				require.Equal(t, tt.shouldMatch, match)
			}

		})
	}
}

// test match functionality returns true for function with variadic argument
func TestScalarFunctionsVariadicMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_testvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        variadic:
          min: 1
          max: 2
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_testvariadic
  supported_kernels:
  - i32_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}

	fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_testvariadic"), 2)

	argTypes := []types.Type{int32Nullable, int32Nullable}
	require.Len(t, fv, 1)
	match, err := fv[0].Match(argTypes)
	require.NoError(t, err)
	assert.True(t, match)

	// test MatchAt
	for pos, typ := range argTypes {
		match, err = fv[0].MatchAt(typ, pos)
		require.NoError(t, err)
		assert.True(t, match)
	}
}

func TestScalarFunctionsVariadicMin0(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: and
    impls:
      - args:
          - value: boolean?
            name: a
        variadic:
          min: 0
        return: boolean?
`

	dialectYaml := `
name: test
type: sql
dependencies:
  boolean:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
scalar_functions:
  - name: boolean.and
    local_name: and
    supported_kernels:
      - bool
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	booleanNullable := &types.BooleanType{Nullability: types.NullabilityNullable}
	int8Nullable := &types.Int8Type{Nullability: types.NullabilityNullable}

	tests := []struct {
		name     string
		argTypes []types.Type
		isMatch  bool
	}{
		{"No Arguments", []types.Type{}, true},
		{"One Argument", []types.Type{booleanNullable}, true},
		{"Two Arguments", []types.Type{booleanNullable, booleanNullable}, true},
		{"Three Arguments", []types.Type{booleanNullable, booleanNullable, booleanNullable}, true},
		{"Wrong Arguments", []types.Type{int8Nullable}, false},
		{"Wrong Arguments", []types.Type{booleanNullable, int8Nullable}, false},
		{"Wrong Arguments", []types.Type{int8Nullable, booleanNullable}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := localRegistry.GetScalarFunctions(LocalFunctionName("and"), len(tt.argTypes))

			require.Len(t, fv, 1)
			match, err := fv[0].Match(tt.argTypes)
			require.NoError(t, err)
			assert.Equal(t, tt.isMatch, match)
			if !tt.isMatch {
				return
			}

			for pos, typ := range tt.argTypes {
				match, err = fv[0].MatchAt(typ, pos)
				require.NoError(t, err)
				assert.Equal(t, tt.isMatch, match)
			}
		})
	}
}

// this tests that match functionality returns true for function with variadic argument
// when argument count is greater than variadic min count
func TestScalarFuncVariadicArgMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_testvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i64
          - name: y
            value: i32
        variadic:
          min: 1
          max: 3
          parameterConsistency: INCONSISTENT
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
  i64:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_testvariadic
  supported_kernels:
  - i64_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}
	int64Nullable := &types.Int64Type{Nullability: types.NullabilityNullable}

	fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_testvariadic"), 2)

	// pass third argument as variadic, it should match against last argument type
	argTypes := []types.Type{int64Nullable, int32Nullable, int32Nullable}
	require.Len(t, fv, 1)
	require.Equal(t, 3, fv[0].MinArgumentCount())
	require.Equal(t, 5, fv[0].MaxArgumentCount())
	match, err := fv[0].Match(argTypes)
	require.NoError(t, err)
	assert.True(t, match)

	// test MatchAt
	for pos, typ := range argTypes {
		match, err = fv[0].MatchAt(typ, pos)
		require.NoError(t, err)
		assert.True(t, match)
	}
}

// this tests that match functionality returns true for function with variadic argument
// when argument count is greater than variadic min count
func TestScalarFuncVariadicArgMisMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_testvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i64
          - name: y
            value: i32
        variadic:
          min: 1
          max: 3
          parameterConsistency: INCONSISTENT
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
  i64:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_testvariadic
  supported_kernels:
  - i64_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}
	int64Nullable := &types.Int64Type{Nullability: types.NullabilityNullable}

	fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_testvariadic"), 2)

	// pass third argument as variadic but different from 2nd argument, it should match against last argument type
	argTypes := []types.Type{int64Nullable, int32Nullable, int64Nullable}
	require.Len(t, fv, 1)
	match, err := fv[0].Match(argTypes)
	require.NoError(t, err)
	assert.False(t, match)

	// last argument shouldn't match
	match, err = fv[0].MatchAt(argTypes[2], 2)
	require.NoError(t, err)
	assert.False(t, match)
}

// test match functionality returns true for function with variadic argument
// if argument count is lesser than variadic min count
func TestScalarFuncVariadicMismatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_testvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: decimal<P1,S1>
          - name: y
            value: decimal<P2,S2>
        variadic:
          min: 3
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  dec:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_testvariadic
  supported_kernels:
  - dec_dec
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	int32Nullable := &types.Int32Type{Nullability: types.NullabilityNullable}

	fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_testvariadic"), 2)

	argTypes := []types.Type{int32Nullable, int32Nullable}
	require.Len(t, fv, 1)
	match, err := fv[0].Match(argTypes)
	require.NoError(t, err)
	assert.False(t, match)

	// test MatchAt
	for pos, typ := range argTypes {
		match, err = fv[0].MatchAt(typ, pos)
		require.NoError(t, err)
		assert.False(t, match)
	}
}

// test match functionality returns false if consistency check for argument fails
// when function implementation has "CONSISTENCY" property for parameter consistency
func TestScalarFuncVariadicConsistencyCheckMisMatch(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
scalar_functions:
  -
    name: "func_testvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: decimal<P1,S1>
          - name: y
            value: decimal<P2,S2>
        variadic:
          min: 1
          max: 2
          parameterConsistency: CONSISTENT
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  dec:
    sql_type_name: INTEGER
scalar_functions:
- name: arithmetic.func_testvariadic
  supported_kernels:
  - dec_dec
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	dec38_2 := &types.DecimalType{Precision: 38, Scale: 2, Nullability: types.NullabilityNullable}
	dec38_5 := &types.DecimalType{Precision: 38, Scale: 5, Nullability: types.NullabilityNullable}

	fv := localRegistry.GetScalarFunctions(LocalFunctionName("func_testvariadic"), 2)

	// one type is int32 and other int64, since concrete type is not consistent so match should fail
	argTypes := []types.Type{dec38_2, dec38_2, dec38_5}
	require.Len(t, fv, 1)
	match, err := fv[0].Match(argTypes)
	require.NoError(t, err)
	// match should fail because of concrete type are different
	// even though function argument allows decimal(P, S)
	assert.False(t, match)
}

func TestAggregateFuncMinMax(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
aggregate_functions:
  -
    name: "func_nonvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
  -
    name: "func_variadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        variadic:
          min: 1
          max: 3
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
aggregate_functions:
- name: arithmetic.func_nonvariadic
  supported_kernels:
  - i32_i32
- name: arithmetic.func_variadic
  supported_kernels:
  - i32_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	// test non-variadic min-max
	fv := localRegistry.GetAggregateFunctions(LocalFunctionName("func_nonvariadic"), 2)
	require.Len(t, fv, 1)
	require.Equal(t, 2, fv[0].MinArgumentCount())
	require.Equal(t, 2, fv[0].MaxArgumentCount())

	// test variadic min-max
	fv = localRegistry.GetAggregateFunctions(LocalFunctionName("func_variadic"), 2)
	require.Len(t, fv, 1)
	require.Equal(t, 3, fv[0].MinArgumentCount())
	require.Equal(t, 5, fv[0].MaxArgumentCount())
}

func TestWindowFuncMinMax(t *testing.T) {
	const urn = "extension:test:def"
	const defYaml = `---
urn: extension:test:def
window_functions:
  -
    name: "func_nonvariadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        return: i32
  -
    name: "func_variadic"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i32
          - name: y
            value: i32
        variadic:
          min: 1
          max: 3
        return: i32
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:test:def
supported_types:
  i32:
    sql_type_name: INTEGER
window_functions:
- name: arithmetic.func_nonvariadic
  supported_kernels:
  - i32_i32
- name: arithmetic.func_variadic
  supported_kernels:
  - i32_i32
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(urn, strings.NewReader(defYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)

	// test non-variadic min-max
	fv := localRegistry.GetWindowFunctions(LocalFunctionName("func_nonvariadic"), 2)
	require.Len(t, fv, 1)
	require.Equal(t, 2, fv[0].MinArgumentCount())
	require.Equal(t, 2, fv[0].MaxArgumentCount())

	// test variadic min-max
	fv = localRegistry.GetWindowFunctions(LocalFunctionName("func_variadic"), 2)
	require.Len(t, fv, 1)
	require.Equal(t, 3, fv[0].MinArgumentCount())
	require.Equal(t, 5, fv[0].MaxArgumentCount())
}

func TestDecimalScalarFunctionsLookup(t *testing.T) {
	baseUrn := "extension:io.substrait:"
	decArithmeticUrn := baseUrn + "functions_arithmetic_decimal"
	allFunctions := gFunctionRegistry.GetAllFunctions()

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:io.substrait:functions_arithmetic_decimal
supported_types:
  decimal:
    sql_type_name: INTEGER
    supported_as_column: true
scalar_functions:
- name: arithmetic.add
  local_name: +
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - dec_dec
- name: arithmetic.subtract
  local_name: '-'
  infix: true
  required_options:
    overflow: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - dec_dec
`
	decType30S2 := &types.DecimalType{Precision: 30, Scale: 2, Nullability: types.NullabilityNullable}
	decType32S2 := &types.DecimalType{Precision: 32, Scale: 2, Nullability: types.NullabilityNullable}
	decType33S2 := &types.DecimalType{Precision: 33, Scale: 2, Nullability: types.NullabilityNullable}
	decType10S4 := &types.DecimalType{Precision: 10, Scale: 4, Nullability: types.NullabilityNullable}
	decType12S5 := &types.DecimalType{Precision: 12, Scale: 5, Nullability: types.NullabilityNullable}
	decType8S3 := &types.DecimalType{Precision: 8, Scale: 3, Nullability: types.NullabilityNullable}
	decType9S2 := &types.DecimalType{Precision: 9, Scale: 2, Nullability: types.NullabilityNullable}
	decType11S3 := &types.DecimalType{Precision: 11, Scale: 3, Nullability: types.NullabilityNullable}
	decType20S10 := &types.DecimalType{Precision: 20, Scale: 10, Nullability: types.NullabilityNullable}
	decType21S10 := &types.DecimalType{Precision: 21, Scale: 10, Nullability: types.NullabilityNullable}
	decType35S30 := &types.DecimalType{Precision: 35, Scale: 30, Nullability: types.NullabilityNullable}
	decType36S30 := &types.DecimalType{Precision: 36, Scale: 30, Nullability: types.NullabilityNullable}
	decType38S20 := &types.DecimalType{Precision: 38, Scale: 20, Nullability: types.NullabilityNullable}
	decType38S19 := &types.DecimalType{Precision: 38, Scale: 19, Nullability: types.NullabilityNullable}
	decType10S5 := &types.DecimalType{Precision: 10, Scale: 5, Nullability: types.NullabilityNullable}
	dectype12S8 := &types.DecimalType{Precision: 12, Scale: 8, Nullability: types.NullabilityNullable}
	decType14S8 := &types.DecimalType{Precision: 14, Scale: 8, Nullability: types.NullabilityNullable}
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, gFunctionRegistry)
	tests := []struct {
		numArgs            int
		localName          string
		substraitName      string
		args               []types.Type
		expectedReturnType types.Type
		expectedUrn        string
		expectedNames      []string
		expectedNotation   FunctionNotation
		isOverflowError    bool
	}{
		{2, "+", "add", []types.Type{decType30S2, decType32S2}, decType33S2, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "+", "add", []types.Type{decType10S4, decType10S5}, decType12S5, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "+", "add", []types.Type{decType8S3, decType9S2}, decType11S3, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "+", "add", []types.Type{decType20S10, decType20S10}, decType21S10, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "+", "add", []types.Type{decType35S30, decType35S30}, decType36S30, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "+", "add", []types.Type{decType38S20, decType38S20}, decType38S19, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "+", "add", []types.Type{decType10S5, dectype12S8}, decType14S8, decArithmeticUrn, []string{"add:dec_dec"}, INFIX, true},
		{2, "-", "subtract", []types.Type{decType30S2, decType32S2}, decType33S2, decArithmeticUrn, []string{"subtract:dec_dec"}, INFIX, true},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			var fv []*LocalScalarFunctionVariant
			fv = localRegistry.GetScalarFunctions(LocalFunctionName(tt.localName), tt.numArgs)

			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.expectedNotation, fv[0].Notation())
			assert.Equal(t, tt.isOverflowError, fv[0].IsOptionSupported("overflow", "ERROR"))
			assert.False(t, fv[0].IsOptionSupported("overflow", "SILENT"))
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)
			retType, err := fv[0].ResolveType(tt.args, extensions.NewSet())
			require.NoError(t, err)
			assert.Equal(t, tt.expectedReturnType, retType)
			fv = localRegistry.GetScalarFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			scalarFunctions := gFunctionRegistry.GetScalarFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
			scalarFunctions = gFunctionRegistry.GetScalarFunctionsByName(tt.substraitName)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestScalarFunctionLookupWithAnyReturnType(t *testing.T) {
	baseUrn := "extension:io.substrait:"
	decArithmeticUrn := baseUrn + "functions_comparison"
	allFunctions := gFunctionRegistry.GetAllFunctions()

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic:
    extension:io.substrait:functions_arithmetic
  comparison:
    extension:io.substrait:functions_comparison
supported_types:
  decimal:
    sql_type_name: INTEGER
    supported_as_column: true
scalar_functions:
  - name: comparison.coalesce
    supported_kernels:
    - any1
    variadic: 2
  - name: comparison.least
    local_name: least
    supported_kernels:
    - any1
    variadic: 2
  - name: comparison.nullif
    supported_kernels:
    - any1_any1
`
	varcharL50 := &types.VarCharType{Length: 50, Nullability: types.NullabilityNullable}
	decType30S2 := &types.DecimalType{Precision: 30, Scale: 2, Nullability: types.NullabilityNullable}
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, gFunctionRegistry)

	tests := []struct {
		numArgs            int
		localName          string
		substraitName      string
		args               []types.Type
		expectedReturnType types.Type
		expectedUrn        string
		expectedNames      []string
	}{
		{2, "least", "least", []types.Type{decType30S2, decType30S2}, decType30S2, decArithmeticUrn, []string{"least:any"}},
		{3, "coalesce", "coalesce", []types.Type{varcharL50, varcharL50, varcharL50}, varcharL50, decArithmeticUrn, []string{"coalesce:any"}},
		{2, "nullif", "nullif", []types.Type{decType30S2, decType30S2}, decType30S2, decArithmeticUrn, []string{"nullif:any_any"}},
	}
	for _, tt := range tests {
		t.Run(tt.substraitName, func(t *testing.T) {
			var fv []*LocalScalarFunctionVariant
			fv = localRegistry.GetScalarFunctions(LocalFunctionName(tt.localName), tt.numArgs)

			require.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.False(t, fv[0].IsOptionSupported("overflow", "SILENT"))
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)
			retType, err := fv[0].ResolveType(tt.args, extensions.NewSet())
			require.NoError(t, err)
			assert.Equal(t, tt.expectedReturnType, retType)
			fv = localRegistry.GetScalarFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			assert.Greater(t, len(fv), 0)
			assert.Equal(t, tt.expectedUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getScalarCompoundNames(fv), tt.expectedNames)

			scalarFunctions := gFunctionRegistry.GetScalarFunctions(tt.substraitName, tt.numArgs)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
			scalarFunctions = gFunctionRegistry.GetScalarFunctionsByName(tt.substraitName)
			assert.Greater(t, len(scalarFunctions), 0)
			for _, f := range scalarFunctions {
				assert.Contains(t, allFunctions, f)
			}
		})
	}
}

func TestScalarFunctionsWithVariantsWithSameFuncName(t *testing.T) {
	const arithmeticUrn = "extension:test:arithmetic"
	const decimalUrn = "extension:test:decimal"
	const arithmeticYaml = `---
urn: extension:test:arithmetic
scalar_functions:
  -
    name: "sqrt"
    description: "Square root of the value"
    impls:
      - args:
          - name: x
            value: i64
        options:
          rounding:
            values: [ TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR ]
          on_domain_error:
            values: [ NAN, ERROR ]
        return: fp64
      - args:
          - name: x
            value: fp32
        options:
          rounding:
            values: [ TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR ]
          on_domain_error:
            values: [ NAN, ERROR ]
        return: fp32
      - args:
          - name: x
            value: fp64
        options:
          rounding:
            values: [ TIE_TO_EVEN, TIE_AWAY_FROM_ZERO, TRUNCATE, CEILING, FLOOR ]
          on_domain_error:
            values: [ NAN, ERROR ]
        return: fp64
`
	const decimalYaml = `---
urn: extension:test:decimal
scalar_functions:
  - name: "sqrt"
    description: Square root of the value. Sqrt of 0 is 0 and sqrt of negative values will raise an error.
    impls:
      - args:
          - name: x
            value: "DECIMAL<P,S>"
        return: fp64
`

	dialectYaml := `
name: test
type: sql
dependencies:
  arithmetic: 
    extension:test:arithmetic
  decimal_arithmetic: 
    extension:test:decimal
supported_types:
  dec:
    sql_type_name: numeric
    supported_as_column: true
  fp32:
    sql_type_name: FLOAT
    supported_as_column: true
  fp64:
    sql_type_name: DOUBLE
    supported_as_column: true
scalar_functions:
- name: arithmetic.sqrt
  local_name: sqrt
  required_options:
    on_domain_error: ERROR
    rounding: TIE_TO_EVEN
  supported_kernels:
  - fp32
  - fp64
- name: decimal_arithmetic.sqrt
  supported_kernels:
  - dec
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(arithmeticUrn, strings.NewReader(arithmeticYaml)))
	require.NoError(t, c.Load(decimalUrn, strings.NewReader(decimalYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)
	allFunctions := funcRegistry.GetAllFunctions()

	var fv []*LocalScalarFunctionVariant
	fv = localRegistry.GetScalarFunctions(LocalFunctionName("sqrt"), 1)

	expectedUrns := []string{arithmeticUrn, decimalUrn}
	expectedNames := []string{"sqrt:fp64", "sqrt:fp32", "sqrt:dec"}
	assert.Equal(t, len(fv), 3)

	urnsFound := make(map[string]bool)
	for _, f := range fv {
		assert.Equal(t, "sqrt", f.LocalName())
		assert.Equal(t, "sqrt", f.Name())
		assert.Contains(t, expectedUrns, f.URN())
		urnsFound[f.URN()] = true
	}
	checkCompoundNames(t, getScalarCompoundNames(fv), expectedNames)
	assert.Len(t, urnsFound, len(expectedUrns))
	for k := range urnsFound {
		assert.Contains(t, expectedUrns, k)
	}

	urnsFound = make(map[string]bool)
	fv = localRegistry.GetScalarFunctions(SubstraitFunctionName("sqrt"), 1)
	assert.Equal(t, len(fv), 3)
	for _, f := range fv {
		assert.Equal(t, "sqrt", f.LocalName())
		assert.Equal(t, "sqrt", f.Name())
		assert.Contains(t, expectedUrns, f.URN())
		urnsFound[f.URN()] = true
	}
	assert.Len(t, urnsFound, len(expectedUrns))
	checkCompoundNames(t, getScalarCompoundNames(fv), []string{})

	scalarFunctions := funcRegistry.GetScalarFunctions("sqrt", 1)
	assert.Greater(t, len(scalarFunctions), 0)
	for _, f := range scalarFunctions {
		assert.Contains(t, allFunctions, f)
	}
	scalarFunctions = funcRegistry.GetScalarFunctionsByName("sqrt")
	assert.Greater(t, len(scalarFunctions), 0)
	for _, f := range scalarFunctions {
		assert.Contains(t, allFunctions, f)
	}
}

func TestAggregateFunctionsWithSameName(t *testing.T) {
	const arithmeticUrn = "extension:test:arithmetic"
	const decimalUrn = "extension:test:decimal"
	const decimalYaml = `---
urn: extension:test:decimal
aggregate_functions:
  - name: "count"
    description: Count a set of values. Result is returned as a decimal instead of i64.
    impls:
      - args:
          - name: x
            value: any
        options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: decimal<38,0>
        return: decimal<38,0>
  - name: "count"
    description: "Count a set of records (not field referenced). Result is returned as a decimal instead of i64."
    impls:
      - options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: decimal<38,0>
        return: decimal<38,0>

`
	const arithmeticYaml = `---
urn: extension:test:arithmetic
aggregate_functions:
  - name: "count"
    description: Count a set of values
    impls:
      - args:
          - name: x
            value: any
        options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i64
  - name: "count"
    description: "Count a set of records (not field referenced)"
    impls:
      - options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i64

`

	dialectYaml := `
name: test
type: sql
dependencies:
  aggregate:
    extension:test:arithmetic
  aggdec:
    extension:test:decimal
supported_types:
  dec:
    sql_type_name: numeric
    supported_as_column: true
  i64:
    sql_type_name: BIGINT
    supported_as_column: true
aggregate_functions:
- name: aggdec.count
  local_name: count
  aggregate: true
  supported_kernels:
    - any
- name: aggdec.count
  local_name: count_rows
  aggregate: true
  supported_kernels:
    - ""
`
	// get substrait function registry
	var c extensions.Collection
	require.NoError(t, c.Load(arithmeticUrn, strings.NewReader(arithmeticYaml)))
	require.NoError(t, c.Load(decimalUrn, strings.NewReader(decimalYaml)))
	funcRegistry := NewFunctionRegistry(&c)
	localRegistry := getLocalFunctionRegistry(t, dialectYaml, funcRegistry)
	allFunctions := funcRegistry.GetAllFunctions()

	testcases := []struct {
		numArgs               int
		localName             string
		substraitName         string
		signature             string
		numSubstraitFunctions int
	}{
		{1, "count", "count", "count:any", 2},
		{0, "count_rows", "count", "count:", 2},
	}
	for _, tt := range testcases {
		t.Run(tt.localName, func(t *testing.T) {
			var fv []*LocalAggregateFunctionVariant
			fv = localRegistry.GetAggregateFunctions(LocalFunctionName(tt.localName), tt.numArgs)

			require.Greater(t, len(fv), 0)
			assert.Equal(t, decimalUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getAggregateCompoundNames(fv), []string{tt.signature})

			fv = localRegistry.GetAggregateFunctions(SubstraitFunctionName(tt.substraitName), tt.numArgs)
			require.Greater(t, len(fv), 0)
			assert.Equal(t, decimalUrn, fv[0].URN())
			assert.Equal(t, tt.localName, fv[0].LocalName())
			assert.Equal(t, tt.substraitName, fv[0].Name())
			checkCompoundNames(t, getAggregateCompoundNames(fv), []string{tt.signature})

			aggregateFunctions := funcRegistry.GetAggregateFunctions(tt.substraitName, tt.numArgs)
			assert.Equal(t, tt.numSubstraitFunctions, len(aggregateFunctions))
			for _, f := range aggregateFunctions {
				assert.Contains(t, allFunctions, f)
			}
			aggregateFunctions = funcRegistry.GetAggregateFunctionsByName(tt.substraitName)
			assert.Equal(t, 4, len(aggregateFunctions))
			urnMap := map[string]int{
				arithmeticUrn: 0,
				decimalUrn:    0,
			}
			for _, f := range aggregateFunctions {
				assert.Contains(t, allFunctions, f)
				urnMap[f.URN()]++
			}
			assert.Equal(t, 2, urnMap[arithmeticUrn])
			assert.Equal(t, 2, urnMap[decimalUrn])
		})
	}
}
