package test

import (
	"testing"
)

// TestOgnlException 测试 OGNL 异常相关的表达式
// 对应 Java 的 OgnlExceptionTest.java
// 测试可能导致异常或错误条件的 OGNL 表达式解析
func TestOgnlException(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "错误1 报错才是正确的",
			input: "45ac", // 这个表达式会导致解析错误
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: "45ac",
			},
		},
		{
			name:  "Division by zero expression",
			input: "1 / 0",
			expected: ExpectedNode{
				Type:     "ASTDivide",
				Fragment: "1 / 0",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			name:  "Null pointer access - method call on null",
			input: "null.toString()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "null.toString()",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "null"},
					{Type: "ASTMethod", Fragment: "toString()"},
				},
			},
		},
		{
			name:  "错误2 报错才是正确的",
			input: "5.length()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "5.length()",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
					{Type: "ASTMethod", Fragment: "length()"},
				},
			},
		},
		{
			name:  "Array index out of bounds - negative index",
			input: "array[-1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[-1]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "array",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"array\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[-1]",
						Children: []ExpectedNode{
							{Type: "ASTNegate", Fragment: "-1"},
						},
					},
				},
			},
		},
		{
			name:  "Invalid type conversion - string to number",
			input: "\"abc\" + 5",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "\"abc\" + 5",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"abc\""},
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			name:  "Invalid boolean conversion",
			input: "\"true\" && true",
			expected: ExpectedNode{
				Type:     "ASTAnd",
				Fragment: "\"true\" && true",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"true\""},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
		{
			name:  "Invalid property access on null",
			input: "null.property",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "null.property",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "null"},
					{
						Type:     "ASTProperty",
						Fragment: "property",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"property\""},
						},
					},
				},
			},
		},
		{
			name:  "Invalid method parameters - wrong type",
			input: "method(\"string\", 123)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "method(\"string\", 123)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"string\""},
					{Type: "ASTConst", Fragment: "123"},
				},
			},
		},
		{
			name:  "Invalid static field access",
			input: "@java.lang.String@INVALID_FIELD",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@java.lang.String@INVALID_FIELD",
			},
		},
		{
			name:  "Invalid static method call",
			input: "@java.lang.Math@invalidMethod()",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@java.lang.Math@invalidMethod()",
			},
		},
		{
			name:  "Complex invalid expression - nested null access",
			input: "null.field.method()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "null.field.method()",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "null"},
					{
						Type:     "ASTProperty",
						Fragment: "field",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"field\""},
						},
					},
					{Type: "ASTMethod", Fragment: "method()"},
				},
			},
		},
		{
			name:  "Invalid constructor call",
			input: "new NonExistentClass()",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new NonExistentClass()",
			},
		},
		{
			name:  "Invalid enum comparison",
			input: "@java.lang.Thread.State@NEW == \"INVALID\"",
			expected: ExpectedNode{
				Type:     "ASTEq",
				Fragment: "@java.lang.Thread.State@NEW == \"INVALID\"",
				Children: []ExpectedNode{
					{Type: "ASTStaticField", Fragment: "@java.lang.Thread.State@NEW"},
					{Type: "ASTConst", Fragment: "\"INVALID\""},
				},
			},
		},
		{
			name:  "Invalid regex pattern",
			input: "\"test\" =~ \"[invalid regex\"",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "\"test\" = ~\"[invalid regex\"",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"test\""},
					{Type: "ASTBitNegate",
						Fragment: "~\"[invalid regex\"",
						Children: []ExpectedNode{{Type: "ASTConst", Fragment: "\"[invalid regex\""}},
					},
				},
			},
		},
		{
			name:  "Invalid projection expression",
			input: "list.{invalid}",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list.{invalid}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{
						Type:     "ASTProject",
						Fragment: "{invalid}",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "invalid",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"invalid\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Invalid selection expression",
			input: "list.{? invalid}",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list.{? invalid}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{
						Type:     "ASTSelectFirst",
						Fragment: "{? invalid}",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "invalid",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"invalid\""},
								},
							},
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
