// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"encoding/hex"
	"flag"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
