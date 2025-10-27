package test

import (
	"testing"
)

// TestIndexAccess 测试索引访问表达式
// 对应 Java 的 IndexAccessTest.java
func TestIndexAccess(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "List index - list[index]",
			input: "list[index]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[index]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "List object index - list[objectIndex]",
			input: "list[objectIndex]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[objectIndex]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[objectIndex]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "objectIndex",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"objectIndex\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Array object index - array[objectIndex]",
			input: "array[objectIndex]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[objectIndex]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "array",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"array\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[objectIndex]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "objectIndex",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"objectIndex\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Array get object index - array[getObjectIndex()]",
			input: "array[getObjectIndex()]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[getObjectIndex()]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "array",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"array\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[getObjectIndex()]",
						Children: []ExpectedNode{
							{Type: "ASTMethod", Fragment: "getObjectIndex()"},
						},
					},
				},
			},
		},
		{
			name:  "Array generic index - array[genericIndex]",
			input: "array[genericIndex]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[genericIndex]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "array",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"array\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[genericIndex]",
						Children: []ExpectedNode{
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
			},
		},
		{
			name:  "Boolean array self object index - booleanArray[self.objectIndex]",
			input: "booleanArray[self.objectIndex]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "booleanArray[self.objectIndex]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "booleanArray",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"booleanArray\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[self.objectIndex]",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "self.objectIndex",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "self",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"self\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "objectIndex",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"objectIndex\""},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Boolean array get object index - booleanArray[getObjectIndex()]",
			input: "booleanArray[getObjectIndex()]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "booleanArray[getObjectIndex()]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "booleanArray",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"booleanArray\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[getObjectIndex()]",
						Children: []ExpectedNode{
							{Type: "ASTMethod", Fragment: "getObjectIndex()"},
						},
					},
				},
			},
		},
		{
			name:  "Boolean array null index - booleanArray[nullIndex]",
			input: "booleanArray[nullIndex]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "booleanArray[nullIndex]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "booleanArray",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"booleanArray\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[nullIndex]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "nullIndex",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"nullIndex\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "List size minus one - list[size() - 1]",
			input: "list[size() - 1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[size() - 1]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[size() - 1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTSubtract",
								Fragment: "size() - 1",
								Children: []ExpectedNode{
									{Type: "ASTMethod", Fragment: "size()"},
									{Type: "ASTConst", Fragment: "1"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Toggle toggle selected - (index == (array.length - 3)) ? 'toggle toggleSelected' : 'toggle'",
			input: "(index == (array.length - 3)) ? 'toggle toggleSelected' : 'toggle'",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "(index == (array.length - 3)) ? \"toggle toggleSelected\" : \"toggle\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTEq",
						Fragment: "index == (array.length - 3)",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index\""},
								},
							},
							{
								Type:     "ASTSubtract",
								Fragment: "array.length - 3",
								Children: []ExpectedNode{
									{
										Type:     "ASTChain",
										Fragment: "array.length",
										Children: []ExpectedNode{
											{
												Type:     "ASTProperty",
												Fragment: "array",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"array\""},
												},
											},
											{
												Type:     "ASTProperty",
												Fragment: "length",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"length\""},
												},
											},
										},
									},
									{Type: "ASTConst", Fragment: "3"},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "\"toggle toggleSelected\""},
					{Type: "ASTConst", Fragment: "\"toggle\""},
				},
			},
		},
		{
			name:  "Toggle display - \"return toggleDisplay('excdisplay\"+index+\"', this)\"",
			input: "\"return toggleDisplay('excdisplay\"+index+\"', this)\"",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "(\"return toggleDisplay('excdisplay\" + index) + \"', this)\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "\"return toggleDisplay('excdisplay\" + index",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"return toggleDisplay('excdisplay\""},
							{
								Type:     "ASTProperty",
								Fragment: "index",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index\""},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "\"', this)\""},
				},
			},
		},
		{
			name:  "Map map key split - map[mapKey].split('=')[0]",
			input: "map[mapKey].split('=')[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[mapKey].split('=')[0]",
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
						Fragment: "[mapKey]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "mapKey",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"mapKey\""},
								},
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "split('=')",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "'='"},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[0]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
		{
			name:  "Boolean values index1 index2 - booleanValues[index1][index2]",
			input: "booleanValues[index1][index2]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "booleanValues[index1][index2]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "booleanValues",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"booleanValues\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index1",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index1\""},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index2]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index2",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index2\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Tab search criteria display name - tab.searchCriteria[index1].displayName",
			input: "tab.searchCriteria[index1].displayName",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "tab.searchCriteria[index1].displayName",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "tab",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"tab\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "searchCriteria",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"searchCriteria\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index1",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index1\""},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "displayName",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"displayName\""},
						},
					},
				},
			},
		},
		{
			name:  "Tab search criteria selections - tab.searchCriteriaSelections[index1][index2]",
			input: "tab.searchCriteriaSelections[index1][index2]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "tab.searchCriteriaSelections[index1][index2]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "tab",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"tab\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "searchCriteriaSelections",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"searchCriteriaSelections\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index1",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index1\""},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index2]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index2",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index2\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Tab search criteria selections set value - tab.searchCriteriaSelections[index1][index2]",
			input: "tab.searchCriteriaSelections[index1][index2]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "tab.searchCriteriaSelections[index1][index2]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "tab",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"tab\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "searchCriteriaSelections",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"searchCriteriaSelections\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index1",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index1\""},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index2]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "index2",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"index2\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Map bar value - map['bar'].value",
			input: "map['bar'].value",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[\"bar\"].value",
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
						Fragment: "[\"bar\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
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
		{
			name:  "Indexed set thing x val - thing[\"x\"].val",
			input: "thing[\"x\"].val",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "thing[\"x\"].val",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "thing",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"thing\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"x\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"x\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "val",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"val\""},
						},
					},
				},
			},
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
