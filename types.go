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

var TypeCreator = struct {
	Bool         func(Nullability) *BooleanType
	I8           func(Nullability) *Int8Type
	I16          func(Nullability) *Int16Type
	I32          func(Nullability) *Int32Type
	I64          func(Nullability) *Int64Type
	Float32      func(Nullability) *Float32Type
	Float64      func(Nullability) *Float64Type
	String       func(Nullability) *StringType
	Binary       func(Nullability) *BinaryType
	Timestamp    func(Nullability) *TimestampType
	Date         func(Nullability) *DateType
	Time         func(Nullability) *TimeType
	TimestampTz  func(Nullability) *TimestampTzType
	IntervalYear func(Nullability) *IntervalYearType
	IntervalDay  func(Nullability) *IntervalDayType
	UUID         func(Nullability) *UUIDType
	FixedChar    func(length int32, n Nullability) *FixedCharType
	VarChar      func(length int32, n Nullability) *VarCharType
	FixedBinary  func(length int32, n Nullability) *FixedBinaryType
	Decimal      func(scale, prec int32, n Nullability) *DecimalType
	Struct       func([]Type, Nullability) *StructType
	List         func(Type, Nullability) *ListType
	Map          func(Key, Value Type, n Nullability) *MapType
}{}

func TypeFromProto(t *proto.Type) Type {
	switch t := t.Kind.(type) {
	case *proto.Type_Bool:
		ret := TypeCreator.Bool(t.Bool.Nullability)
		ret.TypeVariationRef = t.Bool.TypeVariationReference
		return ret
	case *proto.Type_I8_:
		ret := TypeCreator.I8(t.I8.Nullability)
		ret.TypeVariationRef = t.I8.TypeVariationReference
		return ret
	case *proto.Type_I16_:
		ret := TypeCreator.I16(t.I16.Nullability)
		ret.TypeVariationRef = t.I16.TypeVariationReference
		return ret
	case *proto.Type_I32_:
		ret := TypeCreator.I32(t.I32.Nullability)
		ret.TypeVariationRef = t.I32.TypeVariationReference
		return ret
	case *proto.Type_I64_:
		ret := TypeCreator.I64(t.I64.Nullability)
		ret.TypeVariationRef = t.I64.TypeVariationReference
		return ret
	case *proto.Type_Fp32:
		ret := TypeCreator.Float32(t.Fp32.Nullability)
		ret.TypeVariationRef = t.Fp32.TypeVariationReference
		return ret
	case *proto.Type_Fp64:
		ret := TypeCreator.Float64(t.Fp64.Nullability)
		ret.TypeVariationRef = t.Fp64.TypeVariationReference
		return ret
	case *proto.Type_String_:
		ret := TypeCreator.String(t.String_.Nullability)
		ret.TypeVariationRef = t.String_.TypeVariationReference
		return ret
	case *proto.Type_Binary_:
		ret := TypeCreator.Binary(t.Binary.Nullability)
		ret.TypeVariationRef = t.Binary.TypeVariationReference
		return ret
	case *proto.Type_Timestamp_:
		ret := TypeCreator.Timestamp(t.Timestamp.Nullability)
		ret.TypeVariationRef = t.Timestamp.TypeVariationReference
		return ret
	case *proto.Type_Date_:
		ret := TypeCreator.Date(t.Date.Nullability)
		ret.TypeVariationRef = t.Date.TypeVariationReference
		return ret
	case *proto.Type_Time_:
		ret := TypeCreator.Time(t.Time.Nullability)
		ret.TypeVariationRef = t.Time.TypeVariationReference
		return ret
	case *proto.Type_IntervalYear_:
		ret := TypeCreator.IntervalYear(t.IntervalYear.Nullability)
		ret.TypeVariationRef = t.IntervalYear.TypeVariationReference
		return ret
	case *proto.Type_IntervalDay_:
		ret := TypeCreator.IntervalDay(t.IntervalDay.Nullability)
		ret.TypeVariationRef = t.IntervalDay.TypeVariationReference
		return ret
	case *proto.Type_TimestampTz:
		ret := TypeCreator.TimestampTz(t.TimestampTz.Nullability)
		ret.TypeVariationRef = t.TimestampTz.TypeVariationReference
		return ret
	case *proto.Type_Uuid:
		ret := TypeCreator.UUID(t.Uuid.Nullability)
		ret.TypeVariationRef = t.Uuid.TypeVariationReference
		return ret
	case *proto.Type_FixedBinary_:
		ret := TypeCreator.FixedBinary(t.FixedBinary.Length, t.FixedBinary.Nullability)
		ret.TypeVariationRef = t.FixedBinary.TypeVariationReference
		return ret
	case *proto.Type_FixedChar_:
		ret := TypeCreator.FixedChar(t.FixedChar.Length, t.FixedChar.Nullability)
		ret.TypeVariationRef = t.FixedChar.TypeVariationReference
		return ret
	case *proto.Type_Decimal_:
		ret := TypeCreator.Decimal(t.Decimal.Scale, t.Decimal.Precision, t.Decimal.Nullability)
		ret.TypeVariationRef = t.Decimal.TypeVariationReference
		return ret
	case *proto.Type_Struct_:
		fields := make([]Type, len(t.Struct.Types))
		for i, f := range t.Struct.Types {
			fields[i] = TypeFromProto(f)
		}
		ret := TypeCreator.Struct(fields, t.Struct.Nullability)
		ret.TypeVariationRef = t.Struct.TypeVariationReference
		return ret
	case *proto.Type_List_:
		ret := TypeCreator.List(TypeFromProto(t.List.Type), t.List.Nullability)
		ret.TypeVariationRef = t.List.TypeVariationReference
		return ret
	case *proto.Type_Map_:
		ret := TypeCreator.Map(TypeFromProto(t.Map.Key), TypeFromProto(t.Map.Value), t.Map.Nullability)
		ret.TypeVariationRef = t.Map.TypeVariationReference
		return ret
	case *proto.Type_UserDefined_:
		params := make([]TypeParam, len(t.UserDefined.TypeParameters))
		for i, p := range t.UserDefined.TypeParameters {
			params[i] = TypeParamFromProto(p)
		}
		return &UserDefinedType{
			baseType: baseType{
				Nullability:      t.UserDefined.Nullability,
				TypeVariationRef: t.UserDefined.TypeVariationReference,
			},
			TypeReference:  t.UserDefined.TypeReference,
			TypeParameters: params,
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
		ToProto() *proto.Type
		GetNullability() Nullability
		GetTypeVariationReference() uint32
		Equals(Type) bool
	}
)

type baseType struct {
	funcArg
	Nullability      Nullability
	TypeVariationRef uint32
}

func (b *baseType) GetNullability() Nullability       { return b.Nullability }
func (b *baseType) GetTypeVariationReference() uint32 { return b.TypeVariationRef }

type BooleanType struct {
	baseType
}

func (t *BooleanType) Equals(rhs Type) bool {
	if b, ok := rhs.(*BooleanType); ok {
		return *t == *b
	}
	return false
}

func (t *BooleanType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Bool{
		Bool: &proto.Type_Boolean{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type Int8Type struct {
	baseType
}

func (t *Int8Type) Equals(rhs Type) bool {
	if b, ok := rhs.(*Int8Type); ok {
		return *t == *b
	}
	return false
}

func (t *Int8Type) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_I8_{
		I8: &proto.Type_I8{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type Int16Type struct {
	baseType
}

func (t *Int16Type) Equals(rhs Type) bool {
	if b, ok := rhs.(*Int16Type); ok {
		return *t == *b
	}
	return false
}

func (t *Int16Type) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_I16_{
		I16: &proto.Type_I16{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type Int32Type struct {
	baseType
}

func (t *Int32Type) Equals(rhs Type) bool {
	if b, ok := rhs.(*Int32Type); ok {
		return *t == *b
	}
	return false
}

func (t *Int32Type) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_I32_{
		I32: &proto.Type_I32{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type Int64Type struct {
	baseType
}

func (t *Int64Type) Equals(rhs Type) bool {
	if b, ok := rhs.(*Int64Type); ok {
		return *t == *b
	}
	return false
}

func (t *Int64Type) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_I64_{
		I64: &proto.Type_I64{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type Float32Type struct {
	baseType
}

func (t *Float32Type) Equals(rhs Type) bool {
	if b, ok := rhs.(*Float32Type); ok {
		return *t == *b
	}
	return false
}

func (t *Float32Type) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Fp32{
		Fp32: &proto.Type_FP32{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type Float64Type struct {
	baseType
}

func (t *Float64Type) Equals(rhs Type) bool {
	if b, ok := rhs.(*Float64Type); ok {
		return *t == *b
	}
	return false
}

func (t *Float64Type) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Fp64{
		Fp64: &proto.Type_FP64{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type StringType struct {
	baseType
}

func (t *StringType) Equals(rhs Type) bool {
	if b, ok := rhs.(*StringType); ok {
		return *t == *b
	}
	return false
}

func (t *StringType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_String_{
		String_: &proto.Type_String{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type BinaryType struct {
	baseType
}

func (t *BinaryType) Equals(rhs Type) bool {
	if b, ok := rhs.(*BinaryType); ok {
		return *t == *b
	}
	return false
}

func (t *BinaryType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Binary_{
		Binary: &proto.Type_Binary{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type TimestampType struct {
	baseType
}

func (t *TimestampType) Equals(rhs Type) bool {
	if b, ok := rhs.(*TimestampType); ok {
		return *t == *b
	}
	return false
}

func (t *TimestampType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Timestamp_{
		Timestamp: &proto.Type_Timestamp{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type DateType struct {
	baseType
}

func (t *DateType) Equals(rhs Type) bool {
	if b, ok := rhs.(*DateType); ok {
		return *t == *b
	}
	return false
}

func (t *DateType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Date_{
		Date: &proto.Type_Date{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type TimeType struct {
	baseType
}

func (t *TimeType) Equals(rhs Type) bool {
	if b, ok := rhs.(*TimeType); ok {
		return *t == *b
	}
	return false
}

func (t *TimeType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Time_{
		Time: &proto.Type_Time{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type TimestampTzType struct {
	baseType
}

func (t *TimestampTzType) Equals(rhs Type) bool {
	if b, ok := rhs.(*TimestampTzType); ok {
		return *t == *b
	}
	return false
}

func (t *TimestampTzType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_TimestampTz{
		TimestampTz: &proto.Type_TimestampTZ{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type IntervalYearType struct {
	baseType
}

func (t *IntervalYearType) Equals(rhs Type) bool {
	if b, ok := rhs.(*IntervalYearType); ok {
		return *t == *b
	}
	return false
}

func (t *IntervalYearType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_IntervalYear_{
		IntervalYear: &proto.Type_IntervalYear{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type IntervalDayType struct {
	baseType
}

func (t *IntervalDayType) Equals(rhs Type) bool {
	if b, ok := rhs.(*IntervalDayType); ok {
		return *t == *b
	}
	return false
}

func (t *IntervalDayType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_IntervalDay_{
		IntervalDay: &proto.Type_IntervalDay{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type UUIDType struct {
	baseType
}

func (t *UUIDType) Equals(rhs Type) bool {
	if b, ok := rhs.(*UUIDType); ok {
		return *t == *b
	}
	return false
}

func (t *UUIDType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Uuid{
		Uuid: &proto.Type_UUID{Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type FixedCharType struct {
	baseType
	Length int32
}

func (t *FixedCharType) Equals(rhs Type) bool {
	if b, ok := rhs.(*FixedCharType); ok {
		return *t == *b
	}
	return false
}

func (t *FixedCharType) Len() int32 { return t.Length }

func (t *FixedCharType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_FixedChar_{
		FixedChar: &proto.Type_FixedChar{Length: t.Length, Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type VarCharType struct {
	baseType
	Length int32
}

func (t *VarCharType) Equals(rhs Type) bool {
	if b, ok := rhs.(*VarCharType); ok {
		return *t == *b
	}
	return false
}

func (t *VarCharType) Len() int32 { return t.Length }

func (t *VarCharType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_Varchar{
		Varchar: &proto.Type_VarChar{Length: t.Length, Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type FixedBinaryType struct {
	baseType
	Length int32
}

func (t *FixedBinaryType) Equals(rhs Type) bool {
	if b, ok := rhs.(*FixedBinaryType); ok {
		return *t == *b
	}
	return false
}

func (t *FixedBinaryType) Len() int32 { return t.Length }

func (t *FixedBinaryType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_FixedBinary_{
		FixedBinary: &proto.Type_FixedBinary{Length: t.Length, Nullability: t.Nullability,
			TypeVariationReference: t.TypeVariationRef}}}
}

type DecimalType struct {
	baseType
	Scale, Precision int32
}

func (t *DecimalType) Equals(rhs Type) bool {
	if b, ok := rhs.(*DecimalType); ok {
		return *t == *b
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
	baseType
	Types []Type
}

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
		children[i] = c.ToProto()
	}

	return &proto.Type{Kind: &proto.Type_Struct_{
		Struct: &proto.Type_Struct{Types: children,
			TypeVariationReference: t.TypeVariationRef,
			Nullability:            t.Nullability}}}
}

type ListType struct {
	baseType

	Type Type
}

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
			Type:                   t.Type.ToProto(),
			TypeVariationReference: t.TypeVariationRef}}}
}

type MapType struct {
	baseType

	Key, Value Type
}

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
			Key:                    t.Key.ToProto(),
			Value:                  t.Value.ToProto()}}}
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
		DataType: d.Type.ToProto()}}
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
	baseType
	TypeReference  uint32
	TypeParameters []TypeParam
}

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
