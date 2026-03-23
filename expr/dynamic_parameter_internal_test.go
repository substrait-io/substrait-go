// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"testing"

	"github.com/substrait-io/substrait-go/v8/types"
)

func TestDynamicParameterIsRootRef(t *testing.T) {
	dp := &DynamicParameter{
		OutputType:         &types.Int32Type{Nullability: types.NullabilityRequired},
		ParameterReference: 0,
	}

	// Verify DynamicParameter satisfies RootRefType
	var _ RootRefType = dp
	dp.isRootRef()
}
