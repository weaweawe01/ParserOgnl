package test

import (
	"testing"
)

// TestDefaultClassResolver 测试默认类解析器
// 对应 Java 的 DefaultClassResolverTest.java
// 注意：这些测试主要验证类名解析的 AST 结构，而不是运行时类加载
func TestDefaultClassResolver(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Class in default package - new ClassInDefaultPackage()",
			input: "new ClassInDefaultPackage()",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ClassInDefaultPackage()",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "Static method reference - @java.lang.Math@max(1, 2)",
			input: "@java.lang.Math@max(1, 2)",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@java.lang.Math@max(1, 2)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "1"},
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		{
			name:  "Static field reference - @java.lang.Integer@MAX_VALUE",
			input: "@java.lang.Integer@MAX_VALUE",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@java.lang.Integer@MAX_VALUE",
			},
		},
		{
			name:  "Constructor with package name - new java.util.ArrayList()",
			input: "new java.util.ArrayList()",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new java.util.ArrayList()",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "Constructor with package and arguments - new java.lang.String(\"test\")",
			input: "new java.lang.String(\"test\")",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new java.lang.String(\"test\")",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"test\""},
				},
			},
		},
		{
			name:  "Static method with multiple package levels - @com.example.utils.Helper@format(value)",
			input: "@com.example.utils.Helper@format(value)",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@com.example.utils.Helper@format(value)",
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
		{
			name:  "Class name with single letter - @X@Y",
			input: "@X@Y",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@X@Y",
			},
		},
		{
			name:  "Deep package hierarchy - @org.apache.commons.lang3.StringUtils@isEmpty(str)",
			input: "@org.apache.commons.lang3.StringUtils@isEmpty(str)",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@org.apache.commons.lang3.StringUtils@isEmpty(str)",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "str",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"str\""},
						},
					},
				},
			},
		},
		{
			name:  "Constructor with array type - new int[10]",
			input: "new int[10]",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[10]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "10"},
				},
			},
		},
		{
			name:  "Constructor with array initializer - new int[] { 1, 2, 3 }",
			input: "new int[] { 1, 2, 3 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[]{ 1, 2, 3 }",
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
			name:  "Static method chained with instance method - @Factory@create().process()",
			input: "@Factory@create().process()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@Factory@create().process()",
				Children: []ExpectedNode{
					{Type: "ASTStaticMethod", Fragment: "@Factory@create()"},
					{Type: "ASTMethod", Fragment: "process()"},
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

// TestClassResolverErrorCases 测试类解析器的错误情况
// 注意：由于 Go 版本的 OGNL 只做语法解析，不做运行时类加载，
// 所以这里只测试语法解析是否正确，不测试 ClassNotFoundException
func TestClassResolverErrorCases(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Non-existent class reference - @no.such.Class@method()",
			input: "@no.such.Class@method()",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@no.such.Class@method()",
				Children: []ExpectedNode{},
			},
		},
		{
			name:  "Bogus class constructor - new BogusClass()",
			input: "new BogusClass()",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new BogusClass()",
				Children: []ExpectedNode{},
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
