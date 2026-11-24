// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"slices"

	"github.com/substrait-io/substrait-go/v9/types"
)

func setRecordTypesMatch(left, right types.RecordType) bool {
	if left.FieldCount() != right.FieldCount() {
		return false
	}
	for i, field := range left.Types() {
		if !field.WithNullability(types.NullabilityUnspecified).Equals(
			right.Types()[i].WithNullability(types.NullabilityUnspecified)) {
			return false
		}
	}
	return true
}

func (s *SetRel) directOutputSchema() types.RecordType {
	primary := s.inputs[0].RecordType()
	switch s.op {
	case SetOpUnionAll, SetOpUnionDistinct, SetOpIntersectionPrimary,
		SetOpIntersectionMultiset, SetOpIntersectionMultisetAll:
	default:
		return primary
	}

	inputs := make([]types.RecordType, len(s.inputs))
	inputs[0] = primary
	for i, input := range s.inputs[1:] {
		inputs[i+1] = input.RecordType()
	}
	fields := slices.Clone(primary.Types())
	for column, field := range fields {
		nullable := s.op == SetOpIntersectionMultiset || s.op == SetOpIntersectionMultisetAll
		for i, input := range inputs {
			inputNullable := input.Types()[column].GetNullability() == types.NullabilityNullable
			switch s.op {
			case SetOpUnionAll, SetOpUnionDistinct:
				nullable = nullable || inputNullable
			case SetOpIntersectionMultiset, SetOpIntersectionMultisetAll:
				nullable = nullable && inputNullable
			case SetOpIntersectionPrimary:
				if i > 0 {
					nullable = nullable || inputNullable
				}
			}
		}
		if s.op == SetOpIntersectionPrimary {
			nullable = nullable && field.GetNullability() == types.NullabilityNullable
		}
		if nullable != (field.GetNullability() == types.NullabilityNullable) {
			n := types.NullabilityRequired
			if nullable {
				n = types.NullabilityNullable
			}
			fields[column] = field.WithNullability(n)
		}
	}
	return *types.NewRecordTypeFromTypes(fields)
}
