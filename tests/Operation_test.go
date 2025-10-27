package test

import (
	"testing"
)

// TestOperation 测试 OGNL 操作表达式
// 对应 Java 的 OperationTest.java
// 测试各种被认为是"操作"的 OGNL 表达式，与简单的值引用相对
func TestOperation(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		// 非操作 - 简单变量引用
		{
			name:  "Simple variable reference - not an operation",
			input: "#name",
			expected: ExpectedNode{
				Type:     "ASTVarRef",
				Fragment: "#name",
			},
		},

		// 赋值操作
		{
			name:  "Assignment operation",
			input: "#name = 'boo'",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#name = \"boo\"",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#name"},
					{Type: "ASTConst", Fragment: "\"boo\""},
				},
			},
		},
		{
			name:  "Indexed assignment operation",
			input: "#name['foo'] = 'bar'",
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
		{
			name:  "Property assignment with concatenation",
			input: "#name.foo = 'bar' + 'foo'",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#name.foo = \"bar\" + \"foo\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "#name.foo",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#name"},
							{
								Type:     "ASTProperty",
								Fragment: "foo",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"foo\""},
								},
							},
						},
					},
					{
						Type:     "ASTAdd",
						Fragment: "\"bar\" + \"foo\"",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
							{Type: "ASTConst", Fragment: "\"foo\""},
						},
					},
				},
			},
		},
		{
			name:  "Sequence with assignment and method call",
			input: "{name.foo = 'bar' + 'foo', #name.foo()}",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ name.foo = \"bar\" + \"foo\", #name.foo() }",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "name.foo = \"bar\" + \"foo\"",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "name.foo",
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
										Fragment: "foo",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"foo\""},
										},
									},
								},
							},
							{
								Type:     "ASTAdd",
								Fragment: "\"bar\" + \"foo\"",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"bar\""},
									{Type: "ASTConst", Fragment: "\"foo\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#name.foo()",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#name"},
							{Type: "ASTMethod", Fragment: "foo()"},
						},
					},
				},
			},
		},
		{
			name:  "Parenthesized sequence with operations",
			input: "('bar' + 'foo', #name.foo())",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "(\"bar\" + \"foo\"), #name.foo()",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "\"bar\" + \"foo\"",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
							{Type: "ASTConst", Fragment: "\"foo\""},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#name.foo()",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#name"},
							{Type: "ASTMethod", Fragment: "foo()"},
						},
					},
				},
			},
		},

		// 一元运算操作
		{
			name:  "Unary minus with variable - operation",
			input: "-bar",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-bar",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bar",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
						},
					},
				},
			},
		},
		{
			name:  "Unary minus with parenthesized variable - operation",
			input: "-(#bar)",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-#bar",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#bar"},
				},
			},
		},
		{
			name:  "Unary minus with literal - not an operation",
			input: "-1",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			name:  "Unary minus with parenthesized expression - operation",
			input: "-(#bar+#foo)",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-(#bar + #foo)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "#bar + #foo",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#bar"},
							{Type: "ASTVarRef", Fragment: "#foo"},
						},
					},
				},
			},
		},

		// 序列和赋值操作
		{
			name:  "Complex sequence with assignment and arithmetic",
			input: "#bar=3,#foo=4(#bar-#foo)",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#bar = 3, #foo = (4)(#bar - #foo)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#bar = 3",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#bar"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
					{
						Type:     "ASTAssign",
						Fragment: "#foo = (4)(#bar - #foo)",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#foo"},
							{
								Type:     "ASTEval",
								Fragment: "(4)(#bar - #foo)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "4"},
									{
										Type:     "ASTSubtract",
										Fragment: "#bar - #foo",
										Children: []ExpectedNode{
											{Type: "ASTVarRef", Fragment: "#bar"},
											{Type: "ASTVarRef", Fragment: "#foo"},
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
			name:  "Simple arithmetic operation",
			input: "#bar-3",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "#bar - 3",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#bar"},
					{Type: "ASTConst", Fragment: "3"},
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
