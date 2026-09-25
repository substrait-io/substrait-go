// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
)

// nilFilterSchema is a one-column required-int64 schema for building the inputs
// used by the nil-filter omission tests.
func nilFilterSchema() types.NamedStruct {
	return types.NamedStruct{
		Names: []string{"a"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int64Type{}},
		},
	}
}

// A nil aggregate-measure filter must be omitted from the encoded proto, not
// replaced by the default true literal the read-time getter substitutes. The
// encode->decode->encode round-trip cannot catch a regression that swaps the
// raw accessor for the substituting getter (the default would apply on both
// sides and still compare equal), so read the encoded bytes directly.
func TestAggRelMeasureNilFilterOmitted(t *testing.T) {
	b := plan.NewBuilderDefault()
	aggCount, err := b.AggregateFn(extensions.SubstraitDefaultURNPrefix+"functions_aggregate_generic", "count", nil)
	require.NoError(t, err)
	scan := b.NamedScan([]string{"test"}, nilFilterSchema())
	agg, err := b.AggregateColumns(scan, []plan.AggRelMeasure{b.Measure(aggCount, nil)}, 0)
	require.NoError(t, err)

	measures := wire.RelToProto(agg).GetAggregate().GetMeasures()
	require.Len(t, measures, 1)
	assert.Nil(t, measures[0].GetFilter(), "a nil measure filter must be absent from the encoded proto")
}
