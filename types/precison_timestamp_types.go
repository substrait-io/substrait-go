package types

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/substrait-io/substrait-go/proto"
)

type TimePrecision int

const (
	UnknownPrecision TimePrecision = iota
	Seconds
	Milliseconds
	Microseconds
	Nanoseconds
)

func timePrecisionToProtoVal(val TimePrecision) int32 {
	switch val {
	case Seconds:
		return 0
	case Milliseconds:
		return 3
	case Microseconds:
		return 6
	case Nanoseconds:
		return 9
	}
	panic("unreachable")
}

func ProtoToTimePrecision(val int32) (TimePrecision, error) {
	switch val {
	case 0:
		return Seconds, nil
	case 3:
		return Milliseconds, nil
	case 6:
		return Microseconds, nil
	case 9:
		return Nanoseconds, nil
	}
	return UnknownPrecision, errors.Newf("invalid TimePrecision value %d", val)
}

type PrecisionTimeStampType struct {
	precision        TimePrecision
	typeVariationRef uint32
	nullability      Nullability
}

func NewPrecisionTimestamp(precision TimePrecision) PrecisionTimeStampType {
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

type PrecisionTimeStampTzType struct {
	PrecisionTimeStampType
}

// creates a new precision timestamp with nullability as Nullable
func NewPrecisionTimestampTz(precision TimePrecision) PrecisionTimeStampTzType {
	return PrecisionTimeStampTzType{
		PrecisionTimeStampType: PrecisionTimeStampType{
			precision:   precision,
			nullability: NullabilityNullable,
		},
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
