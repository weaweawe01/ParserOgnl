package test

import (
	"testing"
)

// TestPublicMember 测试公共成员表达式（基于 Java 的 PublicMemberTest.java）
func TestPublicMember(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Public property accessor - publicProperty",
			input: "publicProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "publicProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"publicProperty\""},
				},
			},
		},
		{
			name:  "Public field access - _publicProperty",
			input: "_publicProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_publicProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_publicProperty\""},
				},
			},
		},
		{
			name:  "Public final property accessor - publicFinalProperty",
			input: "publicFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "publicFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"publicFinalProperty\""},
				},
			},
		},
		{
			name:  "Public final field access - _publicFinalProperty",
			input: "_publicFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_publicFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_publicFinalProperty\""},
				},
			},
		},
		{
			name:  "Public static property accessor - publicStaticProperty",
			input: "publicStaticProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "publicStaticProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"publicStaticProperty\""},
				},
			},
		},
		{
			name:  "Public static final property accessor - publicStaticFinalProperty",
			input: "publicStaticFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "publicStaticFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"publicStaticFinalProperty\""},
				},
			},
		},
		{
			name:  "Public static field access - _publicStaticProperty",
			input: "_publicStaticProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_publicStaticProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_publicStaticProperty\""},
				},
			},
		},
		{
			name:  "Public static final field access - _publicStaticFinalProperty",
			input: "_publicStaticFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_publicStaticFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_publicStaticFinalProperty\""},
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
