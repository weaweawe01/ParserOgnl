package test

import (
	"testing"
)

// TestOperator 测试运算符表达式
// 对应 Java 的 OperatorTest.java
// 主要测试各种比较运算符和它们的文本形式（如 gt, gte, lt, lte, eq）
func TestOperator(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		// testStringComparisons - 字符串比较运算符
		{
			input: "\"one\" > \"two\"",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "\"one\" > \"two\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"one\""},
					{Type: "ASTConst", Fragment: "\"two\""},
				},
			},
		},
		{
			input: "\"one\" >= \"two\"",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "\"one\" >= \"two\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"one\""},
					{Type: "ASTConst", Fragment: "\"two\""},
				},
			},
		},
		{
			input: "\"one\" < \"two\"",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "\"one\" < \"two\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"one\""},
					{Type: "ASTConst", Fragment: "\"two\""},
				},
			},
		},
		{
			input: "\"one\" <= \"two\"",
			expected: ExpectedNode{
				Type:     "ASTLessEq",
				Fragment: "\"one\" <= \"two\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"one\""},
					{Type: "ASTConst", Fragment: "\"two\""},
				},
			},
		},
		{
			input: "\"one\" == \"two\"",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "\"one\" == \"two\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"one\""},
					{Type: "ASTConst", Fragment: "\"two\""},
				},
			},
		},
		{
			input: "\"o\" > \"o\"",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "\"o\" > \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" gt \"o\"",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "\"o\" > \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" >= \"o\"",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "\"o\" >= \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" gte \"o\"",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "\"o\" >= \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" < \"o\"",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "\"o\" < \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" lt \"o\"",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "\"o\" < \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" <= \"o\"",
			expected: ExpectedNode{
				Type:     "ASTLessEq",
				Fragment: "\"o\" <= \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" lte \"o\"",
			expected: ExpectedNode{
				Type:     "ASTLessEq",
				Fragment: "\"o\" <= \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" == \"o\"",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "\"o\" == \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		{
			input: "\"o\" eq \"o\"",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "\"o\" == \"o\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"o\""},
					{Type: "ASTConst", Fragment: "\"o\""},
				},
			},
		},
		// 额外测试：不等于运算符
		{
			input: "a != b",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "a != b",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "b",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"b\""},
						},
					},
				},
			},
		},
		{
			input: "a neq b",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "a != b",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "b",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"b\""},
						},
					},
				},
			},
		},
		// 额外测试：数字比较
		{
			input: "5 > 3",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "5 > 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "10 lt 20",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "10 < 20",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "10"},
					{Type: "ASTConst", Fragment: "20"},
				},
			},
		},
		// 额外测试：逻辑运算符
		{
			input: "a && b",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "a && b",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "b",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"b\""},
						},
					},
				},
			},
		},
		{
			input: "a and b",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "a && b",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "b",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"b\""},
						},
					},
				},
			},
		},
		{
			input: "a || b",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "a || b",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "b",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"b\""},
						},
					},
				},
			},
		},
		{
			input: "a or b",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "a || b",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "b",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"b\""},
						},
					},
				},
			},
		},
		{
			input: "!a",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!a",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
				},
			},
		},
		{
			input: "not a",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!a",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "a",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"a\""},
						},
					},
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
