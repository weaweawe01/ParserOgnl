package test

import (
	"testing"
)

// TestNumericConversion 测试数字类型转换相关的表达式
// 对应 Java 的 NumericConversionTest.java
// 这个测试主要验证各种数字字面量的解析
func TestNumericConversion(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		// 整数字面量
		{
			input: "55",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55",
			},
		},
		{
			input: "55L",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55L",
			},
		},
		// 浮点数字面量
		{
			input: "55.0",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.0",
			},
		},
		{
			input: "55.1234",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234",
			},
		},
		{
			input: "55.1234f",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234",
			},
		},
		{
			input: "55.1234F",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234",
			},
		},
		{
			input: "55.1234d",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234",
			},
		},
		{
			input: "55.1234D",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234",
			},
		},
		{
			input: "55.",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.0",
			},
		},
		{
			input: "2.",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "2.0",
			},
		},
		// 布尔字面量
		{
			input: "true",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "true",
			},
		},
		{
			input: "false",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "false",
			},
		},
		// BigInteger 和 BigDecimal 字面量（带后缀）
		{
			input: "55h",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55H",
			},
		},
		{
			input: "55H",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55H",
			},
		},
		{
			input: "55.1234b",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234B",
			},
		},
		{
			input: "55.1234B",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.1234B",
			},
		},
		// 字符字面量
		{
			input: "'A'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'A'",
			},
		},
		{
			input: "'7'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'7'",
			},
		},
		// 字节字面量
		{
			input: "55",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55",
			},
		},
		// 字符串字面量（用于转换）
		{
			input: "\"55\"",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"55\"",
			},
		},
		{
			input: "\"55.1234\"",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"55.1234\"",
			},
		},
		{
			input: "\"true\"",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"true\"",
			},
		},
		// 组合表达式：类型转换的常见场景
		{
			input: "1 + 2.0",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "1 + 2.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "2.0"},
				},
			},
		},
		{
			input: "55L + 1",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "55L + 1",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "55L"},
					{Type: "ASTConst", Fragment: "1"},
				},
			},
		},
		{
			input: "55.0f + 1.0",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "55.0 + 1.0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "55.0"},
					{Type: "ASTConst", Fragment: "1.0"},
				},
			},
		},
		{
			input: "true ? 1 : 0",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "true ? 1 : 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		// 负数字面量
		{
			input: "-55",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-55",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "55"},
				},
			},
		},
		{
			input: "-55.1234",
			expected: ExpectedNode{
				Type:     "ASTNegate",
				Fragment: "-55.1234",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "55.1234"},
				},
			},
		},
		{
			input: "+55",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55",
			},
		},
		{
			input: "+55.0",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "55.0",
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
