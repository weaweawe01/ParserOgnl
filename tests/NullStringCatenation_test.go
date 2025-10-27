package test

import (
	"testing"
)

// TestNullStringCatenation 测试字符串拼接表达式（特别是与 null 值的拼接）
// 对应 Java 的 NullStringCatenationTest.java
func TestNullStringCatenation(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			// testCatenateNullToString: "bar" + null
			input: "\"bar\" + null",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "\"bar\" + null",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"bar\""},
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		{
			// testCatenateNullObjectToString: "bar" + nullObject
			input: "\"bar\" + nullObject",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "\"bar\" + nullObject",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"bar\""},
					{
						Type:     "ASTProperty",
						Fragment: "nullObject",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"nullObject\""},
						},
					},
				},
			},
		},
		{
			// testCatenateNullObjectToNumber: 20.56 + nullObject
			input: "20.56 + nullObject",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "20.56 + nullObject",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "20.56"},
					{
						Type:     "ASTProperty",
						Fragment: "nullObject",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"nullObject\""},
						},
					},
				},
			},
		},
		{
			// testConditionalCatenation: (true ? 'tabHeader' : '') + (false ? 'tabHeader' : '')
			input: "(true ? 'tabHeader' : '') + (false ? 'tabHeader' : '')",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "(true ? 'tabHeader' : '') + (false ? 'tabHeader' : '')",
				Children: []ExpectedNode{
					{
						Type:     "ASTTest",
						Fragment: "true ? 'tabHeader' : ''",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "true"},
							{Type: "ASTConst", Fragment: "'tabHeader'"},
							{Type: "ASTConst", Fragment: "''"},
						},
					},
					{
						Type:     "ASTTest",
						Fragment: "false ? 'tabHeader' : ''",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "false"},
							{Type: "ASTConst", Fragment: "'tabHeader'"},
							{Type: "ASTConst", Fragment: "''"},
						},
					},
				},
			},
		},
		{
			// testConditionalCatenationWithInt: theInt == 0 ? '5%' : theInt + '%'
			input: "theInt == 0 ? '5%' : theInt + '%'",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "theInt == 0 ? '5%' : theInt + '%'",
				Children: []ExpectedNode{
					{
						Type:     "ASTEq",
						Fragment: "theInt == 0",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "theInt",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"theInt\""},
								},
							},
							{Type: "ASTConst", Fragment: "0"},
						},
					},
					{Type: "ASTConst", Fragment: "'5%'"},
					{
						Type:     "ASTAdd",
						Fragment: "theInt + '%'",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "theInt",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"theInt\""},
								},
							},
							{Type: "ASTConst", Fragment: "'%'"},
						},
					},
				},
			},
		},
		{
			// testCatenateWidth: 'width:' + width + ';'
			input: "'width:' + width + ';'",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "'width:' + width + ';'",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "'width:' + width",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "'width:'"},
							{
								Type:     "ASTProperty",
								Fragment: "width",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"width\""},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "';'"},
				},
			},
		},
		{
			// testCatenateLongAndIndex: theLong + '_' + index
			input: "theLong + '_' + index",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "theLong + '_' + index",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "theLong + '_'",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "theLong",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"theLong\""},
								},
							},
							{Type: "ASTConst", Fragment: "'_'"},
						},
					},
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
		{
			// testCatenateWithStaticField: 'javascript:' + @ognl.test.NullStringCatenationTest@MESSAGE
			input: "'javascript:' + @ognl.test.NullStringCatenationTest@MESSAGE",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "'javascript:' + @ognl.test.NullStringCatenationTest@MESSAGE",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "'javascript:'"},
					{
						Type:     "ASTStaticField",
						Fragment: "@ognl.test.NullStringCatenationTest@MESSAGE",
					},
				},
			},
		},
		{
			// testConditionalCatenationWithMethodCall (简化版，只解析表达式结构):
			// printDelivery ? '' : 'javascript:deliverySelected(' + property.carrier + ',' + currentDeliveryId + ')'
			input: "printDelivery ? '' : 'javascript:deliverySelected(' + property.carrier + ',' + currentDeliveryId + ')'",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "printDelivery ? '' : 'javascript:deliverySelected(' + property.carrier + ',' + currentDeliveryId + ')'",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "printDelivery",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"printDelivery\""},
						},
					},
					{Type: "ASTConst", Fragment: "''"},
					{
						Type:     "ASTAdd",
						Fragment: "'javascript:deliverySelected(' + property.carrier + ',' + currentDeliveryId + ')'",
						Children: []ExpectedNode{
							{
								Type:     "ASTAdd",
								Fragment: "'javascript:deliverySelected(' + property.carrier + ','",
								Children: []ExpectedNode{
									{
										Type:     "ASTAdd",
										Fragment: "'javascript:deliverySelected(' + property.carrier",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "'javascript:deliverySelected('"},
											{
												Type:     "ASTChain",
												Fragment: "property.carrier",
												Children: []ExpectedNode{
													{
														Type:     "ASTProperty",
														Fragment: "property",
														Children: []ExpectedNode{
															{Type: "ASTConst", Fragment: "\"property\""},
														},
													},
													{
														Type:     "ASTProperty",
														Fragment: "carrier",
														Children: []ExpectedNode{
															{Type: "ASTConst", Fragment: "\"carrier\""},
														},
													},
												},
											},
										},
									},
									{Type: "ASTConst", Fragment: "','"},
								},
							},
							{
								Type:     "ASTAdd",
								Fragment: "currentDeliveryId + ')'",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "currentDeliveryId",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"currentDeliveryId\""},
										},
									},
									{Type: "ASTConst", Fragment: "')'"},
								},
							},
						},
					},
				},
			},
		},
		{
			// testCatenateBeanIdAndInt: bean2.id + '_' + theInt
			input: "bean2.id + '_' + theInt",
			expected: ExpectedNode{
				Type:     "ASTAdd",
				Fragment: "bean2.id + '_' + theInt",
				Children: []ExpectedNode{
					{
						Type:     "ASTAdd",
						Fragment: "bean2.id + '_'",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "bean2.id",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "bean2",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"bean2\""},
										},
									},
									{
										Type:     "ASTProperty",
										Fragment: "id",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"id\""},
										},
									},
								},
							},
							{Type: "ASTConst", Fragment: "'_'"},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "theInt",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"theInt\""},
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
