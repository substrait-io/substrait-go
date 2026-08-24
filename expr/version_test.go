// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// An extended expression parsed with no version is accepted but flagged: HasVersion is false and
// Version is the UnsetVersion sentinel that renders as "0.0.0 (UNSET)".
func TestExtendedWithoutAVersion(t *testing.T) {
	result, err := expr.ExtendedFromProto(&proto.ExtendedExpression{}, ext.GetDefaultCollectionWithNoError())
	require.NoError(t, err)

	assert.False(t, result.HasVersion())
	assert.Equal(t, types.UnsetVersion, result.Version)
	assert.Equal(t, "0.0.0 (UNSET)", result.Version.String())

	// the sentinel is written back out, so the missing version stays visible on the wire
	assert.Equal(t, "UNSET", result.ToProto().Version.GetProducer())
}
