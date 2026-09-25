// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
)

// createJoinInput builds a named table read rel with three required int64 columns.
func createJoinInput(name string) *NamedTableReadRel {
	schema := types.NamedStruct{
		Names: []string{"x", "y", "z"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.Int64Type{}},
		},
	}
	return &NamedTableReadRel{names: []string{name}, baseReadRel: baseReadRel{baseSchema: schema}}
}

func keyRef(t *testing.T, rel Rel, idx int32) *expr.FieldReference {
	t.Helper()
	base := rel.RecordType()
	ref, err := expr.NewRootFieldRef(expr.NewStructFieldRef(idx), &base)
	require.NoError(t, err)
	return ref
}

// All joins below have a left and right input with three int64 columns each.
func joinInputs() (Rel, Rel) {
	return createJoinInput("L"), createJoinInput("R")
}

func eqKeys(t *testing.T, left, right Rel) []*ComparisonJoinKey {
	return []*ComparisonJoinKey{
		NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 2)),
		NewEqualityJoinKey(keyRef(t, left, 1), keyRef(t, right, 0)),
	}
}

// The deprecated accessors derive their values from the keys.
func TestJoinDeprecatedKeyAccessors(t *testing.T) {
	left, right := joinInputs()
	keys := eqKeys(t, left, right)
	wantLeft := []*expr.FieldReference{keyRef(t, left, 0), keyRef(t, left, 1)}
	wantRight := []*expr.FieldReference{keyRef(t, right, 2), keyRef(t, right, 0)}

	hash := &HashJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}
	assert.Equal(t, wantLeft, hash.LeftKeys())
	assert.Equal(t, wantRight, hash.RightKeys())

	merge := &MergeJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}
	assert.Equal(t, wantLeft, merge.LeftKeys())
	assert.Equal(t, wantRight, merge.RightKeys())
}
