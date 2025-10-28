package ast

import (
	"fmt"
	"strings"
)

func Version() string {
	return "v1.0.1"
}

// Node AST节点接口
type Node interface {
	String() string
	Type() string
}

// Expression 表达式节点接口
type Expression interface {
	Node
	expressionNode()
}

// Statement 语句节点接口
type Statement interface {
	Node
	statementNode()
}

// =============================================================================
// 表达式节点实现
// =============================================================================

// BaseExpression 基础表达式结构
type BaseExpression struct{}

func (be *BaseExpression) expressionNode() {}

// SequenceExpression 序列表达式 (逗号分隔)
type SequenceExpression struct {
	BaseExpression
	Expressions []Expression
}

func (se *SequenceExpression) String() string {
	var exprs []string
	for _, expr := range se.Expressions {
		exprStr := expr.String()
		// 根据 Java OGNL 的行为，序列中的某些表达式需要添加括号
		// 例如：二元表达式、条件表达式等需要括号以增强可读性
		if needsParenthesesInSequence(expr) {
			exprStr = "(" + exprStr + ")"
		}
		exprs = append(exprs, exprStr)
	}
	return strings.Join(exprs, ", ")
}

// needsParenthesesInSequence 判断表达式在序列中是否需要括号
func needsParenthesesInSequence(expr Expression) bool {
	switch expr.(type) {
	case *BinaryExpression:
		// 二元表达式在序列中需要括号，如 (#five + #six)
		return true
	case *ConditionalExpression:
		// 条件表达式在序列中也可能需要括号
		return true
	default:
		return false
	}
}

func (se *SequenceExpression) Type() string { return "ASTSequence" }

// AssignmentExpression 赋值表达式
type AssignmentExpression struct {
	BaseExpression
	Left  Expression
	Right Expression
}

func (ae *AssignmentExpression) String() string {
	if ae.Right != nil {
		return fmt.Sprintf("%s = %s", ae.Left.String(), ae.Right.String())
	}
	return ae.Left.String()
}

func (ae *AssignmentExpression) Type() string { return "ASTAssign" }

// ConditionalExpression 三元条件表达式 (? :)
type ConditionalExpression struct {
	BaseExpression
	Test        Expression
	Consequent  Expression
	Alternative Expression
}

func (ce *ConditionalExpression) String() string {
	if ce.Consequent != nil && ce.Alternative != nil {
		testStr := ce.Test.String()
		// 如果测试条件是二元表达式,添加括号以增强可读性
		if _, ok := ce.Test.(*BinaryExpression); ok {
			testStr = "(" + testStr + ")"
		}
		return fmt.Sprintf("%s ? %s : %s", testStr, ce.Consequent.String(), ce.Alternative.String())
	}
	return ce.Test.String()
}

func (ce *ConditionalExpression) Type() string { return "ASTTest" }

// BinaryExpression 二元表达式
type BinaryExpression struct {
	BaseExpression
	Left     Expression
	Operator TokenType
	Right    Expression
}

func (be *BinaryExpression) String() string {
	// 顶层表达式不需要外层括号
	return be.StringWithContext(0, false)
}

// StringWithParentPrecedence 根据父节点优先级生成字符串表示
func (be *BinaryExpression) StringWithParentPrecedence(parentPrecedence int) string {
	return be.StringWithContext(parentPrecedence, false)
}

// StringWithContext 根据上下文生成字符串表示
// parentPrecedence: 父节点的优先级
// isRightChild: 是否作为父节点的右子节点
func (be *BinaryExpression) StringWithContext(parentPrecedence int, isRightChild bool) string {
	leftStr := be.Left.String()
	rightStr := be.Right.String()

	// 如果左子节点是二元表达式,添加括号以增强可读性
	if _, ok := be.Left.(*BinaryExpression); ok {
		leftStr = fmt.Sprintf("(%s)", leftStr)
	}

	// 如果右子节点是二元表达式,添加括号以增强可读性
	if _, ok := be.Right.(*BinaryExpression); ok {
		rightStr = fmt.Sprintf("(%s)", rightStr)
	}

	result := fmt.Sprintf("%s %s %s", leftStr, be.operatorString(), rightStr)

	// 决定是否需要为整个表达式添加括号
	needsParentheses := false

	// 情况1: 当前节点优先级低于父节点优先级
	if parentPrecedence > be.operatorPrecedence() {
		needsParentheses = true
	}

	// 情况2: 当前节点优先级高于父节点优先级
	// 为了明确表达式结构,也加括号
	// 例如: (5L & 3) | (5 ^ 3) 而不是 5L & 3 | 5 ^ 3
	if parentPrecedence > 0 && parentPrecedence < be.operatorPrecedence() {
		needsParentheses = true
	}

	// 情况3: 作为右子节点出现时,为了明确结构总是加括号
	// (这适用于在 printASTStructure 中显示详细结构时)
	if isRightChild {
		needsParentheses = true
	}

	if needsParentheses {
		result = fmt.Sprintf("(%s)", result)
	}

	return result
}

