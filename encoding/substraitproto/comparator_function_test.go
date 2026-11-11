// SPDX-License-Identifier: Apache-2.0

package substraitproto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	gproto "google.golang.org/protobuf/proto"
)

func TestComparatorsRequireLoadedScalarDefinitions(t *testing.T) {
	const urn = "extension:test:comparators"
	var collection extensions.Collection
	require.NoError(t, collection.Load(strings.NewReader(`urn: extension:test:comparators
scalar_functions:
  - name: sort_compare
    impls:
      - args: [{value: i64}, {value: i64}]
        return: i32
  - name: join_compare
    impls:
      - args: [{value: i64}, {value: i64}]
        return: boolean
aggregate_functions:
  - name: aggregate_only
    impls:
      - args: [{value: i64}]
        return: i64
`)))
	for _, kind := range []string{"sort", "hash", "merge"} {
		t.Run(kind, func(t *testing.T) {
			loadedName := "join_compare:i64_i64"
			if kind == "sort" {
				loadedName = "sort_compare:i64_i64"
			}
			for _, tc := range []struct {
				name, urn, function string
				loaded              bool
			}{
				{"missing definition", urn, "absent:i64_i64", false},
				{"wrong kind", urn, "aggregate_only:i64", false},
				{"missing urn", "extension:test:unloaded", "absent:i64_i64", false},
				{"loaded", urn, loadedName, true},
			} {
				t.Run(tc.name, func(t *testing.T) {
					reg := expr.NewEmptyExtensionRegistry(&collection)
					anchor := reg.GetFuncAnchor(extensions.FunctionID{URN: tc.urn, Name: tc.function})
					var wire *proto.Rel
					if kind == "sort" {
						wire = &proto.Rel{RelType: &proto.Rel_Sort{Sort: &proto.SortRel{
							Input: scanProtoRel("left"), Sorts: []*proto.SortField{{
								Expr:     postJoinFilterField(0),
								SortKind: &proto.SortField_ComparisonFunctionReference{ComparisonFunctionReference: anchor},
							}},
						}}}
					} else {
						keys := []*proto.ComparisonJoinKey{{Left: fieldRefProto(0), Right: fieldRefProto(0),
							Comparison: &proto.ComparisonJoinKey_ComparisonType{
								InnerType: &proto.ComparisonJoinKey_ComparisonType_CustomFunctionReference{CustomFunctionReference: anchor},
							},
						}}
						if kind == "hash" {
							wire = &proto.Rel{RelType: &proto.Rel_HashJoin{HashJoin: &proto.HashJoinRel{
								Left: scanProtoRel("left"), Right: scanProtoRel("right"), Keys: keys, Type: proto.HashJoinRel_JOIN_TYPE_INNER,
							}}}
						} else {
							wire = &proto.Rel{RelType: &proto.Rel_MergeJoin{MergeJoin: &proto.MergeJoinRel{
								Left: scanProtoRel("left"), Right: scanProtoRel("right"), Keys: keys, Type: proto.MergeJoinRel_JOIN_TYPE_INNER,
							}}}
						}
					}
					out, err := RelFromProto(wire, reg)
					if !tc.loaded {
						require.ErrorIs(t, err, substraitgo.ErrNotFound)
						require.ErrorContains(t, err, tc.function)
						require.ErrorContains(t, err, tc.urn)
						return
					}
					require.NoError(t, err)
					// Normalize common.direct before comparing the encoded plan.
					encoded := RelToProto(out)
					roundTrip, err := RelFromProto(encoded, reg)
					require.NoError(t, err)
					require.True(t, gproto.Equal(encoded, RelToProto(roundTrip)))
				})
			}
		})
	}
}

func TestComparatorRejectsUndeclaredAnchor(t *testing.T) {
	reg := joinTestRegistry()
	_, err := SortFieldFromProto(&proto.SortField{
		Expr:     ExprToProto(expr.NewPrimitiveLiteral(int64(1), false)),
		SortKind: &proto.SortField_ComparisonFunctionReference{ComparisonFunctionReference: 42},
	}, types.NewRecordTypeFromTypes(nil), reg)
	require.ErrorIs(t, err, substraitgo.ErrNotFound)
	_, err = joinKeyComparisonFromProto(&proto.ComparisonJoinKey_ComparisonType{
		InnerType: &proto.ComparisonJoinKey_ComparisonType_CustomFunctionReference{CustomFunctionReference: 42},
	}, reg)
	require.ErrorIs(t, err, substraitgo.ErrNotFound)
}
