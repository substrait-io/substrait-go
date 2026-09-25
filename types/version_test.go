// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestVersionString(t *testing.T) {
	for _, td := range []struct {
		name     string
		version  *types.Version
		expected string
	}{
		{"numbers only", &types.Version{MinorNumber: 29}, "0.29.0"},
		{"zero value", &types.Version{}, "0.0.0"},
		{"with producer", &types.Version{MinorNumber: 29, Producer: "substrait-go v8"}, "0.29.0 (substrait-go v8)"},
		{"with git hash", &types.Version{MinorNumber: 29, GitHash: "abc123"}, "0.29.0+abc123"},
		{
			"everything",
			&types.Version{MajorNumber: 1, MinorNumber: 2, PatchNumber: 3, GitHash: "abc123", Producer: "acme"},
			"1.2.3+abc123 (acme)",
		},
	} {
		t.Run(td.name, func(t *testing.T) {
			assert.Equal(t, td.expected, td.version.String())
		})
	}
}