// operatorPrecedence 返回运算符优先级 (数字越大优先级越高)
func (be *BinaryExpression) operatorPrecedence() int {
	switch be.Operator {
	case OR:
		return 1
	case AND:
		return 2
	case BIT_OR:
		return 3
	case XOR:
		return 4
	case BIT_AND:
		return 5
	case EQ, NOT_EQ:
		return 6
	case LT, GT, LT_EQ, GT_EQ, IN, NOT_IN, INSTANCEOF:
		return 7
	case SHL, SHR, USHR:
		return 8
	case PLUS, MINUS:
		return 9
	case MULTIPLY, DIVIDE, MODULO:
		return 10
	default:
		return 0
	}
}

// GetPrecedence 导出的方法,返回运算符优先级
func (be *BinaryExpression) GetPrecedence() int {
	return be.operatorPrecedence()
}

// isLeftAssociative 判断运算符是否是左结合
func (be *BinaryExpression) isLeftAssociative() bool {
	// 大多数二元运算符都是左结合的
	return true
}

func (be *BinaryExpression) operatorString() string {
	switch be.Operator {
	case OR:
		return "||"
	case AND:
		return "&&"
	case BIT_OR:
		return "|"
	case XOR:
		return "^"
	case BIT_AND:
		return "&"
	case EQ:
		return "=="
	case NOT_EQ:
		return "!="
	case LT:
		return "<"
	case GT:
		return ">"
	case LT_EQ:
		return "<="
	case GT_EQ:
		return ">="
	case IN:
		return "in"
	case NOT_IN:
		return "not in"
	case SHL:
		return "<<"
	case SHR:
		return ">>"
	case USHR:
		return ">>>"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	case MODULO:
		return "%"
	default:
		return "UNKNOWN"
	}
}

func (be *BinaryExpression) Type() string {
	switch be.Operator {
	case PLUS:
		return "ASTAdd"
	case MINUS:
		return "ASTSubtract"
	case MULTIPLY:
		return "ASTMultiply"
	case DIVIDE:
		return "ASTDivide"
	case MODULO:
		return "ASTRemainder"
	case EQ:
		return "ASTEq"
	case NOT_EQ:
		return "ASTNotEq"
	case LT:
		return "ASTLess"
	case LT_EQ:
		return "ASTLessEq"
	case GT:
		return "ASTGreater"
	case GT_EQ:
		return "ASTGreaterEq"
	case AND:
		return "ASTAnd"
	case OR:
		return "ASTOr"
	case BIT_AND:
		return "ASTBitAnd"
	case BIT_OR:
		return "ASTBitOr"
	case XOR:
		return "ASTXor"
	case SHL:
		return "ASTShiftLeft"
	case SHR:
		return "ASTShiftRight"
	case USHR:
		return "ASTUnsignedShiftRight"
	case IN:
		return "ASTIn"
	case NOT_IN:
		return "ASTNotIn"
	case INSTANCEOF:
		return "ASTInstanceof"
	default:
		return "BinaryExpression"
	}
}

// UnaryExpression 一元表达式
type UnaryExpression struct {
	BaseExpression
	Operator TokenType
	Operand  Expression
}

func (ue *UnaryExpression) String() string {
	operandStr := ue.Operand.String()

	// 如果操作数是二元表达式，需要添加括号
	if _, ok := ue.Operand.(*BinaryExpression); ok {
		operandStr = fmt.Sprintf("(%s)", operandStr)
	}

	return fmt.Sprintf("%s%s", ue.operatorString(), operandStr)
}

