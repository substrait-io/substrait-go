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
	DeciSeconds
	CentiSeconds
	MilliSeconds
	EMinus4Seconds // 10^-4 of seconds
	EMinus5Seconds // 10^-5 of seconds
	MicroSeconds
	EMinus7Seconds // 10^-7 of seconds
	EMinus8Seconds // 10^-8 of seconds
	NanoSeconds
)

func timePrecisionToProtoVal(val TimePrecision) int32 {
	switch val {
	case Seconds:
		return 0
	case DeciSeconds:
		return 1
	case CentiSeconds:
		return 2
	case MilliSeconds:
		return 3
	case EMinus4Seconds:
		return 4
	case EMinus5Seconds:
		return 5
	case MicroSeconds:
		return 6
	case EMinus7Seconds:
		return 7
	case EMinus8Seconds:
		return 8
	case NanoSeconds:
		return 9
	}
	panic("unreachable")
}

func ProtoToTimePrecision(val int32) (TimePrecision, error) {
	switch val {
	case 0:
		return Seconds, nil
	case 1:
		return DeciSeconds, nil
	case 2:
		return CentiSeconds, nil
	case 3:
		return MilliSeconds, nil
	case 4:
		return EMinus4Seconds, nil
	case 5:
		return EMinus5Seconds, nil
	case 6:
		return MicroSeconds, nil
	case 7:
		return EMinus7Seconds, nil
	case 8:
		return EMinus8Seconds, nil
	case 9:
		return NanoSeconds, nil
	}
	return UnknownPrecision, errors.Newf("invalid TimePrecision value %d", val)
}

type PrecisionTimeStampType struct {
	precision        TimePrecision
	typeVariationRef uint32
	nullability      Nullability
}

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

type PrecisionTimeStampTzType struct {
	PrecisionTimeStampType
}

// creates a new precision timestamp with nullability as Nullable
func NewPrecisionTimestampTzType(precision TimePrecision) PrecisionTimeStampTzType {
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
