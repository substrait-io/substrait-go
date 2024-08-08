package expr

import (
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestToProtoLiteral(t *testing.T) {
	for _, tc := range []struct {
		name                      string
		constructedLiteral        *ProtoLiteral
		expectedExpressionLiteral *proto.Expression_Literal
	}{
		{"TimeStampType",
			&ProtoLiteral{Value: uint64(12345678), Type: types.NewPrecisionTimestampType(types.EMinus4Seconds).WithNullability(types.NullabilityNullable)},
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestamp_{PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{Precision: 4, Value: 12345678}}, Nullable: true},
		},
		{"TimeStampTzType",
			&ProtoLiteral{Value: uint64(12345678), Type: types.NewPrecisionTimestampTzType(types.NanoSeconds).WithNullability(types.NullabilityNullable)},
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestampTz{PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{Precision: 9, Value: 12345678}}, Nullable: true},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			toProto := tc.constructedLiteral.ToProtoLiteral()
			if diff := cmp.Diff(toProto, tc.expectedExpressionLiteral, protocmp.Transform()); diff != "" {
				t.Errorf("proto didn't match, diff:\n%v", diff)
			}
		})

	}
}

func TestLiteralFromProto(t *testing.T) {
	for _, tc := range []struct {
		name             string
		constructedProto *proto.Expression_Literal
		expectedLiteral  interface{}
	}{
		{"TimeStampType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestamp_{PrecisionTimestamp: &proto.Expression_Literal_PrecisionTimestamp{Precision: 4, Value: 12345678}}, Nullable: true},
			&ProtoLiteral{Value: uint64(12345678), Type: types.NewPrecisionTimestampType(types.EMinus4Seconds).WithNullability(types.NullabilityNullable)},
		},
		{"TimeStampTzType",
			&proto.Expression_Literal{LiteralType: &proto.Expression_Literal_PrecisionTimestampTz{PrecisionTimestampTz: &proto.Expression_Literal_PrecisionTimestamp{Precision: 9, Value: 12345678}}, Nullable: true},
			&ProtoLiteral{Value: uint64(12345678), Type: types.NewPrecisionTimestampTzType(types.NanoSeconds).WithNullability(types.NullabilityNullable)},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			literal := LiteralFromProto(tc.constructedProto)
			assert.Equal(t, tc.expectedLiteral, literal)
		})

	}
}
