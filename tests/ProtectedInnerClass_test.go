package test

import (
	"testing"
)

// TestProtectedInnerClass 测试受保护内部类表达式（基于 Java 的 ProtectedInnerClassTest.java）
func TestProtectedInnerClass(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "List size method call - list.size()",
			input: "list.size()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list.size()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{Type: "ASTMethod", Fragment: "size()"},
				},
			},
		},
		{
			name:  "List element access - list[0]",
			input: "list[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[0]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
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
