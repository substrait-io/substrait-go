// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"bytes"

	"github.com/substrait-io/substrait-go/proto"
	"golang.org/x/exp/slices"
	pb "google.golang.org/protobuf/proto"
)

type primitiveLiteral interface {
	bool | int8 | int16 | ~int32 | ~int64 |
		float32 | float64 | ~string
}

type nestedLiteral interface {
	StructLiteralValue | ListLiteralValue
}

type (
	IntervalYearToMonth = proto.Expression_Literal_IntervalYearToMonth
	IntervalDayToSecond = proto.Expression_Literal_IntervalDayToSecond
	VarChar             = proto.Expression_Literal_VarChar
	Decimal             = proto.Expression_Literal_Decimal
	StructLiteralValue  []Literal
	ListLiteralValue    []Literal
	MapLiteralValue     []struct {
		Key   Literal
		Value Literal
	}
	UserDefinedLiteral = proto.Expression_Literal_UserDefined
	Null               proto.Type
)

type baseLiteral struct {
	funcArg

	Nullable         bool
	TypeVariationRef uint32
}

func (b *baseLiteral) Nullability() bool     { return b.Nullable }
func (b *baseLiteral) TypeVariation() uint32 { return b.TypeVariationRef }

type Literal interface {
	FuncArg

	Nullability() bool
	TypeVariation() uint32
	ValueType() Type
	Equals(Expression) bool
	ToProto() *proto.Expression
	ToProtoLiteral() *proto.Expression_Literal
}

type NullLiteral struct {
	funcArg
	Type Type
}

func (n *NullLiteral) ValueType() Type       { return n.Type }
func (n *NullLiteral) Nullability() bool     { return true }
func (n *NullLiteral) TypeVariation() uint32 { return n.Type.GetTypeVariationReference() }
func (n *NullLiteral) ToProtoLiteral() *proto.Expression_Literal {
	return &proto.Expression_Literal{
		Nullable:               true,
		TypeVariationReference: n.Type.GetTypeVariationReference(),
		LiteralType:            &proto.Expression_Literal_Null{Null: n.Type.ToProto()},
	}
}

func (n *NullLiteral) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Literal_{Literal: n.ToProtoLiteral()},
	}
}

func (n *NullLiteral) Equals(rhs Expression) bool {
	if nl, ok := rhs.(*NullLiteral); ok {
		return nl.Type == n.Type
	}

	return false
}

type PrimitiveLiteral[T primitiveLiteral] struct {
	baseLiteral

	Value T
	Type  Type
}

func (t *PrimitiveLiteral[T]) ValueType() Type { return t.Type }
func (t *PrimitiveLiteral[T]) ToProtoLiteral() *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               t.Nullable,
		TypeVariationReference: t.TypeVariationRef,
	}

	switch v := any(t.Value).(type) {
	case bool:
		lit.LiteralType = &proto.Expression_Literal_Boolean{Boolean: v}
	case int8:
		lit.LiteralType = &proto.Expression_Literal_I8{I8: int32(v)}
	case int16:
		lit.LiteralType = &proto.Expression_Literal_I16{I16: int32(v)}
	case int32:
		lit.LiteralType = &proto.Expression_Literal_I32{I32: v}
	case int64:
		lit.LiteralType = &proto.Expression_Literal_I64{I64: v}
	case float32:
		lit.LiteralType = &proto.Expression_Literal_Fp32{Fp32: v}
	case float64:
		lit.LiteralType = &proto.Expression_Literal_Fp64{Fp64: v}
	case string:
		lit.LiteralType = &proto.Expression_Literal_String_{String_: v}
	case Timestamp:
		lit.LiteralType = &proto.Expression_Literal_Timestamp{Timestamp: int64(v)}
	case Date:
		lit.LiteralType = &proto.Expression_Literal_Date{Date: int32(v)}
	case Time:
		lit.LiteralType = &proto.Expression_Literal_Time{Time: int64(v)}
	case FixedChar:
		lit.LiteralType = &proto.Expression_Literal_FixedChar{FixedChar: string(v)}
	case TimestampTz:
		lit.LiteralType = &proto.Expression_Literal_TimestampTz{TimestampTz: int64(v)}
	default:
		panic("invalid primitive literal type")
	}

	return lit
}

