// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func TestFileOrFilesFormatRoundTrip(t *testing.T) {
	cases := []struct {
		name   string
		format plan.FileFormat
		check  func(t *testing.T, p *proto.ReadRel_LocalFiles_FileOrFiles)
	}{
		{"parquet", &plan.ParquetReadOptions{}, func(t *testing.T, p *proto.ReadRel_LocalFiles_FileOrFiles) {
			assert.NotNil(t, p.GetParquet())
		}},
		{"arrow", &plan.ArrowReadOptions{}, func(t *testing.T, p *proto.ReadRel_LocalFiles_FileOrFiles) {
			assert.NotNil(t, p.GetArrow())
		}},
		{"orc", &plan.OrcReadOptions{}, func(t *testing.T, p *proto.ReadRel_LocalFiles_FileOrFiles) {
			assert.NotNil(t, p.GetOrc())
		}},
		{"dwrf", &plan.DwrfReadOptions{}, func(t *testing.T, p *proto.ReadRel_LocalFiles_FileOrFiles) {
			assert.NotNil(t, p.GetDwrf())
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			in := &plan.FileOrFiles{PathType: plan.URIFile, Path: "/tmp/x", Format: tc.format}
			p := fileOrFilesToProto(in)
			require.NotNil(t, p)
			tc.check(t, p)

			got := fileOrFilesFromProto(p)
			assert.IsType(t, tc.format, got.Format)
			assert.Equal(t, plan.URIFile, got.PathType)
			assert.Equal(t, "/tmp/x", got.Path)
		})
	}
}
