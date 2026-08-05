// SPDX-License-Identifier: Apache-2.0

// The only file here that calls a core conversion, so it goes away with them.

package codec_test

import (
	"encoding/hex"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	protobuf "google.golang.org/protobuf/proto"
)

// coreWire is the encode path codec is replacing.
func coreWire(t *testing.T, n types.Nullability) []byte {
	t.Helper()
	b, err := protobuf.MarshalOptions{Deterministic: true}.Marshal(types.TypeToProto(&types.BooleanType{Nullability: n}))
	require.NoError(t, err)
	return b
}

// TestNullabilityMatchesCore checks codec against the core over the whole range the domain type
// can hold, not just the values the spec names, since both sides are meant to be casts.
func TestNullabilityMatchesCore(t *testing.T) {
	// Every golden row, plus values the golden does not carry.
	for _, n := range append(slices.Clone(nullabilityValues), 7, 127) {
		t.Run(strconv.Itoa(int(n)), func(t *testing.T) {
			wire := coreWire(t, n)
			assert.Equal(t, hex.EncodeToString(wire), hex.EncodeToString(codecWire(t, n)))

			var decoded proto.Type
			require.NoError(t, protobuf.Unmarshal(wire, &decoded))
			assert.Equal(t, types.TypeFromProto(&decoded).GetNullability(),
				codec.NullabilityFromProto(decoded.GetBool().GetNullability()))
		})
	}
}
