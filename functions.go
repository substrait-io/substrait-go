// SPDX-License-Identifier: Apache-2.0

package substraitgo

import (
	"strings"

	"github.com/substrait-io/substrait-go/proto"
	"golang.org/x/exp/slices"
	pb "google.golang.org/protobuf/proto"
)

func FuncArgsEqual(a, b FuncArg) bool {
	if a == b {
		return true
	}

	switch lhs := a.(type) {
	case Expression:
		rhs, ok := b.(Expression)
		if !ok {
			return false
		}

		return lhs.Equals(rhs)
	case Type:
		rhs, ok := b.(Type)
		if !ok {
			return false
		}

		return lhs.Equals(rhs)
	case Enum:
		rhs, ok := b.(Enum)
		if !ok {
			return false
		}

		return lhs == rhs
	}

	return false
}

type (
	SortField struct {
		Expr Expression
		Kind SortKind
	}

	Bound interface {
		ToProto() *proto.Expression_WindowFunction_Bound
	}

	PrecedingBound int64
	FollowingBound int64
	CurrentRow     struct{}
	Unbounded      struct{}
)

func (s *SortField) ToProto() *proto.SortField {
	ret := &proto.SortField{Expr: s.Expr.ToProto()}
	switch k := s.Kind.(type) {
	case SortDirection:
		ret.SortKind = &proto.SortField_Direction{
			Direction: proto.SortField_SortDirection(k)}
	case FunctionRef:
		ret.SortKind = &proto.SortField_ComparisonFunctionReference{
			ComparisonFunctionReference: uint32(k)}
	}

	return ret
}

func SortFieldFromProto(f *proto.SortField, baseSchema Type, ext ExtensionRegistry) (sf SortField, err error) {
	sf.Expr, err = ExprFromProto(f.Expr, baseSchema, ext)
	if err != nil {
		return
	}

	switch k := f.SortKind.(type) {
	case *proto.SortField_Direction:
		sf.Kind = SortDirection(k.Direction)
	case *proto.SortField_ComparisonFunctionReference:
		sf.Kind = FunctionRef(k.ComparisonFunctionReference)
	default:
		err = ErrNotImplemented
	}
	return
}

func (fb PrecedingBound) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_Preceding_{
			Preceding: &proto.Expression_WindowFunction_Bound_Preceding{Offset: int64(fb)},
		},
	}
}

func (fb FollowingBound) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_Following_{
			Following: &proto.Expression_WindowFunction_Bound_Following{Offset: int64(fb)},
		},
	}
}

func (CurrentRow) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_CurrentRow_{
			CurrentRow: &proto.Expression_WindowFunction_Bound_CurrentRow{},
		},
	}
}

func (Unbounded) ToProto() *proto.Expression_WindowFunction_Bound {
	return &proto.Expression_WindowFunction_Bound{
		Kind: &proto.Expression_WindowFunction_Bound_Unbounded_{
			Unbounded: &proto.Expression_WindowFunction_Bound_Unbounded{},
		}}
}

func BoundFromProto(b *proto.Expression_WindowFunction_Bound) Bound {
	switch t := b.Kind.(type) {
	case *proto.Expression_WindowFunction_Bound_Preceding_:
		return PrecedingBound(t.Preceding.Offset)
	case *proto.Expression_WindowFunction_Bound_CurrentRow_:
		return CurrentRow{}
	case *proto.Expression_WindowFunction_Bound_Following_:
		return FollowingBound(t.Following.Offset)
	case *proto.Expression_WindowFunction_Bound_Unbounded_:
		return Unbounded{}
	}

	// bound is optional
	return nil
}

type ScalarFunction struct {
	FuncRef uint32
	ID      ExtID

	Args       []FuncArg
	Options    []*FunctionOption
	OutputType Type
}

func (*ScalarFunction) isRootRef() {}

