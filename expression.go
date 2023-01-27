// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
)

func FuncArgFromProto(e *proto.FunctionArgument, baseSchema Type, ext ExtensionRegistry) (FuncArg, error) {
	switch et := e.ArgType.(type) {
	case *proto.FunctionArgument_Enum:
		return Enum(et.Enum), nil
	case *proto.FunctionArgument_Type:
		return TypeFromProto(et.Type), nil
	case *proto.FunctionArgument_Value:
		return ExprFromProto(et.Value, baseSchema, ext)
	}
	return nil, ErrNotImplemented
}

func ExprFromProto(e *proto.Expression, baseSchema Type, ext ExtensionRegistry) (Expression, error) {
	switch et := e.RexType.(type) {
	case *proto.Expression_Literal_:
		return LiteralFromProto(et.Literal), nil
	case *proto.Expression_Selection:
		return FieldReferenceFromProto(et.Selection, baseSchema, ext)
	case *proto.Expression_ScalarFunction_:
		var err error
		args := make([]FuncArg, len(et.ScalarFunction.Arguments))
		for i, a := range et.ScalarFunction.Arguments {
			if args[i], err = FuncArgFromProto(a, baseSchema, ext); err != nil {
				return nil, err
			}
		}

		id, ok := ext.DecodeFunc(et.ScalarFunction.FunctionReference)
		if !ok {
			return nil, ErrNotFound
		}

		return &ScalarFunction{
			FuncRef:    et.ScalarFunction.FunctionReference,
			ID:         id,
			Args:       args,
			Options:    et.ScalarFunction.Options,
			OutputType: TypeFromProto(et.ScalarFunction.OutputType),
		}, nil
	case *proto.Expression_WindowFunction_:
		var err error
		args := make([]FuncArg, len(et.WindowFunction.Arguments))
		for i, a := range et.WindowFunction.Arguments {
			if args[i], err = FuncArgFromProto(a, baseSchema, ext); err != nil {
				return nil, err
			}
		}

		parts := make([]Expression, len(et.WindowFunction.Partitions))
		for i, p := range et.WindowFunction.Partitions {
			if parts[i], err = ExprFromProto(p, baseSchema, ext); err != nil {
				return nil, err
			}
		}

		sorts := make([]SortField, len(et.WindowFunction.Sorts))
		for i, s := range et.WindowFunction.Sorts {
			if sorts[i], err = SortFieldFromProto(s, baseSchema, ext); err != nil {
				return nil, err
			}
		}

		id, ok := ext.DecodeFunc(et.WindowFunction.FunctionReference)
		if !ok {
			return nil, ErrNotFound
		}

		return &WindowFunction{
			FuncRef:    et.WindowFunction.FunctionReference,
			ID:         id,
			Args:       args,
			Options:    et.WindowFunction.Options,
			OutputType: TypeFromProto(et.WindowFunction.OutputType),
			Phase:      et.WindowFunction.Phase,
			Invocation: et.WindowFunction.Invocation,
			Partitions: parts,
			Sorts:      sorts,
			LowerBound: BoundFromProto(et.WindowFunction.LowerBound),
			UpperBound: BoundFromProto(et.WindowFunction.UpperBound),
		}, nil
	case *proto.Expression_IfThen_:
	case *proto.Expression_SwitchExpression_:
	case *proto.Expression_SingularOrList_:
	case *proto.Expression_MultiOrList_:
	case *proto.Expression_Cast_:
	case *proto.Expression_Nested_:
	case *proto.Expression_Enum_:
		return nil, fmt.Errorf("%w: deprecated", ErrNotImplemented)
	case *proto.Expression_Subquery_:
	}

	return nil, ErrNotImplemented
}

type VisitFunc func(Expression) Expression

// Expression can be one of many different things as a generalized
// expression. It could be:
//
//  - A literal
//  - A Field Reference Selection
//  - A Scalar Function expression
//  - A Window Function expression
//  - An If-Then statement
//  - A Switch Expression
//  - A Singular Or List
//  - A Multiple Or List
//  - A Cast expression
//  - A Subquery
//  - A Nested expression
type Expression interface {
	// an Expression can also be a function argument
	FuncArg
	// an expression can also be the root of a reference
	RootRefType

	// GetType returns the output type of this expression
	GetType() Type
	// ToProto converts this Expression and its arguments
	// to the equivalent Protobuf objects.
	ToProto() *proto.Expression
	// Equals returns true if this expression and all of its
	// arguments and their children etc. are equal to the passed
	// in Expression.
	Equals(Expression) bool
	// Visit invokes the passed visit function for each child of the
	// expression. The visit function can return its input expression
	// as-is with no changes, or it can construct and return a
	// replacement expression. If any children have been replaced, Visit
	// will construct and return a new instance of this expression using
	// the new children. Callers can use the Visit method to traverse
	// and potentially rewrite the expression tree, in either pre or post
	// order. Here is a pre-order example:
	//
	//   func preOrderVisit(e Expression) Expression {
	//     // Replace some scalar function, leave everything else
	//     // as-is. This check is before the call to Visit, so
	//     // it's a pre-order traversal
	//     if f, ok := e.(*ScalarFunction); ok {
	//       return &ScalarFunction{
	//         ID: ExtID{URI: "some other uri", Name: "some other func"},
	//         Args: f.Args,
	//         Options: f.Options,
	//         OutputType: f.OutputType,
	//       }
	//     }
	//     return e.Visit(preOrderVisit)
	//   }
	//   newExpr := preOrderVisit(oldExpr)
	Visit(VisitFunc) Expression
}

type IfThen struct {
	IFs []struct {
		If   Expression
		Then Expression
	}
	Else Expression
}

type Cast struct {
	Type            Type
	Input           Expression
	FailureBehavior CastFailBehavior
}

type SwitchExpr struct {
	Match Expression
	IFs   []struct {
		If   Literal
		Then Expression
	}
	Else Expression
}

type SinglularOrList struct {
	Value   Expression
	Options []Expression
}

type MultiOrList struct {
	Value   Expression
	Options []struct {
		Fields []Expression
	}
}

type NestedExpr interface {
	FuncArg

	IsNullable() bool
	TypeVariation() uint32
}

type MapExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	KeyValues        []struct{ Key, Value Expression }
}

func (m *MapExpr) IsNullable() bool      { return m.Nullable }
func (m *MapExpr) TypeVariation() uint32 { return m.TypeVariationRef }

type StructExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Fields           []Expression
}

func (s *StructExpr) IsNullable() bool      { return s.Nullable }
func (s *StructExpr) TypeVariation() uint32 { return s.TypeVariationRef }

type ListExpr struct {
	Nullable         bool
	TypeVariationRef uint32
	Values           []Expression
}

func (l *ListExpr) IsNullable() bool      { return l.Nullable }
func (l *ListExpr) TypeVariation() uint32 { return l.TypeVariationRef }
