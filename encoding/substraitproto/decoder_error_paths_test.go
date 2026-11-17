// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	pb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/proto"
)

func TestNarrowIntegerLiteralBounds(t *testing.T) {
	for _, tc := range []struct {
		name     string
		min, max int32
		literal  func(int32) *pb.Expression_Literal
	}{
		{"i8", -128, 127, func(v int32) *pb.Expression_Literal {
			return &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_I8{I8: v}}
		}},
		{"i16", -32768, 32767, func(v int32) *pb.Expression_Literal {
			return &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_I16{I16: v}}
		}},
	} {
		for _, v := range []int32{tc.min - 1, tc.min, 0, tc.max, tc.max + 1} {
			t.Run(fmt.Sprintf("%s/%d", tc.name, v), func(t *testing.T) {
				input := tc.literal(v)
				decoded, err := substraitproto.LiteralFromProto(input)
				if v < tc.min || v > tc.max {
					require.Nil(t, decoded)
					require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
					return
				}
				require.NoError(t, err)
				require.True(t, proto.Equal(input, substraitproto.LiteralToProto(decoded)))
			})
		}
	}
}

func TestLiteralDecoderPreservesNestedErrors(t *testing.T) {
	i64 := &pb.Type{Kind: &pb.Type_I64_{I64: &pb.Type_I64{}}}
	valid := &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_I64{I64: 1}}
	for _, tc := range []struct {
		name  string
		input *pb.Expression_Literal
		want  error
	}{
		{"map literal key 0", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_Map_{Map: &pb.Expression_Literal_Map{KeyValues: []*pb.Expression_Literal_Map_KeyValue{{Key: &pb.Expression_Literal{}, Value: valid}}}}}, substraitgo.ErrInvalidExpr},
		{"map literal value 0", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_Map_{Map: &pb.Expression_Literal_Map{KeyValues: []*pb.Expression_Literal_Map_KeyValue{{Key: valid, Value: &pb.Expression_Literal{}}}}}}, substraitgo.ErrInvalidExpr},
		{"empty map key type", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_EmptyMap{EmptyMap: &pb.Type_Map{Value: i64}}}, substraitgo.ErrInvalidType},
		{"empty map value type", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_EmptyMap{EmptyMap: &pb.Type_Map{Key: i64}}}, substraitgo.ErrInvalidType},
		{"literal type parameter 0", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_UserDefined_{UserDefined: &pb.Expression_Literal_UserDefined{TypeParameters: []*pb.Type_Parameter{{}}}}}, substraitgo.ErrInvalidType},
		{"user-defined literal: struct field 0", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_UserDefined_{UserDefined: &pb.Expression_Literal_UserDefined{Val: &pb.Expression_Literal_UserDefined_Struct{Struct: &pb.Expression_Literal_Struct{Fields: []*pb.Expression_Literal{{}}}}}}}, substraitgo.ErrInvalidExpr},
		{"nil value wrapper", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_UserDefined_{UserDefined: &pb.Expression_Literal_UserDefined{Val: (*pb.Expression_Literal_UserDefined_Value)(nil)}}}, substraitgo.ErrInvalidExpr},
		{"negative value", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_PrecisionTime_{PrecisionTime: &pb.Expression_Literal_PrecisionTime{Value: -1}}}, substraitgo.ErrInvalidExpr},
		{"invalid timestamp precision", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_PrecisionTimestamp_{PrecisionTimestamp: &pb.Expression_Literal_PrecisionTimestamp{Precision: 10}}}, substraitgo.ErrInvalidExpr},
		{"invalid timestamp tz precision", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_PrecisionTimestampTz{PrecisionTimestampTz: &pb.Expression_Literal_PrecisionTimestamp{Precision: 10}}}, substraitgo.ErrInvalidExpr},
		{"invalid interval precision", &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_IntervalDayToSecond_{IntervalDayToSecond: &pb.Expression_Literal_IntervalDayToSecond{PrecisionMode: &pb.Expression_Literal_IntervalDayToSecond_Precision{Precision: 10}}}}, substraitgo.ErrInvalidExpr},
	} {
		t.Run(tc.name, func(t *testing.T) {
			value, err := substraitproto.LiteralFromProto(tc.input)
			require.Nil(t, value)
			require.ErrorIs(t, err, tc.want)
			// Check location context where the decoder adds it, too.
			if tc.want == substraitgo.ErrInvalidType || tc.name == "map literal key 0" || tc.name == "map literal value 0" || tc.name == "user-defined literal: struct field 0" {
				require.ErrorContains(t, err, tc.name)
			}
		})
	}
}

