// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
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

	FuncArg interface {
		isFuncArg()
	}

	SortKind interface {
		isSortKind()
	}

	funcArg struct{}

	Type interface {
		FuncArg
		GetNullability() Nullability
		GetTypeVariationReference() uint32
		Equals(Type) bool
	}
)

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

type primitiveType[T PrimitiveLiteralValue | []byte | IntervalYearToMonth | IntervalDayToSecond | UUID] struct {
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32
}

func (s *primitiveType[T]) GetNullability() Nullability       { return s.Nullability }
func (s *primitiveType[T]) GetTypeVariationReference() uint32 { return s.TypeVariationRef }
func (s *primitiveType[T]) Equals(rhs Type) bool {
	if o, ok := rhs.(*primitiveType[T]); ok {
		return *o == *s
	}

	return false
}

type (
	BooleanType      = primitiveType[bool]
	Int8Type         = primitiveType[int8]
	Int16Type        = primitiveType[int16]
	Int32Type        = primitiveType[int32]
	Int64Type        = primitiveType[int64]
	Float32Type      = primitiveType[float32]
	Float64Type      = primitiveType[float64]
	StringType       = primitiveType[string]
	BinaryType       = primitiveType[[]byte]
	TimestampType    = primitiveType[Timestamp]
	DateType         = primitiveType[Date]
	TimeType         = primitiveType[Time]
	TimestampTzType  = primitiveType[TimestampTz]
	IntervalYearType = primitiveType[IntervalYearToMonth]
	IntervalDayType  = primitiveType[IntervalDayToSecond]
	UUIDType         = primitiveType[UUID]
	FixedCharType    = fixedLenType[FixedChar]
	VarCharType      = fixedLenType[VarChar]
	FixedBinaryType  = fixedLenType[FixedBinary]
)

type fixedLenType[T FixedChar | VarChar | FixedBinary] struct {
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32
	Length           int32
}

func (s *fixedLenType[T]) GetNullability() Nullability       { return s.Nullability }
func (s *fixedLenType[T]) GetTypeVariationReference() uint32 { return s.TypeVariationRef }
func (s *fixedLenType[T]) Equals(rhs Type) bool {
	if o, ok := rhs.(*fixedLenType[T]); ok {
		return *o == *s
	}

	return false
}

type DecimalType struct {
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32
	Scale, Precision int32
}

func (s *DecimalType) GetNullability() Nullability       { return s.Nullability }
func (s *DecimalType) GetTypeVariationReference() uint32 { return s.TypeVariationRef }
func (s *DecimalType) Equals(rhs Type) bool {
	if o, ok := rhs.(*DecimalType); ok {
		return *o == *s
	}

	return false
}

func (t *DecimalType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Decimal_{
		Decimal: &proto.Type_Decimal{
			Scale: t.Scale, Precision: t.Precision,
			Nullability:            t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type StructType struct {
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32
	Types            []Type
}

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

type ListType struct {
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32

	Type Type
}

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

type MapType struct {
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32
	Key, Value       Type
}

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

type TypeParam interface {
	ToProto() *proto.Type_Parameter
	Equals(TypeParam) bool
}

var nullTypeParam = &proto.Type_Parameter_Null{}

type NullParameter struct{}

func (NullParameter) Equals(p TypeParam) bool {
	_, ok := p.(NullParameter)
	return ok
}

func (NullParameter) ToProto() *proto.Type_Parameter {
	return &proto.Type_Parameter{Parameter: nullTypeParam}
}

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
	funcArg

	Nullability      Nullability
	TypeVariationRef uint32
	TypeReference    uint32
	TypeParameters   []TypeParam
}

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

func (funcArg) isFuncArg() {}

func (Enum) isFuncArg() {}
