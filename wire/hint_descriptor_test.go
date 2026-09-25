// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type descriptorField struct {
	number protoreflect.FieldNumber
	kind   protoreflect.Kind
}

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
