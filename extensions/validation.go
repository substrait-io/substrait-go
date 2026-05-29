// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types"
)

// validateUserDefinedTypeReferences ensures every user-defined type used in a
// function signature is declared in that same file's types section. It does
// not check type structures or type variation parents.
func (s SimpleExtensionFile) validateUserDefinedTypeReferences() error {
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
func (s SimpleExtensionFile) functionTypes() []types.FuncDefArgType {
	var out []types.FuncDefArgType

	addImpl := func(impl ScalarFunctionImpl) {
		for _, arg := range impl.Args {
			if typ := arg.GetTypeExpression(); typ != nil {
				out = append(out, typ)
			}
		}
		if impl.Return != nil && impl.Return.ValueType != nil {
			out = append(out, impl.Return.ValueType)
		}
	}

	for _, fn := range s.ScalarFunctions {
		for _, impl := range fn.Impls {
			addImpl(impl)
		}
	}
	for _, fn := range s.AggregateFunctions {
		for _, impl := range fn.Impls {
			addImpl(impl.ScalarFunctionImpl)
			if impl.Intermediate.ValueType != nil {
				out = append(out, impl.Intermediate.ValueType)
			}
		}
	}
	for _, fn := range s.WindowFunctions {
		for _, impl := range fn.Impls {
			addImpl(impl.ScalarFunctionImpl)
			if impl.Intermediate.ValueType != nil {
				out = append(out, impl.Intermediate.ValueType)
			}
		}
	}
	return out
}
