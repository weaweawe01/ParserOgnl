package test

import (
	"testing"
)

// TestPrivateAccessor 测试私有访问器表达式
// 对应 Java 的 PrivateAccessorTest.java
// 测试私有属性和方法的 AST 结构（虽然 Go 版本不测试运行时访问控制）
func TestPrivateAccessor(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Private accessor method call - getPrivateAccessorIntValue()",
			input: "getPrivateAccessorIntValue()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getPrivateAccessorIntValue()",
			},
		},
		{
			name:  "Private accessor property access - privateAccessorIntValue",
			input: "privateAccessorIntValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "privateAccessorIntValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"privateAccessorIntValue\""},
				},
			},
		},
		{
			name:  "Private accessor property assignment - privateAccessorIntValue = 100",
			input: "privateAccessorIntValue = 100",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "privateAccessorIntValue = 100",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "privateAccessorIntValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"privateAccessorIntValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "100"},
				},
			},
		},
		{
			name:  "Private accessor property 2 access - privateAccessorIntValue2",
			input: "privateAccessorIntValue2",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "privateAccessorIntValue2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"privateAccessorIntValue2\""},
				},
			},
		},
		{
			name:  "Private accessor property 2 assignment - privateAccessorIntValue2 = 100",
			input: "privateAccessorIntValue2 = 100",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "privateAccessorIntValue2 = 100",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "privateAccessorIntValue2",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"privateAccessorIntValue2\""},
						},
					},
					{Type: "ASTConst", Fragment: "100"},
				},
			},
		},
		{
			name:  "Private accessor property 3 access - privateAccessorIntValue3",
			input: "privateAccessorIntValue3",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "privateAccessorIntValue3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"privateAccessorIntValue3\""},
				},
			},
		},
		{
			name:  "Private accessor property 3 assignment - privateAccessorIntValue3 = 100",
			input: "privateAccessorIntValue3 = 100",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "privateAccessorIntValue3 = 100",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "privateAccessorIntValue3",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"privateAccessorIntValue3\""},
						},
					},
					{Type: "ASTConst", Fragment: "100"},
				},
			},
		},
		{
			name:  "Private accessor boolean property access - privateAccessorBooleanValue",
			input: "privateAccessorBooleanValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "privateAccessorBooleanValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"privateAccessorBooleanValue\""},
				},
			},
		},
		{
			name:  "Private accessor boolean property assignment - privateAccessorBooleanValue = false",
			input: "privateAccessorBooleanValue = false",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "privateAccessorBooleanValue = false",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "privateAccessorBooleanValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"privateAccessorBooleanValue\""},
						},
					},
					{Type: "ASTConst", Fragment: "false"},
				},
			},
		},
		{
			name:  "Root object property access - root.privateAccessorIntValue",
			input: "root.privateAccessorIntValue",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "root.privateAccessorIntValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "root",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"root\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "privateAccessorIntValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"privateAccessorIntValue\""},
						},
					},
				},
			},
		},
		{
			name:  "Root object method call - root.getPrivateAccessorIntValue()",
			input: "root.getPrivateAccessorIntValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "root.getPrivateAccessorIntValue()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "root",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"root\""},
						},
					},
					{Type: "ASTMethod", Fragment: "getPrivateAccessorIntValue()"},
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
