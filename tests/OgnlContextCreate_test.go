package test

import (
	"testing"
)

// TestOgnlContextCreate 测试 OGNL 上下文创建相关的表达式
// 对应 Java 的 OgnlContextCreateTest.java
// 主要测试变量引用（#variable）和静态方法调用
func TestOgnlContextCreate(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			// createContext, createContextWithRoot, createContextWithNullRoot: #test
			// 变量引用表达式
			input: "#test",
			expected: ExpectedNode{
				Type:     "ASTVarRef",
				Fragment: "#test",
			},
		},
		{
			// createContextWithClassResolver 等: @ognl.test.MyClass@getValue()
			// 静态方法调用表达式
			input: "@ognl.test.MyClass@getValue()",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@ognl.test.MyClass@getValue()",
			},
		},
		{
			// 更复杂的静态方法调用，带参数
			input: "@java.lang.Math@max(1, 2)",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@java.lang.Math@max(1, 2)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			// 静态字段访问
			input: "@java.lang.Math@PI",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@java.lang.Math@PI",
			},
		},
		{
			// 变量与属性链式访问
			input: "#root.stringValue",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#root.stringValue",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
					{
						Type:     "ASTProperty",
						Fragment: "stringValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"stringValue\""},
						},
					},
				},
			},
		},
		{
			// 变量与方法调用
			input: "#root.getValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#root.getValue()",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
					{Type: "ASTMethod", Fragment: "getValue()"},
				},
			},
		},
		{
			// 多个变量引用
			input: "#test1 + #test2",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "#test1 + #test2",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#test1"},
					{Type: "ASTVarRef", Fragment: "#test2"},
				},
			},
		},
		{
			// 变量作为方法参数
			input: "setValue(#test)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "setValue(#test)",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#test"},
				},
			},
		},
		{
			// 静态方法调用后的链式访问
			input: "@ognl.test.MyClass@getValue().length()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@ognl.test.MyClass@getValue().length()",
				Children: []ExpectedNode{
					{Type: "ASTStaticMethod", Fragment: "@ognl.test.MyClass@getValue()"},
					{Type: "ASTMethod", Fragment: "length()"},
				},
			},
		},
		{
			// 变量在条件表达式中使用
			input: "#test ? 'yes' : 'no'",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "#test ? \"yes\" : \"no\"",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#test"},
					{Type: "ASTConst", Fragment: "\"yes\""},
					{Type: "ASTConst", Fragment: "\"no\""},
				},
			},
		},
		{
			// 变量赋值
			input: "#result = getValue()",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#result = getValue()",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#result"},
					{Type: "ASTMethod", Fragment: "getValue()"},
				},
			},
		},
		{
			// 特殊变量 #this
			input: "#this",
			expected: ExpectedNode{
				Type:     "ASTThisVarRef",
				Fragment: "#this",
			},
		},
		{
			// 特殊变量 #root
			input: "#root",
			expected: ExpectedNode{
				Type:     "ASTRootVarRef",
				Fragment: "#root",
			},
		},
		{
			// 静态方法调用作为索引
			input: "array[@java.lang.Math@max(0, index)]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[@java.lang.Math@max(0, index)]",
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
						Fragment: "[@java.lang.Math@max(0, index)]",
						Children: []ExpectedNode{
							{
								Type:     "ASTStaticMethod",
								Fragment: "@java.lang.Math@max(0, index)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "0"},
									{
										Type:     "ASTProperty",
										Fragment: "index",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"index\""},
										},
									},
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
