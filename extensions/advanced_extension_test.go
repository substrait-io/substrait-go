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
	"google.golang.org/protobuf/types/known/anypb"
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
		Optimization: []*anypb.Any{{TypeUrl: "opt", Value: []byte("o")}},
		Enhancement:  &anypb.Any{TypeUrl: "enh", Value: []byte("e")},
	}
	assert.Equal(t, d, extensions.AdvancedExtensionFromProto(extensions.AdvancedExtensionToProto(d)))

	assert.Nil(t, extensions.AdvancedExtensionToProto(nil))
	assert.Nil(t, extensions.AdvancedExtensionFromProto(nil))
}

func TestAdvancedExtensionGetters(t *testing.T) {
	// nil receiver is safe and returns zero values.
	var nilExt *extensions.AdvancedExtension
	assert.Nil(t, nilExt.GetOptimization())
	assert.Nil(t, nilExt.GetEnhancement())

	opt := []*anypb.Any{{TypeUrl: "opt"}}
	enh := &anypb.Any{TypeUrl: "enh"}
	ext := &extensions.AdvancedExtension{Optimization: opt, Enhancement: enh}
	assert.Equal(t, opt, ext.GetOptimization())
	assert.Equal(t, enh, ext.GetEnhancement())
}
