// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	ext "github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/wire"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// TestLambdaProtoRoundTrip tests that lambda expressions can be round-tripped
// through protobuf serialization without losing information.
func TestLambdaProtoRoundTrip(t *testing.T) {
	files, err := filepath.Glob("./testdata/lambda/*.json")
	require.NoError(t, err)

	collection := ext.GetDefaultCollectionWithNoError()

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			data, err := os.ReadFile(file)
			require.NoError(t, err)

			var originalPlan proto.Plan
			require.NoError(t, protojson.Unmarshal(data, &originalPlan))

			goPlan, err := wire.PlanFromProto(&originalPlan, collection)
			require.NoError(t, err)

			resultPlan, err := wire.PlanToProto(goPlan)
			require.NoError(t, err)

			require.True(t, pb.Equal(&originalPlan, resultPlan))
		})
	}
}
