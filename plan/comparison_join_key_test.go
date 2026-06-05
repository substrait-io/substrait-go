// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
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

func joinTestRegistry() expr.ExtensionRegistry {
	return expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
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

// Producing a join always writes the new keys field, and additionally
// mirrors the deprecated left_keys/right_keys when every key is a plain EQ
// comparison so old consumers keep working for equality joins.
func TestJoinWritesLegacyKeysForEqualityOnly(t *testing.T) {
	left, right := joinInputs()
	wantLeft := []*proto.Expression_FieldReference{
		keyRef(t, left, 0).ToProtoFieldRef(), keyRef(t, left, 1).ToProtoFieldRef()}
	wantRight := []*proto.Expression_FieldReference{
		keyRef(t, right, 2).ToProtoFieldRef(), keyRef(t, right, 0).ToProtoFieldRef()}

	// All-EQ keys: new keys field plus mirrored deprecated fields.
	t.Run("all EQ mirrors deprecated fields", func(t *testing.T) {
		keys := eqKeys(t, left, right)

		hash := (&HashJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}).ToProto().GetHashJoin()
		assert.Len(t, hash.GetKeys(), 2)
		assert.Empty(t, cmp.Diff(wantLeft, hash.GetLeftKeys(), protocmp.Transform()))
		assert.Empty(t, cmp.Diff(wantRight, hash.GetRightKeys(), protocmp.Transform()))

		merge := (&MergeJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}).ToProto().GetMergeJoin()
		assert.Len(t, merge.GetKeys(), 2)
		assert.Empty(t, cmp.Diff(wantLeft, merge.GetLeftKeys(), protocmp.Transform()))
		assert.Empty(t, cmp.Diff(wantRight, merge.GetRightKeys(), protocmp.Transform()))
	})

	// A non-EQ comparison anywhere makes the legacy fields lossy, so they are
	// omitted entirely and only the new keys field is written.
	for _, tc := range []struct {
		name       string
		comparison JoinKeyComparison
	}{
		{"is not distinct from", SimpleComparison{Type: SimpleComparisonTypeIsNotDistinctFrom}},
		{"might equal", SimpleComparison{Type: SimpleComparisonTypeMightEqual}},
		{"custom", CustomComparison{FunctionReference: 7}},
	} {
		t.Run("non-EQ omits deprecated fields: "+tc.name, func(t *testing.T) {
			keys := []*ComparisonJoinKey{
				NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 2)),
				NewComparisonJoinKey(keyRef(t, left, 1), keyRef(t, right, 0), tc.comparison),
			}

			hash := (&HashJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}).ToProto().GetHashJoin()
			assert.Len(t, hash.GetKeys(), 2)
			assert.Empty(t, hash.GetLeftKeys())
			assert.Empty(t, hash.GetRightKeys())

			merge := (&MergeJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}).ToProto().GetMergeJoin()
			assert.Len(t, merge.GetKeys(), 2)
			assert.Empty(t, merge.GetLeftKeys())
			assert.Empty(t, merge.GetRightKeys())
		})
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

// A plan from a legacy producer (only deprecated fields set) is consumed and
// mapped to keys with EQ comparisons.
func TestJoinConsumesLegacyKeys(t *testing.T) {
	left, right := joinInputs()
	reg := joinTestRegistry()

	for _, tc := range []struct {
		name  string
		build func(hj *proto.HashJoinRel) *proto.Rel
		keys  func(rel Rel) []*ComparisonJoinKey
	}{
		{
			name:  "hash",
			build: func(hj *proto.HashJoinRel) *proto.Rel { return &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: hj}} },
			keys:  func(rel Rel) []*ComparisonJoinKey { return rel.(*HashJoinRel).Keys() },
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			legacy := tc.build(&proto.HashJoinRel{
				Common: &proto.RelCommon{},
				Left:   left.ToProto(),
				Right:  right.ToProto(),
				Type:   proto.HashJoinRel_JOIN_TYPE_INNER,
				LeftKeys: []*proto.Expression_FieldReference{
					keyRef(t, left, 0).ToProtoFieldRef(), keyRef(t, left, 1).ToProtoFieldRef()},
				RightKeys: []*proto.Expression_FieldReference{
					keyRef(t, right, 2).ToProtoFieldRef(), keyRef(t, right, 0).ToProtoFieldRef()},
			})

			rel, err := RelFromProto(legacy, reg)
			require.NoError(t, err)

			keys := tc.keys(rel)
			require.Len(t, keys, 2)
			for _, k := range keys {
				assert.Equal(t, SimpleComparison{Type: SimpleComparisonTypeEq}, k.Comparison())
			}

			// Re-emitting sets the new keys field and, because these are all
			// EQ comparisons, mirrors the deprecated fields back as well.
			hj := rel.ToProto().GetHashJoin()
			assert.Len(t, hj.GetKeys(), 2)
			assert.Len(t, hj.GetLeftKeys(), 2)
			assert.Len(t, hj.GetRightKeys(), 2)
		})
	}
}

