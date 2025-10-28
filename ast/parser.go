package ast

import (
	"fmt"
	"strconv"
)

const (
	// MaxParseIterations 最大解析迭代次数，防止死循环
	MaxParseIterations = 20000
)

// Parser OGNL递归下降解析器
type Parser struct {
	lexer          *Lexer
	current        Token
	peek           Token
	errors         []string
	position       int
	iterationCount int // 解析迭代计数器
}

// New 创建新的解析器
func New(l *Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	// 读取两个token，current和peek
	p.nextToken()
	p.nextToken()

	return p
}

// Errors 返回解析错误
func (p *Parser) Errors() []string {
	return p.errors
}

// CurrentToken 返回当前 token (用于调试)
func (p *Parser) CurrentToken() Token {
	return p.current
}

// PeekToken 返回下一个 token (用于调试)
func (p *Parser) PeekToken() Token {
	return p.peek
}

// nextToken 前进到下一个token
func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lexer.NextToken()
	p.position++
}

// currentTokenIs 检查当前token类型
func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.current.Type == t
}

// peekTokenIs 检查下一个token类型
func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peek.Type == t
}

// expectPeek 检查并消费下一个token
func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// peekError 添加peek错误
func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		TokenTypeNames[t], TokenTypeNames[p.peek.Type])
	p.errors = append(p.errors, msg)
}

// currentError 添加当前token错误
func (p *Parser) currentError(msg string) {
	p.errors = append(p.errors, fmt.Sprintf("at token %s: %s",
		TokenTypeNames[p.current.Type], msg))
}

// checkIterationLimit 检查是否超过迭代限制，防止死循环
func (p *Parser) checkIterationLimit() bool {
	p.iterationCount++
	if p.iterationCount > MaxParseIterations {
		p.currentError(fmt.Sprintf("parse iteration limit exceeded (%d), possible infinite loop", MaxParseIterations))
		return false
	}
	return true
}

// skipWhitespace 跳过空白字符
func (p *Parser) skipWhitespace() {
	for p.current.Type == WHITESPACE {
		p.nextToken()
	}
}

// =============================================================================
// 主要解析方法 - 对应JavaCC解析器的层次结构
// =============================================================================

// ParseTopLevelExpression 解析顶级表达式
func (p *Parser) ParseTopLevelExpression() (Expression, error) {
	expr := p.parseExpression()
	// parseExpression 返回后，current 指向表达式之后的 token
	// 应该检查 current 而不是 peek
	if p.current.Type != EOF {
		p.currentError(fmt.Sprintf("expected EOF at end of expression, got %s", TokenTypeNames[p.current.Type]))
	}
	if len(p.errors) > 0 {
		return expr, fmt.Errorf("parser errors: %v", p.errors)
	}
	return expr, nil
}

// parseExpression 解析表达式序列 (对应expression)
func (p *Parser) parseExpression() Expression {
	expr := p.parseAssignmentExpression()

	if p.current.Type != COMMA {
		return expr
	}

	// 处理逗号分隔的表达式序列
	exprs := []Expression{expr}

	for p.current.Type == COMMA {
		p.nextToken() // move to next expression
		exprs = append(exprs, p.parseAssignmentExpression())
	}

	return &SequenceExpression{Expressions: exprs}
}

// parseAssignmentExpression 解析赋值表达式 (对应assignmentExpression)
func (p *Parser) parseAssignmentExpression() Expression {
	expr := p.parseConditionalTestExpression()

	if p.current.Type == ASSIGN {
		p.nextToken() // move to right side
		right := p.parseAssignmentExpression()
		return &AssignmentExpression{Left: expr, Right: right}
	}

	return expr
}

// parseConditionalTestExpression 解析条件表达式 (对应conditionalTestExpression)
func (p *Parser) parseConditionalTestExpression() Expression {
	expr := p.parseLogicalOrExpression()

	if p.current.Type == QUESTION {
		p.nextToken() // move to consequent
		consequent := p.parseConditionalTestExpression()

		if p.current.Type != COLON {
			p.currentError(fmt.Sprintf("expected : in ternary expression, got %s", p.current.Value))
			return nil
		}
		p.nextToken() // move to alternative
		alternative := p.parseConditionalTestExpression()

		return &ConditionalExpression{
			Test:        expr,
			Consequent:  consequent,
			Alternative: alternative,
		}
	}

	return expr
}

