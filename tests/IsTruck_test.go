package test

import (
	"testing"
)

// TestIsTruck 测试 isTruck 属性访问
// 对应 Java 的 IsTruckTest.java
// 测试访问 getIsTruck() 方法时使用 "isTruck" 属性名
func TestIsTruck(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "IsTruck property access - isTruck",
			input: "isTruck",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "isTruck",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"isTruck\""},
				},
			},
		},
		{
			name:  "GetIsTruck method call - getIsTruck()",
			input: "getIsTruck()",
			expected: ExpectedNode{
				Type:     "ASTMethod",
				Fragment: "getIsTruck()",
			},
		},
		{
			name:  "IsTruck with object - holder.isTruck",
			input: "holder.isTruck",
			expected: ExpectedNode{
				Type:     "ASTChain",
				Fragment: "holder.isTruck",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "holder",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"holder\""},
						},
					},
					{
						Type:     "ASTProperty",
						Fragment: "isTruck",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"isTruck\""},
						},
					},
				},
			},
		},
		{
			name:  "Negated isTruck - !isTruck",
			input: "!isTruck",
			expected: ExpectedNode{
				Type:     "ASTNot",
				Fragment: "!isTruck",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "isTruck",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"isTruck\""},
						},
					},
				},
			},
		},
		{
			name:  "IsTruck in conditional - isTruck ? 'yes' : 'no'",
			input: "isTruck ? 'yes' : 'no'",
			expected: ExpectedNode{
				Type:     "ASTTest",
				Fragment: "isTruck ? \"yes\" : \"no\"",
				Children: []ExpectedNode{
					{
						Type:     "ASTProperty",
						Fragment: "isTruck",
						Children: []ExpectedNode{
							{Type: "ASTConst", Fragment: "\"isTruck\""},
						},
					},
					{Type: "ASTConst", Fragment: "\"yes\""},
					{Type: "ASTConst", Fragment: "\"no\""},
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
