// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/substrait-io/substrait-go/proto"
)

type Nullability = proto.Type_Nullability

const (
	NullabilityUnspecified = proto.Type_NULLABILITY_UNSPECIFIED
	NullabilityNullable    = proto.Type_NULLABILITY_NULLABLE
	NullabilityRequired    = proto.Type_NULLABILITY_REQUIRED
)

type AggregationPhase = proto.AggregationPhase

const (
	AggPhaseUnspecified                = proto.AggregationPhase_AGGREGATION_PHASE_UNSPECIFIED
	AggPhaseInitialToIntermediate      = proto.AggregationPhase_AGGREGATION_PHASE_INITIAL_TO_INTERMEDIATE
	AggPhaseIntermediateToIntermediate = proto.AggregationPhase_AGGREGATION_PHASE_INTERMEDIATE_TO_INTERMEDIATE
	AggPhaseInitialToResult            = proto.AggregationPhase_AGGREGATION_PHASE_INITIAL_TO_RESULT
	AggPhaseIntermediateToResult       = proto.AggregationPhase_AGGREGATION_PHASE_INTERMEDIATE_TO_RESULT
)

type AggregationInvocation = proto.AggregateFunction_AggregationInvocation

const (
	AggInvocationUnspecified = proto.AggregateFunction_AGGREGATION_INVOCATION_UNSPECIFIED
	AggInvocationAll         = proto.AggregateFunction_AGGREGATION_INVOCATION_ALL
	AggInvocationDistinct    = proto.AggregateFunction_AGGREGATION_INVOCATION_DISTINCT
)

type SortDirection proto.SortField_SortDirection

const (
	SortUnspecified    = SortDirection(proto.SortField_SORT_DIRECTION_UNSPECIFIED)
	SortAscNullsFirst  = SortDirection(proto.SortField_SORT_DIRECTION_ASC_NULLS_FIRST)
	SortAscNullsLast   = SortDirection(proto.SortField_SORT_DIRECTION_ASC_NULLS_LAST)
	SortDescNullsFirst = SortDirection(proto.SortField_SORT_DIRECTION_DESC_NULLS_FIRST)
	SortDescNullsLast  = SortDirection(proto.SortField_SORT_DIRECTION_ASC_NULLS_LAST)
	SortClustered      = SortDirection(proto.SortField_SORT_DIRECTION_CLUSTERED)
)

func (SortDirection) isSortKind() {}

type FunctionRef uint32

func (FunctionRef) isSortKind() {}

type CastFailBehavior = proto.Expression_Cast_FailureBehavior

const (
	BehaviorUnspecified    = proto.Expression_Cast_FAILURE_BEHAVIOR_UNSPECIFIED
	BehaviorReturnNil      = proto.Expression_Cast_FAILURE_BEHAVIOR_RETURN_NULL
	BehaviorThrowException = proto.Expression_Cast_FAILURE_BEHAVIOR_THROW_EXCEPTION
)

