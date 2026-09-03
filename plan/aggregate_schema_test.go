// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/proto"
)

func TestAggregateGroupingSchema(t *testing.T) {
	stringType := &types.StringType{Nullability: types.NullabilityRequired}
	boolType := &types.BooleanType{Nullability: types.NullabilityRequired}
	floatType := &types.Float64Type{Nullability: types.NullabilityNullable}
	countType := &types.Int64Type{Nullability: types.NullabilityRequired}
	groupingIndexType := &types.Int32Type{Nullability: types.NullabilityRequired}

	for _, tc := range []struct {
		name       string
		groups     [][]uint32
		keyCount   int
		noMeasures bool
		want       []types.Type
	}{
		{
			name: "permuted references preserve declaration order", groups: [][]uint32{{2, 0, 1}}, keyCount: 3,
			want: []types.Type{stringType, boolType, floatType, countType, floatType},
		},
		{
			name: "overlapping sets share columns", groups: [][]uint32{{2, 0}, {0, 1}}, keyCount: 3,
			want: []types.Type{stringType, &types.BooleanType{Nullability: types.NullabilityNullable}, floatType, countType, floatType, groupingIndexType},
		},
		{
			name: "grand total makes keys nullable", groups: [][]uint32{{0, 1, 2}, {}}, keyCount: 3,
			want: []types.Type{&types.StringType{Nullability: types.NullabilityNullable}, &types.BooleanType{Nullability: types.NullabilityNullable}, floatType, countType, floatType, groupingIndexType},
		},
		{
			name: "identical sets retain required keys", groups: [][]uint32{{0, 1, 2}, {2, 1, 0}}, keyCount: 3,
			want: []types.Type{stringType, boolType, floatType, countType, floatType, groupingIndexType},
		},
		{
			name: "single empty grouping set", groups: [][]uint32{{}},
			want: []types.Type{countType, floatType},
		},
		{
			name: "no grouping sets",
			want: []types.Type{countType, floatType},
		},
		{
			name: "grouping sets without measures", groups: [][]uint32{{0}, {}}, keyCount: 1, noMeasures: true,
			want: []types.Type{&types.StringType{Nullability: types.NullabilityNullable}, groupingIndexType},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := plan.NewBuilderDefault()
			scan := b.NamedScan([]string{"test"}, types.NamedStruct{
				Names: []string{"a", "b", "c"},
				Struct: types.StructType{Nullability: types.NullabilityRequired,
					Types: []types.Type{stringType, boolType, floatType}},
			})
			var measures []plan.AggRelMeasure
			if !tc.noMeasures {
				count, err := b.AggregateFn(extensions.SubstraitDefaultURNPrefix+"functions_aggregate_generic", "count", nil)
				require.NoError(t, err)
				value, err := b.RootFieldRef(scan, 2)
				require.NoError(t, err)
				avg, err := b.AggregateFn(extensions.SubstraitDefaultURNPrefix+"functions_arithmetic", "avg", nil, value)
				require.NoError(t, err)
				measures = []plan.AggRelMeasure{b.Measure(count, nil), b.Measure(avg, nil)}
			}
			arb := b.GetRelBuilder().AggregateRel(scan, measures)
			for i := 0; i < tc.keyCount; i++ {
				key, err := b.RootFieldRef(scan, int32(i))
				require.NoError(t, err)
				arb.AddExpression(key)
			}
			for _, group := range tc.groups {
				arb.AddGroupingSet(group)
			}
			aggregate, err := arb.Build()
			require.NoError(t, err)
			assert.Equal(t, tc.want, aggregate.RecordType().Types())

			// Valid root names must be accepted for the specified output schema,
			// including after serializing and decoding the entire plan.
			names := []string{"a", "b", "c", "d", "e", "f"}[:len(tc.want)]
			p, err := b.Plan(aggregate, names)
			require.NoError(t, err)
			serialized, err := p.ToProto()
			require.NoError(t, err)
			wire, err := proto.Marshal(serialized)
			require.NoError(t, err)
			decoded := &substraitproto.Plan{}
			require.NoError(t, proto.Unmarshal(wire, decoded))
			roundTrip, err := plan.FromProto(decoded, extensions.GetDefaultCollectionWithNoError())
			require.NoError(t, err)
			assert.Equal(t, tc.want, roundTrip.GetRoots()[0].Input().RecordType().Types())
		})
	}
}

func TestAggregateGroupingSchemaEmit(t *testing.T) {
	b := plan.NewBuilderDefault()
	scan := b.NamedScan([]string{"test"}, baseSchema)
	count, err := b.AggregateFn(extensions.SubstraitDefaultURNPrefix+"functions_aggregate_generic", "count", nil)
	require.NoError(t, err)
	key, err := b.RootFieldRef(scan, 0)
	require.NoError(t, err)
	aggregate, err := b.AggregateExprs(scan, []plan.AggRelMeasure{b.Measure(count, nil)}, []expr.Expression{key}, []expr.Expression{})
	require.NoError(t, err)
	remapped, err := aggregate.Remap(2, 1, 0)
	require.NoError(t, err)
	want := []types.Type{
		&types.Int32Type{Nullability: types.NullabilityRequired},
		&types.Int64Type{Nullability: types.NullabilityRequired},
		&types.StringType{Nullability: types.NullabilityNullable},
	}
	assert.Equal(t, want, remapped.RecordType().Types())
	p, err := b.Plan(remapped, []string{"grouping", "count", "key"})
	require.NoError(t, err)
	serialized, err := p.ToProto()
	require.NoError(t, err)
	assert.Equal(t, []int32{2, 1, 0}, serialized.Relations[0].GetRoot().Input.GetAggregate().Common.GetEmit().OutputMapping)
	roundTrip, err := plan.FromProto(serialized, extensions.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Equal(t, want, roundTrip.GetRoots()[0].Input().RecordType().Types())
}

func TestAggregateGroupingSchemaPreservesInputTypes(t *testing.T) {
	for _, inputType := range []types.Type{
		&types.StringType{Nullability: types.NullabilityRequired, TypeVariationRef: 7},
		&types.IntervalDayType{Precision: types.PrecisionMicroSeconds, Nullability: types.NullabilityRequired, TypeVariationRef: 7},
		&types.PrecisionTimestampType{Precision: types.PrecisionNanoSeconds, Nullability: types.NullabilityRequired, TypeVariationRef: 7},
	} {
		t.Run(inputType.String(), func(t *testing.T) {
			b := plan.NewBuilderDefault()
			scan := b.NamedScan([]string{"test"}, types.NamedStruct{
				Names:  []string{"key"},
				Struct: types.StructType{Nullability: types.NullabilityRequired, Types: []types.Type{inputType}},
			})
			key, err := b.RootFieldRef(scan, 0)
			require.NoError(t, err)
			aggregate, err := b.AggregateExprs(scan, nil, []expr.Expression{key}, []expr.Expression{})
			require.NoError(t, err)
			before := proto.Clone(types.TypeToProto(inputType))
			for i := 0; i < 2; i++ {
				output := aggregate.RecordType().Types()[0]
				assert.Equal(t, types.NullabilityNullable, output.GetNullability())
				assert.Equal(t, uint32(7), output.GetTypeVariationReference())
				assert.Equal(t, inputType.GetParameters(), output.GetParameters())
				assert.True(t, proto.Equal(before, types.TypeToProto(inputType)), "computing the aggregate schema must not mutate the input type")
				assert.Equal(t, types.NullabilityRequired, key.GetType().GetNullability())
			}
		})
	}
}
