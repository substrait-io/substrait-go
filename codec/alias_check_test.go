// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// A core that still aliases a domain type to its wire enum gives the matching
// conversion the same input and output type, so it compiles and does nothing.
// Listing both halves as type switch cases is then a duplicate case, which does
// not compile.
func TestDomainTypesAreDistinctFromTheWireTypes(t *testing.T) {
	var v any
	switch v.(type) {
	case types.Nullability:
	case proto.Type_Nullability:
	}
}
