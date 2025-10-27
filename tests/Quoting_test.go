package test

import (
	"testing"
)

// TestQuoting 测试引用表达式（基于 Java 的 QuotingTest.java）
func TestQuoting(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Single character quoting - 'c'",
			input: "'c'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'c'",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "Single character quoting - 's'",
			input: "'s'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'s'",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "String quoting - 'string'",
			input: "'string'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"string\"",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "String quoting with double quotes - \"string\"",
			input: `"string"`,
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"string\"",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "String concatenation with empty string - '' + 'bar'",
			input: "'' + 'bar'",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "\"\" + \"bar\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"\""},
					{Type: "ASTConst", Fragment: "\"bar\""},
				},
			},
		},
		{
			name:  "Unicode string quoting - 'yyyy年MM月dd日'",
			input: "'yyyy年MM月dd日'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"yyyy年MM月dd日\"",
				Children: []ExpectedNode{},
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
