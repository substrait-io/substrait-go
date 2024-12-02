package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/types"
)

type testcase struct {
	name    string
	expr    string
	want    types.Type
	wantErr assert.ErrorAssertionFunc
}

func parseAndTestTypeDerivation(t *testing.T, tt *testcase) {
	resultType, err := ParseType(tt.expr)
	require.NoError(t, err)
	require.NotNil(t, resultType)
	derivation, ok := resultType.(*OutputDerivation)
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

func Test_getBinaryOpType(t *testing.T) {
	tests := []struct {
		name string
		want BinaryOp
	}{
		{"and", And},
		{"or", Or},
		{"+", Plus},
		{"-", Minus},
		{"*", Multiply},
		{"/", Divide},
		{"<", LT},
		{">", GT},
		{"<=", LTE},
		{">=", GTE},
		{"=", EQ},
		{"!=", NEQ},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getBinaryOpType(tt.name), "getBinaryOpType(%v)", tt.name)
		})
	}
}