func TestLiteralDecoderContainerRoundTrips(t *testing.T) {
	item := &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_I64{I64: 7}}
	for name, input := range map[string]*pb.Expression_Literal{
		"null": {Nullable: true, LiteralType: &pb.Expression_Literal_Null{Null: &pb.Type{Kind: &pb.Type_I64_{I64: &pb.Type_I64{Nullability: pb.Type_NULLABILITY_NULLABLE}}}}},
		"list": {LiteralType: &pb.Expression_Literal_List_{List: &pb.Expression_Literal_List{Values: []*pb.Expression_Literal{item}}}},
		"map":  {LiteralType: &pb.Expression_Literal_Map_{Map: &pb.Expression_Literal_Map{KeyValues: []*pb.Expression_Literal_Map_KeyValue{{Key: item, Value: item}}}}},
	} {
		t.Run(name, func(t *testing.T) {
			value, err := substraitproto.LiteralFromProto(input)
			require.NoError(t, err)
			require.NotNil(t, value)
			require.True(t, proto.Equal(input, substraitproto.LiteralToProto(value)))
		})
	}
}

func TestReferenceSegmentDecoderErrorsAndRoundTrips(t *testing.T) {
	key := &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_String_{String_: "key"}}
	wraps := map[string]func(*pb.Expression_ReferenceSegment) *pb.Expression_ReferenceSegment{
		"map": func(child *pb.Expression_ReferenceSegment) *pb.Expression_ReferenceSegment {
			return &pb.Expression_ReferenceSegment{ReferenceType: &pb.Expression_ReferenceSegment_MapKey_{MapKey: &pb.Expression_ReferenceSegment_MapKey{MapKey: key, Child: child}}}
		},
		"struct": func(child *pb.Expression_ReferenceSegment) *pb.Expression_ReferenceSegment {
			return &pb.Expression_ReferenceSegment{ReferenceType: &pb.Expression_ReferenceSegment_StructField_{StructField: &pb.Expression_ReferenceSegment_StructField{Field: 2, Child: child}}}
		},
		"list": func(child *pb.Expression_ReferenceSegment) *pb.Expression_ReferenceSegment {
			return &pb.Expression_ReferenceSegment{ReferenceType: &pb.Expression_ReferenceSegment_ListElement_{ListElement: &pb.Expression_ReferenceSegment_ListElement{Offset: 3, Child: child}}}
		},
	}
	for name, wrap := range wraps {
		t.Run(name, func(t *testing.T) {
			for _, child := range []*pb.Expression_ReferenceSegment{nil, wraps["struct"](nil)} {
				input := wrap(child)
				value, err := substraitproto.RefSegmentFromProto(input)
				require.NoError(t, err)
				require.NotNil(t, value)
				require.True(t, proto.Equal(input, substraitproto.RefSegmentToProto(value)))
			}
			input := wrap(&pb.Expression_ReferenceSegment{})
			value, err := substraitproto.RefSegmentFromProto(input)
			require.Nil(t, value)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
		})
	}
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	for name, segment := range map[string]*pb.Expression_ReferenceSegment{
		"missing": nil,
		"unset":   {},
		"payload": {ReferenceType: &pb.Expression_ReferenceSegment_ListElement_{}},
	} {
		t.Run(name, func(t *testing.T) {
			value, err := substraitproto.FieldReferenceFromProto(&pb.Expression_FieldReference{
				RootType:      &pb.Expression_FieldReference_RootReference_{RootReference: &pb.Expression_FieldReference_RootReference{}},
				ReferenceType: &pb.Expression_FieldReference_DirectReference{DirectReference: segment},
			}, nil, registry)
			require.Nil(t, value)
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
		})
	}
}

