package test

import (
	"testing"
)

// TestInheritedMethods 测试继承方法表达式
// 对应 Java 的 InheritedMethodsTest.java
// 测试转换继承方法表达式的功能
func TestInheritedMethods(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Base inheritance - map.bean.name",
			input: "map.bean.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.bean.name",
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
						Fragment: "bean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean\""},
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
		{
			name:  "Method call on inherited object - bean.getName()",
			input: "bean.getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "bean.getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "bean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean\""},
						},
					},
					{Type: "ASTMethod", Fragment: "getName()"},
				},
			},
		},
		{
			name:  "Property access on base class - baseBean.value",
			input: "baseBean.value",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "baseBean.value",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "baseBean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"baseBean\""},
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
			name:  "Chained property access - first.second.name",
			input: "first.second.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "first.second.name",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "first",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"first\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "second",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"second\""},
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
		{
			name:  "Method on map value - map['key'].getValue()",
			input: "map['key'].getValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map[\"key\"].getValue()",
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
						Fragment: "[\"key\"]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"key\""},
						},
					},
					{Type: "ASTMethod", Fragment: "getValue()"},
				},
			},
		},
		{
			name:  "Nested property access - root.map.bean.name",
			input: "root.map.bean.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "root.map.bean.name",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "root",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"root\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "map",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"map\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "bean",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"bean\""},
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
		{
			name:  "Property getter method - object.getName().toString()",
			input: "object.getName().toString()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "object.getName().toString()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "object",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"object\""},
						},
					},
					{Type: "ASTMethod", Fragment: "getName()"},
					{Type: "ASTMethod", Fragment: "toString()"},
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
