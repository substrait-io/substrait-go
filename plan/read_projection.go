// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"fmt"

	substraitgo "github.com/substrait-io/substrait-go/v9"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/types"
)

// A read projection removes fields before emit remapping. Preserve the row
// struct and nested containers; maintain_singular_struct only controls the
// outer result of a general mask expression, and a read always returns a row.
func projectReadStruct(input types.StructType, selection expr.MaskStructSelect) (types.StructType, error) {
	output := input
	output.Types = make([]types.Type, len(selection))
	for i, item := range selection {
		field := item.Field()
		if field < 0 || int(field) >= len(input.Types) {
			return types.StructType{}, fmt.Errorf("%w: read projection field %d out of range for struct with %d fields", substraitgo.ErrInvalidRel, field, len(input.Types))
		}
		projected, err := projectReadType(input.Types[field], item.Child())
		if err != nil {
			return types.StructType{}, err
		}
		output.Types[i] = projected
	}
	return output, nil
}

func projectReadType(input types.Type, selection expr.MaskSelect) (types.Type, error) {
	if selection == nil {
		return input, nil
	}
	// Copy containers before replacing children so the base schema, nullability
	// and type variations remain intact, including on repeated RecordType calls.
	switch selection := selection.(type) {
	case expr.MaskStructSelect:
		if input, ok := input.(*types.StructType); ok {
			output, err := projectReadStruct(*input, selection)
			if err != nil {
				return nil, err
			}
			return &output, nil
		}
	case *expr.MaskListSelect:
		if input, ok := input.(*types.ListType); ok {
			output := *input
			var err error
			output.Type, err = projectReadType(input.Type, selection.Child())
			if err != nil {
				return nil, err
			}
			return &output, nil
		}
	case *expr.MaskMapSelect:
		if input, ok := input.(*types.MapType); ok {
			output := *input
			var err error
			output.Value, err = projectReadType(input.Value, selection.Child())
			if err != nil {
				return nil, err
			}
			return &output, nil
		}
	}
	return nil, fmt.Errorf("%w: read projection %T cannot select from %s", substraitgo.ErrInvalidRel, selection, input)
}
