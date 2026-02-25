package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	. "github.com/substrait-io/substrait-go/v7/types"
)

func TestAnyType(t *testing.T) {
	decP30S9 := &DecimalType{Precision: 30, Scale: 9, Nullability: NullabilityRequired}
	varchar37 := &VarCharType{Length: 37}
	for _, td := range []struct {
		testName           string
		argName            string
		parameters         []FuncDefArgType
		args               []Type
		concreteReturnType Type
		nullability        Nullability
		expectedString     string
		expectedErr        string
	}{
		{
			testName:           "any",
			argName:            "any",
			parameters:         []FuncDefArgType{&AnyType{Name: "any"}},
			args:               []Type{decP30S9},
			concreteReturnType: decP30S9,
			nullability:        NullabilityNullable,
			expectedString:     "any?",
		},
		{
			testName:           "anyrequired",
			argName:            "any2",
			parameters:         []FuncDefArgType{&Int16Type{}, &AnyType{Name: "any2"}},
			args:               []Type{&Int16Type{}, &Int64Type{}},
			concreteReturnType: &Int64Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any2",
		},
		{
			testName: "list<any1>",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&Int16Type{},
				&ParameterizedListType{Type: &AnyType{Name: "any1"}},
			},
			args:               []Type{&Int16Type{}, &ListType{Type: &Int64Type{}}},
			concreteReturnType: &Int64Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any1",
		},
		{
			testName: "wrong_list<any1>",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&Int16Type{},
				&ParameterizedListType{Type: &AnyType{Name: "any1"}},
			},
			args:           []Type{&Int16Type{}, &ListType{}},
			nullability:    NullabilityRequired,
			expectedString: "any1",
			expectedErr:    "expected ListType to have non-nil 1 parameter, found [<nil>]",
		},
		{
			testName: "map<string, any1>",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&Int16Type{},
				&ParameterizedMapType{Key: &StringType{}, Value: &AnyType{Name: "any1"}},
			},
			args: []Type{
				&Int16Type{},
				&MapType{Key: &StringType{}, Value: &ListType{Type: &Int64Type{}}},
			},
			concreteReturnType: &ListType{Type: &Int64Type{}},
			nullability:        NullabilityRequired,
			expectedString:     "any1",
		},
		{
			testName: "map<any2, string>",
			argName:  "any2",
			parameters: []FuncDefArgType{
				&Int16Type{},
				&ParameterizedMapType{Key: &AnyType{Name: "any2"}, Value: &StringType{}},
			},
			args: []Type{
				&Int16Type{},
				&MapType{Key: &ListType{Type: &Int64Type{}}, Value: &StringType{}},
			},
			concreteReturnType: &ListType{Type: &Int64Type{}},
			nullability:        NullabilityRequired,
			expectedString:     "any2",
		},
		{
			testName: "wrong_map<any2, string>",
			argName:  "any2",
			parameters: []FuncDefArgType{
				&Int16Type{},
				&ParameterizedMapType{Key: &AnyType{Name: "any2"}, Value: &StringType{}},
			},
			args: []Type{
				&Int16Type{},
				&MapType{Value: &StringType{}},
			},
			nullability:    NullabilityRequired,
			expectedString: "any2",
			expectedErr:    "expected MapType to have 2 non-nil parameters, found [<nil> string]",
		},
		{
			testName: "struct<string, any1, i64>",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedStructType{Types: []FuncDefArgType{
					&StringType{}, &AnyType{Name: "any1"}, &Int64Type{},
				}},
			},
			args: []Type{
				&StructType{Types: []Type{&StringType{}, &VarCharType{Length: 37}, &Int64Type{}}},
			},
			concreteReturnType: &VarCharType{Length: 37},
			nullability:        NullabilityRequired,
			expectedString:     "any1",
		},
		{
			testName: "wrong_struct<string, any1, i64>",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedStructType{Types: []FuncDefArgType{
					&StringType{}, &AnyType{Name: "any1"}, &Int64Type{},
				}},
			},
			args: []Type{
				&StructType{Types: []Type{&StringType{}, &Int64Type{}, nil, &Int64Type{}}},
			},
			expectedString: "any1",
			expectedErr:    "expected StructType to have 3 non-nil parameters, found [string i64 <nil> i64]",
		},
		{
			testName: "func<any1 -> any2> resolve any2 from return",
			argName:  "any2",
			parameters: []FuncDefArgType{
				&ParameterizedListType{Type: &AnyType{Name: "any1"}},
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{&AnyType{Name: "any1"}},
					Return:     &AnyType{Name: "any2"},
				},
			},
			args: []Type{
				&ListType{Type: &Int64Type{}},
				&FuncType{ParameterTypes: []Type{&Int64Type{}}, ReturnType: &StringType{}},
			},
			concreteReturnType: &StringType{},
			nullability:        NullabilityRequired,
			expectedString:     "any2",
		},
		{
			testName: "func<any1 -> any1> resolve any1 from param",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{&AnyType{Name: "any1"}},
					Return:     &AnyType{Name: "any1"},
				},
			},
			args: []Type{
				&FuncType{ParameterTypes: []Type{&Int32Type{}}, ReturnType: &Int32Type{}},
			},
			concreteReturnType: &Int32Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any1",
		},
		{
			testName: "func<any1, any2 -> any1> resolve any2 from second param",
			argName:  "any2",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{&AnyType{Name: "any1"}, &AnyType{Name: "any2"}},
					Return:     &AnyType{Name: "any1"},
				},
			},
			args: []Type{
				&FuncType{ParameterTypes: []Type{&Int64Type{}, &Float64Type{}}, ReturnType: &Int64Type{}},
			},
			concreteReturnType: &Float64Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any2",
		},
		{
			// Int64Type.GetParameters() returns nil (no type components),
			// but ParameterizedFuncType expects 2 (1 param + 1 return)
			testName: "wrong_func_not_func_type",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{&AnyType{Name: "any1"}},
					Return:     &AnyType{Name: "any1"},
				},
			},
			args:           []Type{&Int64Type{}},
			nullability:    NullabilityRequired,
			expectedString: "any1",
			expectedErr:    "expected FuncType to have 2 non-nil parameters, found []",
		},
		{
			// FuncType with 1 param + 1 return = 2 components,
			// but ParameterizedFuncType has 2 params + 1 return = 3
			testName: "wrong_func_param_count_mismatch",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{&AnyType{Name: "any1"}, &AnyType{Name: "any1"}},
					Return:     &AnyType{Name: "any1"},
				},
			},
			args: []Type{
				&FuncType{ParameterTypes: []Type{&Int64Type{}}, ReturnType: &Int64Type{}},
			},
			nullability:    NullabilityRequired,
			expectedString: "any1",
			expectedErr:    "expected FuncType to have 3 non-nil parameters, found [i64 i64]",
		},
		{
			// Nested parameterized type inside func: func<list<any1> -> any1>
			// Tests that unwrapAnyTypeWithName recurses into the func's parameter
			// types when they themselves are parameterized (list<any1>)
			testName: "func<list<any1> -> any1> resolve any1 from nested list param",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{
						&ParameterizedListType{Type: &AnyType{Name: "any1"}},
					},
					Return: &AnyType{Name: "any1"},
				},
			},
			args: []Type{
				&FuncType{
					ParameterTypes: []Type{&ListType{Type: &StringType{}}},
					ReturnType:     &StringType{},
				},
			},
			concreteReturnType: &StringType{},
			nullability:        NullabilityRequired,
			expectedString:     "any1",
		},
		{
			// Nested parameterized type in func return: func<any1 -> list<any2>>
			// The return type wraps any2 in a list, so unwrapAnyTypeWithName must
			// recurse through the func's return type, then into the list
			testName: "func<any1 -> list<any2>> resolve any2 from nested list return",
			argName:  "any2",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{&AnyType{Name: "any1"}},
					Return:     &ParameterizedListType{Type: &AnyType{Name: "any2"}},
				},
			},
			args: []Type{
				&FuncType{
					ParameterTypes: []Type{&Int64Type{}},
					ReturnType:     &ListType{Type: &Float64Type{}},
				},
			},
			concreteReturnType: &Float64Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any2",
		},
		{
			// Deeply nested: func<map<any1, any2> -> any2>
			// Tests recursion through func param → map value → any2
			testName: "func<map<any1, any2> -> any2> resolve any2 from nested map param",
			argName:  "any2",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{
						&ParameterizedMapType{
							Key:   &AnyType{Name: "any1"},
							Value: &AnyType{Name: "any2"},
						},
					},
					Return: &AnyType{Name: "any2"},
				},
			},
			args: []Type{
				&FuncType{
					ParameterTypes: []Type{
						&MapType{Key: &StringType{}, Value: &Int32Type{}},
					},
					ReturnType: &Int32Type{},
				},
			},
			concreteReturnType: &Int32Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any2",
		},
		{
			// Func with struct param: func<struct<any1, i64> -> any1>
			testName: "func<struct<any1, i64> -> any1> resolve any1 from nested struct param",
			argName:  "any1",
			parameters: []FuncDefArgType{
				&ParameterizedFuncType{
					Parameters: []FuncDefArgType{
						&ParameterizedStructType{
							Types: []FuncDefArgType{&AnyType{Name: "any1"}, &Int64Type{}},
						},
					},
					Return: &AnyType{Name: "any1"},
				},
			},
			args: []Type{
				&FuncType{
					ParameterTypes: []Type{
						&StructType{Types: []Type{&Float64Type{}, &Int64Type{}}},
					},
					ReturnType: &Float64Type{},
				},
			},
			concreteReturnType: &Float64Type{},
			nullability:        NullabilityRequired,
			expectedString:     "any1",
		},
		{
			testName:           "anyOtherName",
			argName:            "any1",
			parameters:         []FuncDefArgType{&AnyType{Name: "any1"}, &Int32Type{}},
			args:               []Type{varchar37, &Int32Type{}},
			concreteReturnType: varchar37,
			nullability:        NullabilityNullable,
			expectedString:     "any1?",
		},
		{
			testName:       "T name",
			argName:        "T",
			parameters:     []FuncDefArgType{&AnyType{Name: "U"}},
			args:           []Type{varchar37},
			nullability:    NullabilityNullable,
			expectedString: "T?",
		},
	} {
		t.Run(td.testName, func(t *testing.T) {
			anyBase := &AnyType{
				Name:        td.argName,
				Nullability: td.nullability,
			}
			anyType := anyBase.SetNullability(td.nullability)
			require.Equal(t, td.nullability, anyType.GetNullability())
			require.Equal(t, "any", anyType.ShortString())
			require.Equal(t, td.expectedString, anyType.String())
			returnType, err := anyType.ReturnType(td.parameters, td.args)
			if td.concreteReturnType != nil {
				require.NoError(t, err)
				require.Equal(t, td.concreteReturnType, returnType)
			} else {
				require.Error(t, err)
				if td.expectedErr != "" {
					require.Equal(t, td.expectedErr, err.Error())
				}
			}
		})
	}
}