func (t *PrimitiveLiteral[T]) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Literal_{Literal: t.ToProtoLiteral()},
	}
}

func (t *PrimitiveLiteral[T]) Equals(rhs Expression) bool {
	if other, ok := rhs.(*PrimitiveLiteral[T]); ok {
		return t.Value == other.Value
	}
	return false
}

type NestedLiteral[T nestedLiteral] struct {
	baseLiteral

	Value T
	Type  Type
}

func (t *NestedLiteral[T]) ValueType() Type { return t.Type }
func (t *NestedLiteral[T]) ToProtoLiteral() *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               t.Nullable,
		TypeVariationReference: t.TypeVariationRef,
	}

	vals := make([]*proto.Expression_Literal, len(t.Value))
	for i, l := range t.Value {
		vals[i] = l.ToProtoLiteral()
	}

	switch any(t.Value).(type) {
	case StructLiteralValue:
		lit.LiteralType = &proto.Expression_Literal_Struct_{
			Struct: &proto.Expression_Literal_Struct{
				Fields: vals,
			},
		}
	case ListLiteralValue:
		if len(vals) == 0 {
			lit.LiteralType = &proto.Expression_Literal_EmptyList{
				EmptyList: t.Type.ToProto().GetList(),
			}
		} else {
			lit.LiteralType = &proto.Expression_Literal_List_{
				List: &proto.Expression_Literal_List{
					Values: vals,
				},
			}
		}
	}

	return lit
}

func (t *NestedLiteral[T]) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Literal_{Literal: t.ToProtoLiteral()},
	}
}

func (t *NestedLiteral[T]) Equals(rhs Expression) bool {
	if other, ok := rhs.(*NestedLiteral[T]); ok {
		return slices.EqualFunc(t.Value, other.Value, func(a, b Literal) bool {
			return a.Equals(b)
		})
	}
	return false
}

type MapLiteral struct {
	baseLiteral

	Value MapLiteralValue
	Type  Type
}

func (t *MapLiteral) ValueType() Type { return t.Type }
func (t *MapLiteral) ToProtoLiteral() *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               t.Nullable,
		TypeVariationReference: t.TypeVariationRef,
	}

	if len(t.Value) == 0 {
		lit.LiteralType = &proto.Expression_Literal_EmptyMap{
			EmptyMap: t.Type.ToProto().GetMap(),
		}
	} else {
		kv := make([]*proto.Expression_Literal_Map_KeyValue, len(t.Value))
		for i, v := range t.Value {
			kv[i].Key = v.Key.ToProtoLiteral()
			kv[i].Value = v.Value.ToProtoLiteral()
		}

		lit.LiteralType = &proto.Expression_Literal_Map_{
			Map: &proto.Expression_Literal_Map{KeyValues: kv},
		}
	}

	return lit
}

func (t *MapLiteral) ToProto() *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Literal_{
		Literal: t.ToProtoLiteral(),
	}}
}

func (t *MapLiteral) Equals(rhs Expression) bool {
	if other, ok := rhs.(*MapLiteral); ok {
		if t.Type != other.Type {
			return false
		}

		return false
	}
	return false
}

type ByteSliceLiteral[T ~[]byte] struct {
	baseLiteral

	Value T
	Type  Type
}

func (t *ByteSliceLiteral[T]) ValueType() Type { return t.Type }
func (t *ByteSliceLiteral[T]) ToProtoLiteral() *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               t.Nullable,
		TypeVariationReference: t.TypeVariationRef,
	}

	switch v := any(t.Value).(type) {
	case []byte:
		lit.LiteralType = &proto.Expression_Literal_Binary{Binary: v}
	case FixedBinary:
		lit.LiteralType = &proto.Expression_Literal_FixedBinary{FixedBinary: v}
	case UUID:
		lit.LiteralType = &proto.Expression_Literal_Uuid{Uuid: v}
	}

	return lit
}

func (t *ByteSliceLiteral[T]) ToProto() *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Literal_{
		Literal: t.ToProtoLiteral(),
	}}
}

