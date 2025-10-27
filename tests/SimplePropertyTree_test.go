package test

import (
	"testing"
)

// TestSimplePropertyTree 测试简单属性树（基于 Java 的 SimplePropertyTreeTest.java）
func TestSimplePropertyTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple property - property",
			input: "property",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "property",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"property\""},
				},
			},
		},
		{
			name:  "Property with underscore - my_property",
			input: "my_property",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "my_property",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"my_property\""},
				},
			},
		},
		{
			name:  "Camel case property - camelCaseProperty",
			input: "camelCaseProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "camelCaseProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"camelCaseProperty\""},
				},
			},
		},
		{
			name:  "Mixed case property - MixedCaseProperty",
			input: "MixedCaseProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "MixedCaseProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"MixedCaseProperty\""},
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
