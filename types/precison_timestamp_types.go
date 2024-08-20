package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
)

// TimePrecision is used to represent the precision of a timestamp
type TimePrecision int32

const (
	PrecisionUnknown TimePrecision = -1
	// below precision values are proto values
	PrecisionSeconds        TimePrecision = 0
	PrecisionDeciSeconds    TimePrecision = 1
	PrecisionCentiSeconds   TimePrecision = 2
	PrecisionMilliSeconds   TimePrecision = 3
	PrecisionEMinus4Seconds TimePrecision = 4 // 10^-4 of seconds
	PrecisionEMinus5Seconds TimePrecision = 5 // 10^-5 of seconds
	PrecisionMicroSeconds   TimePrecision = 6
	PrecisionEMinus7Seconds TimePrecision = 7 // 10^-7 of seconds
	PrecisionEMinus8Seconds TimePrecision = 8 // 10^-8 of seconds
	PrecisionNanoSeconds    TimePrecision = 9
)

func (m TimePrecision) ToProtoVal() int32 {
	return int32(m)
}

func ProtoToTimePrecision(val int32) (TimePrecision, error) {
	if val < PrecisionSeconds.ToProtoVal() || val > PrecisionNanoSeconds.ToProtoVal() {
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
	return m.precision.ToProtoVal()
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
			Precision:              m.precision.ToProtoVal(),
			Nullability:            m.nullability,
			TypeVariationReference: m.typeVariationRef}}}
}

func (PrecisionTimeStampType) ShortString() string { return "prets" }
func (m PrecisionTimeStampType) String() string {
	return fmt.Sprintf("precisiontimestamp%s<%d>", strNullable(m),
		m.precision.ToProtoVal())
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
			Precision:              m.precision.ToProtoVal(),
			Nullability:            m.nullability,
			TypeVariationReference: m.typeVariationRef}}}
}

func (m PrecisionTimeStampTzType) String() string {
	return fmt.Sprintf("precisiontimestamptz%s<%d>", strNullable(m),
		m.precision.ToProtoVal())
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
