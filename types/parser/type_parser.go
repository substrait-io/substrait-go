// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/types"
	"github.com/substrait-io/substrait-go/types/integer_parameters"
)

var defaultParser *Parser

type TypeExpression struct {
	Expr Expression `parser:"@@"`
}

func (t TypeExpression) String() string { return t.Expr.String() }

func (t TypeExpression) Type() (types.FuncDefArgType, error) {
	typeDef, ok := t.Expr.(Def)
	if !ok {
		return nil, errors.New("type expression doesn't represent type")
	}
	return typeDef.ArgType()
}

func (t TypeExpression) MarshalYAML() (interface{}, error) {
	return t.Expr.String(), nil
}

func (t *TypeExpression) UnmarshalYAML(fn func(interface{}) error) error {
	type Alias any
	var alias Alias
	if err := fn(&alias); err != nil {
		return err
	}

	if defaultParser == nil {
		defaultParser, _ = New()
	}

	switch v := alias.(type) {
	case string:
		exp, err := defaultParser.ParseString(v)
		if err != nil {
			return err
		}
		*t = *exp
		return nil
	}

	return substraitgo.ErrNotImplemented
}

type Expression interface {
	String() string
}

// TODO: implement UnaryOp, BinaryOp, IfElse, and ReturnProgram

type ParamName struct {
	Name string `parser:"@Identifier"`
}

func (p *ParamName) String() string {
	return p.Name
}

type IntegerLiteral struct {
	Value int32 `parser:"@Int"`
}

func (l *IntegerLiteral) String() string {
	return strconv.Itoa(int(l.Value))
}

type Type struct {
	TypeDef Def `parser:"@@"`
}

func (t *Type) ShortType() string {
	return t.TypeDef.ShortType()
}

func (t *Type) Optional() bool { return t.TypeDef.Optional() }

func (t *Type) String() string {
	return t.TypeDef.String()
}

func (t *Type) ArgType() (types.FuncDefArgType, error) {
	return t.TypeDef.ArgType()
}

func (t *Type) RetType() (types.Type, error) {
	return t.TypeDef.RetType()
}

type Def interface {
	String() string
	ShortType() string
	// ArgType indicates argument type
	ArgType() (types.FuncDefArgType, error)
	Optional() bool
	// TODO RetType indicates return type
	// This should be replaced with TypeDerivation method. Currently it just returns concrete type
	RetType() (types.Type, error)
}

type typename string

func (t *typename) Capture(values []string) error {
	*t = typename(strings.ToLower(values[0]))
	return nil
}

type nonParamType struct {
	TypeName    typename `parser:"@(IntType | Boolean | FPType | Temporal | BinaryType | UserDefinedType)"`
	Nullability bool     `parser:"@'?'?"`
	// Variation   int      `parser:"'[' @\d+ ']'?"`
}

func (t *nonParamType) Optional() bool { return t.Nullability }

func (t *nonParamType) String() string {
	opt := string(t.TypeName)
	if t.Nullability {
		opt += "?"
	}
	return opt
}

func (t *nonParamType) ShortType() string {
	return types.GetShortTypeName(types.TypeName(t.TypeName))
}

func (t *nonParamType) RetType() (types.Type, error) {
	var n types.Nullability
	if t.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	typ, err := types.SimpleTypeNameToType(types.TypeName(t.TypeName))
	if err == nil {
		return typ.WithNullability(n), nil
	}
	return nil, err
}

func (t *nonParamType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability
	if t.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	typ, err := types.SimpleTypeNameToType(types.TypeName(t.TypeName))
	if err != nil {
		return nil, err
	}
	funcArgType := typ.(types.FuncDefArgType)
	return funcArgType.SetNullability(n), nil
}

type listType struct {
	Nullability bool           `parser:"'list' @'?'?"`
	ElemType    TypeExpression `parser:"'<' @@ '>'"`
}

func (*listType) ShortType() string { return "list" }

