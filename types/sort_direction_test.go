// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestSortDirectionString(t *testing.T) {
	for _, td := range []struct {
		d        types.SortDirection
		expected string
	}{
		{types.SortUnspecified, "SORT_DIRECTION_UNSPECIFIED"},
		{types.SortAscNullsFirst, "SORT_DIRECTION_ASC_NULLS_FIRST"},
		{types.SortAscNullsLast, "SORT_DIRECTION_ASC_NULLS_LAST"},
		{types.SortDescNullsFirst, "SORT_DIRECTION_DESC_NULLS_FIRST"},
		{types.SortDescNullsLast, "SORT_DIRECTION_DESC_NULLS_LAST"},
		{types.SortClustered, "SORT_DIRECTION_CLUSTERED"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.d.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.d))
		})
	}
}

func TestSortDirectionMatchesProto(t *testing.T) {
	cases := []struct {
		domain types.SortDirection
		pb     proto.SortField_SortDirection
	}{
		{types.SortUnspecified, proto.SortField_SORT_DIRECTION_UNSPECIFIED},
		{types.SortAscNullsFirst, proto.SortField_SORT_DIRECTION_ASC_NULLS_FIRST},
		{types.SortAscNullsLast, proto.SortField_SORT_DIRECTION_ASC_NULLS_LAST},
		{types.SortDescNullsFirst, proto.SortField_SORT_DIRECTION_DESC_NULLS_FIRST},
		{types.SortDescNullsLast, proto.SortField_SORT_DIRECTION_DESC_NULLS_LAST},
		{types.SortClustered, proto.SortField_SORT_DIRECTION_CLUSTERED},
	}
	// fail if the spec adds a value we don't mirror
	assert.Equal(t, proto.SortField_SORT_DIRECTION_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"a proto sort direction value is not covered")
	for _, td := range cases {
		assert.EqualValues(t, td.pb, td.domain, "wire number for %s", td.pb)
		assert.Equal(t, td.pb.String(), td.domain.String(), "enum name for %s", td.pb)
	}
}
