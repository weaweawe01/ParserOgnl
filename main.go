package main

import (
	"fmt"

	"github.com/weaweawe01/ParserOgnl/lexer"
	"github.com/weaweawe01/ParserOgnl/parser"
)

func main() {
	input := "new ognl.test.objects.Simple(new Object[5])"
	fmt.Println("输入表达式:", input)
	// 创建词法分析器和解析器
	l := lexer.NewLexer(input)
	p := parser.New(l)
	// 解析表达式
	expr, err := p.ParseTopLevelExpression()
	// 检查错误
	if err != nil {
		fmt.Printf("Parser errors: %v\n", err)
		return
	}
	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		return
	}
	if expr == nil {
		fmt.Println("Failed to parse expression")
		return
	}
	// 输出详细的AST结构
	fmt.Println("AST 结构:")
	parser.PrintASTStructure(expr, 0)

}
