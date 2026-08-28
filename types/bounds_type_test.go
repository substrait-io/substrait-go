// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestBoundsTypeString(t *testing.T) {
	for _, td := range []struct {
		b        types.BoundsType
		expected string
	}{
		{types.BoundsTypeUnspecified, "BOUNDS_TYPE_UNSPECIFIED"},
		{types.BoundsTypeRows, "BOUNDS_TYPE_ROWS"},
		{types.BoundsTypeRange, "BOUNDS_TYPE_RANGE"},
		{types.BoundsType(99), "99"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.b.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.b))
		})
	}
}

func TestBoundsTypeMatchesProto(t *testing.T) {
	cases := []struct {
		domain types.BoundsType
		pb     proto.Expression_WindowFunction_BoundsType
	}{
		{types.BoundsTypeUnspecified, proto.Expression_WindowFunction_BOUNDS_TYPE_UNSPECIFIED},
		{types.BoundsTypeRows, proto.Expression_WindowFunction_BOUNDS_TYPE_ROWS},
		{types.BoundsTypeRange, proto.Expression_WindowFunction_BOUNDS_TYPE_RANGE},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.Expression_WindowFunction_BOUNDS_TYPE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto bounds type value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