// TypeFromProto returns the appropriate Type object from a protobuf
// type message.
func TypeFromProto(t *proto.Type) Type {
	switch t := t.Kind.(type) {
	case *proto.Type_Bool:
		return &BooleanType{
			Nullability:      t.Bool.Nullability,
			TypeVariationRef: t.Bool.TypeVariationReference,
		}
	case *proto.Type_I8_:
		return &Int8Type{
			Nullability:      t.I8.Nullability,
			TypeVariationRef: t.I8.TypeVariationReference,
		}
	case *proto.Type_I16_:
		return &Int16Type{
			Nullability:      t.I16.Nullability,
			TypeVariationRef: t.I16.TypeVariationReference,
		}
	case *proto.Type_I32_:
		return &Int32Type{
			Nullability:      t.I32.Nullability,
			TypeVariationRef: t.I32.TypeVariationReference,
		}
	case *proto.Type_I64_:
		return &Int64Type{
			Nullability:      t.I64.Nullability,
			TypeVariationRef: t.I64.TypeVariationReference,
		}
	case *proto.Type_Fp32:
		return &Float32Type{
			Nullability:      t.Fp32.Nullability,
			TypeVariationRef: t.Fp32.TypeVariationReference,
		}
	case *proto.Type_Fp64:
		return &Float64Type{
			Nullability:      t.Fp64.Nullability,
			TypeVariationRef: t.Fp64.TypeVariationReference,
		}
	case *proto.Type_String_:
		return &StringType{
			Nullability:      t.String_.Nullability,
			TypeVariationRef: t.String_.TypeVariationReference,
		}
	case *proto.Type_Binary_:
		return &BinaryType{
			Nullability:      t.Binary.Nullability,
			TypeVariationRef: t.Binary.TypeVariationReference,
		}
	case *proto.Type_Timestamp_:
		return &TimestampType{
			Nullability:      t.Timestamp.Nullability,
			TypeVariationRef: t.Timestamp.TypeVariationReference,
		}
	case *proto.Type_Date_:
		return &DateType{
			Nullability:      t.Date.Nullability,
			TypeVariationRef: t.Date.TypeVariationReference,
		}
	case *proto.Type_Time_:
		return &TimeType{
			Nullability:      t.Time.Nullability,
			TypeVariationRef: t.Time.TypeVariationReference,
		}
	case *proto.Type_IntervalYear_:
		return &IntervalYearType{
			Nullability:      t.IntervalYear.Nullability,
			TypeVariationRef: t.IntervalYear.TypeVariationReference,
		}
	case *proto.Type_IntervalDay_:
		return &IntervalDayType{
			Nullability:      t.IntervalDay.Nullability,
			TypeVariationRef: t.IntervalDay.TypeVariationReference,
		}
	case *proto.Type_TimestampTz:
		return &TimestampTzType{
			Nullability:      t.TimestampTz.Nullability,
			TypeVariationRef: t.TimestampTz.TypeVariationReference,
		}
	case *proto.Type_Uuid:
		return &UUIDType{
			Nullability:      t.Uuid.Nullability,
			TypeVariationRef: t.Uuid.TypeVariationReference,
		}
	case *proto.Type_FixedBinary_:
		return &FixedBinaryType{
			Nullability:      t.FixedBinary.Nullability,
			TypeVariationRef: t.FixedBinary.TypeVariationReference,
			Length:           t.FixedBinary.Length,
		}
	case *proto.Type_FixedChar_:
		return &FixedCharType{
			Nullability:      t.FixedChar.Nullability,
			TypeVariationRef: t.FixedChar.TypeVariationReference,
			Length:           t.FixedChar.Length,
		}
	case *proto.Type_Decimal_:
		return &DecimalType{
			Nullability:      t.Decimal.Nullability,
			TypeVariationRef: t.Decimal.TypeVariationReference,
			Scale:            t.Decimal.Scale,
			Precision:        t.Decimal.Precision,
		}
	case *proto.Type_Struct_:
		fields := make([]Type, len(t.Struct.Types))
		for i, f := range t.Struct.Types {
			fields[i] = TypeFromProto(f)
		}
		return &StructType{
			Nullability:      t.Struct.Nullability,
			TypeVariationRef: t.Struct.TypeVariationReference,
			Types:            fields,
		}
	case *proto.Type_List_:
		return &ListType{
			Nullability:      t.List.Nullability,
			TypeVariationRef: t.List.TypeVariationReference,
			Type:             TypeFromProto(t.List.Type),
		}
	case *proto.Type_Map_:
		return &MapType{
			Nullability:      t.Map.Nullability,
			TypeVariationRef: t.Map.TypeVariationReference,
			Key:              TypeFromProto(t.Map.Key),
			Value:            TypeFromProto(t.Map.Value),
		}
	case *proto.Type_UserDefined_:
		params := make([]TypeParam, len(t.UserDefined.TypeParameters))
		for i, p := range t.UserDefined.TypeParameters {
			params[i] = TypeParamFromProto(p)
		}
		return &UserDefinedType{
			Nullability:      t.UserDefined.Nullability,
			TypeVariationRef: t.UserDefined.TypeVariationReference,
			TypeReference:    t.UserDefined.TypeReference,
			TypeParameters:   params,
		}
	}
	panic("unimplemented type from proto")
}

type (
	Date        int32
	FixedChar   string
	Time        int64
	Timestamp   int64
	TimestampTz int64
	FixedBinary []byte
	UUID        []byte
	Enum        string

	FunctionOption = proto.FunctionOption

	// FuncArg corresponds to the protobuf FunctionArgument. Anything
	// which could be a function argument should meet this interface.
	// This is either an Expression, a Type, or an Enum (string).
	FuncArg interface {
		fmt.Stringer
		ToProtoFuncArg() *proto.FunctionArgument
	}

	SortKind interface {
		isSortKind()
	}

	// Type corresponds to the proto.Type message and represents
	// a specific type.
	Type interface {
		FuncArg
		RootRefType
		fmt.Stringer
		GetType() Type
		GetNullability() Nullability
		GetTypeVariationReference() uint32
		Equals(Type) bool
	}
)

