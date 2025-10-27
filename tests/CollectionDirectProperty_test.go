package test

import (
	"testing"
)

// TestCollectionDirectProperty 测试集合直接属性访问（基于 Java 的 CollectionDirectPropertyTest.java）
func TestCollectionDirectProperty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Size property",
			input: "size",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "size",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"size\""},
				},
			},
		},
		{
			name:  "IsEmpty property",
			input: "isEmpty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "isEmpty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"isEmpty\""},
				},
			},
		},
		{
			name:  "Iterator next",
			input: "iterator.next",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "iterator.next",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "iterator",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"iterator\""},
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
			name:  "Iterator hasNext",
			input: "iterator.hasNext",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "iterator.hasNext",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "iterator",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"iterator\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "hasNext",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"hasNext\""},
						},
					},
				},
			},
		},
		{
			name:  "Iterator hasNext after two nexts",
			input: "#it = iterator, #it.next, #it.next, #it.hasNext",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#it = iterator, #it.next, #it.next, #it.hasNext",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#it = iterator",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
							{
								Type:     "ASTProperty",
								Fragment: "iterator",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"iterator\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#it.next",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
							{
								Type:     "ASTProperty",
								Fragment: "next",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"next\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#it.next",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
							{
								Type:     "ASTProperty",
								Fragment: "next",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"next\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#it.hasNext",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
							{
								Type:     "ASTProperty",
								Fragment: "hasNext",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"hasNext\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Iterator next after two nexts",
			input: "#it = iterator, #it.next, #it.next",
			expected: ExpectedNode{
				Type:     "ASTSequence",
				Fragment: "#it = iterator, #it.next, #it.next",
				Children: []ExpectedNode{
					{
						Type:     "ASTAssign",
						Fragment: "#it = iterator",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
							{
								Type:     "ASTProperty",
								Fragment: "iterator",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"iterator\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#it.next",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
							{
								Type:     "ASTProperty",
								Fragment: "next",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"next\""},
								},
							},
						},
					},
					{
						Type:     "ASTChain",
						Fragment: "#it.next",
						Children: []ExpectedNode{
							{Type: "ASTVarRef", Fragment: "#it"},
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
			},
		},
		{
			name:  "Root map test",
			input: "map[\"test\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[\"test\"]",
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
						Fragment: "[\"test\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"test\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map size",
			input: "map.size",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.size",
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
						Fragment: "size",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"size\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map keySet",
			input: "map.keySet",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.keySet",
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
						Fragment: "keySet",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"keySet\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map values",
			input: "map.values",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.values",
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
						Fragment: "values",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"values\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map keys size",
			input: "map.keys.size",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.keys.size",
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
						Fragment: "keys",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"keys\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "size",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"size\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map size value",
			input: "map[\"size\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[\"size\"]",
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
						Fragment: "[\"size\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"size\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map isEmpty",
			input: "map.isEmpty",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.isEmpty",
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
						Fragment: "isEmpty",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"isEmpty\""},
						},
					},
				},
			},
		},
		{
			name:  "Root map isEmpty key",
			input: "map[\"isEmpty\"]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[\"isEmpty\"]",
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
						Fragment: "[\"isEmpty\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"isEmpty\""},
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
