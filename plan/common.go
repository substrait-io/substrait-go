// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/expr"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// DynamicParameterBinding maps a parameter anchor to a literal value
// for use with DynamicParameter expressions in a plan.
//
// When bindings are provided via PlanWithBindings, the builder validates
// that each binding's literal type matches (ignoring nullability) the
// OutputType declared on the corresponding DynamicParameter expression.
type DynamicParameterBinding struct {
	ParameterAnchor uint32
	Value           expr.Literal
}

// ValidateParameterBindings checks that every binding's literal type matches
// the OutputType of the corresponding DynamicParameter expression found in
// the relation tree. Type comparison ignores nullability so that a required
// parameter can be bound to a nullable literal and vice-versa.
//
// Returns an error for:
//   - A binding whose anchor does not correspond to any DynamicParameter in the tree.
//   - A binding whose value type does not match the parameter's declared type
//     (ignoring nullability).
func ValidateParameterBindings(root Rel, bindings []DynamicParameterBinding) error {
	if len(bindings) == 0 {
		return nil
	}

	// Collect all DynamicParameter output types keyed by anchor.
	paramTypes := make(map[uint32]types.Type)
	WalkRels(root, nil, func(e expr.Expression) {
		if dp, ok := e.(*expr.DynamicParameter); ok {
			paramTypes[dp.ParameterReference] = dp.OutputType
		}
	})

	for _, b := range bindings {
		declaredType, ok := paramTypes[b.ParameterAnchor]
		if !ok {
			return fmt.Errorf("%w: parameter binding references anchor %d, "+
				"but no DynamicParameter with that reference exists in the plan",
				substraitgo.ErrInvalidPlan, b.ParameterAnchor)
		}

		// Compare ignoring nullability.
		bindingType := b.Value.GetType().WithNullability(types.NullabilityUnspecified)
		expectedType := declaredType.WithNullability(types.NullabilityUnspecified)
		if !bindingType.Equals(expectedType) {
			return fmt.Errorf("%w: parameter binding for anchor %d has type %s, "+
				"but DynamicParameter declares type %s",
				substraitgo.ErrInvalidPlan, b.ParameterAnchor, b.Value.GetType(), declaredType)
		}
	}

	return nil
}

type (
	Hint              = proto.RelCommon_Hint
	Stats             = proto.RelCommon_Hint_Stats
	RuntimeConstraint = proto.RelCommon_Hint_RuntimeConstraint
)

// RelCommon is the common fields of all relational operators and is
// embedded in all of them.
type RelCommon struct {
	hint         *Hint
	mapping      []int32
	advExtension *extensions.AdvancedExtension
}

func (rc *RelCommon) fromProtoCommon(c *proto.RelCommon) {
	rc.hint = c.Hint
	rc.advExtension = c.AdvancedExtension

	if emit, ok := c.GetEmitKind().(*proto.RelCommon_Emit_); ok {
		rc.mapping = emit.Emit.OutputMapping
	} else {
		rc.mapping = nil
	}
}

func (rc *RelCommon) remap(initial types.RecordType) types.RecordType {
	if rc.mapping == nil {
		return initial
	}

	outTypes := make([]types.Type, len(rc.mapping))

	for i, m := range rc.mapping {
		outTypes[i] = initial.GetFieldRef(m)
	}

	return *types.NewRecordTypeFromTypes(outTypes)
}

func (rc *RelCommon) OutputMapping() []int32 {
	if rc.mapping == nil {
		return nil
	}
	// Make a copy of the output mapping to prevent accidental modification.
	mapCopy := make([]int32, len(rc.mapping))
	copy(mapCopy, rc.mapping)
	return mapCopy
}

func (rc *RelCommon) setMapping(mapping []int32) {
	rc.mapping = mapping
}

func (rc *RelCommon) GetAdvancedExtension() *extensions.AdvancedExtension {
	return rc.advExtension
}

func (rc *RelCommon) SetAdvancedExtension(advExtension *extensions.AdvancedExtension) *extensions.AdvancedExtension {
	existing := rc.advExtension
	rc.advExtension = advExtension
	return existing
}

func (rc *RelCommon) Hint() *Hint {
	return rc.hint
}

func (rc *RelCommon) toProto() *proto.RelCommon {
	ret := &proto.RelCommon{
		Hint:              rc.hint,
		AdvancedExtension: rc.advExtension,
	}

	if rc.mapping == nil {
		ret.EmitKind = &proto.RelCommon_Direct_{
			Direct: &proto.RelCommon_Direct{},
		}
	} else {
		ret.EmitKind = &proto.RelCommon_Emit_{
			Emit: &proto.RelCommon_Emit{OutputMapping: rc.mapping},
		}
	}
	return ret
}
