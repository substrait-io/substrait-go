// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"github.com/substrait-io/substrait-go/v9/plan"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

func statsToProto(s *plan.Stats) *proto.RelCommon_Hint_Stats {
	if s == nil {
		return nil
	}
	return &proto.RelCommon_Hint_Stats{
		RowCount:          s.RowCount,
		RecordSize:        s.RecordSize,
		AdvancedExtension: advancedExtensionToProto(s.AdvancedExtension),
	}
}

func statsFromProto(s *proto.RelCommon_Hint_Stats) *plan.Stats {
	if s == nil {
		return nil
	}
	return &plan.Stats{
		RowCount:          s.RowCount,
		RecordSize:        s.RecordSize,
		AdvancedExtension: advancedExtensionFromProto(s.AdvancedExtension),
	}
}

func runtimeConstraintToProto(rc *plan.RuntimeConstraint) *proto.RelCommon_Hint_RuntimeConstraint {
	if rc == nil {
		return nil
	}
	return &proto.RelCommon_Hint_RuntimeConstraint{
		AdvancedExtension: advancedExtensionToProto(rc.AdvancedExtension),
	}
}

func runtimeConstraintFromProto(rc *proto.RelCommon_Hint_RuntimeConstraint) *plan.RuntimeConstraint {
	if rc == nil {
		return nil
	}
	return &plan.RuntimeConstraint{
		AdvancedExtension: advancedExtensionFromProto(rc.AdvancedExtension),
	}
}

func savedComputationToProto(s *plan.SavedComputation) *proto.RelCommon_Hint_SavedComputation {
	if s == nil {
		return nil
	}
	return &proto.RelCommon_Hint_SavedComputation{
		ComputationId:     s.ComputationID,
		Type:              proto.RelCommon_Hint_ComputationType(s.Type),
		AdvancedExtension: advancedExtensionToProto(s.AdvancedExtension),
	}
}

func savedComputationFromProto(s *proto.RelCommon_Hint_SavedComputation) *plan.SavedComputation {
	if s == nil {
		return nil
	}
	return &plan.SavedComputation{
		ComputationID:     s.ComputationId,
		Type:              plan.ComputationType(s.Type),
		AdvancedExtension: advancedExtensionFromProto(s.AdvancedExtension),
	}
}

func loadedComputationToProto(l *plan.LoadedComputation) *proto.RelCommon_Hint_LoadedComputation {
	if l == nil {
		return nil
	}
	return &proto.RelCommon_Hint_LoadedComputation{
		ComputationIdReference: l.ComputationIDReference,
		Type:                   proto.RelCommon_Hint_ComputationType(l.Type),
		AdvancedExtension:      advancedExtensionToProto(l.AdvancedExtension),
	}
}

func loadedComputationFromProto(l *proto.RelCommon_Hint_LoadedComputation) *plan.LoadedComputation {
	if l == nil {
		return nil
	}
	return &plan.LoadedComputation{
		ComputationIDReference: l.ComputationIdReference,
		Type:                   plan.ComputationType(l.Type),
		AdvancedExtension:      advancedExtensionFromProto(l.AdvancedExtension),
	}
}

func hintToProto(h *plan.Hint) *proto.RelCommon_Hint {
	if h == nil {
		return nil
	}
	var saved []*proto.RelCommon_Hint_SavedComputation
	if h.SavedComputations != nil {
		saved = make([]*proto.RelCommon_Hint_SavedComputation, len(h.SavedComputations))
		for i, s := range h.SavedComputations {
			saved[i] = savedComputationToProto(s)
		}
	}
	var loaded []*proto.RelCommon_Hint_LoadedComputation
	if h.LoadedComputations != nil {
		loaded = make([]*proto.RelCommon_Hint_LoadedComputation, len(h.LoadedComputations))
		for i, l := range h.LoadedComputations {
			loaded[i] = loadedComputationToProto(l)
		}
	}
	return &proto.RelCommon_Hint{
		Stats:              statsToProto(h.Stats),
		Constraint:         runtimeConstraintToProto(h.Constraint),
		Alias:              h.Alias,
		OutputNames:        h.OutputNames,
		AdvancedExtension:  advancedExtensionToProto(h.AdvancedExtension),
		SavedComputations:  saved,
		LoadedComputations: loaded,
	}
}

func hintFromProto(h *proto.RelCommon_Hint) *plan.Hint {
	if h == nil {
		return nil
	}
	var saved []*plan.SavedComputation
	if h.SavedComputations != nil {
		saved = make([]*plan.SavedComputation, len(h.SavedComputations))
		for i, s := range h.SavedComputations {
			saved[i] = savedComputationFromProto(s)
		}
	}
	var loaded []*plan.LoadedComputation
	if h.LoadedComputations != nil {
		loaded = make([]*plan.LoadedComputation, len(h.LoadedComputations))
		for i, l := range h.LoadedComputations {
			loaded[i] = loadedComputationFromProto(l)
		}
	}
	return &plan.Hint{
		Stats:              statsFromProto(h.Stats),
		Constraint:         runtimeConstraintFromProto(h.Constraint),
		Alias:              h.Alias,
		OutputNames:        h.OutputNames,
		AdvancedExtension:  advancedExtensionFromProto(h.AdvancedExtension),
		SavedComputations:  saved,
		LoadedComputations: loaded,
	}
}
