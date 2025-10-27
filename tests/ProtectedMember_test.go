package test

import (
	"testing"
)

// TestProtectedMember 测试受保护成员表达式（基于 Java 的 ProtectedMemberTest.java）
func TestProtectedMember(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected ExpectedNode
	}{
		{
			name:  "Protected property accessor - protectedProperty",
			input: "protectedProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "protectedProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"protectedProperty\""},
				},
			},
		},
		{
			name:  "Protected field access - _protectedProperty",
			input: "_protectedProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_protectedProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_protectedProperty\""},
				},
			},
		},
		{
			name:  "Protected final property accessor - protectedFinalProperty",
			input: "protectedFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "protectedFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"protectedFinalProperty\""},
				},
			},
		},
		{
			name:  "Protected final field access - _protectedFinalProperty",
			input: "_protectedFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_protectedFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_protectedFinalProperty\""},
				},
			},
		},
		{
			name:  "Protected static property accessor - protectedStaticProperty",
			input: "protectedStaticProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "protectedStaticProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"protectedStaticProperty\""},
				},
			},
		},
		{
			name:  "Protected static final property accessor - protectedStaticFinalProperty",
			input: "protectedStaticFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "protectedStaticFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"protectedStaticFinalProperty\""},
				},
			},
		},
		{
			name:  "Protected static field access - _protectedStaticProperty",
			input: "_protectedStaticProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_protectedStaticProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_protectedStaticProperty\""},
				},
			},
		},
		{
			name:  "Protected static final field access - _protectedStaticFinalProperty",
			input: "_protectedStaticFinalProperty",
			expected: ExpectedNode{
				Type:     "ASTProperty",
				Fragment: "_protectedStaticFinalProperty",
				Children: []ExpectedNode{
					{Type: "ASTConst", Fragment: "\"_protectedStaticFinalProperty\""},
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
