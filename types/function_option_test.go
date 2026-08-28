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

func TestFunctionOptionMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"name":       {1, protoreflect.StringKind},
		"preference": {2, protoreflect.StringKind},
	}

	fields := (&proto.FunctionOption{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec function option field set changed")
	require.Equal(t, len(want), reflect.TypeOf(types.FunctionOption{}).NumField(), "types.FunctionOption field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

func TestFunctionOptionsRoundTrip(t *testing.T) {
	// a nil slice stays nil rather than becoming empty
	assert.Nil(t, types.FunctionOptionsToProto(nil))
	assert.Nil(t, types.FunctionOptionsFromProto(nil))

	opts := []*types.FunctionOption{
		{Name: "rounding", Preference: []string{"TIE_TO_EVEN", "TIE_AWAY_FROM_ZERO"}},
		{Name: "overflow"},
	}
	pb := types.FunctionOptionsToProto(opts)
	require.Len(t, pb, len(opts))
	assert.Equal(t, "rounding", pb[0].Name)
	assert.Equal(t, []string{"TIE_TO_EVEN", "TIE_AWAY_FROM_ZERO"}, pb[0].Preference)
	assert.Equal(t, opts, types.FunctionOptionsFromProto(pb))
}
