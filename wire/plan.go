// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// PlanToProto encodes a plan as its protobuf message.
func PlanToProto(p *plan.Plan) (*proto.Plan, error) {
	reg := p.ExtensionRegistry()
	urns, decls := ExtensionsToProto(reg)

	rels := p.Relations()
	relations := make([]*proto.PlanRel, len(rels))
	for i := range rels {
		relations[i] = relationToProto(&rels[i])
	}

	var bindings []*proto.DynamicParameterBinding
	if bs := p.ParameterBindings(); len(bs) > 0 {
		bindings = make([]*proto.DynamicParameterBinding, len(bs))
		for i, b := range bs {
			bindings[i] = &proto.DynamicParameterBinding{
				ParameterAnchor: b.ParameterAnchor,
				Value:           LiteralToProto(b.Value),
			}
		}
	}

	return &proto.Plan{
		Version:            VersionToProto(p.Version()),
		ExpectedTypeUrls:   p.ExpectedTypeURLs(),
		AdvancedExtensions: advancedExtensionToProto(p.GetAdvancedExtension()),
		Relations:          relations,
		Extensions:         decls,
		ExtensionUrns:      urns,
		ParameterBindings:  bindings,
	}, nil
}

func relationToProto(r *plan.Relation) *proto.PlanRel {
	if r.IsRoot() {
		return rootToProto(r.Root())
	}
	return relToPlanRelProto(r.Rel())
}

func rootToProto(root *plan.Root) *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Root{
			Root: &proto.RelRoot{
				Input: RelToProto(root.Input()),
				Names: root.Names(),
			},
		},
	}
}

func relToPlanRelProto(rel plan.Rel) *proto.PlanRel {
	return &proto.PlanRel{
		RelType: &proto.PlanRel_Rel{Rel: RelToProto(rel)},
	}
}

func isRecordTypeSupported(rel plan.Rel) bool {
	switch r := rel.(type) {
	case *plan.ExtensionSingleRel:
		_, undecoded := r.Definition().(*plan.UndecodedExtension)
		return !undecoded
	case *plan.ExtensionLeafRel:
		_, undecoded := r.Definition().(*plan.UndecodedExtension)
		return !undecoded
	case *plan.ExtensionMultiRel:
		_, undecoded := r.Definition().(*plan.UndecodedExtension)
		return !undecoded
	case *plan.NamedTableWriteRel:
		return false // TODO(#210): panics when outputMode is unspecified
	}
	return true
}

func validateRootNamesForSchema(recordType types.RecordType, names []string) error {
	expected := recordType.AsStructType().DepthFirstNameCount()
	if len(names) != expected {
		return fmt.Errorf("%w: root relation has %d output name(s) but the output schema requires %d",
			substraitgo.ErrInvalidRel, len(names), expected)
	}
	return nil
}

// RelationFromProto decodes a top-level plan relation (a root or a plain relation).
func RelationFromProto(p *proto.PlanRel, reg expr.ExtensionRegistry) (plan.Relation, error) {
	switch rel := p.RelType.(type) {
	case *proto.PlanRel_Rel:
		input, err := RelFromProto(rel.Rel, reg)
		if err != nil {
			return plan.Relation{}, err
		}
		return plan.NewRelation(nil, input), nil
	case *proto.PlanRel_Root:
		input, err := RelFromProto(rel.Root.Input, reg)
		if err != nil {
			return plan.Relation{}, err
		}

		names := rel.Root.Names
		if isRecordTypeSupported(input) {
			if err := validateRootNamesForSchema(input.RecordType(), names); err != nil {
				return plan.Relation{}, err
			}
		}

		return plan.NewRelation(plan.NewRoot(input, names), nil), nil
	}

	return plan.Relation{}, fmt.Errorf("%w: no rel or root set", substraitgo.ErrInvalidRel)
}

// PlanFromProto decodes a plan from its protobuf form using the extension collection c.
func PlanFromProto(p *proto.Plan, c *extensions.Collection) (*plan.Plan, error) {
	return PlanFromProtoWithDecoder(p, c, nil)
}

// PlanFromProtoWithDecoder is like PlanFromProto but registers per-typeURL
// ExtensionRelDecoders on the registry before parsing relations, allowing
// extension rels to be decoded into typed ExtensionRelDefinitions rather than
// UndecodedExtension.
func PlanFromProtoWithDecoder(p *proto.Plan, c *extensions.Collection, decoders map[string]expr.ExtensionRelDecoder) (*plan.Plan, error) {
	extSet, err := GetExtensionSet(p, c)
	if err != nil {
		return nil, err
	}
	version := VersionFromProto(p.Version)

	reg := expr.NewExtensionRegistry(extSet, c)
	for typeURL, dec := range decoders {
		if err := reg.SetExtensionRelDecoder(typeURL, dec); err != nil {
			return nil, err
		}
	}

	relations := make([]plan.Relation, len(p.Relations))
	for i, r := range p.Relations {
		relations[i], err = RelationFromProto(r, reg)
		if err != nil {
			return nil, err
		}
	}

	var parameterBindings []plan.DynamicParameterBinding
	if len(p.ParameterBindings) > 0 {
		parameterBindings = make([]plan.DynamicParameterBinding, len(p.ParameterBindings))
		for i, pb := range p.ParameterBindings {
			parameterBindings[i] = plan.DynamicParameterBinding{
				ParameterAnchor: pb.ParameterAnchor,
				Value:           LiteralFromProto(pb.Value),
			}
		}
	}

	return plan.NewPlan(version, extSet, advancedExtensionFromProto(p.AdvancedExtensions), p.ExpectedTypeUrls, relations, parameterBindings, reg), nil
}
