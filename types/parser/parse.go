package parser

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/goccy/go-yaml"
	"github.com/substrait-io/substrait-go/v8/types"
	baseparser2 "github.com/substrait-io/substrait-go/v8/types/parser/baseparser"
	"github.com/substrait-io/substrait-go/v8/types/parser/util"
)

type TypeExpression struct {
	ValueType types.FuncDefArgType
}

type UserDefinedTypeResolver func(name string) (urn string, err error)

type userDefinedTypeResolverKey struct{}

func ContextWithUserDefinedTypeResolver(ctx context.Context, resolver UserDefinedTypeResolver) context.Context {
	return context.WithValue(ctx, userDefinedTypeResolverKey{}, resolver)
}

func UserDefinedTypeResolverFromContext(ctx context.Context) UserDefinedTypeResolver {
	resolver, _ := ctx.Value(userDefinedTypeResolverKey{}).(UserDefinedTypeResolver)
	return resolver
}

func (t *TypeExpression) MarshalYAML() (interface{}, error) {
	return t.ValueType.String(), nil
}

func (t *TypeExpression) UnmarshalYAML(ctx context.Context, data []byte) error {
	var typeString string
	if err := yaml.UnmarshalContext(ctx, data, &typeString); err != nil {
		return err
	}

	exp, err := ParseType(typeString, UserDefinedTypeResolverFromContext(ctx))
	if err != nil {
		return err
	}
	t.ValueType = exp
	return nil
}

func ParseType(input string, resolveUserDefinedType UserDefinedTypeResolver) (types.FuncDefArgType, error) {
	is := antlr.NewInputStream(input)
	lexer := baseparser2.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser2.NewSubstraitTypeParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	visitor := &TypeVisitor{ResolveUserDefinedType: resolveUserDefinedType, ErrorListener: errorListener}
	ret, err := parseType(input, p, errorListener, visitor)
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

func parseType(input string, p *baseparser2.SubstraitTypeParser, errorListener *util.SimpleErrorListener, visitor *TypeVisitor) (any, error) {
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
