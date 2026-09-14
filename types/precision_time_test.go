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

func TestPrecisionTimeMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"precision": {1, protoreflect.Int32Kind},
		"value":     {2, protoreflect.Int64Kind},
	}

	fields := (&proto.Expression_Literal_PrecisionTime{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec precision_time field set changed")
	require.Equal(t, len(want), reflect.TypeOf(types.PrecisionTime{}).NumField(), "types.PrecisionTime field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

func TestPrecisionTimestampMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"precision": {1, protoreflect.Int32Kind},
		"value":     {2, protoreflect.Int64Kind},
	}

	fields := (&proto.Expression_Literal_PrecisionTimestamp{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec precision_timestamp field set changed")
	require.Equal(t, len(want), reflect.TypeOf(types.PrecisionTimestamp{}).NumField(), "types.PrecisionTimestamp field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

func TestPrecisionTimestampTzMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"precision": {1, protoreflect.Int32Kind},
		"value":     {2, protoreflect.Int64Kind},
	}

	fields := (&proto.Expression_Literal_PrecisionTimestamp{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec precision_timestamp_tz field set changed")
	require.Equal(t, len(want), reflect.TypeOf(types.PrecisionTimestampTz{}).NumField(), "types.PrecisionTimestampTz field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}
