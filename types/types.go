// SPDX-License-Identifier: Apache-2.0

//lint:file-ignore SA1019 Using a deprecated function, variable, constant or field

package types

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	substraitgo "github.com/substrait-io/substrait-go/v9"
)

// Version is the Substrait version a plan or extended expression was built against.
type Version struct {
	MajorNumber uint32
	MinorNumber uint32
	PatchNumber uint32
	GitHash     string
	Producer    string
}

// String reports a readable version like "0.29.0+abc123 (producer)".
func (v Version) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%d.%d.%d", v.MajorNumber, v.MinorNumber, v.PatchNumber)
	if v.GitHash != "" {
		b.WriteString("+" + v.GitHash)
	}
	if v.Producer != "" {
		b.WriteString(" (" + v.Producer + ")")
	}
	return b.String()
}

// FunctionOption is a named function behavior option: its name and the producer's ordered list of
// acceptable preference values, mirroring the fields of the Substrait FunctionOption message.
type FunctionOption struct {
	Name       string
	Preference []string
}

// Nullability indicates whether values of a Substrait type may be null.
type Nullability int32

const (
	NullabilityUnspecified Nullability = 0
	NullabilityNullable    Nullability = 1
	NullabilityRequired    Nullability = 2
)

// String returns the protobuf enum name for the nullability value.
func (n Nullability) String() string {
	switch n {
	case NullabilityUnspecified:
		return "NULLABILITY_UNSPECIFIED"
	case NullabilityNullable:
		return "NULLABILITY_NULLABLE"
	case NullabilityRequired:
		return "NULLABILITY_REQUIRED"
	default:
		return strconv.Itoa(int(n))
	}
}

type TypeName string

const (
	TypeNameI8               TypeName = "i8"
	TypeNameI16              TypeName = "i16"
	TypeNameI32              TypeName = "i32"
	TypeNameI64              TypeName = "i64"
	TypeNameFp32             TypeName = "fp32"
	TypeNameFp64             TypeName = "fp64"
	TypeNameString           TypeName = "string"
	TypeNameBinary           TypeName = "binary"
	TypeNameBoolean          TypeName = "boolean"
	TypeNameDate             TypeName = "date"
	TypeNameTime             TypeName = "time"
	TypeNameTimestamp        TypeName = "timestamp"
	TypeNameTimestampTz      TypeName = "timestamp_tz"
	TypeNameIntervalYear     TypeName = "interval_year"
	TypeNameIntervalDay      TypeName = "interval_day"
	TypeNameIntervalCompound TypeName = "interval_compound"
	TypeNameUUID             TypeName = "uuid"
	TypeNameUDT              TypeName = "u!"

	TypeNameFixedBinary          TypeName = "fixedbinary"
	TypeNameFixedChar            TypeName = "fixedchar"
	TypeNameVarChar              TypeName = "varchar"
	TypeNameDecimal              TypeName = "decimal"
	TypeNamePrecisionTime        TypeName = "precision_time"
	TypeNamePrecisionTimestamp   TypeName = "precision_timestamp"
	TypeNamePrecisionTimestampTz TypeName = "precision_timestamp_tz"
)

var simpleTypeNameMap = map[TypeName]Type{
	TypeNameI8:           &Int8Type{},
	TypeNameI16:          &Int16Type{},
	TypeNameI32:          &Int32Type{},
	TypeNameI64:          &Int64Type{},
	TypeNameFp32:         &Float32Type{},
	TypeNameFp64:         &Float64Type{},
	TypeNameString:       &StringType{},
	TypeNameBinary:       &BinaryType{},
	TypeNameBoolean:      &BooleanType{},
	TypeNameDate:         &DateType{},
	TypeNameTime:         &TimeType{},
	TypeNameTimestamp:    &TimestampType{},
	TypeNameTimestampTz:  &TimestampTzType{},
	TypeNameIntervalYear: &IntervalYearType{},
	TypeNameUUID:         &UUIDType{},
	TypeNameUDT:          &UserDefinedType{},
}

var fixedTypeNameMap = map[TypeName]FixedType{
	TypeNameFixedBinary: &FixedBinaryType{},
	TypeNameFixedChar:   &FixedCharType{},
	TypeNameVarChar:     &VarCharType{},
}

var shortTypeNames = map[TypeName]string{
	TypeNameString:           "str",
	TypeNameBinary:           "vbin",
	TypeNameBoolean:          "bool",
	TypeNameTimestamp:        "ts",
	TypeNameTimestampTz:      "tstz",
	TypeNameIntervalYear:     "iyear",
	TypeNameIntervalDay:      "iday",
	TypeNameIntervalCompound: "icompound",

	TypeNameFixedBinary: "fbin",
	TypeNameFixedChar:   "fchar",
	TypeNameVarChar:     "vchar",

	TypeNameDecimal:              "dec",
	TypeNamePrecisionTime:        "pt",
	TypeNamePrecisionTimestamp:   "pts",
	TypeNamePrecisionTimestampTz: "ptstz",
}

func GetShortTypeName(name TypeName) string {
	if n, ok := shortTypeNames[name]; ok {
		return n
	}
	return string(name)
}

func SimpleTypeNameToType(name TypeName) (Type, error) {
	if t, ok := simpleTypeNameMap[name]; ok {
		return t, nil
	}
	return nil, substraitgo.ErrNotFound
}

