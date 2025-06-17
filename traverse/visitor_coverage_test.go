// SPDX-License-Identifier: Apache-2.0
package traverse

import (
	"testing"

	"github.com/stretchr/testify/require"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
)

// TestAllRelationTypes ensures every relation type switch case is covered
func TestAllRelationTypes(t *testing.T) {

	tests := []struct {
		name          string
		rel           *proto.Rel
		expectedCount int
	}{
		{
			name: "Rel_Project",
			rel: &proto.Rel{
				RelType: &proto.Rel_Project{
					Project: &proto.ProjectRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					},
				},
			},
			expectedCount: 2, // Project + ExtensionLeaf
		},
		{
			name: "Rel_Fetch_StaticCount",
			rel: &proto.Rel{
				RelType: &proto.Rel_Fetch{
					Fetch: &proto.FetchRel{
						Input:      &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						CountMode:  &proto.FetchRel_Count{Count: 10},
						OffsetMode: &proto.FetchRel_Offset{Offset: 0},
					},
				},
			},
			expectedCount: 2, // Fetch + ExtensionLeaf
		},
		{
			name: "Rel_Fetch_ExpressionCount",
			rel: &proto.Rel{
				RelType: &proto.Rel_Fetch{
					Fetch: &proto.FetchRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						CountMode: &proto.FetchRel_CountExpr{
							CountExpr: &proto.Expression{
								RexType: &proto.Expression_Literal_{},
							},
						},
						OffsetMode: &proto.FetchRel_OffsetExpr{
							OffsetExpr: &proto.Expression{
								RexType: &proto.Expression_Literal_{},
							},
						},
					},
				},
			},
			expectedCount: 2, // Fetch + ExtensionLeaf
		},
		{
			name: "Rel_Filter",
			rel: &proto.Rel{
				RelType: &proto.Rel_Filter{
					Filter: &proto.FilterRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Condition: &proto.Expression{
							RexType: &proto.Expression_Literal_{},
						},
					},
				},
			},
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Rel_Aggregate",
			rel: &proto.Rel{
				RelType: &proto.Rel_Aggregate{
					Aggregate: &proto.AggregateRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Groupings: []*proto.AggregateRel_Grouping{
							{
								GroupingExpressions: []*proto.Expression{
									{RexType: &proto.Expression_Literal_{}},
								},
							},
						},
						Measures: []*proto.AggregateRel_Measure{
							{
								Measure: &proto.AggregateFunction{
									Arguments: []*proto.FunctionArgument{
										{
											ArgType: &proto.FunctionArgument_Value{
												Value: &proto.Expression{
													RexType: &proto.Expression_Literal_{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Aggregate + ExtensionLeaf
		},
		{
			name: "Rel_Sort",
			rel: &proto.Rel{
				RelType: &proto.Rel_Sort{
					Sort: &proto.SortRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Sorts: []*proto.SortField{
							{
								Expr: &proto.Expression{
									RexType: &proto.Expression_Literal_{},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Sort + ExtensionLeaf
		},
		{
			name: "Rel_ExtensionSingle",
			rel: &proto.Rel{
				RelType: &proto.Rel_ExtensionSingle{
					ExtensionSingle: &proto.ExtensionSingleRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					},
				},
			},
			expectedCount: 2, // ExtensionSingle + ExtensionLeaf
		},
		{
			name: "Rel_Window",
			rel: &proto.Rel{
				RelType: &proto.Rel_Window{
					Window: &proto.ConsistentPartitionWindowRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						WindowFunctions: []*proto.ConsistentPartitionWindowRel_WindowRelFunction{
							{
								Arguments: []*proto.FunctionArgument{
									{
										ArgType: &proto.FunctionArgument_Value{
											Value: &proto.Expression{
												RexType: &proto.Expression_Literal_{},
											},
										},
									},
								},
							},
						},
						PartitionExpressions: []*proto.Expression{
							{RexType: &proto.Expression_Literal_{}},
						},
						Sorts: []*proto.SortField{
							{
								Expr: &proto.Expression{
									RexType: &proto.Expression_Literal_{},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Window + ExtensionLeaf
		},
		{
			name: "Rel_Exchange_ScatterByFields",
			rel: &proto.Rel{
				RelType: &proto.Rel_Exchange{
					Exchange: &proto.ExchangeRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						ExchangeKind: &proto.ExchangeRel_ScatterByFields{
							ScatterByFields: &proto.ExchangeRel_ScatterFields{
								Fields: []*proto.Expression_FieldReference{
									{
										ReferenceType: &proto.Expression_FieldReference_DirectReference{},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Exchange + ExtensionLeaf
		},
		{
			name: "Rel_Exchange_SingleTarget",
			rel: &proto.Rel{
				RelType: &proto.Rel_Exchange{
					Exchange: &proto.ExchangeRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						ExchangeKind: &proto.ExchangeRel_SingleTarget{
							SingleTarget: &proto.ExchangeRel_SingleBucketExpression{
								Expression: &proto.Expression{
									RexType: &proto.Expression_Literal_{},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Exchange + ExtensionLeaf
		},
		{
			name: "Rel_Exchange_MultiTarget",
			rel: &proto.Rel{
				RelType: &proto.Rel_Exchange{
					Exchange: &proto.ExchangeRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						ExchangeKind: &proto.ExchangeRel_MultiTarget{
							MultiTarget: &proto.ExchangeRel_MultiBucketExpression{
								Expression: &proto.Expression{
									RexType: &proto.Expression_Literal_{},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Exchange + ExtensionLeaf
		},
		{
			name: "Rel_Exchange_RoundRobin",
			rel: &proto.Rel{
				RelType: &proto.Rel_Exchange{
					Exchange: &proto.ExchangeRel{
						Input:        &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						ExchangeKind: &proto.ExchangeRel_RoundRobin_{},
					},
				},
			},
			expectedCount: 2, // Exchange + ExtensionLeaf
		},
		{
			name: "Rel_Exchange_Broadcast",
			rel: &proto.Rel{
				RelType: &proto.Rel_Exchange{
					Exchange: &proto.ExchangeRel{
						Input:        &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						ExchangeKind: &proto.ExchangeRel_Broadcast_{},
					},
				},
			},
			expectedCount: 2, // Exchange + ExtensionLeaf
		},
		{
			name: "Rel_Expand_ConsistentField",
			rel: &proto.Rel{
				RelType: &proto.Rel_Expand{
					Expand: &proto.ExpandRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Fields: []*proto.ExpandRel_ExpandField{
							{
								FieldType: &proto.ExpandRel_ExpandField_ConsistentField{
									ConsistentField: &proto.Expression{
										RexType: &proto.Expression_Literal_{},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Expand + ExtensionLeaf
		},
		{
			name: "Rel_Expand_SwitchingField",
			rel: &proto.Rel{
				RelType: &proto.Rel_Expand{
					Expand: &proto.ExpandRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Fields: []*proto.ExpandRel_ExpandField{
							{
								FieldType: &proto.ExpandRel_ExpandField_SwitchingField{
									SwitchingField: &proto.ExpandRel_SwitchingField{
										Duplicates: []*proto.Expression{
											{RexType: &proto.Expression_Literal_{}},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 2, // Expand + ExtensionLeaf
		},
		{
			name: "Rel_Write",
			rel: &proto.Rel{
				RelType: &proto.Rel_Write{
					Write: &proto.WriteRel{
						Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					},
				},
			},
			expectedCount: 2, // Write + ExtensionLeaf
		},
		{
			name: "Rel_Update",
			rel: &proto.Rel{
				RelType: &proto.Rel_Update{
					Update: &proto.UpdateRel{
						Condition: &proto.Expression{
							RexType: &proto.Expression_Literal_{},
						},
					},
				},
			},
			expectedCount: 1, // Update only (no input relation)
		},
		{
			name: "Rel_Join",
			rel: &proto.Rel{
				RelType: &proto.Rel_Join{
					Join: &proto.JoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Expression: &proto.Expression{
							RexType: &proto.Expression_Literal_{},
						},
					},
				},
			},
			expectedCount: 3, // Join + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "Rel_HashJoin",
			rel: &proto.Rel{
				RelType: &proto.Rel_HashJoin{
					HashJoin: &proto.HashJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					},
				},
			},
			expectedCount: 3, // HashJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "Rel_MergeJoin",
			rel: &proto.Rel{
				RelType: &proto.Rel_MergeJoin{
					MergeJoin: &proto.MergeJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					},
				},
			},
			expectedCount: 3, // MergeJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "Rel_NestedLoopJoin",
			rel: &proto.Rel{
				RelType: &proto.Rel_NestedLoopJoin{
					NestedLoopJoin: &proto.NestedLoopJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Expression: &proto.Expression{
							RexType: &proto.Expression_Literal_{},
						},
					},
				},
			},
			expectedCount: 3, // NestedLoopJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "Rel_Cross",
			rel: &proto.Rel{
				RelType: &proto.Rel_Cross{
					Cross: &proto.CrossRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					},
				},
			},
			expectedCount: 3, // Cross + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "Rel_Set",
			rel: &proto.Rel{
				RelType: &proto.Rel_Set{
					Set: &proto.SetRel{
						Inputs: []*proto.Rel{
							{RelType: &proto.Rel_ExtensionLeaf{}},
							{RelType: &proto.Rel_ExtensionLeaf{}},
						},
					},
				},
			},
			expectedCount: 3, // Set + 2 ExtensionLeaf inputs
		},
		{
			name: "Rel_ExtensionMulti",
			rel: &proto.Rel{
				RelType: &proto.Rel_ExtensionMulti{
					ExtensionMulti: &proto.ExtensionMultiRel{
						Inputs: []*proto.Rel{
							{RelType: &proto.Rel_ExtensionLeaf{}},
						},
					},
				},
			},
			expectedCount: 2, // ExtensionMulti + ExtensionLeaf
		},
		{
			name: "Rel_Read",
			rel: &proto.Rel{
				RelType: &proto.Rel_Read{
					Read: &proto.ReadRel{
						Filter: &proto.Expression{
							RexType: &proto.Expression_Literal_{},
						},
						BestEffortFilter: &proto.Expression{
							RexType: &proto.Expression_Literal_{},
						},
						Projection: &proto.Expression_MaskExpression{
							Select: &proto.Expression_MaskExpression_StructSelect{
								StructItems: []*proto.Expression_MaskExpression_StructItem{
									{Field: 1},
								},
							},
						},
					},
				},
			},
			expectedCount: 1, // Read only (no input relations)
		},
		{
			name: "Rel_ExtensionLeaf",
			rel: &proto.Rel{
				RelType: &proto.Rel_ExtensionLeaf{
					ExtensionLeaf: &proto.ExtensionLeafRel{},
				},
			},
			expectedCount: 1, // ExtensionLeaf only
		},
		{
			name: "Rel_Reference",
			rel: &proto.Rel{
				RelType: &proto.Rel_Reference{
					Reference: &proto.ReferenceRel{},
				},
			},
			expectedCount: 1, // Reference only
		},
		{
			name: "Rel_Ddl",
			rel: &proto.Rel{
				RelType: &proto.Rel_Ddl{
					Ddl: &proto.DdlRel{},
				},
			},
			expectedCount: 1, // Ddl only
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
				return count + 1
			})

			Walk(tt.rel, counter)

			// Should visit the exact expected number of nodes
			require.Equal(t, tt.expectedCount, counter.Result(), "Should visit exactly %d nodes", tt.expectedCount)
		})
	}
}

// TestAllExpressionTypes ensures every expression type switch case is covered
func TestAllExpressionTypes(t *testing.T) {
	// Helper to create a simple rel with an expression
	createRelWithExpr := func(expr *proto.Expression) *proto.Rel {
		return &proto.Rel{
			RelType: &proto.Rel_Filter{
				Filter: &proto.FilterRel{
					Input:     &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
					Condition: expr,
				},
			},
		}
	}

	tests := []struct {
		name          string
		rel           *proto.Rel
		expectedCount int
	}{
		{
			name: "Expression_Subquery_Scalar",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Subquery_{
					Subquery: &proto.Expression_Subquery{
						SubqueryType: &proto.Expression_Subquery_Scalar_{
							Scalar: &proto.Expression_Subquery_Scalar{
								Input: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
							},
						},
					},
				},
			}),
			expectedCount: 3, // Filter + ExtensionLeaf + Subquery ExtensionLeaf
		},
		{
			name: "Expression_Subquery_InPredicate",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Subquery_{
					Subquery: &proto.Expression_Subquery{
						SubqueryType: &proto.Expression_Subquery_InPredicate_{
							InPredicate: &proto.Expression_Subquery_InPredicate{
								Needles: []*proto.Expression{
									{RexType: &proto.Expression_Literal_{}},
								},
								Haystack: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
							},
						},
					},
				},
			}),
			expectedCount: 3, // Filter + ExtensionLeaf + Subquery ExtensionLeaf
		},
		{
			name: "Expression_Subquery_SetPredicate",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Subquery_{
					Subquery: &proto.Expression_Subquery{
						SubqueryType: &proto.Expression_Subquery_SetPredicate_{
							SetPredicate: &proto.Expression_Subquery_SetPredicate{
								Tuples: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
							},
						},
					},
				},
			}),
			expectedCount: 3, // Filter + ExtensionLeaf + Subquery ExtensionLeaf
		},
		{
			name: "Expression_Subquery_SetComparison",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Subquery_{
					Subquery: &proto.Expression_Subquery{
						SubqueryType: &proto.Expression_Subquery_SetComparison_{
							SetComparison: &proto.Expression_Subquery_SetComparison{
								Left:  &proto.Expression{RexType: &proto.Expression_Literal_{}},
								Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
							},
						},
					},
				},
			}),
			expectedCount: 3, // Filter + ExtensionLeaf + Subquery ExtensionLeaf
		},
		{
			name: "Expression_ScalarFunction",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_ScalarFunction_{
					ScalarFunction: &proto.Expression_ScalarFunction{
						Arguments: []*proto.FunctionArgument{
							{
								ArgType: &proto.FunctionArgument_Value{
									Value: &proto.Expression{RexType: &proto.Expression_Literal_{}},
								},
							},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_WindowFunction",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_WindowFunction_{
					WindowFunction: &proto.Expression_WindowFunction{
						Arguments: []*proto.FunctionArgument{
							{
								ArgType: &proto.FunctionArgument_Value{
									Value: &proto.Expression{RexType: &proto.Expression_Literal_{}},
								},
							},
						},
						Partitions: []*proto.Expression{
							{RexType: &proto.Expression_Literal_{}},
						},
						Sorts: []*proto.SortField{
							{
								Expr: &proto.Expression{RexType: &proto.Expression_Literal_{}},
							},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_IfThen",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_IfThen_{
					IfThen: &proto.Expression_IfThen{
						Ifs: []*proto.Expression_IfThen_IfClause{
							{
								If:   &proto.Expression{RexType: &proto.Expression_Literal_{}},
								Then: &proto.Expression{RexType: &proto.Expression_Literal_{}},
							},
						},
						Else: &proto.Expression{RexType: &proto.Expression_Literal_{}},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_SwitchExpression",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_SwitchExpression_{
					SwitchExpression: &proto.Expression_SwitchExpression{
						Match: &proto.Expression{RexType: &proto.Expression_Literal_{}},
						Ifs: []*proto.Expression_SwitchExpression_IfValue{
							{
								If:   &proto.Expression_Literal{},
								Then: &proto.Expression{RexType: &proto.Expression_Literal_{}},
							},
						},
						Else: &proto.Expression{RexType: &proto.Expression_Literal_{}},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_SingularOrList",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_SingularOrList_{
					SingularOrList: &proto.Expression_SingularOrList{
						Value: &proto.Expression{RexType: &proto.Expression_Literal_{}},
						Options: []*proto.Expression{
							{RexType: &proto.Expression_Literal_{}},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_MultiOrList",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_MultiOrList_{
					MultiOrList: &proto.Expression_MultiOrList{
						Value: []*proto.Expression{
							{RexType: &proto.Expression_Literal_{}},
						},
						Options: []*proto.Expression_MultiOrList_Record{
							{
								Fields: []*proto.Expression{
									{RexType: &proto.Expression_Literal_{}},
								},
							},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Cast",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Cast_{
					Cast: &proto.Expression_Cast{
						Input: &proto.Expression{RexType: &proto.Expression_Literal_{}},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Nested_Struct",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Nested_{
					Nested: &proto.Expression_Nested{
						NestedType: &proto.Expression_Nested_Struct_{
							Struct: &proto.Expression_Nested_Struct{
								Fields: []*proto.Expression{
									{RexType: &proto.Expression_Literal_{}},
								},
							},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Nested_List",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Nested_{
					Nested: &proto.Expression_Nested{
						NestedType: &proto.Expression_Nested_List_{
							List: &proto.Expression_Nested_List{
								Values: []*proto.Expression{
									{RexType: &proto.Expression_Literal_{}},
								},
							},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Nested_Map",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Nested_{
					Nested: &proto.Expression_Nested{
						NestedType: &proto.Expression_Nested_Map_{
							Map: &proto.Expression_Nested_Map{
								KeyValues: []*proto.Expression_Nested_Map_KeyValue{
									{
										Key:   &proto.Expression{RexType: &proto.Expression_Literal_{}},
										Value: &proto.Expression{RexType: &proto.Expression_Literal_{}},
									},
								},
							},
						},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Selection",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Selection{
					Selection: &proto.Expression_FieldReference{
						ReferenceType: &proto.Expression_FieldReference_DirectReference{},
					},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Literal",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Literal_{
					Literal: &proto.Expression_Literal{},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
		{
			name: "Expression_Enum",
			rel: createRelWithExpr(&proto.Expression{
				RexType: &proto.Expression_Enum_{
					Enum: &proto.Expression_Enum{},
				},
			}),
			expectedCount: 2, // Filter + ExtensionLeaf
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
				return count + 1
			})
			Walk(tt.rel, counter)

			// Should visit the exact expected number of nodes
			require.Equal(t, tt.expectedCount, counter.Result(), "Should visit exactly %d nodes", tt.expectedCount)
		})
	}
}

// TestMaskExpressionTypes ensures mask expression coverage
func TestMaskExpressionTypes(t *testing.T) {
	tests := []struct {
		name          string
		rel           *proto.Rel
		expectedCount int
	}{
		{
			name: "MaskExpression_Struct",
			rel: &proto.Rel{
				RelType: &proto.Rel_Read{
					Read: &proto.ReadRel{
						Projection: &proto.Expression_MaskExpression{
							Select: &proto.Expression_MaskExpression_StructSelect{
								StructItems: []*proto.Expression_MaskExpression_StructItem{
									{
										Field: 1,
										Child: &proto.Expression_MaskExpression_Select{
											Type: &proto.Expression_MaskExpression_Select_Struct{
												Struct: &proto.Expression_MaskExpression_StructSelect{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 1, // Read only
		},
		{
			name: "MaskExpression_List",
			rel: &proto.Rel{
				RelType: &proto.Rel_Read{
					Read: &proto.ReadRel{
						Projection: &proto.Expression_MaskExpression{
							Select: &proto.Expression_MaskExpression_StructSelect{
								StructItems: []*proto.Expression_MaskExpression_StructItem{
									{
										Field: 1,
										Child: &proto.Expression_MaskExpression_Select{
											Type: &proto.Expression_MaskExpression_Select_List{
												List: &proto.Expression_MaskExpression_ListSelect{
													Child: &proto.Expression_MaskExpression_Select{
														Type: &proto.Expression_MaskExpression_Select_Struct{
															Struct: &proto.Expression_MaskExpression_StructSelect{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 1, // Read only
		},
		{
			name: "MaskExpression_Map",
			rel: &proto.Rel{
				RelType: &proto.Rel_Read{
					Read: &proto.ReadRel{
						Projection: &proto.Expression_MaskExpression{
							Select: &proto.Expression_MaskExpression_StructSelect{
								StructItems: []*proto.Expression_MaskExpression_StructItem{
									{
										Field: 1,
										Child: &proto.Expression_MaskExpression_Select{
											Type: &proto.Expression_MaskExpression_Select_Map{
												Map: &proto.Expression_MaskExpression_MapSelect{
													Child: &proto.Expression_MaskExpression_Select{
														Type: &proto.Expression_MaskExpression_Select_Struct{
															Struct: &proto.Expression_MaskExpression_StructSelect{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedCount: 1, // Read only
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
				return count + 1
			})

			Walk(tt.rel, counter)
			require.Equal(t, tt.expectedCount, counter.Result(), "Should visit exactly %d nodes", tt.expectedCount)
		})
	}
}

// TestFieldReferenceTypes ensures field reference coverage
func TestFieldReferenceTypes(t *testing.T) {
	tests := []struct {
		name          string
		rel           *proto.Rel
		expectedCount int
	}{
		{
			name: "FieldReference_DirectReference",
			rel: &proto.Rel{
				RelType: &proto.Rel_HashJoin{
					HashJoin: &proto.HashJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Keys: []*proto.ComparisonJoinKey{
							{
								Left: &proto.Expression_FieldReference{
									ReferenceType: &proto.Expression_FieldReference_DirectReference{
										DirectReference: &proto.Expression_ReferenceSegment{},
									},
								},
								Right: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
			expectedCount: 3, // HashJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "FieldReference_MaskedReference",
			rel: &proto.Rel{
				RelType: &proto.Rel_HashJoin{
					HashJoin: &proto.HashJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Keys: []*proto.ComparisonJoinKey{
							{
								Left: &proto.Expression_FieldReference{
									ReferenceType: &proto.Expression_FieldReference_MaskedReference{
										MaskedReference: &proto.Expression_MaskExpression{
											Select: &proto.Expression_MaskExpression_StructSelect{},
										},
									},
								},
								Right: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
			expectedCount: 3, // HashJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "FieldReference_Expression",
			rel: &proto.Rel{
				RelType: &proto.Rel_HashJoin{
					HashJoin: &proto.HashJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Keys: []*proto.ComparisonJoinKey{
							{
								Left: &proto.Expression_FieldReference{
									RootType: &proto.Expression_FieldReference_Expression{
										Expression: &proto.Expression{
											RexType: &proto.Expression_Literal_{},
										},
									},
								},
								Right: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
			expectedCount: 3, // HashJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "FieldReference_RootReference",
			rel: &proto.Rel{
				RelType: &proto.Rel_HashJoin{
					HashJoin: &proto.HashJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Keys: []*proto.ComparisonJoinKey{
							{
								Left: &proto.Expression_FieldReference{
									RootType: &proto.Expression_FieldReference_RootReference_{},
								},
								Right: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
			expectedCount: 3, // HashJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
		{
			name: "FieldReference_OuterReference",
			rel: &proto.Rel{
				RelType: &proto.Rel_HashJoin{
					HashJoin: &proto.HashJoinRel{
						Left:  &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Right: &proto.Rel{RelType: &proto.Rel_ExtensionLeaf{}},
						Keys: []*proto.ComparisonJoinKey{
							{
								Left: &proto.Expression_FieldReference{
									RootType: &proto.Expression_FieldReference_OuterReference_{},
								},
								Right: &proto.Expression_FieldReference{},
							},
						},
					},
				},
			},
			expectedCount: 3, // HashJoin + Left ExtensionLeaf + Right ExtensionLeaf
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := newStatefulVisitor(0, func(count int, rel *proto.Rel) int {
				return count + 1
			})

			Walk(tt.rel, counter)
			require.Equal(t, tt.expectedCount, counter.Result(), "Should visit exactly %d nodes", tt.expectedCount)
		})
	}
}
