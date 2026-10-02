// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// An extended expression parsed with no version is accepted; the missing version surfaces as
// "0.0.0 (UNSET)", including on the way back out.
func TestExtendedWithoutAVersion(t *testing.T) {
	result, err := wire.ExtendedFromProto(&proto.ExtendedExpression{}, ext.GetDefaultCollectionWithNoError())
	require.NoError(t, err)
	assert.Equal(t, "0.0.0 (UNSET)", result.Version.String())
	assert.Equal(t, "UNSET", wire.ExtendedToProto(result).Version.GetProducer())
}
