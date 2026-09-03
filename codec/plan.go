// SPDX-License-Identifier: Apache-2.0

package codec

import (
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	substraitpb "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// PlanToProto is plan.Plan.ToProto moved into codec: identical logic, reading the
// plan's state through its public accessors instead of unexported fields.
func PlanToProto(p *plan.Plan) (*substraitpb.Plan, error) {
	reg := p.ExtensionRegistry()
	urns, decls := reg.ExtensionsToProto()

	rels := p.Relations()
	relations := make([]*substraitpb.PlanRel, len(rels))
	for i, r := range rels {
		relations[i] = r.ToProto()
	}

	var bindings []*substraitpb.DynamicParameterBinding
	if binds := p.ParameterBindings(); len(binds) > 0 {
		bindings = make([]*substraitpb.DynamicParameterBinding, len(binds))
		for i, b := range binds {
			bindings[i] = &substraitpb.DynamicParameterBinding{
				ParameterAnchor: b.ParameterAnchor,
				Value:           b.Value.ToProtoLiteral(),
			}
		}
	}

	ae, _ := p.AdvancedExtension().(*extensions.AdvancedExtension)
	return &substraitpb.Plan{
		Version:            types.VersionToProto(p.Version()),
		ExpectedTypeUrls:   p.ExpectedTypeURLs(),
		AdvancedExtensions: ae,
		Relations:          relations,
		Extensions:         decls,
		ExtensionUrns:      urns,
		ParameterBindings:  bindings,
	}, nil
}
