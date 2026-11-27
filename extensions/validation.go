// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/types/parser"
)

// validateUserDefinedTypeReferences ensures every user-defined type used in a
// function signature is declared in that same file's types section. It does
// not check type structures or type variation parents.
func (s *SimpleExtensionFile) validateUserDefinedTypeReferences() error {
	declared := make(map[string]struct{}, len(s.Types))
	for _, t := range s.Types {
		declared[t.Name] = struct{}{}
	}

	// Local references are checked against this file's own declared types,
	// before they are committed to the collection registry, so a failed load
	// leaves the collection untouched.
	//
	// Dependency-qualified references (non-nil DependencyAlias) are skipped for
	// now. A later PR will resolve them against the registry (c.GetType under
	// the dependency's URN), since a dependency's types are already loaded by
	// the time this file is.
	for _, typ := range s.functionTypes() {
		for _, ref := range types.ReferencedUserDefinedTypes(typ) {
			if ref.DependencyAlias == nil {
				if _, ok := declared[ref.Name]; !ok {
					return fmt.Errorf("%w: user-defined type %q is not declared",
						substraitgo.ErrInvalidSimpleExtention, ref.Name)
				}
			}
		}
	}
	return nil
}

// functionTypes returns every type referenced by the file's function signatures.
func (s *SimpleExtensionFile) functionTypes() []types.FuncDefArgType {
	var out []types.FuncDefArgType

	// Extension impl structs embed their shared fields in layers:
	// ScalarFunctionImpl <- AggregateFunctionImpl <- WindowFunctionImpl.
	// Collect the shared args/return fields once, then add aggregate/window
	// intermediate types through the aggregate layer. Table functions have
	// their own impl struct, whose schema plays the role of the return type.
	addSignature := func(args FuncParameterList, ret *parser.TypeExpression) {
		for _, arg := range args {
			if typ := arg.GetTypeExpression(); typ != nil {
				out = append(out, typ)
			}
		}
		if ret != nil && ret.ValueType != nil {
			out = append(out, ret.ValueType)
		}
	}

	addAggregateImpl := func(impl AggregateFunctionImpl) {
		addSignature(impl.Args, impl.Return)
		if impl.Intermediate.ValueType != nil {
			out = append(out, impl.Intermediate.ValueType)
		}
	}

	for _, fn := range s.ScalarFunctions {
		for _, impl := range fn.Impls {
			addSignature(impl.Args, impl.Return)
		}
	}
	for _, fn := range s.AggregateFunctions {
		for _, impl := range fn.Impls {
			addAggregateImpl(impl)
		}
	}
	for _, fn := range s.WindowFunctions {
		for _, impl := range fn.Impls {
			addAggregateImpl(impl.AggregateFunctionImpl)
		}
	}
	for _, fn := range s.TableFunctions {
		for _, impl := range fn.Impls {
			addSignature(impl.Args, impl.Schema)
		}
	}
	return out
}

// validateTableFunctionSchemas ensures every declared table function schema
// is a non-nullable struct.
func (s *SimpleExtensionFile) validateTableFunctionSchemas() error {
	for _, fn := range s.TableFunctions {
		for _, impl := range fn.Impls {
			if impl.Schema == nil || impl.Schema.ValueType == nil {
				continue
			}
			schema, ok := impl.Schema.ValueType.(*types.ParameterizedStructType)
			if !ok || schema.Nullability != types.NullabilityRequired {
				return fmt.Errorf("%w: table function %q schema must be a non-nullable struct, got %s",
					substraitgo.ErrInvalidSimpleExtention, fn.Name, impl.Schema.ValueType)
			}
		}
	}
	return nil
}
