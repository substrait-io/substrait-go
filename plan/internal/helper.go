package internal

import (
	"github.com/substrait-io/substrait-go/v7/expr"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func VirtualTableExpressionFromProto(s *proto.Expression_Nested_Struct, reg expr.ExtensionRegistry) (expr.VirtualTableExpressionValue, error) {
	fields := make(expr.VirtualTableExpressionValue, len(s.Fields))
	for i, f := range s.Fields {
		val, err := expr.ExprFromProto(f, nil, reg)
		if err != nil {
			return nil, err
		}
		fields[i] = val
	}
	return fields, nil
}

func VirtualTableExprFromLiteralProto(s *proto.Expression_Literal_Struct) expr.VirtualTableExpressionValue {
	fields := make(expr.VirtualTableExpressionValue, len(s.Fields))
	for i, f := range s.Fields {
		fields[i] = expr.LiteralFromProto(f)
	}
	return fields
}
