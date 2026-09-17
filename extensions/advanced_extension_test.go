// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestAdvancedExtensionMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"optimization": {1, protoreflect.MessageKind},
		"enhancement":  {2, protoreflect.MessageKind},
	}

	fields := (&extensionspb.AdvancedExtension{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec advanced extension field set changed")
	require.Equal(t, len(want), reflect.TypeOf(extensions.AdvancedExtension{}).NumField(), "extensions.AdvancedExtension field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}

// Round-trip through the proto boundary so a dropped/mismapped field fails here.
func TestAdvancedExtensionRoundTrip(t *testing.T) {
	d := &extensions.AdvancedExtension{
		Optimizations: []*extensions.Optimization{{TypeUrl: "opt", Value: []byte("o")}},
		Enhancement:   &extensions.Enhancement{TypeUrl: "enh", Value: []byte("e")},
	}
	assert.Equal(t, d, extensions.AdvancedExtensionFromProto(extensions.AdvancedExtensionToProto(d)))

	assert.Nil(t, extensions.AdvancedExtensionToProto(nil))
	assert.Nil(t, extensions.AdvancedExtensionFromProto(nil))

	// Absent optimizations stay nil through the boundary rather than becoming an
	// empty slice.
	noOpt := &extensions.AdvancedExtension{Enhancement: &extensions.Enhancement{TypeUrl: "enh"}}
	assert.Nil(t, extensions.AdvancedExtensionFromProto(extensions.AdvancedExtensionToProto(noOpt)).Optimizations)
}

func TestAdvancedExtensionGetters(t *testing.T) {
	// nil receiver is safe and returns zero values.
	var nilExt *extensions.AdvancedExtension
	assert.Nil(t, nilExt.GetOptimizations())
	assert.Nil(t, nilExt.GetEnhancement())

	opt := []*extensions.Optimization{{TypeUrl: "opt"}}
	enh := &extensions.Enhancement{TypeUrl: "enh"}
	ext := &extensions.AdvancedExtension{Optimizations: opt, Enhancement: enh}
	assert.Equal(t, opt, ext.GetOptimizations())
	assert.Equal(t, enh, ext.GetEnhancement())
}
