// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
)

func TestExtensionSetPreservesImportedAnchors(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()
	const arithmeticURN = extensions.SubstraitDefaultURNPrefix + "functions_arithmetic"
	const variationURN = extensions.SubstraitDefaultURNPrefix + "type_variations"

	t.Run("function", func(t *testing.T) {
		testImportedAnchors(t, c, arithmeticURN,
			extensions.FunctionID{URN: arithmeticURN, Name: "add:i32_i32"},
			extensions.FunctionID{URN: arithmeticURN, Name: "subtract:i32_i32"},
			extensions.Set.GetFuncAnchor, extensions.Set.DecodeFunc,
			func(anchor uint32, id extensions.FunctionID) *extensionspb.SimpleExtensionDeclaration {
				return &extensionspb.SimpleExtensionDeclaration{MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 1, FunctionAnchor: anchor, Name: id.Name,
					},
				}}
			}, true)
	})

	t.Run("type", func(t *testing.T) {
		var c extensions.Collection
		require.NoError(t, c.Load(strings.NewReader(sampleYAML)))
		const urn = "extension:test:sample"
		testImportedAnchors(t, &c, urn,
			extensions.TypeID{URN: urn, Name: "point"},
			extensions.TypeID{URN: urn, Name: "line"},
			extensions.Set.GetTypeAnchor, extensions.Set.DecodeType,
			func(anchor uint32, id extensions.TypeID) *extensionspb.SimpleExtensionDeclaration {
				return &extensionspb.SimpleExtensionDeclaration{MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionType_{
					ExtensionType: &extensionspb.SimpleExtensionDeclaration_ExtensionType{
						ExtensionUrnReference: 1, TypeAnchor: anchor, Name: id.Name,
					},
				}}
			}, true)
	})

	t.Run("type variation", func(t *testing.T) {
		testImportedAnchors(t, c, variationURN,
			extensions.TypeVariationID{URN: variationURN, Name: "dict4"},
			extensions.TypeVariationID{URN: variationURN, Name: "bigoffset"},
			extensions.Set.GetTypeVariationAnchor, extensions.Set.DecodeTypeVariation,
			func(anchor uint32, id extensions.TypeVariationID) *extensionspb.SimpleExtensionDeclaration {
				return &extensionspb.SimpleExtensionDeclaration{MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionTypeVariation_{
					ExtensionTypeVariation: &extensionspb.SimpleExtensionDeclaration_ExtensionTypeVariation{
						ExtensionUrnReference: 1, TypeVariationAnchor: anchor, Name: id.Name,
					},
				}}
			}, false)
	})
}

func testImportedAnchors[T comparable](t *testing.T, c *extensions.Collection, urn string, importedID, newID T,
	getAnchor func(extensions.Set, T) uint32, decode func(extensions.Set, uint32) (T, bool),
	declaration func(uint32, T) *extensionspb.SimpleExtensionDeclaration, zeroAllowed bool) {
	t.Helper()
	anchors := []uint32{2, math.MaxUint32}
	if zeroAllowed {
		anchors = append(anchors, 0)
	}
	for _, importedAnchor := range anchors {
		t.Run(fmt.Sprint(importedAnchor), func(t *testing.T) {
			plan := &proto.Plan{
				ExtensionUrns: []*extensionspb.SimpleExtensionURN{{ExtensionUrnAnchor: 1, Urn: urn}},
				Extensions:    []*extensionspb.SimpleExtensionDeclaration{declaration(importedAnchor, importedID)},
			}
			s, err := extensions.GetExtensionSet(plan, c)
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
			plan.ExtensionUrns, plan.Extensions = s.ToProto(c)
			require.Len(t, plan.Extensions, 2)
			s, err = extensions.GetExtensionSet(plan, c)
			require.NoError(t, err)
			check(s)
		})
	}
}

func TestExtensionSetPreservesImportedURNAnchors(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()
	importedID := extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + "functions_arithmetic", Name: "add:i32_i32"}
	newID := extensions.FunctionID{URN: extensions.SubstraitDefaultURNPrefix + "functions_boolean", Name: "and:bool_bool"}
	for _, importedAnchor := range []uint32{0, 2, math.MaxUint32} {
		t.Run(fmt.Sprint(importedAnchor), func(t *testing.T) {
			plan := &proto.Plan{
				ExtensionUrns: []*extensionspb.SimpleExtensionURN{{ExtensionUrnAnchor: importedAnchor, Urn: importedID.URN}},
				Extensions: []*extensionspb.SimpleExtensionDeclaration{{MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: importedAnchor, FunctionAnchor: 1, Name: importedID.Name,
					},
				}}},
			}
			s, err := extensions.GetExtensionSet(plan, c)
			require.NoError(t, err)
			var functionAnchor uint32
			require.NotPanics(t, func() { functionAnchor = s.GetFuncAnchor(newID) })
			newAnchor, ok := s.FindURN(newID.URN)
			require.True(t, ok)
			assert.NotZero(t, newAnchor)
			assert.NotEqual(t, importedAnchor, newAnchor)
			anchor, ok := s.FindURN(importedID.URN)
			require.True(t, ok)
			assert.Equal(t, importedAnchor, anchor)

			plan.ExtensionUrns, plan.Extensions = s.ToProto(c)
			require.Len(t, plan.ExtensionUrns, 2)
			s, err = extensions.GetExtensionSet(plan, c)
			require.NoError(t, err)
			id, ok := s.DecodeFunc(1)
			require.True(t, ok)
			assert.Equal(t, importedID, id)
			id, ok = s.DecodeFunc(functionAnchor)
			require.True(t, ok)
			assert.Equal(t, newID, id)
		})
	}
}
