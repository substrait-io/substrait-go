// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v8/types"
)

func TestVersionGetters(t *testing.T) {
	// distinct values so a getter reading the wrong field shows up
	v := &types.Version{
		MajorNumber: 1,
		MinorNumber: 29,
		PatchNumber: 3,
		GitHash:     "0123456789abcdef",
		Producer:    "substrait-go",
	}

	assert.Equal(t, uint32(1), v.GetMajorNumber())
	assert.Equal(t, uint32(29), v.GetMinorNumber())
	assert.Equal(t, uint32(3), v.GetPatchNumber())
	assert.Equal(t, "0123456789abcdef", v.GetGitHash())
	assert.Equal(t, "substrait-go", v.GetProducer())
}

// Plan.Version() hands out a nil *Version through an interface, so callers reach the getters
// without a way to check for nil first.
func TestVersionGettersOnNil(t *testing.T) {
	var v *types.Version

	assert.Zero(t, v.GetMajorNumber())
	assert.Zero(t, v.GetMinorNumber())
	assert.Zero(t, v.GetPatchNumber())
	assert.Empty(t, v.GetGitHash())
	assert.Empty(t, v.GetProducer())
	assert.Equal(t, "<nil>", v.String())
}

func TestVersionString(t *testing.T) {
	for _, td := range []struct {
		name     string
		version  *types.Version
		expected string
	}{
		{"numbers only", &types.Version{MajorNumber: 0, MinorNumber: 29, PatchNumber: 0}, "0.29.0"},
		{"zero value", &types.Version{}, "0.0.0"},
		{"with producer", &types.Version{MinorNumber: 29, Producer: "substrait-go v8"}, "0.29.0 (substrait-go v8)"},
		{"with git hash", &types.Version{MinorNumber: 29, GitHash: "abc123"}, "0.29.0+abc123"},
		{
			"everything",
			&types.Version{MajorNumber: 1, MinorNumber: 2, PatchNumber: 3, GitHash: "abc123", Producer: "acme"},
			"1.2.3+abc123 (acme)",
		},
		{"nil", nil, "<nil>"},
	} {
		t.Run(td.name, func(t *testing.T) {
			assert.Equal(t, td.expected, td.version.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.version))
		})
	}
}
