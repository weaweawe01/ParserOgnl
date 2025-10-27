package test

import (
	"testing"
)

// 辅助函数: 解析 OGNL 表达式
// func parseExpression(t *testing.T, input string) ast.Expression {
// 	l := lexer.NewLexer(input)
// 	p := parser.New(l)
// 	expr, err := p.ParseTopLevelExpression()

// 	if err != nil {
// 		t.Fatalf("解析错误: %v", err)
// 	}

// 	if len(p.Errors()) > 0 {
// 		t.Fatalf("解析器错误: %v", p.Errors())
// 	}

// 	if expr == nil {
// 		t.Fatal("表达式解析结果为 nil")
// 	}

// 	return expr
// }

// ============================================================================
// Double 值的算术表达式测试
// ============================================================================

func TestDoubleValuedArithmeticExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "-1d",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-1.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1.0"},
				},
			},
		},
		{
			input: "+1d",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "1.0",
			},
		},
		{
			input: "--1f",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "--1.0",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-1.0",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1.0"},
						},
					},
				},
			},
		},
		{
			input: "2*2.0",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "2 * 2.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "2"},
					{Type: "ASTConst", Fragment: "2.0"},
				},
			},
		},
		{
			input: "5/2.",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "5 / 2.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2.0"},
				},
			},
		},
		{
			input: "5+2D",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5 + 2.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2.0"},
				},
			},
		},
		{
			input: "5f-2F",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "5.0 - 2.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5.0"},
					{Type: "ASTConst", Fragment: "2.0"},
				},
			},
		},
		{
			input: "5.+2*3",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5.0 + (2 * 3)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5.0"},
					{
						Type:     "ASTMultiply",
						Fragment: "(2 * 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			input: "(5.+2)*3",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "(5.0 + 2) * 3",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "(5.0 + 2)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5.0"},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// ============================================================================
// BigDecimal 值的算术表达式测试
// ============================================================================

func TestBigDecimalValuedArithmeticExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "-1b",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-1B",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1B"},
				},
			},
		},
		{
			input: "+1b",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "1B",
			},
		},
		{
			input: "--1b",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "--1B",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-1B",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1B"},
						},
					},
				},
			},
		},
		{
			input: "2*2.0b",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "2 * 2.0B",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "2"},
					{Type: "ASTConst", Fragment: "2.0B"},
				},
			},
		},
		{
			input: "5/2.B",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "5 / 2B",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2B"},
				},
			},
		},
		{
			input: "5.0B/2",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "5.0B / 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5.0B"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5+2b",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5 + 2B",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2B"},
				},
			},
		},
		{
			input: "5-2B",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "5 - 2B",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2B"},
				},
			},
		},
		{
			input: "5.+2b*3",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5.0 + (2B * 3)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5.0"},
					{
						Type:     "ASTMultiply",
						Fragment: "(2B * 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2B"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			input: "(5.+2b)*3",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "(5.0 + 2B) * 3",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "(5.0 + 2B)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5.0"},
							{Type: "ASTConst", Fragment: "2B"},
						},
					},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// ============================================================================
// Integer 值的算术表达式测试
// ============================================================================

func TestIntegerValuedArithmeticExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "-1",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "+1",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "1",
			},
		},
		{
			input: "--1",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "--1",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-1",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1"},
						},
					},
				},
			},
		},
		{
			input: "2*2",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "2 * 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "2"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5/2",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "5 / 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5+2",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5 + 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5-2",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "5 - 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5+2*3",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5 + (2 * 3)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{
						Type:     "ASTMultiply",
						Fragment: "(2 * 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			input: "(5+2)*3",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "(5 + 2) * 3",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "(5 + 2)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5"},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "~1",
			expected: ExpectedNode{
				Type:     "ASTBitNegate",
				Fragment: "~1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "5%2",
			expected: ExpectedNode{
				Type:     "ASTRemainder",
				Fragment: "5 % 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5<<2",
			expected: ExpectedNode{
				Type:     "ASTShiftLeft",
				Fragment: "5 << 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5>>2",
			expected: ExpectedNode{
				Type:     "ASTShiftRight",
				Fragment: "5 >> 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5>>1+1",
			expected: ExpectedNode{
				Type:     "ASTShiftRight",
				Fragment: "5 >> (1 + 1)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{
						Type:     "ASTAdd",
						Fragment: "(1 + 1)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1"},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
				},
			},
		},
		{
			input: "-5>>>2",
			expected: ExpectedNode{
				Type:     "ASTUnsignedShiftRight",
				Fragment: "-5 >>> 2",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-5",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5"},
						},
					},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "-5L>>>2",
			expected: ExpectedNode{
				Type:     "ASTUnsignedShiftRight",
				Fragment: "-5L >>> 2",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-5L",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5L"},
						},
					},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5. & 3",
			expected: ExpectedNode{
				Type:     "ASTBitAnd",
				Fragment: "5.0 & 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5.0"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "5 ^3",
			expected: ExpectedNode{
				Type:     "ASTXor",
				Fragment: "5 ^ 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "5l&3|5^3",
			expected: ExpectedNode{
				Type:     "ASTBitOr",
				Fragment: "(5L & 3) | (5 ^ 3)",
				Children: []ExpectedNode{
					{
						Type:     "ASTBitAnd",
						Fragment: "(5L & 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5L"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
					{
						Type:     "ASTXor",
						Fragment: "(5 ^ 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			input: "5&(3|5^3)",
			expected: ExpectedNode{
				Type:     "ASTBitAnd",
				Fragment: "5 & (3 | (5 ^ 3))",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{
						Type:     "ASTBitOr",
						Fragment: "(3 | (5 ^ 3))",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "3"},
							{
								Type:     "ASTXor",
								Fragment: "(5 ^ 3)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "5"},
									{Type: "ASTConst", Fragment: "3"},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "true ? 1 : 1/0",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "true ? 1 : 1 / 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "1"},
					{
						Type:     "ASTDivide",
						Fragment: "1 / 0",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1"},
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// ============================================================================
// BigInteger 值的算术表达式测试
// ============================================================================

func TestBigIntegerValuedArithmeticExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "-1h",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-1H",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1H"},
				},
			},
		},
		{
			input: "+1H",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "1H",
			},
		},
		{
			input: "--1h",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "--1H",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-1H",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1H"},
						},
					},
				},
			},
		},
		{
			input: "2h*2",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "2H * 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "2H"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5/2h",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "5 / 2H",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2H"},
				},
			},
		},
		{
			input: "5h+2",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5H + 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5-2h",
			expected: ExpectedNode{
				Type:     "ASTSubtract",
				Fragment: "5 - 2H",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2H"},
				},
			},
		},
		{
			input: "5+2H*3",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "5 + (2H * 3)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{
						Type:     "ASTMultiply",
						Fragment: "(2H * 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2H"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			input: "(5+2H)*3",
			expected: ExpectedNode{
				Type:     "ASTMultiply",
				Fragment: "(5 + 2H) * 3",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "(5 + 2H)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5"},
							{Type: "ASTConst", Fragment: "2H"},
						},
					},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "~1h",
			expected: ExpectedNode{
				Type:     "ASTBitNegate",
				Fragment: "~1H",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1H"},
				},
			},
		},
		{
			input: "5h%2",
			expected: ExpectedNode{
				Type:     "ASTRemainder",
				Fragment: "5H % 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5h<<2",
			expected: ExpectedNode{
				Type:     "ASTShiftLeft",
				Fragment: "5H << 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5h>>2",
			expected: ExpectedNode{
				Type:     "ASTShiftRight",
				Fragment: "5H >> 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5h>>1+1",
			expected: ExpectedNode{
				Type:     "ASTShiftRight",
				Fragment: "5H >> (1 + 1)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{
						Type:     "ASTAdd",
						Fragment: "(1 + 1)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1"},
							{Type: "ASTConst", Fragment: "1"},
						},
					},
				},
			},
		},
		{
			input: "-5h>>>2",
			expected: ExpectedNode{
				Type:     "ASTUnsignedShiftRight",
				Fragment: "-5H >>> 2",
				Children: []ExpectedNode{
					{
						Type:     "ASTNegate",
						Fragment: "-5H",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5H"},
						},
					},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5.b & 3",
			expected: ExpectedNode{
				Type:     "ASTBitAnd",
				Fragment: "5B & 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5B"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "5h ^3",
			expected: ExpectedNode{
				Type:     "ASTXor",
				Fragment: "5H ^ 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "5h&3|5^3",
			expected: ExpectedNode{
				Type:     "ASTBitOr",
				Fragment: "(5H & 3) | (5 ^ 3)",
				Children: []ExpectedNode{
					{
						Type:     "ASTBitAnd",
						Fragment: "(5H & 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5H"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
					{
						Type:     "ASTXor",
						Fragment: "(5 ^ 3)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5"},
							{Type: "ASTConst", Fragment: "3"},
						},
					},
				},
			},
		},
		{
			input: "5H&(3|5^3)",
			expected: ExpectedNode{
				Type:     "ASTBitAnd",
				Fragment: "5H & (3 | (5 ^ 3))",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5H"},
					{
						Type:     "ASTBitOr",
						Fragment: "(3 | (5 ^ 3))",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "3"},
							{
								Type:     "ASTXor",
								Fragment: "(5 ^ 3)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "5"},
									{Type: "ASTConst", Fragment: "3"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// ============================================================================
// 逻辑表达式测试
// ============================================================================

func TestLogicalExpressions(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "!1",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "!null",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!null",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		{
			input: "5<2",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "5 < 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5>2",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "5 > 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "5<=5",
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
			input: "5>=3",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "5 >= 3",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "3"},
				},
			},
		},
		{
			input: "5<-5>>>2",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "5 < (-5 >>> 2)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{
						Type:     "ASTUnsignedShiftRight",
						Fragment: "(-5 >>> 2)",
						Children: []ExpectedNode{
							{
								Type:     "ASTNegate",
								Fragment: "-5",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "5"},
								},
							},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
				},
			},
		},
		{
			input: "5==5.0",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "5 == 5.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "5.0"},
				},
			},
		},
		{
			input: "5!=5.0",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "5 != 5.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTConst", Fragment: "5.0"},
				},
			},
		},
		{
			input: "!false || true",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "!false || true",
				Children: []ExpectedNode{
					{
						Type:     "ASTNot",
						Fragment: "!false",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "false"},
						},
					},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
		{
			input: "!(true && true)",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!(true && true)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAnd",
						Fragment: "true && true",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "true"},
							{Type: "ASTConst", Fragment: "true"},
						},
					},
				},
			},
		},
		{
			input: "(1 > 0 && true) || 2 > 0",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "((1 > 0) && true) || (2 > 0)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAnd",
						Fragment: "((1 > 0) && true)",
						Children: []ExpectedNode{
							{
								Type:     "ASTGreater",
								Fragment: "(1 > 0)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "1"},
									{Type: "ASTConst", Fragment: "0"},
								},
							},
							{Type: "ASTConst", Fragment: "true"},
						},
					},
					{
						Type:     "ASTGreater",
						Fragment: "(2 > 0)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2"},
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// ============================================================================
// 逻辑表达式字符串版本测试 (or, and, eq, neq, lt, gt, etc.)
// ============================================================================

