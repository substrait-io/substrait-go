package protoenc

import (
	"fmt"

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
			params[i] = TypeParamToProto(p, extSet)
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
	case *types.FuncType:
		params := make([]*proto.Type, len(t.ParameterTypes))
		for i, param := range t.ParameterTypes {
			params[i] = TypeToProto(param, extSet)
		}
		return &proto.Type{Kind: &proto.Type_Func_{Func: &proto.Type_Func{
			Nullability:    t.Nullability,
			ParameterTypes: params,
			ReturnType:     TypeToProto(t.ReturnType, extSet),
		}}}
	default:
		return types.TypeToProto(t)
	}
}

func TypeParamToProto(p types.TypeParam, extSet extensions.Set) *proto.Type_Parameter {
	if data, ok := p.(*types.DataTypeParameter); ok {
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_DataType{
			DataType: TypeToProto(data.Type, extSet),
		}}
	}
	return p.ToProto()
}

func TypeFromProto(t *proto.Type, extSet extensions.Set) (types.Type, error) {
	switch t := t.Kind.(type) {
	case *proto.Type_UserDefined_:
		id, ok := extSet.DecodeType(t.UserDefined.TypeReference)
		if !ok {
			return nil, fmt.Errorf("user-defined type anchor %d is not registered", t.UserDefined.TypeReference)
		}
		params, err := typeParamsFromProto(t.UserDefined.TypeParameters, extSet)
		if err != nil {
			return nil, err
		}
		return &types.UserDefinedType{
			Nullability:      t.UserDefined.Nullability,
			TypeVariationRef: t.UserDefined.TypeVariationReference,
			ID:               id,
			TypeParameters:   params,
		}, nil
	case *proto.Type_Struct_:
		children := make([]types.Type, len(t.Struct.Types))
		for i, child := range t.Struct.Types {
			var err error
			children[i], err = TypeFromProto(child, extSet)
			if err != nil {
				return nil, err
			}
		}
		return &types.StructType{Nullability: t.Struct.Nullability, TypeVariationRef: t.Struct.TypeVariationReference, Types: children}, nil
	case *proto.Type_List_:
		child, err := TypeFromProto(t.List.Type, extSet)
		if err != nil {
			return nil, err
		}
		return &types.ListType{Nullability: t.List.Nullability, TypeVariationRef: t.List.TypeVariationReference, Type: child}, nil
	case *proto.Type_Map_:
		key, err := TypeFromProto(t.Map.Key, extSet)
		if err != nil {
			return nil, err
		}
		value, err := TypeFromProto(t.Map.Value, extSet)
		if err != nil {
			return nil, err
		}
		return &types.MapType{Nullability: t.Map.Nullability, TypeVariationRef: t.Map.TypeVariationReference, Key: key, Value: value}, nil
	case *proto.Type_Func_:
		params := make([]types.Type, len(t.Func.ParameterTypes))
		for i, param := range t.Func.ParameterTypes {
			var err error
			params[i], err = TypeFromProto(param, extSet)
			if err != nil {
				return nil, err
			}
		}
		ret, err := TypeFromProto(t.Func.ReturnType, extSet)
		if err != nil {
			return nil, err
		}
		return &types.FuncType{Nullability: t.Func.Nullability, ParameterTypes: params, ReturnType: ret}, nil
	default:
		return types.TypeFromProto(&proto.Type{Kind: t}), nil
	}
}

func TypeParamFromProto(p *proto.Type_Parameter, extSet extensions.Set) (types.TypeParam, error) {
	if data := p.GetDataType(); data != nil {
		t, err := TypeFromProto(data, extSet)
		if err != nil {
			return nil, err
		}
		return &types.DataTypeParameter{Type: t}, nil
	}
	return types.TypeParamFromProto(p), nil
}

func typeParamsFromProto(params []*proto.Type_Parameter, extSet extensions.Set) ([]types.TypeParam, error) {
	out := make([]types.TypeParam, len(params))
	for i, param := range params {
		var err error
		out[i], err = TypeParamFromProto(param, extSet)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
