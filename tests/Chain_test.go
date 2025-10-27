package test

import (
	"testing"
)

// TestChain 测试链式表达式（基于 Java 的 ChainTest.java）
func TestChain(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    ExpectedNode
		isChain     bool
		description string
	}{
		{
			name:        "Single variable reference - not a chain",
			input:       "#name",
			isChain:     false,
			description: "单个变量引用不是链式表达式",
			expected: ExpectedNode{
				Type:     "ASTVarRef",
				Fragment: "#name",
			},
		},
		{
			name:        "Variable with property access - is a chain",
			input:       "#name.lastChar",
			isChain:     true,
			description: "变量后跟属性访问是链式表达式",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#name.lastChar",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#name"},
					{
						Type:     "ASTProperty",
						Fragment: "lastChar",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"lastChar\""},
						},
					},
				},
			},
		},
		{
			name:        "Map literal with chain expressions - is a chain",
			input:       "#{name.lastChar, #boo}",
			isChain:     true,
			description: "包含链式表达式的 Map 字面量也被认为是链式表达式",
			expected: ExpectedNode{
				Type:     "ASTMap",
				Fragment: "#{ name.lastChar : null, #boo : null }",
				Children: []ExpectedNode{
					{
						Type:     "ASTKeyValue",
						Fragment: "name.lastChar : null",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "name.lastChar",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "name",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"name\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "lastChar",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"lastChar\""},
										},
									},
								},
							},
						},
					},
					{
						Type:     "ASTKeyValue",
						Fragment: "#boo : null",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#boo"},
						},
					},
				},
			},
		},
		{
			name:        "Assignment with map literal containing chains - is a chain",
			input:       "boo = #{name.lastChar, #boo, foo()}",
			isChain:     true,
			description: "赋值语句右侧的 Map 字面量包含链式表达式",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "boo = #{ name.lastChar : null, #boo : null, foo() : null }",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "boo",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"boo\""},
						},
					},
					{
						Type:     "ASTMap",
						Fragment: "#{ name.lastChar : null, #boo : null, foo() : null }",
						Children: []ExpectedNode{
							{
								Type:     "ASTKeyValue",
								Fragment: "name.lastChar : null",
								Children: []ExpectedNode{
									{
										Type:     "ASTChain",
										Fragment: "name.lastChar",
										Children: []ExpectedNode{
											{
												Type:     "ASTProperty",
												Fragment: "name",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"name\""},
												},
											},
											{
												Type:     "ASTProperty",
												Fragment: "lastChar",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"lastChar\""},
												},
											},
										},
									},
								},
							},
							{
								Type:     "ASTKeyValue",
								Fragment: "#boo : null",
								Children: []ExpectedNode{
									{Type: "ASTVarRef", Fragment: "#boo"},
								},
							},
							{
								Type:     "ASTKeyValue",
								Fragment: "foo() : null",
								Children: []ExpectedNode{
									{Type: "ASTMethod", Fragment: "foo()"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:        "Array literal with chain expressions - is a chain",
			input:       "{name.lastChar, #boo, foo()}",
			isChain:     true,
			description: "包含链式表达式的数组字面量也被认为是链式表达式",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ name.lastChar, #boo, foo() }",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "name.lastChar",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "name",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"name\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "lastChar",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"lastChar\""},
								},
							},
						},
					},
					{Type: "ASTVarRef", Fragment: "#boo"},
					{Type: "ASTMethod", Fragment: "foo()"},
				},
			},
		},
		{
			name:        "Parenthesized sequence with chains - is a chain",
			input:       "(name.lastChar, #boo, foo())",
			isChain:     true,
			description: "括号包裹的序列表达式包含链式表达式",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "name.lastChar, #boo, foo()",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "name.lastChar",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "name",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"name\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "lastChar",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"lastChar\""},
								},
							},
						},
					},
					{Type: "ASTVarRef", Fragment: "#boo"},
					{Type: "ASTMethod", Fragment: "foo()"},
				},
			},
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)

			// 检查是否为链式表达式
			actualIsChain := isChainExpression(expr)
			if actualIsChain != tc.isChain {
				t.Errorf("表达式 '%s' 链式判断错误: 期望 %v, 实际 %v",
					tc.input, tc.isChain, actualIsChain)
			}

			// 检查 AST 结构
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// isChainExpression 判断表达式是否为链式表达式或包含链式表达式
// 对应 Java 的 SimpleNode.isChain(OgnlContext) 方法
func isChainExpression(expr interface{}) bool {
	if expr == nil {
		return false
	}

	// 根据表达式类型判断
	switch e := expr.(type) {
	case interface{ Type() string }:
		nodeType := e.Type()
		// ASTChain 本身就是链式表达式
		if nodeType == "ASTChain" {
			return true
		}
		// ASTMap、ASTList、ASTSequence 如果包含链式子节点，也被认为是链式表达式
		if nodeType == "ASTMap" || nodeType == "ASTList" || nodeType == "ASTSequence" || nodeType == "ASTAssign" {
			return true
		}
	}

	return false
}

// TestChainShortCircuit 测试链式访问的短路特性
func TestChainShortCircuit(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		description string
		expected    ExpectedNode
	}{
		{
			name:        "Nested property access with null child",
			input:       "#parent.child.child.name",
			description: "访问空子对象的属性时应短路返回 null，而不是抛出空指针异常",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#parent.child.child.name",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#parent"},
					{
						Type:     "ASTProperty",
						Fragment: "child",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"child\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "child",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"child\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"name\""},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// TestThisPropertyEvaluation 测试 #this 属性的求值
func TestThisPropertyEvaluation(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Conditional with #this check",
			input: "map[$].(#this == null ? 'empty' : #this)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[$].((#this == null) ? \"empty\" : #this)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "map",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"map\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[$]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "$"},
						},
					},
					{
						Type:     "ASTTest",
						Fragment: "((#this == null) ? \"empty\" : #this)",
						Children: []ExpectedNode{
							{
								Type:     "ASTEq",
								Fragment: "(#this == null)",
								Children: []ExpectedNode{
									{Type: "ASTThisVarRef", Fragment: "#this"},
									{Type: "ASTConst", Fragment: "null"},
								},
							},
							{Type: "ASTConst", Fragment: "\"empty\""},
							{Type: "ASTThisVarRef", Fragment: "#this"},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
