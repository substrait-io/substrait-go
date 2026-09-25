// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"
	"strconv"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

// DynamicParameterBinding maps a parameter anchor to a literal value
// for use with DynamicParameter expressions in a plan.
//
// When bindings are provided via PlanWithBindings, the builder validates
// that each binding's literal type matches (ignoring nullability) the
// OutputType declared on the corresponding DynamicParameter expression.
type DynamicParameterBinding struct {
	ParameterAnchor uint32
	Value           expr.Literal
}

// ValidateParameterBindings checks that every binding's literal type matches
// the OutputType of the corresponding DynamicParameter expression found in
// the relation tree. Type comparison ignores nullability so that a required
// parameter can be bound to a nullable literal and vice-versa.
//
// Returns an error for:
//   - A binding whose anchor does not correspond to any DynamicParameter in the tree.
//   - A binding whose value type does not match the parameter's declared type
//     (ignoring nullability).
func ValidateParameterBindings(root Rel, bindings []DynamicParameterBinding) error {
	if len(bindings) == 0 {
		return nil
	}

	// Collect all DynamicParameter output types keyed by anchor.
	paramTypes := make(map[uint32]types.Type)
	collectDynamicParams(root, paramTypes)

	for _, b := range bindings {
		declaredType, ok := paramTypes[b.ParameterAnchor]
		if !ok {
			return fmt.Errorf("%w: parameter binding references anchor %d, "+
				"but no DynamicParameter with that reference exists in the plan",
				substraitgo.ErrInvalidPlan, b.ParameterAnchor)
		}

		// Compare ignoring nullability.
		bindingType := b.Value.GetType().WithNullability(types.NullabilityUnspecified)
		expectedType := declaredType.WithNullability(types.NullabilityUnspecified)
		if !bindingType.Equals(expectedType) {
			return fmt.Errorf("%w: parameter binding for anchor %d has type %s, "+
				"but DynamicParameter declares type %s",
				substraitgo.ErrInvalidPlan, b.ParameterAnchor, b.Value.GetType(), declaredType)
		}
	}

	return nil
}

// collectDynamicParams walks a relation tree and records the OutputType
// of every DynamicParameter expression it encounters.
func collectDynamicParams(rel Rel, out map[uint32]types.Type) {
	if rel == nil {
		return
	}

	// Walk child relations first.
	for _, child := range rel.GetInputs() {
		collectDynamicParams(child, out)
	}

	// Walk expressions owned by this relation.
	walkRelExprs(rel, func(e expr.Expression) {
		walkExpr(e, func(inner expr.Expression) {
			if dp, ok := inner.(*expr.DynamicParameter); ok {
				out[dp.ParameterReference] = dp.OutputType
			}
		})
	})
}

// walkRelExprs invokes fn for every top-level expression in a relation.
func walkRelExprs(rel Rel, fn func(expr.Expression)) {
	switch r := rel.(type) {
	case *FilterRel:
		fn(r.Condition())
	case *ProjectRel:
		for _, e := range r.Expressions() {
			fn(e)
		}
	case *JoinRel:
		fn(r.Expr())
		if pjf := r.PostJoinFilter(); pjf != nil {
			fn(pjf)
		}
	case *SortRel:
		for _, sf := range r.Sorts() {
			fn(sf.Expr)
		}
	}
}

// walkExpr recursively visits every node in an expression tree.
func walkExpr(e expr.Expression, fn func(expr.Expression)) {
	if e == nil {
		return
	}
	fn(e)
	e.Visit(func(child expr.Expression) expr.Expression {
		walkExpr(child, fn)
		return child
	})
}

// Stats are the statistics related to a Hint (physical properties of records).
type Stats struct {
	// RowCount is the estimated number of records produced by the relation.
	RowCount float64
	// RecordSize is the estimated physical size of each record.
	RecordSize float64
	// AdvancedExtension carries implementation-specific statistics details.
	AdvancedExtension *extensions.AdvancedExtension
}

// RuntimeConstraint describes constraints on the runtime environment carried by a Hint.
type RuntimeConstraint struct {
	// AdvancedExtension carries implementation-specific runtime constraints.
	AdvancedExtension *extensions.AdvancedExtension
}

