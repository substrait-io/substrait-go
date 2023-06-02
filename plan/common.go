// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/proto"
)

type (
	Hint              = proto.RelCommon_Hint
	Stats             = proto.RelCommon_Hint_Stats
	RuntimeConstraint = proto.RelCommon_Hint_RuntimeConstraint
)

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

func (rc *RelCommon) OutputMapping() []int32 { return rc.mapping }

func (rc *RelCommon) GetAdvancedExtension() *extensions.AdvancedExtension {
	return rc.advExtension
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
