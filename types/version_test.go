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

// A version-less proto maps to the unset placeholder, which renders visibly and serializes back out
// as a producer "UNSET" version rather than an absent one.
func TestVersionFromProtoNilIsUnset(t *testing.T) {
	unset := types.VersionFromProto(nil)
	assert.Equal(t, "0.0.0 (UNSET)", unset.String())
	assert.Equal(t, "UNSET", types.VersionToProto(unset).GetProducer())
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

// Guard against spec drift: fail if the Version proto message gains, drops, or renumbers a field
// that the hand-written types.Version does not mirror.
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

func TestVersionRoundTrip(t *testing.T) {
	for _, td := range []struct {
		name         string
		protoVersion *proto.Version
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
			domain := types.VersionFromProto(td.protoVersion)
			assert.Equal(t, td.protoVersion.GetMajorNumber(), domain.MajorNumber)
			assert.Equal(t, td.protoVersion.GetMinorNumber(), domain.MinorNumber)
			assert.Equal(t, td.protoVersion.GetPatchNumber(), domain.PatchNumber)
			assert.Equal(t, td.protoVersion.GetGitHash(), domain.GitHash)
			assert.Equal(t, td.protoVersion.GetProducer(), domain.Producer)

			if diff := cmp.Diff(td.protoVersion, types.VersionToProto(domain), protocmp.Transform()); diff != "" {
				t.Errorf("version proto didn't match, diff:\n%v", diff)
			}
		})
	}
}
