package test

import (
	"testing"
)

// TestSimpleNavigationChainTree 测试简单导航链树（基于 Java 的 SimpleNavigationChainTreeTest.java）
// 这个测试验证表达式的 AST 结构，用于判断是否为"简单导航链"
// 简单导航链定义：只包含属性访问的链式表达式（如 name.foo.bar）
// 不是简单导航链的情况：包含索引访问、算术运算等复杂操作
func TestSimpleNavigationChainTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		// testName - 测试单个属性名（是简单导航链）
		{
			name:  "simple property name",
			input: "name",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "name",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"name\""},
				},
			},
		},
		// testNameWithIndex - 测试带索引的属性访问（不是简单导航链）
		{
			name:  "property with index",
			input: "name[i]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name[i]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[i]",
						Children: []ExpectedNode{
							{Type: "ASTProperty", Fragment: "i",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"i\""},
								},
							},
						},
					},
				},
			},
		},
		// testNameWithAddition - 测试带加法运算的表达式（不是简单导航链）
		{
			name:  "property with addition",
			input: "name + foo",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "name + foo",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "foo",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"foo\""},
						},
					},
				},
			},
		},
		// testNameWithProperty - 测试链式属性访问（是简单导航链）
		{
			name:  "chained property access",
			input: "name.foo",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "name.foo",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "foo",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"foo\""},
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
