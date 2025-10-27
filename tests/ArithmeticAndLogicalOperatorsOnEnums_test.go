package test

import (
	"testing"

	"github.com/weaweawe01/ParserOgnl/ast"

	"github.com/weaweawe01/ParserOgnl/lexer"
	"github.com/weaweawe01/ParserOgnl/parser"
)

// 模拟枚举类型 (Go 中使用 iota 实现)
type EnumNoBody int

const (
	ENUM1_NoBody EnumNoBody = iota
	ENUM2_NoBody
)

type EnumEmptyBody int

const (
	ENUM1_EmptyBody EnumEmptyBody = iota
	ENUM2_EmptyBody
)

type EnumBasicBody int

const (
	ENUM1_BasicBody EnumBasicBody = iota
	ENUM2_BasicBody
)

func (e EnumBasicBody) Value() int {
	switch e {
	case ENUM1_BasicBody:
		return 10
	case ENUM2_BasicBody:
		return 20
	default:
		return 0
	}
}

// 辅助函数: 解析 OGNL 表达式
func parseExpression(t *testing.T, input string) ast.Expression {
	l := lexer.NewLexer(input)
	p := parser.New(l)
	expr, err := p.ParseTopLevelExpression()

	if err != nil {
		t.Fatalf("解析错误: %v", err)
	}

	if len(p.Errors()) > 0 {
		t.Fatalf("解析器错误: %v", p.Errors())
	}

	if expr == nil {
		t.Fatal("表达式解析结果为 nil")
	}

	return expr
}

// 辅助函数: 创建静态字段访问的期望节点
func createStaticFieldExpected(className, field string) ExpectedNode {
	return ExpectedNode{
		Type:     "ASTStaticField",
		Fragment: "@" + className + "@" + field,
	}
}

// 辅助函数: 创建二元表达式的期望节点
func createBinaryExpected(exprType, fragment string, leftType, leftFragment, rightType, rightFragment string) ExpectedNode {
	return ExpectedNode{
		Type:     exprType,
		Fragment: fragment,
		Children: []ExpectedNode{
			{Type: leftType, Fragment: leftFragment},
			{Type: rightType, Fragment: rightFragment},
		},
	}
}

// ============================================================================
// EnumNoBody 测试用例
// ============================================================================