// When both the deprecated fields and the new keys are present, keys wins.
func TestJoinPrefersNewKeysOverDeprecated(t *testing.T) {
	left, right := joinInputs()
	reg := joinTestRegistry()
	keys := eqKeys(t, left, right)

	both := &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
		Common: &proto.RelCommon{},
		Left:   left.ToProto(),
		Right:  right.ToProto(),
		Type:   proto.HashJoinRel_JOIN_TYPE_INNER,
		Keys:   comparisonJoinKeysToProto(keys),
		// Bogus deprecated keys pointing at different fields than the real keys.
		LeftKeys:  []*proto.Expression_FieldReference{keyRef(t, left, 2).ToProtoFieldRef()},
		RightKeys: []*proto.Expression_FieldReference{keyRef(t, right, 1).ToProtoFieldRef()},
	}}}

	rel, err := RelFromProto(both, reg)
	require.NoError(t, err)
	got := rel.ToProto().GetHashJoin().GetKeys()
	if diff := cmp.Diff(comparisonJoinKeysToProto(keys), got, protocmp.Transform()); diff != "" {
		t.Errorf("expected new keys to win, diff:\n%v", diff)
	}
}

// Equality, non-EQ simple comparisons and custom comparison functions all
// survive a round trip through proto.
func TestJoinKeysRoundTrip(t *testing.T) {
	left, right := joinInputs()
	reg := joinTestRegistry()

	keys := []*ComparisonJoinKey{
		NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 2)),
		NewComparisonJoinKey(keyRef(t, left, 1), keyRef(t, right, 0),
			SimpleComparison{Type: SimpleComparisonTypeIsNotDistinctFrom}),
		NewComparisonJoinKey(keyRef(t, left, 2), keyRef(t, right, 1),
			CustomComparison{FunctionReference: 42}),
	}

	for _, tc := range []struct {
		name string
		rel  Rel
	}{
		{"hash", &HashJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}},
		{"merge", &MergeJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.rel.ToProto()
			roundTripped, err := RelFromProto(out, reg)
			require.NoError(t, err)

			if diff := cmp.Diff(out, roundTripped.ToProto(), protocmp.Transform()); diff != "" {
				t.Errorf("join did not round trip, diff:\n%v", diff)
			}

			// The comparison kinds are preserved on the model side.
			var gotKeys []*ComparisonJoinKey
			switch r := roundTripped.(type) {
			case *HashJoinRel:
				gotKeys = r.Keys()
			case *MergeJoinRel:
				gotKeys = r.Keys()
			}
			require.Len(t, gotKeys, 3)
			assert.Equal(t, SimpleComparison{Type: SimpleComparisonTypeEq}, gotKeys[0].Comparison())
			assert.Equal(t, SimpleComparison{Type: SimpleComparisonTypeIsNotDistinctFrom}, gotKeys[1].Comparison())
			assert.Equal(t, CustomComparison{FunctionReference: 42}, gotKeys[2].Comparison())
		})
	}
}

// badFieldRef builds a proto field reference to an out-of-range struct field,
// which fails to resolve against the join inputs' three-column schema.
func badFieldRef() *proto.Expression_FieldReference {
	return &proto.Expression_FieldReference{
		ReferenceType: &proto.Expression_FieldReference_DirectReference{
			DirectReference: &proto.Expression_ReferenceSegment{
				ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
					StructField: &proto.Expression_ReferenceSegment_StructField{Field: 99},
				},
			},
		},
		RootType: &proto.Expression_FieldReference_RootReference_{
			RootReference: &proto.Expression_FieldReference_RootReference{},
		},
	}
}

// joinFields holds the key-bearing fields shared by HashJoinRel and
// MergeJoinRel so an error case can be exercised against both join types.
type joinFields struct {
	keys      []*proto.ComparisonJoinKey
	leftKeys  []*proto.Expression_FieldReference
	rightKeys []*proto.Expression_FieldReference
}

func (f joinFields) hashProto(left, right Rel) *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
		Common: &proto.RelCommon{}, Type: proto.HashJoinRel_JOIN_TYPE_INNER,
		Left: left.ToProto(), Right: right.ToProto(),
		Keys: f.keys, LeftKeys: f.leftKeys, RightKeys: f.rightKeys,
	}}}
}

func (f joinFields) mergeProto(left, right Rel) *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
		Common: &proto.RelCommon{}, Type: proto.MergeJoinRel_JOIN_TYPE_INNER,
		Left: left.ToProto(), Right: right.ToProto(),
		Keys: f.keys, LeftKeys: f.leftKeys, RightKeys: f.rightKeys,
	}}}
}

