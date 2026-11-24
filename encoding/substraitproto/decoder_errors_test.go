// SPDX-License-Identifier: Apache-2.0

package substraitproto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/encoding/substraitproto"
	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	pb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/proto"
)

func TestMalformedTypeArgumentReturnsError(t *testing.T) {
	for name, input := range map[string]*pb.Type{
		"nil":                     nil,
		"unset":                   {},
		"nil_bool_wrapper":        {Kind: (*pb.Type_Bool)(nil)},
		"missing_bool_payload":    {Kind: &pb.Type_Bool{}},
		"missing_function_result": {Kind: &pb.Type_Func_{Func: &pb.Type_Func{}}},
		"missing_map_value":       {Kind: &pb.Type_Map_{Map: &pb.Type_Map{Key: &pb.Type{Kind: &pb.Type_I64_{I64: &pb.Type_I64{}}}}}},
		"invalid_udt_parameter":   {Kind: &pb.Type_UserDefined_{UserDefined: &pb.Type_UserDefined{TypeParameters: []*pb.Type_Parameter{{}}}}},
		"missing_list_element":    {Kind: &pb.Type_List_{List: &pb.Type_List{}}},
		"invalid_nested_type":     {Kind: &pb.Type_Struct_{Struct: &pb.Type_Struct{Types: []*pb.Type{{}}}}},
		"invalid_precision":       {Kind: &pb.Type_PrecisionTimestamp_{PrecisionTimestamp: &pb.Type_PrecisionTimestamp{Precision: -1}}},
	} {
		t.Run(name, func(t *testing.T) {
			registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
			var err error
			require.NotPanics(t, func() {
				_, err = substraitproto.FuncArgFromProto(&pb.FunctionArgument{ArgType: &pb.FunctionArgument_Type{Type: input}}, nil, registry)
			})
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
		})
	}
}

func TestMalformedLiteralExpressionReturnsError(t *testing.T) {
	for name, input := range map[string]*pb.Expression_Literal{
		"nil":                        nil,
		"unset":                      {},
		"nil_literal_wrapper":        {LiteralType: (*pb.Expression_Literal_Boolean)(nil)},
		"missing_null_type":          {LiteralType: &pb.Expression_Literal_Null{}},
		"missing_struct_payload":     {LiteralType: &pb.Expression_Literal_Struct_{}},
		"invalid_struct_field":       {LiteralType: &pb.Expression_Literal_Struct_{Struct: &pb.Expression_Literal_Struct{Fields: []*pb.Expression_Literal{{}}}}},
		"invalid_list_item":          {LiteralType: &pb.Expression_Literal_List_{List: &pb.Expression_Literal_List{Values: []*pb.Expression_Literal{{}}}}},
		"missing_empty_list_type":    {LiteralType: &pb.Expression_Literal_EmptyList{EmptyList: &pb.Type_List{}}},
		"invalid_map_entry":          {LiteralType: &pb.Expression_Literal_Map_{Map: &pb.Expression_Literal_Map{KeyValues: []*pb.Expression_Literal_Map_KeyValue{nil}}}},
		"nil_udt_value_wrapper":      {LiteralType: &pb.Expression_Literal_UserDefined_{UserDefined: &pb.Expression_Literal_UserDefined{Val: (*pb.Expression_Literal_UserDefined_Struct)(nil)}}},
		"nil_precision_wrapper":      {LiteralType: &pb.Expression_Literal_IntervalDayToSecond_{IntervalDayToSecond: &pb.Expression_Literal_IntervalDayToSecond{PrecisionMode: (*pb.Expression_Literal_IntervalDayToSecond_Precision)(nil)}}},
		"missing_udt_value":          {LiteralType: &pb.Expression_Literal_UserDefined_{UserDefined: &pb.Expression_Literal_UserDefined{}}},
		"missing_interval_precision": {LiteralType: &pb.Expression_Literal_IntervalDayToSecond_{IntervalDayToSecond: &pb.Expression_Literal_IntervalDayToSecond{}}},
		"empty_untyped_list":         {LiteralType: &pb.Expression_Literal_List_{List: &pb.Expression_Literal_List{}}},
		"empty_untyped_map":          {LiteralType: &pb.Expression_Literal_Map_{Map: &pb.Expression_Literal_Map{}}},
		"invalid_precision":          {LiteralType: &pb.Expression_Literal_PrecisionTime_{PrecisionTime: &pb.Expression_Literal_PrecisionTime{Precision: -1}}},
	} {
		t.Run(name, func(t *testing.T) {
			registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
			var err error
			require.NotPanics(t, func() {
				_, err = substraitproto.ExprFromProto(&pb.Expression{RexType: &pb.Expression_Literal_{Literal: input}}, nil, registry)
			})
			if name == "missing_empty_list_type" {
				require.ErrorIs(t, err, substraitgo.ErrInvalidType)
			} else {
				require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
			}
		})
	}
}

