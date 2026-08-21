// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/expr"
	ext "github.com/substrait-io/substrait-go/v8/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// An absent version must stay absent through a round trip, not become an empty message. The
// populated case is already covered by TestRoundTripExtendedExpression.
func TestExtendedWithoutAVersion(t *testing.T) {
	in := &proto.ExtendedExpression{}
	result, err := expr.ExtendedFromProto(in, ext.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Nil(t, result.Version)

	assert.Nil(t, result.ToProto().Version)
}
