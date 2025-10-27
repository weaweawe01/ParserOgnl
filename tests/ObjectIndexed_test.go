package test

import (
	"testing"
)

// TestObjectIndexed 测试对象索引访问表达式
// 对应 Java 的 ObjectIndexedTest.java
// 测试对象索引访问表达式的功能，包括变量引用、属性链和索引访问
// 主要测试 #variable.property[index] 这种复杂的表达式结构
func TestObjectIndexed(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "testObjectIndexAccess",
			input: "#ka.sunk[#root]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#ka.sunk[#root]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#ka"},
					{
						Type:     "ASTProperty",
						Fragment: "sunk",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"sunk\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[#root]",
						Children: []ExpectedNode{
							{Type: "ASTRootVarRef", Fragment: "#root"},
						},
					},
				},
			},
		},
		{
			name:  "testObjectIndexInSubclass",
			input: "#ka.sunk[#root]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#ka.sunk[#root]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#ka"},
					{
						Type:     "ASTProperty",
						Fragment: "sunk",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"sunk\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[#root]",
						Children: []ExpectedNode{
							{Type: "ASTRootVarRef", Fragment: "#root"},
						},
					},
				},
			},
		},
		{
			name:  "testMultipleObjectIndexGetters",
			input: "#ka.sunk[#root]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#ka.sunk[#root]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#ka"},
					{
						Type:     "ASTProperty",
						Fragment: "sunk",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"sunk\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[#root]",
						Children: []ExpectedNode{
							{Type: "ASTRootVarRef", Fragment: "#root"},
						},
					},
				},
			},
		},
		{
			name:  "simpleVariableIndexAccess",
			input: "#map[key]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#map[key]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#map"},
					{
						Type:     "ASTProperty",
						Fragment: "[key]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "key",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"key\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "variableWithStringIndexAccess",
			input: "#list[\"item\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#list[\"item\"]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#list"},
					{
						Type:     "ASTProperty",
						Fragment: "[\"item\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"item\""},
						},
					},
				},
			},
		},
		{
			name:  "variableWithNumericIndexAccess",
			input: "#array[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#array[0]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#array"},
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
			name:  "nestedVariableIndexAccess",
			input: "#outer[#inner]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#outer[#inner]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#outer"},
					{
						Type:     "ASTProperty",
						Fragment: "[#inner]",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#inner"},
						},
					},
				},
			},
		},
		{
			name:  "variableWithPropertyChainAndIndex",
			input: "#obj.field1.field2[index]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#obj.field1.field2[index]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#obj"},
					{
						Type:     "ASTProperty",
						Fragment: "field1",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"field1\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "field2",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"field2\""},
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
			name:  "variableWithExpressionIndex",
			input: "#data[i + 1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#data[i + 1]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#data"},
					{
						Type:     "ASTProperty",
						Fragment: "[i + 1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTAdd",
								Fragment: "i + 1",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "i",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"i\""},
										},
									},
									{Type: "ASTConst", Fragment: "1"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "thisVariableWithIndexAccess",
			input: "#this.items[#root]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#this.items[#root]",
				Children: []ExpectedNode{
					{Type: "ASTThisVarRef", Fragment: "#this"},
					{
						Type:     "ASTProperty",
						Fragment: "items",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"items\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[#root]",
						Children: []ExpectedNode{
							{Type: "ASTRootVarRef", Fragment: "#root"},
						},
					},
				},
			},
		},
		{
			name:  "multipleIndexAccess",
			input: "#matrix[row][col]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#matrix[row][col]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#matrix"},
					{
						Type:     "ASTProperty",
						Fragment: "[row]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "row",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"row\""},
								},
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[col]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "col",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"col\""},
								},
							},
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
