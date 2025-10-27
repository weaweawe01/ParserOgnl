package test

import (
	"testing"
)

// TestStaticsAndConstructors 测试静态方法和构造函数（基于 Java 的 StaticsAndConstructorsTest.java）
// 测试静态成员访问、静态方法调用、构造函数调用等表达式，包括：
// - 静态方法调用（@ClassName@methodName()）
// - 静态字段访问（@ClassName@FIELD_NAME）
// - 构造函数调用（new ClassName()）
// - class 属性访问
// - 枚举值访问
func TestStaticsAndConstructors(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		// testClassForName - 测试静态方法调用 Class.forName
		{
			name:  "Class.forName static method",
			input: "@java.lang.Class@forName(\"java.lang.Object\")",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@java.lang.Class@forName(\"java.lang.Object\")",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"java.lang.Object\""},
				},
			},
		},
		// testIntegerMaxValue - 测试静态字段访问
		{
			name:  "Integer.MAX_VALUE static field",
			input: "@java.lang.Integer@MAX_VALUE",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@java.lang.Integer@MAX_VALUE",
			},
		},
		// testMaxFunction - 测试 Math.max 静态方法（使用 @@ 简写）
		{
			name:  "Math.max with @@ shorthand",
			input: "@@max(3,4)",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@java.lang.Math@max(3, 4)",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "3"},
					{Type: "ASTConst", Fragment: "4"},
				},
			},
		},
		// testStringBuffer - 测试构造函数和链式方法调用
		{
			name:  "StringBuffer constructor with method chain",
			input: "new java.lang.StringBuffer().append(55).toString()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "new java.lang.StringBuffer().append(55).toString()",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new java.lang.StringBuffer()",
					},
					{
						Type:     "ASTMethod",
						Fragment: "append(55)",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "55"},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "toString()",
					},
				},
			},
		},
		// testClass - 测试 class 属性访问
		{
			name:  "class property",
			input: "class",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "class",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"class\""},
				},
			},
		},
		// testRootClass - 测试静态 class 属性访问
		{
			name:  "static class property",
			input: "@ognl.test.objects.Root@class",
			expected: ExpectedNode{
				Type:     "ASTStaticField",
				Fragment: "@ognl.test.objects.Root@class",
			},
		},
		// testClassName - 测试 class.getName() 方法调用
		{
			name:  "class.getName()",
			input: "class.getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "class.getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "class",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"class\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		// testRootClassName - 测试静态 class.getName()
		{
			name:  "static class.getName()",
			input: "@ognl.test.objects.Root@class.getName()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@ognl.test.objects.Root@class.getName()",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@ognl.test.objects.Root@class",
					},
					{
						Type:     "ASTMethod",
						Fragment: "getName()",
					},
				},
			},
		},
		// testRootClassNameProperty - 测试静态 class.name 属性
		{
			name:  "static class.name property",
			input: "@ognl.test.objects.Root@class.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "@ognl.test.objects.Root@class.name",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@ognl.test.objects.Root@class",
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
		// testClassSuperclass - 测试 class.getSuperclass()
		{
			name:  "class.getSuperclass()",
			input: "class.getSuperclass()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "class.getSuperclass()",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "class",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"class\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getSuperclass()",
					},
				},
			},
		},
		// testClassSuperclassProperty - 测试 class.superclass 属性
		{
			name:  "class.superclass property",
			input: "class.superclass",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "class.superclass",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "class",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"class\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "superclass",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"superclass\""},
						},
					},
				},
			},
		},
		// testClassNameProperty - 测试 class.name 属性
		{
			name:  "class.name property",
			input: "class.name",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "class.name",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "class",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"class\""},
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
		// testStaticInt - 测试静态方法调用
		{
			name:  "static method getStaticInt()",
			input: "getStaticInt()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getStaticInt()",
			},
		},
		// testRootStaticInt - 测试完全限定的静态方法调用
		{
			name:  "fully qualified static method",
			input: "@ognl.test.objects.Root@getStaticInt()",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@ognl.test.objects.Root@getStaticInt()",
			},
		},
		// testSimpleStringValue - 测试带参数的构造函数
		{
			name:  "constructor with property argument",
			input: "new ognl.test.objects.Simple(property).getStringValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "new ognl.test.objects.Simple(property).getStringValue()",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new ognl.test.objects.Simple(property)",
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
					{
						Type:     "ASTMethod",
						Fragment: "getStringValue()",
					},
				},
			},
		},
		// testSimpleStringValueWithMap - 测试构造函数中的复杂表达式参数
		{
			name:  "constructor with map chain argument",
			input: "new ognl.test.objects.Simple(map['test'].property).getStringValue()",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "new ognl.test.objects.Simple(map[\"test\"].property).getStringValue()",
				Children: []ExpectedNode{
					{
						Type:     "ASTCtor",
						Fragment: "new ognl.test.objects.Simple(map[\"test\"].property)",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "map[\"test\"].property",
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
					{
						Type:     "ASTMethod",
						Fragment: "getStringValue()",
					},
				},
			},
		},
		// testMapCurrentClass - 测试复杂的静态字段引用作为参数
		{
			name:  "method with static field as argument",
			input: "map.test.getCurrentClass(@ognl.test.StaticsAndConstructorsTest@KEY.toString())",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "map.test.getCurrentClass(@ognl.test.StaticsAndConstructorsTest@KEY.toString())",
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
						Fragment: "test",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"test\""},
						},
					},
					{
						Type:     "ASTMethod",
						Fragment: "getCurrentClass(@ognl.test.StaticsAndConstructorsTest@KEY.toString())",
						Children: []ExpectedNode{
							{
								Type:     "ASTChain",
								Fragment: "@ognl.test.StaticsAndConstructorsTest@KEY.toString()",
								Children: []ExpectedNode{
									{
										Type:     "ASTStaticField",
										Fragment: "@ognl.test.StaticsAndConstructorsTest@KEY",
									},
									{
										Type:     "ASTMethod",
										Fragment: "toString()",
									},
								},
							},
						},
					},
				},
			},
		},
		// testIntWrapper - 测试构造函数带属性参数
		{
			name:  "constructor with property index",
			input: "new ognl.test.StaticsAndConstructorsTest$IntWrapper(index)",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.StaticsAndConstructorsTest$IntWrapper(index)",
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
		// testIntObjectWrapper - 测试内部类构造函数
		{
			name:  "inner class constructor",
			input: "new ognl.test.StaticsAndConstructorsTest$IntObjectWrapper(index)",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.StaticsAndConstructorsTest$IntObjectWrapper(index)",
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
		// testA - 测试构造函数带变量引用参数
		{
			name:  "constructor with #root variable",
			input: "new ognl.test.StaticsAndConstructorsTest$A(#root)",
			expected: ExpectedNode{
				Type:     "ASTCtor",
				Fragment: "new ognl.test.StaticsAndConstructorsTest$A(#root)",
				Children: []ExpectedNode{
					{Type: "ASTRootVarRef", Fragment: "#root"},
				},
			},
		},
		// testAnimalsValues - 测试枚举的 values() 方法和属性访问
		{
			name:  "enum values() method and length",
			input: "@ognl.test.StaticsAndConstructorsTest$Animals@values().length != 2",
			expected: ExpectedNode{
				Type:     "ASTNotEq",
				Fragment: "@ognl.test.StaticsAndConstructorsTest$Animals@values().length != 2",
				Children: []ExpectedNode{
					{
						Type:     "ASTChain",
						Fragment: "@ognl.test.StaticsAndConstructorsTest$Animals@values().length",
						Children: []ExpectedNode{
							{
								Type:     "ASTStaticMethod",
								Fragment: "@ognl.test.StaticsAndConstructorsTest$Animals@values()",
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
					{Type: "ASTConst", Fragment: "2"},
				},
			},
		},
		// testIsOk - 测试静态字段作为方法参数
		{
			name:  "method with static enum field and null",
			input: "isOk(@ognl.test.objects.SimpleEnum@ONE, null)",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "isOk(@ognl.test.objects.SimpleEnum@ONE, null)",
				Children: []ExpectedNode{
					{
						Type:     "ASTStaticField",
						Fragment: "@ognl.test.objects.SimpleEnum@ONE",
					},
					{Type: "ASTConst", Fragment: "null"},
				},
			},
		},
		// testStaticMethod - 测试接口的静态方法调用
		{
			name:  "interface static method",
			input: "@ognl.test.objects.StaticInterface@staticMethod()",
			expected: ExpectedNode{
				Type:     "ASTStaticMethod",
				Fragment: "@ognl.test.objects.StaticInterface@staticMethod()",
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