func FixedTypeNameToType(name TypeName) (FixedType, error) {
	if t, ok := fixedTypeNameMap[name]; ok {
		return t, nil
	}
	return nil, substraitgo.ErrInvalidType
}

func GetTypeNameToTypeMap() map[string]Type {
	typeMap := make(map[string]Type)
	for k, v := range simpleTypeNameMap {
		typeMap[string(k)] = v
	}
	for k, v := range fixedTypeNameMap {
		typeMap[string(k)] = v
	}
	typeMap[string(TypeNameDecimal)] = &DecimalType{}
	typeMap[string(TypeNameIntervalDay)] = &IntervalDayType{}
	typeMap[string(TypeNamePrecisionTime)] = &PrecisionTimeType{}
	typeMap[string(TypeNamePrecisionTimestamp)] = &PrecisionTimestampType{}
	typeMap[string(TypeNamePrecisionTimestampTz)] = &PrecisionTimestampTzType{}
	return typeMap
}

// AggregationPhase describes which part of an aggregation or window function to
// perform within the context of distributed algorithms.
type AggregationPhase int32

const (
	AggregationPhaseUnspecified                AggregationPhase = 0
	AggregationPhaseInitialToIntermediate      AggregationPhase = 1
	AggregationPhaseIntermediateToIntermediate AggregationPhase = 2
	AggregationPhaseInitialToResult            AggregationPhase = 3
	AggregationPhaseIntermediateToResult       AggregationPhase = 4
)

// String returns the protobuf enum name for the aggregation phase.
func (p AggregationPhase) String() string {
	switch p {
	case AggregationPhaseUnspecified:
		return "AGGREGATION_PHASE_UNSPECIFIED"
	case AggregationPhaseInitialToIntermediate:
		return "AGGREGATION_PHASE_INITIAL_TO_INTERMEDIATE"
	case AggregationPhaseIntermediateToIntermediate:
		return "AGGREGATION_PHASE_INTERMEDIATE_TO_INTERMEDIATE"
	case AggregationPhaseInitialToResult:
		return "AGGREGATION_PHASE_INITIAL_TO_RESULT"
	case AggregationPhaseIntermediateToResult:
		return "AGGREGATION_PHASE_INTERMEDIATE_TO_RESULT"
	default:
		return strconv.Itoa(int(p))
	}
}

// AggregationInvocation describes the method in which equivalent records are
// merged before being aggregated.
type AggregationInvocation int32

const (
	AggregationInvocationUnspecified AggregationInvocation = 0
	AggregationInvocationAll         AggregationInvocation = 1
	AggregationInvocationDistinct    AggregationInvocation = 2
)

// String returns the protobuf enum name for the aggregation invocation.
func (i AggregationInvocation) String() string {
	switch i {
	case AggregationInvocationUnspecified:
		return "AGGREGATION_INVOCATION_UNSPECIFIED"
	case AggregationInvocationAll:
		return "AGGREGATION_INVOCATION_ALL"
	case AggregationInvocationDistinct:
		return "AGGREGATION_INVOCATION_DISTINCT"
	default:
		return strconv.Itoa(int(i))
	}
}

// BoundsType indicates whether a window frame's bounds are measured in rows or in a range of values.
type BoundsType int32

const (
	BoundsTypeUnspecified BoundsType = 0
	BoundsTypeRows        BoundsType = 1
	BoundsTypeRange       BoundsType = 2
)

// String returns the protobuf enum name for the bounds type.
func (b BoundsType) String() string {
	switch b {
	case BoundsTypeUnspecified:
		return "BOUNDS_TYPE_UNSPECIFIED"
	case BoundsTypeRows:
		return "BOUNDS_TYPE_ROWS"
	case BoundsTypeRange:
		return "BOUNDS_TYPE_RANGE"
	default:
		return strconv.Itoa(int(b))
	}
}

type SortDirection int32

const (
	SortUnspecified    SortDirection = 0
	SortAscNullsFirst  SortDirection = 1
	SortAscNullsLast   SortDirection = 2
	SortDescNullsFirst SortDirection = 3
	SortDescNullsLast  SortDirection = 4
	SortClustered      SortDirection = 5
)

// String returns the protobuf enum name for the sort direction.
func (s SortDirection) String() string {
	switch s {
	case SortUnspecified:
		return "SORT_DIRECTION_UNSPECIFIED"
	case SortAscNullsFirst:
		return "SORT_DIRECTION_ASC_NULLS_FIRST"
	case SortAscNullsLast:
		return "SORT_DIRECTION_ASC_NULLS_LAST"
	case SortDescNullsFirst:
		return "SORT_DIRECTION_DESC_NULLS_FIRST"
	case SortDescNullsLast:
		return "SORT_DIRECTION_DESC_NULLS_LAST"
	case SortClustered:
		return "SORT_DIRECTION_CLUSTERED"
	default:
		return strconv.Itoa(int(s))
	}
}

func (SortDirection) isSortKind() {}

type FunctionRef uint32

func (f FunctionRef) String() string { return "comparison_func_ref: " + strconv.Itoa(int(f)) }

func (FunctionRef) isSortKind() {}

// CastFailBehavior indicates how a cast behaves when the input can't be converted to the target type.
type CastFailBehavior int32

