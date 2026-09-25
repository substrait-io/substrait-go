// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
)

const extMgrSampleYAML = `---
urn: extension:test:sample
types:
  - name: point
    structure:
      latitude: i32
      longitude: i32
  - name: line
    structure:
      start: point
      end: point
scalar_functions:
  -
    name: "add"
    description: "Add two values."
    impls:
      - args:
          - name: x
            value: i8
          - name: y
            value: i8
        options:
          overflow:
            values: [ SILENT, SATURATE, ERROR ]
        return: i8
  -
    name: "subtract"
    description: "Subtract one value from another."
    impls:
      - args:
          - name: x
            value: i8
          - name: y
            value: i8
        options:
          overflow:
            values: [ SILENT, SATURATE, ERROR ]
        return: i8
      - args:
          - name: x
            value: i16
          - name: y
            value: i16
        options:
          overflow:
            values: [SILENT, SATURATE, ERROR ]
        return: i16
aggregate_functions:
  - name: "count"
    description: Count a set of values
    impls:
      - args:
          - name: x
            value: any1
        options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i64
  - name: "count"
    description: "Count a set of records (not field referenced)"
    impls:
      - options:
          overflow:
            values: [SILENT, SATURATE, ERROR]
        nullability: DECLARED_OUTPUT
        decomposable: MANY
        intermediate: i64
        return: i64
`

func TestToProtoPopulatesURN(t *testing.T) {
	c := &extensions.Collection{}
	require.NoError(t, c.Load(strings.NewReader(extMgrSampleYAML)))

	plan := &proto.Plan{
		ExtensionUrns: []*extensionspb.SimpleExtensionURN{
			{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
		},
		Extensions: []*extensionspb.SimpleExtensionDeclaration{
			{
				MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
					ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
						ExtensionUrnReference: 1,
						FunctionAnchor:        1,
						Name:                  "add:i8_i8",
					},
				},
			},
		},
	}

	extSet, err := wire.GetExtensionSet(plan, c)
	require.NoError(t, err)

	urns, decls := wire.ExtensionSetToProto(extSet)

	require.Len(t, urns, 1)
	assert.Equal(t, "extension:test:sample", urns[0].Urn)
	require.Len(t, decls, 1)
	assert.Equal(t, "add:i8_i8", decls[0].GetExtensionFunction().Name)
	assert.EqualValues(t, 1, decls[0].GetExtensionFunction().ExtensionUrnReference)
}

func TestResolveRefToURNAllConditions(t *testing.T) {
	c := &extensions.Collection{}
	err := c.Load(strings.NewReader(extMgrSampleYAML))
	require.NoError(t, err)

	t.Run("non-zero URN reference found and valid", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 1, Urn: "extension:test:sample"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 1,
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		extSet, err := wire.GetExtensionSet(plan, c)
		require.NoError(t, err)
		require.NotNil(t, extSet)

		id, ok := extSet.DecodeFunc(1)
		require.True(t, ok)
		assert.Equal(t, "extension:test:sample", id.URN)
	})

	t.Run("non-zero URN reference found but invalid", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 1, Urn: "extension:nonexistent:urn"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 1,
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := wire.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "URN 'extension:nonexistent:urn' not found in extension collection")
	})

	t.Run("zero URN reference with zero anchor defined", func(t *testing.T) {
		plan := &proto.Plan{
			ExtensionUrns: []*extensionspb.SimpleExtensionURN{
				{ExtensionUrnAnchor: 0, Urn: "extension:test:sample"},
			},
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							FunctionAnchor: 1,
							Name:           "test_function",
						},
					},
				},
			},
		}

		extSet, err := wire.GetExtensionSet(plan, c)
		require.NoError(t, err)
		require.NotNil(t, extSet)

		id, ok := extSet.DecodeFunc(1)
		require.True(t, ok)
		assert.Equal(t, "extension:test:sample", id.URN)
	})

	t.Run("URN reference not resolvable", func(t *testing.T) {
		plan := &proto.Plan{
			Extensions: []*extensionspb.SimpleExtensionDeclaration{
				{
					MappingType: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction_{
						ExtensionFunction: &extensionspb.SimpleExtensionDeclaration_ExtensionFunction{
							ExtensionUrnReference: 99,
							FunctionAnchor:        1,
							Name:                  "test_function",
						},
					},
				},
			},
		}

		_, err := wire.GetExtensionSet(plan, c)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unable to resolve extension reference: URN reference 99 could not be resolved")
	})
}
