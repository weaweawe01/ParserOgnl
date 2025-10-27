package test

import (
	"testing"
)

// TestMethod 测试方法调用表达式（基于 Java 的 MethodTest.java）
func TestMethod(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple method call - hashCode()",
			input: "hashCode()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "hashCode()",
			},
		},
		{
			name:  "Method call with conditional - getBooleanValue() ? \"here\" : \"\"",
			input: "getBooleanValue() ? \"here\" : \"\"",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "getBooleanValue() ? \"here\" : \"\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getBooleanValue()",
					},
					{
						Type:     "ASTConst",
						Fragment: "\"here\"",
					},
					{
						Type:     "ASTConst",
						Fragment: "\"\"",
					},
				},
			},
		},
		{
			name:  "Method call with negation - getValueIsTrue(!false) ? \"\" : \"here\"",
			input: "getValueIsTrue(!false) ? \"\" : \"here\" ",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "getValueIsTrue(!false) ? \"\" : \"here\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTMethod",
						Fragment: "getValueIsTrue(!false)",
						Children: []ExpectedNode{
							{
								Type:     "ASTNot",
								Fragment: "!false",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "false",
									},
								},
							},
						},
					},
					{
						Type:     "ASTConst",
						Fragment: "\"\"",
					},
					{
						Type:     "ASTConst",
						Fragment: "\"here\"",
					},
				},
			},
		},
		{
			name:  "Chained method call - messages.format('ShowAllCount', one)",
			input: "messages.format('ShowAllCount', one)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "messages.format(\"ShowAllCount\", one)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "messages",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"messages\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "format(\"ShowAllCount\", one)",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"ShowAllCount\"",
							},
							{
								Type:     "ASTProperty",
								Fragment: "one",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"one\"",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call with array literal - messages.format('ShowAllCount', {one})",
			input: "messages.format('ShowAllCount', {one})",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "messages.format(\"ShowAllCount\", { one })",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "messages",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"messages\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "format(\"ShowAllCount\", { one })",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"ShowAllCount\"",
							},
							{
								Type:     "ASTList",
								Fragment: "{ one }",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "one",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"one\"",
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
			name:  "Method call with multiple array elements - messages.format('ShowAllCount', {one, two})",
			input: "messages.format('ShowAllCount', {one, two})",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "messages.format(\"ShowAllCount\", { one, two })",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "messages",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"messages\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "format(\"ShowAllCount\", { one, two })",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"ShowAllCount\"",
							},
							{
								Type:     "ASTList",
								Fragment: "{ one, two }",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "one",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"one\"",
											},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "two",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"two\"",
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
			name:  "Method call with multiple arguments - messages.format('ShowAllCount', one, two)",
			input: "messages.format('ShowAllCount', one, two)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "messages.format(\"ShowAllCount\", one, two)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "messages",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"messages\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "format(\"ShowAllCount\", one, two)",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"ShowAllCount\"",
							},
							{
								Type:     "ASTProperty",
								Fragment: "one",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"one\"",
									},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "two",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"two\"",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call with static field argument - getTestValue(@ognl.test.objects.SimpleEnum@ONE.value)",
			input: "getTestValue(@ognl.test.objects.SimpleEnum@ONE.value)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getTestValue(@ognl.test.objects.SimpleEnum@ONE.value)",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "@ognl.test.objects.SimpleEnum@ONE.value",
						Children: []ExpectedNode{
							{
								Type:     "ASTStaticField",
								Fragment: "@ognl.test.objects.SimpleEnum@ONE",
							},
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
						},
					},
				},
			},
		},
		{
			name:  "Static method call chain - @ognl.test.MethodTest@getA().isProperty()",
			input: "@ognl.test.MethodTest@getA().isProperty()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@ognl.test.MethodTest@getA().isProperty()",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticMethod",
						Fragment: "@ognl.test.MethodTest@getA()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "isProperty()",
					},
				},
			},
		},
		{
			name:  "Simple method call - isDisabled()",
			input: "isDisabled()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isDisabled()",
			},
		},
		{
			name:  "Property access instead of method call - isTruck",
			input: "isTruck",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "isTruck",
				Children: []ExpectedNode{
					{
						Type:     "ASTConst",
						Fragment: "\"isTruck\"",
					},
				},
			},
		},
		{
			name:  "Method call - isEditorDisabled()",
			input: "isEditorDisabled()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isEditorDisabled()",
			},
		},
		{
			name:  "Method call on different root - addValue(name)",
			input: "addValue(name)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "addValue(name)",
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
				},
			},
		},
		{
			name:  "Method call with property chain - getDisplayValue(methodsTest.allowDisplay)",
			input: "getDisplayValue(methodsTest.allowDisplay)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getDisplayValue(methodsTest.allowDisplay)",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "methodsTest.allowDisplay",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "methodsTest",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"methodsTest\"",
									},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "allowDisplay",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "\"allowDisplay\"",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call with arguments - isThisVarArgsWorking(three, rootValue)",
			input: "isThisVarArgsWorking(three, rootValue)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isThisVarArgsWorking(three, rootValue)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "three",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"three\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "rootValue",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"rootValue\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call without arguments - isThisVarArgsWorking()",
			input: "isThisVarArgsWorking()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isThisVarArgsWorking()",
			},
		},
		{
			name:  "Chained method call - service.getFullMessageFor(value, null)",
			input: "service.getFullMessageFor(value, null)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "service.getFullMessageFor(value, null)",
				Children: []ExpectedNode{
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
						Fragment: "getFullMessageFor(value, null)",
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
								Fragment: "null",
							},
						},
					},
				},
			},
		},
		{
			name:  "Chained method call - testMethods.getBean('TestBean')",
			input: "testMethods.getBean('TestBean')",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.getBean(\"TestBean\")",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getBean(\"TestBean\")",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"TestBean\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Property access instead of method call - testMethods.testProperty",
			input: "testMethods.testProperty",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.testProperty",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "testProperty",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testProperty\"",
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call with array argument - testMethods.argsTest1({one})",
			input: "testMethods.argsTest1({one})",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.argsTest1({ one })",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "argsTest1({ one })",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ one }",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "one",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"one\"",
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
			name:  "Method call with list argument - testMethods.argsTest2({one})",
			input: "testMethods.argsTest2({one})",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.argsTest2({ one })",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "argsTest2({ one })",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ one }",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "one",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"one\"",
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
			name:  "Method call with array argument - testMethods.argsTest3({one})",
			input: "testMethods.argsTest3({one})",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.argsTest3({ one })",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "argsTest3({ one })",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ one }",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "one",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"one\"",
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
			name:  "Nested method calls - testMethods.showList(testMethods.getObjectList())",
			input: "testMethods.showList(testMethods.getObjectList())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.showList(testMethods.getObjectList())",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "showList(testMethods.getObjectList())",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "testMethods.getObjectList()",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "testMethods",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"testMethods\"",
											},
										},
									},
									{
										Type:     "ASTMethod",
										Fragment: "getObjectList()",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Nested method calls - testMethods.showList(testMethods.getStringList())",
			input: "testMethods.showList(testMethods.getStringList())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.showList(testMethods.getStringList())",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "showList(testMethods.getStringList())",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "testMethods.getStringList()",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "testMethods",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"testMethods\"",
											},
										},
									},
									{
										Type:     "ASTMethod",
										Fragment: "getStringList()",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Nested method calls - testMethods.showList(testMethods.getStringArray())",
			input: "testMethods.showList(testMethods.getStringArray())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.showList(testMethods.getStringArray())",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "showList(testMethods.getStringArray())",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "testMethods.getStringArray()",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "testMethods",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"testMethods\"",
											},
										},
									},
									{
										Type:     "ASTMethod",
										Fragment: "getStringArray()",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Method call with array conversion - testMethods.showStringList(testMethods.getStringList().toArray(new String[0]))",
			input: "testMethods.showStringList(testMethods.getStringList().toArray(new String[0]))",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.showStringList(testMethods.getStringList().toArray(new String[0]))",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "showStringList(testMethods.getStringList().toArray(new String[0]))",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "testMethods.getStringList().toArray(new String[0])",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "testMethods",
										Children: []ExpectedNode{
											{
												Type:     "ASTConst",
												Fragment: "\"testMethods\"",
											},
										},
									},
									{
										Type:     "ASTMethod",
										Fragment: "getStringList()",
									},
									{
										Type:     "ASTMethod",
										Fragment: "toArray(new String[0])",
										Children: []ExpectedNode{
											{
												Type:     "ASTCtor",
												Fragment: "new String[0]",
												Children: []ExpectedNode{
													{
														Type:     "ASTConst",
														Fragment: "0",
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
			},
		},
		{
			name:  "Method call with array literal - testMethods.avg({ 5, 5 })",
			input: "testMethods.avg({ 5, 5 })",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "testMethods.avg({ 5, 5 })",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "testMethods",
						Children: []ExpectedNode{
							{
								Type:     "ASTConst",
								Fragment: "\"testMethods\"",
							},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "avg({ 5, 5 })",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ 5, 5 }",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: "5",
									},
									{
										Type:     "ASTConst",
										Fragment: "5",
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
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
