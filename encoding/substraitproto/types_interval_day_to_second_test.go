// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestIntervalDayToSecondMatchesDescriptor(t *testing.T) {
	wantScalars := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"days":       {1, protoreflect.Int32Kind},
		"seconds":    {2, protoreflect.Int32Kind},
		"precision":  {4, protoreflect.Int32Kind},
		"subseconds": {5, protoreflect.Int64Kind},
	}

	desc := (&proto.Expression_Literal_IntervalDayToSecond{}).ProtoReflect().Descriptor()
	fields := desc.Fields()
	require.Equal(t, len(wantScalars), fields.Len(), "spec interval_day_to_second field set changed")

	require.Equal(t, len(wantScalars), reflect.TypeOf(types.IntervalDayToSecond{}).NumField(),
		"types.IntervalDayToSecond field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := wantScalars[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

// Round-trip domain->proto->domain, so a swapped or dropped field fails here.
func TestIntervalDayToSecondRoundTrip(t *testing.T) {
	t.Run("precision round trip", func(t *testing.T) {
		in := &types.IntervalDayToSecond{Days: 1, Seconds: 2, Subseconds: 3, Precision: types.PrecisionNanoSeconds}
		p := substraitproto.IntervalDayToSecondToProto(in)
		assert.EqualValues(t, 9, p.Precision)
		got, err := substraitproto.IntervalDayToSecondFromProto(p)
		require.NoError(t, err)
		assert.Equal(t, in, got)
	})

	// An unset precision is 0, which means the value is in whole seconds.
	t.Run("unset precision is seconds", func(t *testing.T) {
		p := &proto.Expression_Literal_IntervalDayToSecond{Days: 1, Seconds: 2}
		got, err := substraitproto.IntervalDayToSecondFromProto(p)
		require.NoError(t, err)
		assert.Equal(t, &types.IntervalDayToSecond{Days: 1, Seconds: 2, Precision: types.PrecisionSeconds}, got)
	})

	assert.Nil(t, substraitproto.IntervalDayToSecondToProto(nil))
	got, err := substraitproto.IntervalDayToSecondFromProto(nil)
	require.NoError(t, err)
	assert.Nil(t, got)

	// GetPrecisionProtoVal is nil-safe and yields the precision's protobuf value.
	assert.Zero(t, (*types.IntervalDayToSecond)(nil).GetPrecisionProtoVal())
	assert.EqualValues(t, 9, (&types.IntervalDayToSecond{Precision: types.PrecisionNanoSeconds}).GetPrecisionProtoVal())
}
