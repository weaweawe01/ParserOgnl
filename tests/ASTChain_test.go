package test

import (
	"testing"
)

// TestArrayCreation 测试数组创建表达式
func TestASTChain(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "thing[\"x\"].val",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "thing[\"x\"].val",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "thing",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"thing\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"x\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"x\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "val",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"val\""},
						},
					},
				},
			},
		},
	}
	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
