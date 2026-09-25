// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/extensions"
)

func TestAdvancedExtensionGetters(t *testing.T) {
	// nil receiver is safe and returns zero values.
	var nilExt *extensions.AdvancedExtension
	assert.Nil(t, nilExt.GetOptimizations())
	assert.Nil(t, nilExt.GetEnhancement())

	opt := []*extensions.Optimization{{TypeUrl: "opt"}}
	enh := &extensions.Enhancement{TypeUrl: "enh"}
	ext := &extensions.AdvancedExtension{Optimizations: opt, Enhancement: enh}
	assert.Equal(t, opt, ext.GetOptimizations())
	assert.Equal(t, enh, ext.GetEnhancement())
}
