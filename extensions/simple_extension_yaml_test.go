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
			assert.Contains(t, err.Error(), "args[1]."+tt.field)
		})
	}
}

func TestUnmarshalFunctionArgumentMissingDefinition(t *testing.T) {
	tests := []struct {
		name, argument string
	}{
		{"empty argument", "{}"},
		{"name only", "{name: only_a_name}"},
		{"misspelled value", "{name: x, valu: i64}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			assert.Contains(t, err.Error(), "args[1]: expected one of options, value or type")
		})
	}
}

func TestUnmarshalFunctionValueArgument(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		arg, ok := unmarshalFunctionArgument(t, "{value: i64}").(extensions.ValueArg)
		require.True(t, ok)
		assert.Empty(t, arg.Name)
		assert.Empty(t, arg.Description)
		assert.Equal(t, "i64", arg.Value.ValueType.String())
		assert.False(t, arg.Constant)
	})

	t.Run("constant argument", func(t *testing.T) {
		arg, ok := unmarshalFunctionArgument(t,
			"{name: input, description: input value, value: i32?, constant: true}",
		).(extensions.ValueArg)
		require.True(t, ok)
		assert.Equal(t, "input", arg.Name)
		assert.Equal(t, "input value", arg.Description)
		assert.Equal(t, "i32?", arg.Value.ValueType.String())
		assert.True(t, arg.Constant)
	})
}

func TestUnmarshalFunctionEnumArgument(t *testing.T) {
	arg := unmarshalFunctionArgument(t, "{name: mode, description: selected mode, options: [A, B]}")

	assert.Equal(t, extensions.EnumArg{
		Name: "mode", Description: "selected mode", Options: []string{"A", "B"},
	}, arg)
}

func TestUnmarshalFunctionTypeArgument(t *testing.T) {
	arg, ok := unmarshalFunctionArgument(t, `{name: target, description: output type, type: "decimal<P,S>"}`).(extensions.TypeArg)

	require.True(t, ok)
	assert.Equal(t, "target", arg.Name)
	assert.Equal(t, "output type", arg.Description)
	assert.Equal(t, "decimal<P,S>", arg.Type.ValueType.String())
}

func unmarshalFunctionArgument(t *testing.T, argument string) extensions.FuncParameter {
	t.Helper()
	definition := fmt.Sprintf(`
urn: extension:test:argument_values
scalar_functions:
  - name: example
    impls:
      - args:
          - %s
        return: i64
`, argument)
	var extension extensions.SimpleExtensionFile
	require.NoError(t, yaml.Unmarshal([]byte(definition), &extension))
	require.Len(t, extension.ScalarFunctions, 1)
	require.Len(t, extension.ScalarFunctions[0].Impls, 1)
	args := extension.ScalarFunctions[0].Impls[0].Args
	require.Len(t, args, 1)
	return args[0]
}
