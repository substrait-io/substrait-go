package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v3/types"
	"github.com/substrait-io/substrait-go/v3/types/parser"
)

type testcase struct {
	name    string
	expr    string
	want    types.Type
	wantErr assert.ErrorAssertionFunc
}

func parseAndTestTypeDerivation(t *testing.T, tt *testcase) {
	resultType, err := parser.ParseType(tt.expr)
	require.NoError(t, err)
	require.NotNil(t, resultType)
	derivation, ok := resultType.(*types.OutputDerivation)
	require.True(t, ok)
	got, err := derivation.ReturnType(nil, nil)
	if !tt.wantErr(t, err, fmt.Sprintf("Evaluate(%v)", tt.expr)) {
		return
	}
	if tt.want == nil {
		require.Nil(t, got)
		return
	}
	assert.Equalf(t, tt.want.WithNullability(types.NullabilityRequired), got, "Evaluate(%v)", tt.expr)
}

func TestBinaryExpr_Evaluate(t *testing.T) {
	tests := []testcase{
		{"+", "x = 1 + 2\nvarchar<x>", &types.VarCharType{Length: 3}, assert.NoError},
		{"+", "x = 1 + 2\ndecimal<20, x>", &types.DecimalType{Precision: 20, Scale: 3}, assert.NoError},
		{"+", "P1 = 30 / 2\nS1 = 3\ndecimal<P1, S1>", &types.DecimalType{Precision: 15, Scale: 3}, assert.NoError},
		{"-", "L1 = 9\nL2 = 4\nL3 = L1 - L2\nvarchar<L3>", &types.VarCharType{Length: 5}, assert.NoError},
		{"*", "l2 = 5\nl3 = 6\nl4 = l2 * l3\nvarchar<l4>", &types.VarCharType{Length: 30}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseAndTestTypeDerivation(t, &tt)
		})
	}
}

func TestFunctionCallExpr_Evaluate(t *testing.T) {
	tests := []testcase{
		{"max", "x = max(1, 2)\nvarchar<x>", &types.VarCharType{Length: 2}, assert.NoError},
		{"min", "L1 = min(3, 4)\nvarchar<L1>", &types.VarCharType{Length: 3}, assert.NoError},
		{"abs", "l2 = abs(-5)\nvarchar<l2>", &types.VarCharType{Length: 5}, assert.NoError},
		{"abs", "l2 = abs(5)\nvarchar<l2>", &types.VarCharType{Length: 5}, assert.NoError},
		{"max", "L1 = 10\nL2 = 20\nL3 = max(L1, L2)\nvarchar<L3>", &types.VarCharType{Length: 20}, assert.NoError},
		{"unknown", "l2 = unknown(5)\nvarchar<l2>", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseAndTestTypeDerivation(t, &tt)
		})
	}
}

func TestIfExpr_Evaluate(t *testing.T) {
	tests := []testcase{
		{"if", "x = 1\ny = 2\nz = if x > y then x else y\nvarchar<z>", &types.VarCharType{Length: 2}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = if !(x < y) then x * 3 else y * 4\nvarchar<z>", &types.VarCharType{Length: 8}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = if x < y then x * 3 else y * 4\nvarchar<z>", &types.VarCharType{Length: 3}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = (x < y) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 3}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = (x <= y) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 3}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = (x = y) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 8}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = (x != y) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 3}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = (x >= y) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 8}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = ((x < y) or (x > y)) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 3}, assert.NoError},
		{"if", "x = 1\ny = 2\nz = ((x < y) and (x > y)) ? x * 3 : y * 4\nvarchar<z>", &types.VarCharType{Length: 8}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseAndTestTypeDerivation(t, &tt)
		})
	}
}

