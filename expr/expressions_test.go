// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestExtensionRegistrySetAndGetDecoder(t *testing.T) {
	const typeURL = "type.googleapis.com/test.Ext"
	c := ext.GetDefaultCollectionWithNoError()

	t.Run("nil before set", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(c)
		require.Nil(t, reg.ExtensionRelDecoderFor(typeURL))
	})

	t.Run("returns decoder after set", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(c)
		dec := &stubRelDecoder{}
		require.NoError(t, reg.SetExtensionRelDecoder(typeURL, dec))
		require.Equal(t, dec, reg.ExtensionRelDecoderFor(typeURL))
	})

	t.Run("duplicate typeURL returns error", func(t *testing.T) {
		reg := expr.NewEmptyExtensionRegistry(c)
		require.NoError(t, reg.SetExtensionRelDecoder(typeURL, &stubRelDecoder{}))
		err := reg.SetExtensionRelDecoder(typeURL, &stubRelDecoder{})
		require.ErrorContains(t, err, typeURL)
	})
}

// stubRelDecoder is a minimal ExtensionRelDecoder that always returns an error.
type stubRelDecoder struct{}

func (s *stubRelDecoder) DecodeExtensionRel(_ *anypb.Any) (any, error) {
	return nil, errors.New("stub")
}

func TestStructFieldRefGetTypeBounds(t *testing.T) {
	st := &types.StructType{
		Types: []types.Type{
			&types.Int64Type{Nullability: types.NullabilityRequired},
			&types.Int32Type{Nullability: types.NullabilityRequired},
		},
	}

	t.Run("valid index returns type", func(t *testing.T) {
		ref := &expr.StructFieldRef{Field: 0}
		got, err := ref.GetType(st)
		require.NoError(t, err)
		assert.IsType(t, &types.Int64Type{}, got)
	})

	t.Run("negative index returns error", func(t *testing.T) {
		ref := &expr.StructFieldRef{Field: -1}
		_, err := ref.GetType(st)
		require.ErrorIs(t, err, substraitgo.ErrInvalidType)
	})

	t.Run("index equal to len returns error", func(t *testing.T) {
		ref := &expr.StructFieldRef{Field: int32(len(st.Types))}
		_, err := ref.GetType(st)
		require.ErrorIs(t, err, substraitgo.ErrInvalidType)
	})

	t.Run("index beyond len returns error", func(t *testing.T) {
		ref := &expr.StructFieldRef{Field: int32(len(st.Types) + 1)}
		_, err := ref.GetType(st)
		require.ErrorIs(t, err, substraitgo.ErrInvalidType)
	})
}
