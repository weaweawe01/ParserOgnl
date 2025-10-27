package test

import (
	"testing"
)

// TestASTSequence 测试序列表达式（基于 Java 的 ASTSequenceTest.java）
func TestASTSequence(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    ExpectedNode
		isSequence  bool
		description string
	}{
		{
			name:        "Single variable reference - not a sequence",
			input:       "#name",
			isSequence:  false,
			description: "单个变量引用不是序列",
			expected: ExpectedNode{
				Type:     "ASTVarRef",
				Fragment: "#name",
			},
		},
		{
			name:        "Assignment and method call - is a sequence",
			input:       "#name = 'boo', System.out.println(#name)",
			isSequence:  true,
			description: "赋值和方法调用用逗号分隔，是序列表达式",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#name = \"boo\", System.out.println(#name)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#name = \"boo\"",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#name"},
							{Type: "ASTConst", Fragment: "\"boo\""},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "System.out.println(#name)",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "System",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"System\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "out",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"out\""},
								},
							},
							{
								Type:     "ASTMethod",
								Fragment: "println(#name)",
								Children: []ExpectedNode{
									{Type: "ASTVarRef", Fragment: "#name"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:        "Indexed assignment - not a sequence",
			input:       "#name['foo'] = 'bar'",
			isSequence:  false,
			description: "带索引的赋值不是序列",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#name[\"foo\"] = \"bar\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "#name[\"foo\"]",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#name"},
							{
								Type:     "ASTProperty",
								Fragment: "[\"foo\"]",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"foo\""},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "\"bar\""},
				},
			},
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)

			// 检查是否为序列表达式
			actualIsSequence := (expr.Type() == "ASTSequence")
			if actualIsSequence != tc.isSequence {
				t.Errorf("表达式 '%s' 序列判断错误: 期望 %v, 实际 %v",
					tc.input, tc.isSequence, actualIsSequence)
			}

			// 检查 AST 结构
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// TestSequenceExpressions 测试更多序列表达式的情况
func TestSequenceExpressions(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple two-expression sequence",
			input: "a = 1, b = 2",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "a = 1, b = 2",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "a = 1",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "a",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"a\""},
								},
							},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
					{
						Type:     "ASTAssign",
						Fragment: "b = 2",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "b",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"b\""},
								},
							},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
				},
			},
		},
		{
			name:  "Three-expression sequence",
			input: "x = 1, y = 2, x + y",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "x = 1, y = 2, (x + y)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "x = 1",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "x",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"x\""},
								},
							},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
					{
						Type:     "ASTAssign",
						Fragment: "y = 2",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "y",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"y\""},
								},
							},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
					{
						Type:     "ASTAdd",
						Fragment: "x + y",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "x",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"x\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "y",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"y\""},
								},
							},
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
