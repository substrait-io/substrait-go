package types

import (
	"fmt"
	"github.com/substrait-io/substrait-go/proto"
)

// TimePrecision is used to represent the precision of a timestamp
type TimePrecision int32

const PrecisionUnknown TimePrecision = -1

// precision values are proto values
const (
	PrecisionSeconds TimePrecision = iota
	PrecisionDeciSeconds
	PrecisionCentiSeconds
	PrecisionMilliSeconds
	PrecisionEMinus4Seconds // 10^-4 of seconds
	PrecisionEMinus5Seconds // 10^-5 of seconds
	PrecisionMicroSeconds
	PrecisionEMinus7Seconds // 10^-7 of seconds
	PrecisionEMinus8Seconds // 10^-8 of seconds
	PrecisionNanoSeconds
)

func timePrecisionToProtoVal(val TimePrecision) int32 {
	return int32(val)
}

func ProtoToTimePrecision(val int32) (TimePrecision, error) {
	if val < 0 || val > 9 {
		return PrecisionUnknown, fmt.Errorf("invalid TimePrecision value %d", val)
	}
	return TimePrecision(val), nil
}

// PrecisionTimeStampType this is used to represent a type of precision timestamp
type PrecisionTimeStampType struct {
	precision        TimePrecision
	typeVariationRef uint32
	nullability      Nullability
}

// NewPrecisionTimestampType creates a type of new precision timestamp.
// Created type has nullability as Nullable
func NewPrecisionTimestampType(precision TimePrecision) PrecisionTimeStampType {
	return PrecisionTimeStampType{
		precision:   precision,
		nullability: NullabilityNullable,
	}
}

func (m PrecisionTimeStampType) GetPrecisionProtoVal() int32 {
	return timePrecisionToProtoVal(m.precision)
}

func (PrecisionTimeStampType) isRootRef() {}
func (m PrecisionTimeStampType) WithNullability(n Nullability) Type {
	return m.withNullability(n)
}

func (m PrecisionTimeStampType) withNullability(n Nullability) PrecisionTimeStampType {
	return PrecisionTimeStampType{
		precision:   m.precision,
		nullability: n,
	}
}

func (m PrecisionTimeStampType) GetType() Type                     { return m }
func (m PrecisionTimeStampType) GetNullability() Nullability       { return m.nullability }
func (m PrecisionTimeStampType) GetTypeVariationReference() uint32 { return m.typeVariationRef }
func (m PrecisionTimeStampType) Equals(rhs Type) bool {
	if o, ok := rhs.(PrecisionTimeStampType); ok {
		return o == m
	}
	return false
}

func (m PrecisionTimeStampType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: m.ToProto()},
	}
}

func (m PrecisionTimeStampType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_PrecisionTimestamp_{
		PrecisionTimestamp: &proto.Type_PrecisionTimestamp{
			Precision:              timePrecisionToProtoVal(m.precision),
			Nullability:            m.nullability,
			TypeVariationReference: m.typeVariationRef}}}
}

func (PrecisionTimeStampType) ShortString() string { return "prets" }
func (m PrecisionTimeStampType) String() string {
	return fmt.Sprintf("precisiontimestamp%s<%d>", strNullable(m),
		timePrecisionToProtoVal(m.precision))
}

// PrecisionTimeStampTzType this is used to represent a type of precision timestamp with TimeZone
type PrecisionTimeStampTzType struct {
	PrecisionTimeStampType
}

// NewPrecisionTimestampTzType creates a type of new precision timestamp with TimeZone.
// Created type has nullability as Nullable
func NewPrecisionTimestampTzType(precision TimePrecision) PrecisionTimeStampTzType {
	return PrecisionTimeStampTzType{
		PrecisionTimeStampType: PrecisionTimeStampType{
			precision:   precision,
			nullability: NullabilityNullable,
		},
	}
}

func (m PrecisionTimeStampTzType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: m.ToProto()},
	}
}

func (m PrecisionTimeStampTzType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_PrecisionTimestampTz{
		PrecisionTimestampTz: &proto.Type_PrecisionTimestampTZ{
			Precision:              timePrecisionToProtoVal(m.precision),
			Nullability:            m.nullability,
			TypeVariationReference: m.typeVariationRef}}}
}

func (m PrecisionTimeStampTzType) String() string {
	return fmt.Sprintf("precisiontimestamptz%s<%d>", strNullable(m),
		timePrecisionToProtoVal(m.precision))
}

func (m PrecisionTimeStampTzType) WithNullability(n Nullability) Type {
	return PrecisionTimeStampTzType{
		PrecisionTimeStampType: m.PrecisionTimeStampType.withNullability(n),
	}
}

func (m PrecisionTimeStampTzType) Equals(rhs Type) bool {
	if o, ok := rhs.(PrecisionTimeStampTzType); ok {
		return o == m
	}
	return false
}
func (PrecisionTimeStampTzType) ShortString() string { return "pretstz" }
