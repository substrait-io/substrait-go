package types

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/parser/types/baseparser"
	"github.com/substrait-io/substrait-go/types"
)

type TypeExpression struct {
	ValueType types.FuncDefArgType
}

func (t *TypeExpression) MarshalYAML() (interface{}, error) {
	return t.ValueType.String(), nil
}

func (t *TypeExpression) UnmarshalYAML(fn func(interface{}) error) error {
	type Alias any
	var alias Alias
	if err := fn(&alias); err != nil {
		return err
	}

	switch v := alias.(type) {
	case string:
		exp, err := ParseFuncDefArgType(v)
		if err != nil {
			return err
		}
		t.ValueType = exp
		return nil
	}

	return substraitgo.ErrNotImplemented
}

type simpleErrorListener struct {
	errorCount int
	errors     []string
}

func (l *simpleErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errorCount++
	l.errors = append(l.errors, fmt.Sprintf("Syntax error at line %d:%d: %s ", line, column, msg))
}

func (l *simpleErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *simpleErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *simpleErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

func (l *simpleErrorListener) ErrorCount() int {
	return l.errorCount
}

func (l *simpleErrorListener) GetError() string {
	return l.errors[0]
}

func (l *simpleErrorListener) GetErrors() []string {
	return l.errors
}

func newErrorListener() *simpleErrorListener {
	return new(simpleErrorListener)
}

func transformPanicToError(err *error, input, ctxStr string) {
	if r := recover(); r != nil {
		switch t := r.(type) {
		case string:
			*err = fmt.Errorf("failed %s %s with error: %s", ctxStr, input, t)
		case error:
			*err = t
		default:
			*err = fmt.Errorf("failed %s %s with unknown panic", ctxStr, input)
		}
	}
}

func ParseSubstraitType(input string) (types.Type, error) {
	var err error
	defer transformPanicToError(&err, input, "ParseExpr")
	is := antlr.NewInputStream(input)
	lexer := baseparser.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser.NewSubstraitTypeParser(stream)
	errorListener := newErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	visitor := &MyVisitor{}
	context := p.TypeStatement()
	if errorListener.ErrorCount() > 0 {
		fmt.Printf("ParseTree: %v", antlr.TreesStringTree(context, []string{}, p))
		return nil, fmt.Errorf(errorListener.GetError())
	}
	ret := visitor.Visit(context)
	return ret.(types.Type), nil
}

func ParseFuncDefArgType(input string) (types.FuncDefArgType, error) {
	var err error
	defer transformPanicToError(&err, input, "ParseExpr")
	is := antlr.NewInputStream(input)
	lexer := baseparser.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser.NewSubstraitTypeParser(stream)
	errorListener := newErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	visitor := &MyVisitor{}
	context := p.StartRule()
	if errorListener.ErrorCount() > 0 {
		fmt.Printf("ParseTree: %v", antlr.TreesStringTree(context, []string{}, p))
		return nil, fmt.Errorf(errorListener.GetError())
	}
	ret := visitor.Visit(context)
	retType, ok := ret.(types.FuncDefArgType)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s as FuncDefArgType", input)
	}
	return retType, nil
}
