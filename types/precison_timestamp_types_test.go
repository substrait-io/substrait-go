// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
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

func TestNewPrecisionTimeType(t *testing.T) {
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedPrecisionTimeType := PrecisionTimeType{Precision: precision, Nullability: nullability}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(&expectedPrecisionTimeType), precision.ToProtoVal())

			parameters := expectedPrecisionTimeType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			// verify PrecisionTimeType
			createdPrecTimeType := NewPrecisionTimeType(precision).WithNullability(nullability)
			createdPrecTime := createdPrecTimeType.(*PrecisionTimeType)
			assert.True(t, createdPrecTime.Equals(&expectedPrecisionTimeType))
			assert.Equal(t, expectedProtoValMap[precision], createdPrecTime.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdPrecTime.GetNullability())
			assert.Zero(t, createdPrecTime.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("precision_time%s", expectedFormatString), createdPrecTime.String())
			assert.Equal(t, "pt", createdPrecTime.ShortString())
			assertPrecisionTimeProto(t, precision, nullability, *createdPrecTime)
		}
	}
}

func assertPrecisionTimeProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability,
	toVerifyType PrecisionTimeType) {

	expectedTypeProto := &proto.Type{Kind: &proto.Type_PrecisionTime_{
		PrecisionTime: &proto.Type_PrecisionTime{
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

func TestNewPrecisionTimestampType(t *testing.T) {
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedPrecisionTimeStampType := PrecisionTimestampType{Precision: precision, Nullability: nullability}
			expectedPrecisionTimeStampTzType := PrecisionTimestampTzType{PrecisionTimestampType: expectedPrecisionTimeStampType}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(&expectedPrecisionTimeStampType), precision.ToProtoVal())

			parameters := expectedPrecisionTimeStampType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			parameters = expectedPrecisionTimeStampTzType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			// verify PrecisionTimestampType
			createdPrecTimeStampType := NewPrecisionTimestampType(precision).WithNullability(nullability)
			createdPrecTimeStamp := createdPrecTimeStampType.(*PrecisionTimestampType)
			assert.True(t, createdPrecTimeStamp.Equals(&expectedPrecisionTimeStampType))
			assert.Equal(t, expectedProtoValMap[precision], createdPrecTimeStamp.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdPrecTimeStamp.GetNullability())
			assert.Zero(t, createdPrecTimeStamp.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("precision_timestamp%s", expectedFormatString), createdPrecTimeStamp.String())
			assert.Equal(t, "pts", createdPrecTimeStamp.ShortString())
			assertPrecisionTimeStampProto(t, precision, nullability, *createdPrecTimeStamp)

			// verify PrecisionTimestampTzType
			createdPrecTimeStampTzType := NewPrecisionTimestampTzType(precision).WithNullability(nullability)
			createdPrecTimeStampTz := createdPrecTimeStampTzType.(*PrecisionTimestampTzType)
			assert.True(t, createdPrecTimeStampTz.Equals(&expectedPrecisionTimeStampTzType))
			assert.Equal(t, expectedProtoValMap[precision], createdPrecTimeStampTz.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdPrecTimeStampTz.GetNullability())
			assert.Zero(t, createdPrecTimeStampTz.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("precision_timestamp_tz%s", expectedFormatString), createdPrecTimeStampTz.String())
			assert.Equal(t, "ptstz", createdPrecTimeStampTz.ShortString())
			assertPrecisionTimeStampTzProto(t, precision, nullability, *createdPrecTimeStampTz)

			// assert that both types are not equal
			assert.False(t, createdPrecTimeStampType.Equals(createdPrecTimeStampTzType))
			assert.False(t, createdPrecTimeStampTzType.Equals(createdPrecTimeStampType))
		}
	}
}

func assertPrecisionTimeStampProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability,
	toVerifyType PrecisionTimestampType) {

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

func assertPrecisionTimeStampTzProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability, toVerifyType PrecisionTimestampTzType) {
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

func TestSubSecondsToDuration(t *testing.T) {
	tests := []struct {
		name       string
		subSeconds int64
		precision  TimePrecision
		want       time.Duration
	}{
		{"0.000000001s", 1, PrecisionNanoSeconds, time.Nanosecond},
		{"0.00000001s", 1, PrecisionEMinus8Seconds, time.Nanosecond * 10},
		{"0.0000001s", 1, PrecisionEMinus7Seconds, time.Nanosecond * 100},
		{"0.000001s", 1, PrecisionMicroSeconds, time.Microsecond},
		{"0.00001s", 1, PrecisionEMinus5Seconds, time.Microsecond * 10},
		{"0.0001s", 1, PrecisionEMinus4Seconds, time.Microsecond * 100},
		{"0.001s", 1, PrecisionMilliSeconds, time.Millisecond},
		{"0.01s", 1, PrecisionCentiSeconds, time.Millisecond * 10},
		{"0.1s", 1, PrecisionDeciSeconds, time.Millisecond * 100},
		{"1s", 1, PrecisionSeconds, time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SubSecondsToDuration(tt.subSeconds, tt.precision), "SubSecondsToDuration(%v, %v)", tt.subSeconds, tt.precision)
		})
	}
}
