// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"github.com/substrait-io/substrait-go/v9/types"
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
