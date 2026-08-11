// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/expr"
	ext "github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// The version field is optional on the wire. An absent one has to stay absent through the domain,
// otherwise a round trip turns it into an empty message. TestRoundTripExtendedExpression covers the
// populated case.
func TestExtendedWithoutAVersion(t *testing.T) {
	in := &proto.ExtendedExpression{}
	result, err := expr.ExtendedFromProto(in, ext.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Nil(t, result.Version)

	assert.Nil(t, result.ToProto().Version)
}

func TestExtendedVersionRoundTrip(t *testing.T) {
	result, err := expr.ExtendedFromProto(&proto.ExtendedExpression{
		Version: &proto.Version{MinorNumber: 29, Producer: "substraitgo-test"},
	}, ext.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Equal(t, &types.Version{MinorNumber: 29, Producer: "substraitgo-test"}, result.Version)

	out := result.ToProto().Version
	require.NotNil(t, out)
	assert.Equal(t, uint32(29), out.MinorNumber)
	assert.Equal(t, "substraitgo-test", out.Producer)
}
