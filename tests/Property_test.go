package test

import (
	"testing"
)

// TestProperty 测试属性访问表达式
func TestProperty(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{},
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
