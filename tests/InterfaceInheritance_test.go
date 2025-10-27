package test

import (
	"testing"
)

// TestInterfaceInheritance 测试接口继承表达式
// 对应 Java 的 InterfaceInheritanceTest.java
func TestInterfaceInheritance(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "MyMap - myMap",
			input: "myMap",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "myMap",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"myMap\""},
				},
			},
		},
		{
			name:  "MyMap test - myMap.test",
			input: "myMap.test",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap.test",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "test",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"test\""},
						},
					},
				},
			},
		},
		{
			name:  "List - list",
			input: "list",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "list",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"list\""},
				},
			},
		},
		{
			name:  "MyMap array 0 - myMap.array[0]",
			input: "myMap.array[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap.array[0]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
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
			name:  "MyMap list 1 - myMap.list[1]",
			input: "myMap.list[1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap.list[1]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
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
		{
			name:  "MyMap caret (first) - myMap[^]",
			input: "myMap[^]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap[^]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[^]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "^"},
						},
					},
				},
			},
		},
		{
			name:  "MyMap dollar (last) - myMap[$]",
			input: "myMap[$]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap[$]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
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
			name:  "Array dollar - array[$]",
			input: "array[$]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array[$]",
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
						Fragment: "[$]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "$"},
						},
					},
				},
			},
		},
		{
			name:  "MyMap string index - [\"myMap\"]",
			input: "[\"myMap\"]",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "[\"myMap\"]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"myMap\""},
				},
			},
		},
		{
			name:  "MyMap null - myMap[null]",
			input: "myMap[null]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap[null]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[null]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "null"},
						},
					},
				},
			},
		},
		{
			name:  "MyMap x null - myMap[#x = null]",
			input: "myMap[#x = null]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap[#x = null]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[#x = null]",
						Children: []ExpectedNode{
							{
								Type:     "ASTAssign",
								Fragment: "#x = null",
								Children: []ExpectedNode{
									{Type: "ASTVarRef", Fragment: "#x"},
									{Type: "ASTConst", Fragment: "null"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "MyMap null test - myMap.(null,test)",
			input: "myMap.(null,test)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myMap.null, test",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myMap\""},
						},
					},
					{
						Type:     "ASTSequence",
						Fragment: "null, test",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "null"},
							{
								Type:     "ASTProperty",
								Fragment: "test",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"test\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "MyMap null assign 25 - myMap[null] = 25",
			input: "myMap[null] = 25",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "myMap[null] = 25",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "myMap[null]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "myMap",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"myMap\""},
								},
							},
							{
								Type:     "ASTProperty",
								Fragment: "[null]",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "null"},
								},
							},
						},
					},
					{Type: "ASTConst", Fragment: "25"},
				},
			},
		},
		{
			name:  "Beans test bean - beans.testBean",
			input: "beans.testBean",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "beans.testBean",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "beans",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"beans\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "testBean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"testBean\""},
						},
					},
				},
			},
		},
		{
			name:  "Beans even odd next - beans.evenOdd.next",
			input: "beans.evenOdd.next",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "beans.evenOdd.next",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "beans",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"beans\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "evenOdd",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"evenOdd\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "next",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"next\""},
						},
					},
				},
			},
		},
		{
			name:  "Map comp form client id - map.comp.form.clientId",
			input: "map.comp.form.clientId",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.comp.form.clientId",
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
						Fragment: "comp",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"comp\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "form",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"form\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "clientId",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"clientId\""},
						},
					},
				},
			},
		},
		{
			name:  "Map comp get count - map.comp.getCount(genericIndex)",
			input: "map.comp.getCount(genericIndex)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.comp.getCount(genericIndex)",
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
						Fragment: "comp",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"comp\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getCount(genericIndex)",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "genericIndex",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"genericIndex\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Map custom list total - map.customList.total",
			input: "map.customList.total",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.customList.total",
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
						Fragment: "customList",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"customList\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "total",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"total\""},
						},
					},
				},
			},
		},
		{
			name:  "My test the map key - myTest.theMap['key']",
			input: "myTest.theMap['key']",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "myTest.theMap[\"key\"]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "myTest",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"myTest\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "theMap",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"theMap\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[\"key\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"key\""},
						},
					},
				},
			},
		},
		{
			name:  "Content provider has children - contentProvider.hasChildren(property)",
			input: "contentProvider.hasChildren(property)",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "contentProvider.hasChildren(property)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "contentProvider",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"contentProvider\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "hasChildren(property)",
						Children: []ExpectedNode{
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
			},
		},
		{
			name:  "Object index instanceof - objectIndex instanceof java.lang.Object",
			input: "objectIndex instanceof java.lang.Object",
			expected: ExpectedNode{
				Type:     "ASTInstanceof",
				Fragment: "objectIndex instanceof java.lang.Object",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "objectIndex",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"objectIndex\""},
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