const (
	CastFailBehaviorUnspecified    CastFailBehavior = 0
	CastFailBehaviorReturnNull     CastFailBehavior = 1
	CastFailBehaviorThrowException CastFailBehavior = 2
)

// String returns the protobuf enum name for the cast failure behavior.
func (b CastFailBehavior) String() string {
	switch b {
	case CastFailBehaviorUnspecified:
		return "FAILURE_BEHAVIOR_UNSPECIFIED"
	case CastFailBehaviorReturnNull:
		return "FAILURE_BEHAVIOR_RETURN_NULL"
	case CastFailBehaviorThrowException:
		return "FAILURE_BEHAVIOR_THROW_EXCEPTION"
	default:
		return strconv.Itoa(int(b))
	}
}

// VarChar is a variable-length character literal: its value and length, mirroring the fields of
// the Substrait VarChar literal message.
type VarChar struct {
	Value  string
	Length uint32
}

// Decimal is a decimal literal: the value as a 16-byte little-endian two's-complement integer
// (ignoring precision), together with the precision and scale.
type Decimal struct {
	Value     []byte
	Precision int32
	Scale     int32
}

// IntervalYearToMonth is an interval literal expressed in years and months, mirroring the
// fields of the Substrait IntervalYearToMonth literal message.
type IntervalYearToMonth struct {
	Years  int32
	Months int32
}

// PrecisionTime is a time-of-day literal: the number of precision units past
// midnight, mirroring the fields of the Substrait PrecisionTime literal message.
type PrecisionTime struct {
	Precision TimePrecision
	Value     int64
}

// PrecisionTimestamp is a timestamp literal in an unspecified time zone: the
// number of precision units since the UNIX epoch, mirroring the fields of the
// Substrait PrecisionTimestamp literal message.
type PrecisionTimestamp struct {
	Precision TimePrecision
	Value     int64
}

// PrecisionTimestampTz is a UTC timestamp literal: the number of precision units
// since the UNIX epoch, mirroring the fields of the Substrait PrecisionTimestamp
// literal message that backs the precision_timestamp_tz field.
type PrecisionTimestampTz struct {
	Precision TimePrecision
	Value     int64
}

// IntervalDayToSecond is an interval literal spanning days down to sub-seconds, mirroring the
// fields of the Substrait IntervalDayToSecond literal message. Its sub-second value is Subseconds
// interpreted at Precision fractional-second digits.
type IntervalDayToSecond struct {
	Days       int32
	Seconds    int32
	Subseconds int64
	Precision  TimePrecision
}

// GetPrecisionProtoVal returns the sub-second precision as its protobuf value, and 0 for a nil
// receiver.
func (i *IntervalDayToSecond) GetPrecisionProtoVal() int32 {
	if i == nil {
		return 0
	}
	return i.Precision.ToProtoVal()
}

