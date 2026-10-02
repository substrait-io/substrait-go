// SPDX-License-Identifier: Apache-2.0

package extensions_test

import (
	"fmt"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/substrait-io/substrait-go/v9/extensions"
)

func TestUnmarshalFunctionArgumentErrors(t *testing.T) {
	tests := []struct {
		name, argument, field string
	}{
		{"numeric name", "{name: 7, value: i64}", "name"},
		{"null name", "{name: null, value: i64}", "name"},
		{"numeric description", "{description: 7, type: i64}", "description"},
		{"null description", "{description: null, options: [A]}", "description"},
		{"string constant", `{value: i64, constant: "true"}`, "constant"},
		{"numeric constant", "{value: i64, constant: 1}", "constant"},
		{"null constant", "{value: i64, constant: null}", "constant"},
		{"scalar options", "{options: A}", "options"},
		{"mapping options", "{options: {A: 1}}", "options"},
		{"null options", "{options: null}", "options"},
		{"numeric option", "{options: [A, 7]}", "options[1]"},
		{"null option", "{options: [A, null]}", "options[1]"},
		{"nested option", "{options: [A, [B]]}", "options[1]"},
		{"null value", "{value: null}", "value"},
		{"null type", "{type: null}", "type"},
		{"numeric value", "{value: 7}", "value"},
		{"mapping type", "{type: {name: i64}}", "type"},
		{"invalid value expression", `{value: "list<"}`, "value"},
		{"invalid type expression", `{type: "list<"}`, "type"},
		{"empty argument", "{}", ""},
		{"name only", "{name: only_a_name}", ""},
		{"misspelled value", "{name: x, valu: i64}", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Put the malformed argument second to check its reported position.
			definition := fmt.Sprintf(`
urn: extension:test:argument_errors
scalar_functions:
  - name: example
    impls:
      - args:
          - value: i64
          - %s
        return: i64
`, tt.argument)
			var extension extensions.SimpleExtensionFile
			var err error
			require.NotPanics(t, func() {
				err = yaml.Unmarshal([]byte(definition), &extension)
			})
			require.Error(t, err)
			if tt.field == "" {
				assert.Contains(t, err.Error(), "args[1]: expected one of value, type or options")
			} else {
				assert.Contains(t, err.Error(), "args[1]."+tt.field)
			}
		})
	}
}

func TestUnmarshalFunctionArgumentValues(t *testing.T) {
	const definition = `
urn: extension:test:argument_values
scalar_functions:
  - name: example
    impls:
      - args:
          - {value: i64}
          - {name: input, description: input value, value: i32?, constant: true}
          - {name: "", description: "", value: string, constant: false}
          - {name: mode, description: selected mode, options: [A, B]}
          - {name: target, description: output type, type: "decimal<P,S>"}
        return: i64
`
	var extension extensions.SimpleExtensionFile
	require.NoError(t, yaml.Unmarshal([]byte(definition), &extension))
	require.Len(t, extension.ScalarFunctions, 1)
	require.Len(t, extension.ScalarFunctions[0].Impls, 1)
	args := extension.ScalarFunctions[0].Impls[0].Args
	require.Len(t, args, 5)

	for i, expected := range []struct {
		name, description, typ string
		constant               bool
	}{
		{"", "", "i64", false},
		{"input", "input value", "i32?", true},
		{"", "", "string", false},
	} {
		arg, ok := args[i].(extensions.ValueArg)
		require.True(t, ok)
		assert.Equal(t, expected.name, arg.Name)
		assert.Equal(t, expected.description, arg.Description)
		assert.Equal(t, expected.typ, arg.Value.ValueType.String())
		assert.Equal(t, expected.constant, arg.Constant)
	}
	assert.Equal(t, extensions.EnumArg{
		Name: "mode", Description: "selected mode", Options: []string{"A", "B"},
	}, args[3])
	typeArg, ok := args[4].(extensions.TypeArg)
	require.True(t, ok)
	assert.Equal(t, "target", typeArg.Name)
	assert.Equal(t, "output type", typeArg.Description)
	assert.Equal(t, "decimal<P,S>", typeArg.Type.ValueType.String())
}
