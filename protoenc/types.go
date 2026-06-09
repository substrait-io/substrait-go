package protoenc

import (
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// TypeToProto serializes a type using extSet to assign plan-local anchors for
// user-defined types. Non-UDT types continue to use the regular types encoder.
func TypeToProto(t types.Type, extSet extensions.Set) *proto.Type {
	switch t := t.(type) {
	case *types.UserDefinedType:
		params := make([]*proto.Type_Parameter, len(t.TypeParameters))
		for i, p := range t.TypeParameters {
			params[i] = typeParamToProto(p, extSet)
		}
		return &proto.Type{Kind: &proto.Type_UserDefined_{
			UserDefined: &proto.Type_UserDefined{
				Nullability:            t.Nullability,
				TypeVariationReference: t.TypeVariationRef,
				TypeReference:          extSet.GetTypeAnchor(t.ID),
				TypeParameters:         params,
			},
		}}
	case *types.StructType:
		children := make([]*proto.Type, len(t.Types))
		for i, child := range t.Types {
			children[i] = TypeToProto(child, extSet)
		}
		return &proto.Type{Kind: &proto.Type_Struct_{Struct: &proto.Type_Struct{
			Nullability:            t.Nullability,
			TypeVariationReference: t.TypeVariationRef,
			Types:                  children,
		}}}
	case *types.ListType:
		return &proto.Type{Kind: &proto.Type_List_{List: &proto.Type_List{
			Nullability:            t.Nullability,
			TypeVariationReference: t.TypeVariationRef,
			Type:                   TypeToProto(t.Type, extSet),
		}}}
	case *types.MapType:
		return &proto.Type{Kind: &proto.Type_Map_{Map: &proto.Type_Map{
			Nullability:            t.Nullability,
			TypeVariationReference: t.TypeVariationRef,
			Key:                    TypeToProto(t.Key, extSet),
			Value:                  TypeToProto(t.Value, extSet),
		}}}
	default:
		return types.TypeToProto(t)
	}
}

func typeParamToProto(p types.TypeParam, extSet extensions.Set) *proto.Type_Parameter {
	if data, ok := p.(*types.DataTypeParameter); ok {
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_DataType{
			DataType: TypeToProto(data.Type, extSet),
		}}
	}
	return p.ToProto()
}
