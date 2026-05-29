// SPDX-License-Identifier: Apache-2.0

package types

// ReferencedUserDefinedTypes returns the names of every user-defined type
// referenced by t, including those nested inside lists, maps, structs,
// function types, and user-defined type parameters.
func ReferencedUserDefinedTypes(t FuncDefArgType) []string {
	var names []string
	collectUserDefinedTypes(t, &names)
	return names
}

func collectUserDefinedTypes(t FuncDefArgType, names *[]string) {
	switch t := t.(type) {
	case *ParameterizedUserDefinedType:
		*names = append(*names, t.Name)
		for _, param := range t.TypeParameters {
			if dataParam, ok := param.(*DataTypeUDTParam); ok {
				collectUserDefinedTypes(dataParam.Type, names)
			}
		}
	case *ParameterizedListType:
		collectUserDefinedTypes(t.Type, names)
	case *ParameterizedMapType:
		collectUserDefinedTypes(t.Key, names)
		collectUserDefinedTypes(t.Value, names)
	case *ParameterizedStructType:
		for _, field := range t.Types {
			collectUserDefinedTypes(field, names)
		}
	case *ParameterizedFuncType:
		for _, param := range t.Parameters {
			collectUserDefinedTypes(param, names)
		}
		collectUserDefinedTypes(t.Return, names)
	}
}
