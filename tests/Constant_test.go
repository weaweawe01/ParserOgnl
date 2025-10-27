package test

import (
	"testing"
)

// TestConstant 测试常量表达式（基于 Java 的 ConstantTest.java）
func TestConstant(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Integer 12345",
			input: "12345",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "12345",
			},
		},
		{
			name:  "Hexadecimal 0x100",
			input: "0x100",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "256",
			},
		},
		{
			name:  "Hexadecimal 0xfE",
			input: "0xfE",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "254",
			},
		},
		{
			name:  "Octal 01000",
			input: "01000",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "512",
			},
		},
		{
			name:  "Long 1234L",
			input: "1234L",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "1234L",
			},
		},
		{
			name:  "Float 12.34",
			input: "12.34",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "12.34",
			},
		},
		{
			name:  "Float .1234",
			input: ".1234",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "0.1234",
			},
		},
		{
			name:  "Float 12.34f",
			input: "12.34f",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "12.34",
			},
		},
		{
			name:  "Float 12.",
			input: "12.",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "12.0",
			},
		},
		{
			name:  "Scientific notation 12e+1d",
			input: "12e+1d",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "120.0",
			},
		},
		{
			name:  "Character 'x'",
			input: "'x'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'x'",
			},
		},
		{
			name:  "Character newline '\\n'",
			input: "'\\n'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'\n'",
			},
		},
		{
			name:  "Unicode character '\\u048c'",
			input: "'\\u048c'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'Ҍ'",
			},
		},
		{
			name:  "Octal character '\\47'",
			input: "'\\47'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'\\47'",
			},
		},
		{
			name:  "Octal character '\\367'",
			input: "'\\367'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "'\\367'",
			},
		},
		{
			name:  "String hello world",
			input: "\"hello world\"",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"hello world\"",
			},
		},
		{
			name:  "Unicode string with escapes",
			input: "\"\\u00a0\\u0068ell\\'o\\\\\\n\\r\\f\\t\\b\\\"\\167orld\\\"\"",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"\\u00a0\\u0068ell\\'o\\\\\\n\\r\\f\\t\\b\\\"\\167orld\\\"\"",
			},
		},
		{
			name:  "Null constant",
			input: "null",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "null",
			},
		},
		{
			name:  "Boolean true",
			input: "true",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "true",
			},
		},
		{
			name:  "Boolean false",
			input: "false",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "false",
			},
		},
		{
			name:  "Array literal with mixed types",
			input: "{ false, true, null, 0, 1. }",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ false, true, null, 0, 1.0 }",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "false"},
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "null"},
					{Type: "ASTConst", Fragment: "0"},
					{Type: "ASTConst", Fragment: "1.0"},
				},
			},
		},
		{
			name:  "HTML public string",
			input: "'HTML PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\"'",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "\"HTML PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\"",
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
