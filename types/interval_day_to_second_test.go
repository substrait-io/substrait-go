// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestIntervalDayToSecondMatchesDescriptor(t *testing.T) {
	// Scalar fields are pinned by wire number and kind; the precision_mode oneof is pinned
	// separately since the domain type collapses its two arms into a single PrecisionMode field.
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

	// 3 scalars + 1 PrecisionMode field standing in for the 2-arm oneof.
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

// Round-trip both precision_mode arms domain->proto->domain, asserting each arm maps to the
// correct proto oneof field so a swapped or dropped arm fails here.
func TestIntervalDayToSecondRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name    string
		in      *types.IntervalDayToSecond
		checkPB func(t *testing.T, p *proto.Expression_Literal_IntervalDayToSecond)
	}{
		{
			name: "precision arm",
			in:   &types.IntervalDayToSecond{Days: 1, Seconds: 2, Subseconds: 3, PrecisionMode: types.IntervalDayToSecondPrecision(9)},
			checkPB: func(t *testing.T, p *proto.Expression_Literal_IntervalDayToSecond) {
				require.IsType(t, &proto.Expression_Literal_IntervalDayToSecond_Precision{}, p.PrecisionMode)
				assert.EqualValues(t, 9, p.GetPrecision())
			},
		},
		{
			name: "microseconds arm",
			in:   &types.IntervalDayToSecond{Days: 4, Seconds: 5, PrecisionMode: types.IntervalDayToSecondMicroseconds(7)},
			checkPB: func(t *testing.T, p *proto.Expression_Literal_IntervalDayToSecond) {
				require.IsType(t, &proto.Expression_Literal_IntervalDayToSecond_Microseconds{}, p.PrecisionMode)
				assert.EqualValues(t, 7, p.GetMicroseconds())
			},
		},
		{
			name: "no precision mode",
			in:   &types.IntervalDayToSecond{Days: 1, Seconds: 2, Subseconds: 3},
			checkPB: func(t *testing.T, p *proto.Expression_Literal_IntervalDayToSecond) {
				assert.Nil(t, p.PrecisionMode)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := types.IntervalDayToSecondToProto(tc.in)
			tc.checkPB(t, p)
			assert.Equal(t, tc.in, types.IntervalDayToSecondFromProto(p))
		})
	}

	assert.Nil(t, types.IntervalDayToSecondToProto(nil))
	assert.Nil(t, types.IntervalDayToSecondFromProto(nil))

	// GetPrecision is nil-safe and yields the exponent only for the Precision arm.
	assert.Zero(t, (*types.IntervalDayToSecond)(nil).GetPrecision())
	assert.EqualValues(t, 9, (&types.IntervalDayToSecond{PrecisionMode: types.IntervalDayToSecondPrecision(9)}).GetPrecision())
	assert.Zero(t, (&types.IntervalDayToSecond{PrecisionMode: types.IntervalDayToSecondMicroseconds(7)}).GetPrecision())
}
