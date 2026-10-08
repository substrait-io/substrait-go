// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	protobuf "google.golang.org/protobuf/proto"
)

// The fixtures are unchanged DuckDB 1.5.5 SQL-produced plans over
// t(c0 BIGINT NOT NULL, c1 VARCHAR, c2 BOOLEAN), in schema and reversed order.
func TestDuckDBReadProjection(t *testing.T) {
	i64 := &types.Int64Type{Nullability: types.NullabilityRequired}
	boolean := &types.BooleanType{Nullability: types.NullabilityNullable}
	for _, tc := range []struct {
		file  string
		names []string
		types []types.Type
	}{
		{"read_projection_duckdb.json", []string{"c0", "c2"}, []types.Type{i64, boolean}},
		{"read_projection_reordered_duckdb.json", []string{"c2", "c0"}, []types.Type{boolean, i64}},
	} {
		t.Run(tc.file, func(t *testing.T) {
			data, err := os.ReadFile("testdata/" + tc.file)
			require.NoError(t, err)
			var input proto.Plan
			require.NoError(t, protojson.Unmarshal(data, &input))
			p, err := PlanFromProto(&input, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)
			root := p.GetRoots()[0]
			assert.Equal(t, tc.names, root.Names())
			assert.Equal(t, tc.types, root.RecordType().Struct.Types)
			output, err := PlanToProto(p)
			require.NoError(t, err)
			// The encoder makes the default direct emit explicit.
			want := protobuf.Clone(&input).(*proto.Plan)
			want.Relations[0].GetRoot().Input.GetRead().Common = &proto.RelCommon{
				EmitKind: &proto.RelCommon_Direct_{Direct: &proto.RelCommon_Direct{}},
			}
			assert.True(t, protobuf.Equal(want, output))
			roundTrip, err := PlanFromProto(output, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)
			assert.Equal(t, root.RecordType(), roundTrip.GetRoots()[0].RecordType())
		})
	}
}

func TestReadProjectionFiltersUseBaseSchema(t *testing.T) {
	data, err := os.ReadFile("testdata/read_projection_duckdb.json")
	require.NoError(t, err)
	var input proto.Plan
	require.NoError(t, protojson.Unmarshal(data, &input))
	read := input.Relations[0].GetRoot().Input.GetRead()
	base := plan.NewBuilderDefault().NamedScan([]string{"t"}, NamedStructFromProto(read.BaseSchema))
	predicate := ExprToProto(keyRef(t, base, 2))
	read.Filter, read.BestEffortFilter = predicate, predicate
	read.Projection.Select.StructItems = read.Projection.Select.StructItems[:1]
	input.Relations[0].GetRoot().Names = []string{"c0"}
	p, err := PlanFromProto(&input, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	decoded := p.GetRoots()[0].Input().(plan.ReadRel)
	assert.Equal(t, &types.BooleanType{Nullability: types.NullabilityNullable}, decoded.Filter().GetType())
	assert.Equal(t, decoded.Filter().GetType(), decoded.BestEffortFilter().GetType())
	assert.Equal(t, int32(1), decoded.RecordType().FieldCount())
	assert.Len(t, decoded.BaseSchema().Struct.Types, 3)
}

func TestReadProjectionDecodeErrors(t *testing.T) {
	data, err := os.ReadFile("testdata/read_projection_duckdb.json")
	require.NoError(t, err)
	var input proto.Plan
	require.NoError(t, protojson.Unmarshal(data, &input))
	for _, tc := range []struct {
		name string
		edit func(*proto.ReadRel)
		want error
	}{
		{"negative field", func(r *proto.ReadRel) { r.Projection.Select.StructItems[0].Field = -1 }, substraitgo.ErrInvalidExpr},
		{"field out of range", func(r *proto.ReadRel) { r.Projection.Select.StructItems[1].Field = 3 }, substraitgo.ErrInvalidExpr},
		{"nested mask", func(r *proto.ReadRel) {
			r.BaseSchema.Names = []string{"s", "a", "b"}
			r.BaseSchema.Struct.Types = []*proto.Type{{Kind: &proto.Type_Struct_{Struct: &proto.Type_Struct{
				Nullability: proto.Type_NULLABILITY_REQUIRED,
				Types:       r.BaseSchema.Struct.Types[:2],
			}}}}
			r.Projection.Select.StructItems = r.Projection.Select.StructItems[:1]
			r.Projection.Select.StructItems[0].Child = &proto.Expression_MaskExpression_Select{
				Type: &proto.Expression_MaskExpression_Select_Struct{Struct: &proto.Expression_MaskExpression_StructSelect{
					StructItems: []*proto.Expression_MaskExpression_StructItem{{Field: 0}, {Field: 1}},
				}},
			}
		}, substraitgo.ErrNotImplemented},
		{"emit out of range", func(r *proto.ReadRel) {
			r.Common = &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{2}}}}
		}, substraitgo.ErrInvalidRel},
		{"negative emit", func(r *proto.ReadRel) {
			r.Common = &proto.RelCommon{EmitKind: &proto.RelCommon_Emit_{Emit: &proto.RelCommon_Emit{OutputMapping: []int32{-1}}}}
		}, substraitgo.ErrInvalidRel},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := protobuf.Clone(&input).(*proto.Plan)
			tc.edit(p.Relations[0].GetRoot().Input.GetRead())
			assert.NotPanics(t, func() {
				_, err := PlanFromProto(p, extensions.GetDefaultCollectionWithNoError())
				assert.ErrorIs(t, err, tc.want)
			})
		})
	}
	input.Relations[0].GetRoot().Names = []string{"c0", "c1", "c2"}
	_, err = PlanFromProto(&input, extensions.GetDefaultCollectionWithNoError())
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
}
