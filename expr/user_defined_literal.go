// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"

	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/types/known/anypb"
)

// UserDefinedLiteralValue is the value of a user-defined literal: the `val` oneof
// of the Substrait UserDefined literal message. It is one of UserDefinedValueAny
// or UserDefinedValueStruct.
type UserDefinedLiteralValue interface {
	isUserDefinedLiteralValue()
}

// UserDefinedValueAny holds a user-defined literal serialized as a type-specific
// protobuf message.
type UserDefinedValueAny struct {
	Value *anypb.Any
}

func (UserDefinedValueAny) isUserDefinedLiteralValue() {}

// UserDefinedValueStruct holds a user-defined literal serialized using the struct
// definition from its type declaration.
type UserDefinedValueStruct struct {
	Value StructLiteralValue
}

func (UserDefinedValueStruct) isUserDefinedLiteralValue() {}

// UserDefinedLiteral is a literal of a user-defined type, mirroring the fields of
// the Substrait UserDefined literal message: a type reference, the type's
// parameters, and the value.
type UserDefinedLiteral struct {
	TypeReference  uint32
	TypeParameters []types.TypeParam
	Val            UserDefinedLiteralValue
}

// userDefinedValueFromProto decodes the val oneof of a UserDefined literal into
// its domain form.
func userDefinedValueFromProto(ud *proto.Expression_Literal_UserDefined) UserDefinedLiteralValue {
	switch v := ud.Val.(type) {
	case *proto.Expression_Literal_UserDefined_Value:
		return UserDefinedValueAny{Value: v.Value}
	case *proto.Expression_Literal_UserDefined_Struct:
		//lint:ignore SA1019 StructLiteralFromProto is the struct-literal decoder available within this package
		return UserDefinedValueStruct{Value: StructLiteralFromProto(v.Struct)}
	}
	return nil
}

// setUserDefinedVal encodes a domain UserDefinedLiteralValue onto the val oneof of
// a protobuf UserDefined literal.
func setUserDefinedVal(ud *proto.Expression_Literal_UserDefined, v UserDefinedLiteralValue) {
	switch val := v.(type) {
	case UserDefinedValueAny:
		ud.Val = &proto.Expression_Literal_UserDefined_Value{Value: val.Value}
	case UserDefinedValueStruct:
		ud.Val = &proto.Expression_Literal_UserDefined_Struct{Struct: val.Value.ToProto()}
	default:
		panic(fmt.Sprintf("unhandled UserDefinedLiteralValue %T", v))
	}
}
