package parser

import (
	"fmt"
	"strings"

	"github.com/weaweawe01/ParserOgnl/ast"
)

// PrintASTStructure 打印详细的AST结构
func PrintASTStructure(expr ast.Expression, depth int) {
	indent := strings.Repeat("  ", depth)

	switch node := expr.(type) {
	case *ast.BinaryExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		fmt.Printf("%s    ", indent)
		// 打印左子节点
		printBinaryExpressionChild(node.Left, depth+1, node, false)
		fmt.Printf("%s    ", indent)
		// 打印右子节点
		printBinaryExpressionChild(node.Right, depth+1, node, true)
	case *ast.UnaryExpression:
		// 一元表达式的片段只显示操作数,不包含运算符(与 Java OGNL 一致)
		fmt.Printf("%s%s 表达式片段: %s\n", indent, node.Type(), node.Operand.String())
		fmt.Printf("%s    ", indent)
		PrintASTStructure(node.Operand, depth+1)
	case *ast.Literal:
		fmt.Printf("%s 表达式片段: %s\n", node.Type(), node.String())
	case *ast.LambdaLiteral:
		// LambdaLiteral 是 ASTConst 类型，有一个子节点（Lambda body）
		fmt.Printf("%s 表达式片段: %s\n", node.Type(), node.String())
		if node.Body != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Body, depth+1)
		}
	case *ast.Identifier:
		fmt.Printf("%s 表达式片段: %s\n", node.Type(), node.Value)
		// 打印 Identifier 的 NameNode 子节点 (对应Java的ASTConst)
		if node.NameNode != nil {
			fmt.Printf("%s   ", indent)
			PrintASTStructure(node.NameNode, depth+1)
		}
	case *ast.ConditionalExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		fmt.Printf("%s    ", indent)
		PrintASTStructure(node.Test, depth+1)
		fmt.Printf("%s    ", indent)
		PrintASTStructure(node.Consequent, depth+1)
		fmt.Printf("%s    ", indent)
		PrintASTStructure(node.Alternative, depth+1)
	case *ast.AssignmentExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		fmt.Printf("%s    ", indent)
		PrintASTStructure(node.Left, depth+1)
		fmt.Printf("%s    ", indent)
		PrintASTStructure(node.Right, depth+1)
	case *ast.SequenceExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		for _, expr := range node.Expressions {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(expr, depth+1)
		}
	case *ast.ChainExpression:
		if depth == 0 {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		} else {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		}
		// 遍历所有子节点
		for _, child := range node.Children {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(child, depth+1)
		}
	case *ast.IndexExpression:
		// IndexExpression 是 ASTProperty 的 Go 实现
		if node.Object == nil {
			// 作为 ChainExpression 的子节点，只显示索引部分
			fmt.Printf("%s 表达式片段: %s\n", node.Type(), node.String())
			if node.Index != nil {
				fmt.Printf("%s    ", indent)
				PrintASTStructure(node.Index, depth+1)
			}
		} else {
			// 独立的索引表达式
			fmt.Printf("%s 表达式片段: %s\n", node.Type(), node.String())
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Object, depth+1)
			if node.Index != nil {
				fmt.Printf("%s    ", indent)
				PrintASTStructure(node.Index, depth+1)
			}
		}
	case *ast.CallExpression:
		if depth == 0 {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		} else {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		}
		if len(node.Arguments) > 0 {
			fmt.Printf("%s    ", indent)
			for _, arg := range node.Arguments {
				PrintASTStructure(arg, depth+1)
			}
		}
	case *ast.StaticMethodExpression:
		if depth == 0 {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		} else {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		}
		if len(node.Arguments) > 0 {
			fmt.Printf("%s    ", indent)
			for _, arg := range node.Arguments {
				PrintASTStructure(arg, depth+1)
			}
		}
	case *ast.ArrayExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if len(node.Elements) > 0 {
			for _, elem := range node.Elements {
				fmt.Printf("%s    ", indent)
				PrintASTStructure(elem, depth+1)
			}
		}
	case *ast.MapExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if len(node.Pairs) > 0 {
			for _, pair := range node.Pairs {
				fmt.Printf("%s    ", indent)
				PrintASTStructure(pair, depth+1)
			}
		}
	case *ast.KeyValueExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if node.Key != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Key, depth+1)
		}
		if node.Value != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Value, depth+1)
		}
	case *ast.ConstructorExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if len(node.Arguments) > 0 {
			for _, arg := range node.Arguments {
				fmt.Printf("%s    ", indent)
				PrintASTStructure(arg, depth+1)
			}
		}
	case *ast.EvalExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if node.Target != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Target, depth+1)
		}
		if node.Argument != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Argument, depth+1)
		}
	case *ast.VariableExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
	case *ast.ThisExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
	case *ast.RootExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
	case *ast.SelectionExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if node.Expression != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Expression, depth+1)
		}
	case *ast.ProjectionExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if node.Expression != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Expression, depth+1)
		}
	case *ast.InstanceofExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if node.Operand != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Operand, depth+1)
		}
	case *ast.LambdaExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		if node.Body != nil {
			fmt.Printf("%s    ", indent)
			PrintASTStructure(node.Body, depth+1)
		}
	case *ast.DynamicSubscriptExpression:
		fmt.Printf("%s  %s 表达式片段: %s\n", indent, node.Type(), node.String())
		// DynamicSubscriptExpression 没有子节点,只是一个特殊的下标标记
	default:
		if depth == 0 {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, expr.Type(), expr.String())
		} else {
			fmt.Printf("%s  %s 表达式片段: %s\n", indent, expr.Type(), expr.String())
		}
	}
}

// printBinaryExpressionChild 打印二元表达式的子节点,考虑父节点上下文
func printBinaryExpressionChild(expr ast.Expression, depth int, parentBinary *ast.BinaryExpression, isRightChild bool) {
	indent := strings.Repeat("  ", depth)

	// 如果是二元表达式,使用 StringWithContext
	if binExpr, ok := expr.(*ast.BinaryExpression); ok {
		exprStr := binExpr.StringWithContext(parentBinary.GetPrecedence(), isRightChild)
		fmt.Printf("%s 表达式片段: %s\n", binExpr.Type(), exprStr)
		// 递归打印子节点
		fmt.Printf("%s      ", indent)
		printBinaryExpressionChild(binExpr.Left, depth+1, binExpr, false)
		fmt.Printf("%s      ", indent)
		printBinaryExpressionChild(binExpr.Right, depth+1, binExpr, true)
	} else {
		// 否则使用普通的打印方法
		PrintASTStructure(expr, depth)
	}
}
