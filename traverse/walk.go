// SPDX-License-Identifier: Apache-2.0
package traverse

import (
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// walk implements the core traversal logic for Substrait relations.
func walk(root *proto.Rel, visitor Visitor) {
	// Pre-compute interface capabilities to avoid repeated type assertions
	relVisitor, canVisitRel := visitor.(RelVisitor)
	exprVisitor, canVisitExpr := visitor.(ExprVisitor)

	var visitRel func(rel *proto.Rel)
	var visitExpr func(expr *proto.Expression)
	var visitFieldRef func(ref *proto.Expression_FieldReference)
	var traverseMaskExpression func(mask *proto.Expression_MaskExpression)
	var traverseMaskSelect func(sel *proto.Expression_MaskExpression_Select)

	visitRel = func(rel *proto.Rel) {
		if rel == nil {
			return
		}

		// Call visitor if it implements RelVisitor
		if canVisitRel {
			relVisitor.VisitRel(rel)
		}

		// Traverse children based on type
		switch typed := rel.RelType.(type) {
		case *proto.Rel_Read:
			visitExpr(typed.Read.Filter)
			visitExpr(typed.Read.BestEffortFilter)
			traverseMaskExpression(typed.Read.Projection)

		case *proto.Rel_Filter:
			visitRel(typed.Filter.Input)
			visitExpr(typed.Filter.Condition)

		case *proto.Rel_Fetch:
			visitRel(typed.Fetch.Input)
			switch off := typed.Fetch.OffsetMode.(type) {
			case *proto.FetchRel_OffsetExpr:
				visitExpr(off.OffsetExpr)
			}
			switch cm := typed.Fetch.CountMode.(type) {
			case *proto.FetchRel_CountExpr:
				visitExpr(cm.CountExpr)
			}

		case *proto.Rel_Project:
			visitRel(typed.Project.Input)
			for _, expr := range typed.Project.Expressions {
				visitExpr(expr)
			}

		case *proto.Rel_Aggregate:
			visitRel(typed.Aggregate.Input)

			// Visit the grouping expressions at the relation level
			// These are referenced by index from the Grouping messages
			for _, expr := range typed.Aggregate.GroupingExpressions {
				visitExpr(expr)
			}

			// Visit expressions in each grouping set
			// Note: GroupingExpressions field is deprecated in favor of ExpressionReferences
			for _, grp := range typed.Aggregate.Groupings {
				// Still visit the deprecated field for backwards compatibility
				for _, e := range grp.GroupingExpressions {
					visitExpr(e)
				}
			}

			// Visit measure expressions
			for _, m := range typed.Aggregate.Measures {
				if m.Measure != nil {
					for _, arg := range m.Measure.Arguments {
						if arg.GetValue() != nil {
							visitExpr(arg.GetValue())
						}
					}
					for _, s := range m.Measure.Sorts {
						visitExpr(s.Expr)
					}
				}
				visitExpr(m.Filter)
			}

		case *proto.Rel_Sort:
			visitRel(typed.Sort.Input)
			for _, s := range typed.Sort.Sorts {
				visitExpr(s.Expr)
			}

		case *proto.Rel_Window:
			visitRel(typed.Window.Input)
			for _, wf := range typed.Window.WindowFunctions {
				for _, arg := range wf.Arguments {
					if arg.GetValue() != nil {
						visitExpr(arg.GetValue())
					}
				}
			}
			for _, part := range typed.Window.PartitionExpressions {
				visitExpr(part)
			}
			for _, s := range typed.Window.Sorts {
				visitExpr(s.Expr)
			}

		case *proto.Rel_Exchange:
			visitRel(typed.Exchange.Input)
			switch ek := typed.Exchange.ExchangeKind.(type) {
			case *proto.ExchangeRel_ScatterByFields:
				for _, f := range ek.ScatterByFields.Fields {
					visitFieldRef(f)
				}
			case *proto.ExchangeRel_SingleTarget:
				visitExpr(ek.SingleTarget.Expression)
			case *proto.ExchangeRel_MultiTarget:
				visitExpr(ek.MultiTarget.Expression)
			}

		case *proto.Rel_Expand:
			visitRel(typed.Expand.Input)
			for _, fld := range typed.Expand.Fields {
				switch ft := fld.FieldType.(type) {
				case *proto.ExpandRel_ExpandField_ConsistentField:
					visitExpr(ft.ConsistentField)
				case *proto.ExpandRel_ExpandField_SwitchingField:
					for _, dup := range ft.SwitchingField.Duplicates {
						visitExpr(dup)
					}
				}
			}

		case *proto.Rel_Join:
			visitRel(typed.Join.Left)
			visitRel(typed.Join.Right)
			visitExpr(typed.Join.Expression)
			visitExpr(typed.Join.PostJoinFilter)

		case *proto.Rel_HashJoin:
			visitRel(typed.HashJoin.Left)
			visitRel(typed.HashJoin.Right)
			for _, k := range typed.HashJoin.Keys {
				visitFieldRef(k.Left)
				visitFieldRef(k.Right)
			}
			for _, lk := range typed.HashJoin.LeftKeys {
				visitFieldRef(lk)
			}
			for _, rk := range typed.HashJoin.RightKeys {
				visitFieldRef(rk)
			}
			visitExpr(typed.HashJoin.PostJoinFilter)

		case *proto.Rel_MergeJoin:
			visitRel(typed.MergeJoin.Left)
			visitRel(typed.MergeJoin.Right)
			for _, k := range typed.MergeJoin.Keys {
				visitFieldRef(k.Left)
				visitFieldRef(k.Right)
			}
			for _, lk := range typed.MergeJoin.LeftKeys {
				visitFieldRef(lk)
			}
			for _, rk := range typed.MergeJoin.RightKeys {
				visitFieldRef(rk)
			}
			visitExpr(typed.MergeJoin.PostJoinFilter)

		case *proto.Rel_NestedLoopJoin:
			visitRel(typed.NestedLoopJoin.Left)
			visitRel(typed.NestedLoopJoin.Right)
			visitExpr(typed.NestedLoopJoin.Expression)

		case *proto.Rel_Cross:
			visitRel(typed.Cross.Left)
			visitRel(typed.Cross.Right)

		case *proto.Rel_Set:
			for _, in := range typed.Set.Inputs {
				visitRel(in)
			}

		case *proto.Rel_ExtensionSingle:
			visitRel(typed.ExtensionSingle.Input)

		case *proto.Rel_ExtensionMulti:
			for _, in := range typed.ExtensionMulti.Inputs {
				visitRel(in)
			}

		case *proto.Rel_Write:
			visitRel(typed.Write.Input)

		case *proto.Rel_Update:
			visitExpr(typed.Update.Condition)
			for _, tf := range typed.Update.Transformations {
				visitExpr(tf.Transformation)
			}

		case *proto.Rel_Ddl, *proto.Rel_ExtensionLeaf, *proto.Rel_Reference:
			// leaf nodes
		}
	}

	visitExpr = func(expr *proto.Expression) {
		if expr == nil {
			return
		}

		// Call visitor if it implements ExprVisitor
		if canVisitExpr {
			exprVisitor.VisitExpr(expr)
		}

		// Don't traverse into literals - they have no children
		if _, isLit := expr.RexType.(*proto.Expression_Literal_); isLit {
			return
		}

		// Traverse based on expression type
		switch typed := expr.RexType.(type) {
		case *proto.Expression_Subquery_:
			// Handle different subquery types
			switch st := typed.Subquery.SubqueryType.(type) {
			case *proto.Expression_Subquery_Scalar_:
				visitRel(st.Scalar.Input)

			case *proto.Expression_Subquery_InPredicate_:
				for _, needle := range st.InPredicate.Needles {
					visitExpr(needle)
				}
				visitRel(st.InPredicate.Haystack)

			case *proto.Expression_Subquery_SetPredicate_:
				visitRel(st.SetPredicate.Tuples)

			case *proto.Expression_Subquery_SetComparison_:
				visitExpr(st.SetComparison.Left)
				visitRel(st.SetComparison.Right)
			}

		case *proto.Expression_ScalarFunction_:
			for _, arg := range typed.ScalarFunction.Arguments {
				if arg.GetValue() != nil {
					visitExpr(arg.GetValue())
				}
			}

		case *proto.Expression_WindowFunction_:
			for _, arg := range typed.WindowFunction.Arguments {
				if arg.GetValue() != nil {
					visitExpr(arg.GetValue())
				}
			}
			for _, p := range typed.WindowFunction.Partitions {
				visitExpr(p)
			}
			for _, s := range typed.WindowFunction.Sorts {
				visitExpr(s.Expr)
			}

		case *proto.Expression_IfThen_:
			for _, ifc := range typed.IfThen.Ifs {
				visitExpr(ifc.If)
				visitExpr(ifc.Then)
			}
			visitExpr(typed.IfThen.Else)

		case *proto.Expression_SwitchExpression_:
			visitExpr(typed.SwitchExpression.Match)
			for _, iv := range typed.SwitchExpression.Ifs {
				visitExpr(iv.Then)
			}
			visitExpr(typed.SwitchExpression.Else)

		case *proto.Expression_SingularOrList_:
			visitExpr(typed.SingularOrList.Value)
			for _, o := range typed.SingularOrList.Options {
				visitExpr(o)
			}

		case *proto.Expression_MultiOrList_:
			for _, v := range typed.MultiOrList.Value {
				visitExpr(v)
			}
			for _, rec := range typed.MultiOrList.Options {
				for _, f := range rec.Fields {
					visitExpr(f)
				}
			}

		case *proto.Expression_Cast_:
			visitExpr(typed.Cast.Input)

		case *proto.Expression_Nested_:
			switch nt := typed.Nested.NestedType.(type) {
			case *proto.Expression_Nested_Struct_:
				for _, f := range nt.Struct.Fields {
					visitExpr(f)
				}
			case *proto.Expression_Nested_List_:
				for _, v := range nt.List.Values {
					visitExpr(v)
				}
			case *proto.Expression_Nested_Map_:
				for _, kv := range nt.Map.KeyValues {
					visitExpr(kv.Key)
					visitExpr(kv.Value)
				}
			}

		case *proto.Expression_Selection:
			visitFieldRef(typed.Selection)
		}
	}

	visitFieldRef = func(ref *proto.Expression_FieldReference) {
		if ref == nil {
			return
		}

		// Just traverse any expression roots in the field reference
		if rt := ref.GetRootType(); rt != nil {
			if re, ok := rt.(*proto.Expression_FieldReference_Expression); ok {
				visitExpr(re.Expression)
			}
		}
	}

	traverseMaskExpression = func(mask *proto.Expression_MaskExpression) {
		if mask == nil {
			return
		}
		if mask.Select != nil {
			for _, item := range mask.Select.StructItems {
				if item.Child != nil {
					traverseMaskSelect(item.Child)
				}
			}
		}
	}

	traverseMaskSelect = func(sel *proto.Expression_MaskExpression_Select) {
		if sel == nil {
			return
		}

		switch typed := sel.Type.(type) {
		case *proto.Expression_MaskExpression_Select_Struct:
			if typed.Struct != nil {
				for _, item := range typed.Struct.StructItems {
					if item.Child != nil {
						traverseMaskSelect(item.Child)
					}
				}
			}
		case *proto.Expression_MaskExpression_Select_List:
			if typed.List != nil && typed.List.Child != nil {
				traverseMaskSelect(typed.List.Child)
			}
		case *proto.Expression_MaskExpression_Select_Map:
			if typed.Map != nil && typed.Map.Child != nil {
				traverseMaskSelect(typed.Map.Child)
			}
		}
	}

	// Start traversal
	visitRel(root)
}
