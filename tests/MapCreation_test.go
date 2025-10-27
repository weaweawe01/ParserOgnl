package test

import (
	"testing"
)

// TestMapCreation 测试Map创建表达式
func TestMapCreation(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Simple map creation with one key-value pair",
			input: `#{ "foo" : "bar" }`,
			expected: ExpectedNode{
				Type:     "ASTMap",
				Fragment: `#{ "foo" : "bar" }`,
				Children: []ExpectedNode{
					{
						Type:     "ASTKeyValue",
						Fragment: `"foo" : "bar"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"foo"`},
							{Type: "ASTConst", Fragment: `"bar"`},
						},
					},
				},
			},
		},
		{
			name:  "Map creation with multiple key-value pairs",
			input: `#{ "foo" : "bar", "bar" : "baz" }`,
			expected: ExpectedNode{
				Type:     "ASTMap",
				Fragment: `#{ "foo" : "bar", "bar" : "baz" }`,
				Children: []ExpectedNode{
					{
						Type:     "ASTKeyValue",
						Fragment: `"foo" : "bar"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"foo"`},
							{Type: "ASTConst", Fragment: `"bar"`},
						},
					},
					{
						Type:     "ASTKeyValue",
						Fragment: `"bar" : "baz"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"bar"`},
							{Type: "ASTConst", Fragment: `"baz"`},
						},
					},
				},
			},
		},
		{
			name:  "Map creation with null value",
			input: `#{ "foo", "bar" : "baz" }`,
			expected: ExpectedNode{
				Type:     "ASTMap",
				Fragment: `#{ "foo" : null, "bar" : "baz" }`,
				Children: []ExpectedNode{
					{
						Type:     "ASTKeyValue",
						Fragment: `"foo" : null`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"foo"`},
						},
					},
					{
						Type:     "ASTKeyValue",
						Fragment: `"bar" : "baz"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"bar"`},
							{Type: "ASTConst", Fragment: `"baz"`},
						},
					},
				},
			},
		},
		{
			name:  "Map creation with LinkedHashMap type",
			input: `#@java.util.LinkedHashMap@{ "foo" : "bar", "bar" : "baz" }`,
			expected: ExpectedNode{
				Type:     "ASTMap",
				Fragment: `#@java.util.LinkedHashMap@{ "foo" : "bar", "bar" : "baz" }`,
				Children: []ExpectedNode{
					{
						Type:     "ASTKeyValue",
						Fragment: `"foo" : "bar"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"foo"`},
							{Type: "ASTConst", Fragment: `"bar"`},
						},
					},
					{
						Type:     "ASTKeyValue",
						Fragment: `"bar" : "baz"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"bar"`},
							{Type: "ASTConst", Fragment: `"baz"`},
						},
					},
				},
			},
		},
		{
			name:  "Map creation with TreeMap type",
			input: `#@java.util.TreeMap@{ "foo" : "bar", "bar" : "baz" }`,
			expected: ExpectedNode{
				Type:     "ASTMap",
				Fragment: `#@java.util.TreeMap@{ "foo" : "bar", "bar" : "baz" }`,
				Children: []ExpectedNode{
					{
						Type:     "ASTKeyValue",
						Fragment: `"foo" : "bar"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"foo"`},
							{Type: "ASTConst", Fragment: `"bar"`},
						},
					},
					{
						Type:     "ASTKeyValue",
						Fragment: `"bar" : "baz"`,
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: `"bar"`},
							{Type: "ASTConst", Fragment: `"baz"`},
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
