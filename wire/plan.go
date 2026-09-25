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
