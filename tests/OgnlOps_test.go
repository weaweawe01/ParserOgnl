package test

import (
	"testing"
)

// TestOgnlOps 测试 OGNL 操作符和运算
// 对应 Java 的 OgnlOpsTest.java
// 测试各种 OGNL 操作符和运算表达式
func TestOgnlOps(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		// 相等比较操作
		{
			name:  "Equal - string comparison",
			input: "\"hello\" == \"hello\"",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "\"hello\" == \"hello\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"hello\""},
					{Type: "ASTConst", Fragment: "\"hello\""},
				},
			},
		},
		{
			name:  "Equal - float comparison",
			input: "1.5 == 1.5",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "1.5 == 1.5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1.5"},
					{Type: "ASTConst", Fragment: "1.5"},
				},
			},
		},
		{
			name:  "Equal - integer comparison",
			input: "42 == 42",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "42 == 42",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "42"},
					{Type: "ASTConst", Fragment: "42"},
				},
			},
		},
		{
			name:  "Equal - null comparison",
			input: "null == null",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "null == null",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "null"},
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},

		// 位运算操作
		{
			name:  "Shift left",
			input: "8 << 2",
			expected: ExpectedNode{
				Type:     "ASTShiftLeft",
				Fragment: "8 << 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "8"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			name:  "Shift right",
			input: "8 >> 2",
			expected: ExpectedNode{
				Type:     "ASTShiftRight",
				Fragment: "8 >> 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "8"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			name:  "Unsigned shift right",
			input: "8 >>> 2",
			expected: ExpectedNode{
				Type:     "ASTUnsignedShiftRight",
				Fragment: "8 >>> 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "8"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},

		// 算术运算操作
		{
			name:  "Add - integers",
			input: "5 + 3",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5 + 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			name:  "Subtract - integers",
			input: "10 - 4",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "10 - 4",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "10"},
					{Type: "ASTConst", Fragment: "4"},
				},
			},
		},
		{
			name:  "Multiply - integers",
			input: "6 * 7",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "6 * 7",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "6"},
					{Type: "ASTConst", Fragment: "7"},
				},
			},
		},
		{
			name:  "Divide - integers",
			input: "20 / 4",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "20 / 4",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "20"},
					{Type: "ASTConst", Fragment: "4"},
				},
			},
		},
		{
			name:  "Remainder - integers",
			input: "17 % 5",
			expected: ExpectedNode{
				Type:     "ASTRemainder",
				Fragment: "17 % 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "17"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},

		// 一元运算操作
		{
			name:  "Negate - positive number",
			input: "-42",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-42",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "42"},
				},
			},
		},
		{
			name:  "Negate - negative number",
			input: "-(-10)",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "--10",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-10",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "10"},
						},
					},
				},
			},
		},
		{
			name:  "Bitwise NOT",
			input: "~15",
			expected: ExpectedNode{
				Type:     "ASTBitNegate",
				Fragment: "~15",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "15"},
				},
			},
		},

		// 逻辑运算操作
		{
			name:  "Logical AND",
			input: "true && false",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "true && false",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "false"},
				},
			},
		},
		{
			name:  "Logical OR",
			input: "true || false",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "true || false",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "false"},
				},
			},
		},

		// 比较运算操作
		{
			name:  "Less than",
			input: "3 < 7",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "3 < 7",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "3"},
					{Type: "ASTConst", Fragment: "7"},
				},
			},
		},
		{
			name:  "Greater than",
			input: "8 > 3",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "8 > 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "8"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			name:  "Less than or equal",
			input: "5 <= 5",
			expected: ExpectedNode{
				Type:     "ASTLessEq",
				Fragment: "5 <= 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			name:  "Greater than or equal",
			input: "7 >= 7",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "7 >= 7",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "7"},
					{Type: "ASTConst", Fragment: "7"},
				},
			},
		},
		{
			name:  "Not equal",
			input: "4 != 5",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "4 != 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "4"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},

		// 位运算操作
		{
			name:  "Bitwise AND",
			input: "12 & 7",
			expected: ExpectedNode{
				Type:     "ASTBitAnd",
				Fragment: "12 & 7",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "12"},
					{Type: "ASTConst", Fragment: "7"},
				},
			},
		},
		{
			name:  "Bitwise XOR",
			input: "12 ^ 7",
			expected: ExpectedNode{
				Type:     "ASTXor",
				Fragment: "12 ^ 7",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "12"},
					{Type: "ASTConst", Fragment: "7"},
				},
			},
		},
		{
			name:  "Bitwise OR",
			input: "12 | 7",
			expected: ExpectedNode{
				Type:     "ASTBitOr",
				Fragment: "12 | 7",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "12"},
					{Type: "ASTConst", Fragment: "7"},
				},
			},
		},

		// 字符串操作
		{
			name:  "String concatenation",
			input: "\"hello\" + \" world\"",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "\"hello\" + \" world\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"hello\""},
					{Type: "ASTConst", Fragment: "\" world\""},
				},
			},
		},

		// 复杂表达式
		{
			name:  "Complex arithmetic expression",
			input: "(2 + 3) * 4 - 1",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "((2 + 3) * 4) - 1",
				Children: []ExpectedNode{
					{
						Type:     "ASTMultiply",
						Fragment: "((2 + 3) * 4)",
						Children: []ExpectedNode{
							{
								Type:     "ASTAdd",
								Fragment: "(2 + 3)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "2"},
									{Type: "ASTConst", Fragment: "3"},
								},
							},
							{Type: "ASTConst", Fragment: "4"},
						},
					},
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			name:  "Complex logical expression",
			input: "(x > 5) && (y < 10) || z == 0",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "((x > 5) && (y < 10)) || (z == 0)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAnd",
						Fragment: "((x > 5) && (y < 10))",
						Children: []ExpectedNode{
							{
								Type:     "ASTGreater",
								Fragment: "(x > 5)",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "x",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"x\""},
										},
									},
									{Type: "ASTConst", Fragment: "5"},
								},
							},
							{
								Type:     "ASTLess",
								Fragment: "(y < 10)",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "y",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"y\""},
										},
									},
									{Type: "ASTConst", Fragment: "10"},
								},
							},
						},
					},
					{
						Type:     "ASTEq",
						Fragment: "z == 0",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "z",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"z\""},
								},
							},
							{Type: "ASTConst", Fragment: "0"},
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
