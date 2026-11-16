// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestVariadicDeclarationWithoutParametersDoesNotMatch(t *testing.T) {
	id := extensions.FunctionID{URN: "extension:test:variadic", Name: "without_parameters"}
	variadic := &extensions.VariadicBehavior{Min: 1}
	variants := []struct {
		name    string
		variant extensions.FunctionVariant
	}{
		{"scalar", extensions.NewScalarFuncVariantWithProps(id, variadic, false, true)},
		{"aggregate", extensions.NewAggFuncVariantOpts(id, extensions.AggVariantOptions{Variadic: variadic})},
		{"window", extensions.NewWindowFuncVariantOpts(id, extensions.WindowVariantOpts{Variadic: variadic})},
	}
	for _, tc := range variants {
		t.Run(tc.name, func(t *testing.T) {
			for _, args := range [][]types.Type{nil, {&types.Int64Type{}}} {
				matched, err := tc.variant.Match(args)
				require.NoError(t, err)
				require.False(t, matched)
			}
		})
	}
}
