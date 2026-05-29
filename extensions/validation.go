// SPDX-License-Identifier: Apache-2.0

package extensions

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v8"
	"github.com/substrait-io/substrait-go/v8/types/parser"
)

// validateUserDefinedTypeReferences ensures every user-defined type used by a
// function in the file is declared in that same file's types section.
func (s SimpleExtensionFile) validateUserDefinedTypeReferences() error {
	declared := make(map[string]struct{}, len(s.Types))
	for _, t := range s.Types {
		declared[t.Name] = struct{}{}
	}

	for _, expr := range s.functionTypeExpressions() {
		for _, name := range expr.UserDefinedTypes {
			if _, ok := declared[name]; !ok {
				return fmt.Errorf("%w: user-defined type %q is not declared",
					substraitgo.ErrInvalidSimpleExtention, name)
			}
		}
	}
	return nil
}

// functionTypeExpressions returns every type expression referenced by the
// file's function signatures.
func (s SimpleExtensionFile) functionTypeExpressions() []parser.TypeExpression {
	var out []parser.TypeExpression

	addImpl := func(impl ScalarFunctionImpl) {
		for _, arg := range impl.Args {
			switch a := arg.(type) {
			case ValueArg:
				if a.Value != nil {
					out = append(out, *a.Value)
				}
			case TypeArg:
				if a.Type != nil {
					out = append(out, *a.Type)
				}
			}
		}
		if impl.Return != nil {
			out = append(out, *impl.Return)
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
			out = append(out, impl.Intermediate)
		}
	}
	for _, fn := range s.WindowFunctions {
		for _, impl := range fn.Impls {
			addImpl(impl.ScalarFunctionImpl)
			out = append(out, impl.Intermediate)
		}
	}
	return out
}
