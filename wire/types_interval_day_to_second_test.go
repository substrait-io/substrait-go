// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestIntervalDayToSecondMatchesDescriptor(t *testing.T) {
	// Scalar fields are pinned by wire number and kind; the precision_mode oneof is pinned
	// separately since the domain type collapses its two arms into a single Precision field.
	wantScalars := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"days":       {1, protoreflect.Int32Kind},
		"seconds":    {2, protoreflect.Int32Kind},
		"subseconds": {5, protoreflect.Int64Kind},
	}
	wantOneof := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"microseconds": {3, protoreflect.Int32Kind},
		"precision":    {4, protoreflect.Int32Kind},
	}

	desc := (&proto.Expression_Literal_IntervalDayToSecond{}).ProtoReflect().Descriptor()
	fields := desc.Fields()
	require.Equal(t, len(wantScalars)+len(wantOneof), fields.Len(), "spec interval_day_to_second field set changed")

	// 3 scalars + 1 Precision field standing in for the 2-arm oneof.
	require.Equal(t, len(wantScalars)+1, reflect.TypeOf(types.IntervalDayToSecond{}).NumField(),
		"types.IntervalDayToSecond field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		if f.ContainingOneof() != nil {
			assert.Equal(t, protoreflect.Name("precision_mode"), f.ContainingOneof().Name(), "%s oneof", f.Name())
			w, ok := wantOneof[f.Name()]
			require.Truef(t, ok, "unexpected oneof spec field %q", f.Name())
			assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
			assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
			continue
		}
		w, ok := wantScalars[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

// Round-trip domain->proto->domain always uses the non-deprecated precision arm, so a swapped or
// dropped field fails here.
func TestIntervalDayToSecondRoundTrip(t *testing.T) {
	t.Run("precision round trip", func(t *testing.T) {
		in := &types.IntervalDayToSecond{Days: 1, Seconds: 2, Subseconds: 3, Precision: types.PrecisionNanoSeconds}
		p := wire.IntervalDayToSecondToProto(in)
		require.IsType(t, &proto.Expression_Literal_IntervalDayToSecond_Precision{}, p.PrecisionMode)
		got, err := wire.IntervalDayToSecondFromProto(p)
		require.NoError(t, err)
		assert.Equal(t, in, got)
	})

	// The deprecated microseconds precision_mode arm is normalized to microsecond precision on decode.
	t.Run("deprecated microseconds normalized", func(t *testing.T) {
		p := &proto.Expression_Literal_IntervalDayToSecond{
			Days:          4,
			Seconds:       5,
			PrecisionMode: &proto.Expression_Literal_IntervalDayToSecond_Microseconds{Microseconds: 7},
		}
		got, err := wire.IntervalDayToSecondFromProto(p)
		require.NoError(t, err)
		assert.Equal(t,
			&types.IntervalDayToSecond{Days: 4, Seconds: 5, Subseconds: 7, Precision: types.PrecisionMicroSeconds},
			got)
	})

	// An absent precision_mode is malformed: subseconds carries no scale without a precision, so
	// decoding must fail rather than guess.
	t.Run("absent precision_mode is rejected", func(t *testing.T) {
		p := &proto.Expression_Literal_IntervalDayToSecond{Days: 1, Seconds: 2, Subseconds: 3}
		_, err := wire.IntervalDayToSecondFromProto(p)
		assert.Error(t, err)
	})

	assert.Nil(t, wire.IntervalDayToSecondToProto(nil))
	got, err := wire.IntervalDayToSecondFromProto(nil)
	require.NoError(t, err)
	assert.Nil(t, got)

	// GetPrecisionProtoVal is nil-safe and yields the precision's protobuf value.
	assert.Zero(t, (*types.IntervalDayToSecond)(nil).GetPrecisionProtoVal())
	assert.EqualValues(t, 9, (&types.IntervalDayToSecond{Precision: types.PrecisionNanoSeconds}).GetPrecisionProtoVal())
}
