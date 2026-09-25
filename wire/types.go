// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// VersionToProto encodes a version as its protobuf message.
func VersionToProto(v types.Version) *proto.Version {
	return &proto.Version{
		MajorNumber: v.MajorNumber,
		MinorNumber: v.MinorNumber,
		PatchNumber: v.PatchNumber,
		GitHash:     v.GitHash,
		Producer:    v.Producer,
	}
}

// NamedStructToProto encodes a named struct as its protobuf message.
func NamedStructToProto(n types.NamedStruct) *proto.NamedStruct {
	return &proto.NamedStruct{
		Names:  n.Names,
		Struct: structTypeToProto(&n.Struct).GetStruct(),
	}
}

// TypeToProto constructs the protobuf message for the given type.
func TypeToProto(t types.Type) *proto.Type {
	switch t := t.(type) {
	}
	panic("unimplemented type")
}

// TypeParamToProto encodes a user-defined-type parameter.
func TypeParamToProto(p types.TypeParam) *proto.Type_Parameter {
	switch p := p.(type) {
	case types.NullParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Null{}}
	case *types.DataTypeParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_DataType{
			DataType: TypeToProto(p.Type)}}
	case types.BooleanParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Boolean{
			Boolean: bool(p)}}
	case types.IntegerParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Integer{
			Integer: int64(p)}}
	case types.EnumParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Enum{
			Enum: string(p)}}
	case types.StringParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_String_{
			String_: string(p)}}
	}
	panic("unimplemented type parameter")
}

// VersionFromProto decodes a version from its protobuf message.
func VersionFromProto(v *proto.Version) types.Version {
	if v == nil {
		return types.Version{Producer: "UNSET"}
	}
	return types.Version{
		MajorNumber: v.MajorNumber,
		MinorNumber: v.MinorNumber,
		PatchNumber: v.PatchNumber,
		GitHash:     v.GitHash,
		Producer:    v.Producer,
	}
}

// NamedStructFromProto decodes a named struct from its protobuf message.
func NamedStructFromProto(n *proto.NamedStruct) types.NamedStruct {
	if n == nil {
		return types.NamedStruct{}
	}

	fields := make([]types.Type, len(n.Struct.Types))
	for i, f := range n.Struct.Types {
		fields[i] = TypeFromProto(f)
	}

	return types.NamedStruct{
		Names: n.Names,
		Struct: types.StructType{
			Nullability:      types.Nullability(n.Struct.Nullability),
			TypeVariationRef: n.Struct.TypeVariationReference,
			Types:            fields,
		},
	}
}

// TypeParamFromProto decodes a protobuf Type_Parameter message into a TypeParam.
func TypeParamFromProto(p *proto.Type_Parameter) types.TypeParam {
	switch p := p.Parameter.(type) {
	case *proto.Type_Parameter_Null:
		return types.NullParameter{}
	case *proto.Type_Parameter_Boolean:
		return types.BooleanParameter(p.Boolean)
	case *proto.Type_Parameter_DataType:
		return &types.DataTypeParameter{Type: TypeFromProto(p.DataType)}
	case *proto.Type_Parameter_Integer:
		return types.IntegerParameter(p.Integer)
	case *proto.Type_Parameter_Enum:
		return types.EnumParameter(p.Enum)
	case *proto.Type_Parameter_String_:
		return types.StringParameter(p.String_)
	}
	return nil
}

// TypeFromProto returns the appropriate Type object from a protobuf type message.
func TypeFromProto(t *proto.Type) types.Type {
	switch t := t.Kind.(type) {
	}
	panic("unimplemented type from proto")
}
