package parser

import (
	"fmt"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
	baseparser2 "github.com/substrait-io/substrait-go/v8/types/parser/baseparser"
	"github.com/substrait-io/substrait-go/v8/types/parser/util"
)

type TypeExpression struct {
	ValueType types.FuncDefArgType
}

var qualifiedUDTPattern = regexp.MustCompile(`(^|[^A-Za-z0-9_$])([A-Za-z_$][A-Za-z0-9_$]*)\.[uU]!([A-Za-z_$][A-Za-z0-9_$]*)`)

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
	parseInput, aliases := replaceQualifiedUDTs(input)
	is := antlr.NewInputStream(parseInput)
	lexer := baseparser2.NewSubstraitTypeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := baseparser2.NewSubstraitTypeParser(stream)
	errorListener := util.NewSimpleErrorListener()
	p.AddErrorListener(errorListener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	visitor := &TypeVisitor{}
	ret, err := parseType(parseInput, p, errorListener, visitor)
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
	applyQualifiedUDTAliases(retType, aliases)
	return retType, nil
}

func replaceQualifiedUDTs(input string) (string, map[string]qualifiedUDT) {
	aliases := make(map[string]qualifiedUDT)
	idx := 0
	replaced := qualifiedUDTPattern.ReplaceAllStringFunc(input, func(match string) string {
		parts := qualifiedUDTPattern.FindStringSubmatch(match)
		placeholder := fmt.Sprintf("__substrait_qualified_udt_%d", idx)
		idx++
		aliases[placeholder] = qualifiedUDT{Alias: parts[2], Name: parts[3]}
		return parts[1] + "u!" + placeholder
	})
	return replaced, aliases
}

type qualifiedUDT struct {
	Alias string
	Name  string
}

func applyQualifiedUDTAliases(typ types.FuncDefArgType, aliases map[string]qualifiedUDT) {
	if len(aliases) == 0 || typ == nil {
		return
	}

	switch t := typ.(type) {
	case *types.ParameterizedUserDefinedType:
		if qualified, ok := aliases[t.Name]; ok {
			t.DependencyAlias = qualified.Alias
			t.Name = qualified.Name
		}
		for _, param := range t.TypeParameters {
			if dataParam, ok := param.(*types.DataTypeUDTParam); ok {
				applyQualifiedUDTAliases(dataParam.Type, aliases)
			}
		}
	case *types.ParameterizedListType:
		applyQualifiedUDTAliases(t.Type, aliases)
	case *types.ParameterizedMapType:
		applyQualifiedUDTAliases(t.Key, aliases)
		applyQualifiedUDTAliases(t.Value, aliases)
	case *types.ParameterizedStructType:
		for _, field := range t.Types {
			applyQualifiedUDTAliases(field, aliases)
		}
	case *types.ParameterizedFuncType:
		for _, param := range t.Parameters {
			applyQualifiedUDTAliases(param, aliases)
		}
		applyQualifiedUDTAliases(t.Return, aliases)
	}
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
