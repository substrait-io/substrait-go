// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
)

func TestNewLiteralWithIntervalDayToSecond(t *testing.T) {
	_, err := expr.NewLiteral((*types.IntervalDayToSecond)(nil), false)
	require.Error(t, err)

	v := &types.IntervalDayToSecond{Days: 1, Seconds: 2, Subseconds: 3, Precision: types.PrecisionMicroSeconds}
	lit, err := expr.NewLiteral(v, false)
	require.NoError(t, err)
	assert.Equal(t, "1 days, 2 seconds, 3 subseconds", lit.ValueString())
	assert.Equal(t, "P1DT2.000003S", lit.(types.IsoValuePrinter).IsoValueString())

	pb := wire.LiteralToProto(lit).GetIntervalDayToSecond()
	assert.Equal(t, int32(1), pb.GetDays())
	assert.Equal(t, int32(2), pb.GetSeconds())
	assert.Equal(t, int64(3), pb.GetSubseconds())
	assert.Equal(t, int32(6), pb.GetPrecision())
}