// parseLogicalOrExpression 解析逻辑或表达式 (对应logicalOrExpression)
// parseLogicalOrExpression 解析逻辑或表达式 (对应logicalOrExpression)
func (p *Parser) parseLogicalOrExpression() Expression {
	left := p.parseLogicalAndExpression()

	for p.current.Type == OR {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseLogicalAndExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseLogicalAndExpression 解析逻辑与表达式 (对应logicalAndExpression)
func (p *Parser) parseLogicalAndExpression() Expression {
	left := p.parseInclusiveOrExpression()

	for p.current.Type == AND {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseInclusiveOrExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseInclusiveOrExpression 解析按位或表达式 (对应inclusiveOrExpression)
func (p *Parser) parseInclusiveOrExpression() Expression {
	left := p.parseExclusiveOrExpression()

	for p.current.Type == BIT_OR {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseExclusiveOrExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseExclusiveOrExpression 解析异或表达式 (对应exclusiveOrExpression)
func (p *Parser) parseExclusiveOrExpression() Expression {
	left := p.parseAndExpression()

	for p.current.Type == XOR {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseAndExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseAndExpression 解析按位与表达式 (对应andExpression)
func (p *Parser) parseAndExpression() Expression {
	left := p.parseEqualityExpression()

	for p.current.Type == BIT_AND {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseEqualityExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseEqualityExpression 解析相等性表达式 (对应equalityExpression)
func (p *Parser) parseEqualityExpression() Expression {
	left := p.parseRelationalExpression()

	for p.current.Type == EQ || p.current.Type == NOT_EQ {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseRelationalExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseRelationalExpression 解析关系表达式 (对应relationalExpression)
func (p *Parser) parseRelationalExpression() Expression {
	left := p.parseShiftExpression()

	for p.isRelationalOperator(p.current.Type) {
		operator := p.current.Type
		p.nextToken() // consume operator

		// 处理 "not in" 情况
		if operator == NOT && p.peekTokenIs(IN) {
			p.nextToken() // consume "in"
			operator = NOT_IN
		}

		right := p.parseShiftExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// isRelationalOperator 检查是否为关系运算符
func (p *Parser) isRelationalOperator(t TokenType) bool {
	return t == LT || t == GT || t == LT_EQ ||
		t == GT_EQ || t == IN || t == NOT_IN ||
		(t == NOT && p.peek.Type == IN)
}

// parseShiftExpression 解析位移表达式 (对应shiftExpression)
func (p *Parser) parseShiftExpression() Expression {
	left := p.parseAdditiveExpression()

	for p.isShiftOperator(p.current.Type) {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseAdditiveExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// isShiftOperator 检查是否为位移运算符
func (p *Parser) isShiftOperator(t TokenType) bool {
	return t == SHL || t == SHR || t == USHR
}

// parseAdditiveExpression 解析加减表达式 (对应additiveExpression)
func (p *Parser) parseAdditiveExpression() Expression {
	left := p.parseMultiplicativeExpression()

	for p.current.Type == PLUS || p.current.Type == MINUS {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseMultiplicativeExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// parseMultiplicativeExpression 解析乘除表达式 (对应multiplicativeExpression)
func (p *Parser) parseMultiplicativeExpression() Expression {
	left := p.parseUnaryExpression()

	for p.isMultiplicativeOperator(p.current.Type) {
		operator := p.current.Type
		p.nextToken() // move to right operand
		right := p.parseUnaryExpression()
		left = &BinaryExpression{Left: left, Operator: operator, Right: right}
	}

	return left
}

// isMultiplicativeOperator 检查是否为乘除运算符
func (p *Parser) isMultiplicativeOperator(t TokenType) bool {
	return t == MULTIPLY || t == DIVIDE || t == MODULO
}

// parseUnaryExpression 解析一元表达式 (对应unaryExpression)
func (p *Parser) parseUnaryExpression() Expression {
	switch p.current.Type {
	case PLUS:
		// 在 OGNL 中，+号作为正号前缀时直接忽略，返回操作数本身
		p.nextToken() // move to operand
		return p.parseUnaryExpression()
	case MINUS, BIT_NOT, NOT:
		operator := p.current.Type
		p.nextToken() // move to operand
		operand := p.parseUnaryExpression()
		return &UnaryExpression{Operator: operator, Operand: operand}
	default:
		expr := p.parseNavigationChain()

		// 检查instanceof (parseNavigationChain 返回后 current 已经指向下一个token)
		if p.currentTokenIs(INSTANCEOF) {
			p.nextToken() // move to type name
			if p.current.Type != IDENT {
				p.currentError(fmt.Sprintf("expected type name after instanceof, got %s", TokenTypeNames[p.current.Type]))
				return nil
			}

			// 构建类型名
			className := p.current.Value
			for p.peekTokenIs(DOT) {
				p.nextToken() // consume dot
				if !p.expectPeek(IDENT) {
					return nil
				}
				className += "." + p.current.Value
			}

			p.nextToken() // move past the type name

			return &InstanceofExpression{
				Operand:    expr,
				TargetType: className,
				TypeNode:   &Literal{Value: className, Raw: fmt.Sprintf("\"%s\"", className)},
			}
		}

		return expr
	}
}

// parseNavigationChain 解析导航链 (对应navigationChain)
func (p *Parser) parseNavigationChain() Expression {
	left := p.parsePrimaryExpression()
	return p.parseNavigationChainContinue(left)
}

// parseNavigationChainContinue 继续解析导航链（用于静态引用等方法返回后继续链式调用）
func (p *Parser) parseNavigationChainContinue(left Expression) Expression {
	// 收集所有链式操作的子节点
	var children []Expression
	children = append(children, left)

navigationLoop:
	for p.isNavigationOperator(p.current.Type) {
		// 防止死循环：检查迭代次数
		if !p.checkIterationLimit() {
			return nil
		}

		switch p.current.Type {
		case DOT:
			// DOT 已经是当前 token，解析右侧的属性或方法
			right := p.parseChainRightSide()
			if right != nil {
				children = append(children, right)
			} else {
				// 如果 parseChainRightSide 返回 nil，说明遇到错误
				// 必须退出循环，否则会死循环
				p.currentError("failed to parse chain right side, breaking loop to prevent infinite loop")
				break navigationLoop
			}
		case LBRACK:
			// LBRACK 是当前 token，current = [
			p.nextToken() // 移动到索引表达式的第一个 token

			// 检查是否是动态下标 [^], [|], [$]
			if p.isDynamicSubscript() {
				// 创建字符字面量作为索引
				var symbol string
				switch p.current.Type {
				case XOR:
					symbol = "^"
				case BIT_OR:
					symbol = "|"
				case DOLLAR:
					symbol = "$"
				}
				indexLiteral := &Literal{Value: symbol, Raw: symbol}
				p.nextToken() // consume ^, |, or $
				p.nextToken() // consume ]

				// 创建普通的 IndexExpression，索引是字符字面量
				indexExpr := &IndexExpression{Object: nil, Index: indexLiteral}
				children = append(children, indexExpr)
			} else {
				// 解析索引表达式
				index := p.parseExpression()
				// parseExpression 返回后，current 应该指向 ]（表达式之后的第一个 token）
				if p.current.Type != RBRACK {
					p.currentError(fmt.Sprintf("expected ] after index expression, got %s", TokenTypeNames[p.current.Type]))
					return nil
				}
				// 现在 current 是 ]，我们需要移动到下一个 token
				p.nextToken() // 移动过 ]

				// 创建 IndexExpression，不设置 Object（将由 ChainExpression 管理）
				indexExpr := &IndexExpression{Object: nil, Index: index}
				children = append(children, indexExpr)
			}
		case DYNAMIC_SUBSCRIPT:
			// DYNAMIC_SUBSCRIPT 已经是当前 token
			subscriptType := p.parseDynamicSubscriptType(p.current.Value)
			p.nextToken() // consume dynamic subscript
			dynamicExpr := &DynamicSubscriptExpression{
				Object:        nil, // 不设置 Object
				SubscriptType: subscriptType,
			}
			children = append(children, dynamicExpr)
		case LPAREN:
			// 处理 (arg) 的情况
			// 根据前一个节点的类型判断是 ASTEval 还是 ASTMethod
			// 如果前一个节点是 VarRef、Lambda 等可求值表达式，则创建 ASTEval
			// 否则创建 ASTMethod（方法调用）

			p.nextToken() // consume LPAREN, 移动到参数列表

			// 解析参数
			var argument Expression
			if p.current.Type != RPAREN {
				argument = p.parseExpression()
			}

			// parseExpression 返回时，current 应该在 RPAREN 上
			if p.current.Type != RPAREN {
				p.currentError(fmt.Sprintf("expected ) after argument, got %s", TokenTypeNames[p.current.Type]))
				return nil
			}
			p.nextToken() // consume RPAREN

			// 判断是否应该创建 ASTEval
			// 如果 children 中只有一个元素且该元素是可求值的表达式，创建 ASTEval
			if len(children) == 1 && p.isEvaluableExpression(children[0]) {
				// 创建 ASTEval 表达式
				evalExpr := &EvalExpression{
					Target:   children[0],
					Argument: argument,
				}
				// 重置 children，将 evalExpr 作为新的起点
				children = []Expression{evalExpr}
			} else {
				// 创建方法调用表达式
				var arguments []Expression
				if argument != nil {
					arguments = []Expression{argument}
				}
				callExpr := &CallExpression{
					Object:    nil,
					Method:    "",
					Arguments: arguments,
				}
				children = append(children, callExpr)
			}
		}
	}

	// 如果只有一个子节点，直接返回它（没有链式操作）
	if len(children) == 1 {
		return children[0]
	}

	// 创建 ChainExpression 包含所有子节点
	return &ChainExpression{
		Children: children,
	}
}

// parseChainRightSide 解析链式表达式的右侧（不包装在 ChainExpression 中）
func (p *Parser) parseChainRightSide() Expression {
	switch p.peek.Type {
	case LPAREN:
		// eval 表达式: .(expression)
		// 在 Java OGNL 中，eval 表达式直接返回内部表达式，不包装
		p.nextToken() // move to LPAREN
		p.nextToken() // move past LPAREN
		expr := p.parseExpression()
		if p.current.Type != RPAREN {
			p.currentError(fmt.Sprintf("expected ) after eval expression, got %s", TokenTypeNames[p.current.Type]))
			return nil
		}
		p.nextToken() // consume RPAREN
		// 直接返回内部表达式，不包装在 EvalExpression 中
		return expr
	case IDENT:
		p.nextToken() // move to identifier
		methodName := p.current.Value

		if p.peekTokenIs(LPAREN) {
			// 方法调用
			p.nextToken() // move to LPAREN
			p.nextToken() // move to LPAREN 后的第一个token
			arguments := p.parseArgumentList()

			// parseArgumentList 执行完后，current 应该在 RPAREN 上，需要消费它
			if p.current.Type == RPAREN {
				p.nextToken() // consume RPAREN
			}

			// 创建方法调用表达式，不设置 Object（将由 ChainExpression 管理）
			return &CallExpression{
				Object:    nil,
				Method:    methodName,
				Arguments: arguments,
			}
		} else {
			// 属性访问
			identValue := p.current.Value
			property := &Identifier{
				Value:    identValue,
				NameNode: &Literal{Value: identValue, Raw: fmt.Sprintf("%q", identValue)},
			}
			p.nextToken() // consume the identifier
			return property
		}
	case LBRACE:
		// 支持链式投影/选择表达式，如 name.{? foo } 或 name.{ foo }
		// 当前 token 是 DOT, peek 是 LBRACE。先移动到 LBRACE
		p.nextToken() // move to LBRACE
		return p.parseProjectionOrSelection(nil)
	case AT:
		// 支持链式静态引用表达式，如 Thread.@Class@method(...)
		// 当前 token 是 DOT, peek 是 AT
		p.nextToken() // move to AT (current = AT)
		p.nextToken() // move past AT to class name (current = class name)

		if p.current.Type != IDENT {
			p.currentError("expected class name after @")
			return nil
		}

		// 构建完整的类名 (package.Class)
		className := p.current.Value

		// 处理包名中的点号
		dotCount := 0
		for p.peek.Type == DOT {
			dotCount++
			if dotCount > 100 { // 包名最多100层，防止死循环
				p.currentError("package name too deep (>100 levels), possible infinite loop")
				return nil
			}
			p.nextToken() // consume dot
			p.nextToken() // move to next identifier
			if p.current.Type != IDENT {
				p.currentError("expected identifier after .")
				return nil
			}
			className += "." + p.current.Value
		}

		// 期望第二个 @
		if p.peek.Type != AT {
			p.currentError("expected @ after class name")
			return nil
		}
		p.nextToken() // consume @

		// 期望成员名称
		if p.peek.Type != IDENT {
			p.currentError("expected member name after @")
			return nil
		}
		p.nextToken() // move to member name
		memberName := p.current.Value

		// 检查是否是方法调用（有括号）
		if p.peek.Type == LPAREN {
			// 静态方法调用
			p.nextToken() // move to LPAREN
			p.nextToken() // 移动到 LPAREN 后的第一个token
			arguments := p.parseArgumentList()
			// parseArgumentList 返回时，当前应该在 RPAREN 上，需要消费它
			if p.current.Type == RPAREN {
				p.nextToken() // consume RPAREN，移动到下一个 token
			}
			return &StaticMethodExpression{
				ClassName: className,
				Method:    memberName,
				Arguments: arguments,
			}
		} else {
			// 静态字段访问
			p.nextToken() // consume the field name (move past it)
			return &StaticFieldExpression{
				ClassName: className,
				Field:     memberName,
			}
		}
	default:
		p.currentError("expected identifier, {, or @ after .")
		return nil
	}
}

// isNavigationOperator 检查是否为导航运算符
func (p *Parser) isNavigationOperator(t TokenType) bool {
	return t == DOT || t == LBRACK || t == DYNAMIC_SUBSCRIPT || t == LPAREN
}

// isEvaluableExpression 检查表达式是否可求值
// 根据 Java OGNL 的行为，以下类型的表达式后跟 (arg) 时应该创建 ASTEval：
// - ASTVarRef (变量引用)
// - ASTConst (包括 Lambda 字面量)
// - 其他可以被调用的表达式
func (p *Parser) isEvaluableExpression(expr Expression) bool {
	switch expr.(type) {
	case *VariableExpression:
		return true
	case *LambdaLiteral:
		return true
	case *Literal:
		return true
	default:
		return false
	}
}

// parseChainedExpression 解析链式表达式的右侧
func (p *Parser) parseChainedExpression(left Expression) Expression {
	switch p.peek.Type {
	case IDENT:
		// 使用新的扁平结构
		right := p.parseChainRightSide()
		if right == nil {
			return nil
		}
		return &ChainExpression{Children: []Expression{left, right}}
	case LBRACE:
		// 投影或选择
		return p.parseProjectionOrSelection(left)
	default:
		p.currentError("expected identifier or { after .")
		return nil
	}
}

// parseEvalOrIndex 解析求值或索引表达式
// 注意：这个函数目前未被使用，保留以备将来可能的用途
// 在 OGNL 中，[expr] 通常表示索引访问，而不是求值
func (p *Parser) parseEvalOrIndex(left Expression) Expression {
	p.nextToken() // move past [
	expr := p.parseExpression()
	if !p.expectPeek(RBRACK) {
		return nil
	}
	// 创建索引表达式而不是求值表达式
	return &IndexExpression{Object: left, Index: expr}
}

// parseDynamicSubscriptType 解析动态下标类型
func (p *Parser) parseDynamicSubscriptType(value string) DynamicSubscriptType {
	switch value {
	case "^":
		return FIRST
	case "|":
		return MID
	case "$":
		return LAST
	default:
		return ALL
	}
}

// isDynamicSubscript 检查当前是否是动态下标 [^], [|], [$]
func (p *Parser) isDynamicSubscript() bool {
	// current 应该是 ^, |, $ 之一，且 peek 应该是 ]
	return (p.current.Type == XOR || p.current.Type == BIT_OR || p.current.Type == DOLLAR) &&
		p.peekTokenIs(RBRACK)
}

// parseDynamicSubscriptFromTokens 从当前 token 解析动态下标类型
func (p *Parser) parseDynamicSubscriptFromTokens() DynamicSubscriptType {
	var subscriptType DynamicSubscriptType
	switch p.current.Type {
	case XOR: // ^
		subscriptType = FIRST
	case OR: // |
		subscriptType = MID
	case DOLLAR: // $
		subscriptType = LAST
	default:
		subscriptType = ALL
	}
	p.nextToken() // consume ^, |, or $
	return subscriptType
}

// parseMethodCall 解析方法调用
func (p *Parser) parseMethodCall(object Expression) Expression {
	methodName := p.current.Value
	p.nextToken() // consume method name

	// 检查当前token是否是左括号（此时current应该在LPAREN上）
	if p.current.Type != LPAREN {
		p.currentError(fmt.Sprintf("expected ( after method name, got %s", p.current.Value))
		return nil
	}

	p.nextToken() // 移动到 LPAREN 后的第一个token
	arguments := p.parseArgumentList()

	// parseArgumentList 执行完后，current 应该在 RPAREN 上，需要消费它
	if p.current.Type == RPAREN {
		p.nextToken() // consume RPAREN
	}

	return &CallExpression{
		Object:    object,
		Method:    methodName,
		Arguments: arguments,
	}
}

// parseArgumentList 解析参数列表 (调用时 current 应该在 LPAREN 后的第一个 token)
func (p *Parser) parseArgumentList() []Expression {
	args := []Expression{}

	// 检查是否是空参数列表
	if p.current.Type == RPAREN {
		// 空参数列表，直接返回
		return args
	}

	// 解析第一个参数
	args = append(args, p.parseAssignmentExpression())

	// 处理后续参数
	for p.current.Type == COMMA {
		p.nextToken() // consume comma, move to next argument
		args = append(args, p.parseAssignmentExpression())
	}

	// 此时应该在 RPAREN 上
	if p.current.Type != RPAREN {
		p.currentError(fmt.Sprintf("expected ) after arguments, got %s", p.current.Value))
		return nil
	}

	return args
} // parseProjectionOrSelection 解析投影或选择表达式
func (p *Parser) parseProjectionOrSelection(object Expression) Expression {
	p.nextToken() // consume {

	if p.currentTokenIs(QUESTION) {
		// 选择表达式 {? expr}
		p.nextToken()                         // move to expression
		expr := p.parseAssignmentExpression() // Use parseAssignmentExpression to avoid comma-sequence handling
		if p.current.Type != RBRACE {
			p.currentError(fmt.Sprintf("expected } after selection expression, got %s", TokenTypeNames[p.current.Type]))
			return nil
		}
		p.nextToken() // consume }
		return &SelectionExpression{
			Object:     object,
			Expression: expr,
			SelectType: "all",
		}
	} else {
		// 投影表达式 {expr}
		expr := p.parseAssignmentExpression() // Use parseAssignmentExpression to avoid comma-sequence handling
		if p.current.Type != RBRACE {
			p.currentError(fmt.Sprintf("expected } after projection expression, got %s", TokenTypeNames[p.current.Type]))
			return nil
		}
		p.nextToken() // consume }
		return &ProjectionExpression{
			Object:     object,
			Expression: expr,
		}
	}
}

// parsePrimaryExpression 解析主表达式 (对应primaryExpression)
func (p *Parser) parsePrimaryExpression() Expression {
	switch p.current.Type {
	case IDENT:
		identValue := p.current.Value
		p.nextToken() // consume identifier

		// 检查是否是方法调用
		if p.current.Type == LPAREN {
			// 这是一个方法调用
			p.nextToken() // consume LPAREN
			arguments := p.parseArgumentList()

			// parseArgumentList 执行完后，current 应该在 RPAREN 上
			if p.current.Type == RPAREN {
				p.nextToken() // consume RPAREN
			}

			return &CallExpression{
				Object:    nil, // 简单方法调用没有对象
				Method:    identValue,
				Arguments: arguments,
			}
		}

		// 否则是普通的标识符
		return &Identifier{
			Value:    identValue,
			NameNode: &Literal{Value: identValue, Raw: fmt.Sprintf("%q", identValue)},
		}
	case INT_LITERAL:
		literal := p.parseIntegerLiteral()
		p.nextToken() // consume integer
		return literal
	case FLT_LITERAL:
		literal := p.parseFloatLiteral()
		p.nextToken() // consume float
		return literal
	case STR_LITERAL:
		literal := p.parseStringLiteral()
		p.nextToken() // consume string
		return literal
	case CHAR_LITERAL:
		literal := p.parseCharLiteral()
		p.nextToken() // consume char
		return literal
	case TRUE:
		literal := &Literal{Value: true, Raw: "true"}
		p.nextToken() // consume true
		return literal
	case FALSE:
		literal := &Literal{Value: false, Raw: "false"}
		p.nextToken() // consume false
		return literal
	case NULL:
		literal := &Literal{Value: nil, Raw: "null"}
		p.nextToken() // consume null
		return literal
	case THIS:
		expr := &ThisExpression{}
		p.nextToken() // consume this
		return expr
	case ROOT:
		expr := &RootExpression{}
		p.nextToken() // consume root
		return expr
	case HASH:
		// 需要判断是变量引用、Map 字面量还是带类型的 Map 构造
		if p.peekTokenIs(LBRACE) {
			// Map 字面量 #{}
			p.nextToken() // 移动到 {
			return p.parseMapLiteralWithHash()
		} else if p.peekTokenIs(AT) {
			// 带类型的 Map 构造 #@ClassName@{...}
			p.nextToken() // 移动到 @
			return p.parseTypedMapConstruction()
		} else {
			// 变量引用 #variable
			return p.parseVariableReference()
		}
	case DOLLAR:
		// $ 是一个常量，表示当前上下文
		literal := &Literal{Value: "$", Raw: "$"}
		p.nextToken() // consume $
		return literal
	case LPAREN:
		return p.parseGroupedExpression()
	case LBRACK:
		// 以 [ 开头的索引表达式，如 ["values"] 或 [0]
		p.nextToken() // consume [
		index := p.parseExpression()
		if p.current.Type != RBRACK {
			p.currentError(fmt.Sprintf("expected ] after index expression, got %s", TokenTypeNames[p.current.Type]))
			return nil
		}
		p.nextToken() // consume ]
		return &IndexExpression{Object: nil, Index: index}
	case LBRACE:
		return p.parseArrayOrMapLiteral()
	case NEW:
		return p.parseConstructorCall()
	case AT:
		return p.parseStaticReference()
	case COLON:
		// Lambda 表达式 :[expression]
		return p.parseLambdaExpression()
	default:
		p.currentError(fmt.Sprintf("unexpected token: %s", p.current.Value))
		return nil
	}
}

// parseIntegerLiteral 解析整数字面量
// parseIntegerLiteral 解析整数字面量
func (p *Parser) parseIntegerLiteral() Expression {
	valueStr := p.current.Value

	// 移除整数后缀 (l, L, h, H)
	if len(valueStr) > 0 {
		lastChar := valueStr[len(valueStr)-1]
		if lastChar == 'l' || lastChar == 'L' || lastChar == 'h' || lastChar == 'H' {
			valueStr = valueStr[:len(valueStr)-1]
		}
	}

	value, err := strconv.ParseInt(valueStr, 0, 64)
	if err != nil {
		p.currentError(fmt.Sprintf("could not parse %q as integer", p.current.Value))
		return nil
	}
	return &Literal{Value: value, Raw: p.current.Value}
}

// parseFloatLiteral 解析浮点数字面量
func (p *Parser) parseFloatLiteral() Expression {
	valueStr := p.current.Value

	// 移除浮点数后缀 (d, D, f, F, b, B)
	if len(valueStr) > 0 {
		lastChar := valueStr[len(valueStr)-1]
		if lastChar == 'd' || lastChar == 'D' || lastChar == 'f' || lastChar == 'F' || lastChar == 'b' || lastChar == 'B' {
			valueStr = valueStr[:len(valueStr)-1]
		}
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		p.currentError(fmt.Sprintf("could not parse %q as float", p.current.Value))
		return nil
	}
	return &Literal{Value: value, Raw: p.current.Value}
}

// parseStringLiteral 解析字符串字面量
func (p *Parser) parseStringLiteral() Expression {
	// p.current.Value 已经是不带引号的字符串内容
	value := p.current.Value
	// Raw 应该包含引号，用于显示
	raw := fmt.Sprintf("\"%s\"", value)
	return &Literal{Value: value, Raw: raw}
}

// parseCharLiteral 解析字符字面量
func (p *Parser) parseCharLiteral() Expression {
	// p.current.Value 已经是不带引号的字符内容
	value := p.current.Value
	// Raw 应该包含单引号，用于显示
	raw := fmt.Sprintf("'%s'", value)
	if len(value) == 1 {
		return &Literal{Value: rune(value[0]), Raw: raw}
	}
	// 处理转义字符
	return &Literal{Value: value, Raw: raw}
}

// parseVariableReference 解析变量引用
func (p *Parser) parseVariableReference() Expression {
	if !p.expectPeek(IDENT) {
		return nil
	}
	varName := p.current.Value

	// 检查是否是特殊的 #this 或 #root
	if varName == "this" {
		expr := &ThisExpression{}
		p.nextToken() // consume this
		return expr
	} else if varName == "root" {
		expr := &RootExpression{}
		p.nextToken() // consume root
		return expr
	}

	varExpr := &VariableExpression{Name: varName}
	p.nextToken() // consume the identifier
	return varExpr
}

// parseGroupedExpression 解析括号表达式
func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken() // move past (
	expr := p.parseExpression()
	// parseExpression 返回后，current 应该已经在 ) 上或者已经移动过 )
	// 需要检查 current 是否是 )
	if p.current.Type != RPAREN {
		p.currentError(fmt.Sprintf("expected ) after expression, got %s ('%s')", TokenTypeNames[p.current.Type], p.current.Value))
		return nil
	}
	p.nextToken() // consume RPAREN
	return expr
}

// parseArrayOrMapLiteral 解析数组或Map字面量
func (p *Parser) parseArrayOrMapLiteral() Expression {
	// 当前 token 是 {
	// 先向前看，判断是空数组还是有内容
	if p.peekTokenIs(RBRACE) {
		// 空数组 {}
		p.nextToken() // current 移动到 }
		p.nextToken() // current 移动到 } 之后的 token
		return &ArrayExpression{Elements: []Expression{}}
	}

	p.nextToken() // 移动到第一个元素
	firstExpr := p.parseAssignmentExpression()

	// parseAssignmentExpression 返回后，current 指向表达式之后的 token
	if p.current.Type == COLON {
		// 这是一个Map
		return p.parseMapLiteral(firstExpr)
	} else {
		// 这是一个数组
		return p.parseArrayLiteral(firstExpr)
	}
}

// parseMapLiteralWithHash 解析带 # 前缀的 Map 字面量
func (p *Parser) parseMapLiteralWithHash() Expression {
	// 当前 token 是 {
	// 检查是否为空 Map
	if p.peekTokenIs(RBRACE) {
		// 空 Map #{}
		p.nextToken() // current 移动到 }
		p.nextToken() // current 移动到 } 之后的 token
		return &MapExpression{Pairs: []Expression{}}
	}

	p.nextToken() // 移动到第一个键
	firstKey := p.parseAssignmentExpression()

	// parseAssignmentExpression 返回后，current 应该是 :
	return p.parseMapLiteral(firstKey)
}

// parseTypedMapConstruction 解析带类型的 Map 构造 #@ClassName@{...}
func (p *Parser) parseTypedMapConstruction() Expression {
	// 当前 token 是 @，移动到类名
	p.nextToken()

	if p.current.Type != IDENT {
		p.currentError("expected class name after @")
		return nil
	}

	// 构建完整的类名 (package.Class)
	className := p.current.Value

	// 处理包名中的点号
	for p.peek.Type == DOT {
		p.nextToken() // consume dot
		p.nextToken() // move to next identifier
		if p.current.Type != IDENT {
			p.currentError("expected identifier after .")
			return nil
		}
		className += "." + p.current.Value
	}

	// 期望第二个 @
	if p.peek.Type != AT {
		p.currentError("expected @ after class name")
		return nil
	}
	p.nextToken() // consume @

	// 期望 {
	if p.peek.Type != LBRACE {
		p.currentError("expected { after @ in typed map construction")
		return nil
	}
	p.nextToken() // move to {

	// 解析 Map 字面量内容
	var mapExpr *MapExpression

	// 检查是否为空 Map
	if p.peekTokenIs(RBRACE) {
		// 空 Map
		p.nextToken() // current 移动到 }
		p.nextToken() // current 移动到 } 之后的 token
		mapExpr = &MapExpression{
			Pairs:     []Expression{},
			ClassName: className,
		}
	} else {
		p.nextToken() // 移动到第一个键
		firstKey := p.parseAssignmentExpression()
		mapExpr = p.parseMapLiteral(firstKey).(*MapExpression)
		// 设置类型名
		mapExpr.ClassName = className
	}

	// 直接返回带类型的 MapExpression，而不是包装在 ConstructorExpression 中
	return mapExpr
}

// parseMapLiteral 解析Map字面量
func (p *Parser) parseMapLiteral(firstKey Expression) Expression {
	pairs := []Expression{}

	// 处理第一个键值对
	// parseAssignmentExpression 返回后，current 可能是 : , 或 }
	if p.current.Type == COLON {
		p.nextToken() // 移动到值
		value := p.parseAssignmentExpression()
		pairs = append(pairs, &KeyValueExpression{Key: firstKey, Value: value})
	} else {
		// 没有冒号，说明只有键，值为 nil
		pairs = append(pairs, &KeyValueExpression{Key: firstKey, Value: nil})
	}

	// 处理后续键值对
	// parseAssignmentExpression 返回后，current 指向表达式之后的 token（可能是 , 或 }）
	for p.current.Type == COMMA {
		p.nextToken() // 移动到键
		key := p.parseAssignmentExpression()

		// parseAssignmentExpression 返回后，current 应该是 : , 或 }
		if p.current.Type == COLON {
			p.nextToken() // 移动到值
			value := p.parseAssignmentExpression()
			pairs = append(pairs, &KeyValueExpression{Key: key, Value: value})
		} else {
			pairs = append(pairs, &KeyValueExpression{Key: key, Value: nil})
		}
	}

	// 现在 current 应该是 }
	if p.current.Type != RBRACE {
		p.currentError(fmt.Sprintf("expected } at end of map, got %s", TokenTypeNames[p.current.Type]))
		return nil
	}

	// 移动过 }
	p.nextToken()

	return &MapExpression{Pairs: pairs}
}

// parseArrayLiteral 解析数组字面量
func (p *Parser) parseArrayLiteral(firstElement Expression) Expression {
	elements := []Expression{firstElement}

	// parseAssignmentExpression 返回后，current 指向表达式之后的 token（可能是 , 或 }）
	for p.current.Type == COMMA {
		p.nextToken() // 移动到下一个元素
		elements = append(elements, p.parseAssignmentExpression())
		// parseAssignmentExpression 返回后，current 又指向表达式之后的 token
	}

	// 现在 current 应该是 }
	if p.current.Type != RBRACE {
		p.currentError(fmt.Sprintf("expected } at end of array, got %s", TokenTypeNames[p.current.Type]))
		return nil
	}

	// 移动过 }
	p.nextToken()

	return &ArrayExpression{Elements: elements}
}

// parseConstructorCall 解析构造器调用
func (p *Parser) parseConstructorCall() Expression {
	if !p.expectPeek(IDENT) {
		return nil
	}

	className := p.current.Value
	// 处理包名和内部类（支持 . 和 $ 符号）
	for p.peekTokenIs(DOT) || p.peekTokenIs(DOLLAR) {
		separator := p.peek.Value // 保存分隔符（'.' 或 '$'）
		p.nextToken()             // consume separator
		if !p.expectPeek(IDENT) {
			return nil
		}
		className += separator + p.current.Value
	}

	if p.peekTokenIs(LPAREN) {
		// 普通构造器调用
		p.nextToken() // consume (
		p.nextToken() // 移动到 ( 后的第一个token
		arguments := p.parseArgumentList()

		// parseArgumentList 执行完后，current 应该在 RPAREN 上，需要消费它
		if p.current.Type == RPAREN {
			p.nextToken() // consume RPAREN
		}

		return &ConstructorExpression{
			ClassName: className,
			Arguments: arguments,
			IsArray:   false,
		}
	} else if p.peekTokenIs(LBRACK) {
		// 数组构造器: new Type[] { ... } 或 new Type[size]
		p.nextToken() // consume [
		p.nextToken() // move to content inside []

		if p.current.Type == RBRACK {
			// new Type[] { ... } 形式 - 数组初始化
			if !p.expectPeek(LBRACE) {
				return nil
			}
			// 现在 current = LBRACE

			elements := []Expression{}
			// 检查是否是空数组
			if !p.peekTokenIs(RBRACE) {
				p.nextToken() // move to first element
				elements = append(elements, p.parseAssignmentExpression())

				// 处理后续元素
				for p.current.Type == COMMA {
					p.nextToken() // consume comma, move to next element
					elements = append(elements, p.parseAssignmentExpression())
				}
				// 此时 current 应该在最后一个元素之后
			} else {
				// 空数组: current = LBRACE, peek = RBRACE
				// 需要移动到 RBRACE
				p.nextToken() // move from LBRACE to RBRACE
			}

			// 现在 current 应该在 RBRACE 上
			if p.current.Type != RBRACE {
				p.currentError(fmt.Sprintf("expected } after array elements, got %s", TokenTypeNames[p.current.Type]))
				return nil
			}
			p.nextToken() // consume RBRACE

			// 创建 ArrayExpression 包装元素列表
			arrayExpr := &ArrayExpression{Elements: elements}

			return &ConstructorExpression{
				ClassName: className,
				Arguments: []Expression{arrayExpr},
				IsArray:   true,
			}
		} else {
			// new Type[size] 形式 - 指定长度的数组
			// 解析大小表达式
			sizeExpr := p.parseAssignmentExpression()
			if sizeExpr == nil {
				return nil
			}

			// 期望 ]
			if p.current.Type != RBRACK {
				p.currentError(fmt.Sprintf("expected ] after array size, got %s", TokenTypeNames[p.current.Type]))
				return nil
			}
			p.nextToken() // consume ]

			// 对于 new Type[size] 形式，我们返回一个特殊的构造器表达式
			// Arguments 包含一个表示大小的表达式
			return &ConstructorExpression{
				ClassName: className,
				Arguments: []Expression{sizeExpr},
				IsArray:   true,
			}
		}
	}

	p.currentError("expected ( or [ after constructor class name")
	return nil
}

// parseStaticReference 解析静态引用 (@package.Class@member)
func (p *Parser) parseStaticReference() Expression {
	// 当前 token 是 @，移动到下一个 token
	p.nextToken()

	// 检查是否是 @@ 简写语法 (等价于 @java.lang.Math@)
	var className string
	if p.current.Type == AT {
		// @@ 简写，默认使用 java.lang.Math
		className = "java.lang.Math"
		p.nextToken() // consume 第二个 @，移动到方法名

		// 期望方法名
		if p.current.Type != IDENT {
			p.currentError("expected method name after @@")
			return nil
		}
		memberName := p.current.Value

		// @@ 后只能是方法调用，不能是字段访问
		if p.peek.Type != LPAREN {
			p.currentError("@@ can only be used with method calls, not field access")
			return nil
		}

		// 静态方法调用
		p.nextToken() // consume (，current 现在是 LPAREN
		p.nextToken() // 移动到 LPAREN 后的第一个token
		arguments := p.parseArgumentList()
		// parseArgumentList 返回时，当前应该在 RPAREN 上，需要消费它
		if p.current.Type == RPAREN {
			p.nextToken() // consume RPAREN，移动到下一个 token
		}
		result := &StaticMethodExpression{
			ClassName: className,
			Method:    memberName,
			Arguments: arguments,
		}
		return p.parseNavigationChainContinue(result)
	}

	// 标准 @ClassName@ 语法
	if p.current.Type != IDENT {
		p.currentError("expected class name after @")
		return nil
	}

	// 构建完整的类名 (package.Class 或 package.Class$InnerClass)
	className = p.current.Value

	// 处理包名中的点号和内部类的 $ 符号
	for p.peek.Type == DOT || p.peek.Type == DOLLAR {
		separator := p.peek.Value // '.' 或 '$'
		p.nextToken()             // consume . or $
		p.nextToken()             // move to next identifier
		if p.current.Type != IDENT {
			p.currentError(fmt.Sprintf("expected identifier after %s", separator))
			return nil
		}
		className += separator + p.current.Value
	}

	// 期望第二个 @
	if p.peek.Type != AT {
		p.currentError("expected @ after class name")
		return nil
	}
	p.nextToken() // consume @

	// 期望成员名称
	if p.peek.Type != IDENT {
		p.currentError("expected member name after @")
		return nil
	}
	p.nextToken() // move to member name
	memberName := p.current.Value

	// 检查是否是方法调用（有括号）
	if p.peek.Type == LPAREN {
		// 静态方法调用
		p.nextToken() // consume (，current 现在是 LPAREN
		p.nextToken() // 移动到 LPAREN 后的第一个token
		arguments := p.parseArgumentList()
		// parseArgumentList 返回时，当前应该在 RPAREN 上，需要消费它
		if p.current.Type == RPAREN {
			p.nextToken() // consume RPAREN，移动到下一个 token (可能是 . 或 EOF)
		}
		result := &StaticMethodExpression{
			ClassName: className,
			Method:    memberName,
			Arguments: arguments,
		}

		// 关键修复：继续处理可能的链式调用
		return p.parseNavigationChainContinue(result)
	} else {
		// 静态字段访问
		result := &StaticFieldExpression{
			ClassName: className,
			Field:     memberName,
		}
		p.nextToken() // consume the field name (move past it)

		// 关键修复：继续处理可能的链式调用
		return p.parseNavigationChainContinue(result)
	}
}

// parseLambdaExpression 解析 Lambda 表达式 :[expression]
// Lambda 表达式用于定义匿名函数
// 例如: :[#this * 2] 定义一个将参数乘以2的函数
// 例如: #fact=:[#this <= 1 ? 1 : #fact(#this-1) * #this] 定义递归阶乘函数
// 如果 Lambda 后紧跟 (arg)，则解析为 ASTEval 表达式
// 注意：根据Java OGNL的实现，Lambda表达式被包装在ASTConst节点中
func (p *Parser) parseLambdaExpression() Expression {
	// current 是 COLON
	if !p.expectPeek(LBRACK) {
		return nil
	}
	// current 现在是 LBRACK
	p.nextToken() // 移动到 [ 后的第一个 token

	// 解析 Lambda 函数体
	body := p.parseExpression()
	if body == nil {
		p.currentError("expected expression in lambda body")
		return nil
	}

	// 期望 ]
	if p.current.Type != RBRACK {
		p.currentError(fmt.Sprintf("expected ] after lambda body, got %s", TokenTypeNames[p.current.Type]))
		return nil
	}
	p.nextToken() // consume ]

	// 根据Java OGNL的实现，Lambda表达式应该包装在ASTConst中
	// 在Go中，我们使用LambdaLiteral来表示这个ASTConst节点
	lambdaConst := &LambdaLiteral{
		Body: body,
	}

	// 检查是否紧跟 (，如果是则解析为 ASTEval 表达式
	if p.current.Type == LPAREN {
		p.nextToken() // consume LPAREN
		// 解析参数
		argument := p.parseExpression()
		if argument == nil {
			p.currentError("expected argument in eval expression")
			return nil
		}
		// 期望 )
		if p.current.Type != RPAREN {
			p.currentError(fmt.Sprintf("expected ) after eval argument, got %s", TokenTypeNames[p.current.Type]))
			return nil
		}
		p.nextToken() // consume RPAREN

		// 创建 ASTEval 表达式
		evalExpr := &EvalExpression{
			Target:   lambdaConst,
			Argument: argument,
		}

		// 继续处理可能的链式调用，如 :[33](20).longValue()
		return p.parseNavigationChainContinue(evalExpr)
	}

	// 如果没有紧跟 (，继续处理可能的链式调用
	return p.parseNavigationChainContinue(lambdaConst)
}
