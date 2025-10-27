package test

import (
	"testing"
)

// TestNumberFormatException 测试数字格式异常相关的表达式解析（基于 Java 的 NumberFormatExceptionTest.java）
func TestNumberFormatException(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Float value assignment - floatValue = 10f",
			input: "floatValue = 10f",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "floatValue = 10.0",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "floatValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"floatValue\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "10.0",
					},
				},
			},
		},
		{
			name:  "Float value access - floatValue",
			input: "floatValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "floatValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "\"floatValue\"",
					},
				},
			},
		},
		{
			name:  "Int value assignment - intValue = 34",
			input: "intValue = 34",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "intValue = 34",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "intValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"intValue\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "34",
					},
				},
			},
		},
		{
			name:  "Int value access - intValue",
			input: "intValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "intValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "\"intValue\"",
					},
				},
			},
		},
		{
			name:  "BigInt value assignment - bigIntValue = 34",
			input: "bigIntValue = 34",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "bigIntValue = 34",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bigIntValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"bigIntValue\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "34",
					},
				},
			},
		},
		{
			name:  "BigInt value access - bigIntValue",
			input: "bigIntValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "bigIntValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "\"bigIntValue\"",
					},
				},
			},
		},
		{
			name:  "BigInt null assignment - bigIntValue = null",
			input: "bigIntValue = null",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "bigIntValue = null",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bigIntValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"bigIntValue\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "null",
					},
				},
			},
		},
		{
			name:  "BigDec value assignment - bigDecValue = 34.55",
			input: "bigDecValue = 34.55",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "bigDecValue = 34.55",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bigDecValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"bigDecValue\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "34.55",
					},
				},
			},
		},
		{
			name:  "BigDec value access - bigDecValue",
			input: "bigDecValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "bigDecValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "\"bigDecValue\"",
					},
				},
			},
		},
		{
			name:  "BigDec null assignment - bigDecValue = null",
			input: "bigDecValue = null",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "bigDecValue = null",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bigDecValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"bigDecValue\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "null",
					},
				},
			},
		},
		{
			name:  "String to number conversion - value = \"123\"",
			input: `value = "123"`,
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: `value = "123"`,
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: `"123"`,
					},
				},
			},
		},
		{
			name:  "Invalid string to number - value = \"foobar\"",
			input: `value = "foobar"`,
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: `value = "foobar"`,
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: `"foobar"`,
					},
				},
			},
		},
		{
			name:  "Empty string to number - value = \"\"",
			input: `value = ""`,
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: `value = ""`,
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: `""`,
					},
				},
			},
		},
		{
			name:  "Whitespace string to number - value = \"   \\t\"",
			input: `value = "   \t"`,
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "value = \"   \t\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: `"   \t"`,
					},
				},
			},
		},
		{
			name:  "Valid whitespace string - value = \"   \\t1234\\t\\t\"",
			input: `value = "   \t1234\t\t"`,
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: `value = "   \t1234\t\t"`,
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: `"   \t1234\t\t"`,
					},
				},
			},
		},
		{
			name:  "Float literal with suffix - value = 10f",
			input: "value = 10f",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "value = 10.0",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "10.0",
					},
				},
			},
		},
		{
			name:  "Invalid float format - value = \"x10x\"",
			input: `value = "x10x"`,
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: `value = "x10x"`,
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"value\"",
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: `"x10x"`,
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
