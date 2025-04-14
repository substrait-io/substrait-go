package types

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIntervalDayType(t *testing.T) {
	anotherType := &FixedCharType{Length: 10, Nullability: NullabilityNullable}
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedIntervalDayType := &IntervalDayType{Precision: precision, Nullability: nullability}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(expectedIntervalDayType), precision.ToProtoVal())

			parameters := expectedIntervalDayType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			// verify IntervalDayType
			createdIntervalDayTypeIfc := (&IntervalDayType{Precision: precision}).WithNullability(nullability)
			createdIntervalDayType := createdIntervalDayTypeIfc.(*IntervalDayType)
			assert.True(t, createdIntervalDayType.Equals(expectedIntervalDayType))
			assert.Equal(t, expectedProtoValMap[precision], createdIntervalDayType.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdIntervalDayType.GetNullability())
			assert.Zero(t, createdIntervalDayType.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("interval_day%s", expectedFormatString), createdIntervalDayType.String())
			assert.Equal(t, "iday", createdIntervalDayType.ShortString())
			assert.Equal(t, "interval_day", createdIntervalDayType.BaseString())
			assert.Equal(t, precision, createdIntervalDayType.GetPrecision())
			expectedParameterString := fmt.Sprintf("%d", precision.ToProtoVal())
			assert.Equal(t, expectedParameterString, createdIntervalDayType.ParameterString())
			assertIntervalDayTypeProto(t, precision, nullability, createdIntervalDayType)
			assert.False(t, createdIntervalDayTypeIfc.Equals(anotherType))
		}
	}
}

func assertIntervalDayTypeProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability,
	toVerifyType *IntervalDayType) {

	expectedPrecisionProtoVal := expectedPrecision.ToProtoVal()
	expectedTypeProto := &proto.Type{Kind: &proto.Type_IntervalDay_{
		IntervalDay: &proto.Type_IntervalDay{
			Precision:   &expectedPrecisionProtoVal,
			Nullability: expectedNullability,
		},
	}}
	if diff := cmp.Diff(toVerifyType.ToProto(), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalDayType proto didn't match, diff:\n%v", diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(toVerifyType.ToProtoFuncArg(), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalDayType func arg proto didn't match, diff:\n%v", diff)
	}
}