func TestDecoderErrorsReachPlanAndRelationCallers(t *testing.T) {
	collection := extensions.GetDefaultCollectionWithNoError()
	registry := expr.NewEmptyExtensionRegistry(collection)
	for name, input := range map[string]*pb.Plan{
		"parameter_binding": {ParameterBindings: []*pb.DynamicParameterBinding{{Value: &pb.Expression_Literal{}}}},
		"virtual_table":     {Relations: []*pb.PlanRel{{RelType: &pb.PlanRel_Rel{Rel: &pb.Rel{RelType: &pb.Rel_Read{Read: &pb.ReadRel{ReadType: &pb.ReadRel_VirtualTable_{VirtualTable: &pb.ReadRel_VirtualTable{Values: []*pb.Expression_Literal_Struct{{Fields: []*pb.Expression_Literal{{}}}}}}}}}}}}},
		"nil_virtual_row":   {Relations: []*pb.PlanRel{{RelType: &pb.PlanRel_Rel{Rel: &pb.Rel{RelType: &pb.Rel_Read{Read: &pb.ReadRel{ReadType: &pb.ReadRel_VirtualTable_{VirtualTable: &pb.ReadRel_VirtualTable{Values: []*pb.Expression_Literal_Struct{nil}}}}}}}}}},
	} {
		t.Run(name, func(t *testing.T) {
			var err error
			require.NotPanics(t, func() { _, err = substraitproto.PlanFromProto(input, collection) })
			require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
		})
	}
	for name, input := range map[string]*pb.Rel{
		"missing_struct":       {RelType: &pb.Rel_Read{Read: &pb.ReadRel{BaseSchema: &pb.NamedStruct{}, ReadType: &pb.ReadRel_NamedTable_{NamedTable: &pb.ReadRel_NamedTable{Names: []string{"t"}}}}}},
		"invalid_schema_field": {RelType: &pb.Rel_Read{Read: &pb.ReadRel{BaseSchema: &pb.NamedStruct{Struct: &pb.Type_Struct{Types: []*pb.Type{{}}}}, ReadType: &pb.ReadRel_NamedTable_{NamedTable: &pb.ReadRel_NamedTable{Names: []string{"t"}}}}}},
	} {
		t.Run(name, func(t *testing.T) {
			var err error
			require.NotPanics(t, func() { _, err = substraitproto.RelFromProto(input, registry) })
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
		})
	}
}

func TestDecoderErrorsReachExpressionCallers(t *testing.T) {
	registry := expr.NewEmptyExtensionRegistry(extensions.GetDefaultCollectionWithNoError())
	boolean := &pb.Expression{RexType: &pb.Expression_Literal_{Literal: &pb.Expression_Literal{LiteralType: &pb.Expression_Literal_Boolean{Boolean: true}}}}
	for name, input := range map[string]*pb.Expression{
		"cast":              {RexType: &pb.Expression_Cast_{Cast: &pb.Expression_Cast{Input: boolean, Type: &pb.Type{}}}},
		"dynamic_parameter": {RexType: &pb.Expression_DynamicParameter{DynamicParameter: &pb.DynamicParameter{Type: &pb.Type{}}}},
		"scalar_output":     {RexType: &pb.Expression_ScalarFunction_{ScalarFunction: &pb.Expression_ScalarFunction{OutputType: &pb.Type{}}}},
		"window_output":     {RexType: &pb.Expression_WindowFunction_{WindowFunction: &pb.Expression_WindowFunction{OutputType: &pb.Type{}}}},
		"lambda_parameter":  {RexType: &pb.Expression_Lambda_{Lambda: &pb.Expression_Lambda{Parameters: &pb.Type_Struct{Types: []*pb.Type{{}}}, Body: boolean}}},
	} {
		t.Run(name, func(t *testing.T) {
			var err error
			require.NotPanics(t, func() { _, err = substraitproto.ExprFromProto(input, nil, registry) })
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
		})
	}
	_, err := substraitproto.AggregateFunctionFromProto(&pb.AggregateFunction{OutputType: &pb.Type{}}, nil, registry)
	require.ErrorIs(t, err, substraitgo.ErrInvalidType)
	_, err = substraitproto.RefSegmentFromProto(&pb.Expression_ReferenceSegment{ReferenceType: &pb.Expression_ReferenceSegment_MapKey_{MapKey: &pb.Expression_ReferenceSegment_MapKey{MapKey: &pb.Expression_Literal{}}}})
	require.ErrorIs(t, err, substraitgo.ErrInvalidExpr)
}

