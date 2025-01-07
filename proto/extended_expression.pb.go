// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: substrait/extended_expression.proto

package proto

import (
	extensions "github.com/substrait-io/substrait-go/v3/proto/extensions"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExpressionReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ExprType:
	//
	//	*ExpressionReference_Expression
	//	*ExpressionReference_Measure
	ExprType isExpressionReference_ExprType `protobuf_oneof:"expr_type"`
	// Field names in depth-first order
	OutputNames []string `protobuf:"bytes,3,rep,name=output_names,json=outputNames,proto3" json:"output_names,omitempty"`
}

func (x *ExpressionReference) Reset() {
	*x = ExpressionReference{}
	mi := &file_substrait_extended_expression_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpressionReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionReference) ProtoMessage() {}

func (x *ExpressionReference) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extended_expression_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionReference.ProtoReflect.Descriptor instead.
func (*ExpressionReference) Descriptor() ([]byte, []int) {
	return file_substrait_extended_expression_proto_rawDescGZIP(), []int{0}
}

func (m *ExpressionReference) GetExprType() isExpressionReference_ExprType {
	if m != nil {
		return m.ExprType
	}
	return nil
}

func (x *ExpressionReference) GetExpression() *Expression {
	if x, ok := x.GetExprType().(*ExpressionReference_Expression); ok {
		return x.Expression
	}
	return nil
}

func (x *ExpressionReference) GetMeasure() *AggregateFunction {
	if x, ok := x.GetExprType().(*ExpressionReference_Measure); ok {
		return x.Measure
	}
	return nil
}

func (x *ExpressionReference) GetOutputNames() []string {
	if x != nil {
		return x.OutputNames
	}
	return nil
}

type isExpressionReference_ExprType interface {
	isExpressionReference_ExprType()
}

type ExpressionReference_Expression struct {
	Expression *Expression `protobuf:"bytes,1,opt,name=expression,proto3,oneof"`
}

type ExpressionReference_Measure struct {
	Measure *AggregateFunction `protobuf:"bytes,2,opt,name=measure,proto3,oneof"`
}

func (*ExpressionReference_Expression) isExpressionReference_ExprType() {}

func (*ExpressionReference_Measure) isExpressionReference_ExprType() {}

// Describe a set of operations to complete.
// For compactness sake, identifiers are normalized at the plan level.
type ExtendedExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Substrait version of the expression. Optional up to 0.17.0, required for later
	// versions.
	Version *Version `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	// a list of yaml specifications this expression may depend on
	ExtensionUris []*extensions.SimpleExtensionURI `protobuf:"bytes,1,rep,name=extension_uris,json=extensionUris,proto3" json:"extension_uris,omitempty"`
	// a list of extensions this expression may depend on
	Extensions []*extensions.SimpleExtensionDeclaration `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions,omitempty"`
	// one or more expression trees with same order in plan rel
	ReferredExpr []*ExpressionReference `protobuf:"bytes,3,rep,name=referred_expr,json=referredExpr,proto3" json:"referred_expr,omitempty"`
	BaseSchema   *NamedStruct           `protobuf:"bytes,4,opt,name=base_schema,json=baseSchema,proto3" json:"base_schema,omitempty"`
	// additional extensions associated with this expression.
	AdvancedExtensions *extensions.AdvancedExtension `protobuf:"bytes,5,opt,name=advanced_extensions,json=advancedExtensions,proto3" json:"advanced_extensions,omitempty"`
	// A list of com.google.Any entities that this plan may use. Can be used to
	// warn if some embedded message types are unknown. Note that this list may
	// include message types that are ignorable (optimizations) or that are
	// unused. In many cases, a consumer may be able to work with a plan even if
	// one or more message types defined here are unknown.
	ExpectedTypeUrls []string `protobuf:"bytes,6,rep,name=expected_type_urls,json=expectedTypeUrls,proto3" json:"expected_type_urls,omitempty"`
}

func (x *ExtendedExpression) Reset() {
	*x = ExtendedExpression{}
	mi := &file_substrait_extended_expression_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtendedExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedExpression) ProtoMessage() {}

func (x *ExtendedExpression) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extended_expression_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedExpression.ProtoReflect.Descriptor instead.
func (*ExtendedExpression) Descriptor() ([]byte, []int) {
	return file_substrait_extended_expression_proto_rawDescGZIP(), []int{1}
}

func (x *ExtendedExpression) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ExtendedExpression) GetExtensionUris() []*extensions.SimpleExtensionURI {
	if x != nil {
		return x.ExtensionUris
	}
	return nil
}

