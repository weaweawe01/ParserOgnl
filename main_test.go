package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/weaweawe01/ParserOgnl/lexer"
	"github.com/weaweawe01/ParserOgnl/parser"
	"github.com/weaweawe01/ParserOgnl/token"

	"github.com/weaweawe01/ParserOgnl/ast"
)

// 测试表达式的结构体
type ExpressionTest struct {
	input       string
	expectedAST string
	shouldPass  bool
	description string
}

// 辅助函数：打印AST结构用于测试验证
func printASTForTest(expr ast.Expression, depth int) string {
	// 添加 nil 检查
	if expr == nil {
		return ""
	}

	indent := strings.Repeat("  ", depth)
	var result strings.Builder

	switch node := expr.(type) {
	case *ast.BinaryExpression:
		result.WriteString(fmt.Sprintf("%s%s(%s)\n", indent, node.Type(), token.TokenTypeNames[node.Operator]))
		result.WriteString(printASTForTest(node.Left, depth+1))
		result.WriteString(printASTForTest(node.Right, depth+1))
	case *ast.UnaryExpression:
		result.WriteString(fmt.Sprintf("%s%s(%s)\n", indent, node.Type(), token.TokenTypeNames[node.Operator]))
		result.WriteString(printASTForTest(node.Operand, depth+1))
	case *ast.Literal:
		result.WriteString(fmt.Sprintf("%s%s(%v)\n", indent, node.Type(), node.Value))
	case *ast.Identifier:
		result.WriteString(fmt.Sprintf("%s%s(%s)\n", indent, node.Type(), node.Value))
	case *ast.ConditionalExpression:
		result.WriteString(fmt.Sprintf("%s%s\n", indent, node.Type()))
		result.WriteString(printASTForTest(node.Test, depth+1))
		result.WriteString(printASTForTest(node.Consequent, depth+1))
		result.WriteString(printASTForTest(node.Alternative, depth+1))
	case *ast.ChainExpression:
		result.WriteString(fmt.Sprintf("%s%s\n", indent, node.Type()))
		// 使用 Children 数组而不是 Object 和 Property
		for _, child := range node.Children {
			if child != nil {
				result.WriteString(printASTForTest(child, depth+1))
			}
		}
	case *ast.CallExpression:
		result.WriteString(fmt.Sprintf("%s%s(%s)\n", indent, node.Type(), node.Method))
		if node.Object != nil {
			result.WriteString(printASTForTest(node.Object, depth+1))
		}
		for _, arg := range node.Arguments {
			result.WriteString(printASTForTest(arg, depth+1))
		}
	case *ast.StaticMethodExpression:
		result.WriteString(fmt.Sprintf("%s%s(%s@%s)\n", indent, node.Type(), node.ClassName, node.Method))
		for _, arg := range node.Arguments {
			result.WriteString(printASTForTest(arg, depth+1))
		}
	case *ast.StaticFieldExpression:
		result.WriteString(fmt.Sprintf("%s%s(%s@%s)\n", indent, node.Type(), node.ClassName, node.Field))
	case *ast.AssignmentExpression:
		result.WriteString(fmt.Sprintf("%s%s\n", indent, node.Type()))
		result.WriteString(printASTForTest(node.Left, depth+1))
		result.WriteString(printASTForTest(node.Right, depth+1))
	case *ast.SequenceExpression:
		result.WriteString(fmt.Sprintf("%s%s\n", indent, node.Type()))
		for _, expr := range node.Expressions {
			result.WriteString(printASTForTest(expr, depth+1))
		}
	default:
		result.WriteString(fmt.Sprintf("%s%s\n", indent, expr.Type()))
	}

	return result.String()
}

// 辅助函数：执行表达式解析测试
func testExpression(t *testing.T, test ExpressionTest) {
	t.Run(test.description, func(t *testing.T) {
		l := lexer.NewLexer(test.input)
		p := parser.New(l)

		expr, err := p.ParseTopLevelExpression()

		if test.shouldPass {
			// 期望解析成功
			if err != nil {
				t.Errorf("Expected successful parsing, but got error: %v", err)
				return
			}
			if len(p.Errors()) > 0 {
				t.Errorf("Expected no parser errors, but got: %v", p.Errors())
				return
			}
			if expr == nil {
				t.Errorf("Expected non-nil expression, but got nil")
				return
			}

			// 打印成功解析的AST
			fmt.Printf("\n=== %s ===\n", test.description)
			fmt.Printf("Input: %s\n", test.input)
			fmt.Printf("AST Type: %s\n", expr.Type())
			fmt.Printf("AST String: %s\n", expr.String())
			fmt.Printf("AST Structure:\n%s", printASTForTest(expr, 0))

		} else {
			// 期望解析失败
			if err == nil && len(p.Errors()) == 0 && expr != nil {
				t.Errorf("Expected parsing to fail, but it succeeded")
				return
			}
			fmt.Printf("\n=== %s (Expected Failure) ===\n", test.description)
			fmt.Printf("Input: %s\n", test.input)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			if len(p.Errors()) > 0 {
				fmt.Printf("Parser Errors: %v\n", p.Errors())
			}
		}
	})
}

