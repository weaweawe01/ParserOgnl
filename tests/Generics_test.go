package test

import (
	"testing"
)

// TestGenerics 测试泛型相关的表达式解析
// 对应 Java 的 GenericsTest.java
// 注意：Go 版本主要测试语法解析，不涉及运行时泛型类型推断
func TestGenerics(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple property access - ids",
			input: "ids",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "ids",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"ids\""},
				},
			},
		},
		{
			name:  "Generic object property - value.data",
			input: "value.data",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "value.data",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"value\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "data",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"data\""},
						},
					},
				},
			},
		},
		{
			name:  "Array index access - ids[0]",
			input: "ids[0]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "ids[0]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "ids",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"ids\""},
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
			name:  "Array length property - ids.length",
			input: "ids.length",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "ids.length",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "ids",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"ids\""},
						},
					},
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
		{
			name:  "Generic list access - items[0].value",
			input: "items[0].value",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "items[0].value",
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
						Fragment: "[0]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"value\""},
						},
					},
				},
			},
		},
		{
			name:  "Generic method call - getItems().size()",
			input: "getItems().size()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getItems().size()",
				Children: []ExpectedNode{
					{Type: "ASTMethod", Fragment: "getItems()"},
					{Type: "ASTMethod", Fragment: "size()"},
				},
			},
		},
		{
			name:  "Generic getter - getValue().getData()",
			input: "getValue().getData()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "getValue().getData()",
				Children: []ExpectedNode{
					{Type: "ASTMethod", Fragment: "getValue()"},
					{Type: "ASTMethod", Fragment: "getData()"},
				},
			},
		},
		{
			name:  "Nested generic access - container.items[0].id",
			input: "container.items[0].id",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "container.items[0].id",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "container",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"container\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "items",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"items\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "[0]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
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
			name:  "Generic with projection - items.{id}",
			input: "items.{id}",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "items.{id}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "items",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"items\""},
						},
					},
					{
						Type:     "ASTProject",
						Fragment: "{id}",
						Children: []ExpectedNode{
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
			},
		},
		{
			name:  "Generic with selection - items.{? #this.id > 100}",
			input: "items.{? #this.id > 100}",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "items.{? (#this.id > 100)}",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "items",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"items\""},
						},
					},
					{
						Type:     "ASTSelect",
						Fragment: "{? (#this.id > 100)}",
						Children: []ExpectedNode{
							{
								Type:     "ASTGreater",
								Fragment: "#this.id > 100",
								Children: []ExpectedNode{
									{
										Type:     "ASTChain",
										Fragment: "#this.id",
										Children: []ExpectedNode{
											{Type: "ASTThisVarRef", Fragment: "#this"},
											{
												Type:     "ASTProperty",
												Fragment: "id",
												Children: []ExpectedNode{
													{Type: "ASTConst", Fragment: "\"id\""},
												},
											},
										},
									},
									{Type: "ASTConst", Fragment: "100"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Array literal with Long type - new Long[] { 1L, 101L }",
			input: "new Long[] { 1L, 101L }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new Long[]{ 1L, 101L }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ 1L, 101L }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1L"},
							{Type: "ASTConst", Fragment: "101L"},
						},
					},
				},
			},
		},
		{
			name:  "Generic type comparison - value instanceof GameGenericObject",
			input: "value instanceof GameGenericObject",
			expected: ExpectedNode{
				Type:     "ASTInstanceof",
				Fragment: "value instanceof GameGenericObject",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"value\""},
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

// TestGenericAssignment 测试泛型对象的赋值表达式
func TestGenericAssignment(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Assign array to ids - ids = new Long[] { 1L, 101L }",
			input: "ids = new Long[] { 1L, 101L }",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "ids = new Long[]{ 1L, 101L }",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "ids",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"ids\""},
						},
					},
					{
						Type:     "ASTCtor",
						Fragment: "new Long[]{ 1L, 101L }",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ 1L, 101L }",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "1L"},
									{Type: "ASTConst", Fragment: "101L"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Assign to array element - ids[0] = 200L",
			input: "ids[0] = 200L",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "ids[0] = 200L",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "ids[0]",
						Children: []ExpectedNode{
							{
								Type:     "ASTProperty",
								Fragment: "ids",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"ids\""},
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
					{Type: "ASTConst", Fragment: "200L"},
				},
			},
		},
		{
			name:  "Assign generic value - value = new GameGenericObject()",
			input: "value = new GameGenericObject()",
			expected: ExpectedNode{
				Type:     "ASTAssign",
				Fragment: "value = new GameGenericObject()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "value",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"value\""},
						},
					},
					{
						Type:     "ASTCtor",
						Fragment: "new GameGenericObject()",
						Children: []ExpectedNode{},
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
