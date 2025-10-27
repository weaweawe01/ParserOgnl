package test

import (
	"testing"
)

// TestPrimitiveArray 测试原始类型数组创建表达式
// 对应 Java 的 PrimitiveArrayTest.java
// 测试各种原始类型数组的创建，包括 boolean、char、byte、short、int、long、float、double
func TestPrimitiveArray(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		// Boolean 数组创建
		{
			name:  "Boolean array with size",
			input: "new boolean[5]",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new boolean[5]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "5"},
				},
			},
		},
		{
			name:  "Boolean array with initializer",
			input: "new boolean[] { true, false }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new boolean[]{ true, false }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ true, false }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "true"},
							{Type: "ASTConst", Fragment: "false"},
						},
					},
				},
			},
		},
		{
			name:  "Boolean array with numeric conversion",
			input: "new boolean[] { 0, 1, 5.5 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new boolean[]{ 0, 1, 5.5 }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ 0, 1, 5.5 }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "0"},
							{Type: "ASTConst", Fragment: "1"},
							{Type: "ASTConst", Fragment: "5.5"},
						},
					},
				},
			},
		},

		// Char 数组创建
		{
			name:  "Char array with character literals",
			input: "new char[] { 'a', 'b' }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new char[]{ 'a', 'b' }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ 'a', 'b' }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "'a'"},
							{Type: "ASTConst", Fragment: "'b'"},
						},
					},
				},
			},
		},
		{
			name:  "Char array with numeric values",
			input: "new char[] { 10, 11 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new char[]{ 10, 11 }",
				Children: []ExpectedNode{
					{
						Type:     "ASTList",
						Fragment: "{ 10, 11 }",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "10"},
							{Type: "ASTConst", Fragment: "11"},
						},
					},
				},
			},
		},

		// Byte 数组创建
		{
			name:  "Byte array with initializer",
			input: "new byte[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new byte[]{ 1, 2 }",
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

		// Short 数组创建
		{
			name:  "Short array with initializer",
			input: "new short[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new short[]{ 1, 2 }",
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

		// Int 数组创建
		{
			name:  "Int array with size from variable",
			input: "new int[six]",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[six]",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "six",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"six\""},
						},
					},
				},
			},
		},
		{
			name:  "Int array with size from root variable",
			input: "new int[#root.six]",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[#root.six]",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "#root.six",
						Children: []ExpectedNode{
							{Type: "ASTRootVarRef", Fragment: "#root"},
							{
								Type:     "ASTProperty",
								Fragment: "six",
								Children: []ExpectedNode{
									{Type: "ASTConst", Fragment: "\"six\""},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "Int array with numeric size",
			input: "new int[6]",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[6]",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "6"},
				},
			},
		},
		{
			name:  "Int array with initializer",
			input: "new int[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new int[]{ 1, 2 }",
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

		// Long 数组创建
		{
			name:  "Long array with initializer",
			input: "new long[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new long[]{ 1, 2 }",
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

		// Float 数组创建
		{
			name:  "Float array with initializer",
			input: "new float[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new float[]{ 1, 2 }",
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

		// Double 数组创建
		{
			name:  "Double array with initializer",
			input: "new double[] { 1, 2 }",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new double[]{ 1, 2 }",
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
