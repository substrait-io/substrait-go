package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	baseparser2 "github.com/substrait-io/substrait-go/v8/types/parser/baseparser"
	"github.com/substrait-io/substrait-go/v8/types/parser/util"
)

type TypeExpression struct {
	ValueType types.FuncDefArgType
	// UserDefinedTypes holds the names of user-defined types referenced by this
	// expression, captured while parsing.
	UserDefinedTypes []string
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
		exp, refs, err := parseType(v)
		if err != nil {
			return err
		}
		t.ValueType = exp
		t.UserDefinedTypes = refs
		return nil
	}

	return substraitgo.ErrNotImplemented
}

func ParseType(input string) (types.FuncDefArgType, error) {
	typ, _, err := parseType(input)
	return typ, err
}

func parseType(input string) (types.FuncDefArgType, []string, error) {
	is := antlr.NewInputStream(input)
	lexer := baseparser2.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser2.NewSubstraitTypeParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	visitor := &TypeVisitor{ErrorListener: errorListener}
	ret, err := visitType(input, p, errorListener, visitor)
	if err != nil {
		return nil, nil, err
	}
	if errorListener.ErrorCount() > 0 {
		return nil, nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	retType, ok := ret.(types.FuncDefArgType)
	if !ok {
		return nil, nil, fmt.Errorf("failed to parse %s as FuncDefArgType", input)
	}
	return retType, visitor.userDefinedTypes, nil
}

func visitType(input string, p *baseparser2.SubstraitTypeParser, errorListener *util.SimpleErrorListener, visitor *TypeVisitor) (any, error) {
	var err error
	defer util.TransformPanicToError(&err, input, "ParseExpr", errorListener)
	context := p.StartRule()
	if errorListener.ErrorCount() > 0 {
		fmt.Printf("ParseTree: %v", antlr.TreesStringTree(context, []string{}, p))
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	ret := visitor.Visit(context)

	return ret, err
}