func TestReturnType(t *testing.T) {
	multiplyOutput := `init_scale = max(6, ((S1 + P2) + 1))
init_prec = (((P1 - S1) + P2) + init_scale)
min_scale = min(init_scale, 6)
delta = (init_prec - 38)
prec = min(init_prec, 38)
scale_after_borrow = max((init_scale - delta), min_scale)
scale = (init_prec > 38) ? scale_after_borrow : init_scale
decimal<prec,scale>`
	decimalUniqueParameters := []string{"decimal<P1, S1>", "decimal<P2, S2>"}

	moduloOutput := `init_scale = max(S1, S2)
init_prec = (min((P1 - S1), (P2 - S2)) + init_scale)
min_scale = min(init_scale, 6)
delta = (init_prec - 38)
prec = min(init_prec, 38)
scale_after_borrow = max((init_scale - delta), min_scale)
scale = (init_prec > 38) ? scale_after_borrow : init_scale
decimal<prec,scale>`

	tests := []struct {
		name       string
		parameters []string
		args       []string
		expr       string
		want       types.Type
		wantErr    assert.ErrorAssertionFunc
	}{
		{"SameAsInput", []string{"varchar<L1>"}, []string{"varchar<8>"}, "varchar<L1>",
			&types.VarCharType{Length: 8, Nullability: types.NullabilityRequired}, assert.NoError},
		{"SameAsInput?", []string{"varchar<L1>"}, []string{"varchar<8>"}, "varchar?<L1>",
			&types.VarCharType{Length: 8, Nullability: types.NullabilityNullable}, assert.NoError},
		{"decFixedOutput?", []string{"decimal<P1, S1>"}, []string{"decimal<20,5>"}, "decimal<38,0>",
			&types.DecimalType{Precision: 38, Scale: 0, Nullability: types.NullabilityRequired}, assert.NoError},
		{"decFixedScalePrecisionSameAsInput?", []string{"decimal<P1, S1>"}, []string{"decimal<20,5>"}, "decimal<38,S1>",
			&types.DecimalType{Precision: 38, Scale: 5, Nullability: types.NullabilityRequired}, assert.NoError},
		{"decScaleSameAsInput?", []string{"decimal<P1, 5>"}, []string{"decimal<20,5>"}, "decimal<P1,5>",
			&types.DecimalType{Precision: 20, Scale: 5, Nullability: types.NullabilityRequired}, assert.NoError},
		{"decPrecisisonSameAsInput?", []string{"decimal<25, S1>"}, []string{"decimal<25,9>"}, "decimal<25,S1>",
			&types.DecimalType{Precision: 25, Scale: 9, Nullability: types.NullabilityRequired}, assert.NoError},
		{"decSameAsInput?", []string{"decimal<P1, S1>"}, []string{"decimal<20,10>"}, "decimal?<P1,S1>",
			&types.DecimalType{Precision: 20, Scale: 10, Nullability: types.NullabilityNullable}, assert.NoError},
		{"max", []string{"varchar<L1>", "varchar<L2>"}, []string{"varchar<9>", "varchar<8>"}, "x = max(L1, L2)\nvarchar<x>",
			&types.VarCharType{Length: 9, Nullability: types.NullabilityRequired}, assert.NoError},
		{"maxPrecision", []string{"decimal<P1, 0>", "decimal<P2, 0>"}, []string{"decimal<18,0>", "decimal<27,0>"}, "max_p = max(P1, P2)\ndecimal?<max_p,0>",
			&types.DecimalType{Precision: 27, Scale: 0, Nullability: types.NullabilityNullable}, assert.NoError},
		{"maxScale", []string{"decimal<38, S1>", "decimal<38, S2>"}, []string{"decimal<20,10>", "decimal<20,12>"}, "S3 = max(S1, S2)\ndecimal?<38,S3>",
			&types.DecimalType{Precision: 38, Scale: 12, Nullability: types.NullabilityNullable}, assert.NoError},
		{"primitive", []string{"i8"}, []string{"i8"}, "i8", &types.Int8Type{Nullability: types.NullabilityRequired}, assert.NoError},
		{"multiply1", decimalUniqueParameters, []string{"decimal<20,10>", "decimal<20,10>"}, multiplyOutput,
			&types.DecimalType{Precision: 38, Scale: 8, Nullability: types.NullabilityRequired}, assert.NoError},
		{"multiply2", decimalUniqueParameters, []string{"decimal<10,4>", "decimal<15,7>"}, multiplyOutput,
			&types.DecimalType{Precision: 38, Scale: 17, Nullability: types.NullabilityRequired}, assert.NoError},
		{"multiply3", decimalUniqueParameters, []string{"decimal<5,2>", "decimal<5,2>"}, multiplyOutput,
			&types.DecimalType{Precision: 16, Scale: 8, Nullability: types.NullabilityRequired}, assert.NoError},
		{"multiply4", decimalUniqueParameters, []string{"decimal<38,37>", "decimal<38,37>"}, multiplyOutput,
			&types.DecimalType{Precision: 38, Scale: 6, Nullability: types.NullabilityRequired}, assert.NoError},
		{"multiply5", decimalUniqueParameters, []string{"decimal<20,10>", "decimal<15,5>"}, multiplyOutput,
			&types.DecimalType{Precision: 38, Scale: 13, Nullability: types.NullabilityRequired}, assert.NoError},
		{"multiply6", decimalUniqueParameters, []string{"decimal<12,6>", "decimal<8,4>"}, multiplyOutput,
			&types.DecimalType{Precision: 29, Scale: 15, Nullability: types.NullabilityRequired}, assert.NoError},
		{"modulo1", decimalUniqueParameters, []string{"decimal<10,4>", "decimal<15, 7>"}, moduloOutput,
			&types.DecimalType{Precision: 13, Scale: 7, Nullability: types.NullabilityRequired}, assert.NoError},
		{"modulo2", decimalUniqueParameters, []string{"decimal<5,2>", "decimal<5,2>"}, moduloOutput,
			&types.DecimalType{Precision: 5, Scale: 2, Nullability: types.NullabilityRequired}, assert.NoError},
		{"modulo3", decimalUniqueParameters, []string{"decimal<38,37>", "decimal<38,37>"}, moduloOutput,
			&types.DecimalType{Precision: 38, Scale: 37, Nullability: types.NullabilityRequired}, assert.NoError},
		{"modulo4", decimalUniqueParameters, []string{"decimal<20,10>", "decimal<15,5>"}, moduloOutput,
			&types.DecimalType{Precision: 20, Scale: 10, Nullability: types.NullabilityRequired}, assert.NoError},
		{"modulo5", decimalUniqueParameters, []string{"decimal<12,6>", "decimal<8,4>"}, moduloOutput,
			&types.DecimalType{Precision: 10, Scale: 6, Nullability: types.NullabilityRequired}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			funcParameters := parseFuncParameters(t, tt.parameters)
			funcArguments := parseFuncArguments(t, tt.args)
			resultType, err := parser.ParseType(tt.expr)
			require.NoError(t, err)
			require.NotNil(t, resultType)
			derivation, ok := resultType.(*types.OutputDerivation)
			if !ok {
				derivation = &types.OutputDerivation{FinalType: resultType}
			}
			require.NotNil(t, derivation)
			got, err := derivation.ReturnType(funcParameters, funcArguments)
			if !tt.wantErr(t, err, fmt.Sprintf("Evaluate(%v)", tt.expr)) {
				return
			}
			if tt.want == nil {
				require.Nil(t, got)
				return
			}
			assert.Equalf(t, tt.want, got, "Evaluate(%v)", tt.expr)
			assert.Equal(t, tt.expr, derivation.String())
			assert.Equal(t, tt.want.GetNullability(), derivation.GetNullability())
			assert.Equal(t, tt.want.ShortString(), derivation.ShortString())
			assert.GreaterOrEqual(t, len(tt.want.GetParameters()), len(derivation.GetParameterizedParams()))
		})
	}
}

