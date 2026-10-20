// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec/v9"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// An extended expression parsed with no version is accepted; the missing version surfaces as
// "0.0.0 (UNSET)", including on the way back out.
func TestExtendedWithoutAVersion(t *testing.T) {
	result, err := codec.ExtendedFromProto(&proto.ExtendedExpression{}, ext.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Equal(t, "0.0.0 (UNSET)", result.Version.String())
	assert.Equal(t, "UNSET", codec.ExtendedToProto(result).Version.GetProducer())
}
