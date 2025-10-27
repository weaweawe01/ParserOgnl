package test

import (
	"testing"
)

// TestMethodWithConversion 测试方法调用时的参数类型转换
// 对应 Java 的 MethodWithConversionTest.java
func TestMethodWithConversion(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			// testSetValues: setValues(10, "10.56", 34.225D)
			input: "setValues(10, \"10.56\", 34.225D)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "setValues(10, \"10.56\", 34.225)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "10"},
					{Type: "ASTConst", Fragment: "\"10.56\""},
					{Type: "ASTConst", Fragment: "34.225"},
				},
			},
		},
		{
			// testStringValue: stringValue
			input: "stringValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "stringValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"stringValue\""},
				},
			},
		},
		{
			// testStringValue: floatValue
			input: "floatValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "floatValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"floatValue\""},
				},
			},
		},
		{
			// testStringValue: intValue
			input: "intValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "intValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"intValue\""},
				},
			},
		},
		{
			// testSetStringValue: setStringValue('x')
			input: "setStringValue('x')",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "setStringValue('x')",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "'x'"},
				},
			},
		},
		{
			// testGetValueIsTrue: getValueIsTrue(rootValue)
			input: "getValueIsTrue(rootValue)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getValueIsTrue(rootValue)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "rootValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"rootValue\""},
						},
					},
				},
			},
		},
		{
			// testMessagesFormat: messages.format('Testing', one, two, three)
			input: "messages.format('Testing', one, two, three)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "messages.format(\"Testing\", one, two, three)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "messages",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"messages\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "format(\"Testing\", one, two, three)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"Testing\""},
							{
								Type:     "ASTProperty",
								Fragment: "one",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"one\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "two",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"two\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "three",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"three\""},
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
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
