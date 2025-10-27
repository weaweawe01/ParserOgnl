package test

import (
	"testing"
)

// TestContextVariable 测试上下文变量表达式
// 对应 Java 的 ContextVariableTest.java
func TestContextVariable(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Root variable reference",
			input: "#root",
			expected: ExpectedNode{
				Type:     "ASTRootVarRef",
				Fragment: "#root",
			},
		},
		{
			name:  "This variable reference",
			input: "#this",
			expected: ExpectedNode{
				Type:     "ASTThisVarRef",
				Fragment: "#this",
			},
		},
		{
			name:  "Sum of five and six - #f=5, #s=6, #f + #s",
			input: "#f=5, #s=6, #f + #s",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#f = 5, #s = 6, (#f + #s)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#f = 5",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#f"},
							{Type: "ASTConst", Fragment: "5"},
						},
					},
					{
						Type:     "ASTAssign",
						Fragment: "#s = 6",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#s"},
							{Type: "ASTConst", Fragment: "6"},
						},
					},
					{
						Type:     "ASTAdd",
						Fragment: "#f + #s",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#f"},
							{Type: "ASTVarRef", Fragment: "#s"},
						},
					},
				},
			},
		},
		{
			name:  "Sum with intermediate assignment - #six=(#five=5, 6), #five + #six",
			input: "#six=(#five=5, 6), #five + #six",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#six = #five = 5, 6, (#five + #six)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#six = #five = 5, 6",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#six"},
							{
								Type:     "ASTSequence",
								Fragment: "#five = 5, 6",
								Children: []ExpectedNode{
									{
										Type:     "ASTAssign",
										Fragment: "#five = 5",
										Children: []ExpectedNode{
											{Type: "ASTVarRef", Fragment: "#five"},
											{Type: "ASTConst", Fragment: "5"},
										},
									},
									{Type: "ASTConst", Fragment: "6"},
								},
							},
						},
					},
					{
						Type:     "ASTAdd",
						Fragment: "#five + #six",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#five"},
							{Type: "ASTVarRef", Fragment: "#six"},
						},
					},
				},
			},
		},
		{
			name:  "Simple variable assignment",
			input: "#var = 100",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#var = 100",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#var"},
					{Type: "ASTConst", Fragment: "100"},
				},
			},
		},
		{
			name:  "Variable in expression",
			input: "#x + #y",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "#x + #y",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#x"},
					{Type: "ASTVarRef", Fragment: "#y"},
				},
			},
		},
		{
			name:  "Variable with property access",
			input: "#user.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#user.name",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#user"},
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
				},
			},
		},
		{
			name:  "Variable with method call",
			input: "#obj.toString()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#obj.toString()",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#obj"},
					{Type: "ASTMethod", Fragment: "toString()"},
				},
			},
		},
		{
			name:  "Variable with array access",
			input: "#list[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#list[0]",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#list"},
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
			name:  "Multiple variable assignments in sequence",
			input: "#a=1, #b=2, #c=3",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#a = 1, #b = 2, #c = 3",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#a = 1",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#a"},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
					{
						Type:     "ASTAssign",
						Fragment: "#b = 2",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#b"},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
					{
						Type:     "ASTAssign",
						Fragment: "#c = 3",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#c"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			name:  "Variable in conditional expression",
			input: "#flag ? #a : #b",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "#flag ? #a : #b",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#flag"},
					{Type: "ASTVarRef", Fragment: "#a"},
					{Type: "ASTVarRef", Fragment: "#b"},
				},
			},
		},
		{
			name:  "Chain assignment",
			input: "#a = #b = #c = 5",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#a = #b = #c = 5",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#a"},
					{
						Type:     "ASTAssign",
						Fragment: "#b = #c = 5",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#b"},
							{
								Type:     "ASTAssign",
								Fragment: "#c = 5",
								Children: []ExpectedNode{
									{Type: "ASTVarRef", Fragment: "#c"},
									{Type: "ASTConst", Fragment: "5"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Variable with complex expression",
			input: "#result = #a * 2 + #b",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#result = (#a * 2) + #b",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#result"},
					{
						Type:     "ASTAdd",
						Fragment: "(#a * 2) + #b",
						Children: []ExpectedNode{
							{
								Type:     "ASTMultiply",
								Fragment: "(#a * 2)",
								Children: []ExpectedNode{
									{Type: "ASTVarRef", Fragment: "#a"},
									{Type: "ASTConst", Fragment: "2"},
								},
							},
							{Type: "ASTVarRef", Fragment: "#b"},
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
