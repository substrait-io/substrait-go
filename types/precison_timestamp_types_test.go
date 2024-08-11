package types

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

var expectedProtoValMap = map[TimePrecision]int32{
	PrecisionSeconds:        0,
	PrecisionDeciSeconds:    1,
	PrecisionCentiSeconds:   2,
	PrecisionMilliSeconds:   3,
	PrecisionEMinus4Seconds: 4,
	PrecisionEMinus5Seconds: 5,
	PrecisionMicroSeconds:   6,
	PrecisionEMinus7Seconds: 7,
	PrecisionEMinus8Seconds: 8,
	PrecisionNanoSeconds:    9,
}

func TestProtoToTimePrecision(t *testing.T) {
	for expectedTimePrecision, expectedProtoVal := range expectedProtoValMap {
		got, err := ProtoToTimePrecision(expectedProtoVal)
		assert.NoError(t, err)
		assert.Equal(t, expectedTimePrecision, got)
	}

	got, err := ProtoToTimePrecision(-1)
	assert.Error(t, err)
	assert.Equal(t, PrecisionUnknown, got)
	_, err = ProtoToTimePrecision(10)
	assert.Error(t, err)
	assert.Equal(t, PrecisionUnknown, got)
}

func TestNewPrecisionTimestampType(t *testing.T) {
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedPrecisionTimeStampType := PrecisionTimeStampType{precision: precision, nullability: nullability}
			expectedPrecisionTimeStampTzType := PrecisionTimeStampTzType{PrecisionTimeStampType: expectedPrecisionTimeStampType}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(expectedPrecisionTimeStampType), precision.ToProtoVal())
			// verify PrecisionTimeStampType
			createdPrecTimeStampType := NewPrecisionTimestampType(precision).WithNullability(nullability)
			createdPrecTimeStamp := createdPrecTimeStampType.(PrecisionTimeStampType)
			assert.True(t, createdPrecTimeStamp.Equals(expectedPrecisionTimeStampType))
			assert.Equal(t, expectedProtoValMap[precision], createdPrecTimeStamp.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdPrecTimeStamp.GetNullability())
			assert.Zero(t, createdPrecTimeStamp.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("precisiontimestamp%s", expectedFormatString), createdPrecTimeStamp.String())
			assert.Equal(t, "prets", createdPrecTimeStamp.ShortString())
			assertPrecisionTimeStampProto(t, precision, nullability, createdPrecTimeStamp)

			// verify PrecisionTimeStampTzType
			createdPrecTimeStampTzType := NewPrecisionTimestampTzType(precision).WithNullability(nullability)
			createdPrecTimeStampTz := createdPrecTimeStampTzType.(PrecisionTimeStampTzType)
			assert.True(t, createdPrecTimeStampTz.Equals(expectedPrecisionTimeStampTzType))
			assert.Equal(t, expectedProtoValMap[precision], createdPrecTimeStampTz.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdPrecTimeStampTz.GetNullability())
			assert.Zero(t, createdPrecTimeStampTz.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("precisiontimestamptz%s", expectedFormatString), createdPrecTimeStampTz.String())
			assert.Equal(t, "pretstz", createdPrecTimeStampTz.ShortString())
			assertPrecisionTimeStampTzProto(t, precision, nullability, createdPrecTimeStampTz)

			// assert that both types are not equal
			assert.False(t, createdPrecTimeStampType.Equals(createdPrecTimeStampTzType))
			assert.False(t, createdPrecTimeStampTzType.Equals(createdPrecTimeStampType))
		}
	}
}

func assertPrecisionTimeStampProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability,
	toVerifyType PrecisionTimeStampType) {

	expectedTypeProto := &proto.Type{Kind: &proto.Type_PrecisionTimestamp_{
		PrecisionTimestamp: &proto.Type_PrecisionTimestamp{
			Precision:   expectedPrecision.ToProtoVal(),
			Nullability: expectedNullability,
		},
	}}
	if diff := cmp.Diff(toVerifyType.ToProto(), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("precisionTimeStamp proto didn't match, diff:\n%v", diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(toVerifyType.ToProtoFuncArg(), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("precisionTimeStamp proto didn't match, diff:\n%v", diff)
	}
}

func assertPrecisionTimeStampTzProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability, toVerifyType PrecisionTimeStampTzType) {
	expectedTypeProto := &proto.Type{Kind: &proto.Type_PrecisionTimestampTz{
		PrecisionTimestampTz: &proto.Type_PrecisionTimestampTZ{
			Precision:   expectedPrecision.ToProtoVal(),
			Nullability: expectedNullability,
		},
	}}
	if diff := cmp.Diff(toVerifyType.ToProto(), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("precisionTimeStampTz proto didn't match, diff:\n%v", diff)
	}
	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(toVerifyType.ToProtoFuncArg(), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("precisionTimeStampTz proto didn't match, diff:\n%v", diff)
	}
}
