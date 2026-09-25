// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"errors"
	"fmt"

	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// IntervalDayToSecondToProto encodes the domain interval as its protobuf literal message. It always
// writes the precision precision_mode arm; the deprecated microseconds arm is never emitted.
func IntervalDayToSecondToProto(v *types.IntervalDayToSecond) *proto.Expression_Literal_IntervalDayToSecond {
	if v == nil {
		return nil
	}
	return &proto.Expression_Literal_IntervalDayToSecond{
		Days:       v.Days,
		Seconds:    v.Seconds,
		Subseconds: v.Subseconds,
		PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Precision{
			Precision: v.Precision.ToProtoVal(),
		},
	}
}

// IntervalDayToSecondFromProto decodes a protobuf interval literal message into the domain type.
// An absent precision_mode is rejected: subseconds has no scale without a precision.
func IntervalDayToSecondFromProto(p *proto.Expression_Literal_IntervalDayToSecond) (*types.IntervalDayToSecond, error) {
	if p == nil {
		return nil, nil
	}
	v := &types.IntervalDayToSecond{
		Days:       p.GetDays(),
		Seconds:    p.GetSeconds(),
		Subseconds: p.GetSubseconds(),
	}
	switch m := p.PrecisionMode.(type) {
	case *proto.Expression_Literal_IntervalDayToSecond_Precision:
		v.Precision = types.TimePrecision(m.Precision)
	case *proto.Expression_Literal_IntervalDayToSecond_Microseconds:
		v.Subseconds = int64(m.Microseconds)
		v.Precision = types.PrecisionMicroSeconds
	default:
		return nil, errors.New("interval day to second literal is missing its precision_mode")
	}
	return v, nil
}

// LiteralToProto encodes a literal as its protobuf message.
func LiteralToProto(l expr.Literal) *proto.Expression_Literal {
	switch l := l.(type) {
	case *expr.NullLiteral:
		return nullLiteralToProto(l)
	case *expr.PrimitiveLiteral[bool]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[int8]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[int16]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[int32]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[int64]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[float32]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[float64]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[string]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[types.Timestamp]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[types.Date]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[types.Time]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[types.FixedChar]:
		return primitiveLiteralToProto(l)
	case *expr.PrimitiveLiteral[types.TimestampTz]:
		return primitiveLiteralToProto(l)
	case *expr.NestedLiteral[expr.StructLiteralValue]:
		return nestedLiteralToProto(l)
	case *expr.NestedLiteral[expr.ListLiteralValue]:
		return nestedLiteralToProto(l)
	case *expr.ByteSliceLiteral[[]byte]:
		return byteSliceLiteralToProto(l)
	case *expr.ByteSliceLiteral[types.FixedBinary]:
		return byteSliceLiteralToProto(l)
	case *expr.ByteSliceLiteral[types.UUID]:
		return byteSliceLiteralToProto(l)
	case *expr.MapLiteral:
		return mapLiteralToProto(l)
	case *expr.ProtoLiteral:
		return protoLiteralToProto(l)
	case expr.IntervalCompoundLiteral:
		return intervalCompoundLiteralToProto(l)
	case expr.IntervalYearToMonthLiteral:
		return intervalYearToMonthLiteralToProto(l)
	default:
		panic(fmt.Sprintf("wire: unhandled literal %T", l))
	}
}

func nullLiteralToProto(n *expr.NullLiteral) *proto.Expression_Literal {
	return &proto.Expression_Literal{
		Nullable:               true,
		TypeVariationReference: n.Type.GetTypeVariationReference(),
		LiteralType:            &proto.Expression_Literal_Null{Null: TypeToProto(n.Type)},
	}
}

func primitiveLiteralToProto[T expr.PrimitiveLiteralValue](l *expr.PrimitiveLiteral[T]) *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               l.Type.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: l.Type.GetTypeVariationReference(),
	}

	switch v := any(l.Value).(type) {
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
	case types.Timestamp:
		lit.LiteralType = &proto.Expression_Literal_Timestamp{Timestamp: int64(v)}
	case types.Date:
		lit.LiteralType = &proto.Expression_Literal_Date{Date: int32(v)}
	case types.Time:
		lit.LiteralType = &proto.Expression_Literal_Time{Time: int64(v)}
	case types.FixedChar:
		lit.LiteralType = &proto.Expression_Literal_FixedChar{FixedChar: string(v)}
	case types.TimestampTz:
		lit.LiteralType = &proto.Expression_Literal_TimestampTz{TimestampTz: int64(v)}
	default:
		panic("invalid primitive literal type")
	}

	return lit
}

