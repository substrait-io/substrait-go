// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/wire"
	substraitpb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestBoundFromProto(t *testing.T) {
	for _, tc := range []struct {
		name        string
		proto       *substraitpb.Expression_WindowFunction_Bound
		expected    expr.Bound
		expectedStr string
	}{
		{
			name: "nil",
		},
		{
			name:  "nil kind",
			proto: &substraitpb.Expression_WindowFunction_Bound{},
		},
		{
			name: "unbounded",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_Unbounded_{},
			},
			expected:    expr.Unbounded{},
			expectedStr: "UNBOUNDED",
		},
		{
			name: "current row",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_CurrentRow_{},
			},
			expected:    expr.CurrentRow{},
			expectedStr: "CURRENT ROW",
		},
		{
			name: "preceding 42",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_Preceding_{
					Preceding: &substraitpb.Expression_WindowFunction_Bound_Preceding{
						Offset: 42,
					},
				},
			},
			expected:    expr.PrecedingBound(42),
			expectedStr: "42 PRECEDING",
		},
		{
			name: "following 42",
			proto: &substraitpb.Expression_WindowFunction_Bound{
				Kind: &substraitpb.Expression_WindowFunction_Bound_Following_{
					Following: &substraitpb.Expression_WindowFunction_Bound_Following{
						Offset: 42,
					},
				},
			},
			expected:    expr.FollowingBound(42),
			expectedStr: "42 FOLLOWING",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bound := wire.BoundFromProto(tc.proto)
			require.Equal(t, tc.expected, bound)
			if bound != nil {
				require.Equal(t, tc.expectedStr, bound.String())
			}
		})
	}
}
