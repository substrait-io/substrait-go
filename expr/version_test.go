// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/expr"
	ext "github.com/substrait-io/substrait-go/v8/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// version is required, so an extended expression without one is rejected at parse.
func TestExtendedWithoutAVersion(t *testing.T) {
	_, err := expr.ExtendedFromProto(&proto.ExtendedExpression{}, ext.GetDefaultCollectionWithNoError())
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
}
