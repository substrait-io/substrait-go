// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"bytes"

	"github.com/substrait-io/substrait-go/proto"
	"golang.org/x/exp/slices"
	pb "google.golang.org/protobuf/proto"
)

type PrimitiveLiteralValue interface {
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
		LiteralType:            &proto.Expression_Literal_Null{Null: TypeToProto(n.Type)},
	}
}

func (n *NullLiteral) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Literal_{Literal: n.ToProtoLiteral()},
	}
}

func (n *NullLiteral) Equals(rhs Expression) bool {
	if nl, ok := rhs.(*NullLiteral); ok {
		return nl.Type.Equals(n.Type)
	}

	return false
}

type PrimitiveLiteral[T PrimitiveLiteralValue] struct {
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Value            T
	Type             Type
}

func (b *PrimitiveLiteral[T]) Nullability() bool     { return b.Nullable }
func (b *PrimitiveLiteral[T]) TypeVariation() uint32 { return b.TypeVariationRef }

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
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Value            T
	Type             Type
}

func (b *NestedLiteral[T]) Nullability() bool     { return b.Nullable }
func (b *NestedLiteral[T]) TypeVariation() uint32 { return b.TypeVariationRef }

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
				EmptyList: TypeToProto(t.Type).GetList(),
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
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Value            MapLiteralValue
	Type             Type
}

func (b *MapLiteral) Nullability() bool     { return b.Nullable }
func (b *MapLiteral) TypeVariation() uint32 { return b.TypeVariationRef }

func (t *MapLiteral) ValueType() Type { return t.Type }
func (t *MapLiteral) ToProtoLiteral() *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               t.Nullable,
		TypeVariationReference: t.TypeVariationRef,
	}

	if len(t.Value) == 0 {
		lit.LiteralType = &proto.Expression_Literal_EmptyMap{
			EmptyMap: TypeToProto(t.Type).GetMap(),
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
		return t.Type.Equals(other.Type)
	}
	return false
}

type ByteSliceLiteral[T ~[]byte] struct {
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Value            T
	Type             Type
}

func (b *ByteSliceLiteral[T]) Nullability() bool     { return b.Nullable }
func (b *ByteSliceLiteral[T]) TypeVariation() uint32 { return b.TypeVariationRef }

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
	funcArg

	Nullable         bool
	TypeVariationRef uint32
	Value            pb.Message
	Type             Type
}

func (b *ProtoLiteral) Nullability() bool     { return b.Nullable }
func (b *ProtoLiteral) TypeVariation() uint32 { return b.TypeVariationRef }

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

func NewPrimitiveLiteral[T PrimitiveLiteralValue](val T, nullable bool) Literal {
	var nullability Nullability
	if nullable {
		nullability = NullabilityNullable
	} else {
		nullability = NullabilityRequired
	}
	return &PrimitiveLiteral[T]{
		Nullable: nullable,
		Value:    val,
		Type: &primitiveType[T]{
			Nullability: nullability,
		},
	}
}

type (
	StructLiteral = NestedLiteral[StructLiteralValue]
	ListLiteral   = NestedLiteral[ListLiteralValue]
)

func NewNestedLiteral[T StructLiteralValue | MapLiteralValue | ListLiteralValue](val T, nullable bool) Literal {
	var nullability Nullability
	if nullable {
		nullability = NullabilityNullable
	} else {
		nullability = NullabilityRequired
	}

	switch v := any(val).(type) {
	case StructLiteralValue:
		typeList := make([]Type, len(v))
		for i, f := range v {
			typeList[i] = f.ValueType()
		}
		return &NestedLiteral[StructLiteralValue]{
			Value:    v,
			Nullable: nullable,
			Type: &StructType{
				Nullability: nullability,
				Types:       typeList,
			}}
	case MapLiteralValue:
		return &MapLiteral{
			Value:    v,
			Nullable: nullable,
			Type: &MapType{
				Nullability: nullability,
				Key:         v[0].Key.ValueType(),
				Value:       v[0].Value.ValueType(),
			}}
	case ListLiteralValue:
		return &NestedLiteral[ListLiteralValue]{
			Value:    v,
			Nullable: nullable,
			Type: &ListType{
				Nullability: nullability,
				Type:        v[0].ValueType(),
			}}
	}
	panic("should not get here")
}

func NewEmptyMapLiteral(key, val Type, nullable bool) *MapLiteral {
	var nullability Nullability
	if nullable {
		nullability = NullabilityNullable
	} else {
		nullability = NullabilityRequired
	}

	return &MapLiteral{
		Nullable: nullable,
		Type: &MapType{
			Nullability: nullability,
			Key:         key,
			Value:       val,
		},
	}
}

func NewEmptyListLiteral(t Type, nullable bool) *ListLiteral {
	var nullability Nullability
	if nullable {
		nullability = NullabilityNullable
	} else {
		nullability = NullabilityRequired
	}

	return &NestedLiteral[ListLiteralValue]{
		Nullable: nullable,
		Type: &ListType{
			Nullability: nullability,
			Type:        t,
		}}
}

