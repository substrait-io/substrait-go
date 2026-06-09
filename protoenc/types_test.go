package protoenc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v8/extensions"
	"github.com/substrait-io/substrait-go/v8/protoenc"
	"github.com/substrait-io/substrait-go/v8/types"
)

func TestTypeToProtoAssignsUserDefinedTypeAnchors(t *testing.T) {
	extSet := extensions.NewSet()
	pointID := extensions.TypeID{URN: "extension:test:types", Name: "point"}

	protoType := protoenc.TypeToProto(&types.UserDefinedType{
		Nullability:      types.NullabilityNullable,
		TypeVariationRef: 9,
		ID:               pointID,
		TypeParameters: []types.TypeParam{
			&types.DataTypeParameter{Type: &types.Int32Type{Nullability: types.NullabilityRequired}},
			types.IntegerParameter(4),
		},
	}, extSet)

	udt := protoType.GetUserDefined()
	require.NotNil(t, udt)
	require.Equal(t, types.NullabilityNullable, udt.Nullability)
	require.Equal(t, uint32(9), udt.TypeVariationReference)
	require.Equal(t, extSet.GetTypeAnchor(pointID), udt.TypeReference)
	require.Len(t, udt.TypeParameters, 2)
	require.NotNil(t, udt.TypeParameters[0].GetDataType().GetI32())
	require.Equal(t, int64(4), udt.TypeParameters[1].GetInteger())
}

func TestTypeToProtoHandlesNestedUserDefinedTypes(t *testing.T) {
	extSet := extensions.NewSet()
	pointID := extensions.TypeID{URN: "extension:test:types", Name: "point"}
	lineID := extensions.TypeID{URN: "extension:test:types", Name: "line"}

	protoType := protoenc.TypeToProto(&types.StructType{
		Nullability: types.NullabilityRequired,
		Types: []types.Type{
			&types.ListType{
				Nullability: types.NullabilityNullable,
				Type: &types.UserDefinedType{
					Nullability: types.NullabilityRequired,
					ID:          pointID,
				},
			},
			&types.MapType{
				Nullability: types.NullabilityRequired,
				Key:         &types.StringType{Nullability: types.NullabilityRequired},
				Value: &types.UserDefinedType{
					Nullability: types.NullabilityNullable,
					ID:          lineID,
				},
			},
		},
	}, extSet)

	fields := protoType.GetStruct().Types
	require.Len(t, fields, 2)
	require.Equal(t, extSet.GetTypeAnchor(pointID), fields[0].GetList().Type.GetUserDefined().TypeReference)
	require.Equal(t, extSet.GetTypeAnchor(lineID), fields[1].GetMap().Value.GetUserDefined().TypeReference)
	require.NotNil(t, fields[1].GetMap().Key.GetString_())
}

func TestTypeToProtoDelegatesNonUserDefinedTypes(t *testing.T) {
	protoType := protoenc.TypeToProto(&types.BooleanType{Nullability: types.NullabilityRequired}, extensions.NewSet())
	require.NotNil(t, protoType.GetBool())
	require.Equal(t, types.NullabilityRequired, protoType.GetBool().Nullability)
}

func TestTypeFromProtoResolvesUserDefinedTypeAnchors(t *testing.T) {
	extSet := extensions.NewSet()
	pointID := extensions.TypeID{URN: "extension:test:types", Name: "point"}
	typeAnchor := extSet.GetTypeAnchor(pointID)

	protoType := protoenc.TypeToProto(&types.FuncType{
		Nullability: types.NullabilityRequired,
		ParameterTypes: []types.Type{
			&types.ListType{Type: &types.UserDefinedType{ID: pointID, Nullability: types.NullabilityNullable}},
		},
		ReturnType: &types.UserDefinedType{
			Nullability: types.NullabilityRequired,
			ID:          pointID,
			TypeParameters: []types.TypeParam{
				&types.DataTypeParameter{Type: &types.Int32Type{Nullability: types.NullabilityRequired}},
			},
		},
	}, extSet)

	decoded, err := protoenc.TypeFromProto(protoType, extSet)
	require.NoError(t, err)

	fn := decoded.(*types.FuncType)
	paramUDT := fn.ParameterTypes[0].(*types.ListType).Type.(*types.UserDefinedType)
	require.Equal(t, pointID, paramUDT.ID)

	returnUDT := fn.ReturnType.(*types.UserDefinedType)
	require.Equal(t, pointID, returnUDT.ID)
	require.Equal(t, typeAnchor, protoType.GetFunc().ReturnType.GetUserDefined().TypeReference)
	require.Len(t, returnUDT.TypeParameters, 1)
}

func TestTypeFromProtoRejectsUnknownUserDefinedTypeAnchor(t *testing.T) {
	protoType := protoenc.TypeToProto(&types.UserDefinedType{ID: extensions.TypeID{URN: "extension:test:types", Name: "point"}}, extensions.NewSet())

	_, err := protoenc.TypeFromProto(protoType, extensions.NewSet())
	require.ErrorContains(t, err, "user-defined type anchor")
}
