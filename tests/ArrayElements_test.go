package test

import (
	"testing"

	"github.com/weaweawe01/ParserOgnl/lexer"

	"github.com/weaweawe01/ParserOgnl/parser"
)

// TestArrayLiteralBasic 测试基本的数组字面量语法
func TestArrayLiteralBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Array with size",
			input: "new Object[5]",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new Object[5]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			name:  "Integer array",
			input: "new int[] { 10, 20 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[]{ 10, 20 }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ 10, 20 }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "10"},
							{Type: "ASTConst", Fragment: "20"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)
			p := parser.New(l)
			expr, err := p.ParseTopLevelExpression()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}
			if !Check(expr, tt.expected) {
				t.Errorf("AST check failed for: %s", tt.input)
			}
		})
	}
}

// TestArrayAccess 测试数组元素访问语法
func TestArrayAccess(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Array element access with variable",
			input: "#root[1]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "#root[1]",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
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
			name:  "Array length property",
			input: "array.length",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "array.length",
				Children: []ExpectedNode{
					{Type: "ASTProperty", Fragment: "array"},
					{Type: "ASTProperty", Fragment: "length"},
				},
			},
		},
		{
			name:  "String to char array",
			input: "\"Tapestry\".toCharArray()[2]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "\"Tapestry\".toCharArray()[2]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"Tapestry\""},
					{Type: "ASTMethod", Fragment: "toCharArray()"},
					{
						Type:     "ASTProperty",
						Fragment: "[2]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "2"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)
			p := parser.New(l)
			expr, err := p.ParseTopLevelExpression()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}
			if !Check(expr, tt.expected) {
				t.Errorf("AST check failed for: %s", tt.input)
			}
		})
	}
}

// TestCharArray 测试字符数组字面量
func TestCharArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Char array literal",
			input: "{'1','2','3'}",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ '1', '2', '3' }",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "'1'"},
					{Type: "ASTConst", Fragment: "'2'"},
					{Type: "ASTConst", Fragment: "'3'"},
				},
			},
		},
		{
			name:  "String char array access",
			input: "\"{Hello}\".toCharArray()[6]",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "\"{Hello}\".toCharArray()[6]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"{Hello}\""},
					{Type: "ASTMethod", Fragment: "toCharArray()"},
					{
						Type:     "ASTProperty",
						Fragment: "[6]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "6"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)
			p := parser.New(l)
			expr, err := p.ParseTopLevelExpression()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}
			if !Check(expr, tt.expected) {
				t.Errorf("AST check failed for: %s", tt.input)
			}
		})
	}
}

// TestBooleanArray 测试布尔数组字面量
func TestBooleanArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Boolean array with expressions",
			input: "{ true, !false }",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ true, !false }",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{
						Type:     "ASTNot",
						Fragment: "!false",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "false"},
						},
					},
				},
			},
		},
		{
			name:  "Boolean array literal",
			input: "{ true, true }",
			expected: ExpectedNode{
				Type:     "ASTList",
				Fragment: "{ true, true }",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "true"},
					{Type: "ASTConst", Fragment: "true"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)
			p := parser.New(l)
			expr, err := p.ParseTopLevelExpression()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}
			if !Check(expr, tt.expected) {
				t.Errorf("AST check failed for: %s", tt.input)
			}
		})
	}
}

// TestArrayNestedConstructor 测试嵌套的数组构造器
func TestArrayNestedConstructor(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Constructor with array parameter",
			input: "new ognl.test.objects.Simple(new Object[5])",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.objects.Simple(new Object[5])",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new Object[5]",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "5"},
						},
					},
				},
			},
		},
		{
			name:  "Constructor with array initializer",
			input: "new ognl.test.objects.Simple(new Object[] { #root, #this })",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.objects.Simple(new Object[]{ #root, #this })",
				Children: []ExpectedNode{
					{
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
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)
			p := parser.New(l)
			expr, err := p.ParseTopLevelExpression()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}
			if !Check(expr, tt.expected) {
				t.Errorf("AST check failed for: %s", tt.input)
			}
		})
	}
}

// TestArrayExpressionOutput 测试数组表达式的输出格式
func TestArrayExpressionOutput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "String array initialization",
			input:    "new String[] { \"one\", \"two\" }",
			expected: "new String[]{ \"one\", \"two\" }",
		},
		{
			name:     "Empty array initialization",
			input:    "new String[] { }",
			expected: "new String[]{  }",
		},
		{
			name:     "Array with size",
			input:    "new Object[5]",
			expected: "new Object[5]",
		},
		{
			name:     "Array literal",
			input:    "{ 1, 2, 3 }",
			expected: "{ 1, 2, 3 }",
		},
		{
			name:     "Empty array literal",
			input:    "{ }",
			expected: "{  }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.input)
			p := parser.New(l)
			expr, err := p.ParseTopLevelExpression()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}
			result := expr.String()
			if result != tt.expected {
				t.Errorf("Output mismatch:\nInput:    %s\nExpected: %s\nGot:      %s",
					tt.input, tt.expected, result)
			}
		})
	}
}