func nestedLiteralToProto[T expr.StructLiteralValue | expr.ListLiteralValue](l *expr.NestedLiteral[T]) *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               l.Type.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: l.Type.GetTypeVariationReference(),
	}

	vals := make([]*proto.Expression_Literal, len(l.Value))
	for i, v := range l.Value {
		vals[i] = LiteralToProto(v)
	}

	switch any(l.Value).(type) {
	case expr.StructLiteralValue:
		lit.LiteralType = &proto.Expression_Literal_Struct_{
			Struct: &proto.Expression_Literal_Struct{Fields: vals},
		}
	case expr.ListLiteralValue:
		if len(vals) == 0 {
			lit.LiteralType = &proto.Expression_Literal_EmptyList{
				EmptyList: TypeToProto(l.Type).GetList(),
			}
		} else {
			lit.LiteralType = &proto.Expression_Literal_List_{
				List: &proto.Expression_Literal_List{Values: vals},
			}
		}
	}

	return lit
}

func mapLiteralToProto(l *expr.MapLiteral) *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               l.Type.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: l.Type.GetTypeVariationReference(),
	}

	if len(l.Value) == 0 {
		lit.LiteralType = &proto.Expression_Literal_EmptyMap{
			EmptyMap: TypeToProto(l.Type).GetMap(),
		}
	} else {
		kv := make([]*proto.Expression_Literal_Map_KeyValue, len(l.Value))
		for i, v := range l.Value {
			kv[i] = &proto.Expression_Literal_Map_KeyValue{
				Key:   LiteralToProto(v.Key),
				Value: LiteralToProto(v.Value),
			}
		}
		lit.LiteralType = &proto.Expression_Literal_Map_{
			Map: &proto.Expression_Literal_Map{KeyValues: kv},
		}
	}

	return lit
}

func byteSliceLiteralToProto[T ~[]byte](l *expr.ByteSliceLiteral[T]) *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               l.Type.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: l.Type.GetTypeVariationReference(),
	}

	switch v := any(l.Value).(type) {
	case []byte:
		lit.LiteralType = &proto.Expression_Literal_Binary{Binary: v}
	case types.FixedBinary:
		lit.LiteralType = &proto.Expression_Literal_FixedBinary{FixedBinary: v}
	case types.UUID:
		lit.LiteralType = &proto.Expression_Literal_Uuid{Uuid: v}
	}

	return lit
}

func protoLiteralToProto(l *expr.ProtoLiteral) *proto.Expression_Literal {
	lit := &proto.Expression_Literal{
		Nullable:               l.Type.GetNullability() == types.NullabilityNullable,
		TypeVariationReference: l.Type.GetTypeVariationReference(),
	}

	switch literalType := l.Type.(type) {
	case *types.IntervalYearType:
		v := l.Value.(*types.IntervalYearToMonth)
		lit.LiteralType = &proto.Expression_Literal_IntervalYearToMonth_{
			IntervalYearToMonth: &proto.Expression_Literal_IntervalYearToMonth{
				Years:  v.Years,
				Months: v.Months,
			},
		}
	case *types.IntervalDayType:
		v := l.Value.(*types.IntervalDayToSecond)
		lit.LiteralType = &proto.Expression_Literal_IntervalDayToSecond_{
			IntervalDayToSecond: IntervalDayToSecondToProto(v),
		}
	case *types.VarCharType:
		v := l.Value.(string)
		lit.LiteralType = &proto.Expression_Literal_VarChar_{
			VarChar: &proto.Expression_Literal_VarChar{
				Value:  v,
				Length: uint32(literalType.Length),
			},
		}
	case *types.DecimalType:
		v := l.Value.([]byte)
		lit.LiteralType = &proto.Expression_Literal_Decimal_{
			Decimal: &proto.Expression_Literal_Decimal{
				Value:     v,
				Precision: literalType.Precision,
				Scale:     literalType.Scale,
			},
		}
	}

	return lit
}

func nullabilityFromBool(nullable bool) types.Nullability {
	if nullable {
		return types.NullabilityNullable
	}
	return types.NullabilityRequired
}

