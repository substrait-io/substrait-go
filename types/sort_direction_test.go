// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
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