func TestTypeDecoderAdditionalNestedErrors(t *testing.T) {
	for name, input := range map[string]*pb.Type{
		"interval day":        {Kind: &pb.Type_IntervalDay_{IntervalDay: &pb.Type_IntervalDay{Precision: proto.Int32(10)}}},
		"interval compound":   {Kind: &pb.Type_IntervalCompound_{IntervalCompound: &pb.Type_IntervalCompound{Precision: 10}}},
		"time":                {Kind: &pb.Type_PrecisionTime_{PrecisionTime: &pb.Type_PrecisionTime{Precision: 10}}},
		"timestamp tz":        {Kind: &pb.Type_PrecisionTimestampTz{PrecisionTimestampTz: &pb.Type_PrecisionTimestampTZ{Precision: 10}}},
		"function parameter":  {Kind: &pb.Type_Func_{Func: &pb.Type_Func{ParameterTypes: []*pb.Type{{}}}}},
		"data type parameter": {Kind: &pb.Type_UserDefined_{UserDefined: &pb.Type_UserDefined{TypeParameters: []*pb.Type_Parameter{{Parameter: &pb.Type_Parameter_DataType{DataType: &pb.Type{}}}}}}},
	} {
		t.Run(name, func(t *testing.T) {
			value, err := substraitproto.TypeFromProto(input)
			require.Nil(t, value)
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
		})
	}
}

func TestTypeParameterDecoderRoundTrips(t *testing.T) {
	for name, input := range map[string]*pb.Type_Parameter{
		"bool":    {Parameter: &pb.Type_Parameter_Boolean{Boolean: true}},
		"integer": {Parameter: &pb.Type_Parameter_Integer{Integer: 7}},
		"enum":    {Parameter: &pb.Type_Parameter_Enum{Enum: "value"}},
		"string":  {Parameter: &pb.Type_Parameter_String_{String_: "value"}},
	} {
		t.Run(name, func(t *testing.T) {
			value, err := substraitproto.TypeParamFromProto(input)
			require.NoError(t, err)
			require.NotNil(t, value)
			require.True(t, proto.Equal(input, substraitproto.TypeParamToProto(value)))
		})
	}
}

func TestNullTypeParameterRoundTrip(t *testing.T) {
	want := &types.UserDefinedType{TypeReference: 7, TypeParameters: []types.TypeParam{types.NullParameter{}}}
	input := substraitproto.TypeToProto(want)
	decoded, err := substraitproto.TypeFromProto(input)
	require.NoError(t, err)
	require.Equal(t, want, decoded)

	// Serializing the protobuf must not change whether the decoder accepts it.
	data, err := proto.Marshal(input)
	require.NoError(t, err)
	var wire pb.Type
	require.NoError(t, proto.Unmarshal(data, &wire))
	decoded, err = substraitproto.TypeFromProto(&wire)
	require.NoError(t, err)
	require.Equal(t, want, decoded)
}

func TestNilParameterBindingReturnsError(t *testing.T) {
	value, err := substraitproto.PlanFromProto(&pb.Plan{ParameterBindings: []*pb.DynamicParameterBinding{nil}}, extensions.GetDefaultCollectionWithNoError())
	require.Nil(t, value)
	require.ErrorIs(t, err, substraitgo.ErrInvalidPlan)
	require.ErrorContains(t, err, "parameter binding 0")
}
