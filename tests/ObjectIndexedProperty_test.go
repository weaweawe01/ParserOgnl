package test

import (
	"testing"
)

// TestObjectIndexedProperty 测试对象索引属性访问表达式
// 对应 Java 的 ObjectIndexedPropertyTest.java
// 测试通过索引（如 obj[key]）访问对象属性的各种场景
func TestObjectIndexedProperty(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			// testGetNonIndexedPropertyThroughAttributesMap: attributes["bar"]
			input: "attributes[\"bar\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "attributes[\"bar\"]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "attributes",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attributes\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"bar\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
						},
					},
				},
			},
		},
		{
			// testGetIndexedProperty: attribute["foo"]
			input: "attribute[\"foo\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "attribute[\"foo\"]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "attribute",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attribute\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"foo\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"foo\""},
						},
					},
				},
			},
		},
		{
			// testSetIndexedProperty: attribute["bar"]
			input: "attribute[\"bar\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "attribute[\"bar\"]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "attribute",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attribute\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"bar\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
						},
					},
				},
			},
		},
		{
			// testGetIndexedPropertyFromIndexedThenThroughOther:
			// attribute["other"].attribute["bar"]
			input: "attribute[\"other\"].attribute[\"bar\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "attribute[\"other\"].attribute[\"bar\"]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "attribute",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attribute\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"other\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"other\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "attribute",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attribute\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"bar\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
						},
					},
				},
			},
		},
		{
			// testGetPropertyBackThroughMapToConfirmFromIndexed:
			// attribute["other"].attributes["bar"]
			input: "attribute[\"other\"].attributes[\"bar\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "attribute[\"other\"].attributes[\"bar\"]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "attribute",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attribute\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"other\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"other\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "attributes",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attributes\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"bar\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bar\""},
						},
					},
				},
			},
		},
		{
			// testIllegalDynamicSubscriptAccessToObjectIndexedProperty: attribute[$]
			// 动态下标访问
			input: "attribute[$]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "attribute[$]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "attribute",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"attribute\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[$]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "$"},
						},
					},
				},
			},
		},
		{
			// testBeanIndexedValue: bean2.bean3.indexedValue[25]
			input: "bean2.bean3.indexedValue[25]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "bean2.bean3.indexedValue[25]",
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
						Fragment: "bean3",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean3\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "indexedValue",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"indexedValue\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[25]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "25"},
						},
					},
				},
			},
		},
		{
			// 额外测试：数字索引
			input: "array[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[0]",
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
						Fragment: "[0]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
						},
					},
				},
			},
		},
		{
			// 额外测试：表达式作为索引
			input: "items[i + 1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "items[i + 1]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "items",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"items\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[i + 1]",
						Children: []ExpectedNode{
							{
								Type:     "ASTAdd",
								Fragment: "i + 1",
								Children: []ExpectedNode{
									{
										Type:     "ASTProperty",
										Fragment: "i",
										Children: []ExpectedNode{
											{Type: "ASTConst", Fragment: "\"i\""},
										},
									},
									{Type: "ASTConst", Fragment: "1"},
								},
							},
						},
					},
				},
			},
		},
		{
			// 额外测试：变量作为索引
			input: "map[key]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[key]",
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
						Fragment: "[key]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "key",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"key\""},
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