func (t *ByteSliceLiteral[T]) Equals(rhs Expression) bool {
	if other, ok := rhs.(*ByteSliceLiteral[T]); ok {
		return bytes.Equal(t.Value, other.Value)
	}

	return false
}

type ProtoLiteral struct {
	baseLiteral

	Value pb.Message
	Type  Type
}

func (t *ProtoLiteral) ValueType() Type { return t.Type }

func (t *ProtoLiteral) ToProtoLiteral() *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               t.Nullable,
		TypeVariationReference: t.TypeVariationRef,
	}

	switch v := t.Value.(type) {
	case *UserDefinedLiteral:
		lit.LiteralType = &proto.Expression_Literal_UserDefined_{UserDefined: v}
	case *IntervalYearToMonth:
		lit.LiteralType = &proto.Expression_Literal_IntervalYearToMonth_{
			IntervalYearToMonth: v,
		}
	case *IntervalDayToSecond:
		lit.LiteralType = &proto.Expression_Literal_IntervalDayToSecond_{
			IntervalDayToSecond: v,
		}
	case *VarChar:
		lit.LiteralType = &proto.Expression_Literal_VarChar_{VarChar: v}
	case *Decimal:
		lit.LiteralType = &proto.Expression_Literal_Decimal_{Decimal: v}
	}

	return lit
}

func (t *ProtoLiteral) ToProto() *proto.Expression {
	return &proto.Expression{RexType: &proto.Expression_Literal_{
		Literal: t.ToProtoLiteral(),
	}}
}

func (t *ProtoLiteral) Equals(rhs Expression) bool {
	if other, ok := rhs.(*ProtoLiteral); ok {
		return pb.Equal(t.Value, other.Value)
	}
	return false
}

