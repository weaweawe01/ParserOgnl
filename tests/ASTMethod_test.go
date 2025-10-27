package test

import (
	"testing"
)

// TestASTMethod 测试方法调用表达式（基于 Java 的 ASTMethodTest.java）
func TestASTMethod(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple method call",
			input: "execute()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "execute()",
			},
		},
		{
			name:  "Method call with string argument",
			input: "get(\"value\")",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "get(\"value\")",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"value\""},
				},
			},
		},
		{
			name:  "Chained method call",
			input: "bean.execute()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "bean.execute()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "execute()",
					},
				},
			},
		},
		{
			name:  "Method call on property chain",
			input: "aaa.get(\"value\")",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "aaa.get(\"value\")",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "aaa",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"aaa\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "get(\"value\")",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"value\""},
						},
					},
				},
			},
		},
		{
			name:  "Method call with multiple arguments",
			input: "method(\"arg1\", 123)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "method(\"arg1\", 123)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"arg1\""},
					{Type: "ASTConst", Fragment: "123"},
				},
			},
		},
		{
			name:  "Property access after method call",
			input: "bean.getBean3().value",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "bean.getBean3().value",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getBean3()",
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
			name:  "Variable reference with property",
			input: "#name.lastChar",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#name.lastChar",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#name"},
					{
						Type:     "ASTProperty",
						Fragment: "lastChar",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"lastChar\""},
						},
					},
				},
			},
		},
		{
			name:  "Array literal with method calls",
			input: "{name.lastChar, foo()}",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ name.lastChar, foo() }",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "name.lastChar",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "name",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"name\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "lastChar",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"lastChar\""},
								},
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "foo()",
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
