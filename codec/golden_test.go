// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	substraitpb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

var updateGolden = flag.Bool("update", false, "rewrite golden files from current output")

// ignoreVersion excludes Version: it is build-info/release-derived, and tested in types.
var ignoreVersion = protocmp.IgnoreMessages(&substraitpb.Version{})

// protoDiff compares two messages semantically (empty == equal); pass opts like ignoreVersion.
func protoDiff(want, got proto.Message, opts ...cmp.Option) string {
	return cmp.Diff(want, got, append([]cmp.Option{protocmp.Transform()}, opts...)...)
}

// checkGolden compares got against its committed golden; -update rewrites it.
func checkGolden(t *testing.T, group, name string, got proto.Message) {
	t.Helper()
	dir := filepath.Join("testdata", group)
	path := filepath.Join(dir, name+".golden.txt")

	if *updateGolden {
		canonical := proto.Clone(got)
		stripVersion(canonical)
		require.NoError(t, os.MkdirAll(dir, 0o755))
		require.NoError(t, os.WriteFile(path, []byte(prototext.Format(canonical)), 0o644))
		return
	}

	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	want := got.ProtoReflect().New().Interface()
	require.NoError(t, prototext.Unmarshal(raw, want))

	if diff := protoDiff(want, got, ignoreVersion); diff != "" {
		t.Errorf("%s/%s golden mismatch (-want +got):\n%s", group, name, diff)
	}
}

func stripVersion(m proto.Message) {
	r := m.ProtoReflect()
	if f := r.Descriptor().Fields().ByName("version"); f != nil {
		r.Clear(f)
	}
}

func TestStripVersion(t *testing.T) {
	m := &substraitpb.Plan{Version: &substraitpb.Version{Producer: "x", MinorNumber: 1}}
	stripVersion(m)
	require.Nil(t, m.Version)
}