func (ue *UnaryExpression) operatorString() string {
	switch ue.Operator {
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case BIT_NOT:
		return "~"
	case NOT:
		return "!"
	default:
		return "UNKNOWN"
	}
}

func (ue *UnaryExpression) Type() string {
	switch ue.Operator {
	case MINUS:
		return "ASTNegate"
	case NOT:
		return "ASTNot"
	case BIT_NOT:
		return "ASTBitNegate"
	case PLUS:
		return "ASTConst"
	default:
		return "UnaryExpression"
	}
}

// InstanceofExpression instanceof表达式
type InstanceofExpression struct {
	BaseExpression
	Operand    Expression
	TargetType string
	TypeNode   Expression // 类型节点(ASTConst),用于AST树结构
}

func (ie *InstanceofExpression) String() string {
	return fmt.Sprintf("%s instanceof %s", ie.Operand.String(), ie.TargetType)
}

func (ie *InstanceofExpression) Type() string { return "ASTInstanceof" }

// LambdaExpression Lambda 表达式 :[expression]
// 在 OGNL 中，Lambda 表达式用于定义匿名函数
// 例如: :[#this * 2] 定义一个将参数乘以2的函数
// 例如: #fact=:[#this <= 1 ? 1 : #fact(#this-1) * #this] 定义递归阶乘函数
type LambdaExpression struct {
	BaseExpression
	Body Expression // Lambda 函数体
}

func (le *LambdaExpression) String() string {
	return fmt.Sprintf(":[%s]", le.Body.String())
}

func (le *LambdaExpression) Type() string { return "ASTLambda" }

// ChainExpression 链式表达式 (导航链)
// 在 Java OGNL 中，ASTChain 可以有多个子节点，不是二叉树结构
type ChainExpression struct {
	BaseExpression
	// 保留 Object 和 Property 以保持向后兼容，但新的实现使用 Children
	Object   Expression
	Property Expression
	// Children 包含链式表达式的所有子节点
	// 例如: "foo".bar()[0] 会有三个子节点: "foo", bar(), [0]
	Children []Expression
}

func (ce *ChainExpression) String() string {
	if len(ce.Children) == 0 {
		return ""
	}

	var parts []string
	for i, child := range ce.Children {
		childStr := child.String()
		if i == 0 {
			parts = append(parts, childStr)
		} else {
			// 如果子节点是索引表达式（以 [ 开头），不添加点号
			if len(childStr) > 0 && childStr[0] == '[' {
				parts = append(parts, childStr)
			} else {
				// 某些表达式作为链的一部分时需要用括号包裹
				// 例如: map[$].(expression) 而不是 map[$].expression
				if needsParenthesesInChain(child) {
					parts = append(parts, ".("+childStr+")")
				} else {
					parts = append(parts, "."+childStr)
				}
			}
		}
	}
	return strings.Join(parts, "")
}

// needsParenthesesInChain 判断表达式作为链的一部分时是否需要括号
func needsParenthesesInChain(expr Expression) bool {
	switch expr.(type) {
	case *ConditionalExpression, *BinaryExpression:
		return true
	}
	return false
}

func (ce *ChainExpression) Type() string { return "ASTChain" }

// IndexExpression 索引表达式
type IndexExpression struct {
	BaseExpression
	Object Expression
	Index  Expression
}

func (ie *IndexExpression) String() string {
	// 当 Object 为 nil 时，这个 IndexExpression 是作为 ChainExpression 的 Property 使用的
	// 在这种情况下，只显示索引部分 [index]
	if ie.Object == nil {
		return fmt.Sprintf("[%s]", ie.Index.String())
	}
	return fmt.Sprintf("%s[%s]", ie.Object.String(), ie.Index.String())
}

func (ie *IndexExpression) Type() string { return "ASTProperty" }

// CallExpression 方法调用表达式
type CallExpression struct {
	BaseExpression
	Object    Expression
	Method    string
	Arguments []Expression
}

func (ce *CallExpression) String() string {
	var args []string
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	if ce.Object != nil {
		return fmt.Sprintf("%s.%s(%s)", ce.Object.String(), ce.Method, strings.Join(args, ", "))
	}
	return fmt.Sprintf("%s(%s)", ce.Method, strings.Join(args, ", "))
}