// Error paths in comparisonJoinKeysFromProto / joinKeyComparisonFromProto,
// exercised against both hash and merge joins.
func TestJoinKeysFromProtoErrors(t *testing.T) {
	left, right := joinInputs()
	reg := joinTestRegistry()
	goodLeft := keyRef(t, left, 0).ToProtoFieldRef()
	goodRight := keyRef(t, right, 0).ToProtoFieldRef()
	eq := SimpleComparison{Type: SimpleComparisonTypeEq}.toProto()

	for _, tc := range []struct {
		name   string
		fields joinFields
	}{
		{
			name:   "keys path: bad left ref",
			fields: joinFields{keys: []*proto.ComparisonJoinKey{{Left: badFieldRef(), Right: goodRight, Comparison: eq}}},
		},
		{
			name:   "keys path: bad right ref",
			fields: joinFields{keys: []*proto.ComparisonJoinKey{{Left: goodLeft, Right: badFieldRef(), Comparison: eq}}},
		},
		{
			name:   "keys path: unset comparison",
			fields: joinFields{keys: []*proto.ComparisonJoinKey{{Left: goodLeft, Right: goodRight, Comparison: &proto.ComparisonJoinKey_ComparisonType{}}}},
		},
		{
			name:   "legacy path: bad left ref",
			fields: joinFields{leftKeys: []*proto.Expression_FieldReference{badFieldRef()}, rightKeys: []*proto.Expression_FieldReference{goodRight}},
		},
		{
			name:   "legacy path: bad right ref",
			fields: joinFields{leftKeys: []*proto.Expression_FieldReference{goodLeft}, rightKeys: []*proto.Expression_FieldReference{badFieldRef()}},
		},
		{
			name:   "legacy path: mismatched lengths",
			fields: joinFields{leftKeys: []*proto.Expression_FieldReference{goodLeft}, rightKeys: []*proto.Expression_FieldReference{}},
		},
	} {
		t.Run("hash/"+tc.name, func(t *testing.T) {
			_, err := RelFromProto(tc.fields.hashProto(left, right), reg)
			require.Error(t, err)
		})
		t.Run("merge/"+tc.name, func(t *testing.T) {
			_, err := RelFromProto(tc.fields.mergeProto(left, right), reg)
			require.Error(t, err)
		})
	}
}

// A merge join whose right input fails to decode surfaces an error.
func TestMergeJoinBadRightInput(t *testing.T) {
	left, _ := joinInputs()
	rel := &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
		Common: &proto.RelCommon{},
		Type:   proto.MergeJoinRel_JOIN_TYPE_INNER,
		Left:   left.ToProto(),
		Right:  &proto.Rel{}, // nil RelType -> RelFromProto errors
	}}}

	_, err := RelFromProto(rel, joinTestRegistry())
	require.ErrorContains(t, err, "right input to MergeJoinRel")
}

// Exercises the join accessors and the ComparisonJoinKey getters, and covers
// the post-join-filter branch of ToProto via a round trip.
func TestJoinAccessorsAndPostJoinFilter(t *testing.T) {
	left, right := joinInputs()
	reg := joinTestRegistry()
	keys := eqKeys(t, left, right)
	postFilter := expr.NewPrimitiveLiteral(true, false)

	k := keys[0]
	assert.Equal(t, keyRef(t, left, 0), k.Left())
	assert.Equal(t, keyRef(t, right, 2), k.Right())
	assert.Equal(t, SimpleComparison{Type: SimpleComparisonTypeEq}, k.Comparison())

	// With no post-join filter set, the accessor returns the default filter.
	noFilter := &HashJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys}
	assert.Equal(t, defFilter, noFilter.PostJoinFilter())

	for _, tc := range []struct {
		name string
		rel  Rel
	}{
		{"hash", &HashJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys, postJoinFilter: postFilter}},
		{"merge", &MergeJoinRel{left: left, right: right, joinType: HashMergeInner, keys: keys, postJoinFilter: postFilter}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			switch r := tc.rel.(type) {
			case *HashJoinRel:
				assert.Equal(t, left, r.Left())
				assert.Equal(t, right, r.Right())
				assert.Equal(t, HashMergeInner, r.Type())
				assert.Equal(t, postFilter, r.PostJoinFilter())
			case *MergeJoinRel:
				assert.Equal(t, left, r.Left())
				assert.Equal(t, right, r.Right())
				assert.Equal(t, HashMergeInner, r.Type())
				assert.Equal(t, postFilter, r.PostJoinFilter())
			}

			out := tc.rel.ToProto()
			roundTripped, err := RelFromProto(out, reg)
			require.NoError(t, err)
			if diff := cmp.Diff(out, roundTripped.ToProto(), protocmp.Transform()); diff != "" {
				t.Errorf("join with post-join filter did not round trip, diff:\n%v", diff)
			}
		})
	}
}