// IntervalDayToSecondToProto encodes the domain interval as its protobuf literal message. It always
// writes the precision precision_mode arm; the deprecated microseconds arm is never emitted.
func IntervalDayToSecondToProto(v *IntervalDayToSecond) *proto.Expression_Literal_IntervalDayToSecond {
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
func IntervalDayToSecondFromProto(p *proto.Expression_Literal_IntervalDayToSecond) (*IntervalDayToSecond, error) {
	if p == nil {
		return nil, nil
	}
	v := &IntervalDayToSecond{
		Days:       p.GetDays(),
		Seconds:    p.GetSeconds(),
		Subseconds: p.GetSubseconds(),
	}
	switch m := p.PrecisionMode.(type) {
	case *proto.Expression_Literal_IntervalDayToSecond_Precision:
		v.Precision = TimePrecision(m.Precision)
	case *proto.Expression_Literal_IntervalDayToSecond_Microseconds:
		v.Subseconds = int64(m.Microseconds)
		v.Precision = PrecisionMicroSeconds
	default:
		return nil, errors.New("interval day to second literal is missing its precision_mode")
	}
	return v, nil
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

	// FuncArg corresponds to the protobuf FunctionArgument. Anything
	// which could be a function argument should meet this interface.
	// This is either an Expression, a Type, or an Enum (string).
	// These are the actual arguments for a function present in a plan.
	FuncArg interface {
		fmt.Stringer
		ToProtoFuncArg() *proto.FunctionArgument
	}

	SortKind interface {
		isSortKind()
		fmt.Stringer
	}

	// Type corresponds to the proto.Type message and represents
	// a specific concrete type. These are types which can be present in plan (are serializable)
	Type interface {
		FuncArg
		isRootRef()
		fmt.Stringer
		ShortString() string
		GetType() Type
		GetNullability() Nullability
		GetTypeVariationReference() uint32
		Equals(Type) bool
		// WithNullability returns a copy of this type but with
		// the nullability set to the passed in value
		WithNullability(Nullability) Type
		// GetParameters returns all parameters of this type and will be used in function return type derivation
		GetParameters() []interface{}
	}

	TimeConverter interface {
		// ToTime converts the current value into a time.Time assuming microsecond precision.
		ToTime() time.Time
		// ToPrecisionTime converts the current value using the provided precision into a time.Time.
		ToPrecisionTime(precision TimePrecision) time.Time
	}

	TimePrinter interface {
		// ToTimeString returns a human consumable version of the current value.
		ToTimeString() string
	}

	IsoTimePrinter interface {
		// ToIsoTimeString returns a human consumable version of the current value in ISO8601 format.
		ToIsoTimeString() string
	}

	IsoValuePrinter interface {
		// IsoValueString API returns the value in ISO8601 format. This is used in normalizing the function testcases
		IsoValueString() string
	}

	// CompositeType this represents a concrete type having components
	CompositeType interface {
		Type
		// ParameterString this returns parameter string
		// for e.g. parameter decimal<P, S>, ParameterString returns "P,S"
		ParameterString() string
		// BaseString this returns long name for parameter string
		// for e.g. parameter decimal<P, S>, BaseString returns "decimal"
		BaseString() string
	}

	// FuncDefArgType represents a type used in a function argument
	// This is an unresolved type of an argument present in the function definition.
	// This type can't be present in plan (not serializable)
	FuncDefArgType interface {
		fmt.Stringer
		//SetNullability set nullability as given argument
		SetNullability(Nullability) FuncDefArgType
		// HasParameterizedParam returns true if the type has at least one parameterized parameters
		// if all parameters are concrete then it returns false
		HasParameterizedParam() bool
		// GetParameterizedParams returns all parameterized parameters
		// it doesn't return concrete parameters
		GetParameterizedParams() []interface{}

		// MatchWithNullability This API return true if Type argument
		// is compatible with this param otherwise it returns false.
		// This method expects that nullability of argument is same as this type.
		MatchWithNullability(ot Type) bool
		// MatchWithoutNullability This API return true if Type argument
		// is compatible with this param otherwise it returns false.
		// This method ignores nullability for matching.
		MatchWithoutNullability(ot Type) bool
		ShortString() string
		GetNullability() Nullability

		// ReturnType resolves this unresolved type to a concrete Type.
		// Called when this type is used as a function return type (or nested within one).
		//
		// funcParameters are the abstract parameter types from the function definition.
		// argumentTypes are the concrete types passed to the function.
		ReturnType(funcParameters []FuncDefArgType, argumentTypes []Type) (Type, error)

		// WithParameters returns a new instance of this type with the given parameters.
		// This is used in function return type derivation
		WithParameters([]interface{}) (Type, error)
	}

	FixedType interface {
		CompositeType
		WithLength(int32) FixedType
		GetLength() int32
	}

	timestampPrecisionType interface {
		CompositeType
		GetPrecision() TimePrecision
	}
)

var CommonEnumType = &EnumType{}

// EnumType represents an enumeration function parameter.
// It supports a fixed set of declared string values as constant arguments.
type EnumType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Name             string
	Options          []string
}

func (e *EnumType) ToProtoFuncArg() *proto.FunctionArgument {
	// FIXME no proto for enum yet
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: TypeToProto(e)},
	}
}

func (e *EnumType) isRootRef() {}

func (e *EnumType) GetType() Type {
	return e
}

func (e *EnumType) GetTypeVariationReference() uint32 {
	return e.TypeVariationRef
}

func (e *EnumType) Equals(t Type) bool {
	return e.MatchWithNullability(t)
}

func (e *EnumType) WithNullability(n Nullability) Type {
	out := *e
	out.Nullability = n
	return &out
}

func (e *EnumType) GetParameters() []interface{} {
	return []interface{}{}
}

func (e *EnumType) String() string {
	return e.Name
}

func (e *EnumType) SetNullability(n Nullability) FuncDefArgType {
	e.Nullability = n
	return e
}

func (e *EnumType) HasParameterizedParam() bool {
	return false
}

func (e *EnumType) GetParameterizedParams() []interface{} {
	return nil
}

func (e *EnumType) MatchWithNullability(ot Type) bool {
	if e.Nullability != ot.GetNullability() {
		return false
	}
	return e.MatchWithoutNullability(ot)
}

func (e *EnumType) MatchWithoutNullability(ot Type) bool {
	if ot == CommonEnumType {
		return true
	}
	if odt, ok := ot.(*EnumType); ok {
		if e.Name != odt.Name {
			return false
		}
		if len(e.Options) != len(odt.Options) {
			return false
		}
		for i, v := range e.Options {
			if v != odt.Options[i] {
				return false
			}
		}
		return true
	}
	return false
}

func (e *EnumType) ShortString() string {
	// Enum args use "req" as their compound function name key, matching the
	// substrait spec's EnumTypeString constant in extensions/simple_extension.go.
	return "req"
}

func (e *EnumType) GetNullability() Nullability {
	return e.Nullability
}

func (e *EnumType) ReturnType(funcParameters []FuncDefArgType, argumentTypes []Type) (Type, error) {
	return e, nil
}

func (e *EnumType) WithParameters(params []interface{}) (Type, error) {
	panic("EnumType.WithParameters not implemented")
}

type primitiveTypeIFace interface {
	bool | int8 | int16 | ~int32 | ~int64 |
		float32 | float64 | ~string |
		[]byte | IntervalYearToMonth | IntervalDayToSecond | UUID
}

var emptyFixedChar FixedChar

