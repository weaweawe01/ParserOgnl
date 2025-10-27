package test

import (
	"testing"
)

// TestShortCircuitingExpression 测试短路表达式（基于 Java 的 ShortCircuitingExpressionTest.java）
func TestShortCircuitingExpression(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "AND short circuit - false && true",
			input: "false && true",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "false && true",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "false"},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
		{
			name:  "AND short circuit - true && false",
			input: "true && false",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "true && false",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "false"},
				},
			},
		},
		{
			name:  "OR short circuit - false || true",
			input: "false || true",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "false || true",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "false"},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
		{
			name:  "OR short circuit - true || false",
			input: "true || false",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "true || false",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "false"},
				},
			},
		},
		{
			name:  "Mixed short circuit - false && (true || false)",
			input: "false && (true || false)",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "false && (true || false)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "false"},
					{
						Type:     "ASTOr",
						Fragment: "true || false",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "true"},
							{Type: "ASTConst", Fragment: "false"},
						},
					},
				},
			},
		},
		{
			name:  "Mixed short circuit - true || (false && true)",
			input: "true || (false && true)",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "true || (false && true)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{
						Type:     "ASTAnd",
						Fragment: "false && true",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "false"},
							{Type: "ASTConst", Fragment: "true"},
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