// TypeToProto properly constructs the appropriate protobuf message
// for the given type.
func TypeToProto(t Type) *proto.Type {
	switch t := t.(type) {
	case *BooleanType:
		return &proto.Type{Kind: &proto.Type_Bool{
			Bool: &proto.Type_Boolean{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *Int8Type:
		return &proto.Type{Kind: &proto.Type_I8_{
			I8: &proto.Type_I8{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *Int16Type:
		return &proto.Type{Kind: &proto.Type_I16_{
			I16: &proto.Type_I16{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *Int32Type:
		return &proto.Type{Kind: &proto.Type_I32_{
			I32: &proto.Type_I32{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *Int64Type:
		return &proto.Type{Kind: &proto.Type_I64_{
			I64: &proto.Type_I64{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *Float32Type:
		return &proto.Type{Kind: &proto.Type_Fp32{
			Fp32: &proto.Type_FP32{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *Float64Type:
		return &proto.Type{Kind: &proto.Type_Fp64{
			Fp64: &proto.Type_FP64{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *StringType:
		return &proto.Type{Kind: &proto.Type_String_{
			String_: &proto.Type_String{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *BinaryType:
		return &proto.Type{Kind: &proto.Type_Binary_{
			Binary: &proto.Type_Binary{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *DateType:
		return &proto.Type{Kind: &proto.Type_Date_{
			Date: &proto.Type_Date{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *TimeType:
		return &proto.Type{Kind: &proto.Type_Time_{
			Time: &proto.Type_Time{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *TimestampTzType:
		return &proto.Type{Kind: &proto.Type_TimestampTz{
			TimestampTz: &proto.Type_TimestampTZ{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *TimestampType:
		return &proto.Type{Kind: &proto.Type_Timestamp_{
			Timestamp: &proto.Type_Timestamp{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *IntervalYearType:
		return &proto.Type{Kind: &proto.Type_IntervalYear_{
			IntervalYear: &proto.Type_IntervalYear{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *IntervalDayType:
		return &proto.Type{Kind: &proto.Type_IntervalDay_{
			IntervalDay: &proto.Type_IntervalDay{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *UUIDType:
		return &proto.Type{Kind: &proto.Type_Uuid{
			Uuid: &proto.Type_UUID{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *FixedCharType:
		return &proto.Type{Kind: &proto.Type_FixedChar_{
			FixedChar: &proto.Type_FixedChar{
				Length:                 t.Length,
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *VarCharType:
		return &proto.Type{Kind: &proto.Type_Varchar{
			Varchar: &proto.Type_VarChar{
				Length:                 t.Length,
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *FixedBinaryType:
		return &proto.Type{Kind: &proto.Type_FixedBinary_{
			FixedBinary: &proto.Type_FixedBinary{
				Length:                 t.Length,
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef}}}
	case *DecimalType:
		return t.ToProto()
	case *StructType:
		return t.ToProto()
	case *ListType:
		return t.ToProto()
	case *MapType:
		return t.ToProto()
	}
	panic("unimplemented type")
}

type primitiveTypeIFace interface {
	PrimitiveLiteralValue | []byte | IntervalYearToMonth |
		IntervalDayToSecond | UUID
}

var primitiveNames = map[reflect.Type]string{
	reflect.PointerTo(reflect.TypeOf(true)):           "boolean",
	reflect.PointerTo(reflect.TypeOf(int8(0))):        "i8",
	reflect.PointerTo(reflect.TypeOf(int16(0))):       "i16",
	reflect.PointerTo(reflect.TypeOf(int32(0))):       "i32",
	reflect.PointerTo(reflect.TypeOf(int64(0))):       "i64",
	reflect.PointerTo(reflect.TypeOf(float32(0))):     "fp32",
	reflect.PointerTo(reflect.TypeOf(float64(0))):     "fp64",
	reflect.PointerTo(reflect.TypeOf([]byte{})):       "binary",
	reflect.PointerTo(reflect.TypeOf("")):             "string",
	reflect.PointerTo(reflect.TypeOf(Timestamp(0))):   "timestamp",
	reflect.PointerTo(reflect.TypeOf(Date(0))):        "date",
	reflect.PointerTo(reflect.TypeOf(Time(0))):        "time",
	reflect.PointerTo(reflect.TypeOf(TimestampTz(0))): "timestamp_tz",
	reflect.PointerTo(reflect.TypeOf(UUID{})):         "uuid",
	reflect.TypeOf(&IntervalYearToMonth{}):            "interval_year",
	reflect.TypeOf(&IntervalDayToSecond{}):            "interval_day",
}

// PrimitiveType is a generic implementation of simple primitive types
// which only need to track if they are nullable and if they are a type
// variation.
type PrimitiveType[T primitiveTypeIFace] struct {
	Nullability      Nullability
	TypeVariationRef uint32
}

func (*PrimitiveType[T]) isRootRef()                          {}
func (s *PrimitiveType[T]) GetType() Type                     { return s }
func (s *PrimitiveType[T]) GetNullability() Nullability       { return s.Nullability }
func (s *PrimitiveType[T]) GetTypeVariationReference() uint32 { return s.TypeVariationRef }
func (s *PrimitiveType[T]) Equals(rhs Type) bool {
	if o, ok := rhs.(*PrimitiveType[T]); ok {
		return *o == *s
	}

	return false
}

func (s *PrimitiveType[T]) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: TypeToProto(s)},
	}
}

func (s *PrimitiveType[T]) String() string {
	var z *T
	if n, ok := primitiveNames[reflect.TypeOf(z)]; ok {
		return n
	}
	return reflect.TypeOf(z).Elem().Name()
}

// create type aliases to the generic structs
type (
	BooleanType      = PrimitiveType[bool]
	Int8Type         = PrimitiveType[int8]
	Int16Type        = PrimitiveType[int16]
	Int32Type        = PrimitiveType[int32]
	Int64Type        = PrimitiveType[int64]
	Float32Type      = PrimitiveType[float32]
	Float64Type      = PrimitiveType[float64]
	StringType       = PrimitiveType[string]
	BinaryType       = PrimitiveType[[]byte]
	TimestampType    = PrimitiveType[Timestamp]
	DateType         = PrimitiveType[Date]
	TimeType         = PrimitiveType[Time]
	TimestampTzType  = PrimitiveType[TimestampTz]
	IntervalYearType = PrimitiveType[IntervalYearToMonth]
	IntervalDayType  = PrimitiveType[IntervalDayToSecond]
	UUIDType         = PrimitiveType[UUID]
	FixedCharType    = FixedLenType[FixedChar]
	VarCharType      = FixedLenType[VarChar]
	FixedBinaryType  = FixedLenType[FixedBinary]
)

// FixedLenType is any of the types which also need to track their specific
// length as they have a fixed length.
type FixedLenType[T FixedChar | VarChar | FixedBinary] struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Length           int32
}

func (*FixedLenType[T]) isRootRef()                          {}
func (s *FixedLenType[T]) GetType() Type                     { return s }
func (s *FixedLenType[T]) GetNullability() Nullability       { return s.Nullability }
func (s *FixedLenType[T]) GetTypeVariationReference() uint32 { return s.TypeVariationRef }
func (s *FixedLenType[T]) Equals(rhs Type) bool {
	if o, ok := rhs.(*FixedLenType[T]); ok {
		return *o == *s
	}

	return false
}

func (s *FixedLenType[T]) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: TypeToProto(s)},
	}
}

func (s *FixedLenType[T]) String() string {
	var z *T
	return fmt.Sprintf("%s<%d>",
		reflect.TypeOf(z).Elem().Name(), s.Length)
}

type DecimalType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Scale, Precision int32
}

func (*DecimalType) isRootRef()                          {}
func (s *DecimalType) GetType() Type                     { return s }
func (s *DecimalType) GetNullability() Nullability       { return s.Nullability }
func (s *DecimalType) GetTypeVariationReference() uint32 { return s.TypeVariationRef }
func (s *DecimalType) Equals(rhs Type) bool {
	if o, ok := rhs.(*DecimalType); ok {
		return *o == *s
	}

	return false
}

func (s *DecimalType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: s.ToProto()},
	}
}

func (t *DecimalType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Decimal_{
		Decimal: &proto.Type_Decimal{
			Scale: t.Scale, Precision: t.Precision,
			Nullability:            t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

func (t *DecimalType) String() string {
	return fmt.Sprintf("decimal<%d, %d>",
		t.Precision, t.Scale)
}

type StructType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Types            []Type
}

func (*StructType) isRootRef()                          {}
func (s *StructType) GetType() Type                     { return s }
func (s *StructType) GetNullability() Nullability       { return s.Nullability }
func (s *StructType) GetTypeVariationReference() uint32 { return s.TypeVariationRef }

func (t *StructType) Equals(rhs Type) bool {
	if b, ok := rhs.(*StructType); ok {
		switch {
		case t.Nullability != b.Nullability:
			return false
		case t.TypeVariationRef != b.TypeVariationRef:
			return false
		case len(t.Types) != len(b.Types):
			return false
		}

		for i := range t.Types {
			if !t.Types[i].Equals(b.Types[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (t *StructType) ToProto() *proto.Type {
	children := make([]*proto.Type, len(t.Types))
	for i, c := range t.Types {
		children[i] = TypeToProto(c)
	}

	return &proto.Type{Kind: &proto.Type_Struct_{
		Struct: &proto.Type_Struct{Types: children,
			TypeVariationReference: t.TypeVariationRef,
			Nullability:            t.Nullability}}}
}

func (t *StructType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (t *StructType) String() string {
	var b strings.Builder
	b.WriteString("struct<")
	for i, f := range t.Types {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(f.String())
	}
	b.WriteByte('>')
	return b.String()
}

type ListType struct {
	Nullability      Nullability
	TypeVariationRef uint32

	Type Type
}

func (*ListType) isRootRef()                          {}
func (s *ListType) GetType() Type                     { return s }
func (s *ListType) GetNullability() Nullability       { return s.Nullability }
func (s *ListType) GetTypeVariationReference() uint32 { return s.TypeVariationRef }

func (t *ListType) Equals(rhs Type) bool {
	if b, ok := rhs.(*ListType); ok {
		switch {
		case t.Nullability != b.Nullability:
			return false
		case t.TypeVariationRef != b.TypeVariationRef:
			return false
		}

		return t.Type.Equals(b.Type)
	}
	return false
}

func (t *ListType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_List_{
		List: &proto.Type_List{Nullability: t.Nullability,
			Type:                   TypeToProto(t.Type),
			TypeVariationReference: t.TypeVariationRef}}}
}

func (t *ListType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (t *ListType) String() string {
	return "list<" + t.Type.String() + ">"
}

type MapType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Key, Value       Type
}

func (*MapType) isRootRef()                          {}
func (s *MapType) GetType() Type                     { return s }
func (s *MapType) GetNullability() Nullability       { return s.Nullability }
func (s *MapType) GetTypeVariationReference() uint32 { return s.TypeVariationRef }

func (t *MapType) Equals(rhs Type) bool {
	if b, ok := rhs.(*MapType); ok {
		switch {
		case t.Nullability != b.Nullability:
			return false
		case t.TypeVariationRef != b.TypeVariationRef:
			return false
		}

		return t.Key.Equals(b.Key) && t.Value.Equals(b.Value)
	}
	return false
}

func (t *MapType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Map_{
		Map: &proto.Type_Map{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef,
			Key:                    TypeToProto(t.Key),
			Value:                  TypeToProto(t.Value)}}}
}

func (t *MapType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (t *MapType) String() string {
	return "map<" + t.Key.String() + " => " + t.Value.String() + ">"
}

// TypeParam represents a type parameter for a user defined type
type TypeParam interface {
	ToProto() *proto.Type_Parameter
	Equals(TypeParam) bool
}

// rather than creating a new one of these for every call ToProto which
// will always be the same empty object we can just create this once
// and return the same one every time.
var nullTypeParam = &proto.Type_Parameter_Null{}

// NullParameter is an explicitly null/unspecified parameter, to select
// the default value (if any).
type NullParameter struct{}

func (NullParameter) Equals(p TypeParam) bool {
	_, ok := p.(NullParameter)
	return ok
}

func (NullParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: nullTypeParam}
}

// DataTypeParameter is like the i32 in LIST<i32>
type DataTypeParameter struct {
	Type
}

func (d *DataTypeParameter) Equals(p TypeParam) bool {
	if dt, ok := p.(*DataTypeParameter); ok {
		return d.Type.Equals(dt.Type)
	}
	return false
}

func (d *DataTypeParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_DataType{
		DataType: TypeToProto(d.Type)}}
}

// BooleanParameter is a type parameter like <true> for a type.
type BooleanParameter bool

func (b BooleanParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(BooleanParameter); ok {
		return b == rhs
	}
	return false
}

func (b BooleanParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Boolean{
		Boolean: bool(b)}}
}

// IntegerParameter is the type parameter like 10 in VARCHAR<10>
type IntegerParameter int64

func (b IntegerParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(IntegerParameter); ok {
		return b == rhs
	}
	return false
}

func (p IntegerParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Integer{
		Integer: int64(p)}}
}

// EnumParameter is a type parameter that is some enum value
type EnumParameter string

func (b EnumParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(EnumParameter); ok {
		return b == rhs
	}
	return false
}

func (p EnumParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Enum{
		Enum: string(p)}}
}

// StringParameter is a type parameter which is a string value
type StringParameter string

func (b StringParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(StringParameter); ok {
		return b == rhs
	}
	return false
}

func (p StringParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_String_{
		String_: string(p)}}
}

// TypeParamFromProto converts a protobuf Type_Parameter message to
// a TypeParam object for processing.
func TypeParamFromProto(p *proto.Type_Parameter) TypeParam {
	switch p := p.Parameter.(type) {
	case *proto.Type_Parameter_Null:
		return NullParameter{}
	case *proto.Type_Parameter_Boolean:
		return BooleanParameter(p.Boolean)
	case *proto.Type_Parameter_DataType:
		return &DataTypeParameter{TypeFromProto(p.DataType)}
	case *proto.Type_Parameter_Integer:
		return IntegerParameter(p.Integer)
	case *proto.Type_Parameter_Enum:
		return EnumParameter(p.Enum)
	case *proto.Type_Parameter_String_:
		return StringParameter(p.String_)
	}
	return nil
}

type UserDefinedType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	TypeReference    uint32
	TypeParameters   []TypeParam
}

func (*UserDefinedType) isRootRef()                          {}
func (s *UserDefinedType) GetType() Type                     { return s }
func (s *UserDefinedType) GetNullability() Nullability       { return s.Nullability }
func (s *UserDefinedType) GetTypeVariationReference() uint32 { return s.TypeVariationRef }

func (t *UserDefinedType) Equals(rhs Type) bool {
	if other, ok := rhs.(*UserDefinedType); ok {
		switch {
		case t.Nullability != other.Nullability:
			return false
		case t.TypeVariationRef != other.TypeVariationRef:
			return false
		case t.TypeReference != other.TypeReference:
			return false
		case len(t.TypeParameters) != len(other.TypeParameters):
			return false
		}

		for i := range t.TypeParameters {
			if !t.TypeParameters[i].Equals(other.TypeParameters[i]) {
				return false
			}
		}
		return true
	}

	return false
}

func (t *UserDefinedType) ToProto() *proto.Type {
	params := make([]*proto.Type_Parameter, len(t.TypeParameters))
	for i, p := range t.TypeParameters {
		params[i] = p.ToProto()
	}

	return &proto.Type{Kind: &proto.Type_UserDefined_{
		UserDefined: &proto.Type_UserDefined{
			Nullability:            t.Nullability,
			TypeVariationReference: t.TypeVariationRef,
			TypeReference:          t.TypeReference,
			TypeParameters:         params,
		}}}
}

func (t *UserDefinedType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (t *UserDefinedType) String() string {
	return "user_defined_type"
}

func (e Enum) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Enum{Enum: string(e)},
	}
}

func (e Enum) String() string { return string(e) }

type NamedStruct struct {
	Names  []string
	Struct StructType
}

func (n *NamedStruct) String() string {
	var b strings.Builder

	// names are in depth-first order
	nameIdx := 0

	var writeType func(t Type)

	writeType = func(t Type) {
		switch t := t.(type) {
		case *StructType:
			b.WriteString("struct<")
			for i, c := range t.Types {
				if i != 0 {
					b.WriteString(", ")
				}
				b.WriteString(n.Names[nameIdx])
				nameIdx++
				b.WriteString(": ")
				writeType(c)
			}
			b.WriteString(">")
		case *MapType:
			b.WriteString("map<")
			writeType(t.Key)
			b.WriteString(",")
			writeType(t.Value)
			b.WriteString(">")
		case *ListType:
			b.WriteString("list<")
			writeType(t.Type)
			b.WriteString(">")
		default:
			b.WriteString(t.String())
		}
	}

	for _, t := range n.Struct.Types {
		b.WriteString("- ")
		b.WriteString(n.Names[nameIdx])
		b.WriteString(": ")
		nameIdx++

		writeType(t)
		b.WriteString("\n")
	}

	return b.String()
}
