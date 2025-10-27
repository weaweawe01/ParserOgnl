package test

import (
	"testing"
)

// TestConstantTree 测试常量表达式树（基于 Java 的 ConstantTreeTest.java）
// 该测试验证各种常量和非常量表达式的 AST 结构
func TestConstantTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Boolean true constant",
			input: "true",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "true",
			},
		},
		{
			name:  "Integer 55 constant",
			input: "55",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55",
			},
		},
		{
			name:  "Static field reference @java.awt.Color@black",
			input: "@java.awt.Color@black",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@java.awt.Color@black",
			},
		},
		{
			name:  "Non-final static variable reference",
			input: "@ognl.test.ConstantTreeTest@nonFinalStaticVariable",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@ognl.test.ConstantTreeTest@nonFinalStaticVariable",
			},
		},
		{
			name:  "Non-final static variable plus 10",
			input: "@ognl.test.ConstantTreeTest@nonFinalStaticVariable + 10",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "(@ognl.test.ConstantTreeTest@nonFinalStaticVariable + 10)",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@ognl.test.ConstantTreeTest@nonFinalStaticVariable",
					},
					{Type: "ASTConst", Fragment: "10"},
				},
			},
		},
		{
			name:  "Addition with static field",
			input: "55 + 24 + @java.awt.Event@ALT_MASK",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "((55 + 24) + @java.awt.Event@ALT_MASK)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "(55 + 24)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "55"},
							{Type: "ASTConst", Fragment: "24"},
						},
					},
					{
						Type:     "ASTStaticField",
						Fragment: "@java.awt.Event@ALT_MASK",
					},
				},
			},
		},
		{
			name:  "Property access - name",
			input: "name",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "name",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"name\""},
				},
			},
		},
		{
			name:  "Array index access - name[i]",
			input: "name[i]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name[i]",
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
						Fragment: "[i]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "i",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"i\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Chained property access - name[i].property",
			input: "name[i].property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name[i].property",
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
						Fragment: "[i]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "i",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"i\""},
								},
							},
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
			name:  "Selection - name.{? foo }",
			input: "name.{? foo }",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name.{? foo}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTSelectFirst",
						Fragment: "{? foo}",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "foo",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"foo\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Projection - name.{ foo }",
			input: "name.{ foo }",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name.{foo}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTProject",
						Fragment: "{foo}",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "foo",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"foo\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Projection with constant - name.{ 25 }",
			input: "name.{ 25 }",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name.{25}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTProject",
						Fragment: "{25}",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "25"},
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