var typeNames = map[reflect.Type]string{
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
	reflect.TypeOf(&FixedBinary{}):                    "fixedbinary",
	reflect.TypeOf(&emptyFixedChar):                   "fixedchar",
	reflect.TypeOf(&VarChar{}):                        "varchar",
	reflect.TypeOf(&PrecisionTimeType{}):              "precision_time",
	reflect.TypeOf(&PrecisionTimestampType{}):         "precision_timestamp",
	reflect.TypeOf(&PrecisionTimestampTzType{}):       "precision_timestamp_tz",
}

var shortNames = map[reflect.Type]string{
	reflect.PointerTo(reflect.TypeOf(true)):           "bool",
	reflect.PointerTo(reflect.TypeOf(int8(0))):        "i8",
	reflect.PointerTo(reflect.TypeOf(int16(0))):       "i16",
	reflect.PointerTo(reflect.TypeOf(int32(0))):       "i32",
	reflect.PointerTo(reflect.TypeOf(int64(0))):       "i64",
	reflect.PointerTo(reflect.TypeOf(float32(0))):     "fp32",
	reflect.PointerTo(reflect.TypeOf(float64(0))):     "fp64",
	reflect.PointerTo(reflect.TypeOf([]byte{})):       "vbin",
	reflect.PointerTo(reflect.TypeOf("")):             "str",
	reflect.PointerTo(reflect.TypeOf(Timestamp(0))):   "ts",
	reflect.PointerTo(reflect.TypeOf(Date(0))):        "date",
	reflect.PointerTo(reflect.TypeOf(Time(0))):        "time",
	reflect.PointerTo(reflect.TypeOf(TimestampTz(0))): "tstz",
	reflect.PointerTo(reflect.TypeOf(UUID{})):         "uuid",
	reflect.TypeOf(&IntervalYearToMonth{}):            "iyear",
	reflect.TypeOf(&IntervalDayToSecond{}):            "iday",
	reflect.TypeOf(&FixedBinary{}):                    "fbin",
	reflect.TypeOf(&emptyFixedChar):                   "fchar",
	reflect.TypeOf(&VarChar{}):                        "vchar",
}

func strNullable(t Type) string {
	return strFromNullability(t.GetNullability())
}

func strFromNullability(nullability Nullability) string {
	if nullability == NullabilityNullable {
		return "?"
	}
	return ""
}

// PrimitiveType is a generic implementation of simple primitive types
// which only need to track if they are nullable and if they are a type
// variation.
type PrimitiveType[T primitiveTypeIFace] struct {
	Nullability      Nullability
	TypeVariationRef uint32
}

func (*PrimitiveType[T]) isRootRef() {}
func (s *PrimitiveType[T]) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *PrimitiveType[T]) GetParameters() []interface{} {
	return []interface{}{}
}

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

func (*PrimitiveType[T]) ShortString() string {
	var z *T
	if n, ok := shortNames[reflect.TypeOf(z)]; ok {
		return n
	}
	return reflect.TypeOf(z).Elem().Name()
}

func (s *PrimitiveType[T]) String() string {
	var z *T
	if n, ok := typeNames[reflect.TypeOf(z)]; ok {
		return n + strNullable(s)
	}
	return reflect.TypeOf(z).Elem().Name() + strNullable(s)
}

func (s *PrimitiveType[T]) HasParameterizedParam() bool {
	// primitive type doesn't have abstract parameters
	return false
}

func (s *PrimitiveType[T]) GetParameterizedParams() []interface{} {
	// primitive type doesn't have any abstract parameters
	return nil
}

func (s *PrimitiveType[T]) SetNullability(n Nullability) FuncDefArgType {
	s.Nullability = n
	return s
}

func (s *PrimitiveType[T]) MatchWithNullability(ot Type) bool {
	if s.Nullability != ot.GetNullability() {
		return false
	}
	return s.MatchWithoutNullability(ot)
}

func (s *PrimitiveType[T]) MatchWithoutNullability(ot Type) bool {
	if _, ok := ot.(*PrimitiveType[T]); ok {
		return true
	}
	return false
}

func (s *PrimitiveType[T]) ReturnType([]FuncDefArgType, []Type) (Type, error) {
	return s, nil
}

func (s *PrimitiveType[T]) WithParameters([]interface{}) (Type, error) {
	return s, nil
}

// create type aliases to the generic structs
//
// The parameterized types below use parameterizedTypeSingleIntegerParam to
// create types that accept a single integer parameter (e.g., precision)
// Other parameterized types (e.g. DecimalType) have their
// own dedicated type definitions (ParameterizedDecimalType) rather than using the generic parameterized type.
type (
	BooleanType                           = PrimitiveType[bool]
	Int8Type                              = PrimitiveType[int8]
	Int16Type                             = PrimitiveType[int16]
	Int32Type                             = PrimitiveType[int32]
	Int64Type                             = PrimitiveType[int64]
	Float32Type                           = PrimitiveType[float32]
	Float64Type                           = PrimitiveType[float64]
	StringType                            = PrimitiveType[string]
	BinaryType                            = PrimitiveType[[]byte]
	TimestampType                         = PrimitiveType[Timestamp]
	DateType                              = PrimitiveType[Date]
	TimeType                              = PrimitiveType[Time]
	TimestampTzType                       = PrimitiveType[TimestampTz]
	IntervalYearType                      = PrimitiveType[IntervalYearToMonth]
	UUIDType                              = PrimitiveType[UUID]
	FixedCharType                         = FixedLenType[FixedChar]
	VarCharType                           = FixedLenType[VarChar]
	FixedBinaryType                       = FixedLenType[FixedBinary]
	ParameterizedVarCharType              = parameterizedTypeSingleIntegerParam[*VarCharType]
	ParameterizedFixedCharType            = parameterizedTypeSingleIntegerParam[*FixedCharType]
	ParameterizedFixedBinaryType          = parameterizedTypeSingleIntegerParam[*FixedBinaryType]
	ParameterizedPrecisionTimeType        = parameterizedTypeSingleIntegerParam[*PrecisionTimeType]
	ParameterizedPrecisionTimestampType   = parameterizedTypeSingleIntegerParam[*PrecisionTimestampType]
	ParameterizedPrecisionTimestampTzType = parameterizedTypeSingleIntegerParam[*PrecisionTimestampTzType]
	ParameterizedIntervalDayType          = parameterizedTypeSingleIntegerParam[*IntervalDayType]
)

