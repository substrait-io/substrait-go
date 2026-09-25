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
	"google.golang.org/protobuf/types/known/anypb"
)

// RelToProto encodes a relation as its protobuf message.
func RelToProto(rel plan.Rel) *proto.Rel {
	switch r := rel.(type) {
	case *plan.NamedTableReadRel:
		return namedTableReadRelToProto(r)
	case *plan.VirtualTableReadRel:
		return virtualTableReadRelToProto(r)
	default:
		panic(fmt.Sprintf("wire: unhandled relation %T", rel))
	}
}

type readRelReader interface {
	BaseSchema() types.NamedStruct
	Filter() expr.Expression
	BestEffortFilter() expr.Expression
	Projection() *expr.MaskExpression
}

func baseReadRelToProto(rc *plan.RelCommon, advExt *extensions.AdvancedExtension, r readRelReader) *proto.ReadRel {
	out := &proto.ReadRel{
		Common:            relCommonToProto(rc),
		BaseSchema:        NamedStructToProto(r.BaseSchema()),
		AdvancedExtension: advancedExtensionToProto(advExt),
	}
	if f := r.Filter(); f != nil {
		out.Filter = ExprToProto(f)
	}
	if f := r.BestEffortFilter(); f != nil {
		out.BestEffortFilter = ExprToProto(f)
	}
	if p := r.Projection(); p != nil {
		out.Projection = MaskExpressionToProto(p)
	}
	return out
}

func namedTableReadRelToProto(n *plan.NamedTableReadRel) *proto.Rel {
	readRel := baseReadRelToProto(&n.RelCommon, n.GetAdvancedExtension(), n)
	readRel.ReadType = &proto.ReadRel_NamedTable_{
		NamedTable: &proto.ReadRel_NamedTable{
			Names:             n.Names(),
			AdvancedExtension: advancedExtensionToProto(n.NamedTableAdvancedExtension()),
		},
	}
	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

func virtualTableReadRelToProto(v *plan.VirtualTableReadRel) *proto.Rel {
	readRel := baseReadRelToProto(&v.RelCommon, v.GetAdvancedExtension(), v)
	values := make([]*proto.Expression_Nested_Struct, len(v.Values()))
	for i, val := range v.Values() {
		values[i] = VirtualTableExpressionValueToProto(val)
	}
	readRel.ReadType = &proto.ReadRel_VirtualTable_{
		VirtualTable: &proto.ReadRel_VirtualTable{Expressions: values},
	}
	return &proto.Rel{RelType: &proto.Rel_Read{Read: readRel}}
}

// relCommonFromProto decodes the common fields shared by every relation.
func relCommonFromProto(c *proto.RelCommon) plan.RelCommon {
	if c == nil {
		return plan.NewRelCommon(nil, nil, nil)
	}
	var mapping []int32
	if emit, ok := c.GetEmitKind().(*proto.RelCommon_Emit_); ok {
		mapping = emit.Emit.OutputMapping
	}
	return plan.NewRelCommon(hintFromProto(c.Hint), mapping, advancedExtensionFromProto(c.AdvancedExtension))
}

// decodedReadRelBase carries the fields common to all read relations while a
// specific read relation is being decoded.
type decodedReadRelBase struct {
	common           plan.RelCommon
	baseSchema       types.NamedStruct
	filter           expr.Expression
	bestEffortFilter expr.Expression
	projection       *expr.MaskExpression
	advExtension     *extensions.AdvancedExtension
}

func readRelBaseFromProto(rel *proto.ReadRel, reg expr.ExtensionRegistry) (decodedReadRelBase, error) {
	var b decodedReadRelBase
	if rel.Common != nil {
		b.common = relCommonFromProto(rel.Common)
	}

	b.baseSchema = NamedStructFromProto(rel.BaseSchema)
	var err error
	if rel.Filter != nil {
		b.filter, err = ExprFromProto(rel.Filter, types.NewRecordTypeFromStruct(b.baseSchema.Struct), reg)
		if err != nil {
			return b, err
		}
	}

	if rel.BestEffortFilter != nil {
		b.bestEffortFilter, err = ExprFromProto(rel.BestEffortFilter, types.NewRecordTypeFromStruct(b.baseSchema.Struct), reg)
		if err != nil {
			return b, err
		}
	}

	if rel.Projection != nil {
		b.projection = MaskExpressionFromProto(rel.Projection)
	}

	b.advExtension = advancedExtensionFromProto(rel.AdvancedExtension)
	return b, nil
}

// fileOrFilesFromProto decodes a single local-file item.

// virtualTableExpressionFromProto decodes an expression-valued virtual table row.
func virtualTableExpressionFromProto(s *proto.Expression_Nested_Struct, reg expr.ExtensionRegistry) (expr.VirtualTableExpressionValue, error) {
	fields := make(expr.VirtualTableExpressionValue, len(s.Fields))
	for i, f := range s.Fields {
		val, err := ExprFromProto(f, nil, reg)
		if err != nil {
			return nil, err
		}
		fields[i] = val
	}
	return fields, nil
}

// virtualTableExprFromLiteralProto decodes a literal-valued virtual table row.
func virtualTableExprFromLiteralProto(s *proto.Expression_Literal_Struct) expr.VirtualTableExpressionValue {
	fields := make(expr.VirtualTableExpressionValue, len(s.Fields))
	for i, f := range s.Fields {
		fields[i] = LiteralFromProto(f)
	}
	return fields
}

// RelFromProto decodes a relation and all of its inputs from protobuf.
func RelFromProto(rel *proto.Rel, reg expr.ExtensionRegistry) (plan.Rel, error) {
	switch rel := rel.RelType.(type) {
	case *proto.Rel_Read:
		// Decode the read type first, then the shared base, so a malformed read
		// type is reported before a base field (Filter/BestEffortFilter) that
		// also fails to decode. build assembles the relation once the base is ready.
		var build func(base decodedReadRelBase) plan.Rel
		switch readType := rel.Read.ReadType.(type) {
		case *proto.ReadRel_NamedTable_:
			names := readType.NamedTable.Names
			advExtension := advancedExtensionFromProto(readType.NamedTable.AdvancedExtension)
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewNamedTableReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), names, advExtension)
			}
		case *proto.ReadRel_VirtualTable_:
			if len(readType.VirtualTable.Values) > 0 && len(readType.VirtualTable.Expressions) > 0 {
				return nil, fmt.Errorf("VirtualTable cannot declare both Values and Expressions")
			}
			var values []expr.VirtualTableExpressionValue
			for _, v := range readType.VirtualTable.Values {
				values = append(values, virtualTableExprFromLiteralProto(v))
			}
			for _, v := range readType.VirtualTable.Expressions {
				row, err := virtualTableExpressionFromProto(v, reg)
				if err != nil {
					return nil, err
				}
				values = append(values, row)
			}
			build = func(b decodedReadRelBase) plan.Rel {
				return plan.NewVirtualTableReadRel(plan.NewBaseReadRel(b.common, b.baseSchema, b.filter, b.bestEffortFilter, b.projection, b.advExtension), values)
			}
		default:
			return nil, fmt.Errorf("%w: unknown ReadRel type", substraitgo.ErrInvalidRel)
		}

		base, err := readRelBaseFromProto(rel.Read, reg)
		if err != nil {
			return nil, err
		}
		return build(base), nil
	case nil:
		return nil, fmt.Errorf("%w: got nil", substraitgo.ErrInvalidRel)
	}

	return nil, substraitgo.ErrNotImplemented
}
