// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/plan"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestFileOrFilesFormatRoundTrip(t *testing.T) {
	cases := []struct {
		name   string
		format plan.FileFormat
	}{
		{"parquet", &plan.ParquetReadOptions{}},
		{"arrow", &plan.ArrowReadOptions{}},
		{"orc", &plan.OrcReadOptions{}},
		{"dwrf", &plan.DwrfReadOptions{}},
		{"extension", (*plan.ExtensionReadOptions)(&anypb.Any{TypeUrl: "urn:test", Value: []byte{1, 2, 3}})},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := &plan.FileOrFiles{
				PathType:  plan.URIFile,
				Path:      "/tmp/data",
				PartIndex: 7,
				Start:     100,
				Len:       200,
				Format:    tc.format,
			}

			got := fileOrFilesFromProto(fileOrFilesToProto(original))
			assert.Equal(t, original, &got)

		})
	}
}