func (ce *CallExpression) Type() string { return "ASTMethod" }

// StaticMethodExpression 静态方法调用表达式
type StaticMethodExpression struct {
	BaseExpression
	ClassName string
	Method    string
	Arguments []Expression
}

func (sme *StaticMethodExpression) String() string {
	var args []string
	for _, arg := range sme.Arguments {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("@%s@%s(%s)", sme.ClassName, sme.Method, strings.Join(args, ", "))
}

func (sme *StaticMethodExpression) Type() string { return "ASTStaticMethod" }

// StaticFieldExpression 静态字段访问表达式
type StaticFieldExpression struct {
	BaseExpression
	ClassName string
	Field     string
}

func (sfe *StaticFieldExpression) String() string {
	return fmt.Sprintf("@%s@%s", sfe.ClassName, sfe.Field)
}

func (sfe *StaticFieldExpression) Type() string { return "ASTStaticField" }

// ConstructorExpression 构造器表达式
type ConstructorExpression struct {
	BaseExpression
	ClassName string
	Arguments []Expression
	IsArray   bool
}

func (ce *ConstructorExpression) String() string {
	if ce.IsArray {
		// 对于数组构造器，区分两种情况：
		// 1. new Type[]{ elem1, elem2 } - 数组初始化
		// 2. new Type[size] - 指定大小的数组
		if len(ce.Arguments) > 0 {
			if arrayExpr, ok := ce.Arguments[0].(*ArrayExpression); ok {
				// 情况1: 数组初始化 - 第一个参数是 ArrayExpression
				var elemStrs []string
				for _, elem := range arrayExpr.Elements {
					elemStrs = append(elemStrs, elem.String())
				}
				return fmt.Sprintf("new %s[]{ %s }", ce.ClassName, strings.Join(elemStrs, ", "))
			} else {
				// 情况2: 指定大小的数组 - 第一个参数是大小表达式
				return fmt.Sprintf("new %s[%s]", ce.ClassName, ce.Arguments[0].String())
			}
		}
		// 如果没有参数，返回空数组初始化
		return fmt.Sprintf("new %s[]{ }", ce.ClassName)
	}
	// 普通构造器
	var args []string
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("new %s(%s)", ce.ClassName, strings.Join(args, ", "))
}

func (ce *ConstructorExpression) Type() string { return "ASTCtor" }

// ProjectionExpression 投影表达式 {expr}
type ProjectionExpression struct {
	BaseExpression
	Object     Expression
	Expression Expression
}

func (pe *ProjectionExpression) String() string {
	exprStr := pe.Expression.String()
	// 只有复杂表达式才需要括号(二元运算、条件、赋值等)
	if needsParenthesesInProjection(pe.Expression) {
		exprStr = fmt.Sprintf("(%s)", exprStr)
	}
	if pe.Object != nil {
		return fmt.Sprintf("%s{%s}", pe.Object.String(), exprStr)
	} else {
		return fmt.Sprintf("{%s}", exprStr)
	}
}

func (pe *ProjectionExpression) Type() string { return "ASTProject" }

// needsParenthesesInProjection 判断在 Projection 中是否需要括号
func needsParenthesesInProjection(expr Expression) bool {
	switch expr.(type) {
	case *BinaryExpression, *ConditionalExpression, *AssignmentExpression, *SequenceExpression:
		return true
	default:
		return false
	}
}

// SelectionExpression 选择表达式 {? expr}
type SelectionExpression struct {
	BaseExpression
	Object     Expression
	Expression Expression
	SelectType string // "all", "first", "last"
}

func (se *SelectionExpression) String() string {
	prefix := "?"
	switch se.SelectType {
	case "first":
		prefix = "^"
	case "last":
		prefix = "$"
	}
	exprStr := se.Expression.String()
	// 只有复杂表达式才需要括号
	if needsParenthesesInSelection(se.Expression) {
		exprStr = fmt.Sprintf("(%s)", exprStr)
	}
	if se.Object != nil {
		return fmt.Sprintf("%s{%s %s}", se.Object.String(), prefix, exprStr)
	} else {
		return fmt.Sprintf("{%s %s}", prefix, exprStr)
	}
}

func (se *SelectionExpression) Type() string { return "ASTSelectFirst" }

// needsParenthesesInSelection 判断在 Selection 中是否需要括号
func needsParenthesesInSelection(expr Expression) bool {
	switch expr.(type) {
	case *BinaryExpression, *ConditionalExpression, *AssignmentExpression, *SequenceExpression:
		return true
	default:
		return false
	}
}

// EvalExpression 求值表达式 (expr)(arg)
// 对应 Java OGNL 的 ASTEval，有两个子节点：
// 1. Target: 被求值的表达式（通常是Lambda表达式）
// 2. Argument: 求值的参数
type EvalExpression struct {
	BaseExpression
	Target   Expression // 被求值的表达式，如 :[33]
	Argument Expression // 参数，如 20
}

func (ee *EvalExpression) String() string {
	return fmt.Sprintf("(%s)(%s)", ee.Target.String(), ee.Argument.String())
}

func (ee *EvalExpression) Type() string { return "ASTEval" }

// =============================================================================
// 字面量和标识符
// =============================================================================

// Identifier 标识符 (对应Java的ASTProperty)
type Identifier struct {
	BaseExpression
	Value string
	// NameNode 存储属性名的子节点 (对应Java ASTProperty的子节点ASTConst)
	NameNode Expression
}

func (i *Identifier) String() string { return i.Value }
func (i *Identifier) Type() string   { return "ASTProperty" }

// Literal 字面量
type Literal struct {
	BaseExpression
	Value interface{}
	Raw   string
}

func (l *Literal) String() string {
	// 对于 nil 值,直接返回 "null"
	if l.Value == nil {
		return "null"
	}

	// 对于 BigDecimal 类型（带 b/B 后缀），保留后缀
	if len(l.Raw) > 0 {
		lastChar := l.Raw[len(l.Raw)-1]
		if lastChar == 'b' || lastChar == 'B' {
			// BigDecimal 类型，保留后缀
			result := l.Raw
			// 将小写 b 统一转换为大写 B
			if lastChar == 'b' {
				result = l.Raw[:len(l.Raw)-1] + "B"
			}
			// 移除 ".B" 中的小数点，如果小数点后面直接跟着 B
			// 例如: "2.B" -> "2B", "2.5B" -> "2.5B"
			if len(result) >= 3 && result[len(result)-2] == '.' {
				// 检查小数点前面是否是数字
				beforeDot := result[:len(result)-2]
				if len(beforeDot) > 0 && (beforeDot[len(beforeDot)-1] >= '0' && beforeDot[len(beforeDot)-1] <= '9') {
					// 移除小数点: "2.B" -> "2B"
					result = beforeDot + "B"
				}
			}
			return result
		}
		// BigInteger 类型（带 h/H 后缀），也保留后缀
		if lastChar == 'h' || lastChar == 'H' {
			// 将小写 h 统一转换为大写 H
			if lastChar == 'h' {
				return l.Raw[:len(l.Raw)-1] + "H"
			}
			return l.Raw
		}
		// Long 类型（带 l/L 后缀），也保留后缀
		if lastChar == 'l' || lastChar == 'L' {
			// 将小写 l 统一转换为大写 L
			if lastChar == 'l' {
				return l.Raw[:len(l.Raw)-1] + "L"
			}
			return l.Raw
		}
	}

	// 对于其他数值类型，返回规范化的表示
	switch v := l.Value.(type) {
	case int, int32, int64:
		// 检查是否是字符字面量（Raw 以单引号开始）
		if len(l.Raw) > 0 && l.Raw[0] == '\'' {
			// 字符类型，使用 Raw 保持单引号格式
			return l.Raw
		}
		return fmt.Sprintf("%v", v)
	case float32:
		// 浮点数始终显示小数点
		s := fmt.Sprintf("%v", v)
		if !strings.Contains(s, ".") && !strings.Contains(s, "e") && !strings.Contains(s, "E") {
			s += ".0"
		}
		return s
	case float64:
		// 浮点数始终显示小数点
		s := fmt.Sprintf("%v", v)
		if !strings.Contains(s, ".") && !strings.Contains(s, "e") && !strings.Contains(s, "E") {
			s += ".0"
		}
		return s
	case bool:
		return fmt.Sprintf("%v", v)
	case string:
		// 字符串保持带引号
		return l.Raw
	default:
		return l.Raw
	}
}
func (l *Literal) Type() string { return "ASTConst" }

// LambdaLiteral Lambda表达式字面量
// 在Java OGNL中，Lambda表达式 :[expr] 被解析为 ASTConst 节点，该节点有一个子节点（Lambda body）
// 在Go中，我们使用 LambdaLiteral 来表示这个特殊的 ASTConst 节点
type LambdaLiteral struct {
	BaseExpression
	Body Expression // Lambda 函数体
}

func (ll *LambdaLiteral) String() string {
	return fmt.Sprintf(":[%s]", ll.Body.String())
}

func (ll *LambdaLiteral) Type() string { return "ASTConst" }

// ThisExpression this引用
type ThisExpression struct {
	BaseExpression
}

func (te *ThisExpression) String() string { return "#this" }
func (te *ThisExpression) Type() string   { return "ASTThisVarRef" }

// RootExpression root引用
type RootExpression struct {
	BaseExpression
}

func (re *RootExpression) String() string { return "#root" }
func (re *RootExpression) Type() string   { return "ASTRootVarRef" }

// VariableExpression 变量引用
type VariableExpression struct {
	BaseExpression
	Name string
}

func (ve *VariableExpression) String() string { return "#" + ve.Name }
func (ve *VariableExpression) Type() string {
	// root 和 this 使用 ASTRootVarRef，其他变量使用 ASTVarRef
	if ve.Name == "root" || ve.Name == "this" {
		return "ASTRootVarRef"
	}
	return "ASTVarRef"
}

// ArrayExpression 数组字面量
type ArrayExpression struct {
	BaseExpression
	Elements []Expression
}

func (ae *ArrayExpression) String() string {
	var elements []string
	for _, elem := range ae.Elements {
		elements = append(elements, elem.String())
	}
	return fmt.Sprintf("{ %s }", strings.Join(elements, ", "))
}

func (ae *ArrayExpression) Type() string { return "ASTList" }

// MapExpression Map字面量
type MapExpression struct {
	BaseExpression
	Pairs     []Expression // 改为 Expression 列表，每个元素是 KeyValueExpression
	ClassName string       // 可选的类型名，用于 #@ClassName@{...} 语法
}

// KeyValueExpression 键值对表达式 (对应 Java 的 ASTKeyValue)
type KeyValueExpression struct {
	BaseExpression
	Key   Expression
	Value Expression
}

func (kve *KeyValueExpression) String() string {
	if kve.Value != nil {
		return fmt.Sprintf("%s : %s", kve.Key.String(), kve.Value.String())
	}
	// 当值为 nil 时，显示 key : null（对应 Java OGNL 的行为）
	return fmt.Sprintf("%s : null", kve.Key.String())
}

func (kve *KeyValueExpression) Type() string { return "ASTKeyValue" }

func (me *MapExpression) String() string {
	var pairs []string
	for _, pair := range me.Pairs {
		pairs = append(pairs, pair.String())
	}
	if me.ClassName != "" {
		// 带类型的 Map: #@ClassName@{ ... }
		return fmt.Sprintf("#@%s@{ %s }", me.ClassName, strings.Join(pairs, ", "))
	}
	// 普通 Map: #{ ... }
	return fmt.Sprintf("#{ %s }", strings.Join(pairs, ", "))
}

func (me *MapExpression) Type() string { return "ASTMap" }

// DynamicSubscriptExpression 动态下标表达式
type DynamicSubscriptExpression struct {
	BaseExpression
	Object        Expression
	SubscriptType DynamicSubscriptType
}

func (dse *DynamicSubscriptExpression) String() string {
	var symbol string
	switch dse.SubscriptType {
	case FIRST:
		symbol = "^"
	case MID:
		symbol = "|"
	case LAST:
		symbol = "$"
	case ALL:
		symbol = "*"
	}
	if dse.Object != nil {
		return fmt.Sprintf("%s[%s]", dse.Object.String(), symbol)
	}
	return fmt.Sprintf("[%s]", symbol)
}

func (dse *DynamicSubscriptExpression) Type() string { return "ASTDynamicSubscript" }