// Hint carries changes to an operation that can influence efficiency/performance
// but should not impact correctness.
type Hint struct {
	// Stats are the physical property estimates for records produced by the relation.
	Stats *Stats
	// Constraint describes runtime requirements for evaluating the relation.
	Constraint *RuntimeConstraint
	// Alias is a name for qualifying or debugging the relation.
	Alias string
	// OutputNames assigns alternative names to the relation's output fields.
	OutputNames []string
	// AdvancedExtension carries implementation-specific hint details.
	AdvancedExtension *extensions.AdvancedExtension
	// SavedComputations describe computations saved by this relation for later reuse.
	SavedComputations []*SavedComputation
	// LoadedComputations describe saved computations loaded by this relation.
	LoadedComputations []*LoadedComputation
}

// ComputationType is the kind of a saved or loaded computation hint.
type ComputationType int32

const (
	ComputationTypeUnspecified ComputationType = 0
	ComputationTypeHashTable   ComputationType = 1
	ComputationTypeBloomFilter ComputationType = 2
	ComputationTypeUnknown     ComputationType = 9999
)

func (c ComputationType) String() string {
	switch c {
	case ComputationTypeUnspecified:
		return "COMPUTATION_TYPE_UNSPECIFIED"
	case ComputationTypeHashTable:
		return "COMPUTATION_TYPE_HASHTABLE"
	case ComputationTypeBloomFilter:
		return "COMPUTATION_TYPE_BLOOM_FILTER"
	case ComputationTypeUnknown:
		return "COMPUTATION_TYPE_UNKNOWN"
	default:
		return strconv.Itoa(int(c))
	}
}

// SavedComputation is a computation the plan saves once and may load multiple times.
type SavedComputation struct {
	// ComputationID is the plan-unique identifier for the saved computation.
	ComputationID int32
	// Type identifies the kind of computation being saved.
	Type ComputationType
	// AdvancedExtension carries implementation-specific saved computation details.
	AdvancedExtension *extensions.AdvancedExtension
}

// LoadedComputation references a previously SavedComputation by ID.
type LoadedComputation struct {
	// ComputationIDReference identifies a previously saved computation.
	ComputationIDReference int32
	// Type identifies the kind of computation being loaded.
	Type ComputationType
	// AdvancedExtension carries implementation-specific loaded computation details.
	AdvancedExtension *extensions.AdvancedExtension
}

// RelCommon is the common fields of all relational operators and is
// embedded in all of them.
type RelCommon struct {
	hint         *Hint
	mapping      []int32
	advExtension *extensions.AdvancedExtension
}

// NewRelCommon builds the common fields embedded in every relation.
func NewRelCommon(hint *Hint, mapping []int32, advExtension *extensions.AdvancedExtension) RelCommon {
	return RelCommon{hint: hint, mapping: mapping, advExtension: advExtension}
}

func (rc *RelCommon) remap(initial types.RecordType) types.RecordType {
	if rc.mapping == nil {
		return initial
	}

	outTypes := make([]types.Type, len(rc.mapping))

	for i, m := range rc.mapping {
		outTypes[i] = initial.GetFieldRef(m)
	}

	return *types.NewRecordTypeFromTypes(outTypes)
}

func (rc *RelCommon) OutputMapping() []int32 {
	if rc.mapping == nil {
		return nil
	}
	// Make a copy of the output mapping to prevent accidental modification.
	mapCopy := make([]int32, len(rc.mapping))
	copy(mapCopy, rc.mapping)
	return mapCopy
}

func (rc *RelCommon) setMapping(mapping []int32) {
	rc.mapping = mapping
}

func (rc *RelCommon) GetAdvancedExtension() *extensions.AdvancedExtension {
	return rc.advExtension
}

func (rc *RelCommon) SetAdvancedExtension(advExtension *extensions.AdvancedExtension) *extensions.AdvancedExtension {
	existing := rc.advExtension
	rc.advExtension = advExtension
	return existing
}

func (rc *RelCommon) Hint() *Hint {
	return rc.hint
}