func LiteralFromProto(l *proto.Expression_Literal) Literal {
	base := baseLiteral{Nullable: l.Nullable, TypeVariationRef: l.TypeVariationReference}
	baseType := baseType{TypeVariationRef: l.TypeVariationReference}
	if l.Nullable {
		baseType.Nullability = NullabilityNullable
	} else {
		baseType.Nullability = NullabilityRequired
	}

	switch lit := l.LiteralType.(type) {
	case *proto.Expression_Literal_Boolean:
		return &PrimitiveLiteral[bool]{base, lit.Boolean, &BooleanType{baseType: baseType}}
	case *proto.Expression_Literal_I8:
		return &PrimitiveLiteral[int8]{base, int8(lit.I8), &Int8Type{baseType: baseType}}
	case *proto.Expression_Literal_I16:
		return &PrimitiveLiteral[int16]{base, int16(lit.I16), &Int16Type{baseType: baseType}}
	case *proto.Expression_Literal_I32:
		return &PrimitiveLiteral[int32]{base, int32(lit.I32), &Int32Type{baseType: baseType}}
	case *proto.Expression_Literal_I64:
		return &PrimitiveLiteral[int64]{base, int64(lit.I64), &Int64Type{baseType: baseType}}
	case *proto.Expression_Literal_Fp32:
		return &PrimitiveLiteral[float32]{base, lit.Fp32, &Float32Type{baseType: baseType}}
	case *proto.Expression_Literal_Fp64:
		return &PrimitiveLiteral[float64]{base, lit.Fp64, &Float64Type{baseType: baseType}}
	case *proto.Expression_Literal_String_:
		return &PrimitiveLiteral[string]{base, lit.String_, &StringType{baseType: baseType}}
	case *proto.Expression_Literal_Binary:
		return &ByteSliceLiteral[[]byte]{base, lit.Binary, &BinaryType{baseType: baseType}}
	case *proto.Expression_Literal_Timestamp:
		return &PrimitiveLiteral[Timestamp]{base, Timestamp(lit.Timestamp), &TimestampType{baseType: baseType}}
	case *proto.Expression_Literal_Date:
		return &PrimitiveLiteral[Date]{base, Date(lit.Date), &DateType{baseType: baseType}}
	case *proto.Expression_Literal_Time:
		return &PrimitiveLiteral[Time]{base, Time(lit.Time), &TimeType{baseType: baseType}}
	case *proto.Expression_Literal_IntervalYearToMonth_:
		return &ProtoLiteral{base, lit.IntervalYearToMonth, &IntervalYearType{baseType: baseType}}
	case *proto.Expression_Literal_IntervalDayToSecond_:
		return &ProtoLiteral{base, lit.IntervalDayToSecond, &IntervalDayType{baseType: baseType}}
	case *proto.Expression_Literal_FixedChar:
		return &PrimitiveLiteral[FixedChar]{base, FixedChar(lit.FixedChar), &FixedCharType{
			baseType: baseType, Length: int32(len(lit.FixedChar))}}
	case *proto.Expression_Literal_VarChar_:
		return &ProtoLiteral{base, lit.VarChar, &VarCharType{baseType: baseType, Length: int32(lit.VarChar.Length)}}
	case *proto.Expression_Literal_FixedBinary:
		return &ByteSliceLiteral[FixedBinary]{base, FixedBinary(lit.FixedBinary), &FixedBinaryType{baseType: baseType,
			Length: int32(len(lit.FixedBinary))}}
	case *proto.Expression_Literal_Decimal_:
		return &ProtoLiteral{base, lit.Decimal, &DecimalType{baseType: baseType,
			Scale: lit.Decimal.Scale, Precision: lit.Decimal.Precision}}
	case *proto.Expression_Literal_TimestampTz:
		return &PrimitiveLiteral[TimestampTz]{base, TimestampTz(lit.TimestampTz), &TimestampTzType{baseType: baseType}}
	case *proto.Expression_Literal_Uuid:
		return &ByteSliceLiteral[UUID]{base, UUID(lit.Uuid), &UUIDType{baseType: baseType}}
	case *proto.Expression_Literal_Null:
		return &NullLiteral{Type: TypeFromProto(lit.Null)}
	case *proto.Expression_Literal_Struct_:
		types := make([]Type, len(lit.Struct.Fields))
		fields := make([]Literal, len(lit.Struct.Fields))
		for i, f := range lit.Struct.Fields {
			fields[i] = LiteralFromProto(f)
			types[i] = fields[i].ValueType()
		}

		return &NestedLiteral[StructLiteralValue]{base, StructLiteralValue(fields),
			&StructType{baseType: baseType, Types: types}}
	case *proto.Expression_Literal_Map_:
		ret := make(MapLiteralValue, len(lit.Map.KeyValues))
		for i, kv := range lit.Map.KeyValues {
			ret[i].Key = LiteralFromProto(kv.Key)
			ret[i].Value = LiteralFromProto(kv.Value)
		}
		return &MapLiteral{base, ret, &MapType{baseType: baseType,
			Key: ret[0].Key.ValueType(), Value: ret[0].Value.ValueType()}}
	case *proto.Expression_Literal_List_:
		ret := make(ListLiteralValue, len(lit.List.Values))
		for i, v := range lit.List.Values {
			ret[i] = LiteralFromProto(v)
		}
		return &NestedLiteral[ListLiteralValue]{base, ret, &ListType{baseType: baseType,
			Type: ret[0].ValueType()}}
	case *proto.Expression_Literal_EmptyList:
		return &NestedLiteral[ListLiteralValue]{base, nil, &ListType{baseType: baseType,
			Type: TypeFromProto(lit.EmptyList.Type)}}
	case *proto.Expression_Literal_EmptyMap:
		return &MapLiteral{base, nil, &MapType{baseType: baseType,
			Key: TypeFromProto(lit.EmptyMap.Key), Value: TypeFromProto(lit.EmptyMap.Value)}}
	case *proto.Expression_Literal_UserDefined_:
		params := make([]TypeParam, len(lit.UserDefined.TypeParameters))
		for i, p := range lit.UserDefined.TypeParameters {
			params[i] = TypeParamFromProto(p)
		}
		return &ProtoLiteral{base,
			lit.UserDefined, &UserDefinedType{baseType: baseType,
				TypeReference:  lit.UserDefined.TypeReference,
				TypeParameters: params}}
	}
	panic("unimplemented literal type")
}