func TestLogicalExpressionsStringVersions(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "2 or 0",
			expected: ExpectedNode{
				Type:     "ASTOr",
				Fragment: "2 || 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "2"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "1 and 0",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "1 && 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "1 bor 0",
			expected: ExpectedNode{
				Type:     "ASTBitOr",
				Fragment: "1 | 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "true && 12",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "true && 12",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "12"},
				},
			},
		},
		{
			input: "1 xor 0",
			expected: ExpectedNode{
				Type:     "ASTXor",
				Fragment: "1 ^ 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "1 band 0",
			expected: ExpectedNode{
				Type:     "ASTBitAnd",
				Fragment: "1 & 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			input: "1 eq 1",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "1 == 1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "1 neq 1",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "1 != 1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "1 lt 5",
			expected: ExpectedNode{
				Type:     "ASTLess",
				Fragment: "1 < 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			input: "1 lte 5",
			expected: ExpectedNode{
				Type:     "ASTLessEq",
				Fragment: "1 <= 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			input: "1 gt 5",
			expected: ExpectedNode{
				Type:     "ASTGreater",
				Fragment: "1 > 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			input: "1 gte 5",
			expected: ExpectedNode{
				Type:     "ASTGreaterEq",
				Fragment: "1 >= 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			input: "1 shl 2",
			expected: ExpectedNode{
				Type:     "ASTShiftLeft",
				Fragment: "1 << 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "4 shr 2",
			expected: ExpectedNode{
				Type:     "ASTShiftRight",
				Fragment: "4 >> 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "4"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "4 ushr 2",
			expected: ExpectedNode{
				Type:     "ASTUnsignedShiftRight",
				Fragment: "4 >>> 2",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "4"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			input: "not null",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!null",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		{
			input: "not 1",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
