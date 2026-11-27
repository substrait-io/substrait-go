// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"
	"slices"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/extensions"
	"github.com/substrait-io/substrait-go/v9/types"
)

// TableFunction is a call to a table function. It is not an Expression: it
// produces rows rather than a value, and is invoked by a table function
// relation.
type TableFunction struct {
	funcRef     uint32
	declaration *extensions.TableFunctionVariant

	args    []types.FuncArg
	options []*types.FunctionOption
	schema  *types.StructType
}

// NewTableFunc validates that the specified ID can be found in the registry
// via GetTableFunc and retrieves a function anchor to use. A simple name is
// resolved against the argument types as in NewScalarFunc.
//
// If schema is nil, it is derived from the implementation's schema expression
// and the argument types, and it is an error for the implementation not to
// declare one. Otherwise schema must be a non-nullable struct and, when the
// implementation declares a schema expression, must match its evaluation.
//
// Options are not validated against the function declaration.
func NewTableFunc(
	reg ExtensionRegistry, id extensions.FunctionID, schema *types.StructType,
	opts []*types.FunctionOption, args ...types.FuncArg,
) (*TableFunction, error) {
	decl, argTypes, err := lookupVariant(id, reg, reg.c.GetTableFunc, args)
	if err != nil {
		return nil, err
	}

	if decl.Schema() == nil {
		if schema == nil {
			return nil, fmt.Errorf("%w: table function %s does not declare a schema, one must be provided",
				substraitgo.ErrInvalidType, decl.CompoundName())
		}
		match, err := decl.Match(argTypes)
		if err != nil {
			return nil, err
		}
		if !match {
			return nil, fmt.Errorf("%w: argument types did not match table function %s",
				substraitgo.ErrInvalidType, decl.CompoundName())
		}
	} else {
		resolved, err := decl.ResolveType(argTypes, reg.Set)
		if err != nil {
			return nil, err
		}
		derived, ok := resolved.(*types.StructType)
		if !ok {
			return nil, fmt.Errorf("%w: table function %s schema resolved to %s, expected a struct",
				substraitgo.ErrInvalidType, decl.CompoundName(), resolved)
		}
		if schema == nil {
			schema = derived
		} else if !schema.Equals(derived) {
			return nil, fmt.Errorf("%w: table function %s schema %s does not match the declared schema %s",
				substraitgo.ErrInvalidType, decl.CompoundName(), schema, derived)
		}
	}

	if err := validateTableFunctionSchema(schema); err != nil {
		return nil, err
	}

	return &TableFunction{
		funcRef:     reg.GetFuncAnchor(decl.ID()),
		declaration: decl,
		args:        args,
		options:     opts,
		schema:      schema,
	}, nil
}

// NewCustomTableFunc doesn't validate that the variant can be found in the
// registry or that the arguments match it. The registry is only used to
// provide a function anchor that can be used when serializing the call to
// Protobuf. The schema must be a non-nullable struct.
func NewCustomTableFunc(
	reg ExtensionRegistry, v *extensions.TableFunctionVariant, schema *types.StructType,
	opts []*types.FunctionOption, args ...types.FuncArg,
) (*TableFunction, error) {
	if err := validateTableFunctionSchema(schema); err != nil {
		return nil, err
	}

	return &TableFunction{
		funcRef:     reg.GetFuncAnchor(v.ID()),
		declaration: v,
		args:        args,
		options:     opts,
		schema:      schema,
	}, nil
}

// NewTableFunctionFromParts assembles a TableFunction from decoded parts
// without the validation NewTableFunc performs. Prefer the builder. (issue #358)
func NewTableFunctionFromParts(
	funcRef uint32, declaration *extensions.TableFunctionVariant,
	args []types.FuncArg, options []*types.FunctionOption, schema *types.StructType,
) *TableFunction {
	return &TableFunction{
		funcRef:     funcRef,
		declaration: declaration,
		args:        args,
		options:     options,
		schema:      schema,
	}
}

func validateTableFunctionSchema(schema *types.StructType) error {
	if schema == nil {
		return fmt.Errorf("%w: table function schema must not be nil", substraitgo.ErrInvalidType)
	}
	if schema.Nullability != types.NullabilityRequired {
		return fmt.Errorf("%w: table function schema must be a non-nullable struct, got %s",
			substraitgo.ErrInvalidType, schema)
	}
	return nil
}

func (t *TableFunction) Name() string                           { return t.declaration.Name() }
func (t *TableFunction) CompoundName() string                   { return t.declaration.CompoundName() }
func (t *TableFunction) ID() extensions.FunctionID              { return t.declaration.ID() }
func (t *TableFunction) Variadic() *extensions.VariadicBehavior { return t.declaration.Variadic() }
func (t *TableFunction) NArgs() int                             { return len(t.args) }
func (t *TableFunction) Arg(i int) types.FuncArg                { return t.args[i] }
func (t *TableFunction) FuncRef() uint32                        { return t.funcRef }

// Schema returns the type of every row the call produces.
func (t *TableFunction) Schema() *types.StructType { return t.schema }

func (t *TableFunction) GetOption(name string) []string {
	for _, o := range t.options {
		if name == o.Name {
			return o.Preference
		}
	}
	return nil
}

func (t *TableFunction) GetOptions() []*types.FunctionOption { return t.options }

func (t *TableFunction) GetArgTypes() []types.Type {
	return getArgTypes(t.args)
}

func (t *TableFunction) Clone() *TableFunction {
	out := *t
	if t.args != nil {
		out.args = slices.Clone(t.args)
	}
	if t.options != nil {
		out.options = slices.Clone(t.options)
	}
	return &out
}

// SetArg sets the specified argument to the provided value. The index is not checked for validity.
func (t *TableFunction) SetArg(i int, arg types.FuncArg) {
	t.args[i] = arg
}

func (t *TableFunction) String() string {
	var b strings.Builder

	b.WriteString(t.Name())
	b.WriteByte('(')

	for i, arg := range t.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}

	if len(t.options) > 0 {
		b.WriteString(", {")
		for i, o := range t.options {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(o.Name)
			b.WriteString(": [")
			b.WriteString(strings.Join(o.Preference, ","))
			b.WriteString("]")
		}
		b.WriteString("}")
	}

	b.WriteString(") => ")
	if t.schema != nil {
		b.WriteString(t.schema.String())
	} else {
		b.WriteString("?")
	}

	return b.String()
}

func (t *TableFunction) Equals(other *TableFunction) bool {
	switch {
	case t == other:
		return true
	case t == nil || other == nil:
		return false
	case t.funcRef != other.funcRef:
		return false
	case !t.schema.Equals(other.schema):
		return false
	case !slices.EqualFunc(t.args, other.args, FuncArgsEqual):
		return false
	}

	return slices.EqualFunc(t.options, other.options, func(a, b *types.FunctionOption) bool {
		if a == nil || b == nil {
			return a == b
		}
		return a.Name == b.Name && slices.Equal(a.Preference, b.Preference)
	})
}
