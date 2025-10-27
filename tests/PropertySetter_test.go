package test

import (
	"testing"
)

// TestPropertySetter 测试属性设置器表达式（基于 Java 的 PropertySetterTest.java）
func TestPropertySetter(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Interface object property access - interfaceObject.property",
			input: "interfaceObject.property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "interfaceObject.property",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "interfaceObject",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"interfaceObject\""},
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
			name:  "Object property access - object.property",
			input: "object.property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "object.property",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "object",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"object\""},
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
			name:  "Nested property access - object.nested.property",
			input: "object.nested.property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "object.nested.property",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "object",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"object\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "nested",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"nested\""},
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
			name:  "Static property access - ClassName.staticProperty",
			input: "ClassName.staticProperty",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "ClassName.staticProperty",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "ClassName",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"ClassName\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "staticProperty",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"staticProperty\""},
						},
					},
				},
			},
		},
		{
			name:  "Complex expression - object.property != null",
			input: "object.property != null",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "object.property != null",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "object.property",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "object",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"object\""},
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
					{
						Type:     "ASTConst",
						Fragment: "null",
						Children: []ExpectedNode{},
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
