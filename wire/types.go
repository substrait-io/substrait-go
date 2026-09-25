// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// TypeToProto constructs the protobuf message for the given type.
func TypeToProto(t types.Type) *proto.Type {
	switch t := t.(type) {
	}
	panic("unimplemented type")
}

// TypeFromProto returns the appropriate Type object from a protobuf type message.
func TypeFromProto(t *proto.Type) types.Type {
	switch t := t.Kind.(type) {
	}
	panic("unimplemented type from proto")
}
