package test

import (
	"testing"
)

// TestInExpression 测试 in 表达式
// 对应 Java 的 InExpressionTest.java (OGNL-118)
func TestInExpression(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "String in list - #name in {\"Greenland\", \"Austin\", \"Africa\", \"Rome\"}",
			input: "#name in {\"Greenland\", \"Austin\", \"Africa\", \"Rome\"}",
			expected: ExpectedNode{
				Type:     "ASTIn",
				Fragment: "#name in { \"Greenland\", \"Austin\", \"Africa\", \"Rome\" }",
				Children: []ExpectedNode{
					{
						Type:     "ASTVarRef",
						Fragment: "#name",
					},
					{
						Type:     "ASTList",
						Fragment: "{ \"Greenland\", \"Austin\", \"Africa\", \"Rome\" }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"Greenland\""},
							{Type: "ASTConst", Fragment: "\"Austin\""},
							{Type: "ASTConst", Fragment: "\"Africa\""},
							{Type: "ASTConst", Fragment: "\"Rome\""},
						},
					},
				},
			},
		},
		{
			name:  "Number in list - 5 in {1, 3, 5, 7, 9}",
			input: "5 in {1, 3, 5, 7, 9}",
			expected: ExpectedNode{
				Type:     "ASTIn",
				Fragment: "5 in { 1, 3, 5, 7, 9 }",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{
						Type:     "ASTList",
						Fragment: "{ 1, 3, 5, 7, 9 }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1"},
							{Type: "ASTConst", Fragment: "3"},
							{Type: "ASTConst", Fragment: "5"},
							{Type: "ASTConst", Fragment: "7"},
							{Type: "ASTConst", Fragment: "9"},
						},
					},
				},
			},
		},
		{
			name:  "Property in list - name in {\"John\", \"Jane\", \"Bob\"}",
			input: "name in {\"John\", \"Jane\", \"Bob\"}",
			expected: ExpectedNode{
				Type:     "ASTIn",
				Fragment: "name in { \"John\", \"Jane\", \"Bob\" }",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTList",
						Fragment: "{ \"John\", \"Jane\", \"Bob\" }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"John\""},
							{Type: "ASTConst", Fragment: "\"Jane\""},
							{Type: "ASTConst", Fragment: "\"Bob\""},
						},
					},
				},
			},
		},
		{
			name:  "Not in expression - #name not in {\"A\", \"B\", \"C\"}",
			input: "#name not in { \"A\", \"B\", \"C\" }",
			expected: ExpectedNode{
				Type:     "ASTNotIn",
				Fragment: "#name not in { \"A\", \"B\", \"C\" }",
				Children: []ExpectedNode{
					{
						Type:     "ASTVarRef",
						Fragment: "#name",
					},
					{
						Type:     "ASTList",
						Fragment: "{ \"A\", \"B\", \"C\" }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"A\""},
							{Type: "ASTConst", Fragment: "\"B\""},
							{Type: "ASTConst", Fragment: "\"C\""},
						},
					},
				},
			},
		},
		{
			name:  "Expression in array - (x + y) in {10, 20, 30}",
			input: "(x + y) in {10, 20, 30}",
			expected: ExpectedNode{
				Type:     "ASTIn",
				Fragment: "(x + y) in { 10, 20, 30 }",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "x + y",
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
					{
						Type:     "ASTList",
						Fragment: "{ 10, 20, 30 }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "10"},
							{Type: "ASTConst", Fragment: "20"},
							{Type: "ASTConst", Fragment: "30"},
						},
					},
				},
			},
		},
		{
			name:  "String in empty list - name in {}",
			input: "name in {}",
			expected: ExpectedNode{
				Type:     "ASTIn",
				Fragment: "name in {  }",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
					{
						Type:     "ASTList",
						Fragment: "{  }",
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
