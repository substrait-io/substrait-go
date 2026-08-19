// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// pin the hand written constants to the generated Type.Nullability descriptor
func TestNullabilityMatchesDescriptor(t *testing.T) {
	ours := map[protoreflect.Name]types.Nullability{
		"NULLABILITY_UNSPECIFIED": types.NullabilityUnspecified,
		"NULLABILITY_NULLABLE":    types.NullabilityNullable,
		"NULLABILITY_REQUIRED":    types.NullabilityRequired,
	}

	values := proto.Type_Nullability(0).Descriptor().Values()
	require.Equal(t, values.Len(), len(ours), "the set of spec nullability values changed")

	for i := 0; i < values.Len(); i++ {
		v := values.Get(i)
		ourValue, ok := ours[v.Name()]
		require.Truef(t, ok, "no types.Nullability constant covers %s", v.Name())
		assert.EqualValues(t, v.Number(), ourValue, "wire number for %s", v.Name())
		assert.Equal(t, string(v.Name()), ourValue.String(), "String() for %s", v.Name())
	}
}
