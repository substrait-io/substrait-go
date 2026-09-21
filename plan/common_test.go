// SPDX-License-Identifier: Apache-2.0

package plan_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestHintMatchesDescriptor(t *testing.T) {
	assertMatchesDescriptor(t, (&proto.RelCommon_Hint{}).ProtoReflect().Descriptor(), reflect.TypeOf(plan.Hint{}), map[protoreflect.Name]descriptorField{
		"stats":               {1, protoreflect.MessageKind},
		"constraint":          {2, protoreflect.MessageKind},
		"alias":               {3, protoreflect.StringKind},
		"output_names":        {4, protoreflect.StringKind},
		"advanced_extension":  {10, protoreflect.MessageKind},
		"saved_computations":  {11, protoreflect.MessageKind},
		"loaded_computations": {12, protoreflect.MessageKind},
	})
}

func TestStatsMatchesDescriptor(t *testing.T) {
	assertMatchesDescriptor(t, (&proto.RelCommon_Hint_Stats{}).ProtoReflect().Descriptor(), reflect.TypeOf(plan.Stats{}), map[protoreflect.Name]descriptorField{
		"row_count":          {1, protoreflect.DoubleKind},
		"record_size":        {2, protoreflect.DoubleKind},
		"advanced_extension": {10, protoreflect.MessageKind},
	})
}

func TestRuntimeConstraintMatchesDescriptor(t *testing.T) {
	assertMatchesDescriptor(t, (&proto.RelCommon_Hint_RuntimeConstraint{}).ProtoReflect().Descriptor(), reflect.TypeOf(plan.RuntimeConstraint{}), map[protoreflect.Name]descriptorField{
		"advanced_extension": {10, protoreflect.MessageKind},
	})
}

func TestSavedComputationMatchesDescriptor(t *testing.T) {
	assertMatchesDescriptor(t, (&proto.RelCommon_Hint_SavedComputation{}).ProtoReflect().Descriptor(), reflect.TypeOf(plan.SavedComputation{}), map[protoreflect.Name]descriptorField{
		"computation_id":     {1, protoreflect.Int32Kind},
		"type":               {2, protoreflect.EnumKind},
		"advanced_extension": {10, protoreflect.MessageKind},
	})
}

func TestLoadedComputationMatchesDescriptor(t *testing.T) {
	assertMatchesDescriptor(t, (&proto.RelCommon_Hint_LoadedComputation{}).ProtoReflect().Descriptor(), reflect.TypeOf(plan.LoadedComputation{}), map[protoreflect.Name]descriptorField{
		"computation_id_reference": {1, protoreflect.Int32Kind},
		"type":                     {2, protoreflect.EnumKind},
		"advanced_extension":       {10, protoreflect.MessageKind},
	})
}

type descriptorField struct {
	number protoreflect.FieldNumber
	kind   protoreflect.Kind
}

func assertMatchesDescriptor(t *testing.T, descriptor protoreflect.MessageDescriptor, domainType reflect.Type, want map[protoreflect.Name]descriptorField) {
	t.Helper()

	fields := descriptor.Fields()
	require.Equal(t, len(want), fields.Len(), "spec %s field set changed", descriptor.FullName())
	require.Equal(t, len(want), domainType.NumField(), "%s field count drifted from the spec", domainType)

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		expected, ok := want[field.Name()]
		require.Truef(t, ok, "unexpected spec field %q", field.Name())
		assert.EqualValues(t, expected.number, field.Number(), "%s wire number", field.Name())
		assert.Equal(t, expected.kind, field.Kind(), "%s kind", field.Name())
	}
}

