package types

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestNewIntervalYearToMonthType(t *testing.T) {
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, nullability := range allPossibleNullability {
		expectedIntervalType := IntervalYearToMonthType{nullability: nullability}

		parameters := expectedIntervalType.GetParameters()
		assert.Len(t, parameters, 0)
		// verify IntervalYearToMonthType
		createdIntervalTypeIfcType := NewIntervalYearToMonthType().WithTypeVariationRef(0).WithNullability(nullability)
		createdIntervalType := createdIntervalTypeIfcType.(IntervalYearToMonthType)
		assert.True(t, createdIntervalType.Equals(expectedIntervalType))
		assert.Equal(t, nullability, createdIntervalType.GetNullability())
		assert.Zero(t, createdIntervalTypeIfcType.GetTypeVariationReference())
		assert.Equal(t, fmt.Sprintf("interval_year%s", strNullable(expectedIntervalType)), createdIntervalType.String())
		assert.Equal(t, "iyear", createdIntervalType.ShortString())
		assertIntervalYearToMonthTypeProto(t, nullability, createdIntervalType)
	}
}

func assertIntervalYearToMonthTypeProto(t *testing.T, expectedNullability Nullability,
	toVerifyType IntervalYearToMonthType) {

	expectedTypeProto := &proto.Type{Kind: &proto.Type_IntervalYear_{
		IntervalYear: &proto.Type_IntervalYear{
			Nullability: expectedNullability,
		},
	}}
	if diff := cmp.Diff(toVerifyType.ToProto(), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalYearToMonthType proto didn't match, diff:\n%v", diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(toVerifyType.ToProtoFuncArg(), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalYearToMonthType func arg proto didn't match, diff:\n%v", diff)
	}
}
