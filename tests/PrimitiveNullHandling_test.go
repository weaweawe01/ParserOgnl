package test

import (
	"testing"
)

// TestPrimitiveNullHandling 测试原始类型空值处理表达式
// 对应 Java 的 PrimitiveNullHandlingTest.java
// 测试原始类型属性访问的 AST 结构（虽然 Go 版本不测试运行时 null 处理行为）
func TestPrimitiveNullHandling(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Float property access - floatValue",
			input: "floatValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "floatValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"floatValue\""},
				},
			},
		},
		{
			name:  "Int property access - intValue",
			input: "intValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "intValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"intValue\""},
				},
			},
		},
		{
			name:  "Boolean property access - booleanValue",
			input: "booleanValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "booleanValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"booleanValue\""},
				},
			},
		},
		{
			name:  "Float property assignment - floatValue = 10.56f",
			input: "floatValue = 10.56f",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "floatValue = 10.56",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "floatValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"floatValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "10.56"},
				},
			},
		},
		{
			name:  "Int property assignment - intValue = 34",
			input: "intValue = 34",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "intValue = 34",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "intValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"intValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "34"},
				},
			},
		},
		{
			name:  "Boolean property assignment - booleanValue = true",
			input: "booleanValue = true",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "booleanValue = true",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "booleanValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"booleanValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
		{
			name:  "Float property assignment with null - floatValue = null",
			input: "floatValue = null",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "floatValue = null",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "floatValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"floatValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		{
			name:  "Int property assignment with null - intValue = null",
			input: "intValue = null",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "intValue = null",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "intValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"intValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		{
			name:  "Boolean property assignment with null - booleanValue = null",
			input: "booleanValue = null",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "booleanValue = null",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "booleanValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"booleanValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		{
			name:  "Chained property access - simple.floatValue",
			input: "simple.floatValue",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "simple.floatValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "simple",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"simple\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "floatValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"floatValue\""},
						},
					},
				},
			},
		},
		{
			name:  "Chained property assignment - simple.intValue = 100",
			input: "simple.intValue = 100",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "simple.intValue = 100",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "simple.intValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "simple",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"simple\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "intValue",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"intValue\""},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "100"},
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
