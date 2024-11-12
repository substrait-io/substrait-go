package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/types"
	baseparser2 "github.com/substrait-io/substrait-go/types/parser/baseparser"
	"github.com/substrait-io/substrait-go/types/parser/util"
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
		exp, err := ParseType(v)
		if err != nil {
			return err
		}
		t.ValueType = exp
		return nil
	}

	return substraitgo.ErrNotImplemented
}

func ParseType(input string) (types.FuncDefArgType, error) {
	var err error
	is := antlr.NewInputStream(input)
	lexer := baseparser2.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser2.NewSubstraitTypeParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	defer util.TransformPanicToError(&err, input, "ParseExpr", errorListener)

	visitor := &TypeVisitor{}
	context := p.StartRule()
	if errorListener.ErrorCount() > 0 {
		fmt.Printf("ParseTree: %v", antlr.TreesStringTree(context, []string{}, p))
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	ret := visitor.Visit(context)
	if errorListener.ErrorCount() > 0 {
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	retType, ok := ret.(types.FuncDefArgType)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s as FuncDefArgType", input)
	}
	return retType, nil
}
