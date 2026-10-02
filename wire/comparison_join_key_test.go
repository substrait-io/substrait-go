// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

// createJoinInput builds a named table read rel with three required int64 columns.
func createJoinInput(name string) *plan.NamedTableReadRel {
	schema := types.NamedStruct{
		Names: []string{"x", "y", "z"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int64Type{}, &types.Int64Type{}, &types.Int64Type{}},
		},
	}
	base := plan.NewBaseReadRel(plan.RelCommon{}, schema, nil, nil, nil, nil)
	return plan.NewNamedTableReadRel(base, []string{name}, nil)
}

func joinTestRegistry() expr.ExtensionRegistry {
	return expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
}

// All joins below have a left and right input with three int64 columns each.
func joinInputs() (plan.Rel, plan.Rel) {
	return createJoinInput("L"), createJoinInput("R")
}

func eqKeys(t *testing.T, left, right plan.Rel) []*plan.ComparisonJoinKey {
	return []*plan.ComparisonJoinKey{
		plan.NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 2)),
		plan.NewEqualityJoinKey(keyRef(t, left, 1), keyRef(t, right, 0)),
	}
}

// i64Type is a required int64 protobuf type.
func i64Type() *proto.Type {
	return &proto.Type{Kind: &proto.Type_I64_{I64: &proto.Type_I64{Nullability: proto.Type_NULLABILITY_REQUIRED}}}
}

// scanProtoRel is the protobuf for a named table read with three int64 columns,
// matching the domain input createJoinInput produces.
func scanProtoRel(name string) *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_Read{Read: &proto.ReadRel{
		Common: &proto.RelCommon{EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}}},
		BaseSchema: &proto.NamedStruct{
			Names:  []string{"x", "y", "z"},
			Struct: &proto.Type_Struct{Nullability: proto.Type_NULLABILITY_REQUIRED, Types: []*proto.Type{i64Type(), i64Type(), i64Type()}},
		},
		ReadType: &proto.ReadRel_NamedTable_{NamedTable: &proto.ReadRel_NamedTable{Names: []string{name}}},
	}}}
}

// fieldRefProto is a root-referenced struct-field protobuf reference.
func fieldRefProto(field int32) *proto.Expression_FieldReference {
	return &proto.Expression_FieldReference{
		ReferenceType: &proto.Expression_FieldReference_DirectReference{
			DirectReference: &proto.Expression_ReferenceSegment{
				ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
					StructField: &proto.Expression_ReferenceSegment_StructField{Field: field}}}},
		RootType: &proto.Expression_FieldReference_RootReference_{RootReference: &proto.Expression_FieldReference_RootReference{}},
	}
}

func eqComparisonProto() *proto.ComparisonJoinKey_ComparisonType {
	return &proto.ComparisonJoinKey_ComparisonType{
		InnerType: &proto.ComparisonJoinKey_ComparisonType_Simple{Simple: proto.ComparisonJoinKey_SIMPLE_COMPARISON_TYPE_EQ}}
}

// fieldIndex returns the struct-field index a join key references on one side.
func fieldIndex(t *testing.T, ref *expr.FieldReference) int32 {
	t.Helper()
	seg, ok := ref.Reference.(*expr.StructFieldRef)
	require.True(t, ok)
	return seg.Field
}

