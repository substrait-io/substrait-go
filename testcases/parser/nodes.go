package parser

import (
	"fmt"
	"strconv"

	"github.com/substrait-io/substrait-go/expr"
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
		return nil, fmt.Errorf("Column index must be greater than or equal to 0")
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
