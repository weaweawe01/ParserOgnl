package test

import (
	"testing"
)

// TestNullRoot 测试空值和空根对象访问表达式（基于 Java 的 NullRootTest.java）
func TestNullRoot(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Null value chain access - key1.key2.key3",
			input: "key1.key2.key3",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "key1.key2.key3",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "key1",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key1\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "key2",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key2\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "key3",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key3\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Empty root chain access - key1.key2.key3 (same expression)",
			input: "key1.key2.key3",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "key1.key2.key3",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "key1",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key1\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "key2",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key2\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "key3",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key3\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Null root chain access - key1.key2.key3 (same expression)",
			input: "key1.key2.key3",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "key1.key2.key3",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "key1",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key1\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "key2",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key2\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "key3",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"key3\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Simple property access - user.name",
			input: "user.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "user.name",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "user",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"user\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"name\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Nested method call - user.getProfile().getName()",
			input: "user.getProfile().getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "user.getProfile().getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "user",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"user\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getProfile()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		{
			name:  "Mixed property and method access - data.items[0].name",
			input: "data.items[0].name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "data.items[0].name",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "data",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"data\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "items",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"items\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[0]",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "0",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"name\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call with null argument - process(null)",
			input: "process(null)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "process(null)",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "null",
					},
				},
			},
		},
		{
			name:  "Property access on null - null.property",
			input: "null.property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "null.property",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "null",
					},
					{
						Type:     "ASTProperty",
						Fragment: "property",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"property\"",
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
