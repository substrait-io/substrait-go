// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"
	"strconv"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
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
	RowCount          float64
	RecordSize        float64
	AdvancedExtension *extensions.AdvancedExtension
}

// RuntimeConstraint describes constraints on the runtime environment carried by a Hint.
type RuntimeConstraint struct {
	AdvancedExtension *extensions.AdvancedExtension
}

// Hint carries changes to an operation that can influence efficiency/performance
// but should not impact correctness.
type Hint struct {
	Stats              *Stats
	Constraint         *RuntimeConstraint
	Alias              string
	OutputNames        []string
	AdvancedExtension  *extensions.AdvancedExtension
	SavedComputations  []*SavedComputation
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
	ComputationID     int32
	Type              ComputationType
	AdvancedExtension *extensions.AdvancedExtension
}

// SavedComputationFromProto converts a protobuf SavedComputation message to the domain type.
func SavedComputationFromProto(s *proto.RelCommon_Hint_SavedComputation) *SavedComputation {
	if s == nil {
		return nil
	}
	return &SavedComputation{
		ComputationID:     s.ComputationId,
		Type:              ComputationType(s.Type),
		AdvancedExtension: extensions.AdvancedExtensionFromProto(s.AdvancedExtension),
	}
}

// SavedComputationToProto encodes a domain SavedComputation as its protobuf message.
func SavedComputationToProto(s *SavedComputation) *proto.RelCommon_Hint_SavedComputation {
	if s == nil {
		return nil
	}
	return &proto.RelCommon_Hint_SavedComputation{
		ComputationId:     s.ComputationID,
		Type:              proto.RelCommon_Hint_ComputationType(s.Type),
		AdvancedExtension: extensions.AdvancedExtensionToProto(s.AdvancedExtension),
	}
}

// LoadedComputation references a previously SavedComputation by ID.
type LoadedComputation struct {
	ComputationIDReference int32
	Type                   ComputationType
	AdvancedExtension      *extensions.AdvancedExtension
}

// LoadedComputationFromProto converts a protobuf LoadedComputation message to the domain type.
func LoadedComputationFromProto(l *proto.RelCommon_Hint_LoadedComputation) *LoadedComputation {
	if l == nil {
		return nil
	}
	return &LoadedComputation{
		ComputationIDReference: l.ComputationIdReference,
		Type:                   ComputationType(l.Type),
		AdvancedExtension:      extensions.AdvancedExtensionFromProto(l.AdvancedExtension),
	}
}

// LoadedComputationToProto encodes a domain LoadedComputation as its protobuf message.
func LoadedComputationToProto(l *LoadedComputation) *proto.RelCommon_Hint_LoadedComputation {
	if l == nil {
		return nil
	}
	return &proto.RelCommon_Hint_LoadedComputation{
		ComputationIdReference: l.ComputationIDReference,
		Type:                   proto.RelCommon_Hint_ComputationType(l.Type),
		AdvancedExtension:      extensions.AdvancedExtensionToProto(l.AdvancedExtension),
	}
}

// StatsFromProto converts a protobuf Stats message to the domain type.
func StatsFromProto(s *proto.RelCommon_Hint_Stats) *Stats {
	if s == nil {
		return nil
	}
	return &Stats{
		RowCount:          s.RowCount,
		RecordSize:        s.RecordSize,
		AdvancedExtension: extensions.AdvancedExtensionFromProto(s.AdvancedExtension),
	}
}

// StatsToProto encodes a domain Stats as its protobuf message.
func StatsToProto(s *Stats) *proto.RelCommon_Hint_Stats {
	if s == nil {
		return nil
	}
	return &proto.RelCommon_Hint_Stats{
		RowCount:          s.RowCount,
		RecordSize:        s.RecordSize,
		AdvancedExtension: extensions.AdvancedExtensionToProto(s.AdvancedExtension),
	}
}

// RuntimeConstraintFromProto converts a protobuf RuntimeConstraint message to the domain type.
func RuntimeConstraintFromProto(rc *proto.RelCommon_Hint_RuntimeConstraint) *RuntimeConstraint {
	if rc == nil {
		return nil
	}
	return &RuntimeConstraint{
		AdvancedExtension: extensions.AdvancedExtensionFromProto(rc.AdvancedExtension),
	}
}

