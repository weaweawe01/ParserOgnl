package test

import (
	"testing"
)

// TestClassMethod 测试类方法调用表达式（基于 Java 的 ClassMethodTest.java）
func TestClassMethod(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Get class name",
			input: "getClass().getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getClass().getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getClass()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		{
			name:  "Get class interfaces",
			input: "getClass().getInterfaces()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getClass().getInterfaces()",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getClass()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getInterfaces()",
					},
				},
			},
		},
		{
			name:  "Get class interfaces length",
			input: "getClass().getInterfaces().length",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getClass().getInterfaces().length",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getClass()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getInterfaces()",
					},
					{
						Type:     "ASTProperty",
						Fragment: "length",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"length\""},
						},
					},
				},
			},
		},
		{
			name:  "System class get interfaces",
			input: "@System@class.getInterfaces()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@System@class.getInterfaces()",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@System@class",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getInterfaces()",
					},
				},
			},
		},
		{
			name:  "Class class get name",
			input: "@Class@class.getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@Class@class.getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@Class@class",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		{
			name:  "ImageObserver class get name",
			input: "@java.awt.image.ImageObserver@class.getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@java.awt.image.ImageObserver@class.getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@java.awt.image.ImageObserver@class",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
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
