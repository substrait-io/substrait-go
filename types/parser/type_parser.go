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
)

var defaultParser *Parser

type TypeExpression struct {
	Expr Expression `parser:"@@"`
}

func (t TypeExpression) String() string { return t.Expr.String() }

func (t TypeExpression) Type() (types.Type, error) {
	typeDef, ok := t.Expr.(Def)
	if !ok {
		return nil, errors.New("type expression doesn't represent type")
	}
	return typeDef.Type()
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

func (t *Type) Type() (types.Type, error) {
	return t.TypeDef.Type()
}

type Def interface {
	String() string
	ShortType() string
	Type() (types.Type, error)
	Optional() bool
}

type typename string

func (t *typename) Capture(values []string) error {
	*t = typename(strings.ToLower(values[0]))
	return nil
}

type nonParamType struct {
	TypeName    typename `parser:"@(IntType | Boolean | FPType | Temporal | BinaryType)"`
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

func (t *nonParamType) Type() (types.Type, error) {
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

func (l *listType) Type() (types.Type, error) {
	var n types.Nullability
	if l.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	if t, ok := l.ElemType.Expr.(*Type); ok {
		ret, err := t.Type()
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

type lengthType struct {
	TypeName     string         `parser:"@LengthType '<'"`
	NumericParam TypeExpression `parser:"@@ '>'"`
}

func (p *lengthType) ShortType() string {
	switch p.TypeName {
	case "fixedchar", "varchar", "fixedbinary":
		return types.GetShortTypeName(types.TypeName(p.TypeName))
	case "precision_timestamp":
		return "prets"
	case "precision_timestamp_tz":
		return "pretstz"
	}
	return ""
}

func (p *lengthType) String() string {
	return p.TypeName + "<" + p.NumericParam.Expr.String() + ">"
}

func (p *lengthType) Optional() bool { return false }

func (p *lengthType) Type() (types.Type, error) {
	var n types.Nullability

	var typ types.Type
	var err error
	switch t := p.NumericParam.Expr.(type) {
	case *IntegerLiteral:
		typ, err = getFixedTypeFromConcreteParam(p.TypeName, t)
	case *ParamName:
		typ, err = getParameterizedTypeSingleParam(p.TypeName, t)
	default:
		return nil, substraitgo.ErrNotImplemented
	}
	if err != nil {
		return nil, err
	}
	return typ.WithNullability(n), nil
}

func getFixedTypeFromConcreteParam(name string, param *IntegerLiteral) (types.Type, error) {
	typeName := types.TypeName(name)
	switch typeName {
	case types.TypeNamePrecisionTimestamp:
		precision, err := types.ProtoToTimePrecision(param.Value)
		if err != nil {
			return nil, err
		}
		return types.NewPrecisionTimestampType(precision), nil
	case types.TypeNamePrecisionTimestampTz:
		precision, err := types.ProtoToTimePrecision(param.Value)
		if err != nil {
			return nil, err
		}
		return types.NewPrecisionTimestampTzType(precision), nil
	}
	typ, err := types.FixedTypeNameToType(typeName)
	if err != nil {
		return nil, err
	}
	return typ.WithLength(param.Value), nil
}

func getParameterizedTypeSingleParam(typeName string, param *ParamName) (types.Type, error) {
	intParam := types.IntegerParam{Name: param.Name}
	switch types.TypeName(typeName) {
	case types.TypeNameVarChar:
		return &types.ParameterizedVarCharType{IntegerOption: intParam}, nil
	case types.TypeNameFixedChar:
		return &types.ParameterizedFixedCharType{IntegerOption: intParam}, nil
	case types.TypeNameFixedBinary:
		return &types.ParameterizedFixedBinaryType{IntegerOption: intParam}, nil
	case types.TypeNamePrecisionTimestamp:
		return &types.ParameterizedPrecisionTimestampType{IntegerOption: intParam}, nil
	case types.TypeNamePrecisionTimestampTz:
		return &types.ParameterizedPrecisionTimestampTzType{IntegerOption: intParam}, nil
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
	return "decimal" + opt + "<" + d.Precision.Expr.String() + ", " + d.Scale.Expr.String() + ">"
}

func (d *decimalType) Optional() bool { return d.Nullability }

func (d *decimalType) Type() (types.Type, error) {
	var n types.Nullability
	pi, ok1 := d.Precision.Expr.(*IntegerLiteral)
	si, ok2 := d.Scale.Expr.(*IntegerLiteral)
	if ok1 && ok2 {
		// concrete decimal param
		return &types.DecimalType{
			Nullability: n,
			Precision:   pi.Value,
			Scale:       si.Value,
		}, nil
	}

	ps, ok1 := d.Precision.Expr.(*ParamName)
	ss, ok2 := d.Scale.Expr.(*ParamName)
	if ok1 && ok2 {
		// parameterized decimal param
		return &types.ParameterizedDecimalType{
			Nullability: n,
			Precision:   types.IntegerParam{Name: ps.Name},
			Scale:       types.IntegerParam{Name: ss.Name},
		}, nil
	}
	return nil, substraitgo.ErrNotImplemented
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

func (t *structType) Type() (types.Type, error) {
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

		if typeList[i], err = tp.Type(); err != nil {
			return nil, err
		}
	}
	return &types.StructType{
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
	return "map" + opt + "<" + m.Key.Expr.String() + "," + m.Value.Expr.String() + ">"
}

func (m *mapType) Optional() bool { return m.Nullability }

func (m *mapType) Type() (types.Type, error) {
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

	key, err := k.Type()
	if err != nil {
		return nil, err
	}

	value, err := v.Type()
	if err != nil {
		return nil, err
	}
	return &types.MapType{
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

func (anyType) Optional() bool { return false }

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

func (t anyType) Type() (types.Type, error) {
	var n types.Nullability
	if t.Nullability {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}
	typeName := string(t.TypeName)
	if strings.HasPrefix(typeName, "any") {
		return &types.AnyType{Name: "any", Nullability: n}, nil
	}
	return &types.AnyType{Name: typeName, Nullability: n}, nil
}

var (
	def = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "whitespace", Pattern: `[ \t]+`},
		{Name: "Template", Pattern: `T`},
		{Name: "AnyType", Pattern: `any[\d]?`},
		{Name: "Boolean", Pattern: `(?i)boolean`},
		{Name: "IntType", Pattern: `i(8|16|32|64)`},
		{Name: "FPType", Pattern: `fp(32|64)`},
		{Name: "Temporal", Pattern: `timestamp(_tz)?|date|time|interval_day|interval_year`},
		{Name: "BinaryType", Pattern: `string|binary|uuid`},
		{Name: "LengthType", Pattern: `fixedchar|varchar|fixedbinary|precision_timestamp_tz|precision_timestamp`},
		{Name: "Int", Pattern: `[-+]?\d+`},
		{Name: "ParamType", Pattern: `(?i)(struct|list|decimal|map)`},
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