func LiteralFromProto(l *proto.Expression_Literal) Literal {
	var nullability Nullability
	if l.Nullable {
		nullability = NullabilityNullable
	} else {
		nullability = NullabilityRequired
	}

	switch lit := l.LiteralType.(type) {
	case *proto.Expression_Literal_Boolean:
		return &PrimitiveLiteral[bool]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.Boolean,
			Type: &BooleanType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I8:
		return &PrimitiveLiteral[int8]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            int8(lit.I8),
			Type: &Int8Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I16:
		return &PrimitiveLiteral[int16]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            int16(lit.I16),
			Type: &Int16Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I32:
		return &PrimitiveLiteral[int32]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.I32,
			Type: &Int32Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I64:
		return &PrimitiveLiteral[int64]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.I64,
			Type: &Int64Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Fp32:
		return &PrimitiveLiteral[float32]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.Fp32,
			Type: &Float32Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Fp64:
		return &PrimitiveLiteral[float64]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.Fp64,
			Type: &Float64Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_String_:
		return &PrimitiveLiteral[string]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.String_,
			Type: &StringType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Binary:
		return &ByteSliceLiteral[[]byte]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.Binary,
			Type: &BinaryType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Timestamp:
		return &PrimitiveLiteral[Timestamp]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            Timestamp(lit.Timestamp),
			Type: &TimestampType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Date:
		return &PrimitiveLiteral[Date]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            Date(lit.Date),
			Type: &DateType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Time:
		return &PrimitiveLiteral[Time]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            Time(lit.Time),
			Type: &TimeType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_IntervalYearToMonth_:
		return &ProtoLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.IntervalYearToMonth,
			Type: &IntervalYearType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_IntervalDayToSecond_:
		return &ProtoLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.IntervalDayToSecond,
			Type: &IntervalDayType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_FixedChar:
		return &PrimitiveLiteral[FixedChar]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            FixedChar(lit.FixedChar),
			Type: &FixedCharType{
				Length:           int32(len(lit.FixedChar)),
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_VarChar_:
		return &ProtoLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.VarChar,
			Type: &VarCharType{
				Length:           int32(lit.VarChar.Length),
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_FixedBinary:
		return &ByteSliceLiteral[FixedBinary]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.FixedBinary,
			Type: &FixedBinaryType{
				Length:           int32(len(lit.FixedBinary)),
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Decimal_:
		return &ProtoLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.Decimal,
			Type: &DecimalType{
				Scale:            lit.Decimal.Scale,
				Precision:        lit.Decimal.Precision,
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_TimestampTz:
		return &PrimitiveLiteral[TimestampTz]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            TimestampTz(lit.TimestampTz),
			Type: &BooleanType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Uuid:
		return &ByteSliceLiteral[UUID]{Nullable: l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.Uuid,
			Type: &BooleanType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Null:
		return &NullLiteral{Type: TypeFromProto(lit.Null)}
	case *proto.Expression_Literal_Struct_:
		types := make([]Type, len(lit.Struct.Fields))
		fields := make([]Literal, len(lit.Struct.Fields))
		for i, f := range lit.Struct.Fields {
			fields[i] = LiteralFromProto(f)
			types[i] = fields[i].ValueType()
		}

		return &NestedLiteral[StructLiteralValue]{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            StructLiteralValue(fields),
			Type: &StructType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Types:            types,
			}}
	case *proto.Expression_Literal_Map_:
		ret := make(MapLiteralValue, len(lit.Map.KeyValues))
		for i, kv := range lit.Map.KeyValues {
			ret[i].Key = LiteralFromProto(kv.Key)
			ret[i].Value = LiteralFromProto(kv.Value)
		}
		return &MapLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            ret,
			Type: &MapType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Key:              ret[0].Key.ValueType(),
				Value:            ret[0].Value.ValueType(),
			}}
	case *proto.Expression_Literal_List_:
		ret := make(ListLiteralValue, len(lit.List.Values))
		for i, v := range lit.List.Values {
			ret[i] = LiteralFromProto(v)
		}
		return &NestedLiteral[ListLiteralValue]{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            ListLiteralValue(ret),
			Type: &ListType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Type:             ret[0].ValueType(),
			}}
	case *proto.Expression_Literal_EmptyList:
		return &NestedLiteral[ListLiteralValue]{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            nil,
			Type: &ListType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Type:             TypeFromProto(lit.EmptyList.Type),
			}}
	case *proto.Expression_Literal_EmptyMap:
		return &MapLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            nil,
			Type: &MapType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Key:              TypeFromProto(lit.EmptyMap.Key),
				Value:            TypeFromProto(lit.EmptyMap.Value),
			}}
	case *proto.Expression_Literal_UserDefined_:
		params := make([]TypeParam, len(lit.UserDefined.TypeParameters))
		for i, p := range lit.UserDefined.TypeParameters {
			params[i] = TypeParamFromProto(p)
		}

		return &ProtoLiteral{
			Nullable:         l.Nullable,
			TypeVariationRef: l.TypeVariationReference,
			Value:            lit.UserDefined,
			Type: &UserDefinedType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				TypeReference:    lit.UserDefined.TypeReference,
				TypeParameters:   params,
			},
		}
	}
	panic("unimplemented literal type")
}