// Producing a join always writes the new keys field, and additionally
// mirrors the deprecated left_keys/right_keys when every key is a plain EQ
// comparison so old consumers keep working for equality joins.
func TestJoinWritesLegacyKeysForEqualityOnly(t *testing.T) {
	left, right := joinInputs()
	wantLeft := []*proto.Expression_FieldReference{
		fieldReferenceRefToProto(keyRef(t, left, 0)), fieldReferenceRefToProto(keyRef(t, left, 1))}
	wantRight := []*proto.Expression_FieldReference{
		fieldReferenceRefToProto(keyRef(t, right, 2)), fieldReferenceRefToProto(keyRef(t, right, 0))}

	// All-EQ keys: new keys field plus mirrored deprecated fields.
	for _, tc := range []struct {
		name string
		keys []*plan.ComparisonJoinKey
	}{
		{name: "value comparison", keys: eqKeys(t, left, right)},
		{name: "pointer comparison", keys: []*plan.ComparisonJoinKey{
			plan.NewComparisonJoinKey(keyRef(t, left, 0), keyRef(t, right, 2), &plan.SimpleComparison{Type: plan.SimpleComparisonTypeEq}),
			plan.NewComparisonJoinKey(keyRef(t, left, 1), keyRef(t, right, 0), &plan.SimpleComparison{Type: plan.SimpleComparisonTypeEq}),
		}},
	} {
		t.Run("all EQ mirrors deprecated fields: "+tc.name, func(t *testing.T) {
			hash := RelToProto(plan.NewHashJoinRel(left, right, tc.keys, plan.HashMergeInner, nil, plan.RelCommon{}, nil)).GetHashJoin()
			assert.Len(t, hash.GetKeys(), 2)
			assert.Empty(t, cmp.Diff(wantLeft, hash.GetLeftKeys(), protocmp.Transform()))
			assert.Empty(t, cmp.Diff(wantRight, hash.GetRightKeys(), protocmp.Transform()))

			merge := RelToProto(plan.NewMergeJoinRel(left, right, tc.keys, plan.HashMergeInner, nil, plan.RelCommon{}, nil)).GetMergeJoin()
			assert.Len(t, merge.GetKeys(), 2)
			assert.Empty(t, cmp.Diff(wantLeft, merge.GetLeftKeys(), protocmp.Transform()))
			assert.Empty(t, cmp.Diff(wantRight, merge.GetRightKeys(), protocmp.Transform()))
		})
	}

	// A non-EQ comparison anywhere makes the legacy fields lossy, so they are
	// omitted entirely and only the new keys field is written.
	for _, tc := range []struct {
		name       string
		comparison plan.JoinKeyComparison
	}{
		{"is not distinct from", plan.SimpleComparison{Type: plan.SimpleComparisonTypeIsNotDistinctFrom}},
		{"might equal", plan.SimpleComparison{Type: plan.SimpleComparisonTypeMightEqual}},
		{"custom", plan.CustomComparison{FunctionReference: 7}},
	} {
		t.Run("non-EQ omits deprecated fields: "+tc.name, func(t *testing.T) {
			keys := []*plan.ComparisonJoinKey{
				plan.NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 2)),
				plan.NewComparisonJoinKey(keyRef(t, left, 1), keyRef(t, right, 0), tc.comparison),
			}

			hash := RelToProto(plan.NewHashJoinRel(left, right, keys, plan.HashMergeInner, nil, plan.RelCommon{}, nil)).GetHashJoin()
			assert.Len(t, hash.GetKeys(), 2)
			assert.Empty(t, hash.GetLeftKeys())
			assert.Empty(t, hash.GetRightKeys())

			merge := RelToProto(plan.NewMergeJoinRel(left, right, keys, plan.HashMergeInner, nil, plan.RelCommon{}, nil)).GetMergeJoin()
			assert.Len(t, merge.GetKeys(), 2)
			assert.Empty(t, merge.GetLeftKeys())
			assert.Empty(t, merge.GetRightKeys())
		})
	}
}

