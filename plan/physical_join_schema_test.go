// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/types"
)

func TestNullableRecordTypePreservesNestedFields(t *testing.T) {
	child := &types.Int64Type{Nullability: types.NullabilityRequired}
	nested := &types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{child}}
	input := types.NewRecordTypeFromTypes([]types.Type{nested})

	output := nullableRecordType(*input)
	assert.Equal(t, &types.StructType{
		Nullability: types.NullabilityNullable,
		Types:       []types.Type{&types.Int64Type{Nullability: types.NullabilityRequired}},
	}, output.Types()[0])
	assert.Equal(t, types.NullabilityRequired, nested.Nullability)
	assert.Equal(t, types.NullabilityRequired, child.Nullability)
}