// 测试基本算术运算
func TestArithmeticOperations(t *testing.T) {
	tests := []ExpressionTest{
		// 加法 (10+)
		{"1 + 2", "ASTAdd", true, "整数加法"},
		{"3.14 + 2.86", "ASTAdd", true, "浮点数加法"},
		{"a + b", "ASTAdd", true, "变量加法"},
		{"1 + 2 + 3", "ASTAdd", true, "多项加法"},
		{"x + y + z", "ASTAdd", true, "三变量加法"},
		{"100 + 200", "ASTAdd", true, "大数加法"},
		{"0 + 1", "ASTAdd", true, "零加法"},
		{"-5 + 10", "ASTAdd", true, "负数加法"},
		{"a + 1", "ASTAdd", true, "变量加常量"},
		{"1.5 + 2.5", "ASTAdd", true, "小数加法"},
		{"0.1 + 0.2", "ASTAdd", true, "精度测试加法"},

		// 减法 (10+)
		{"10 - 5", "ASTSubtract", true, "整数减法"},
		{"x - y", "ASTSubtract", true, "变量减法"},
		{"100 - 50 - 25", "ASTSubtract", true, "连续减法"},
		{"0 - 5", "ASTSubtract", true, "零减法"},
		{"a - 10", "ASTSubtract", true, "变量减常量"},
		{"5.5 - 2.3", "ASTSubtract", true, "小数减法"},
		{"-10 - 5", "ASTSubtract", true, "负数减法"},
		{"x - y - z", "ASTSubtract", true, "三项减法"},
		{"1000 - 1", "ASTSubtract", true, "大数减法"},
		{"a - b - c", "ASTSubtract", true, "多变量减法"},

		// 乘法 (10+)
		{"3 * 4", "ASTMultiply", true, "整数乘法"},
		{"a * b * c", "ASTMultiply", true, "多项乘法"},
		{"2 * 3 * 4", "ASTMultiply", true, "三数乘法"},
		{"0 * 100", "ASTMultiply", true, "零乘法"},
		{"1 * x", "ASTMultiply", true, "单位乘法"},
		{"2.5 * 4", "ASTMultiply", true, "小数乘整数"},
		{"1.5 * 2.0", "ASTMultiply", true, "小数乘法"},
		{"-3 * 5", "ASTMultiply", true, "负数乘法"},
		{"x * y * z * w", "ASTMultiply", true, "四变量乘法"},
		{"10 * 10", "ASTMultiply", true, "平方"},

		// 除法 (10+)
		{"10 / 2", "ASTDivide", true, "整数除法"},
		{"x / y", "ASTDivide", true, "变量除法"},
		{"100 / 10 / 2", "ASTDivide", true, "连续除法"},
		{"a / 2", "ASTDivide", true, "变量除常量"},
		{"5.0 / 2.0", "ASTDivide", true, "小数除法"},
		{"1 / 2", "ASTDivide", true, "分数除法"},
		{"0 / 1", "ASTDivide", true, "零除法"},
		{"x / 1", "ASTDivide", true, "除以1"},
		{"100 / 3", "ASTDivide", true, "不整除"},
		{"a / b / c", "ASTDivide", true, "多项除法"},

		// 求余 (5+)
		{"10 % 3", "ASTRemainder", true, "求余运算"},
		{"x % y", "ASTRemainder", true, "变量求余"},
		{"100 % 7", "ASTRemainder", true, "大数求余"},
		{"5 % 2", "ASTRemainder", true, "奇数求余"},
		{"a % 10", "ASTRemainder", true, "取个位"},

		// 复杂算术表达式 (10+)
		{"1 + 2 * 3", "ASTAdd", true, "运算符优先级"},
		{"(1 + 2) * 3", "ASTMultiply", true, "括号改变优先级"},
		{"-5 + 3", "ASTAdd", true, "负数参与运算"},
		{"2 * 3 + 4 * 5", "ASTAdd", true, "多重乘加"},
		{"10 - 2 * 3", "ASTSubtract", true, "减乘混合"},
		{"(a + b) * (c + d)", "ASTMultiply", true, "复杂括号表达式"},
		{"1 + 2 - 3 * 4 / 5", "ASTAdd", true, "四则混合运算"},
		{"a * b + c * d", "ASTAdd", true, "分配律形式"},
		{"(x + y) / (a - b)", "ASTDivide", true, "分式"},
		{"2 * (3 + 4) - 5", "ASTSubtract", true, "括号优先级"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试比较运算
func TestComparisonOperations(t *testing.T) {
	tests := []ExpressionTest{
		// 相等比较 (10+)
		{"1 == 2", "ASTEq", true, "数字相等比较"},
		{"a == b", "ASTEq", true, "变量相等比较"},
		{"\"hello\" == \"world\"", "ASTEq", true, "字符串相等比较"},
		{"true == false", "ASTEq", true, "布尔相等比较"},
		{"x == 0", "ASTEq", true, "变量等于零"},
		{"1 == 1", "ASTEq", true, "相同数字比较"},
		{"a == a", "ASTEq", true, "自身比较"},
		{"null == null", "ASTEq", true, "null比较"},
		{"5.5 == 5.5", "ASTEq", true, "小数相等"},
		{"'a' == 'b'", "ASTEq", true, "字符相等比较"},

		// 不等比较 (10+)
		{"1 != 2", "ASTNotEq", true, "不等比较"},
		{"x != y", "ASTNotEq", true, "变量不等比较"},
		{"a != 0", "ASTNotEq", true, "非零检查"},
		{"true != false", "ASTNotEq", true, "布尔不等"},
		{"\"a\" != \"b\"", "ASTNotEq", true, "字符串不等"},
		{"null != x", "ASTNotEq", true, "null不等"},
		{"1 != 1", "ASTNotEq", true, "相同值不等"},
		{"x != x", "ASTNotEq", true, "自身不等"},
		{"5 != 10", "ASTNotEq", true, "整数不等"},
		{"1.0 != 2.0", "ASTNotEq", true, "小数不等"},

		// 大于比较 (10+)
		{"5 > 3", "ASTGreater", true, "大于比较"},
		{"x > y", "ASTGreater", true, "变量大于"},
		{"10 > 0", "ASTGreater", true, "大于零"},
		{"a > 100", "ASTGreater", true, "变量大于常量"},
		{"0 > -1", "ASTGreater", true, "零大于负数"},
		{"1.5 > 1.0", "ASTGreater", true, "小数大于"},
		{"100 > 99", "ASTGreater", true, "相近数比较"},
		{"x > 0", "ASTGreater", true, "正数检查"},
		{"a > b", "ASTGreater", true, "两变量比较"},
		{"5 > 5", "ASTGreater", true, "等值大于"},

		// 大于等于比较 (10+)
		{"x >= y", "ASTGreaterEq", true, "大于等于比较"},
		{"5 >= 5", "ASTGreaterEq", true, "等于情况"},
		{"6 >= 5", "ASTGreaterEq", true, "大于情况"},
		{"a >= 0", "ASTGreaterEq", true, "非负检查"},
		{"x >= 100", "ASTGreaterEq", true, "大于等于常量"},
		{"0 >= -5", "ASTGreaterEq", true, "零大于等于负数"},
		{"1.0 >= 0.9", "ASTGreaterEq", true, "小数大于等于"},
		{"a >= a", "ASTGreaterEq", true, "自身大于等于"},
		{"10 >= 5", "ASTGreaterEq", true, "明显大于"},
		{"x >= y", "ASTGreaterEq", true, "变量大于等于"},

		// 小于比较 (10+)
		{"2 < 5", "ASTLess", true, "小于比较"},
		{"x < y", "ASTLess", true, "变量小于"},
		{"0 < 10", "ASTLess", true, "零小于正数"},
		{"-1 < 0", "ASTLess", true, "负数小于零"},
		{"a < 100", "ASTLess", true, "变量小于常量"},
		{"1.0 < 1.5", "ASTLess", true, "小数小于"},
		{"5 < 5", "ASTLess", true, "等值小于"},
		{"x < 0", "ASTLess", true, "负数检查"},
		{"1 < 1000", "ASTLess", true, "大差距比较"},
		{"a < b", "ASTLess", true, "变量小于"},

		// 小于等于比较 (10+)
		{"a <= b", "ASTLessEq", true, "小于等于比较"},
		{"5 <= 5", "ASTLessEq", true, "等于情况"},
		{"3 <= 5", "ASTLessEq", true, "小于情况"},
		{"x <= 0", "ASTLessEq", true, "非正检查"},
		{"0 <= a", "ASTLessEq", true, "零小于等于"},
		{"-5 <= 0", "ASTLessEq", true, "负数小于等于零"},
		{"1.0 <= 1.1", "ASTLessEq", true, "小数小于等于"},
		{"a <= a", "ASTLessEq", true, "自身小于等于"},
		{"1 <= 100", "ASTLessEq", true, "明显小于"},
		{"x <= y", "ASTLessEq", true, "变量小于等于"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试逻辑运算
func TestLogicalOperations(t *testing.T) {
	tests := []ExpressionTest{
		// 与运算 (10+)
		{"true && true", "ASTAnd", true, "两个true的与"},
		{"true && false", "ASTAnd", true, "true与false"},
		{"false && false", "ASTAnd", true, "两个false的与"},
		{"a && b", "ASTAnd", true, "变量与运算"},
		{"x > 0 && y > 0", "ASTAnd", true, "条件与运算"},
		{"true && x", "ASTAnd", true, "true与变量"},
		{"a == b && c == d", "ASTAnd", true, "等式与运算"},
		{"(x > 5) && (y < 10)", "ASTAnd", true, "括号与运算"},
		{"flag && ready", "ASTAnd", true, "布尔变量与"},
		{"1 == 1 && 2 == 2", "ASTAnd", true, "多个等式与"},
		{"a && b && c", "ASTAnd", true, "三个变量与"},
		{"true && true && true", "ASTAnd", true, "多个true"},

		// 或运算 (10+)
		{"true || false", "ASTOr", true, "true或false"},
		{"false || false", "ASTOr", true, "两个false的或"},
		{"true || true", "ASTOr", true, "两个true的或"},
		{"a || b", "ASTOr", true, "变量或运算"},
		{"x > 0 || y > 0", "ASTOr", true, "条件或运算"},
		{"false || x", "ASTOr", true, "false或变量"},
		{"a == b || c == d", "ASTOr", true, "等式或运算"},
		{"(x < 0) || (x > 100)", "ASTOr", true, "范围或运算"},
		{"flag || ready", "ASTOr", true, "布尔变量或"},
		{"1 != 1 || 2 == 2", "ASTOr", true, "混合条件或"},
		{"a || b || c", "ASTOr", true, "三个变量或"},
		{"false || false || true", "ASTOr", true, "多个或运算"},

		// 非运算 (10+)
		{"!true", "ASTNot", true, "true取反"},
		{"!false", "ASTNot", true, "false取反"},
		{"!x", "ASTNot", true, "变量取反"},
		{"!(a > b)", "ASTNot", true, "条件取反"},
		{"!flag", "ASTNot", true, "布尔变量取反"},
		{"!(x == y)", "ASTNot", true, "等式取反"},
		{"!!x", "ASTNot", true, "双重取反"},
		{"!(a && b)", "ASTNot", true, "与运算取反"},
		{"!(a || b)", "ASTNot", true, "或运算取反"},
		{"!ready", "ASTNot", true, "就绪标志取反"},
		{"!(x > 0 && y > 0)", "ASTNot", true, "复杂条件取反"},
		{"!(true)", "ASTNot", true, "括号true取反"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试位运算
func TestBitwiseOperations(t *testing.T) {
	tests := []ExpressionTest{
		// 位与 (10+)
		{"5 & 3", "ASTBitAnd", true, "位与运算"},
		{"a & b", "ASTBitAnd", true, "变量位与"},
		{"0xFF & 0x0F", "ASTBitAnd", true, "十六进制位与"},
		{"255 & 15", "ASTBitAnd", true, "掩码运算"},
		{"x & 1", "ASTBitAnd", true, "奇偶检查"},
		{"a & 0", "ASTBitAnd", true, "清零"},
		{"b & 0xFF", "ASTBitAnd", true, "取低8位"},
		{"(a & b) & c", "ASTBitAnd", true, "多重位与"},
		{"x & y & z", "ASTBitAnd", true, "三个位与"},
		{"128 & 64", "ASTBitAnd", true, "2的幂位与"},
		{"flag & mask", "ASTBitAnd", true, "标志掩码"},

		// 位或 (10+)
		{"5 | 3", "ASTBitOr", true, "位或运算"},
		{"x | y", "ASTBitOr", true, "变量位或"},
		{"0x10 | 0x01", "ASTBitOr", true, "十六进制位或"},
		{"a | 1", "ASTBitOr", true, "设置最低位"},
		{"x | 0", "ASTBitOr", true, "保持原值"},
		{"flag | mask", "ASTBitOr", true, "设置标志位"},
		{"a | b | c", "ASTBitOr", true, "多个位或"},
		{"128 | 64 | 32", "ASTBitOr", true, "组合标志"},
		{"r | g | b", "ASTBitOr", true, "RGB组合"},
		{"(x | y) | z", "ASTBitOr", true, "括号位或"},

		// 位异或 (10+)
		{"5 ^ 3", "ASTXor", true, "位异或运算"},
		{"a ^ b", "ASTXor", true, "变量位异或"},
		{"x ^ 0xFF", "ASTXor", true, "反转所有位"},
		{"a ^ 0", "ASTXor", true, "保持不变"},
		{"x ^ x", "ASTXor", true, "清零运算"},
		{"a ^ b ^ a", "ASTXor", true, "交换运算"},
		{"flag ^ 1", "ASTXor", true, "翻转最低位"},
		{"x ^ y ^ z", "ASTXor", true, "多重异或"},
		{"(a ^ b) ^ c", "ASTXor", true, "括号异或"},
		{"mask ^ 0xFF", "ASTXor", true, "反转掩码"},

		// 位取反 (5+)
		{"~5", "ASTBitNegate", true, "位取反运算"},
		{"~flag", "ASTBitNegate", true, "变量位取反"},
		{"~0", "ASTBitNegate", true, "零取反"},
		{"~mask", "ASTBitNegate", true, "掩码取反"},
		{"~~x", "ASTBitNegate", true, "双重取反"},

		// 左移运算 (10+)
		{"8 << 2", "ASTShiftLeft", true, "左移运算"},
		{"1 << 0", "ASTShiftLeft", true, "左移0位"},
		{"1 << 8", "ASTShiftLeft", true, "左移8位"},
		{"x << 1", "ASTShiftLeft", true, "乘以2"},
		{"a << n", "ASTShiftLeft", true, "变量左移"},
		{"1 << 16", "ASTShiftLeft", true, "大幅左移"},
		{"value << shift", "ASTShiftLeft", true, "值左移"},
		{"2 << 3", "ASTShiftLeft", true, "2的幂左移"},
		{"mask << pos", "ASTShiftLeft", true, "掩码定位"},
		{"(x << 2) << 1", "ASTShiftLeft", true, "连续左移"},

		// 右移运算 (10+)
		{"16 >> 2", "ASTShiftRight", true, "右移运算"},
		{"x >> 1", "ASTShiftRight", true, "除以2"},
		{"a >> n", "ASTShiftRight", true, "变量右移"},
		{"128 >> 3", "ASTShiftRight", true, "大数右移"},
		{"value >> shift", "ASTShiftRight", true, "值右移"},
		{"32 >> 5", "ASTShiftRight", true, "完全右移"},
		{"-8 >> 2", "ASTShiftRight", true, "负数右移"},
		{"(x >> 2) >> 1", "ASTShiftRight", true, "连续右移"},
		{"mask >> pos", "ASTShiftRight", true, "掩码右移"},
		{"x >> 0", "ASTShiftRight", true, "右移0位"},

		// 无符号右移 (5+)
		{"16 >>> 2", "ASTUnsignedShiftRight", true, "无符号右移"},
		{"-1 >>> 1", "ASTUnsignedShiftRight", true, "负数无符号右移"},
		{"x >>> n", "ASTUnsignedShiftRight", true, "变量无符号右移"},
		{"value >>> shift", "ASTUnsignedShiftRight", true, "值无符号右移"},
		{"0x80000000 >>> 1", "ASTUnsignedShiftRight", true, "符号位右移"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试一元运算
func TestUnaryOperations(t *testing.T) {
	tests := []ExpressionTest{
		// 负号
		{"-5", "ASTNegate", true, "负数"},
		{"-x", "ASTNegate", true, "变量取负"},
		{"-(a + b)", "ASTNegate", true, "表达式取负"},

		// 正号
		{"+5", "ASTConst", true, "正数"},
		{"+x", "ASTConst", true, "变量取正"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试条件表达式
func TestConditionalExpressions(t *testing.T) {
	tests := []ExpressionTest{
		// 基本三元条件 (10+)
		{"true ? 1 : 2", "ASTTest", true, "基本三元条件"},
		{"false ? 1 : 2", "ASTTest", true, "false条件"},
		{"x > 0 ? x : -x", "ASTTest", true, "绝对值运算"},
		{"flag ? \"yes\" : \"no\"", "ASTTest", true, "字符串三元条件"},
		{"a == b ? 1 : 0", "ASTTest", true, "相等检查"},
		{"x ? a : b", "ASTTest", true, "变量条件"},
		{"(x > y) ? x : y", "ASTTest", true, "最大值"},
		{"(x < y) ? x : y", "ASTTest", true, "最小值"},
		{"n % 2 == 0 ? \"even\" : \"odd\"", "ASTTest", true, "奇偶判断"},
		{"score >= 60 ? \"pass\" : \"fail\"", "ASTTest", true, "及格判断"},
		{"x != null ? x : 0", "ASTTest", true, "空值检查"},

		// 嵌套条件 (10+)
		{"a > b ? (a > c ? a : c) : (b > c ? b : c)", "ASTTest", true, "三值最大"},
		{"x > 0 ? 1 : (x < 0 ? -1 : 0)", "ASTTest", true, "符号函数"},
		{"a ? (b ? 1 : 2) : 3", "ASTTest", true, "嵌套true分支"},
		{"a ? 1 : (b ? 2 : 3)", "ASTTest", true, "嵌套false分支"},
		{"(a ? b : c) ? x : y", "ASTTest", true, "条件作条件"},
		{"x > 10 ? (x > 20 ? \"high\" : \"mid\") : \"low\"", "ASTTest", true, "范围判断"},
		{"a && b ? 1 : (c || d ? 2 : 3)", "ASTTest", true, "逻辑条件组合"},
		{"(x ? y : z) + 1", "ASTAdd", true, "条件结果运算"},
		{"a ? b ? c : d : e", "ASTTest", true, "连续嵌套"},
		{"x > 0 ? x * 2 : x / 2", "ASTTest", true, "不同运算"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试赋值表达式
func TestAssignmentExpressions(t *testing.T) {
	tests := []ExpressionTest{
		// 简单赋值
		{"x = 5", "ASTAssign", true, "变量赋值"},
		{"name = \"hello\"", "ASTAssign", true, "字符串赋值"},

		// 链式赋值
		{"a = b = c", "ASTAssign", true, "链式赋值"},
		{"x = y = 10", "ASTAssign", true, "多变量赋值"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试序列表达式
func TestSequenceExpressions(t *testing.T) {
	tests := []ExpressionTest{
		// 逗号分隔的表达式
		{"1, 2, 3", "ASTSequence", true, "数字序列"},
		{"a = 1, b = 2", "ASTSequence", true, "赋值序列"},
		{"--y", "ASTSequence", true, "操作序列"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试字面量
func TestLiterals(t *testing.T) {
	tests := []ExpressionTest{
		// 整数字面量 (10+)
		{"42", "ASTConst", true, "整数字面量"},
		{"0", "ASTConst", true, "零"},
		{"-100", "ASTNegate", true, "负整数"},
		{"2147483647", "ASTConst", true, "最大int"},
		{"1000000", "ASTConst", true, "百万"},
		{"999", "ASTConst", true, "三位数"},
		{"1", "ASTConst", true, "最小正整数"},
		{"128", "ASTConst", true, "2的幂"},
		{"255", "ASTConst", true, "字节最大值"},
		{"65535", "ASTConst", true, "short最大值"},

		// 浮点数字面量 (10+)
		{"3.14", "ASTConst", true, "浮点数字面量"},
		{"0.0", "ASTConst", true, "零浮点"},
		{"1.0", "ASTConst", true, "整数浮点"},
		{"0.5", "ASTConst", true, "小数"},
		{"3.14159", "ASTConst", true, "多位小数"},
		{"1.23e10", "ASTConst", true, "科学计数法"},
		{"1.5e-3", "ASTConst", true, "负指数"},
		{"-3.14", "ASTNegate", true, "负浮点"},
		{"0.1", "ASTConst", true, "十分之一"},
		{"99.99", "ASTConst", true, "价格格式"},

		// 十六进制字面量 (5+)
		{"0xFF", "ASTConst", true, "十六进制字面量"},
		{"0x00", "ASTConst", true, "十六进制零"},
		{"0x10", "ASTConst", true, "十六进制16"},
		{"0xDEADBEEF", "ASTConst", true, "大十六进制"},
		{"0xABCD", "ASTConst", true, "混合大小写十六进制"},

		// 八进制字面量 (5+)
		{"077", "ASTConst", true, "八进制字面量"},
		{"00", "ASTConst", true, "八进制零"},
		{"010", "ASTConst", true, "八进制8"},
		{"0777", "ASTConst", true, "大八进制"},
		{"0123", "ASTConst", true, "混合八进制"},

		// 字符串字面量 (10+)
		{"\"hello\"", "ASTConst", true, "字符串字面量"},
		{"\"\"", "ASTConst", true, "空字符串"},
		{"\"world\"", "ASTConst", true, "单词字符串"},
		{"\"Hello, World!\"", "ASTConst", true, "带标点字符串"},
		{"\"123\"", "ASTConst", true, "数字字符串"},
		{"\"true\"", "ASTConst", true, "布尔字符串"},
		{"\"  spaces  \"", "ASTConst", true, "带空格字符串"},
		{"\"line1\\nline2\"", "ASTConst", true, "转义字符"},
		{"\"tab\\there\"", "ASTConst", true, "制表符"},
		{"\"quote\\\"here\"", "ASTConst", true, "引号转义"},

		// 字符字面量 (5+)
		{"'a'", "ASTConst", true, "字符字面量"},
		{"'Z'", "ASTConst", true, "大写字符"},
		{"'0'", "ASTConst", true, "数字字符"},
		{"' '", "ASTConst", true, "空格字符"},
		{"'\\n'", "ASTConst", true, "换行字符"},

		// 布尔字面量 (2+)
		{"true", "ASTConst", true, "true字面量"},
		{"false", "ASTConst", true, "false字面量"},

		// null字面量 (1+)
		{"null", "ASTConst", true, "null字面量"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试标识符和变量引用
func TestIdentifiersAndVariables(t *testing.T) {
	tests := []ExpressionTest{
		// 简单标识符
		{"name", "ASTProperty", true, "简单标识符"},
		{"_variable", "ASTProperty", true, "下划线开头标识符"},
		{"var123", "ASTProperty", true, "数字结尾标识符"},

		// 特殊变量引用
		{"#this", "ASTThisVarRef", true, "this引用"},
		{"#root", "ASTRootVarRef", true, "root引用"},
		{"#variable", "ASTVarRef", true, "变量引用"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试静态引用
func TestStaticReferences(t *testing.T) {
	tests := []ExpressionTest{
		// 静态字段 (10+)
		{"@Math@PI", "ASTStaticField", true, "静态字段引用"},
		{"@System@out", "ASTStaticField", true, "复杂静态字段"},
		{"@Integer@MAX_VALUE", "ASTStaticField", true, "最大值常量"},
		{"@Integer@MIN_VALUE", "ASTStaticField", true, "最小值常量"},
		{"@Double@NaN", "ASTStaticField", true, "NaN常量"},
		{"@Boolean@TRUE", "ASTStaticField", true, "布尔常量"},
		{"@Color@RED", "ASTStaticField", true, "颜色常量"},
		{"@Calendar@JANUARY", "ASTStaticField", true, "月份常量"},
		{"@File@separator", "ASTStaticField", true, "分隔符"},
		{"@Character@MAX_VALUE", "ASTStaticField", true, "字符最大值"},

		// 静态方法 (10+)
		{"@Math@max(1, 2)", "ASTStaticMethod", true, "静态方法调用"},
		{"@System@currentTimeMillis()", "ASTStaticMethod", true, "无参静态方法"},
		{"@String@valueOf(123)", "ASTStaticMethod", true, "带参静态方法"},
		{"@Math@abs(-5)", "ASTStaticMethod", true, "绝对值方法"},
		{"@Math@sqrt(16)", "ASTStaticMethod", true, "平方根方法"},
		{"@Integer@parseInt(\"123\")", "ASTStaticMethod", true, "解析方法"},
		{"@String@format(\"%d\", 42)", "ASTStaticMethod", true, "格式化方法"},
		{"@Arrays@asList(1, 2, 3)", "ASTStaticMethod", true, "数组转列表"},
		{"@Collections@emptyList()", "ASTStaticMethod", true, "空列表"},
		{"@Objects@requireNonNull(obj)", "ASTStaticMethod", true, "非空检查"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试属性访问
func TestPropertyAccess(t *testing.T) {
	tests := []ExpressionTest{
		// 简单属性访问 (10+)
		{"user.name", "ASTChain", true, "简单属性访问"},
		{"person.age", "ASTChain", true, "年龄属性"},
		{"obj.value", "ASTChain", true, "值属性"},
		{"account.balance", "ASTChain", true, "余额属性"},
		{"employee.salary", "ASTChain", true, "薪水属性"},
		{"product.price", "ASTChain", true, "价格属性"},
		{"book.title", "ASTChain", true, "标题属性"},
		{"car.brand", "ASTChain", true, "品牌属性"},
		{"student.score", "ASTChain", true, "分数属性"},
		{"order.status", "ASTChain", true, "状态属性"},

		// 链式属性访问 (10+)
		{"user.address.city", "ASTChain", true, "链式属性访问"},
		{"person.company.name", "ASTChain", true, "公司名称"},
		{"a.b.c", "ASTChain", true, "多层属性"},
		{"obj.field.value", "ASTChain", true, "字段值"},
		{"root.child.grandchild", "ASTChain", true, "祖孙关系"},
		{"order.customer.address", "ASTChain", true, "订单地址"},
		{"book.author.name", "ASTChain", true, "作者名"},
		{"x.y.z.w", "ASTChain", true, "四层属性"},
		{"employee.department.manager", "ASTChain", true, "部门经理"},
		{"student.class.teacher", "ASTChain", true, "班级老师"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试方法调用
func TestMethodCalls(t *testing.T) {
	tests := []ExpressionTest{
		// 无参方法调用 (10+)
		{"obj.toString()", "ASTChain", true, "toString方法"},
		{"list.size()", "ASTChain", true, "size方法"},
		{"str.length()", "ASTChain", true, "length方法"},
		{"user.getName()", "ASTChain", true, "getName方法"},
		{"obj.hashCode()", "ASTChain", true, "hashCode方法"},
		{"array.clone()", "ASTChain", true, "clone方法"},
		{"date.getTime()", "ASTChain", true, "getTime方法"},
		{"collection.isEmpty()", "ASTChain", true, "isEmpty方法"},
		{"map.clear()", "ASTChain", true, "clear方法"},
		{"buffer.flush()", "ASTChain", true, "flush方法"},

		// 带参方法调用 (10+)
		{"str.charAt(0)", "ASTChain", true, "charAt方法"},
		{"list.get(0)", "ASTChain", true, "get方法"},
		{"str.substring(1, 5)", "ASTChain", true, "substring方法"},
		{"map.put(\"key\", value)", "ASTChain", true, "put方法"},
		{"list.add(item)", "ASTChain", true, "add方法"},
		{"str.equals(\"test\")", "ASTChain", true, "equals方法"},
		{"num.compareTo(other)", "ASTChain", true, "compareTo方法"},
		{"str.replace(\"old\", \"new\")", "ASTChain", true, "replace方法"},
		{"list.contains(item)", "ASTChain", true, "contains方法"},
		{"str.indexOf(\"x\")", "ASTChain", true, "indexOf方法"},

		// 链式方法调用 (10+)
		{"str.trim().toLowerCase()", "ASTChain", true, "链式方法调用"},
		{"list.subList(0, 5).size()", "ASTChain", true, "子列表大小"},
		{"user.getName().toUpperCase()", "ASTChain", true, "大写名称"},
		{"str.replace(\"a\", \"b\").length()", "ASTChain", true, "替换后长度"},
		{"obj.toString().charAt(0)", "ASTChain", true, "首字符"},
		{"a.b().c().d()", "ASTChain", true, "多重链式"},
		{"list.get(0).toString()", "ASTChain", true, "元素转字符串"},
		{"map.get(key).process()", "ASTChain", true, "值处理"},
		{"date.getTime().toString()", "ASTChain", true, "时间转字符串"},
		{"str.trim().substring(1)", "ASTChain", true, "修剪后截取"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试数组访问
func TestArrayAccess(t *testing.T) {
	tests := []ExpressionTest{
		// 基本数组访问 (10+)
		{"array[0]", "ASTChain", true, "数组索引访问"},
		{"arr[1]", "ASTChain", true, "数组第二元素"},
		{"list[0]", "ASTChain", true, "列表索引"},
		{"items[i]", "ASTChain", true, "变量索引"},
		{"data[10]", "ASTChain", true, "固定索引"},
		{"array[array.length - 1]", "ASTChain", true, "最后元素"},
		{"matrix[0][0]", "ASTChain", true, "二维数组"},
		{"arr[x + 1]", "ASTChain", true, "表达式索引"},
		{"list[i * 2]", "ASTChain", true, "计算索引"},
		{"data[n]", "ASTChain", true, "变量n索引"},

		// 链式数组访问 (10+)
		{"users[0].name", "ASTChain", true, "数组元素属性"},
		{"list[0].toString()", "ASTChain", true, "数组元素方法"},
		{"arr[i].value", "ASTChain", true, "索引元素值"},
		{"matrix[0][1]", "ASTChain", true, "矩阵元素"},
		{"data[x].process()", "ASTChain", true, "元素处理"},
		{"items[0].field.value", "ASTChain", true, "深层访问"},
		{"array[0][1][2]", "ASTChain", true, "三维数组"},
		{"list[n].get(0)", "ASTChain", true, "嵌套列表"},
		{"map[key].size()", "ASTChain", true, "映射值大小"},
		{"users[i].address.city", "ASTChain", true, "用户城市"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试集合字面量
func TestCollectionLiterals(t *testing.T) {
	tests := []ExpressionTest{
		// 列表字面量 (10+)
		{"{1, 2, 3}", "ASTList", true, "基本列表"},
		{"{}", "ASTList", true, "空列表"},
		{"{a, b, c}", "ASTList", true, "变量列表"},
		{"{\"one\", \"two\", \"three\"}", "ASTList", true, "字符串列表"},
		{"{1, 2, 3, 4, 5}", "ASTList", true, "数字列表"},
		{"{true, false, true}", "ASTList", true, "布尔列表"},
		{"{x, y, z}", "ASTList", true, "xyz列表"},
		{"{1 + 1, 2 * 2, 3 - 1}", "ASTList", true, "表达式列表"},
		{"{obj1, obj2, obj3}", "ASTList", true, "对象列表"},
		{"{a.b, c.d, e.f}", "ASTList", true, "属性列表"},

		// 映射字面量 (10+)
		{"#{\"key\": \"value\"}", "ASTMap", true, "基本映射"},
		{"#{}", "ASTMap", true, "空映射"},
		{"#{\"a\": 1, \"b\": 2}", "ASTMap", true, "多键值对"},
		{"#{1: \"one\", 2: \"two\"}", "ASTMap", true, "数字键"},
		{"#{x: a, y: b}", "ASTMap", true, "变量映射"},
		{"#{\"name\": name, \"age\": age}", "ASTMap", true, "混合映射"},
		{"#{true: 1, false: 0}", "ASTMap", true, "布尔键"},
		{"#{\"k1\": v1, \"k2\": v2, \"k3\": v3}", "ASTMap", true, "三键值对"},
		{"#{a: 1, b: 2, c: 3}", "ASTMap", true, "简单键映射"},
		{"#{\"x\": x + 1, \"y\": y * 2}", "ASTMap", true, "表达式值"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 测试构造器调用
func TestConstructorCalls(t *testing.T) {
	tests := []ExpressionTest{
		// 无参构造器 (5+)
		{"new String()", "ASTCtor", true, "无参构造器"},
		{"new Object()", "ASTCtor", true, "Object构造器"},
		{"new ArrayList()", "ASTCtor", true, "ArrayList构造器"},
		{"new HashMap()", "ASTCtor", true, "HashMap构造器"},
		{"new Date()", "ASTCtor", true, "Date构造器"},

		// 带参构造器 (10+)
		{"new String(\"hello\")", "ASTCtor", true, "带参构造器"},
		{"new Integer(42)", "ASTCtor", true, "Integer构造器"},
		{"new ArrayList(10)", "ASTCtor", true, "指定容量"},
		{"new Date(0)", "ASTCtor", true, "时间戳构造"},
		{"new Point(x, y)", "ASTCtor", true, "坐标构造"},
		{"new Person(name, age)", "ASTCtor", true, "Person构造"},
		{"new Rectangle(0, 0, 100, 100)", "ASTCtor", true, "矩形构造"},
		{"new String(bytes, charset)", "ASTCtor", true, "多参构造"},
		{"new BigDecimal(\"123.45\")", "ASTCtor", true, "BigDecimal构造"},
		{"new File(path)", "ASTCtor", true, "File构造"},
	}

	for _, test := range tests {
		testExpression(t, test)
	}
}

// 运行所有基本类型测试
func TestAllBasicTypes(t *testing.T) {
	fmt.Println("开始运行OGNL基本类型解析测试...")
	fmt.Println("==========================================")

	// 运行各个测试组
	t.Run("算术运算", TestArithmeticOperations)
	t.Run("比较运算", TestComparisonOperations)
	t.Run("逻辑运算", TestLogicalOperations)
	t.Run("位运算", TestBitwiseOperations)
	t.Run("一元运算", TestUnaryOperations)
	t.Run("条件表达式", TestConditionalExpressions)
	t.Run("赋值表达式", TestAssignmentExpressions)
	t.Run("序列表达式", TestSequenceExpressions)
	t.Run("字面量", TestLiterals)
	t.Run("标识符和变量", TestIdentifiersAndVariables)
	t.Run("静态引用", TestStaticReferences)
	t.Run("属性访问", TestPropertyAccess)
	t.Run("方法调用", TestMethodCalls)
	t.Run("数组访问", TestArrayAccess)
	t.Run("集合字面量", TestCollectionLiterals)
	t.Run("构造器调用", TestConstructorCalls)

	fmt.Println("\n==========================================")
	fmt.Println("OGNL基本类型解析测试完成!")
}
