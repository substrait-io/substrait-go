// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// VersionToProto encodes a version as its protobuf message.
func VersionToProto(v types.Version) *proto.Version {
	return &proto.Version{
		MajorNumber: v.MajorNumber,
		MinorNumber: v.MinorNumber,
		PatchNumber: v.PatchNumber,
		GitHash:     v.GitHash,
		Producer:    v.Producer,
	}
}

// NamedStructToProto encodes a named struct as its protobuf message.
func NamedStructToProto(n types.NamedStruct) *proto.NamedStruct {
	return &proto.NamedStruct{
		Names:  n.Names,
		Struct: structTypeToProto(&n.Struct).GetStruct(),
	}
}

// TypeToProto constructs the protobuf message for the given type.
func TypeToProto(t types.Type) *proto.Type {
	switch t := t.(type) {
	case *types.BooleanType:
		return &proto.Type{Kind: &proto.Type_Bool{
			Bool: &proto.Type_Boolean{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.Int8Type:
		return &proto.Type{Kind: &proto.Type_I8_{
			I8: &proto.Type_I8{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.Int16Type:
		return &proto.Type{Kind: &proto.Type_I16_{
			I16: &proto.Type_I16{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.Int32Type:
		return &proto.Type{Kind: &proto.Type_I32_{
			I32: &proto.Type_I32{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.Int64Type:
		return &proto.Type{Kind: &proto.Type_I64_{
			I64: &proto.Type_I64{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.Float32Type:
		return &proto.Type{Kind: &proto.Type_Fp32{
			Fp32: &proto.Type_FP32{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.Float64Type:
		return &proto.Type{Kind: &proto.Type_Fp64{
			Fp64: &proto.Type_FP64{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.StringType:
		return &proto.Type{Kind: &proto.Type_String_{
			String_: &proto.Type_String{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.BinaryType:
		return &proto.Type{Kind: &proto.Type_Binary_{
			Binary: &proto.Type_Binary{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.DateType:
		return &proto.Type{Kind: &proto.Type_Date_{
			Date: &proto.Type_Date{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.TimeType:
		return &proto.Type{Kind: &proto.Type_Time_{
			Time: &proto.Type_Time{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.TimestampTzType:
		return &proto.Type{Kind: &proto.Type_TimestampTz{
			TimestampTz: &proto.Type_TimestampTZ{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.TimestampType:
		return &proto.Type{Kind: &proto.Type_Timestamp_{
			Timestamp: &proto.Type_Timestamp{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.IntervalYearType:
		return &proto.Type{Kind: &proto.Type_IntervalYear_{
			IntervalYear: &proto.Type_IntervalYear{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case types.IntervalYearToMonthType:
		return &proto.Type{Kind: &proto.Type_IntervalYear_{
			IntervalYear: &proto.Type_IntervalYear{
				Nullability:            proto.Type_Nullability(t.GetNullability()),
				TypeVariationReference: t.GetTypeVariationReference()}}}
	case *types.UUIDType:
		return &proto.Type{Kind: &proto.Type_Uuid{
			Uuid: &proto.Type_UUID{
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.FixedCharType:
		return &proto.Type{Kind: &proto.Type_FixedChar_{
			FixedChar: &proto.Type_FixedChar{
				Length:                 t.Length,
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.VarCharType:
		return &proto.Type{Kind: &proto.Type_Varchar{
			Varchar: &proto.Type_VarChar{
				Length:                 t.Length,
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.FixedBinaryType:
		return &proto.Type{Kind: &proto.Type_FixedBinary_{
			FixedBinary: &proto.Type_FixedBinary{
				Length:                 t.Length,
				Nullability:            proto.Type_Nullability(t.Nullability),
				TypeVariationReference: t.TypeVariationRef}}}
	case *types.DecimalType:
		return decimalTypeToProto(t)
	case *types.StructType:
		return structTypeToProto(t)
	case *types.FuncType:
		return funcTypeToProto(t)
	case *types.ListType:
		return listTypeToProto(t)
	case *types.MapType:
		return mapTypeToProto(t)
	case *types.UserDefinedType:
		return userDefinedTypeToProto(t)
	}
	panic("unimplemented type")
}

func decimalTypeToProto(s *types.DecimalType) *proto.Type {
	return &proto.Type{Kind: &proto.Type_Decimal_{
		Decimal: &proto.Type_Decimal{
			Scale: s.Scale, Precision: s.Precision,
			Nullability:            proto.Type_Nullability(s.Nullability),
			TypeVariationReference: s.TypeVariationRef}}}
}

func structTypeToProto(t *types.StructType) *proto.Type {
	children := make([]*proto.Type, len(t.Types))
	for i, c := range t.Types {
		children[i] = TypeToProto(c)
	}

	return &proto.Type{Kind: &proto.Type_Struct_{
		Struct: &proto.Type_Struct{Types: children,
			TypeVariationReference: t.TypeVariationRef,
			Nullability:            proto.Type_Nullability(t.Nullability)}}}
}

func funcTypeToProto(f *types.FuncType) *proto.Type {
	params := make([]*proto.Type, len(f.ParameterTypes))
	for i, p := range f.ParameterTypes {
		params[i] = TypeToProto(p)
	}

	return &proto.Type{Kind: &proto.Type_Func_{
		Func: &proto.Type_Func{
			ParameterTypes: params,
			ReturnType:     TypeToProto(f.ReturnType),
			Nullability:    proto.Type_Nullability(f.Nullability),
		}}}
}

func listTypeToProto(t *types.ListType) *proto.Type {
	return &proto.Type{Kind: &proto.Type_List_{
		List: &proto.Type_List{Nullability: proto.Type_Nullability(t.Nullability),
			Type:                   TypeToProto(t.Type),
			TypeVariationReference: t.TypeVariationRef}}}
}

func mapTypeToProto(t *types.MapType) *proto.Type {
	return &proto.Type{Kind: &proto.Type_Map_{
		Map: &proto.Type_Map{Nullability: proto.Type_Nullability(t.Nullability),
			TypeVariationReference: t.TypeVariationRef,
			Key:                    TypeToProto(t.Key),
			Value:                  TypeToProto(t.Value)}}}
}

func userDefinedTypeToProto(t *types.UserDefinedType) *proto.Type {
	params := make([]*proto.Type_Parameter, len(t.TypeParameters))
	for i, p := range t.TypeParameters {
		params[i] = TypeParamToProto(p)
	}

	return &proto.Type{Kind: &proto.Type_UserDefined_{
		UserDefined: &proto.Type_UserDefined{
			Nullability:            proto.Type_Nullability(t.Nullability),
			TypeVariationReference: t.TypeVariationRef,
			TypeReference:          t.TypeReference,
			TypeParameters:         params,
		}}}
}

// TypeParamToProto encodes a user-defined-type parameter.
func TypeParamToProto(p types.TypeParam) *proto.Type_Parameter {
	switch p := p.(type) {
	case types.NullParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Null{}}
	case *types.DataTypeParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_DataType{
			DataType: TypeToProto(p.Type)}}
	case types.BooleanParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Boolean{
			Boolean: bool(p)}}
	case types.IntegerParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Integer{
			Integer: int64(p)}}
	case types.EnumParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_Enum{
			Enum: string(p)}}
	case types.StringParameter:
		return &proto.Type_Parameter{Parameter: &proto.Type_Parameter_String_{
			String_: string(p)}}
	}
	panic("unimplemented type parameter")
}

// VersionFromProto decodes a version from its protobuf message.
func VersionFromProto(v *proto.Version) types.Version {
	if v == nil {
		return types.Version{Producer: "UNSET"}
	}
	return types.Version{
		MajorNumber: v.MajorNumber,
		MinorNumber: v.MinorNumber,
		PatchNumber: v.PatchNumber,
		GitHash:     v.GitHash,
		Producer:    v.Producer,
	}
}

// FunctionOptionsToProto encodes domain FunctionOptions as their protobuf messages.
func FunctionOptionsToProto(opts []*types.FunctionOption) []*proto.FunctionOption {
	if opts == nil {
		return nil
	}
	out := make([]*proto.FunctionOption, len(opts))
	for i, o := range opts {
		if o == nil {
			continue
		}
		out[i] = &proto.FunctionOption{Name: o.Name, Preference: o.Preference}
	}
	return out
}

// FunctionOptionsFromProto decodes function option messages into domain FunctionOptions.
func FunctionOptionsFromProto(opts []*proto.FunctionOption) []*types.FunctionOption {
	if opts == nil {
		return nil
	}
	out := make([]*types.FunctionOption, len(opts))
	for i, o := range opts {
		if o == nil {
			continue
		}
		out[i] = &types.FunctionOption{Name: o.Name, Preference: o.Preference}
	}
	return out
}

// NamedStructFromProto decodes a named struct from its protobuf message.
func NamedStructFromProto(n *proto.NamedStruct) types.NamedStruct {
	if n == nil {
		return types.NamedStruct{}
	}

	fields := make([]types.Type, len(n.Struct.Types))
	for i, f := range n.Struct.Types {
		fields[i] = TypeFromProto(f)
	}

	return types.NamedStruct{
		Names: n.Names,
		Struct: types.StructType{
			Nullability:      types.Nullability(n.Struct.Nullability),
			TypeVariationRef: n.Struct.TypeVariationReference,
			Types:            fields,
		},
	}
}

// TypeParamFromProto decodes a protobuf Type_Parameter message into a TypeParam.
func TypeParamFromProto(p *proto.Type_Parameter) types.TypeParam {
	switch p := p.Parameter.(type) {
	case *proto.Type_Parameter_Null:
		return types.NullParameter{}
	case *proto.Type_Parameter_Boolean:
		return types.BooleanParameter(p.Boolean)
	case *proto.Type_Parameter_DataType:
		return &types.DataTypeParameter{Type: TypeFromProto(p.DataType)}
	case *proto.Type_Parameter_Integer:
		return types.IntegerParameter(p.Integer)
	case *proto.Type_Parameter_Enum:
		return types.EnumParameter(p.Enum)
	case *proto.Type_Parameter_String_:
		return types.StringParameter(p.String_)
	}
	return nil
}

// TypeFromProto returns the appropriate Type object from a protobuf type message.
func TypeFromProto(t *proto.Type) types.Type {
	switch t := t.Kind.(type) {
	case *proto.Type_Bool:
		return &types.BooleanType{
			Nullability:      types.Nullability(t.Bool.Nullability),
			TypeVariationRef: t.Bool.TypeVariationReference,
		}
	case *proto.Type_I8_:
		return &types.Int8Type{
			Nullability:      types.Nullability(t.I8.Nullability),
			TypeVariationRef: t.I8.TypeVariationReference,
		}
	case *proto.Type_I16_:
		return &types.Int16Type{
			Nullability:      types.Nullability(t.I16.Nullability),
			TypeVariationRef: t.I16.TypeVariationReference,
		}
	case *proto.Type_I32_:
		return &types.Int32Type{
			Nullability:      types.Nullability(t.I32.Nullability),
			TypeVariationRef: t.I32.TypeVariationReference,
		}
	case *proto.Type_I64_:
		return &types.Int64Type{
			Nullability:      types.Nullability(t.I64.Nullability),
			TypeVariationRef: t.I64.TypeVariationReference,
		}
	case *proto.Type_Fp32:
		return &types.Float32Type{
			Nullability:      types.Nullability(t.Fp32.Nullability),
			TypeVariationRef: t.Fp32.TypeVariationReference,
		}
	case *proto.Type_Fp64:
		return &types.Float64Type{
			Nullability:      types.Nullability(t.Fp64.Nullability),
			TypeVariationRef: t.Fp64.TypeVariationReference,
		}
	case *proto.Type_String_:
		return &types.StringType{
			Nullability:      types.Nullability(t.String_.Nullability),
			TypeVariationRef: t.String_.TypeVariationReference,
		}
	case *proto.Type_Binary_:
		return &types.BinaryType{
			Nullability:      types.Nullability(t.Binary.Nullability),
			TypeVariationRef: t.Binary.TypeVariationReference,
		}
	case *proto.Type_Timestamp_:
		return &types.TimestampType{
			Nullability:      types.Nullability(t.Timestamp.Nullability),
			TypeVariationRef: t.Timestamp.TypeVariationReference,
		}
	case *proto.Type_Date_:
		return &types.DateType{
			Nullability:      types.Nullability(t.Date.Nullability),
			TypeVariationRef: t.Date.TypeVariationReference,
		}
	case *proto.Type_Time_:
		return &types.TimeType{
			Nullability:      types.Nullability(t.Time.Nullability),
			TypeVariationRef: t.Time.TypeVariationReference,
		}
	case *proto.Type_IntervalYear_:
		return &types.IntervalYearType{
			Nullability:      types.Nullability(t.IntervalYear.Nullability),
			TypeVariationRef: t.IntervalYear.TypeVariationReference,
		}
	case *proto.Type_TimestampTz:
		return &types.TimestampTzType{
			Nullability:      types.Nullability(t.TimestampTz.Nullability),
			TypeVariationRef: t.TimestampTz.TypeVariationReference,
		}
	case *proto.Type_Uuid:
		return &types.UUIDType{
			Nullability:      types.Nullability(t.Uuid.Nullability),
			TypeVariationRef: t.Uuid.TypeVariationReference,
		}
	case *proto.Type_FixedBinary_:
		return &types.FixedBinaryType{
			Nullability:      types.Nullability(t.FixedBinary.Nullability),
			TypeVariationRef: t.FixedBinary.TypeVariationReference,
			Length:           t.FixedBinary.Length,
		}
	case *proto.Type_FixedChar_:
		return &types.FixedCharType{
			Nullability:      types.Nullability(t.FixedChar.Nullability),
			TypeVariationRef: t.FixedChar.TypeVariationReference,
			Length:           t.FixedChar.Length,
		}
	case *proto.Type_Varchar:
		return &types.VarCharType{
			Nullability:      types.Nullability(t.Varchar.Nullability),
			TypeVariationRef: t.Varchar.TypeVariationReference,
			Length:           t.Varchar.Length,
		}
	case *proto.Type_Decimal_:
		return &types.DecimalType{
			Nullability:      types.Nullability(t.Decimal.Nullability),
			TypeVariationRef: t.Decimal.TypeVariationReference,
			Scale:            t.Decimal.Scale,
			Precision:        t.Decimal.Precision,
		}
	case *proto.Type_Struct_:
		fields := make([]types.Type, len(t.Struct.Types))
		for i, f := range t.Struct.Types {
			fields[i] = TypeFromProto(f)
		}
		return &types.StructType{
			Nullability:      types.Nullability(t.Struct.Nullability),
			TypeVariationRef: t.Struct.TypeVariationReference,
			Types:            fields,
		}
	case *proto.Type_Func_:
		params := make([]types.Type, len(t.Func.ParameterTypes))
		for i, p := range t.Func.ParameterTypes {
			params[i] = TypeFromProto(p)
		}
		return &types.FuncType{
			Nullability:    types.Nullability(t.Func.Nullability),
			ParameterTypes: params,
			ReturnType:     TypeFromProto(t.Func.ReturnType),
		}
	case *proto.Type_List_:
		return &types.ListType{
			Nullability:      types.Nullability(t.List.Nullability),
			TypeVariationRef: t.List.TypeVariationReference,
			Type:             TypeFromProto(t.List.Type),
		}
	case *proto.Type_Map_:
		return &types.MapType{
			Nullability:      types.Nullability(t.Map.Nullability),
			TypeVariationRef: t.Map.TypeVariationReference,
			Key:              TypeFromProto(t.Map.Key),
			Value:            TypeFromProto(t.Map.Value),
		}
	case *proto.Type_UserDefined_:
		params := make([]types.TypeParam, len(t.UserDefined.TypeParameters))
		for i, p := range t.UserDefined.TypeParameters {
			params[i] = TypeParamFromProto(p)
		}
		return &types.UserDefinedType{
			Nullability:      types.Nullability(t.UserDefined.Nullability),
			TypeVariationRef: t.UserDefined.TypeVariationReference,
			TypeReference:    t.UserDefined.TypeReference,
			TypeParameters:   params,
		}
	}
	panic("unimplemented type from proto")
}
