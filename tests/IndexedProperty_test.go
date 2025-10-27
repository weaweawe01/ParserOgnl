package test

import (
	"testing"
)

// TestIndexedProperty 测试索引属性表达式
// 对应 Java 的 IndexedPropertyTest.java
func TestIndexedProperty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Get values - getValues",
			input: "getValues",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "getValues",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"getValues\""},
				},
			},
		},
		{
			name:  "Values property - [\"values\"]",
			input: "[\"values\"]",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "[\"values\"]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"values\""},
				},
			},
		},
		{
			name:  "Values index 0 - [0]",
			input: "[0]",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "[0]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "0"},
				},
			},
		},
		{
			name:  "Get values index 0 - getValues()[0]",
			input: "getValues()[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getValues()[0]",
				Children: []ExpectedNode{
					{Type: "ASTMethod", Fragment: "getValues()"},
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
			name:  "Values index 0 direct - values[0]",
			input: "values[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "values[0]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "values",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"values\""},
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
			name:  "Values caret (first) - values[^]",
			input: "values[^]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "values[^]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "values",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"values\""},
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
			name:  "Values pipe (middle) - values[|]",
			input: "values[|]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "values[|]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "values",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"values\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[|]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "|"},
						},
					},
				},
			},
		},
		{
			name:  "Values dollar (last) - values[$]",
			input: "values[$]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "values[$]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "values",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"values\""},
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
			name:  "Set values index 1 - values[1]",
			input: "values[1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "values[1]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "values",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"values\""},
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
			name:  "Set values index 2 method - setValues(2, \"xxxx\")",
			input: "setValues(2, \"xxxx\")",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "setValues(2, \"xxxx\")",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "2"},
					{Type: "ASTConst", Fragment: "\"xxxx\""},
				},
			},
		},
		{
			name:  "Get title with list size - getTitle(list.size)",
			input: "getTitle(list.size)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getTitle(list.size)",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "list.size",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "list",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"list\""},
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
			},
		},
		{
			name:  "Source total - source.total",
			input: "source.total",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "source.total",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "source",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"source\""},
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
			name:  "Indexer line index - indexer.line[index]",
			input: "indexer.line[index]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "indexer.line[index]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "indexer",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"indexer\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "line",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"line\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[index]",
						Children: []ExpectedNode{
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
		{
			name:  "List long value - list[2].longValue()",
			input: "list[2].longValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "list[2].longValue()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "list",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"list\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[2]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2"},
						},
					},
					{Type: "ASTMethod", Fragment: "longValue()"},
				},
			},
		},
		{
			name:  "Map value id - map.value.id",
			input: "map.value.id",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.value.id",
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
						Fragment: "value",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"value\""},
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
		},
		{
			name:  "Property with string key - property['hoodak']",
			input: "property['hoodak']",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "property[\"hoodak\"]",
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
						Fragment: "[\"hoodak\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"hoodak\""},
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