// FixedLenType is any of the types which also need to track their specific
// length as they have a fixed length.
type FixedLenType[T FixedChar | VarChar | FixedBinary] struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Length           int32
}

func (*FixedLenType[T]) isRootRef() {}
func (s *FixedLenType[T]) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *FixedLenType[T]) GetParameters() []interface{} {
	return []interface{}{int64(s.Length)}
}

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

func (*FixedLenType[T]) ShortString() string {
	var z *T
	return shortNames[reflect.TypeOf(z)]
}

func (s *FixedLenType[T]) String() string {
	var z *T
	return fmt.Sprintf("%s%s<%d>",
		typeNames[reflect.TypeOf(z)], strNullable(s), s.Length)
}

func (s *FixedLenType[T]) ParameterString() string {
	return fmt.Sprintf("%d", s.Length)
}

func (s *FixedLenType[T]) BaseString() string {
	var z *T
	return typeNames[reflect.TypeOf(z)]
}

func (s *FixedLenType[T]) WithLength(length int32) FixedType {
	out := *s
	out.Length = length
	return &out
}

func (s *FixedLenType[T]) GetLength() int32 {
	return s.Length
}

func (s *FixedLenType[T]) GetReturnType(length int32, nullability Nullability) Type {
	out := *s
	out.Length = length
	out.Nullability = nullability
	return &out
}

// DecimalType is a decimal type with concrete precision and scale parameters, e.g. Decimal(10, 2).
type DecimalType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Scale, Precision int32
}

func (*DecimalType) isRootRef() {}
func (s *DecimalType) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *DecimalType) GetParameters() []interface{} {
	return []interface{}{int64(s.Precision), int64(s.Scale)}
}

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

func (*DecimalType) ShortString() string { return "dec" }
func (s *DecimalType) String() string {
	return fmt.Sprintf("decimal%s<%d,%d>", strNullable(s),
		s.Precision, s.Scale)
}

func (s *DecimalType) ParameterString() string {
	return fmt.Sprintf("%d,%d", s.Precision, s.Scale)
}

func (*DecimalType) BaseString() string {
	return "decimal"
}

type StructType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Types            []Type
}

func (*StructType) isRootRef() {}
func (s *StructType) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *StructType) GetParameters() []interface{} {
	params := make([]interface{}, len(s.Types))
	for i, p := range s.Types {
		params[i] = p
	}
	return params
}

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

