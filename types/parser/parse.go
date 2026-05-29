package parser

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	baseparser2 "github.com/substrait-io/substrait-go/v8/types/parser/baseparser"
	"github.com/substrait-io/substrait-go/v8/types/parser/util"
)

type TypeExpression struct {
	ValueType types.FuncDefArgType
}

type userDefinedTypeValidatorKey struct{}

// WithUserDefinedTypeValidator returns a context that validates every
// user-defined type encountered while parsing a type expression. It lets
// callers reject references to types that are not declared.
func WithUserDefinedTypeValidator(ctx context.Context, validate func(name string) error) context.Context {
	return context.WithValue(ctx, userDefinedTypeValidatorKey{}, validate)
}

func userDefinedTypeValidatorFromContext(ctx context.Context) func(string) error {
	validate, _ := ctx.Value(userDefinedTypeValidatorKey{}).(func(string) error)
	return validate
}

func (t *TypeExpression) MarshalYAML() (interface{}, error) {
	return t.ValueType.String(), nil
}

func (t *TypeExpression) UnmarshalYAML(ctx context.Context, fn func(interface{}) error) error {
	type Alias any
	var alias Alias
	if err := fn(&alias); err != nil {
		return err
	}

	switch v := alias.(type) {
	case string:
		exp, err := parseType(ctx, v)
		if err != nil {
			return err
		}
		t.ValueType = exp
		return nil
	}

	return substraitgo.ErrNotImplemented
}

func ParseType(input string) (types.FuncDefArgType, error) {
	return parseType(context.Background(), input)
}

func parseType(ctx context.Context, input string) (types.FuncDefArgType, error) {
	is := antlr.NewInputStream(input)
	lexer := baseparser2.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser2.NewSubstraitTypeParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	visitor := &TypeVisitor{
		ErrorListener:               errorListener,
		validateUserDefinedTypeName: userDefinedTypeValidatorFromContext(ctx),
	}
	ret, err := visitType(input, p, errorListener, visitor)
	if err != nil {
		return nil, err
	}
	if errorListener.ErrorCount() > 0 {
		return nil, fmt.Errorf("error parsing input '%s': %s", input, errorListener.GetErrors())
	}
	retType, ok := ret.(types.FuncDefArgType)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s as FuncDefArgType", input)
	}
	return retType, nil
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
