package test

import (
	"fmt"
	"strings"

	"github.com/weaweawe01/ParserOgnl/ast"
)

// ExpectedNode 表示期望的节点信息
type ExpectedNode struct {
	Type     string         // 节点类型,如 "ASTGreater", "ASTProperty", "ASTConst"
	Fragment string         // 表达式片段,如 "a > n", "a", "n"
	Children []ExpectedNode // 子节点
}

// CheckResult 检查结果
type CheckResult struct {
	Success bool
	Errors  []string
}

// Check 检查AST树是否符合预期
// expr: 要检查的AST表达式
// expected: 期望的节点结构
// returns: true表示匹配成功,false表示不匹配,并打印出不匹配的地方
func Check(expr ast.Expression, expected ExpectedNode) bool {
	result := CheckResult{
		Success: true,
		Errors:  make([]string, 0),
	}

	checkNode(expr, expected, "", &result)

	if !result.Success {
		fmt.Println("=== AST检查失败 ===")
		for _, err := range result.Errors {
			fmt.Println(err)
		}
		fmt.Println("==================")
	} else {
		fmt.Println("✓ AST检查成功")
	}

	return result.Success
}

// checkNode 递归检查节点
func checkNode(node ast.Expression, expected ExpectedNode, path string, result *CheckResult) {
	if node == nil {
		result.Success = false
		result.Errors = append(result.Errors, fmt.Sprintf("[%s] 节点为nil,但期望类型为: %s", path, expected.Type))
		return
	}

	// 检查节点类型
	actualType := node.Type()
	if actualType != expected.Type {
		result.Success = false
		result.Errors = append(result.Errors,
			fmt.Sprintf("[%s] 类型不匹配: 期望 %s, 实际 %s", path, expected.Type, actualType))
		return
	}

	// 检查表达式片段
	actualFragment := node.String()
	if expected.Fragment != "" && !matchFragment(actualFragment, expected.Fragment) {
		result.Success = false
		result.Errors = append(result.Errors,
			fmt.Sprintf("[%s] 表达式片段不匹配: 期望 %q, 实际 %q", path, expected.Fragment, actualFragment))
		return
	}

	// 检查子节点
	if len(expected.Children) > 0 {
		checkChildren(node, expected.Children, path, result)
	}
}

// matchFragment 匹配表达式片段,允许一定的格式差异
func matchFragment(actual, expected string) bool {
	// 去除空格进行比较
	actualTrimmed := strings.TrimSpace(actual)
	expectedTrimmed := strings.TrimSpace(expected)

	// 移除括号进行比较
	actualNoParen := strings.Trim(actualTrimmed, "()")
	expectedNoParen := strings.Trim(expectedTrimmed, "()")

	return actualTrimmed == expectedTrimmed || actualNoParen == expectedNoParen || actualTrimmed == expectedNoParen || actualNoParen == expectedTrimmed
}

// checkChildren 检查子节点
func checkChildren(node ast.Expression, expectedChildren []ExpectedNode, path string, result *CheckResult) {
	var actualChildren []ast.Expression

	// 根据节点类型提取子节点
	switch n := node.(type) {
	case *ast.BinaryExpression:
		actualChildren = []ast.Expression{n.Left, n.Right}
	case *ast.UnaryExpression:
		actualChildren = []ast.Expression{n.Operand}
	case *ast.ChainExpression:
		actualChildren = n.Children
	case *ast.IndexExpression:
		// 如果 Object 为 nil（在 ChainExpression 中），只包含 Index
		if n.Object != nil {
			actualChildren = []ast.Expression{n.Object, n.Index}
		} else {
			actualChildren = []ast.Expression{n.Index}
		}
	case *ast.CallExpression:
		if n.Object != nil {
			actualChildren = append(actualChildren, n.Object)
		}
		actualChildren = append(actualChildren, n.Arguments...)
	case *ast.AssignmentExpression:
		actualChildren = []ast.Expression{n.Left, n.Right}
	case *ast.ConditionalExpression:
		actualChildren = []ast.Expression{n.Test, n.Consequent, n.Alternative}
	case *ast.ProjectionExpression:
		// 只包含非nil的子节点
		if n.Object != nil {
			actualChildren = append(actualChildren, n.Object)
		}
		if n.Expression != nil {
			actualChildren = append(actualChildren, n.Expression)
		}
	case *ast.SelectionExpression:
		// 只包含非nil的子节点
		if n.Object != nil {
			actualChildren = append(actualChildren, n.Object)
		}
		if n.Expression != nil {
			actualChildren = append(actualChildren, n.Expression)
		}
	case *ast.EvalExpression:
		// EvalExpression 有两个子节点: Target 和 Argument
		actualChildren = []ast.Expression{n.Target, n.Argument}
	case *ast.SequenceExpression:
		actualChildren = n.Expressions
	case *ast.ArrayExpression:
		actualChildren = n.Elements
	case *ast.MapExpression:
		actualChildren = n.Pairs
	case *ast.KeyValueExpression:
		// 如果 Value 为 nil，只包含 Key
		if n.Value != nil {
			actualChildren = []ast.Expression{n.Key, n.Value}
		} else {
			actualChildren = []ast.Expression{n.Key}
		}
	case *ast.Identifier:
		// Identifier可能有NameNode子节点
		if n.NameNode != nil {
			actualChildren = []ast.Expression{n.NameNode}
		}
	case *ast.LambdaLiteral:
		// LambdaLiteral 是 ASTConst 类型，有一个子节点（Lambda body）
		actualChildren = []ast.Expression{n.Body}
	case *ast.ConstructorExpression:
		actualChildren = n.Arguments
	case *ast.InstanceofExpression:
		actualChildren = []ast.Expression{n.Operand}
	case *ast.StaticMethodExpression:
		actualChildren = n.Arguments
	case *ast.DynamicSubscriptExpression:
		actualChildren = []ast.Expression{n.Object}
	}

	// 检查子节点数量
	if len(actualChildren) != len(expectedChildren) {
		result.Success = false
		result.Errors = append(result.Errors,
			fmt.Sprintf("[%s] 子节点数量不匹配: 期望 %d 个, 实际 %d 个",
				path, len(expectedChildren), len(actualChildren)))
		return
	}

	// 递归检查每个子节点
	for i, expectedChild := range expectedChildren {
		childPath := fmt.Sprintf("%s[%d]", path, i)
		if path == "" {
			childPath = fmt.Sprintf("child[%d]", i)
		}
		checkNode(actualChildren[i], expectedChild, childPath, result)
	}
}

// CheckSimple 简化版检查函数,只检查类型和片段,不检查子节点
func CheckSimple(expr ast.Expression, expectedType string, expectedFragment string) bool {
	expected := ExpectedNode{
		Type:     expectedType,
		Fragment: expectedFragment,
	}
	return Check(expr, expected)
}