func (t *StructType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (*StructType) ShortString() string { return "struct" }

func (t *StructType) String() string {
	var b strings.Builder
	b.WriteString("struct")
	b.WriteString(strNullable(t))
	b.WriteByte('<')
	for i, f := range t.Types {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(f.String())
	}
	b.WriteByte('>')
	return b.String()
}

func (t *StructType) ParameterString() string {
	sb := strings.Builder{}
	for i, typ := range t.Types {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(typ.String())
	}
	return sb.String()
}

func (*StructType) BaseString() string {
	return "struct"
}

// DepthFirstNameCount returns the number of names required to name all
// fields in the struct in depth-first order. Each field consumes one name,
// and nested StructType fields recursively consume additional names,
// including structs nested inside list or map types.
func (t *StructType) DepthFirstNameCount() int {
	count := 0
	for _, typ := range t.Types {
		count++
		count += nestedNameCount(typ)
	}
	return count
}

// nestedNameCount returns the number of additional names consumed by
// nested types. It mirrors the depth-first traversal in NamedStruct.String().
func nestedNameCount(t Type) int {
	switch t := t.(type) {
	case *StructType:
		return t.DepthFirstNameCount()
	case *ListType:
		return nestedNameCount(t.Type)
	case *MapType:
		return nestedNameCount(t.Key) + nestedNameCount(t.Value)
	default:
		return 0
	}
}

// FuncType represents a function type for higher-order functions.
// It describes a function that takes parameters of specified types and
// returns a value of a specified type.
//
// Note: FuncType does not support type variations (always returns 0 for
// GetTypeVariationReference) because function types are abstract and have
// no physical representation.
type FuncType struct {
	Nullability    Nullability
	ParameterTypes []Type
	ReturnType     Type
}

func (*FuncType) isRootRef() {}

func (f *FuncType) WithNullability(n Nullability) Type {
	out := *f
	out.Nullability = n
	return &out
}

func (f *FuncType) GetParameters() []interface{} {
	// Return all type components: parameter types + return type
	// This is needed for type derivation with parameterized function types like func<T1, T2 -> T3>
	params := make([]interface{}, len(f.ParameterTypes)+1)
	for i, p := range f.ParameterTypes {
		params[i] = p
	}
	params[len(f.ParameterTypes)] = f.ReturnType
	return params
}

func (f *FuncType) GetType() Type                     { return f }
func (f *FuncType) GetNullability() Nullability       { return f.Nullability }
func (f *FuncType) GetTypeVariationReference() uint32 { return 0 } // FuncType doesn't support variations

func (f *FuncType) Equals(rhs Type) bool {
	if b, ok := rhs.(*FuncType); ok {
		switch {
		case f.Nullability != b.Nullability:
			return false
		case len(f.ParameterTypes) != len(b.ParameterTypes):
			return false
		case !f.ReturnType.Equals(b.ReturnType):
			return false
		}

		for i := range f.ParameterTypes {
			if !f.ParameterTypes[i].Equals(b.ParameterTypes[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (f *FuncType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: f.ToProto()},
	}
}

func (*FuncType) ShortString() string { return "func" }

func (f *FuncType) String() string {
	var b strings.Builder
	b.WriteString("func")
	b.WriteString(strNullable(f))
	b.WriteByte('<')
	for i, p := range f.ParameterTypes {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(p.String())
	}
	b.WriteString(" -> ")
	b.WriteString(f.ReturnType.String())
	b.WriteByte('>')
	return b.String()
}

type ListType struct {
	Nullability      Nullability
	TypeVariationRef uint32

	Type Type
}

func (*ListType) isRootRef() {}
func (s *ListType) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *ListType) GetParameters() []interface{} {
	return []interface{}{s.Type}
}

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

func (t *ListType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (*ListType) ShortString() string { return "list" }

func (t *ListType) String() string {
	return "list" + strNullable(t) + "<" + t.Type.String() + ">"
}

func (s *ListType) ParameterString() string {
	return s.Type.String()
}

func (*ListType) BaseString() string {
	return "list"
}

type MapType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	Key, Value       Type
}

func (*MapType) isRootRef() {}
func (s *MapType) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *MapType) GetParameters() []interface{} {
	return []interface{}{s.Key, s.Value}
}

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

func (t *MapType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

func (t *MapType) ShortString() string { return "map" }

func (t *MapType) String() string {
	return "map" + strNullable(t) + "<" + t.Key.String() + ", " + t.Value.String() + ">"
}

func (t *MapType) ParameterString() string {
	return fmt.Sprintf("%s, %s", t.Key.String(), t.Value.String())
}

func (*MapType) BaseString() string {
	return "map"
}

// TypeParam represents a type parameter for a user defined type
type TypeParam interface {
	Equals(TypeParam) bool
}

// NullParameter is an explicitly null/unspecified parameter, to select
// the default value (if any).
type NullParameter struct{}

func (NullParameter) Equals(p TypeParam) bool {
	_, ok := p.(NullParameter)
	return ok
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

// BooleanParameter is a type parameter like <true> for a type.
type BooleanParameter bool

func (b BooleanParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(BooleanParameter); ok {
		return b == rhs
	}
	return false
}

// IntegerParameter is the type parameter like 10 in VARCHAR<10>
type IntegerParameter int64

func (b IntegerParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(IntegerParameter); ok {
		return b == rhs
	}
	return false
}

// EnumParameter is a type parameter that is some enum value
type EnumParameter string

func (b EnumParameter) Equals(p TypeParam) bool {
	if rhs, ok := p.(EnumParameter); ok {
		return b == rhs
	}
	return false
}

// StringParameter is a type parameter which is a string value
type StringParameter string

func (p StringParameter) String() string {
	return string(p)
}

func (p StringParameter) Equals(o TypeParam) bool {
	if rhs, ok := o.(StringParameter); ok {
		return p == rhs
	}
	return false
}

func (p StringParameter) Evaluate(symbolTable map[string]any) (any, error) {
	if v, ok := symbolTable[string(p)]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("symbol not found: stringParameter %s", p)
}

type UserDefinedType struct {
	Nullability      Nullability
	TypeVariationRef uint32
	TypeReference    uint32
	TypeParameters   []TypeParam
}

func (*UserDefinedType) isRootRef() {}
func (s *UserDefinedType) WithNullability(n Nullability) Type {
	out := *s
	out.Nullability = n
	return &out
}

func (s *UserDefinedType) GetParameters() []interface{} {
	params := make([]interface{}, len(s.TypeParameters))
	for i, p := range s.TypeParameters {
		params[i] = p
	}
	return params
}

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

func (t *UserDefinedType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: t.ToProto()},
	}
}

// exists for meeting the interface, but the correct short name for
// a user defined type is "u!name" which requires looking up the
// type first via the type reference to find the name.
func (*UserDefinedType) ShortString() string { return "" }

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

func (n NamedStruct) String() string {
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
			b.WriteString(strNullable(t))
		case *MapType:
			b.WriteString("map<")
			writeType(t.Key)
			b.WriteString(",")
			writeType(t.Value)
			b.WriteString(">")
			b.WriteString(strNullable(t))
		case *ListType:
			b.WriteString("list<")
			writeType(t.Type)
			b.WriteString(">")
			b.WriteString(strNullable(t))
		default:
			b.WriteString(t.String())
		}
	}

	b.WriteString("NSTRUCT<")
	for i, t := range n.Struct.Types {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(n.Names[nameIdx])
		b.WriteString(": ")
		nameIdx++

		writeType(t)
	}
	b.WriteString(">")

	return b.String()
}

