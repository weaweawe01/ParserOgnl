package test

import (
	"testing"
)

// TestNestedMethod 测试嵌套方法调用表达式（基于 Java 的 NestedMethodTest.java）
func TestNestedMethod(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Nested property access - toDisplay.pictureUrl",
			input: "toDisplay.pictureUrl",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "toDisplay.pictureUrl",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "toDisplay",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"toDisplay\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "pictureUrl",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"pictureUrl\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Nested method call with property argument - page.createRelativeAsset(toDisplay.pictureUrl)",
			input: "page.createRelativeAsset(toDisplay.pictureUrl)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "page.createRelativeAsset(toDisplay.pictureUrl)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "page",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"page\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "createRelativeAsset(toDisplay.pictureUrl)",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "toDisplay.pictureUrl",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "toDisplay",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"toDisplay\"",
											},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "pictureUrl",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"pictureUrl\"",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Deep nested property access - component.page.title",
			input: "component.page.title",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "component.page.title",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "component",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"component\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "page",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"page\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "title",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"title\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Nested method call with multiple arguments - user.profile.update(name, email)",
			input: "user.profile.update(name, email)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "user.profile.update(name, email)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "user",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"user\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "profile",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"profile\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "update(name, email)",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "name",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"name\"",
									},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "email",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"email\"",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Complex nested method call with property chain - data.service.process(item.getValue())",
			input: "data.service.process(item.getValue())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "data.service.process(item.getValue())",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "data",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"data\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "service",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"service\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "process(item.getValue())",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "item.getValue()",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "item",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"item\"",
											},
										},
									},
									{
										Type:     "ASTMethod",
										Fragment: "getValue()",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call result property access - getUser().name",
			input: "getUser().name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getUser().name",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getUser()",
					},
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"name\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Chained method calls - getUser().getProfile().getName()",
			input: "getUser().getProfile().getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getUser().getProfile().getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getUser()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getProfile()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		{
			name:  "Property access then method call - user.name.toUpperCase()",
			input: "user.name.toUpperCase()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "user.name.toUpperCase()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "user",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"user\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "name",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"name\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "toUpperCase()",
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
