package test

import (
	"testing"
)

// TestArrayCreation 测试数组创建表达式
func TestArrayCreation(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "new String[] { \"one\", \"two\" }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new String[]{ \"one\", \"two\" }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ \"one\", \"two\" }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"one\""},
							{Type: "ASTConst", Fragment: "\"two\""},
						},
					},
				},
			},
		},
		{
			input: "new String[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new String[]{ 1, 2 }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ 1, 2 }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "1"},
							{Type: "ASTConst", Fragment: "2"},
						},
					},
				},
			},
		},
		{
			input: "new Integer[] { 1, 2, 3 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new Integer[]{ 1, 2, 3 }",
				Children: []ExpectedNode{
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
		{
			input: "new Object[] { #root, #this }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new Object[]{ #root, #this }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ #root, #this }",
						Children: []ExpectedNode{
							{Type: "ASTRootVarRef", Fragment: "#root"},
							{Type: "ASTThisVarRef", Fragment: "#this"},
						},
					},
				},
			},
		},
		{
			input: "new ognl.test.objects.Simple[] { new ognl.test.objects.Simple(), new ognl.test.objects.Simple(\"foo\", 1.0f, 2) }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.objects.Simple[]{ new ognl.test.objects.Simple(), new ognl.test.objects.Simple(\"foo\", 1.0, 2) }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ new ognl.test.objects.Simple(), new ognl.test.objects.Simple(\"foo\", 1.0, 2) }",
						Children: []ExpectedNode{
							{
								Type:     "ASTCtor",
								Fragment: "new ognl.test.objects.Simple()",
							},
							{
								Type:     "ASTCtor",
								Fragment: "new ognl.test.objects.Simple(\"foo\", 1.0, 2)",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"foo\""},
									{Type: "ASTConst", Fragment: "1.0"},
									{Type: "ASTConst", Fragment: "2"},
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

// TestArrayCreationWithSize 测试带大小的数组创建（暂不实现，仅作占位）
// 注意：new String[10] 这种语法需要特殊处理，当前解析器可能不支持
func TestArrayCreationWithSize(t *testing.T) {
	t.Skip("带大小的数组创建语法 (new String[10]) 暂未实现")
}

// TestConditionalArrayCreation 测试条件表达式中的数组创建
func TestConditionalArrayCreation(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "true ? new Object[] { 1, 2 } : new Object[] { 3, 4 }",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "true ? new Object[]{ 1, 2 } : new Object[]{ 3, 4 }",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{
						Type:     "ASTCtor",
						Fragment: "new Object[]{ 1, 2 }",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ 1, 2 }",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "1"},
									{Type: "ASTConst", Fragment: "2"},
								},
							},
						},
					},
					{
						Type:     "ASTCtor",
						Fragment: "new Object[]{ 3, 4 }",
						Children: []ExpectedNode{
							{
								Type:     "ASTList",
								Fragment: "{ 3, 4 }",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "3"},
									{Type: "ASTConst", Fragment: "4"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// TestNestedArrayCreation 测试嵌套的数组创建
func TestNestedArrayCreation(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "new ognl.test.objects.Simple(new Object[5])",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.objects.Simple(new Object[5])",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new Object[5]",
					},
				},
			},
		},
		{
			input: "new ognl.test.objects.Simple(new String[5])",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.objects.Simple(new String[5])",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new String[5]",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}

// TestEmptyArrayCreation 测试空数组创建
func TestEmptyArrayCreation(t *testing.T) {
	testCases := []struct {
		input    string
		expected ExpectedNode
	}{
		{
			input: "new String[] { }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new String[]{  }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{  }",
					},
				},
			},
		},
		{
			input: "new Object[] { }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new Object[]{  }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{  }",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			expr := parseExpression(t, tc.input)
			if !Check(expr, tc.expected) {
				t.Errorf("表达式 '%s' 的 AST 检查失败", tc.input)
			}
		})
	}
}
