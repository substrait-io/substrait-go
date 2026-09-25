// SPDX-License-Identifier: Apache-2.0

//lint:file-ignore SA1019 Using a deprecated function, variable, constant or field

package plan

import (
	"fmt"
	"runtime/debug"
	"slices"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

var CurrentVersion = types.Version{
	MajorNumber: 0,
	MinorNumber: 29,
	PatchNumber: 0,
	Producer:    "substrait-go",
}

func init() {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if strings.HasPrefix(dep.Path, "github.com/substrait-io/substrait-go/v9") {
				CurrentVersion.Producer += " " + dep.Version
				break
			}
		}

		var goarch, goos string
		for _, s := range info.Settings {
			if s.Key == "GOARCH" {
				goarch = s.Value
			} else if s.Key == "GOOS" {
				goos = s.Value
			}
		}

		if goos != "" && goarch != "" {
			CurrentVersion.Producer += " " + goos + "/" + goarch
		}
	}
}

// Relation is either a Root relation (a relation + list of column names)
// or another relation (such as a CTE or other reference).
type Relation struct {
	root *Root
	rel  Rel
}

// NewRelation builds a top-level plan relation from either a root or a plain
// relation (exactly one is non-nil).
func NewRelation(root *Root, rel Rel) Relation {
	return Relation{root: root, rel: rel}
}

// IsRoot returns true if this is the root of the plan Relation tree.
func (r *Relation) IsRoot() bool {
	return r.root != nil
}

func (r *Relation) Root() *Root { return r.root }
func (r *Relation) Rel() Rel    { return r.rel }

type AdvancedExtension interface {
	GetEnhancement() *extensions.Enhancement
	GetOptimizations() []*extensions.Optimization
}

// Plan describes a set of operations to complete. For
// compactness, identifiers are normalized at the plan level.
type Plan struct {
	version           types.Version
	extensions        extensions.Set
	expectedTypeURLs  []string
	advExtension      *extensions.AdvancedExtension
	relations         []Relation
	parameterBindings []DynamicParameterBinding

	reg expr.ExtensionRegistry
}

// NewPlan assembles a decoded plan from its finished parts. The registry is
// built by the caller.
func NewPlan(version types.Version, extSet extensions.Set, advExtension *extensions.AdvancedExtension, expectedTypeURLs []string, relations []Relation, parameterBindings []DynamicParameterBinding, reg expr.ExtensionRegistry) *Plan {
	return &Plan{
		version:           version,
		extensions:        extSet,
		advExtension:      advExtension,
		expectedTypeURLs:  expectedTypeURLs,
		relations:         relations,
		parameterBindings: parameterBindings,
		reg:               reg,
	}
}

// Version returns the plan's version.
func (p *Plan) Version() types.Version { return p.version }

// ExtensionRegistry returns the set of registered extensions for this plan
// that it may depend on.
func (p *Plan) ExtensionRegistry() expr.ExtensionRegistry { return p.reg }

// ExpectedTypeURLs is a list of anypb.Any protobuf entities that this plan
// may use. This can be used to warn if some embedded message types are
// unknown. Note that this list may include message types which are ignorable
// (optimizations) or are unused. In many cases, a consumer may be able to
// work with a plan even if one or more message types defined here are unknown.
//
// This returns a clone of the slice, so that the Plan itself remains
// immutable.
func (p *Plan) ExpectedTypeURLs() []string {
	return slices.Clone(p.expectedTypeURLs)
}

// AdvancedExtension returns optional additional extensions associated with
// this plan such as optimizations or enhancements.
func (p *Plan) AdvancedExtension() AdvancedExtension { return p.advExtension }

// GetAdvancedExtension returns the plan's advanced extension as its concrete
// type, matching the accessor the relations expose.
func (p *Plan) GetAdvancedExtension() *extensions.AdvancedExtension { return p.advExtension }

// Relations returns the full slice of relation trees that are in this plan.
//
// This returns a clone of the internal slice so that the plan itself remains
// immutable.
func (p *Plan) Relations() []Relation {
	return slices.Clone(p.relations)
}

// GetRoots returns a slice containing *only* the relations which are
// considered Root relations from the list (as opposed to CTEs or references).
func (p *Plan) GetRoots() (roots []*Root) {
	roots = make([]*Root, 0, 1)
	for _, r := range p.relations {
		if r.IsRoot() {
			roots = append(roots, r.root)
		}
	}
	return roots
}