// LiteralFromProto constructs the appropriate Literal from a protobuf message.
func LiteralFromProto(l *proto.Expression_Literal) expr.Literal {
	nullability := nullabilityFromBool(l.Nullable)

	switch lit := l.LiteralType.(type) {
	case *proto.Expression_Literal_Null:
		return &expr.NullLiteral{Type: TypeFromProto(lit.Null)}
	case *proto.Expression_Literal_Boolean:
		return &expr.PrimitiveLiteral[bool]{
			Value: lit.Boolean,
			Type: &types.BooleanType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I8:
		return &expr.PrimitiveLiteral[int8]{
			Value: int8(lit.I8),
			Type: &types.Int8Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I16:
		return &expr.PrimitiveLiteral[int16]{
			Value: int16(lit.I16),
			Type: &types.Int16Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I32:
		return &expr.PrimitiveLiteral[int32]{
			Value: lit.I32,
			Type: &types.Int32Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_I64:
		return &expr.PrimitiveLiteral[int64]{
			Value: lit.I64,
			Type: &types.Int64Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Fp32:
		return &expr.PrimitiveLiteral[float32]{
			Value: lit.Fp32,
			Type: &types.Float32Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Fp64:
		return &expr.PrimitiveLiteral[float64]{
			Value: lit.Fp64,
			Type: &types.Float64Type{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_String_:
		return &expr.PrimitiveLiteral[string]{
			Value: lit.String_,
			Type: &types.StringType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Timestamp:
		return &expr.PrimitiveLiteral[types.Timestamp]{
			Value: types.Timestamp(lit.Timestamp),
			Type: &types.TimestampType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Date:
		return &expr.PrimitiveLiteral[types.Date]{
			Value: types.Date(lit.Date),
			Type: &types.DateType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Time:
		return &expr.PrimitiveLiteral[types.Time]{
			Value: types.Time(lit.Time),
			Type: &types.TimeType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_FixedChar:
		return &expr.PrimitiveLiteral[types.FixedChar]{
			Value: types.FixedChar(lit.FixedChar),
			Type: &types.FixedCharType{
				Length:           int32(len(lit.FixedChar)),
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_TimestampTz:
		return &expr.PrimitiveLiteral[types.TimestampTz]{
			Value: types.TimestampTz(lit.TimestampTz),
			Type: &types.TimestampTzType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Struct_:
		typeList := make([]types.Type, len(lit.Struct.Fields))
		fields := make([]expr.Literal, len(lit.Struct.Fields))
		for i, f := range lit.Struct.Fields {
			fields[i] = LiteralFromProto(f)
			typeList[i] = fields[i].GetType()
		}

		return &expr.NestedLiteral[expr.StructLiteralValue]{
			Value: expr.StructLiteralValue(fields),
			Type: &types.StructType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Types:            typeList,
			}}
	case *proto.Expression_Literal_List_:
		ret := make(expr.ListLiteralValue, len(lit.List.Values))
		for i, v := range lit.List.Values {
			ret[i] = LiteralFromProto(v)
		}
		return &expr.NestedLiteral[expr.ListLiteralValue]{
			Value: expr.ListLiteralValue(ret),
			Type: &types.ListType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Type:             ret[0].GetType(),
			}}
	case *proto.Expression_Literal_EmptyList:
		return &expr.NestedLiteral[expr.ListLiteralValue]{
			Value: nil,
			Type: &types.ListType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Type:             TypeFromProto(lit.EmptyList.Type),
			}}
	case *proto.Expression_Literal_Map_:
		ret := make(expr.MapLiteralValue, len(lit.Map.KeyValues))
		for i, kv := range lit.Map.KeyValues {
			ret[i].Key = LiteralFromProto(kv.Key)
			ret[i].Value = LiteralFromProto(kv.Value)
		}
		return &expr.MapLiteral{
			Value: ret,
			Type: &types.MapType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Key:              ret[0].Key.GetType(),
				Value:            ret[0].Value.GetType(),
			}}
	case *proto.Expression_Literal_EmptyMap:
		return &expr.MapLiteral{
			Value: nil,
			Type: &types.MapType{
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
				Key:              TypeFromProto(lit.EmptyMap.Key),
				Value:            TypeFromProto(lit.EmptyMap.Value),
			}}
	case *proto.Expression_Literal_Binary:
		return &expr.ByteSliceLiteral[[]byte]{
			Value: lit.Binary,
			Type: &types.BinaryType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_FixedBinary:
		return &expr.ByteSliceLiteral[types.FixedBinary]{
			Value: lit.FixedBinary,
			Type: &types.FixedBinaryType{
				Length:           int32(len(lit.FixedBinary)),
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_Uuid:
		return &expr.ByteSliceLiteral[types.UUID]{
			Value: lit.Uuid,
			Type: &types.UUIDType{
				TypeVariationRef: l.TypeVariationReference,
				Nullability:      nullability,
			}}
	case *proto.Expression_Literal_IntervalDayToSecond_:
		value, err := IntervalDayToSecondFromProto(lit.IntervalDayToSecond)
		if err != nil {
			return nil
		}
		precision, err := types.ProtoToTimePrecision(value.GetPrecisionProtoVal())
		if err != nil {
			return nil
		}
		return &expr.ProtoLiteral{
			Value: value,
			Type: &types.IntervalDayType{
				Precision:        precision,
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_VarChar_:
		return &expr.ProtoLiteral{
			Value: lit.VarChar.Value,
			Type: &types.VarCharType{
				Length:           int32(lit.VarChar.Length),
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_Decimal_:
		return &expr.ProtoLiteral{
			Value: lit.Decimal.Value,
			Type: &types.DecimalType{
				Scale:            lit.Decimal.Scale,
				Precision:        lit.Decimal.Precision,
				Nullability:      nullability,
				TypeVariationRef: l.TypeVariationReference,
			},
		}
	case *proto.Expression_Literal_IntervalCompound_:
		return intervalCompoundLiteralFromProto(l)
	case *proto.Expression_Literal_IntervalYearToMonth_:
		return intervalYearToMonthLiteralFromProto(l)
	}
	panic("unimplemented literal type")
}