func TestEnumNoBodyEquality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 == @EnumNoBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumNoBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyInequality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 != @EnumNoBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumNoBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyEqualityEnum2(t *testing.T) {
	input := "@EnumNoBody@ENUM2 == @EnumNoBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM2",
		"ASTStaticField", "@EnumNoBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyInequalityEnum2(t *testing.T) {
	input := "@EnumNoBody@ENUM2 != @EnumNoBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM2",
		"ASTStaticField", "@EnumNoBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyDifferentEnums(t *testing.T) {
	input := "@EnumNoBody@ENUM1 != @EnumNoBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumNoBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyDifferentEnumsEquality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 == @EnumNoBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumNoBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

// ============================================================================
// EnumEmptyBody 测试用例
// ============================================================================

func TestEnumEmptyBodyEquality(t *testing.T) {
	input := "@EnumEmptyBody@ENUM1 == @EnumEmptyBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyInequality(t *testing.T) {
	input := "@EnumEmptyBody@ENUM1 != @EnumEmptyBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyEqualityEnum2(t *testing.T) {
	input := "@EnumEmptyBody@ENUM2 == @EnumEmptyBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM2",
		"ASTStaticField", "@EnumEmptyBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyInequalityEnum2(t *testing.T) {
	input := "@EnumEmptyBody@ENUM2 != @EnumEmptyBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM2",
		"ASTStaticField", "@EnumEmptyBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyDifferentEnums(t *testing.T) {
	input := "@EnumEmptyBody@ENUM1 != @EnumEmptyBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
		"ASTStaticField", "@EnumEmptyBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyDifferentEnumsEquality(t *testing.T) {
	input := "@EnumEmptyBody@ENUM1 == @EnumEmptyBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
		"ASTStaticField", "@EnumEmptyBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

// ============================================================================
// EnumBasicBody 测试用例
// ============================================================================

func TestEnumBasicBodyEquality(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 == @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumBasicBodyInequality(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 != @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumBasicBodyEqualityEnum2(t *testing.T) {
	input := "@EnumBasicBody@ENUM2 == @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM2",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumBasicBodyInequalityEnum2(t *testing.T) {
	input := "@EnumBasicBody@ENUM2 != @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM2",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumBasicBodyDifferentEnums(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 != @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumBasicBodyDifferentEnumsEquality(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 == @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

// ============================================================================
// 跨枚举类型测试用例 (预期会有错误,但在 AST 解析阶段应该成功)
// ============================================================================

func TestEnumNoBodyAndEnumEmptyBodyEquality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 == @EnumEmptyBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyAndEnumEmptyBodyInequality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 != @EnumEmptyBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyAndEnumBasicBodyEquality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 == @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumNoBodyAndEnumBasicBodyInequality(t *testing.T) {
	input := "@EnumNoBody@ENUM1 != @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumNoBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyAndEnumBasicBodyEquality(t *testing.T) {
	input := "@EnumEmptyBody@ENUM1 == @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumEmptyBodyAndEnumBasicBodyInequality(t *testing.T) {
	input := "@EnumEmptyBody@ENUM1 != @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTNotEq", input,
		"ASTStaticField", "@EnumEmptyBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

// ============================================================================
// 额外测试: 复杂表达式
// ============================================================================

func TestComplexEnumExpression(t *testing.T) {
	input := "@EnumNoBody@ENUM1 == @EnumNoBody@ENUM1 && @EnumNoBody@ENUM2 != @EnumNoBody@ENUM1"
	expr := parseExpression(t, input)

	// 构建期望的 AST 结构
	expected := ExpectedNode{
		Type:     "ASTAnd",
		Fragment: input,
		Children: []ExpectedNode{
			{
				Type:     "ASTEq",
				Fragment: "@EnumNoBody@ENUM1 == @EnumNoBody@ENUM1",
				Children: []ExpectedNode{
					{Type: "ASTStaticField", Fragment: "@EnumNoBody@ENUM1"},
					{Type: "ASTStaticField", Fragment: "@EnumNoBody@ENUM1"},
				},
			},
			{
				Type:     "ASTNotEq",
				Fragment: "@EnumNoBody@ENUM2 != @EnumNoBody@ENUM1",
				Children: []ExpectedNode{
					{Type: "ASTStaticField", Fragment: "@EnumNoBody@ENUM2"},
					{Type: "ASTStaticField", Fragment: "@EnumNoBody@ENUM1"},
				},
			},
		},
	}

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

// ============================================================================
// 算术运算符测试 (虽然枚举通常不支持算术运算,但 AST 应该能解析)
// ============================================================================

func TestEnumArithmeticAddition(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 + @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTAdd", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumArithmeticSubtraction(t *testing.T) {
	input := "@EnumBasicBody@ENUM2 - @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTSubtract", input,
		"ASTStaticField", "@EnumBasicBody@ENUM2",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumArithmeticMultiplication(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 * @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTMultiply", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumArithmeticDivision(t *testing.T) {
	input := "@EnumBasicBody@ENUM2 / @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTDivide", input,
		"ASTStaticField", "@EnumBasicBody@ENUM2",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

// ============================================================================
// 比较运算符测试
// ============================================================================

func TestEnumLessThan(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 < @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTLess", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumGreaterThan(t *testing.T) {
	input := "@EnumBasicBody@ENUM2 > @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTGreater", input,
		"ASTStaticField", "@EnumBasicBody@ENUM2",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumLessThanOrEqual(t *testing.T) {
	input := "@EnumBasicBody@ENUM1 <= @EnumBasicBody@ENUM2"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTLessEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM1",
		"ASTStaticField", "@EnumBasicBody@ENUM2",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}

func TestEnumGreaterThanOrEqual(t *testing.T) {
	input := "@EnumBasicBody@ENUM2 >= @EnumBasicBody@ENUM1"
	expr := parseExpression(t, input)

	expected := createBinaryExpected(
		"ASTGreaterEq", input,
		"ASTStaticField", "@EnumBasicBody@ENUM2",
		"ASTStaticField", "@EnumBasicBody@ENUM1",
	)

	if !Check(expr, expected) {
		t.Fatal("AST 检查失败")
	}
}
