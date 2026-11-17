// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	substraitpb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/proto"
)

func TestNamedTableWriteOutputModeRoundTrip(t *testing.T) {
	schema := types.NamedStruct{
		Names: []string{"id", "name"},
		Struct: types.StructType{
			Nullability: types.NullabilityRequired,
			Types:       []types.Type{&types.Int64Type{}, &types.StringType{Nullability: types.NullabilityNullable}},
		},
	}
	for _, tc := range []struct {
		name     string
		mode     plan.OutputMode
		wireMode substraitpb.WriteRel_OutputMode
		emit     []int32
		expected []types.Type
	}{
		{"no_output", plan.OutputModeNoOutput, substraitpb.WriteRel_OUTPUT_MODE_NO_OUTPUT, nil, nil},
		{"modified_records", plan.OutputModeModifiedRecords, substraitpb.WriteRel_OUTPUT_MODE_MODIFIED_RECORDS, nil, schema.Struct.Types},
		{"remapped_records", plan.OutputModeModifiedRecords, substraitpb.WriteRel_OUTPUT_MODE_MODIFIED_RECORDS, []int32{1, 0}, []types.Type{schema.Struct.Types[1], schema.Struct.Types[0]}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := plan.NewBuilderDefault().NamedScan([]string{"source"}, schema)
			var original plan.Rel = plan.NewNamedTableWriteRel(
				schema, plan.WriteOpInsert, input, tc.mode, plan.RelCommon{}, []string{"target"}, nil)
			if tc.emit != nil {
				var err error
				original, err = original.Remap(tc.emit...)
				require.NoError(t, err)
			}
			encoded := substraitproto.RelToProto(original)
			require.Equal(t, tc.wireMode, encoded.GetWrite().GetOutput())
			data, err := proto.Marshal(encoded)
			require.NoError(t, err)
			var parsed substraitpb.Rel
			require.NoError(t, proto.Unmarshal(data, &parsed))
			registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
			decoded, err := substraitproto.RelFromProto(&parsed, registry)
			require.NoError(t, err)
			write, ok := decoded.(*plan.NamedTableWriteRel)
			require.True(t, ok)
			require.Equal(t, tc.mode, write.OutputMode())
			require.NotPanics(t, func() {
				require.Equal(t, tc.expected, write.RecordType().Types())
			})
			require.True(t, proto.Equal(encoded, substraitproto.RelToProto(decoded)))
		})
	}
}
