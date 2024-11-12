package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/testcases/parser/baseparser"
	"github.com/substrait-io/substrait-go/types/parser/util"
)

func ParseTestCaseFile(input string) (*TestFile, error) {
	is := antlr.NewInputStream(input)
	lexer := baseparser.NewFuncTestCaseLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser.NewFuncTestCaseParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	testFile, err := parseTestCases(p, errorListener, input)
	if errorListener.ErrorCount() > 0 {
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	return testFile, err
}

func parseTestCases(p *baseparser.FuncTestCaseParser, errorListener util.VisitErrorListener, input string) (*TestFile, error) {
	var err error
	defer util.TransformPanicToError(&err, input, "ParseExpr", errorListener)

	visitor := &TestCaseVisitor{ErrorListener: errorListener}
	context := p.Doc()
	if errorListener.ErrorCount() > 0 {
		fmt.Printf("ParseTree: %v", antlr.TreesStringTree(context, []string{}, p))
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	ret := visitor.Visit(context)
	if errorListener.ErrorCount() > 0 {
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	retType, ok := ret.(*TestFile)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s as FuncDefArgType", input)
	}
	return retType, err
}
