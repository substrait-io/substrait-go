// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Shared by every slice's golden. Declared per slice, the flag registers twice and the test
// binary panics at init.
var update = flag.Bool("update", false, "rewrite testdata golden files")

// goldenRecord is one golden row: a domain value and the bytes it serializes to.
type goldenRecord struct {
	name        string
	domainValue int32
	wire        []byte
}

func readGolden(t *testing.T, path string) []goldenRecord {
	t.Helper()
	raw, err := os.ReadFile(path)
	require.NoError(t, err)

	var records []goldenRecord
	for i, line := range strings.Split(strings.TrimRight(string(raw), "\n"), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		require.Lenf(t, fields, 3, "%s:%d is not \"name domainValue wireHex\"", path, i+1)
		domainValue, err := strconv.ParseInt(fields[1], 10, 32)
		require.NoError(t, err)
		wire, err := hex.DecodeString(fields[2])
		require.NoError(t, err)
		records = append(records, goldenRecord{name: fields[0], domainValue: int32(domainValue), wire: wire})
	}
	require.NotEmpty(t, records)
	return records
}

// enumGolden is the harness for an enum slice. Every enum conversion is a cast over int32, so the
// checks are the same each time and only the message carrying the value changes.
type enumGolden[T ~int32] struct {
	// path is the golden file. note is its one slice-specific header line.
	path, note string
	// values is the golden's row order.
	values []T
	// encode wraps one value in the smallest spec message that carries it, and marshals that.
	encode func(*testing.T, T) []byte
	// decodeWire returns the value the bytes carry, read off the decoded message rather than
	// through the conversion.
	decodeWire func(*testing.T, []byte) int32
	// decodeCodec is the decode half of the conversion under test.
	decodeCodec func(int32) int32
	// spec and constants pin the hand written constants to the generated descriptor.
	spec      protoreflect.EnumValueDescriptors
	constants map[protoreflect.Name]T
}

func (g enumGolden[T]) run(t *testing.T) {
	if *update {
		g.write(t)
		return
	}
	t.Run("GoldenMatchesCodec", g.goldenMatchesCodec)
	t.Run("GoldenWireCarriesTheValue", g.goldenWireCarriesTheValue)
	t.Run("CoversTheSpec", g.coversTheSpec)
}

func (g enumGolden[T]) write(t *testing.T) {
	t.Helper()
	var buf bytes.Buffer
	buf.WriteString("# name domainValue wireHex\n")
	fmt.Fprintf(&buf, "# %s\n", g.note)
	buf.WriteString("# Rows named for a bare number are outside the spec's enum and are cast unchanged.\n")
	for _, v := range g.values {
		fmt.Fprintf(&buf, "%v %d %s\n", v, int32(v), hex.EncodeToString(g.encode(t, v)))
	}
	require.NoError(t, os.WriteFile(g.path, buf.Bytes(), 0o644))
}

// The golden is a snapshot, so a later change to the conversion shows up as a diff.
func (g enumGolden[T]) goldenMatchesCodec(t *testing.T) {
	records := readGolden(t, g.path)
	require.Len(t, records, len(g.values), "the golden does not have a row per value")
	for i, v := range g.values {
		t.Run(fmt.Sprint(v), func(t *testing.T) {
			assert.Equal(t, fmt.Sprint(v), records[i].name)
			assert.EqualValues(t, v, records[i].domainValue)
			assert.Equal(t, hex.EncodeToString(records[i].wire), hex.EncodeToString(g.encode(t, v)))
		})
	}
}

// Reading the wire column off the decoded message keeps this independent of the conversion, so a
// drift in both directions cannot regenerate the golden with wrong bytes and still pass. The
// second assertion is what catches a decode that stops preserving values.
func (g enumGolden[T]) goldenWireCarriesTheValue(t *testing.T) {
	for _, record := range readGolden(t, g.path) {
		t.Run(record.name, func(t *testing.T) {
			onWire := g.decodeWire(t, record.wire)
			assert.EqualValues(t, record.domainValue, onWire, "golden wire bytes do not carry domainValue")
			assert.EqualValues(t, record.domainValue, g.decodeCodec(onWire))
		})
	}
}

// Fails if the spec gains a value, which is where a hand written enum quietly stops matching.
func (g enumGolden[T]) coversTheSpec(t *testing.T) {
	byName := map[string]bool{}
	for _, record := range readGolden(t, g.path) {
		byName[record.name] = true
	}

	require.Equal(t, g.spec.Len(), len(g.constants), "the set of spec values changed")
	for i := 0; i < g.spec.Len(); i++ {
		v := g.spec.Get(i)
		ourValue, ok := g.constants[v.Name()]
		require.Truef(t, ok, "no constant covers %s", v.Name())
		assert.EqualValues(t, v.Number(), ourValue, "wire number for %s", v.Name())
		assert.Equal(t, string(v.Name()), fmt.Sprint(ourValue), "String() for %s", v.Name())
		assert.Containsf(t, byName, string(v.Name()), "the golden has no row for %s", v.Name())
	}
}