// RecordType is the type of a record (or row) comprising a list of fields (or columns).
type RecordType struct {
	types []Type
}

func NewRecordTypeFromTypes(types []Type) *RecordType {
	return &RecordType{types: types}
}

func NewRecordTypeFromStruct(s StructType) *RecordType {
	return &RecordType{types: s.Types}
}

func (r RecordType) Equals(other *RecordType) bool {
	return r.AsStructType().Equals(other.AsStructType())
}

func (r RecordType) String() string {
	return r.AsStructType().String()
}

func (r RecordType) GetFieldRef(index int32) Type {
	return r.types[index]
}

func (r RecordType) FieldCount() int32 {
	return int32(len(r.types))
}

func (r RecordType) AsStructType() *StructType {
	return &StructType{Nullability: NullabilityRequired, Types: r.types}
}

func (r RecordType) Types() []Type {
	return r.types
}

func (r RecordType) Concat(other RecordType) RecordType {
	return RecordType{types: append(r.Types(), other.Types()...)}
}

func (d Date) ToTimeString() string {
	date := civil.Date{Year: 1970, Month: time.January, Day: 1}
	date = date.AddDays(int(d))
	return date.String()
}

func timeFromPrecisionUnits(units int64, precision TimePrecision) time.Time {
	var tm time.Time
	switch precision {
	case PrecisionSeconds:
		tm = time.Unix(units, 0)
	case PrecisionDeciSeconds:
		tm = time.Unix(units/10, units%10*100000000)
	case PrecisionCentiSeconds:
		tm = time.Unix(units/100, units%100*10000000)
	case PrecisionMilliSeconds:
		tm = time.UnixMilli(units)
	case PrecisionEMinus4Seconds:
		tm = time.Unix(units/10000, units%10000*100000)
	case PrecisionEMinus5Seconds:
		tm = time.Unix(units/100000, units%100000*10000)
	case PrecisionMicroSeconds:
		tm = time.UnixMicro(units)
	case PrecisionEMinus7Seconds:
		tm = time.Unix(units/10000000, units%10000000*100)
	case PrecisionEMinus8Seconds:
		tm = time.Unix(units/100000000, units%100000000*10)
	case PrecisionNanoSeconds:
		tm = time.Unix(units/1000000000, units%1000000000)
	default:
		panic("unsupported precision")
	}
	return tm
}

func (t Time) ToTimeString() string {
	tm := time.UnixMicro(int64(t))
	return tm.UTC().Format(time.TimeOnly)
}

func (t Time) ToIsoTimeString() string {
	tm := time.UnixMicro(int64(t))
	return tm.UTC().Format("15:04:05.000000")
}

func (t Time) ToPrecisionTime(precision TimePrecision) time.Time {
	return timeFromPrecisionUnits(int64(t), precision)
}

func (t Timestamp) ToTime() time.Time {
	return time.UnixMicro(int64(t))
}

func (t Timestamp) ToPrecisionTime(precision TimePrecision) time.Time {
	return timeFromPrecisionUnits(int64(t), precision)
}

func (t Timestamp) ToTimeString() string {
	tm := any(t).(TimeConverter).ToTime()
	return tm.UTC().Format("2006-01-02 15:04:05.999999999")
}

func (t Timestamp) ToIsoTimeString() string {
	tm := any(t).(TimeConverter).ToTime()
	return tm.UTC().Format("2006-01-02T15:04:05.999999999")
}

func (t TimestampTz) ToTime() time.Time {
	return time.UnixMicro(int64(t))
}

func (t TimestampTz) ToPrecisionTime(precision TimePrecision) time.Time {
	return timeFromPrecisionUnits(int64(t), precision)
}

func (t TimestampTz) ToTimeString() string {
	tm := any(t).(TimeConverter).ToTime()
	return tm.UTC().Format(time.RFC3339Nano)
}

func (t TimestampTz) ToIsoTimeString() string {
	tm := any(t).(TimeConverter).ToTime()
	return tm.UTC().Format("2006-01-02T15:04:05.999999999")
}

func GetTimeValueByPrecision(tm time.Time, precision TimePrecision) int64 {
	switch precision {
	case PrecisionSeconds:
		return tm.Unix()
	case PrecisionDeciSeconds:
		return tm.UnixMilli() / 100
	case PrecisionCentiSeconds:
		return tm.UnixMilli() / 10
	case PrecisionMilliSeconds:
		return tm.UnixMilli()
	case PrecisionEMinus4Seconds:
		return tm.UnixMicro() / 100
	case PrecisionEMinus5Seconds:
		return tm.UnixMicro() / 10
	case PrecisionMicroSeconds:
		return tm.UnixMicro()
	case PrecisionEMinus7Seconds:
		return tm.UnixNano() / 100
	case PrecisionEMinus8Seconds:
		return tm.UnixNano() / 10
	case PrecisionNanoSeconds:
		return tm.UnixNano()
	default:
		panic(fmt.Sprintf("unknown TimePrecision %v", precision))
	}
}
