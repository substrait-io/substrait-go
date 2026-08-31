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

func TestDecimalMatchesDescriptor(t *testing.T) {
	want := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
	}{
		"value":     {1, protoreflect.BytesKind},
		"precision": {2, protoreflect.Int32Kind},
		"scale":     {3, protoreflect.Int32Kind},
	}

	fields := (&proto.Expression_Literal_Decimal{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(want), fields.Len(), "spec decimal field set changed")
	require.Equal(t, len(want), reflect.TypeOf(types.Decimal{}).NumField(), "types.Decimal field count drifted from the spec")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		w, ok := want[f.Name()]
		require.Truef(t, ok, "unexpected spec field %q", f.Name())
		assert.EqualValues(t, w.number, f.Number(), "%s wire number", f.Name())
		assert.Equal(t, w.kind, f.Kind(), "%s kind", f.Name())
	}
}