func (l *listType) String() string {
	var opt string
	if l.Nullability {
		opt = "?"
	}
	return "list" + opt + "<" + l.ElemType.Expr.String() + ">"
}

func (l *listType) Optional() bool { return false }

func (l *listType) RetType() (types.Type, error) {
	var n types.Nullability
	if l.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	if t, ok := l.ElemType.Expr.(*Type); ok {
		ret, err := t.RetType()
		if err != nil {
			return nil, err
		}
		return &types.ListType{
			Nullability: n,
			Type:        ret,
		}, nil
	}

	return nil, substraitgo.ErrNotImplemented
}

func (l *listType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability
	if l.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	if t, ok := l.ElemType.Expr.(*Type); ok {
		ret, err := t.ArgType()
		if err != nil {
			return nil, err
		}
		return &types.ParameterizedListType{
			Nullability: n,
			Type:        ret,
		}, nil
	}
	return nil, substraitgo.ErrNotImplemented
}

type lengthType struct {
	TypeName     string         `parser:"@LengthType"`
	Nullability  bool           `parser:"@'?'? '<'"`
	NumericParam TypeExpression `parser:"@@ '>'"`
}

func (p *lengthType) ShortType() string {
	switch p.TypeName {
	case "fixedchar", "varchar", "fixedbinary", "interval_day":
		return types.GetShortTypeName(types.TypeName(p.TypeName))
	case "precision_timestamp":
		return "prets"
	case "precision_timestamp_tz":
		return "pretstz"
	}
	return ""
}

func (p *lengthType) String() string {
	var opt string
	if p.Nullability {
		opt = "?"
	}
	return p.TypeName + opt + "<" + p.NumericParam.Expr.String() + ">"
}

func (p *lengthType) Optional() bool { return false }

func (p *lengthType) RetType() (types.Type, error) {
	var n types.Nullability
	if p.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	lit, ok := p.NumericParam.Expr.(*IntegerLiteral)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	typ, err := types.FixedTypeNameToType(types.TypeName(p.TypeName))
	if err != nil {
		return nil, err
	}
	return typ.WithLength(lit.Value).WithNullability(n), nil
}

func (p *lengthType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability

	var leafParam integer_parameters.IntegerParameter
	switch t := p.NumericParam.Expr.(type) {
	case *IntegerLiteral:
		leafParam = integer_parameters.NewConcreteIntParam(t.Value)
	case *ParamName:
		leafParam = integer_parameters.NewVariableIntParam(t.Name)
	default:
		return nil, substraitgo.ErrNotImplemented
	}
	typ, err := getParameterizedTypeSingleParam(p.TypeName, leafParam, n)
	if err != nil {
		return nil, err
	}
	return typ, nil
}

func getParameterizedTypeSingleParam(typeName string, leafParam integer_parameters.IntegerParameter, n types.Nullability) (types.FuncDefArgType, error) {
	switch types.TypeName(typeName) {
	case types.TypeNameVarChar:
		return &types.ParameterizedVarCharType{IntegerOption: leafParam, Nullability: n}, nil
	case types.TypeNameFixedChar:
		return &types.ParameterizedFixedCharType{IntegerOption: leafParam, Nullability: n}, nil
	case types.TypeNameFixedBinary:
		return &types.ParameterizedFixedBinaryType{IntegerOption: leafParam, Nullability: n}, nil
	case types.TypeNamePrecisionTimestamp:
		return &types.ParameterizedPrecisionTimestampType{IntegerOption: leafParam, Nullability: n}, nil
	case types.TypeNamePrecisionTimestampTz:
		return &types.ParameterizedPrecisionTimestampTzType{IntegerOption: leafParam, Nullability: n}, nil
	default:
		return nil, substraitgo.ErrNotImplemented
	}
}

type decimalType struct {
	Nullability bool           `parser:"'decimal' @'?'?"`
	Precision   TypeExpression `parser:"'<' @@"`
	Scale       TypeExpression `parser:"',' @@ '>'"`
}