func (x *ExtendedExpression) GetExtensions() []*extensions.SimpleExtensionDeclaration {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *ExtendedExpression) GetReferredExpr() []*ExpressionReference {
	if x != nil {
		return x.ReferredExpr
	}
	return nil
}

func (x *ExtendedExpression) GetBaseSchema() *NamedStruct {
	if x != nil {
		return x.BaseSchema
	}
	return nil
}

func (x *ExtendedExpression) GetAdvancedExtensions() *extensions.AdvancedExtension {
	if x != nil {
		return x.AdvancedExtensions
	}
	return nil
}

func (x *ExtendedExpression) GetExpectedTypeUrls() []string {
	if x != nil {
		return x.ExpectedTypeUrls
	}
	return nil
}

var File_substrait_extended_expression_proto protoreflect.FileDescriptor

var file_substrait_extended_expression_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x1a, 0x17, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x61, 0x6c, 0x67, 0x65,
	0x62, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a,
	0x13, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x07, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x65, 0x78,
	0x70, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xeb, 0x03, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x49, 0x52, 0x0d,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x73, 0x12, 0x50, 0x0a,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x43, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x45, 0x78, 0x70, 0x72, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x58, 0x0a,
	0x13, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x72, 0x6c, 0x73, 0x42, 0x57, 0x0a, 0x12, 0x69, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x12, 0x53, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_substrait_extended_expression_proto_rawDescOnce sync.Once
	file_substrait_extended_expression_proto_rawDescData = file_substrait_extended_expression_proto_rawDesc
)

func file_substrait_extended_expression_proto_rawDescGZIP() []byte {
	file_substrait_extended_expression_proto_rawDescOnce.Do(func() {
		file_substrait_extended_expression_proto_rawDescData = protoimpl.X.CompressGZIP(file_substrait_extended_expression_proto_rawDescData)
	})
	return file_substrait_extended_expression_proto_rawDescData
}

var file_substrait_extended_expression_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_substrait_extended_expression_proto_goTypes = []any{
	(*ExpressionReference)(nil),                   // 0: substrait.ExpressionReference
	(*ExtendedExpression)(nil),                    // 1: substrait.ExtendedExpression
	(*Expression)(nil),                            // 2: substrait.Expression
	(*AggregateFunction)(nil),                     // 3: substrait.AggregateFunction
	(*Version)(nil),                               // 4: substrait.Version
	(*extensions.SimpleExtensionURI)(nil),         // 5: substrait.extensions.SimpleExtensionURI
	(*extensions.SimpleExtensionDeclaration)(nil), // 6: substrait.extensions.SimpleExtensionDeclaration
	(*NamedStruct)(nil),                           // 7: substrait.NamedStruct
	(*extensions.AdvancedExtension)(nil),          // 8: substrait.extensions.AdvancedExtension
}
var file_substrait_extended_expression_proto_depIdxs = []int32{
	2, // 0: substrait.ExpressionReference.expression:type_name -> substrait.Expression
	3, // 1: substrait.ExpressionReference.measure:type_name -> substrait.AggregateFunction
	4, // 2: substrait.ExtendedExpression.version:type_name -> substrait.Version
	5, // 3: substrait.ExtendedExpression.extension_uris:type_name -> substrait.extensions.SimpleExtensionURI
	6, // 4: substrait.ExtendedExpression.extensions:type_name -> substrait.extensions.SimpleExtensionDeclaration
	0, // 5: substrait.ExtendedExpression.referred_expr:type_name -> substrait.ExpressionReference
	7, // 6: substrait.ExtendedExpression.base_schema:type_name -> substrait.NamedStruct
	8, // 7: substrait.ExtendedExpression.advanced_extensions:type_name -> substrait.extensions.AdvancedExtension
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_substrait_extended_expression_proto_init() }
func file_substrait_extended_expression_proto_init() {
	if File_substrait_extended_expression_proto != nil {
		return
	}
	file_substrait_algebra_proto_init()
	file_substrait_plan_proto_init()
	file_substrait_type_proto_init()
	file_substrait_extended_expression_proto_msgTypes[0].OneofWrappers = []any{
		(*ExpressionReference_Expression)(nil),
		(*ExpressionReference_Measure)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_substrait_extended_expression_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_substrait_extended_expression_proto_goTypes,
		DependencyIndexes: file_substrait_extended_expression_proto_depIdxs,
		MessageInfos:      file_substrait_extended_expression_proto_msgTypes,
	}.Build()
	File_substrait_extended_expression_proto = out.File
	file_substrait_extended_expression_proto_rawDesc = nil
	file_substrait_extended_expression_proto_goTypes = nil
	file_substrait_extended_expression_proto_depIdxs = nil
}
