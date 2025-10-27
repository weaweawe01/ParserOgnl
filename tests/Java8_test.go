package test

import (
	"testing"
)

// TestJava8 测试 Java 8 接口默认方法相关的表达式
// 对应 Java 的 Java8Test.java
// 注意: Go 中没有接口默认方法的概念,这里主要测试方法访问表达式的解析
func TestJava8(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Default method call - defaultMethod()",
			input: "defaultMethod()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "defaultMethod()",
			},
		},
		{
			name:  "Name property access - name",
			input: "name",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "name",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"name\""},
				},
			},
		},
		{
			name:  "Get name method call - getName()",
			input: "getName()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getName()",
			},
		},
		{
			name:  "Interface method with parameters - defaultMethod(param)",
			input: "defaultMethod(param)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "defaultMethod(param)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "param",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"param\""},
						},
					},
				},
			},
		},
		{
			name:  "Chain method call - object.defaultMethod()",
			input: "object.defaultMethod()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "object.defaultMethod()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "object",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"object\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "defaultMethod()",
					},
				},
			},
		},
		{
			name:  "Chain property access - object.name",
			input: "object.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "object.name",
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
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
				},
			},
		},
		{
			name:  "Subclass method access - subClass.getName()",
			input: "subClass.getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "subClass.getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "subClass",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"subClass\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		{
			name:  "Interface default method with return - getValue()",
			input: "getValue()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getValue()",
			},
		},
		{
			name:  "Multiple parameters - process(x, y)",
			input: "process(x, y)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "process(x, y)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "x",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"x\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "y",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"y\""},
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
