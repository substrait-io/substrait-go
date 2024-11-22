package parser

import (
	"embed"
	"fmt"
	"io"
	"io/fs"

	"github.com/antlr4-go/antlr/v4"
	"github.com/substrait-io/substrait-go/testcases/parser/baseparser"
	"github.com/substrait-io/substrait-go/types/parser/util"
)

func ParseTestCaseFileFromFS(fs embed.FS, s string) (*TestFile, error) {
	file, err := fs.Open(s)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseTestCasesFromFile(file)
}

func ParseTestCasesFromFile(input fs.File) (*TestFile, error) {
	buf, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}
	is := antlr.NewInputStream(string(buf))
	return parseTestCasesFromStream(is, fmt.Sprintf("file %s", input))
}

func ParseTestCasesFromString(input string) (*TestFile, error) {
	is := antlr.NewInputStream(input)
	return parseTestCasesFromStream(is, input)
}

func parseTestCasesFromStream(is *antlr.InputStream, debugStr string) (*TestFile, error) {
	lexer := baseparser.NewFuncTestCaseLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser.NewFuncTestCaseParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	testFile, err := parseTestCases(p, errorListener, debugStr)
	if errorListener.ErrorCount() > 0 {
		return nil, fmt.Errorf("error parsing input '%s': %s", debugStr, errorListener.GetErrors())
	}
	return testFile, err
}

func parseTestCases(p *baseparser.FuncTestCaseParser, errorListener util.VisitErrorListener, debugStr string) (*TestFile, error) {
	var err error
	defer util.TransformPanicToError(&err, debugStr, "ParseExpr", errorListener)

	visitor := &TestCaseVisitor{ErrorListener: errorListener}
	context := p.Doc()
	if errorListener.ErrorCount() > 0 {
		fmt.Printf("ParseTree: %v", antlr.TreesStringTree(context, []string{}, p))
		return nil, fmt.Errorf("error parsing input '%s': %s", debugStr, errorListener.GetErrors())
	}
	ret := visitor.Visit(context)
	if errorListener.ErrorCount() > 0 {
		return nil, fmt.Errorf("error parsing input '%s': %s", debugStr, errorListener.GetErrors())
	}
	retType, ok := ret.(*TestFile)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s as FuncDefArgType", debugStr)
	}
	return retType, err
}
