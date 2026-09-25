// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/wire"
	substraitproto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// A write op with no write type is malformed. Decoding must reject it rather
// than silently build a NamedTableWriteRel with empty names.
func TestWriteRelMissingWriteTypeRejected(t *testing.T) {
	b := plan.NewBuilderDefault()
	schema := nilFilterSchema()
	scan := b.NamedScan([]string{"test"}, schema)

	rel := &substraitproto.Rel{RelType: &substraitproto.Rel_Write{Write: &substraitproto.WriteRel{
		Input:       wire.RelToProto(scan),
		TableSchema: wire.NamedStructToProto(schema),
		Op:          substraitproto.WriteRel_WRITE_OP_CTAS,
		// WriteType deliberately left nil.
	}}}

	reg := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	_, err := wire.RelFromProto(rel, reg)
	require.Error(t, err)
	assert.ErrorIs(t, err, substraitgo.ErrInvalidRel)
}
