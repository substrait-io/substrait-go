// SPDX-License-Identifier: Apache-2.0

package substraitproto

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
	keys []*proto.ComparisonJoinKey
}

func (f joinFields) hashProto() *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
		Common: &proto.RelCommon{}, Type: proto.HashJoinRel_JOIN_TYPE_INNER,
		Left: scanProtoRel("L"), Right: scanProtoRel("R"),
		Keys: f.keys,
	}}}
}

func (f joinFields) mergeProto() *proto.Rel {
	return &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
		Common: &proto.RelCommon{}, Type: proto.MergeJoinRel_JOIN_TYPE_INNER,
		Left: scanProtoRel("L"), Right: scanProtoRel("R"),
		Keys: f.keys,
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
