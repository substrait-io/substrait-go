package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/types"
)

type TestFuncType string

const (
	ScalarFuncType    TestFuncType = "scalar"
	AggregateFuncType TestFuncType = "aggregate"
	WindowFuncType    TestFuncType = "window"
)

type CaseLiteral struct {
	Type           types.Type
	ValueText      string
	Value          expr.Literal
	SubstraitError *SubstraitError
}

type TestFileHeader struct {
	Version     string
	FuncType    TestFuncType
	IncludedURI string
}

type TestCase struct {
	BaseURI       string
	GroupDesc     string
	FuncName      string
	Args          []*CaseLiteral
	AggregateArgs []*AggregateArgument
	Result        *CaseLiteral
	Options       FuncOptions
	Columns       [][]expr.Literal
	TableName     string
	ColumnTypes   []types.Type
	FuncType      TestFuncType
}

func (tc *TestCase) GetFunctionOptions() []*types.FunctionOption {
	if len(tc.Options) == 0 {
		return nil
	}
	funcOptions := make([]*types.FunctionOption, 0)
	for key, value := range tc.Options {
		funcOptions = append(funcOptions, &types.FunctionOption{
			Name:       key,
			Preference: []string{value},
		})
	}
	return funcOptions
}

func (tc *TestCase) scalarSignatureKey() string {
	var b strings.Builder
	for i, a := range tc.Args {
		if i != 0 {
			b.WriteByte('_')
		}
		b.WriteString(a.Type.ShortString())
	}
	return b.String()
}

func (tc *TestCase) aggregateSignatureKey() string {
	var b strings.Builder
	for i, a := range tc.AggregateArgs {
		if i != 0 {
			b.WriteByte('_')
		}
		b.WriteString(a.ColumnType.ShortString())
	}
	return b.String()
}

func (tc *TestCase) signatureKey() string {
	if tc.FuncType == ScalarFuncType {
		return tc.scalarSignatureKey()
	}
	return tc.aggregateSignatureKey()
}

func (tc *TestCase) CompoundFunctionName() string {
	return tc.FuncName + ":" + tc.signatureKey()
}

func (tc *TestCase) ID() extensions.ID {
	baseURI := tc.BaseURI
	if !strings.HasPrefix(baseURI, "https") || !strings.HasPrefix(baseURI, "http") {
		baseURI = "https://github.com/substrait-io/substrait/blob/main" + tc.BaseURI
	}
	return extensions.ID{
		URI:  baseURI,
		Name: tc.CompoundFunctionName(),
	}
}

func (tc *TestCase) GetScalarFunctionInvocation(reg *expr.ExtensionRegistry) (*expr.ScalarFunction, error) {
	if tc.FuncType != ScalarFuncType {
		return nil, fmt.Errorf("not a scalar function testcase")
	}
	id := tc.ID()
	args := make([]types.FuncArg, len(tc.Args))
	for i, arg := range tc.Args {
		args[i] = arg.Value
	}

	return expr.NewScalarFunc(*reg, id, tc.GetFunctionOptions(), args...)
}

type TestGroup struct {
	Description string
	TestCases   []*TestCase
}

type TestFile struct {
	Header    *TestFileHeader
	TestCases []*TestCase
}

type FuncOptions map[string]string

type AggregateArgument struct {
	Argument    *CaseLiteral
	TableName   string
	ColumnName  string
	ColumnType  types.Type
	ColumnIndex int
}

func newAggregateArgument(tableName string, columnName string, columnType types.Type) (*AggregateArgument, error) {
	index, err := strconv.ParseInt(columnName[3:], 10, 64)
	if err != nil {
		return nil, err
	}
	if index < 0 {
		return nil, fmt.Errorf("column index must be greater than or equal to 0")
	}
	return &AggregateArgument{
		TableName:   tableName,
		ColumnName:  columnName,
		ColumnType:  columnType,
		ColumnIndex: int(index),
	}, nil
}

type CompactAggregateFuncCall struct {
	FuncName      string
	Rows          [][]expr.Literal
	AggregateArgs []*AggregateArgument
}