// RuntimeConstraintToProto encodes a domain RuntimeConstraint as its protobuf message.
func RuntimeConstraintToProto(rc *RuntimeConstraint) *proto.RelCommon_Hint_RuntimeConstraint {
	if rc == nil {
		return nil
	}
	return &proto.RelCommon_Hint_RuntimeConstraint{
		AdvancedExtension: extensions.AdvancedExtensionToProto(rc.AdvancedExtension),
	}
}

// HintFromProto converts a protobuf Hint message to the domain type.
func HintFromProto(h *proto.RelCommon_Hint) *Hint {
	if h == nil {
		return nil
	}
	var saved []*SavedComputation
	if h.SavedComputations != nil {
		saved = make([]*SavedComputation, len(h.SavedComputations))
		for i, s := range h.SavedComputations {
			saved[i] = SavedComputationFromProto(s)
		}
	}
	var loaded []*LoadedComputation
	if h.LoadedComputations != nil {
		loaded = make([]*LoadedComputation, len(h.LoadedComputations))
		for i, l := range h.LoadedComputations {
			loaded[i] = LoadedComputationFromProto(l)
		}
	}
	return &Hint{
		Stats:              StatsFromProto(h.Stats),
		Constraint:         RuntimeConstraintFromProto(h.Constraint),
		Alias:              h.Alias,
		OutputNames:        h.OutputNames,
		AdvancedExtension:  extensions.AdvancedExtensionFromProto(h.AdvancedExtension),
		SavedComputations:  saved,
		LoadedComputations: loaded,
	}
}

// HintToProto encodes a domain Hint as its protobuf message.
func HintToProto(h *Hint) *proto.RelCommon_Hint {
	if h == nil {
		return nil
	}
	var saved []*proto.RelCommon_Hint_SavedComputation
	if h.SavedComputations != nil {
		saved = make([]*proto.RelCommon_Hint_SavedComputation, len(h.SavedComputations))
		for i, s := range h.SavedComputations {
			saved[i] = SavedComputationToProto(s)
		}
	}
	var loaded []*proto.RelCommon_Hint_LoadedComputation
	if h.LoadedComputations != nil {
		loaded = make([]*proto.RelCommon_Hint_LoadedComputation, len(h.LoadedComputations))
		for i, l := range h.LoadedComputations {
			loaded[i] = LoadedComputationToProto(l)
		}
	}
	return &proto.RelCommon_Hint{
		Stats:              StatsToProto(h.Stats),
		Constraint:         RuntimeConstraintToProto(h.Constraint),
		Alias:              h.Alias,
		OutputNames:        h.OutputNames,
		AdvancedExtension:  extensions.AdvancedExtensionToProto(h.AdvancedExtension),
		SavedComputations:  saved,
		LoadedComputations: loaded,
	}
}

// RelCommon is the common fields of all relational operators and is
// embedded in all of them.
type RelCommon struct {
	hint         *Hint
	mapping      []int32
	advExtension *extensions.AdvancedExtension
}

func (rc *RelCommon) fromProtoCommon(c *proto.RelCommon) {
	rc.hint = HintFromProto(c.GetHint())
	rc.advExtension = extensions.AdvancedExtensionFromProto(c.GetAdvancedExtension())

	if emit, ok := c.GetEmitKind().(*proto.RelCommon_Emit_); ok {
		rc.mapping = emit.Emit.OutputMapping
	} else {
		rc.mapping = nil
	}
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

func (rc *RelCommon) toProto() *proto.RelCommon {
	ret := &proto.RelCommon{
		Hint:              HintToProto(rc.hint),
		AdvancedExtension: extensions.AdvancedExtensionToProto(rc.advExtension),
	}

	if rc.mapping == nil {
		ret.EmitKind = &proto.RelCommon_Direct_{
			Direct: &proto.RelCommon_Direct{},
		}
	} else {
		ret.EmitKind = &proto.RelCommon_Emit_{
			Emit: &proto.RelCommon_Emit{OutputMapping: rc.mapping},
		}
	}
	return ret
}
