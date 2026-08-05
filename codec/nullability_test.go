// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	protobuf "google.golang.org/protobuf/proto"
)

const nullabilityGolden = "testdata/nullability.golden"

// nullabilityValues is the golden's row order. Nullability is an int32, so a caller can hold a
// number the spec has not defined and a plan from a newer spec version can carry one; without
// those rows the golden cannot tell a value-preserving conversion from one that clamps.
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

func renderNullabilityGolden(t *testing.T) string {
	t.Helper()
	var buf bytes.Buffer
	buf.WriteString("# name domainValue wireHex\n")
	buf.WriteString("# substrait.Type{bool:{nullability}} encoded through codec.\n")
	buf.WriteString("# Rows named for a bare number are outside the spec's enum and are cast unchanged.\n")
	for _, n := range nullabilityValues {
		fmt.Fprintf(&buf, "%s %d %s\n", n, int32(n), hex.EncodeToString(codecWire(t, n)))
	}
	return buf.String()
}

// TestNullabilityGoldenMatchesCodec compares the committed bytes against what codec encodes now.
// The golden is a snapshot, so a later change to the conversion shows up as a diff.
func TestNullabilityGoldenMatchesCodec(t *testing.T) {
	if *update {
		require.NoError(t, os.WriteFile(nullabilityGolden, []byte(renderNullabilityGolden(t)), 0o644))
		return
	}

	records := readGolden(t, nullabilityGolden)
	require.Len(t, records, len(nullabilityValues), "the golden does not have a row per value")
	for i, n := range nullabilityValues {
		t.Run(n.String(), func(t *testing.T) {
			assert.Equal(t, n.String(), records[i].name)
			assert.EqualValues(t, n, records[i].domainValue)
			assert.Equal(t, hex.EncodeToString(records[i].wire), hex.EncodeToString(codecWire(t, n)))
		})
	}
}

func TestNullabilityFromProtoRoundTrip(t *testing.T) {
	for _, record := range readGolden(t, nullabilityGolden) {
		t.Run(record.name, func(t *testing.T) {
			var decoded proto.Type
			require.NoError(t, protobuf.Unmarshal(record.wire, &decoded))
			// The wire column has to be checked against the domain column directly. Decoding it
			// through the conversion only proves the pair agrees with itself, so a symmetric drift
			// would regenerate the golden with wrong bytes and still pass.
			n := decoded.GetBool().GetNullability()
			assert.EqualValues(t, record.domainValue, n, "golden wire bytes do not carry domainValue")
			assert.EqualValues(t, record.domainValue, codec.NullabilityFromProto(n))
		})
	}
}

// TestNullabilityGoldenCoversDescriptor fails if the spec gains a nullability value the golden has
// no row for.
func TestNullabilityGoldenCoversDescriptor(t *testing.T) {
	byName := map[string]goldenRecord{}
	for _, record := range readGolden(t, nullabilityGolden) {
		byName[record.name] = record
	}

	values := proto.Type_Nullability(0).Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		v := values.Get(i)
		record, ok := byName[string(v.Name())]
		require.Truef(t, ok, "the golden has no row for %s", v.Name())

		var decoded proto.Type
		require.NoError(t, protobuf.Unmarshal(record.wire, &decoded))
		assert.EqualValues(t, v.Number(), decoded.GetBool().GetNullability(), "wire number for %s", v.Name())
		assert.EqualValues(t, v.Number(), record.domainValue, "domain value for %s", v.Name())
	}
}
