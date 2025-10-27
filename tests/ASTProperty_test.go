package test

import (
	"testing"
)

// TestASTProperty 测试属性访问表达式（基于 Java 的 ASTPropertyTest.java）
func TestASTProperty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple property access",
			input: "nested",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "nested",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"nested\""},
				},
			},
		},
		{
			name:  "Property chain",
			input: "bean.value",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "bean.value",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean\""},
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
			name:  "Indexed property access",
			input: "list[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[0]",
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
						Fragment: "[0]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
		{
			name:  "Nested indexed property",
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
			name:  "Property with string index",
			input: "thing[\"x\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "thing[\"x\"]",
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
				},
			},
		},
		{
			name:  "Property chain with string index",
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
		{
			name:  "Generic property access",
			input: "cracker.param",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "cracker.param",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "cracker",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"cracker\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "param",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"param\""},
						},
					},
				},
			},
		},
		{
			name:  "Property access on method result",
			input: "getList().size",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getList().size",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getList()",
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
		{
			name:  "Complex property chain with index",
			input: "list[genericIndex].value",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[genericIndex].value",
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
			name:  "Variable reference property",
			input: "#root.property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#root.property",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
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

// TestComplicatedList 测试复杂的列表表达式（对应 Java 的 test_Complicated_List）
func TestComplicatedList(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "List with nested constructor calls",
			input: `{ new MenuItem('Home', 'Main'), new MenuItem('Help', 'Help') }`,
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: `{ new MenuItem("Home", "Main"), new MenuItem("Help", "Help") }`,
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new MenuItem(\"Home\", \"Main\")",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"Home\""},
							{Type: "ASTConst", Fragment: "\"Main\""},
						},
					},
					{
						Type:     "ASTCtor",
						Fragment: "new MenuItem(\"Help\", \"Help\")",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"Help\""},
							{Type: "ASTConst", Fragment: "\"Help\""},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