// Equality, non-EQ simple comparisons and custom comparison functions all
// survive a round trip through proto.
func TestJoinKeysRoundTrip(t *testing.T) {
	left, right := joinInputs()
	reg := joinTestRegistry()

	keys := []*plan.ComparisonJoinKey{
		plan.NewEqualityJoinKey(keyRef(t, left, 0), keyRef(t, right, 2)),
		plan.NewComparisonJoinKey(keyRef(t, left, 1), keyRef(t, right, 0),
			plan.SimpleComparison{Type: plan.SimpleComparisonTypeIsNotDistinctFrom}),
		plan.NewComparisonJoinKey(keyRef(t, left, 2), keyRef(t, right, 1),
			plan.CustomComparison{FunctionReference: 42}),
	}

	for _, tc := range []struct {
		name string
		rel  plan.Rel
	}{
		{"hash", plan.NewHashJoinRel(left, right, keys, plan.HashMergeInner, nil, plan.RelCommon{}, nil)},
		{"merge", plan.NewMergeJoinRel(left, right, keys, plan.HashMergeInner, nil, plan.RelCommon{}, nil)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := RelToProto(tc.rel)
			roundTripped, err := RelFromProto(out, reg)
			require.NoError(t, err)

			if diff := cmp.Diff(out, RelToProto(roundTripped), protocmp.Transform()); diff != "" {
				t.Errorf("join did not round trip, diff:\n%v", diff)
			}

			// The comparison kinds are preserved on the model side.
			var gotKeys []*plan.ComparisonJoinKey
			switch r := roundTripped.(type) {
			case *plan.HashJoinRel:
				gotKeys = r.Keys()
			case *plan.MergeJoinRel:
				gotKeys = r.Keys()
			}
			require.Len(t, gotKeys, 3)
			assert.Equal(t, plan.SimpleComparison{Type: plan.SimpleComparisonTypeEq}, gotKeys[0].Comparison())
			assert.Equal(t, plan.SimpleComparison{Type: plan.SimpleComparisonTypeIsNotDistinctFrom}, gotKeys[1].Comparison())
			assert.Equal(t, plan.CustomComparison{FunctionReference: 42}, gotKeys[2].Comparison())
		})
	}
}

// assertConsumedLegacyKeys checks that the keys decoded from a legacy producer
// are the expected pair of EQ comparisons over fields (0,2) and (1,0).
func assertConsumedLegacyKeys(t *testing.T, keys []*plan.ComparisonJoinKey) {
	t.Helper()
	require.Len(t, keys, 2)
	for _, k := range keys {
		assert.Equal(t, plan.SimpleComparison{Type: plan.SimpleComparisonTypeEq}, k.Comparison())
	}
	assert.Equal(t, int32(0), fieldIndex(t, keys[0].Left()))
	assert.Equal(t, int32(2), fieldIndex(t, keys[0].Right()))
	assert.Equal(t, int32(1), fieldIndex(t, keys[1].Left()))
	assert.Equal(t, int32(0), fieldIndex(t, keys[1].Right()))
}

// A hash join plan from a legacy producer (only deprecated fields set) is
// consumed and mapped to keys with EQ comparisons.
func TestHashJoinConsumesLegacyKeys(t *testing.T) {
	reg := joinTestRegistry()
	legacy := &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
		Common:    &proto.RelCommon{},
		Left:      scanProtoRel("L"),
		Right:     scanProtoRel("R"),
		Type:      proto.HashJoinRel_JOIN_TYPE_INNER,
		LeftKeys:  []*proto.Expression_FieldReference{fieldRefProto(0), fieldRefProto(1)},
		RightKeys: []*proto.Expression_FieldReference{fieldRefProto(2), fieldRefProto(0)},
	}}}

	rel, err := RelFromProto(legacy, reg)
	require.NoError(t, err)
	assertConsumedLegacyKeys(t, rel.(*plan.HashJoinRel).Keys())
}

// As TestHashJoinConsumesLegacyKeys, but for merge joins.
func TestMergeJoinConsumesLegacyKeys(t *testing.T) {
	reg := joinTestRegistry()
	legacy := &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
		Common:    &proto.RelCommon{},
		Left:      scanProtoRel("L"),
		Right:     scanProtoRel("R"),
		Type:      proto.MergeJoinRel_JOIN_TYPE_INNER,
		LeftKeys:  []*proto.Expression_FieldReference{fieldRefProto(0), fieldRefProto(1)},
		RightKeys: []*proto.Expression_FieldReference{fieldRefProto(2), fieldRefProto(0)},
	}}}

	rel, err := RelFromProto(legacy, reg)
	require.NoError(t, err)
	assertConsumedLegacyKeys(t, rel.(*plan.MergeJoinRel).Keys())
}

