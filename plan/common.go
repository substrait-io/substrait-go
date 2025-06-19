// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v4/extensions"
	"github.com/substrait-io/substrait-go/v4/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

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

func (rc *RelCommon) SetAdvancedExtension(advExtension *extensions.AdvancedExtension) error {
	if rc.advExtension != nil {
		return fmt.Errorf("AdvancedExtension is already set")
	}
	rc.advExtension = advExtension
	return nil
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
