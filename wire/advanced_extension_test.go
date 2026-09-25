// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/extensions"
)

// Round-trip through the proto boundary so a dropped/mismapped field fails here.
func TestAdvancedExtensionRoundTrip(t *testing.T) {
	d := &extensions.AdvancedExtension{
		Optimizations: []*extensions.Optimization{{TypeUrl: "opt", Value: []byte("o")}},
		Enhancement:   &extensions.Enhancement{TypeUrl: "enh", Value: []byte("e")},
	}
	assert.Equal(t, d, advancedExtensionFromProto(advancedExtensionToProto(d)))

	assert.Nil(t, advancedExtensionToProto(nil))
	assert.Nil(t, advancedExtensionFromProto(nil))

	// Absent optimizations stay nil through the boundary rather than becoming an
	// empty slice.
	noOpt := &extensions.AdvancedExtension{Enhancement: &extensions.Enhancement{TypeUrl: "enh"}}
	assert.Nil(t, advancedExtensionFromProto(advancedExtensionToProto(noOpt)).Optimizations)
}