func (s *ScalarFunction) String() string {
	var b strings.Builder

	b.WriteString(s.ID.Name)
	b.WriteByte('(')

	for i, arg := range s.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}

	b.WriteString(") => ")
	b.WriteString(s.OutputType.String())

	return b.String()
}

func (s *ScalarFunction) GetType() Type { return s.OutputType }
func (s *ScalarFunction) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: s.ToProto(),
		},
	}
}

func (s *ScalarFunction) ToProto() *proto.Expression {
	args := make([]*proto.FunctionArgument, len(s.Args))
	for i, a := range s.Args {
		args[i] = a.ToProtoFuncArg()
	}

	return &proto.Expression{
		RexType: &proto.Expression_ScalarFunction_{
			ScalarFunction: &proto.Expression_ScalarFunction{
				FunctionReference: s.FuncRef,
				Options:           s.Options,
				OutputType:        TypeToProto(s.OutputType),
				Arguments:         args,
			},
		},
	}
}

func (s *ScalarFunction) Equals(rhs Expression) bool {
	other, ok := rhs.(*ScalarFunction)
	if !ok {
		return false
	}

	switch {
	case s.FuncRef != other.FuncRef:
		return false
	case !s.OutputType.Equals(other.OutputType):
		return false
	}

	for i := range s.Args {
		if !FuncArgsEqual(s.Args[i], other.Args[i]) {
			return false
		}
	}

	return slices.EqualFunc(s.Options, other.Options, func(a, b *FunctionOption) bool {
		return pb.Equal(a, b)
	})
}

type WindowFunction struct {
	FuncRef uint32
	ID      ExtID

	Args       []FuncArg
	Options    []*FunctionOption
	OutputType Type

	Phase      AggregationPhase
	Sorts      []SortField
	Invocation AggregationInvocation
	Partitions []Expression

	LowerBound, UpperBound Bound
}

func (*WindowFunction) isRootRef() {}

func (w *WindowFunction) String() string {
	var b strings.Builder

	b.WriteString(w.ID.Name)
	b.WriteByte('(')

	for i, arg := range w.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}

	b.WriteString(") => ")
	b.WriteString(w.OutputType.String())

	return b.String()
}
func (w *WindowFunction) GetType() Type              { return w.OutputType }
func (w *WindowFunction) Equals(rhs Expression) bool { return false }

func (w *WindowFunction) ToProto() *proto.Expression {
	var (
		args       []*proto.FunctionArgument
		sorts      []*proto.SortField
		parts      []*proto.Expression
		upperBound *proto.Expression_WindowFunction_Bound
		lowerBound *proto.Expression_WindowFunction_Bound
	)

	if len(w.Args) > 0 {
		args = make([]*proto.FunctionArgument, len(w.Args))
		for i, a := range w.Args {
			args[i] = a.ToProtoFuncArg()
		}
	}

	if len(w.Sorts) > 0 {
		sorts = make([]*proto.SortField, len(w.Sorts))
		for i, s := range w.Sorts {
			sorts[i] = s.ToProto()
		}
	}

	if len(w.Partitions) > 0 {
		parts = make([]*proto.Expression, len(w.Partitions))
		for i, p := range w.Partitions {
			parts[i] = p.ToProto()
		}
	}

	if w.UpperBound != nil {
		upperBound = w.UpperBound.ToProto()
	}

	if w.LowerBound != nil {
		lowerBound = w.LowerBound.ToProto()
	}

	return &proto.Expression{
		RexType: &proto.Expression_WindowFunction_{
			WindowFunction: &proto.Expression_WindowFunction{
				FunctionReference: w.FuncRef,
				Arguments:         args,
				Options:           w.Options,
				OutputType:        TypeToProto(w.OutputType),
				Phase:             w.Phase,
				Sorts:             sorts,
				Invocation:        w.Invocation,
				Partitions:        parts,
				LowerBound:        lowerBound,
				UpperBound:        upperBound,
			},
		},
	}
}

func (w *WindowFunction) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{
			Value: w.ToProto(),
		},
	}
}
