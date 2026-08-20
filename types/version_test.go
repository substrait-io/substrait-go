// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestVersionGetters(t *testing.T) {
	// distinct values so a getter reading the wrong field shows up
	v := &types.Version{
		MajorNumber: 1,
		MinorNumber: 29,
		PatchNumber: 3,
		GitHash:     "0123456789abcdef",
		Producer:    "substrait-go",
	}

	assert.Equal(t, uint32(1), v.GetMajorNumber())
	assert.Equal(t, uint32(29), v.GetMinorNumber())
	assert.Equal(t, uint32(3), v.GetPatchNumber())
	assert.Equal(t, "0123456789abcdef", v.GetGitHash())
	assert.Equal(t, "substrait-go", v.GetProducer())
}

// Plan.Version() hands out a nil *Version through an interface, so callers reach the getters
// without a way to check for nil first.
func TestVersionGettersOnNil(t *testing.T) {
	var v *types.Version

	assert.Zero(t, v.GetMajorNumber())
	assert.Zero(t, v.GetMinorNumber())
	assert.Zero(t, v.GetPatchNumber())
	assert.Empty(t, v.GetGitHash())
	assert.Empty(t, v.GetProducer())
	assert.Equal(t, "<nil>", v.String())
}

func TestVersionString(t *testing.T) {
	for _, td := range []struct {
		name     string
		version  *types.Version
		expected string
	}{
		{"numbers only", &types.Version{MajorNumber: 0, MinorNumber: 29, PatchNumber: 0}, "0.29.0"},
		{"zero value", &types.Version{}, "0.0.0"},
		{"with producer", &types.Version{MinorNumber: 29, Producer: "substrait-go v8"}, "0.29.0 (substrait-go v8)"},
		{"with git hash", &types.Version{MinorNumber: 29, GitHash: "abc123"}, "0.29.0+abc123"},
		{
			"everything",
			&types.Version{MajorNumber: 1, MinorNumber: 2, PatchNumber: 3, GitHash: "abc123", Producer: "acme"},
			"1.2.3+abc123 (acme)",
		},
		{"nil", nil, "<nil>"},
	} {
		t.Run(td.name, func(t *testing.T) {
			assert.Equal(t, td.expected, td.version.String())
			assert.Equal(t, td.expected, fmt.Sprintf("%v", td.version))
		})
	}
}

// pin the hand written struct to the generated Version descriptor. Field numbers are spelled out
// rather than read from the descriptor so that a renumber in the spec fails here instead of
// silently changing what substrait-go writes on the wire.
func TestVersionMatchesDescriptor(t *testing.T) {
	ours := map[protoreflect.Name]struct {
		number protoreflect.FieldNumber
		field  string
		kind   reflect.Kind
	}{
		"major_number": {1, "MajorNumber", reflect.Uint32},
		"minor_number": {2, "MinorNumber", reflect.Uint32},
		"patch_number": {3, "PatchNumber", reflect.Uint32},
		"git_hash":     {4, "GitHash", reflect.String},
		"producer":     {5, "Producer", reflect.String},
	}

	domain := reflect.TypeOf(types.Version{})
	fields := (&proto.Version{}).ProtoReflect().Descriptor().Fields()
	require.Equal(t, fields.Len(), len(ours), "the set of spec version fields changed")
	require.Equal(t, len(ours), domain.NumField(), "types.Version carries a field the spec does not")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		ourField, ok := ours[f.Name()]
		require.Truef(t, ok, "no types.Version field covers %s", f.Name())
		assert.EqualValues(t, ourField.number, f.Number(), "wire number for %s", f.Name())

		sf, ok := domain.FieldByName(ourField.field)
		require.Truef(t, ok, "types.Version has no %s field for %s", ourField.field, f.Name())
		assert.Equal(t, ourField.kind, sf.Type.Kind(), "Go kind for %s", f.Name())
	}
}

// the conversions carry every field, in both directions, and leave an absent version absent
func TestVersionRoundTrip(t *testing.T) {
	wire := &proto.Version{
		MajorNumber: 1,
		MinorNumber: 29,
		PatchNumber: 3,
		GitHash:     "0123456789abcdef",
		Producer:    "substrait-go",
	}

	domain := types.VersionFromProto(wire)
	require.NotNil(t, domain)
	assert.Equal(t, &types.Version{
		MajorNumber: 1,
		MinorNumber: 29,
		PatchNumber: 3,
		GitHash:     "0123456789abcdef",
		Producer:    "substrait-go",
	}, domain)

	if diff := cmp.Diff(wire, types.VersionToProto(domain), protocmp.Transform()); diff != "" {
		t.Errorf("version proto didn't match, diff:\n%v", diff)
	}

	assert.Nil(t, types.VersionFromProto(nil))
	assert.Nil(t, types.VersionToProto(nil))
}

// The spec asks producers for a 40 character lowercase hex git hash, and the conversions carry
// whatever they are given rather than enforcing that. These cases sit on both sides of the length
// the spec asks for, so a conversion that starts trimming or lowercasing fails here.
func TestVersionCarriesAnyGitHash(t *testing.T) {
	for _, tc := range []struct {
		name    string
		gitHash string
	}{
		{"unset", ""},
		{"not a hash at all", "HEAD"},
		{"the length the spec asks for", "0123456789abcdef0123456789abcdef01234567"},
		{"longer than the spec asks for", "0123456789ABCDEF0123456789abcdef0123456789abcdef"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wire := &proto.Version{MinorNumber: 29, GitHash: tc.gitHash}

			domain := types.VersionFromProto(wire)
			require.NotNil(t, domain)
			assert.Equal(t, tc.gitHash, domain.GetGitHash())

			if diff := cmp.Diff(wire, types.VersionToProto(domain), protocmp.Transform()); diff != "" {
				t.Errorf("version proto didn't match, diff:\n%v", diff)
			}
		})
	}
}
