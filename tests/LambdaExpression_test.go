package test

import (
	"testing"
)

// TestLambdaExpression 测试 Lambda 表达式
// 对应 Java 的 LambdaExpressionTest.java
// Lambda 表达式使用 :[] 语法定义匿名函数
func TestLambdaExpression(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Array length with lambda - #a=:[33](20).longValue().{0}.toArray().length",
			input: "#a=:[33](20).longValue().{0}.toArray().length",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "#a = (:[33])(20).longValue().{0}.toArray().length",
				Children: []ExpectedNode{
					{Type: "ASTVarRef", Fragment: "#a"},
					{
						Type:     "ASTChain",
						Fragment: "(:[33])(20).longValue().{0}.toArray().length",
						Children: []ExpectedNode{
							{
								Type:     "ASTEval",
								Fragment: "(:[33])(20)",
								Children: []ExpectedNode{
									{
										Type:     "ASTConst",
										Fragment: ":[33]",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "33"},
										},
									},
									{Type: "ASTConst", Fragment: "20"},
								},
							},
							{Type: "ASTMethod", Fragment: "longValue()"},
							{
								Type:     "ASTProject",
								Fragment: "{0}",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "0"},
								},
							},
							{Type: "ASTMethod", Fragment: "toArray()"},
							{
								Type:     "ASTProperty",
								Fragment: "length",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"length\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Factorial lambda - #fact=:[#this <=1 ? 1 : #fact(#this-1) * #this], #fact(30)",
			input: "#fact=:[#this <=1 ? 1 : #fact(#this-1) * #this], #fact(30)",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#fact = :[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this], (#fact)(30)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#fact = :[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this]",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#fact"},
							{
								Type:     "ASTConst",
								Fragment: ":[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this]",
								Children: []ExpectedNode{
									{
										Type:     "ASTTest",
										Fragment: "(#this <= 1) ? 1 : (#fact)(#this - 1) * #this",
										Children: []ExpectedNode{
											{
												Type:     "ASTLessEq",
												Fragment: "#this <= 1",
												Children: []ExpectedNode{
													{Type: "ASTThisVarRef", Fragment: "#this"},
													{Type: "ASTConst", Fragment: "1"},
												},
											},
											{Type: "ASTConst", Fragment: "1"},
											{
												Type:     "ASTMultiply",
												Fragment: "(#fact)(#this - 1) * #this",
												Children: []ExpectedNode{
													{
														Type:     "ASTEval",
														Fragment: "(#fact)(#this - 1)",
														Children: []ExpectedNode{
															{Type: "ASTVarRef", Fragment: "#fact"},
															{
																Type:     "ASTSubtract",
																Fragment: "#this - 1",
																Children: []ExpectedNode{
																	{Type: "ASTThisVarRef", Fragment: "#this"},
																	{Type: "ASTConst", Fragment: "1"},
																},
															},
														},
													},
													{Type: "ASTThisVarRef", Fragment: "#this"},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Type:     "ASTEval",
						Fragment: "(#fact)(30)",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#fact"},
							{Type: "ASTConst", Fragment: "30"},
						},
					},
				},
			},
		},
		{
			name:  "Factorial with Long - #fact=:[#this <= 1 ? 1 : #fact(#this-1) * #this], #fact(30L)",
			input: "#fact=:[#this <= 1 ? 1 : #fact(#this-1) * #this], #fact(30L)",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#fact = :[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this], (#fact)(30L)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#fact = :[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this]",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#fact"},
							{
								Type:     "ASTConst",
								Fragment: ":[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this]",
								Children: []ExpectedNode{
									{
										Type:     "ASTTest",
										Fragment: "(#this <= 1) ? 1 : (#fact)(#this - 1) * #this",
									},
								},
							},
						},
					},
					{
						Type:     "ASTEval",
						Fragment: "(#fact)(30L)",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#fact"},
							{Type: "ASTConst", Fragment: "30L"},
						},
					},
				},
			},
		},
		{
			name:  "Factorial with BigInteger - #fact=:[#this <= 1 ? 1 : #fact(#this-1) * #this], #fact(30h)",
			input: "#fact=:[#this <= 1 ? 1 : #fact(#this-1) * #this], #fact(30h)",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#fact = :[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this], (#fact)(30H)",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#fact = :[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this]",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#fact"},
							{
								Type:     "ASTConst",
								Fragment: ":[(#this <= 1) ? 1 : (#fact)(#this - 1) * #this]",
							},
						},
					},
					{
						Type:     "ASTEval",
						Fragment: "(#fact)(30H)",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#fact"},
							{Type: "ASTConst", Fragment: "30H"},
						},
					},
				},
			},
		},
		{
			name:  "Bump lambda - #bump = :[ #this.{ #this + 1 } ], (#bump)({ 1, 2, 3 })",
			input: "#bump = :[ #this.{ #this + 1 } ], (#bump)({ 1, 2, 3 })",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#bump = :[#this.{(#this + 1)}], (#bump)({ 1, 2, 3 })",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#bump = :[#this.{(#this + 1)}]",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#bump"},
							{
								Type:     "ASTConst",
								Fragment: ":[#this.{(#this + 1)}]",
								Children: []ExpectedNode{
									{
										Type:     "ASTChain",
										Fragment: "#this.{(#this + 1)}",
										Children: []ExpectedNode{
											{Type: "ASTThisVarRef", Fragment: "#this"},
											{
												Type:     "ASTProject",
												Fragment: "{(#this + 1)}",
												Children: []ExpectedNode{
													{
														Type:     "ASTAdd",
														Fragment: "#this + 1",
														Children: []ExpectedNode{
															{Type: "ASTThisVarRef", Fragment: "#this"},
															{Type: "ASTConst", Fragment: "1"},
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
						Type:     "ASTEval",
						Fragment: "(#bump)({ 1, 2, 3 })",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#bump"},
							{
								Type:     "ASTList",
								Fragment: "{ 1, 2, 3 }",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "1"},
									{Type: "ASTConst", Fragment: "2"},
									{Type: "ASTConst", Fragment: "3"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Simple lambda - :[#this * 2]",
			input: ":[#this * 2]",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: ":[#this * 2]",
				Children: []ExpectedNode{
					{
						Type:     "ASTMultiply",
						Fragment: "#this * 2",
						Children: []ExpectedNode{
							{Type: "ASTThisVarRef", Fragment: "#this"},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
				},
			},
		},
		{
			name:  "Lambda with parameters - :[#this[0] + #this[1]]",
			input: ":[#this[0] + #this[1]]",
			expected: ExpectedNode{
				Type:     "ASTConst",
				Fragment: ":[#this[0] + #this[1]]",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "#this[0] + #this[1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "#this[0]",
								Children: []ExpectedNode{
									{Type: "ASTThisVarRef", Fragment: "#this"},
									{
										Type:     "ASTProperty",
										Fragment: "[0]",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "0"},
										},
									},
								},
							},
							{
								Type:     "ASTChain",
								Fragment: "#this[1]",
								Children: []ExpectedNode{
									{Type: "ASTThisVarRef", Fragment: "#this"},
									{
										Type:     "ASTProperty",
										Fragment: "[1]",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "1"},
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
		t.Run(tc.name, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
