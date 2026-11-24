// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
)

func TestSetPreservesImportedAnchors(t *testing.T) {
	const urn = "extension:test:sample"
	t.Run("function", func(t *testing.T) {
		testImportedAnchorAllocation(t,
			extensions.FunctionID{URN: urn, Name: "add:i32_i32"},
			extensions.FunctionID{URN: urn, Name: "subtract:i32_i32"},
			func(imported map[uint32]extensions.FunctionID) (extensions.Set, error) {
				return extensions.NewSetFromParts(map[uint32]string{1: urn}, nil, nil, imported)
			}, extensions.Set.GetFuncAnchor, extensions.Set.DecodeFunc, true)
	})
	t.Run("type", func(t *testing.T) {
		testImportedAnchorAllocation(t,
			extensions.TypeID{URN: urn, Name: "point"},
			extensions.TypeID{URN: urn, Name: "line"},
			func(imported map[uint32]extensions.TypeID) (extensions.Set, error) {
				return extensions.NewSetFromParts(map[uint32]string{1: urn}, imported, nil, nil)
			}, extensions.Set.GetTypeAnchor, extensions.Set.DecodeType, true)
	})
	t.Run("type variation", func(t *testing.T) {
		testImportedAnchorAllocation(t,
			extensions.TypeVariationID{URN: urn, Name: "dict4"},
			extensions.TypeVariationID{URN: urn, Name: "bigoffset"},
			func(imported map[uint32]extensions.TypeVariationID) (extensions.Set, error) {
				return extensions.NewSetFromParts(map[uint32]string{1: urn}, nil, imported, nil)
			}, extensions.Set.GetTypeVariationAnchor, extensions.Set.DecodeTypeVariation, false)
	})
}

func testImportedAnchorAllocation[T comparable](t *testing.T, importedID, newID T,
	newSet func(map[uint32]T) (extensions.Set, error),
	getAnchor func(extensions.Set, T) uint32, decode func(extensions.Set, uint32) (T, bool), zeroAllowed bool) {
	t.Helper()
	anchors := []uint32{2, math.MaxUint32}
	if zeroAllowed {
		anchors = append(anchors, 0)
	}
	for _, importedAnchor := range anchors {
		t.Run(fmt.Sprint(importedAnchor), func(t *testing.T) {
			set, err := newSet(map[uint32]T{importedAnchor: importedID})
			require.NoError(t, err)
			newAnchor := getAnchor(set, newID)
			assert.NotZero(t, newAnchor)
			assert.NotEqual(t, importedAnchor, newAnchor)
			assert.Equal(t, newAnchor, getAnchor(set, newID))
			assert.Equal(t, importedAnchor, getAnchor(set, importedID))
			id, ok := decode(set, importedAnchor)
			assert.True(t, ok)
			assert.Equal(t, importedID, id)
			id, ok = decode(set, newAnchor)
			assert.True(t, ok)
			assert.Equal(t, newID, id)
		})
	}
}

func TestSetPreservesImportedURNAnchors(t *testing.T) {
	importedID := extensions.FunctionID{URN: "extension:test:arithmetic", Name: "add:i32_i32"}
	newID := extensions.FunctionID{URN: "extension:test:boolean", Name: "and:bool_bool"}
	for _, importedAnchor := range []uint32{0, 2, math.MaxUint32} {
		t.Run(fmt.Sprint(importedAnchor), func(t *testing.T) {
			set, err := extensions.NewSetFromParts(
				map[uint32]string{importedAnchor: importedID.URN}, nil, nil,
				map[uint32]extensions.FunctionID{1: importedID})
			require.NoError(t, err)
			var functionAnchor uint32
			require.NotPanics(t, func() { functionAnchor = set.GetFuncAnchor(newID) })
			newAnchor, ok := set.FindURN(newID.URN)
			require.True(t, ok)
			assert.NotZero(t, newAnchor)
			assert.NotEqual(t, importedAnchor, newAnchor)
			anchor, ok := set.FindURN(importedID.URN)
			require.True(t, ok)
			assert.Equal(t, importedAnchor, anchor)
			assert.Len(t, set.URNs(), 2)
			id, ok := set.DecodeFunc(1)
			require.True(t, ok)
			assert.Equal(t, importedID, id)
			id, ok = set.DecodeFunc(functionAnchor)
			require.True(t, ok)
			assert.Equal(t, newID, id)
		})
	}
}

func TestSetAllocatesAfterConsecutiveImportedAnchors(t *testing.T) {
	const urn = "extension:test:arithmetic"
	ids := []extensions.FunctionID{
		{URN: urn, Name: "add:i32_i32"},
		{URN: urn, Name: "subtract:i32_i32"},
		{URN: urn, Name: "multiply:i32_i32"},
		{URN: urn, Name: "divide:i32_i32"},
		{URN: urn, Name: "modulus:i32_i32"},
	}
	for _, firstAnchor := range []uint32{1, 3} {
		t.Run(fmt.Sprint(firstAnchor), func(t *testing.T) {
			set, err := extensions.NewSetFromParts(map[uint32]string{1: urn}, nil, nil,
				map[uint32]extensions.FunctionID{firstAnchor: ids[0], firstAnchor + 1: ids[1]})
			require.NoError(t, err)
			for i := 2; i < 4; i++ {
				assert.Equal(t, firstAnchor+uint32(i), set.GetFuncAnchor(ids[i]))
			}
			imported := make(map[uint32]extensions.FunctionID)
			for id, anchor := range set.Functions() {
				imported[anchor] = id
			}
			set, err = extensions.NewSetFromParts(set.URNs(), nil, nil, imported)
			require.NoError(t, err)
			assert.Equal(t, firstAnchor+4, set.GetFuncAnchor(ids[4]))
			for i, id := range ids {
				anchor := firstAnchor + uint32(i)
				assert.Equal(t, anchor, set.GetFuncAnchor(id))
				decoded, ok := set.DecodeFunc(anchor)
				assert.True(t, ok)
				assert.Equal(t, id, decoded)
			}
		})
	}
}
