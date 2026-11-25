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

func TestExtensionSetPreservesImportedAnchors(t *testing.T) {
	const arithmeticURN = extensions.SubstraitDefaultURNPrefix + "functions_arithmetic"
	const variationURN = extensions.SubstraitDefaultURNPrefix + "type_variations"

	t.Run("function", func(t *testing.T) {
		testImportedAnchors(t,
			extensions.FunctionID{URN: arithmeticURN, Name: "add:i32_i32"},
			extensions.FunctionID{URN: arithmeticURN, Name: "subtract:i32_i32"},
			extensions.Set.GetFuncAnchor, extensions.Set.DecodeFunc,
			func(anchor uint32, id extensions.FunctionID) (extensions.Set, error) {
				return extensions.NewSetFromParts(map[uint32]string{1: arithmeticURN}, nil, nil,
					map[uint32]extensions.FunctionID{anchor: id})
			}, true)
	})

	t.Run("type", func(t *testing.T) {
		const urn = "extension:test:sample"
		testImportedAnchors(t,
			extensions.TypeID{URN: urn, Name: "point"},
			extensions.TypeID{URN: urn, Name: "line"},
			extensions.Set.GetTypeAnchor, extensions.Set.DecodeType,
			func(anchor uint32, id extensions.TypeID) (extensions.Set, error) {
				return extensions.NewSetFromParts(map[uint32]string{1: urn},
					map[uint32]extensions.TypeID{anchor: id}, nil, nil)
			}, true)
	})

	t.Run("type variation", func(t *testing.T) {
		testImportedAnchors(t,
			extensions.TypeVariationID{URN: variationURN, Name: "dict4"},
			extensions.TypeVariationID{URN: variationURN, Name: "bigoffset"},
			extensions.Set.GetTypeVariationAnchor, extensions.Set.DecodeTypeVariation,
			func(anchor uint32, id extensions.TypeVariationID) (extensions.Set, error) {
				return extensions.NewSetFromParts(map[uint32]string{1: variationURN}, nil,
					map[uint32]extensions.TypeVariationID{anchor: id}, nil)
			}, false)
	})
}

func testImportedAnchors[T comparable](t *testing.T, importedID, newID T,
	getAnchor func(extensions.Set, T) uint32, decode func(extensions.Set, uint32) (T, bool),
	newSet func(uint32, T) (extensions.Set, error), zeroAllowed bool) {
	t.Helper()
	anchors := []uint32{2, math.MaxUint32}
	if zeroAllowed {
		anchors = append(anchors, 0)
	}
	for _, importedAnchor := range anchors {
		t.Run(fmt.Sprint(importedAnchor), func(t *testing.T) {
			s, err := newSet(importedAnchor, importedID)
			require.NoError(t, err)
			newAnchor := getAnchor(s, newID)
			assert.NotZero(t, newAnchor)
			assert.NotEqual(t, importedAnchor, newAnchor)
			assert.Equal(t, newAnchor, getAnchor(s, newID), "repeated lookup must reuse its anchor")
			assert.Equal(t, importedAnchor, getAnchor(s, importedID), "imported anchors must remain stable")

			check := func(s extensions.Set) {
				id, ok := decode(s, importedAnchor)
				assert.True(t, ok)
				assert.Equal(t, importedID, id)
				id, ok = decode(s, newAnchor)
				assert.True(t, ok)
				assert.Equal(t, newID, id)
			}
			check(s)
			require.Equal(t, 2, len(s.Types())+len(s.TypeVariations())+len(s.Functions()))
			s = reconstructExtensionSet(t, s)
			check(s)
		})
	}
}

func reconstructExtensionSet(t *testing.T, original extensions.Set) extensions.Set {
	t.Helper()
	types := make(map[uint32]extensions.TypeID)
	for id, anchor := range original.Types() {
		types[anchor] = id
	}
	typeVariations := make(map[uint32]extensions.TypeVariationID)
	for id, anchor := range original.TypeVariations() {
		typeVariations[anchor] = id
	}
	functions := make(map[uint32]extensions.FunctionID)
	for id, anchor := range original.Functions() {
		functions[anchor] = id
	}
	reconstructed, err := extensions.NewSetFromParts(original.URNs(), types, typeVariations, functions)
	require.NoError(t, err)
	return reconstructed
}
