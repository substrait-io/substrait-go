// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const nullabilityGolden = "testdata/nullability.golden"

// nullabilityValues is the golden's row order. Nullability is an int32, so a plan from a newer spec
// version can carry a number this one has not defined; without those rows the golden cannot tell a
// value-preserving conversion from one that clamps.
var nullabilityValues = []types.Nullability{
	types.NullabilityUnspecified,
	types.NullabilityNullable,
	types.NullabilityRequired,
	types.Nullability(3),
	types.Nullability(-1),
	types.Nullability(math.MinInt32),
	types.Nullability(math.MaxInt32),
}

// codecWire encodes one nullability through codec. substrait.Type carrying a boolean is the
// smallest spec message with a nullability field.
func codecWire(t *testing.T, n types.Nullability) []byte {
	t.Helper()
	msg := &proto.Type{Kind: &proto.Type_Bool{
		Bool: &proto.Type_Boolean{Nullability: codec.NullabilityToProto(n)}}}
	b, err := protobuf.MarshalOptions{Deterministic: true}.Marshal(msg)
	require.NoError(t, err)
	return b
}

func TestNullability(t *testing.T) {
	enumGolden[types.Nullability]{
		path:   nullabilityGolden,
		note:   "substrait.Type{bool:{nullability}} encoded through codec.",
		values: nullabilityValues,
		encode: codecWire,
		decodeWire: func(t *testing.T, b []byte) int32 {
			var decoded proto.Type
			require.NoError(t, protobuf.Unmarshal(b, &decoded))
			return int32(decoded.GetBool().GetNullability())
		},
		decodeCodec: func(v int32) int32 {
			return int32(codec.NullabilityFromProto(proto.Type_Nullability(v)))
		},
		spec: proto.Type_Nullability(0).Descriptor().Values(),
		constants: map[protoreflect.Name]types.Nullability{
			"NULLABILITY_UNSPECIFIED": types.NullabilityUnspecified,
			"NULLABILITY_NULLABLE":    types.NullabilityNullable,
			"NULLABILITY_REQUIRED":    types.NullabilityRequired,
		},
	}.run(t)
}