// Round-trip a fully populated Hint through the proto boundary, so a field-mapping
// mistake or a dropped field in the converters fails here.
func TestHintRoundTrip(t *testing.T) {
	ae := &extensions.AdvancedExtension{
		Optimizations: []*extensions.Optimization{{TypeUrl: "opt", Value: []byte("o")}},
		Enhancement:   &extensions.Enhancement{TypeUrl: "enh", Value: []byte("e")},
	}
	h := &plan.Hint{
		Stats:              &plan.Stats{RowCount: 10, RecordSize: 20, AdvancedExtension: ae},
		Constraint:         &plan.RuntimeConstraint{AdvancedExtension: ae},
		Alias:              "myrel",
		OutputNames:        []string{"a", "b"},
		AdvancedExtension:  ae,
		SavedComputations:  []*plan.SavedComputation{{ComputationID: 1, Type: plan.ComputationTypeHashTable, AdvancedExtension: ae}},
		LoadedComputations: []*plan.LoadedComputation{{ComputationIDReference: 1, Type: plan.ComputationTypeBloomFilter, AdvancedExtension: ae}},
	}
	assert.Equal(t, h, plan.HintFromProto(plan.HintToProto(h)))

	assert.Nil(t, plan.HintToProto(nil))
	assert.Nil(t, plan.HintFromProto(nil))
}

// TestRelCommonConvertersNil verifies each Hint-tree converter maps a nil message
// to nil rather than panicking.
func TestRelCommonConvertersNil(t *testing.T) {
	assert.Nil(t, plan.HintFromProto(nil))
	assert.Nil(t, plan.HintToProto(nil))
	assert.Nil(t, plan.StatsFromProto(nil))
	assert.Nil(t, plan.StatsToProto(nil))
	assert.Nil(t, plan.RuntimeConstraintFromProto(nil))
	assert.Nil(t, plan.RuntimeConstraintToProto(nil))
	assert.Nil(t, plan.SavedComputationFromProto(nil))
	assert.Nil(t, plan.SavedComputationToProto(nil))
	assert.Nil(t, plan.LoadedComputationFromProto(nil))
	assert.Nil(t, plan.LoadedComputationToProto(nil))
}

func TestComputationTypeString(t *testing.T) {
	for _, td := range []struct {
		c        plan.ComputationType
		expected string
	}{
		{plan.ComputationTypeUnspecified, "COMPUTATION_TYPE_UNSPECIFIED"},
		{plan.ComputationTypeHashTable, "COMPUTATION_TYPE_HASHTABLE"},
		{plan.ComputationTypeBloomFilter, "COMPUTATION_TYPE_BLOOM_FILTER"},
		{plan.ComputationTypeUnknown, "COMPUTATION_TYPE_UNKNOWN"},
		{plan.ComputationType(7), "7"},
	} {
		t.Run(td.expected, func(t *testing.T) {
			assert.Equal(t, td.expected, td.c.String())
		})
	}
}

// TestComputationTypeMatchesProto keeps the domain enum in lockstep with the spec:
// every value maps to its protobuf counterpart, and a new spec value fails the count.
func TestComputationTypeMatchesProto(t *testing.T) {
	cases := []struct {
		domain plan.ComputationType
		pb     proto.RelCommon_Hint_ComputationType
	}{
		{plan.ComputationTypeUnspecified, proto.RelCommon_Hint_COMPUTATION_TYPE_UNSPECIFIED},
		{plan.ComputationTypeHashTable, proto.RelCommon_Hint_COMPUTATION_TYPE_HASHTABLE},
		{plan.ComputationTypeBloomFilter, proto.RelCommon_Hint_COMPUTATION_TYPE_BLOOM_FILTER},
		{plan.ComputationTypeUnknown, proto.RelCommon_Hint_COMPUTATION_TYPE_UNKNOWN},
	}
	assert.Equal(t, proto.RelCommon_Hint_COMPUTATION_TYPE_UNSPECIFIED.Descriptor().Values().Len(), len(cases),
		"the set of spec computation types changed")
	for _, c := range cases {
		assert.EqualValues(t, c.pb, c.domain, "%s wire value", c.domain)
		assert.Equal(t, c.pb.String(), c.domain.String(), "%s name", c.domain)
	}
}
