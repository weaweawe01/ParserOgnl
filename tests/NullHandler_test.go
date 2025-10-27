package test

import (
	"testing"
)

// TestNullHandler 测试 Null 处理器相关的表达式
// 对应 Java 的 NullHandlerTest.java
// 测试不同形式访问属性的表达式解析
func TestNullHandler(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			// stringValue - 简单属性访问
			input: "stringValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "stringValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"stringValue\""},
				},
			},
		},
		{
			// getStringValue() - getter 方法调用
			input: "getStringValue()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getStringValue()",
			},
		},
		{
			// #root.stringValue - 通过 root 变量访问属性
			input: "#root.stringValue",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#root.stringValue",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
					{
						Type:     "ASTProperty",
						Fragment: "stringValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"stringValue\""},
						},
					},
				},
			},
		},
		{
			// #root.getStringValue() - 通过 root 变量调用 getter 方法
			input: "#root.getStringValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#root.getStringValue()",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
					{Type: "ASTMethod", Fragment: "getStringValue()"},
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
