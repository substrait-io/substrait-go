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

// TypeFromProto returns the appropriate Type object from a protobuf type message.
func TypeFromProto(t *proto.Type) types.Type {
	switch t := t.Kind.(type) {
	}
	panic("unimplemented type from proto")
}
