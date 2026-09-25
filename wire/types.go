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

// TypeFromProto returns the appropriate Type object from a protobuf type message.
func TypeFromProto(t *proto.Type) types.Type {
	switch t := t.Kind.(type) {
	}
	panic("unimplemented type from proto")
}
