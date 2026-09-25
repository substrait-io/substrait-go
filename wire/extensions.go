// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/extensions"
	extensionspb "github.com/substrait-io/substrait-protobuf/go/substraitpb/extensions"
	"google.golang.org/protobuf/types/known/anypb"
)

// advancedExtensionToProto encodes a domain AdvancedExtension as its protobuf message.
func advancedExtensionToProto(a *extensions.AdvancedExtension) *extensionspb.AdvancedExtension {
	if a == nil {
		return nil
	}
	var optimizations []*anypb.Any
	for _, o := range a.Optimizations {
		optimizations = append(optimizations, (*anypb.Any)(o))
	}
	return &extensionspb.AdvancedExtension{
		Optimization: optimizations,
		Enhancement:  (*anypb.Any)(a.Enhancement),
	}
}

// advancedExtensionFromProto converts a protobuf AdvancedExtension to the domain type.
func advancedExtensionFromProto(a *extensionspb.AdvancedExtension) *extensions.AdvancedExtension {
	if a == nil {
		return nil
	}
	var optimizations []*extensions.Optimization
	for _, o := range a.Optimization {
		optimizations = append(optimizations, (*extensions.Optimization)(o))
	}
	return &extensions.AdvancedExtension{
		Optimizations: optimizations,
		Enhancement:   (*extensions.Enhancement)(a.Enhancement),
	}
}
