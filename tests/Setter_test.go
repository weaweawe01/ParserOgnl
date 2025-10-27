package test

import (
	"testing"
)

// TestSetter 测试设置器表达式（基于 Java 的 SetterTest.java）
func TestSetter(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Map property access - newValue",
			input: "newValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "newValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"newValue\""},
				},
			},
		},
		{
			name:  "List index access - settableList[0]",
			input: "settableList[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "settableList[0]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "settableList",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"settableList\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: " [0]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
		{
			name:  "List special index - settableList[$]",
			input: "settableList[$]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "settableList[$]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "settableList",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"settableList\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[$]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "$"},
						},
					},
				},
			},
		},
		{
			name:  "Map property access - map.newValue",
			input: "map.newValue",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.newValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "map",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"map\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "newValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"newValue\""},
						},
					},
				},
			},
		},
		{
			name:  "Map access - map",
			input: "map",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "map",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"map\""},
				},
			},
		},
		{
			name:  "List access - selectedList",
			input: "selectedList",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "selectedList",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"selectedList\""},
				},
			},
		},
		{
			name:  "Boolean property - openTransitionWin",
			input: "openTransitionWin",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "openTransitionWin",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"openTransitionWin\""},
				},
			},
		},
		{
			name:  "Invalid expression - 0",
			input: "0",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "0",
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