// When both the deprecated fields and the new keys are present, the new keys win.
func TestJoinPrefersNewKeysOverDeprecated(t *testing.T) {
	reg := joinTestRegistry()
	newKeys := []*proto.ComparisonJoinKey{
		{Left: fieldRefProto(0), Right: fieldRefProto(2), Comparison: eqComparisonProto()},
		{Left: fieldRefProto(1), Right: fieldRefProto(0), Comparison: eqComparisonProto()},
	}
	// Bogus deprecated keys pointing at different fields than the real keys.
	bogusLeft := []*proto.Expression_FieldReference{fieldRefProto(2)}
	bogusRight := []*proto.Expression_FieldReference{fieldRefProto(1)}

	assertNewKeysWon := func(t *testing.T, keys []*plan.ComparisonJoinKey) {
		t.Helper()
		require.Len(t, keys, 2)
		assert.Equal(t, int32(0), fieldIndex(t, keys[0].Left()))
		assert.Equal(t, int32(1), fieldIndex(t, keys[1].Left()))
	}

	t.Run("hash", func(t *testing.T) {
		both := &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
			Common:    &proto.RelCommon{},
			Left:      scanProtoRel("L"),
			Right:     scanProtoRel("R"),
			Type:      proto.HashJoinRel_JOIN_TYPE_INNER,
			Keys:      newKeys,
			LeftKeys:  bogusLeft,
			RightKeys: bogusRight,
		}}}
		rel, err := RelFromProto(both, reg)
		require.NoError(t, err)
		assertNewKeysWon(t, rel.(*plan.HashJoinRel).Keys())
	})

	t.Run("merge", func(t *testing.T) {
		both := &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
			Common:    &proto.RelCommon{},
			Left:      scanProtoRel("L"),
			Right:     scanProtoRel("R"),
			Type:      proto.MergeJoinRel_JOIN_TYPE_INNER,
			Keys:      newKeys,
			LeftKeys:  bogusLeft,
			RightKeys: bogusRight,
		}}}
		rel, err := RelFromProto(both, reg)
		require.NoError(t, err)
		assertNewKeysWon(t, rel.(*plan.MergeJoinRel).Keys())
	})
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

// joinFields holds the key-bearing fields shared by plan.HashJoinRel and
// plan.MergeJoinRel so an error case can be exercised against both join types.
type joinFields struct {
	keys      []*proto.ComparisonJoinKey
	leftKeys  []*proto.Expression_FieldReference
	rightKeys []*proto.Expression_FieldReference
}

func (f joinFields) hashProto() *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
		Common: &proto.RelCommon{}, Type: proto.HashJoinRel_JOIN_TYPE_INNER,
		Left: scanProtoRel("L"), Right: scanProtoRel("R"),
		Keys: f.keys, LeftKeys: f.leftKeys, RightKeys: f.rightKeys,
	}}}
}

func (f joinFields) mergeProto() *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
		Common: &proto.RelCommon{}, Type: proto.MergeJoinRel_JOIN_TYPE_INNER,
		Left: scanProtoRel("L"), Right: scanProtoRel("R"),
		Keys: f.keys, LeftKeys: f.leftKeys, RightKeys: f.rightKeys,
	}}}
}

// Error paths in comparisonJoinKeysFromProto / joinKeyComparisonFromProto,
// exercised against both hash and merge joins.
func TestJoinKeysFromProtoErrors(t *testing.T) {
	reg := joinTestRegistry()
	goodLeft := fieldRefProto(0)
	goodRight := fieldRefProto(0)
	eq := eqComparisonProto()

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
			_, err := RelFromProto(tc.fields.hashProto(), reg)
			require.Error(t, err)
		})
		t.Run("merge/"+tc.name, func(t *testing.T) {
			_, err := RelFromProto(tc.fields.mergeProto(), reg)
			require.Error(t, err)
		})
	}
}
