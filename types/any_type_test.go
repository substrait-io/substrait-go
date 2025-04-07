package types_test

import (
	"testing"
	
	"github.com/stretchr/testify/require"
	. "github.com/substrait-io/substrait-go/v3/types"
)

func TestAnyType(t *testing.T) {
	decP30S9 := &DecimalType{Precision: 30, Scale: 9, Nullability: NullabilityRequired}
	varchar37 := &VarCharType{Length: 37}
	for _, td := range []struct {
		testName           string
		argName            string
		parameters         []FuncDefArgType
		args               []Type
		expectedReturnType Type
		nullability        Nullability
		expectedString     string
	}{
		{"any", "any", []FuncDefArgType{&AnyType{Name: "any"}}, []Type{decP30S9}, decP30S9, NullabilityNullable, "any?"},
		{"List<any>", "any", []FuncDefArgType{&AnyType{Name: "any"}}, []Type{&ListType{Type: decP30S9}}, &ListType{Type: decP30S9}, NullabilityNullable, "any?"},
		{"anyrequired", "any2", []FuncDefArgType{&Int16Type{}, &AnyType{Name: "any2"}}, []Type{&Int16Type{}, &Int64Type{}}, &Int64Type{}, NullabilityRequired, "any2"},
		{"anyOtherName", "any1", []FuncDefArgType{&AnyType{Name: "any1"}, &Int32Type{}}, []Type{varchar37, &Int32Type{}}, varchar37, NullabilityNullable, "any1?"},
		{"T name", "T", []FuncDefArgType{&AnyType{Name: "U"}}, []Type{varchar37}, nil, NullabilityNullable, "T?"},
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
			if td.expectedReturnType != nil {
				require.NoError(t, err)
				require.Equal(t, td.expectedReturnType, returnType)
			} else {
				require.Error(t, err)
			}
		})
	}
}