// GetNonRootRelations returns a slice containing only the relations from
// this plan which are not considered Roots.
func (p *Plan) GetNonRootRelations() (rels []Rel) {
	rels = make([]Rel, 0, 1)
	for _, r := range p.relations {
		if !r.IsRoot() {
			rels = append(rels, r.rel)
		}
	}
	return rels
}

// ParameterBindings returns the list of dynamic parameter bindings for this plan.
// Each binding maps a parameter anchor to a runtime literal value.
//
// This returns a clone of the internal slice so that the plan itself remains immutable.
func (p *Plan) ParameterBindings() []DynamicParameterBinding {
	return slices.Clone(p.parameterBindings)
}

// Root is a relation with output field names.
// This is used as the root of a Rel tree.
type Root struct {
	input Rel
	names []string
}

// NewRoot builds a root relation from its input and output names.
func NewRoot(input Rel, names []string) *Root {
	return &Root{input: input, names: names}
}

func (r *Root) Input() Rel { return r.input }

// Names are the field names in depth-first order.
func (r *Root) Names() []string { return r.names }

func (r *Root) RecordType() types.NamedStruct {
	return types.NamedStruct{
		Names:  r.names,
		Struct: *r.input.RecordType().AsStructType(),
	}
}

type RewriteFunc func(expr.Expression) (expr.Expression, error)

// Rel is a relation tree, representing one of the expected Relation
// types such as Fetch, Sort, Filter, Join, etc.
//
// It contains the common functionality between the different relations
// and should be type switched to determine which relation type it actually
// is for evaluation.
//
// All the exported methods in this interface should be considered constant.
type Rel interface {
	// Hint returns a set of changes to the operation which can influence
	// efficiency and performance but should not impact correctness.
	//
	// This includes things such as Stats and Runtime constraints.
	Hint() *Hint
	// OutputMapping is optional and may be nil. If this is nil, then
	// the result of this relation is the direct output as is (with no
	// reordering or projection of columns). Otherwise this is a slice
	// of indices into the underlying relation's output to map columns
	// to the intended result column order.
	//
	// For example, an output map of [5, 2, 1] means that the expected
	// result should be 3 columns consisting of the 5th, 2nd and 1st
	// output columns from the underlying relation.
	OutputMapping() []int32

	// Remap modifies the current relation by applying the provided
	// mapping to the current relation.  Typically used to remove any
	// unneeded columns or provide them in a different order. If there
	// already is a mapping on this relation, this provides mapping over
	// the current mapping.
	//
	// If any column numbers specified are outside the currently available
	// input range an error is returned and the mapping is left unchanged.
	//
	// If Remap is called with no arguments, an empty mapping will be set
	// on the relation, which removes ALL columns.

	Remap(mapping ...int32) (Rel, error)

	// setMapping sets the current mapping and is for internal use.
	// It performs no checks.  End users should call Remap() instead.
	setMapping(mapping []int32)

	// directOutputSchema returns the output record type of the underlying
	// relation as a struct type.  Mapping is not applied.
	directOutputSchema() types.RecordType
	// RecordType returns the types used by all columns returned by
	// this relation after applying any provided mapping.
	RecordType() types.RecordType

	GetAdvancedExtension() *extensions.AdvancedExtension
	// SetAdvancedExtension sets an AdvancedExtension on this Rel, returning any existing one on this Rel. Use `nil` to remove any existing AdvancedExtension.
	SetAdvancedExtension(extension *extensions.AdvancedExtension) (existing *extensions.AdvancedExtension)

	// Copy creates a copy of this relation with new inputs
	Copy(newInputs ...Rel) (Rel, error)

	// GetInputs returns a list of zero or more inputs for this relation
	GetInputs() []Rel

	// CopyWithExpressionRewrite rewrites all expression trees in this Rel. Returns original Rel
	// if no changes were made, otherwise a newly created rel that includes the given expressions
	CopyWithExpressionRewrite(rewriteFunc RewriteFunc, newInputs ...Rel) (Rel, error)
}

func validateRootNamesForSchema(recordType types.RecordType, names []string) error {
	expected := recordType.AsStructType().DepthFirstNameCount()
	if len(names) != expected {
		return fmt.Errorf("%w: root relation has %d output name(s) but the output schema requires %d",
			substraitgo.ErrInvalidRel, len(names), expected)
	}
	return nil
}
