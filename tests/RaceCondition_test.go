package test

import (
	"testing"
)

// TestRaceCondition 测试竞态条件表达式（基于 Java 的 RaceConditionTest.java）
func TestRaceCondition(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Method call - execute()",
			input: "execute()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "execute",
			},
		},
		{
			name:  "Method call with class - TestAction.execute()",
			input: "TestAction.execute()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "TestAction.execute()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "TestAction",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"TestAction\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "execute()",
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
