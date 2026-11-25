// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	"github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
)

func TestExtensionSetAllocatesAfterConsecutiveImportedAnchors(t *testing.T) {
	c := extensions.GetDefaultCollectionWithNoError()
	const urn = extensions.SubstraitDefaultURNPrefix + "functions_arithmetic"
	ids := []extensions.FunctionID{
		{URN: urn, Name: "add:i32_i32"},
		{URN: urn, Name: "subtract:i32_i32"},
		{URN: urn, Name: "multiply:i32_i32"},
		{URN: urn, Name: "divide:i32_i32"},
		{URN: urn, Name: "modulus:i32_i32"},
	}
	for _, firstAnchor := range []uint32{1, 3} {
		t.Run(fmt.Sprint(firstAnchor), func(t *testing.T) {
			plan := &proto.Plan{
				ExtensionUrns: []*extensionspb.SimpleExtensionURN{{ExtensionUrnAnchor: 1, Urn: urn}},
			}
			for i, id := range ids[:2] {
				plan.Extensions = append(plan.Extensions, &extensionspb.SimpleExtensionDeclaration{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 1, FunctionAnchor: firstAnchor + uint32(i), Name: id.Name,
						},
					},
				})
			}
			s, err := substraitproto.GetExtensionSet(plan, c)
			require.NoError(t, err)
			for i := 2; i < 4; i++ {
				assert.Equal(t, firstAnchor+uint32(i), s.GetFuncAnchor(ids[i]))
			}

			plan.ExtensionUrns, plan.Extensions = substraitproto.ExtensionSetToProto(s)
			require.Len(t, plan.Extensions, 4)
			s, err = substraitproto.GetExtensionSet(plan, c)
			require.NoError(t, err)
			assert.Equal(t, firstAnchor+4, s.GetFuncAnchor(ids[4]), "allocation must resume without replacing a round-tripped declaration")
			for i, id := range ids {
				anchor := firstAnchor + uint32(i)
				assert.Equal(t, anchor, s.GetFuncAnchor(id))
				decoded, ok := s.DecodeFunc(anchor)
				assert.True(t, ok)
				assert.Equal(t, id, decoded)
			}
		})
	}
}