func (*decimalType) ShortType() string { return "dec" }

func (d *decimalType) String() string {
	var opt string
	if d.Nullability {
		opt = "?"
	}
	return "decimal" + opt + "<" + d.Precision.Expr.String() + "," + d.Scale.Expr.String() + ">"
}

func (d *decimalType) Optional() bool { return d.Nullability }

func (d *decimalType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability
	if d.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	var precision integer_parameters.IntegerParameter
	if pi, ok := d.Precision.Expr.(*IntegerLiteral); ok {
		precision = integer_parameters.NewConcreteIntParam(pi.Value)
	} else {
		ps := d.Precision.Expr.(*ParamName)
		precision = integer_parameters.NewVariableIntParam(ps.String())
	}

	var scale integer_parameters.IntegerParameter
	if si, ok := d.Scale.Expr.(*IntegerLiteral); ok {
		scale = integer_parameters.NewConcreteIntParam(si.Value)
	} else {
		ss := d.Scale.Expr.(*ParamName)
		scale = integer_parameters.NewVariableIntParam(ss.String())
	}

	return &types.ParameterizedDecimalType{
		Nullability: n,
		Precision:   precision,
		Scale:       scale,
	}, nil
}

func (d *decimalType) RetType() (types.Type, error) {
	var n types.Nullability
	if d.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	p, ok := d.Precision.Expr.(*IntegerLiteral)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}
	s, ok := d.Scale.Expr.(*IntegerLiteral)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}
	return &types.DecimalType{
		Nullability: n,
		Precision:   p.Value,
		Scale:       s.Value,
	}, nil
}

type structType struct {
	Nullability bool             `parser:"'struct' @'?'?"`
	Types       []TypeExpression `parser:"'<' @@ (',' @@)* '>'"`
}

func (*structType) ShortType() string { return "struct" }

func (s *structType) String() string {
	var b strings.Builder
	b.WriteString("struct")
	if s.Nullability {
		b.WriteByte('?')
	}
	b.WriteByte('<')
	for i, t := range s.Types {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(t.Expr.String())
	}
	b.WriteString(">")
	return b.String()
}

func (t *structType) Optional() bool { return t.Nullability }

func (t *structType) RetType() (types.Type, error) {
	var n types.Nullability
	if t.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	var err error
	typeList := make([]types.Type, len(t.Types))
	for i, typ := range t.Types {
		tp, ok := typ.Expr.(*Type)
		if !ok {
			return nil, substraitgo.ErrNotImplemented
		}

		if typeList[i], err = tp.RetType(); err != nil {
			return nil, err
		}
	}
	return &types.StructType{
		Nullability: n,
		Types:       typeList,
	}, nil
}

func (t *structType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability
	if t.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	var err error
	typeList := make([]types.FuncDefArgType, len(t.Types))
	for i, typ := range t.Types {
		tp, ok := typ.Expr.(*Type)
		if !ok {
			return nil, substraitgo.ErrNotImplemented
		}

		if typeList[i], err = tp.ArgType(); err != nil {
			return nil, err
		}
	}
	return &types.ParameterizedStructType{
		Nullability: n,
		Types:       typeList,
	}, nil
}

type mapType struct {
	Nullability bool           `parser:"'map' @'?'?"`
	Key         TypeExpression `parser:"'<' @@"`
	Value       TypeExpression `parser:"',' @@ '>'"`
}

func (*mapType) ShortType() string { return "map" }

func (m *mapType) String() string {
	var opt string
	if m.Nullability {
		opt = "?"
	}
	return "map" + opt + "<" + m.Key.Expr.String() + ", " + m.Value.Expr.String() + ">"
}

func (m *mapType) Optional() bool { return m.Nullability }

