package test

import (
	"testing"
)

// TestPropertyArithmeticAndLogicalOperators 测试属性访问与算术/逻辑运算符的组合表达式
// 对应 Java 的 PropertyArithmeticAndLogicalOperatorsTest.java
// 测试属性访问、方法调用与各种运算符的复杂组合
func TestPropertyArithmeticAndLogicalOperators(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		// testBooleanExpressions - 布尔表达式
		{
			input: "objectIndex > 0",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "objectIndex > 0",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "objectIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"objectIndex\""},
						},
					},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "false",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "false",
			},
		},
		{
			input: "!false || true",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "!false || true",
				Children: []ExpectedNode{
					{
						Type:     "ASTNot",
						Fragment: "!false",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "false"},
						},
					},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
		{
			input: "property.bean3.value >= 24",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "property.bean3.value >= 24",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "property.bean3.value",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "property",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"property\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "bean3",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"bean3\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "value",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"value\""},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "24"},
				},
			},
		},
		{
			input: "(unassignedCopyModel.optionCount > 0 && canApproveCopy) || entry.copy.size() > 0",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "((unassignedCopyModel.optionCount > 0) && canApproveCopy) || (entry.copy.size() > 0)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAnd",
						Fragment: "((unassignedCopyModel.optionCount > 0) && canApproveCopy)",
						Children: []ExpectedNode{
							{
								Type:     "ASTGreater",
								Fragment: "(unassignedCopyModel.optionCount > 0)",
								Children: []ExpectedNode{
									{
										Type:     "ASTChain",
										Fragment: "unassignedCopyModel.optionCount",
										Children: []ExpectedNode{
											{
												Type:     "ASTProperty",
												Fragment: "unassignedCopyModel",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"unassignedCopyModel\""},
												},
											},
											{
												Type:     "ASTProperty",
												Fragment: "optionCount",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"optionCount\""},
												},
											},
										},
									},
									{Type: "ASTConst", Fragment: "0"},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "canApproveCopy",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"canApproveCopy\""},
								},
							},
						},
					},
					{
						Type:     "ASTGreater",
						Fragment: "(entry.copy.size() > 0)",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "entry.copy.size()",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "entry",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"entry\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "copy",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"copy\""},
										},
									},
									{Type: "ASTMethod", Fragment: "size()"},
								},
							},
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
		{
			input: "!(printDelivery || @Boolean@FALSE)",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!(printDelivery || @Boolean@FALSE)",
				Children: []ExpectedNode{
					{
						Type:     "ASTOr",
						Fragment: "printDelivery || @Boolean@FALSE",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "printDelivery",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"printDelivery\""},
								},
							},
							{Type: "ASTStaticField", Fragment: "@Boolean@FALSE"},
						},
					},
				},
			},
		},
		// testIntegerExpressions - 整数表达式
		{
			input: "genericIndex-1",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "genericIndex - 1",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "genericIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"genericIndex\""},
						},
					},
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "((renderNavigation ? 0 : 1) + map.size) * theInt",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "(renderNavigation ? 0 : 1 + map.size) * theInt",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "renderNavigation ? 0 : 1 + map.size",
						Children: []ExpectedNode{
							{
								Type:     "ASTTest",
								Fragment: "renderNavigation ? 0 : 1",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "renderNavigation",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"renderNavigation\""},
										},
									},
									{Type: "ASTConst", Fragment: "0"},
									{Type: "ASTConst", Fragment: "1"},
								},
							},
							{
								Type:     "ASTChain",
								Fragment: "map.size",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "map",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"map\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "size",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"size\""},
										},
									},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "theInt",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"theInt\""},
						},
					},
				},
			},
		},
		{
			input: "{theInt + 1}",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ theInt + 1 }",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "theInt + 1",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "theInt",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"theInt\""},
								},
							},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
				},
			},
		},
		{
			input: "(getIndexedProperty('nested').size - 1) > genericIndex",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "(getIndexedProperty(\"nested\").size - 1) > genericIndex",
				Children: []ExpectedNode{
					{
						Type:     "ASTSubtract",
						Fragment: "getIndexedProperty(\"nested\").size - 1",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "getIndexedProperty(\"nested\").size",
								Children: []ExpectedNode{
									{
										Type:     "ASTMethod",
										Fragment: "getIndexedProperty(\"nested\")",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"nested\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "size",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"size\""},
										},
									},
								},
							},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "genericIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"genericIndex\""},
						},
					},
				},
			},
		},
		{
			input: "(getIndexedProperty('nested').size + 1) >= genericIndex",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "(getIndexedProperty(\"nested\").size + 1) >= genericIndex",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "getIndexedProperty(\"nested\").size + 1",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "getIndexedProperty(\"nested\").size",
								Children: []ExpectedNode{
									{
										Type:     "ASTMethod",
										Fragment: "getIndexedProperty(\"nested\")",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"nested\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "size",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"size\""},
										},
									},
								},
							},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "genericIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"genericIndex\""},
						},
					},
				},
			},
		},
		{
			input: "(getIndexedProperty('nested').size + 1) == genericIndex",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "(getIndexedProperty(\"nested\").size + 1) == genericIndex",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "getIndexedProperty(\"nested\").size + 1",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "getIndexedProperty(\"nested\").size",
								Children: []ExpectedNode{
									{
										Type:     "ASTMethod",
										Fragment: "getIndexedProperty(\"nested\")",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"nested\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "size",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"size\""},
										},
									},
								},
							},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "genericIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"genericIndex\""},
						},
					},
				},
			},
		},
		{
			input: "map.size * genericIndex",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "map.size * genericIndex",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "map.size",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "map",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"map\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "size",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"size\""},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "genericIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"genericIndex\""},
						},
					},
				},
			},
		},
		{
			input: "property == property",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "property == property",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "property",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"property\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "property",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"property\""},
						},
					},
				},
			},
		},
		{
			input: "property.bean3.value % 2 == 0",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "(property.bean3.value % 2) == 0",
				Children: []ExpectedNode{
					{
						Type:     "ASTRemainder",
						Fragment: "(property.bean3.value % 2)",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "property.bean3.value",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "property",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"property\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "bean3",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"bean3\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "value",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"value\""},
										},
									},
								},
							},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "genericIndex % 3 == 0",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "(genericIndex % 3) == 0",
				Children: []ExpectedNode{
					{
						Type:     "ASTRemainder",
						Fragment: "(genericIndex % 3)",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "genericIndex",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"genericIndex\""},
								},
							},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "genericIndex % theInt == property.bean3.value",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "(genericIndex % theInt) == property.bean3.value",
				Children: []ExpectedNode{
					{
						Type:     "ASTRemainder",
						Fragment: "(genericIndex % theInt)",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "genericIndex",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"genericIndex\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "theInt",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"theInt\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "property.bean3.value",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "property",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"property\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "bean3",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"bean3\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "value",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"value\""},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "theInt / 100.0",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "theInt / 100.0",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "theInt",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"theInt\""},
						},
					},
					{Type: "ASTConst", Fragment: "100.0"},
				},
			},
		},
		{
			input: "@java.lang.Long@valueOf('100') == @java.lang.Long@valueOf('100')",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "@java.lang.Long@valueOf(\"100\") == @java.lang.Long@valueOf(\"100\")",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticMethod",
						Fragment: "@java.lang.Long@valueOf(\"100\")",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"100\""},
						},
					},
					{
						Type:     "ASTStaticMethod",
						Fragment: "@java.lang.Long@valueOf(\"100\")",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"100\""},
						},
					},
				},
			},
		},
		// testDoubleExpressions - 浮点数表达式
		{
			input: "budget - timeBilled",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "budget - timeBilled",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "budget",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"budget\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "timeBilled",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"timeBilled\""},
						},
					},
				},
			},
		},
		{
			input: "(budget % tableSize) == 0",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "(budget % tableSize) == 0",
				Children: []ExpectedNode{
					{
						Type:     "ASTRemainder",
						Fragment: "budget % tableSize",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "budget",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"budget\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "tableSize",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"tableSize\""},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
