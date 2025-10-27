package test

import (
	"testing"
)

// TestVarArgsMethod 测试可变参数方法表达式（基于 Java 的 VarArgsMethodTest.java）
// 测试方法调用中的可变参数（varargs）特性，包括：
// - 无参数的可变参数方法
// - 单个参数的可变参数方法
// - 多个参数的可变参数方法
// - 嵌套调用中的可变参数方法
func TestVarArgsMethod(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		// testNullVarArgs - 测试不传任何参数给可变参数方法
		{
			name:  "null varargs",
			input: "isNullVarArgs()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isNullVarArgs()",
			},
		},
		// testVarArgsWithSingleArg - 测试传递单个参数给可变参数方法
		{
			name:  "varargs with single arg",
			input: "isStringVarArgs(new String())",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isStringVarArgs(new String())",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new String()",
					},
				},
			},
		},
		// testVarArgsWithMultipleArgs - 测试传递多个参数给可变参数方法
		{
			name:  "varargs with multiple args",
			input: "isStringVarArgs(new String(), new String())",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isStringVarArgs(new String(), new String())",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new String()",
					},
					{
						Type:     "ASTCtor",
						Fragment: "new String()",
					},
				},
			},
		},
		// testNestedNullVarArgs - 测试嵌套调用中不传参数的可变参数方法
		{
			name:  "nested null varargs",
			input: "get().request()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "get().request()",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "get()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "request()",
					},
				},
			},
		},
		// testNestedSingleVarArgs - 测试嵌套调用中传递单个参数的可变参数方法
		{
			name:  "nested single varargs",
			input: "get().request(new String())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "get().request(new String())",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "get()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "request(new String())",
						Children: []ExpectedNode{
							{
								Type:     "ASTCtor",
								Fragment: "new String()",
							},
						},
					},
				},
			},
		},
		// testNestedMultipleVarArgs - 测试嵌套调用中传递多个参数的可变参数方法
		{
			name:  "nested multiple varargs",
			input: "get().request(new String(), new String())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "get().request(new String(), new String())",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "get()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "request(new String(), new String())",
						Children: []ExpectedNode{
							{
								Type:     "ASTCtor",
								Fragment: "new String()",
							},
							{
								Type:     "ASTCtor",
								Fragment: "new String()",
							},
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