func TestEmptyCompoundLiteralsRemainSupported(t *testing.T) {
	itemType := &pb.Type{Kind: &pb.Type_I64_{I64: &pb.Type_I64{Nullability: pb.Type_NULLABILITY_REQUIRED}}}
	for name, input := range map[string]*pb.Expression_Literal{
		"struct":     {LiteralType: &pb.Expression_Literal_Struct_{Struct: &pb.Expression_Literal_Struct{}}},
		"typed_list": {LiteralType: &pb.Expression_Literal_EmptyList{EmptyList: &pb.Type_List{Type: itemType, Nullability: pb.Type_NULLABILITY_REQUIRED}}},
		"typed_map":  {LiteralType: &pb.Expression_Literal_EmptyMap{EmptyMap: &pb.Type_Map{Key: itemType, Value: itemType, Nullability: pb.Type_NULLABILITY_REQUIRED}}},
	} {
		t.Run(name, func(t *testing.T) {
			decoded, err := substraitproto.LiteralFromProto(input)
			require.NoError(t, err)
			require.NotNil(t, decoded)
			require.True(t, proto.Equal(input, substraitproto.LiteralToProto(decoded)))
		})
	}
}

func FuzzTypeFromProtoReturnsErrors(f *testing.F) {
	for _, input := range []*pb.Type{
		{},
		{Kind: &pb.Type_Bool{Bool: &pb.Type_Boolean{}}},
		{Kind: &pb.Type_List_{List: &pb.Type_List{}}},
		{Kind: &pb.Type_Map_{Map: &pb.Type_Map{}}},
		{Kind: &pb.Type_Func_{Func: &pb.Type_Func{}}},
		{Kind: &pb.Type_UserDefined_{UserDefined: &pb.Type_UserDefined{TypeParameters: []*pb.Type_Parameter{{}}}}},
	} {
		data, err := proto.Marshal(input)
		require.NoError(f, err)
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		var input pb.Type
		if err := proto.Unmarshal(data, &input); err != nil {
			return
		}
		decoded, err := substraitproto.TypeFromProto(&input)
		if err == nil {
			require.NotNil(t, decoded)
		}
	})
}

func FuzzLiteralFromProtoReturnsErrors(f *testing.F) {
	for _, input := range []*pb.Expression_Literal{
		{},
		{LiteralType: &pb.Expression_Literal_Boolean{}},
		{LiteralType: &pb.Expression_Literal_Null{Null: &pb.Type{}}},
		{LiteralType: &pb.Expression_Literal_List_{List: &pb.Expression_Literal_List{}}},
		{LiteralType: &pb.Expression_Literal_Map_{Map: &pb.Expression_Literal_Map{}}},
		{LiteralType: &pb.Expression_Literal_UserDefined_{UserDefined: &pb.Expression_Literal_UserDefined{}}},
		{LiteralType: &pb.Expression_Literal_IntervalCompound_{IntervalCompound: &pb.Expression_Literal_IntervalCompound{IntervalDayToSecond: &pb.Expression_Literal_IntervalDayToSecond{}}}},
	} {
		data, err := proto.Marshal(input)
		require.NoError(f, err)
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		var input pb.Expression_Literal
		if err := proto.Unmarshal(data, &input); err != nil {
			return
		}
		decoded, err := substraitproto.LiteralFromProto(&input)
		if err == nil {
			require.NotNil(t, decoded)
		}
	})
}

func TestMalformedSchemasReturnErrors(t *testing.T) {
	collection := extensions.GetDefaultCollectionWithNoError()
	registry := expr.NewEmptyExtensionRegistry(collection)
	for name, schema := range map[string]*pb.NamedStruct{
		"missing_struct": {},
		"invalid_field":  {Struct: &pb.Type_Struct{Types: []*pb.Type{{}}}},
	} {
		t.Run(name, func(t *testing.T) {
			_, err := substraitproto.NamedStructFromProto(schema)
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
			_, err = substraitproto.ExtendedFromProto(&pb.ExtendedExpression{BaseSchema: schema}, collection)
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
			read := &pb.Rel{RelType: &pb.Rel_Read{Read: &pb.ReadRel{BaseSchema: schema, ReadType: &pb.ReadRel_NamedTable_{NamedTable: &pb.ReadRel_NamedTable{Names: []string{"t"}}}}}}
			_, err = substraitproto.RelFromProto(read, registry)
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
			input := &pb.Rel{RelType: &pb.Rel_Read{Read: &pb.ReadRel{BaseSchema: &pb.NamedStruct{Struct: &pb.Type_Struct{}}, ReadType: &pb.ReadRel_NamedTable_{NamedTable: &pb.ReadRel_NamedTable{Names: []string{"source"}}}}}}
			write := &pb.Rel{RelType: &pb.Rel_Write{Write: &pb.WriteRel{Input: input, TableSchema: schema, Op: pb.WriteRel_WRITE_OP_INSERT, Output: pb.WriteRel_OUTPUT_MODE_NO_OUTPUT, WriteType: &pb.WriteRel_NamedTable{NamedTable: &pb.NamedObjectWrite{Names: []string{"target"}}}}}}
			_, err = substraitproto.RelFromProto(write, registry)
			require.ErrorIs(t, err, substraitgo.ErrInvalidType)
		})
	}
}
