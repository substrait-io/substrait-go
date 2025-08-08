package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	. "github.com/substrait-io/substrait-go/v5/types"
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
