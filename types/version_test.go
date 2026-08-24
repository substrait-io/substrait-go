// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
)

// UnsetVersion renders as a visible sentinel so a missing version stands out in output.
func TestUnsetVersionString(t *testing.T) {
	assert.Equal(t, "0.0.0 (UNSET)", types.UnsetVersion.String())
}

func TestVersionString(t *testing.T) {
	for _, td := range []struct {
		name     string
		version  *types.Version
		expected string
	}{
		{"numbers only", &types.Version{MinorNumber: 29}, "0.29.0"},
		{"zero value", &types.Version{}, "0.0.0"},
		{"with producer", &types.Version{MinorNumber: 29, Producer: "substrait-go v8"}, "0.29.0 (substrait-go v8)"},
		{"with git hash", &types.Version{MinorNumber: 29, GitHash: "abc123"}, "0.29.0+abc123"},
		{
			"everything",
			&types.Version{MajorNumber: 1, MinorNumber: 2, PatchNumber: 3, GitHash: "abc123", Producer: "acme"},
			"1.2.3+abc123 (acme)",
		},
	} {
		t.Run(td.name, func(t *testing.T) {
			assert.Equal(t, td.expected, td.version.String())
		})
	}
}

// Pin the hand written struct to the generated descriptor: the field set and wire numbers, which
// a spec change can move without breaking the build.
func TestVersionMatchesDescriptor(t *testing.T) {
	ours := map[protoreflect.Name]protoreflect.FieldNumber{
		"major_number": 1,
		"minor_number": 2,
		"patch_number": 3,
		"git_hash":     4,
		"producer":     5,
	}

	fields := (&proto.Version{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, len(ours), fields.Len(), "the set of spec version fields changed")
	require.Equal(t, len(ours), reflect.TypeOf(types.Version{}).NumField(),
		"types.Version carries a field the spec does not")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		number, ok := ours[f.Name()]
		require.Truef(t, ok, "no types.Version field covers %s", f.Name())
		assert.EqualValues(t, number, f.Number(), "wire number for %s", f.Name())
	}
}

// Both directions carry every field and leave an absent version absent. A non-hash git value
// ("HEAD") shows the hash is copied as given, not validated.
func TestVersionRoundTrip(t *testing.T) {
	for _, td := range []struct {
		name string
		wire *proto.Version
	}{
		// distinct numbers, so a conversion reading the wrong field shows up
		{"every field", &proto.Version{
			MajorNumber: 1, MinorNumber: 29, PatchNumber: 3,
			GitHash: "0123456789abcdef", Producer: "substrait-go",
		}},
		{"no git hash", &proto.Version{MinorNumber: 29}},
		{"not a hash at all", &proto.Version{MinorNumber: 29, GitHash: "HEAD"}},
	} {
		t.Run(td.name, func(t *testing.T) {
			domain := types.VersionFromProto(td.wire)
			require.NotNil(t, domain)
			assert.Equal(t, td.wire.GetMajorNumber(), domain.MajorNumber)
			assert.Equal(t, td.wire.GetMinorNumber(), domain.MinorNumber)
			assert.Equal(t, td.wire.GetPatchNumber(), domain.PatchNumber)
			assert.Equal(t, td.wire.GetGitHash(), domain.GitHash)
			assert.Equal(t, td.wire.GetProducer(), domain.Producer)

			if diff := cmp.Diff(td.wire, types.VersionToProto(domain), protocmp.Transform()); diff != "" {
				t.Errorf("version proto didn't match, diff:\n%v", diff)
			}
		})
	}

	assert.Nil(t, types.VersionFromProto(nil))
	assert.Nil(t, types.VersionToProto(nil))
}
