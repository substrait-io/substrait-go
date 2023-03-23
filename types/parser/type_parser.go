// SPDX-License-Identifier: Apache-2.0

package parser

import (
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
	TypeDef  Def  `parser:"@@"`
	Optional bool `parser:"@'?'?"`
}

func (t *Type) ShortType() string {
	return t.TypeDef.ShortType()
}

func (t *Type) String() string {
	var opt string
	if t.Optional {
		opt = "?"
	}
	return t.TypeDef.String() + opt
}

func (t *Type) Type() (types.Type, error) {
	var n types.Nullability
	if t.Optional {
		n = types.NullabilityNullable
	} else {
		n = types.NullabilityRequired
	}

	return t.TypeDef.Type(n)
}

type Def interface {
	String() string
	ShortType() string
	Type(n types.Nullability) (types.Type, error)
}

type typename string

func (t *typename) Capture(values []string) error {
	*t = typename(strings.ToLower(values[0]))
	return nil
}

type nonParamType struct {
	TypeName typename `parser:"@(AnyType | Template | IntType | Boolean | FPType | Temporal | BinaryType)"`
}

func (t *nonParamType) String() string {
	return string(t.TypeName)
}

func (t *nonParamType) ShortType() string {
	if strings.HasPrefix(string(t.TypeName), "any") {
		return "any"
	}

	switch t.TypeName {
	case "timestamp":
		return "ts"
	case "timestamp_tz":
		return "tstz"
	case "interval_day":
		return "iday"
	case "interval_year":
		return "iyear"
	case "string":
		return "str"
	case "binary":
		return "vbin"
	case "boolean":
		return "bool"
	default:
		return string(t.TypeName)
	}

}

func (t *nonParamType) Type(n types.Nullability) (types.Type, error) {
	switch t.TypeName {
	case "i8":
		return &types.Int8Type{Nullability: n}, nil
	case "i16":
		return &types.Int16Type{Nullability: n}, nil
	case "i32":
		return &types.Int32Type{Nullability: n}, nil
	case "i64":
		return &types.Int64Type{Nullability: n}, nil
	case "fp32":
		return &types.Float32Type{Nullability: n}, nil
	case "fp64":
		return &types.Float64Type{Nullability: n}, nil
	case "timestamp":
		return &types.TimestampType{Nullability: n}, nil
	case "timestamp_tz":
		return &types.TimestampTzType{Nullability: n}, nil
	case "date":
		return &types.DateType{Nullability: n}, nil
	case "time":
		return &types.TimeType{Nullability: n}, nil
	case "interval_day":
		return &types.IntervalDayType{Nullability: n}, nil
	case "interval_year":
		return &types.IntervalYearType{Nullability: n}, nil
	case "uuid":
		return &types.UUIDType{Nullability: n}, nil
	case "string":
		return &types.StringType{Nullability: n}, nil
	case "binary":
		return &types.BinaryType{Nullability: n}, nil
	case "boolean":
		return &types.BooleanType{Nullability: n}, nil
	}
	return nil, substraitgo.ErrNotFound
}

type listType struct {
	ElemType TypeExpression `parser:"'list' '<' @@ '>'"`
}

func (*listType) ShortType() string { return "list" }

func (l *listType) String() string { return "list<" + l.ElemType.Expr.String() + ">" }

func (l *listType) Type(n types.Nullability) (types.Type, error) {
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
	case "fixedchar":
		return "fchar"
	case "fixedbinary":
		return "fbin"
	case "varchar":
		return "vchar"
	}
	return ""
}

func (p *lengthType) String() string {
	return p.TypeName + "<" + p.NumericParam.Expr.String() + ">"
}

func (p *lengthType) Type(n types.Nullability) (types.Type, error) {
	lit, ok := p.NumericParam.Expr.(*IntegerLiteral)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	switch p.TypeName {
	case "fixedchar":
		return &types.FixedCharType{
			Length:      lit.Value,
			Nullability: n,
		}, nil
	case "fixedbinary":
		return &types.FixedBinaryType{
			Length:      lit.Value,
			Nullability: n,
		}, nil
	case "varchar":
		return &types.VarCharType{
			Length:      lit.Value,
			Nullability: n,
		}, nil
	default:
		return nil, substraitgo.ErrInvalidType
	}
}

type decimalType struct {
	Precision TypeExpression `parser:"'decimal' '<' @@"`
	Scale     TypeExpression `parser:"',' @@ '>'"`
}

func (*decimalType) ShortType() string { return "dec" }

func (d *decimalType) String() string {
	return "decimal<" + d.Precision.Expr.String() + ", " + d.Scale.Expr.String() + ">"
}

func (d *decimalType) Type(n types.Nullability) (types.Type, error) {
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
	Types []TypeExpression `parser:"'struct' '<' @@ (',' @@)* '>'"`
}

func (*structType) ShortType() string { return "struct" }

func (s *structType) String() string {
	var b strings.Builder
	b.WriteString("struct<")
	for i, t := range s.Types {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(t.Expr.String())
	}
	b.WriteString(">")
	return b.String()
}

func (t *structType) Type(n types.Nullability) (types.Type, error) {
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
	Key   TypeExpression `parser:"'map' '<' @@"`
	Value TypeExpression `parser:"',' @@ '>'"`
}

func (*mapType) ShortType() string { return "map" }

func (m *mapType) String() string {
	return "map<" + m.Key.Expr.String() + "," + m.Value.Expr.String() + ">"
}

func (t *mapType) Type(n types.Nullability) (types.Type, error) {
	k, ok := t.Key.Expr.(*Type)
	if !ok {
		return nil, substraitgo.ErrNotImplemented
	}

	v, ok := t.Value.Expr.(*Type)
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
		{Name: "LengthType", Pattern: `fixedchar|varchar|fixedbinary`},
		{Name: "Int", Pattern: `[-+]?\d+`},
		{Name: "ParamType", Pattern: `(?i)(struct|list|decimal|map)`},
		{Name: "Identifier", Pattern: `[a-zA-Z_$][a-zA-Z_$0-9]*`},
		{Name: "Ident", Pattern: `([a-zA-Z_]\w*)|[><,?]`},
	}, lexer.MatchLongest())
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
		participle.Union[Def](&nonParamType{}, &mapType{}, &listType{}, &structType{}, &lengthType{}, &decimalType{}),
		participle.CaseInsensitive("Boolean", "ParamType", "IntType", "FPType", "Temporal", "BinaryType", "LengthType"),
		participle.Lexer(def),
		participle.UseLookahead(3),
	)
	if err != nil {
		return nil, err
	}

	return &Parser{parser: parser}, nil
}
