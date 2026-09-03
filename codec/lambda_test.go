// SPDX-License-Identifier: Apache-2.0

package codec_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/codec"
	"github.com/substrait-io/substrait-go/v9/extensions"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	protobuf "google.golang.org/protobuf/proto"
)

func TestLambdaProtoRoundTrip(t *testing.T) {
	files, err := filepath.Glob("./testdata/lambda/*.json")
	require.NoError(t, err)
	require.NotEmpty(t, files, "no lambda fixtures found")

	collection := extensions.GetDefaultCollectionWithNoError()

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			data, err := os.ReadFile(file)
			require.NoError(t, err)

			var originalPlan proto.Plan
			require.NoError(t, protojson.Unmarshal(data, &originalPlan))

			goPlan, err := codec.PlanFromProto(&originalPlan, collection)
			require.NoError(t, err)

			resultPlan, err := codec.PlanToProto(goPlan)
			require.NoError(t, err)

			require.True(t, protobuf.Equal(&originalPlan, resultPlan))
		})
	}
}
