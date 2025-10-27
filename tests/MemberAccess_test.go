package test

import (
	"testing"
)

// TestMemberAccess 测试成员访问表达式
// 对应 Java 的 MemberAccessTest.java
func TestMemberAccess(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Static method call - @Runtime@getRuntime()",
			input: "@Runtime@getRuntime()",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@Runtime@getRuntime()",
			},
		},
		{
			name:  "Property access - bigIntValue",
			input: "bigIntValue",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "bigIntValue",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"bigIntValue\""},
				},
			},
		},
		{
			name:  "Method call - getBigIntValue()",
			input: "getBigIntValue()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getBigIntValue()",
			},
		},
		{
			name:  "Static property access - @System@getProperty('java.specification.version')",
			input: "@System@getProperty('java.specification.version')",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@System@getProperty(\"java.specification.version\")",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"java.specification.version\""},
				},
			},
		},
		{
			name:  "Property access - stringValue",
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
			name:  "Method call with parameter - setBigIntValue(25)",
			input: "setBigIntValue(25)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "setBigIntValue(25)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "25"},
				},
			},
		},
		{
			name:  "Chain property access - simple.bigIntValue",
			input: "simple.bigIntValue",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "simple.bigIntValue",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "simple",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"simple\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "bigIntValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bigIntValue\""},
						},
					},
				},
			},
		},
		{
			name:  "Chain method access - simple.getBigIntValue()",
			input: "simple.getBigIntValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "simple.getBigIntValue()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "simple",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"simple\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getBigIntValue()",
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
