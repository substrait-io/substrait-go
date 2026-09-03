// SPDX-License-Identifier: Apache-2.0

package codec

import (
	"github.com/substrait-io/substrait-go/v9/expr"
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

// PlanFromProto decodes a protobuf Plan into the domain model.
func PlanFromProto(p *substraitpb.Plan, c *extensions.Collection) (*plan.Plan, error) {
	return PlanFromProtoWithDecoder(p, c, nil)
}

// PlanFromProtoWithDecoder is like PlanFromProto but installs per-typeURL
// ExtensionRelDecoders, so matching extension rels decode into typed
// ExtensionRelDefinitions instead of UndecodedExtension.
func PlanFromProtoWithDecoder(p *substraitpb.Plan, c *extensions.Collection, decoders map[string]expr.ExtensionRelDecoder) (*plan.Plan, error) {
	extSet, err := extensions.GetExtensionSet(p, c)
	if err != nil {
		return nil, err
	}

	reg := expr.NewExtensionRegistry(extSet, c)
	reg.SetSubqueryConverter(&plan.ExpressionConverter{ExtensionRegistry: reg})
	for typeURL, dec := range decoders {
		if err := reg.SetExtensionRelDecoder(typeURL, dec); err != nil {
			return nil, err
		}
	}

	relations := make([]plan.Relation, len(p.Relations))
	for i, r := range p.Relations {
		if err := relations[i].FromProto(r, reg); err != nil {
			return nil, err
		}
	}

	var bindings []plan.DynamicParameterBinding
	if len(p.ParameterBindings) > 0 {
		bindings = make([]plan.DynamicParameterBinding, len(p.ParameterBindings))
		for i, pb := range p.ParameterBindings {
			bindings[i] = plan.DynamicParameterBinding{
				ParameterAnchor: pb.ParameterAnchor,
				Value:           expr.LiteralFromProto(pb.Value),
			}
		}
	}

	return plan.NewPlan(types.VersionFromProto(p.Version), p.ExpectedTypeUrls, p.AdvancedExtensions, relations, bindings, reg), nil
}