func parseFuncParameters(t *testing.T, params []string) []types.FuncDefArgType {
	result := make([]types.FuncDefArgType, len(params))
	for i, p := range params {
		var err error
		result[i], err = parser.ParseType(p)
		require.NoError(t, err)
	}
	return result
}

func parseFuncArguments(t *testing.T, args []string) []types.Type {
	result := make([]types.Type, len(args))
	for i, a := range args {
		funcDefArgType, err := parser.ParseType(a)
		require.NoError(t, err)
		result[i], err = funcDefArgType.WithParameters(nil)
		require.NoError(t, err)
	}
	return result
}

func Test_getBinaryOpType(t *testing.T) {
	tests := []struct {
		name string
		want types.BinaryOp
	}{
		{"and", types.And},
		{"or", types.Or},
		{"+", types.Plus},
		{"-", types.Minus},
		{"*", types.Multiply},
		{"/", types.Divide},
		{"<", types.LT},
		{">", types.GT},
		{"<=", types.LTE},
		{">=", types.GTE},
		{"=", types.EQ},
		{"!=", types.NEQ},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, types.GetBinaryOpType(tt.name), "GetBinaryOpType(%v)", tt.name)
			assert.Equalf(t, tt.name, tt.want.String(), "getBinaryOpType(%v)", tt.name)
		})
	}
}
