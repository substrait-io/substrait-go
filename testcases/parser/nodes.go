package parser

import (
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/types"
)

type CaseLiteral struct {
	Type           types.Type
	ValueText      string
	Value          expr.Literal
	SubstraitError *SubstraitError
}

type TestFileHeader struct {
	Version     string
	IncludedURI string
}

type TestCase struct {
	BaseURI       string
	GroupDesc     string
	FuncName      string
	Args          []*CaseLiteral
	AggregateArgs []*AggregateArgument
	Result        *CaseLiteral
	FuncOptions   map[string]string
	options       FuncOptions
	Rows          [][]expr.Literal
	TableName     string
	ColumnTypes   []types.FuncDefArgType
}

type TestFile struct {
	Header    TestFileHeader
	TestCases []*TestCase
}

type FuncOptions map[string]string

type AggregateArgument struct {
	Argument   *CaseLiteral
	TableName  string
	ColumnName string
	ColumnType types.Type
}
