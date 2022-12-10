// SPDX-License-Identifier: Apache-2.0

package substraitgo

type (
	SortField struct {
		Expr Expression
		Kind SortKind
	}

	Bound interface {
		isBound()
	}

	PrecedingBound int64
	FollowingBound int64
	CurrentRow     struct{}
	Unbounded      struct{}
)

func (PrecedingBound) isBound() {}
func (FollowingBound) isBound() {}
func (CurrentRow) isBound()     {}
func (Unbounded) isBound()      {}

type ScalarFunction struct {
	funcArg

	FuncRef    uint32
	Args       []FuncArg
	Options    []*FunctionOption
	OutputType Type
}

type WindowFunction struct {
	funcArg

	FuncRef    uint32
	Args       []FuncArg
	Options    []*FunctionOption
	OutputType Type

	Phase      AggregationPhase
	Sorts      []SortField
	Invocation AggregationInvocation
	Partitions []Expression

	LowerBound, UpperBound Bound
}