func (m *mapType) RetType() (types.Type, error) {
	var n types.Nullability
	if m.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}

	k, ok := m.Key.Expr.(*Type)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	v, ok := m.Value.Expr.(*Type)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	key, err := k.RetType()
	if err != nil {
		return nil, err
	}

	value, err := v.RetType()
	if err != nil {
		return nil, err
	}
	return &types.MapType{
		Key:         key,
		Value:       value,
		Nullability: n,
	}, nil
}

func (m *mapType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability
	if m.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}

	k, ok := m.Key.Expr.(*Type)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	v, ok := m.Value.Expr.(*Type)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	key, err := k.ArgType()
	if err != nil {
		return nil, err
	}

	value, err := v.ArgType()
	if err != nil {
		return nil, err
	}

	return &types.ParameterizedMapType{
		Key:         key,
		Value:       value,
		Nullability: n,
	}, nil
}

// parser token for any
type anyType struct {
	TypeName    typename `parser:"@(AnyType|Template)"`
	Nullability bool     `parser:"@'?'?"`
}

func (t anyType) Optional() bool { return t.Nullability }

func (t anyType) String() string {
	opt := string(t.TypeName)
	if t.Nullability {
		opt += "?"
	}
	return opt
}

func (t anyType) ShortType() string {
	if strings.HasPrefix(string(t.TypeName), "any") {
		return "any"
	}
	return string(t.TypeName)
}

func (t anyType) ArgType() (types.FuncDefArgType, error) {
	var n types.Nullability
	if t.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	typeName := string(t.TypeName)
	if strings.HasPrefix(typeName, "any") {
		return types.AnyType{Name: "any", Nullability: n}, nil
	}
	return types.AnyType{Name: typeName, Nullability: n}, nil
}

func (t anyType) RetType() (types.Type, error) {
	panic("any type can't be in return type")
}

var (
	def = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "whitespace", Pattern: `[ \t]+`},
		{Name: "Template", Pattern: `T`},
		{Name: "AnyType", Pattern: `any[\d]?`},
		{Name: "Boolean", Pattern: `(?i)boolean`},
		{Name: "IntType", Pattern: `i(8|16|32|64)`},
		{Name: "FPType", Pattern: `fp(32|64)`},
		{Name: "Temporal", Pattern: `timestamp(_tz)?|date|time|interval_year`},
		{Name: "BinaryType", Pattern: `string|binary|uuid`},
		{Name: "LengthType", Pattern: `fixedchar|varchar|fixedbinary|precision_timestamp_tz|precision_timestamp|interval_day`},
		{Name: "Int", Pattern: `[-+]?\d+`},
		{Name: "ParamType", Pattern: `(?i)(struct|list|decimal|map)`},
		{Name: "UserDefinedType", Pattern: `u![a-zA-Z_][a-zA-Z0-9_]*`},
		{Name: "Identifier", Pattern: `[a-zA-Z_$][a-zA-Z_$0-9]*`},
		{Name: "Ident", Pattern: `([a-zA-Z_]\w*)|[><,?]`},
	})
)

type Parser struct {
	parser *participle.Parser[TypeExpression]
}

func (p *Parser) Parse(r io.Reader) (*TypeExpression, error) {
	return p.parser.Parse("expression", r)
}

func (p *Parser) ParseString(str string) (*TypeExpression, error) {
	return p.parser.ParseString("expression", str)
}

func (p *Parser) ParseBytes(expr []byte) (*TypeExpression, error) {
	return p.parser.ParseBytes("expression", expr)
}

func New() (*Parser, error) {
	parser, err := participle.Build[TypeExpression](
		participle.Union[Expression](&Type{}, &IntegerLiteral{}, &ParamName{}),
		participle.Union[Def](&anyType{}, &nonParamType{}, &mapType{}, &listType{}, &structType{}, &lengthType{}, &decimalType{}),
		participle.CaseInsensitive("Boolean", "ParamType", "IntType", "FPType", "Temporal", "BinaryType", "LengthType"),
		participle.Lexer(def),
		participle.UseLookahead(3),
	)
	if err != nil {
		return nil, err
	}

	return &Parser{parser: parser}, nil
}
