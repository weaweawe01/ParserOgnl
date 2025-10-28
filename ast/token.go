package ast

import "fmt"

// TokenType 表示Token的类型
type TokenType int

const (
	// 特殊Token
	ILLEGAL TokenType = iota
	EOF
	WHITESPACE

	// 标识符和字面量
	IDENT             // 标识符: abc, foo, bar, etc.
	INT_LITERAL       // 整数字面量: 123, 0x1F, 077
	FLT_LITERAL       // 浮点数字面量: 123.45, 1.23e10
	CHAR_LITERAL      // 字符字面量: 'a', '\n'
	STR_LITERAL       // 字符串字面量: "hello"
	BACK_CHAR_LITERAL // 反引号字符: `a`

	// 运算符
	ASSIGN    // =
	COMMA     // ,
	SEMICOLON // ;
	QUESTION  // ?
	COLON     // :

	// 逻辑运算符
	OR      // ||, or
	AND     // &&, and
	BIT_OR  // |, bor
	XOR     // ^, xor
	BIT_AND // &, band

	// 比较运算符
	EQ     // ==, eq
	NOT_EQ // !=, neq
	LT     // <, lt
	GT     // >, gt
	LT_EQ  // <=, lte
	GT_EQ  // >=, gte
	IN     // in
	NOT_IN // not in

	// 位移运算符
	SHL  // <<, shl
	SHR  // >>, shr
	USHR // >>>, ushr

	// 算术运算符
	PLUS     // +
	MINUS    // -
	MULTIPLY // *
	DIVIDE   // /
	MODULO   // %
	BIT_NOT  // ~
	NOT      // !

	// 特殊运算符
	INSTANCEOF // instanceof
	DOT        // .

	// 分隔符
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }
	LBRACK // [
	RBRACK // ]

	// 关键字
	TRUE  // true
	FALSE // false
	NULL  // null
	NEW   // new
	THIS  // #this
	ROOT  // #root
	HASH  // #

	// 特殊字符
	AT     // @
	DOLLAR // $

	// 动态下标
	DYNAMIC_SUBSCRIPT // [^], [|], [$], [*]
)

// Token 表示一个词法单元
type Token struct {
	Type     TokenType
	Value    string
	Literal  interface{} // 存储解析后的字面量值
	Line     int
	Column   int
	Position int
}

// TokenTypeNames Token类型名称映射
var TokenTypeNames = map[TokenType]string{
	ILLEGAL:           "ILLEGAL",
	EOF:               "EOF",
	WHITESPACE:        "WHITESPACE",
	IDENT:             "IDENT",
	INT_LITERAL:       "INT_LITERAL",
	FLT_LITERAL:       "FLT_LITERAL",
	CHAR_LITERAL:      "CHAR_LITERAL",
	STR_LITERAL:       "STR_LITERAL",
	BACK_CHAR_LITERAL: "BACK_CHAR_LITERAL",
	ASSIGN:            "ASSIGN",
	COMMA:             "COMMA",
	SEMICOLON:         "SEMICOLON",
	QUESTION:          "QUESTION",
	COLON:             "COLON",
	OR:                "OR",
	AND:               "AND",
	BIT_OR:            "BIT_OR",
	XOR:               "XOR",
	BIT_AND:           "BIT_AND",
	EQ:                "EQ",
	NOT_EQ:            "NOT_EQ",
	LT:                "LT",
	GT:                "GT",
	LT_EQ:             "LT_EQ",
	GT_EQ:             "GT_EQ",
	IN:                "IN",
	NOT_IN:            "NOT_IN",
	SHL:               "SHL",
	SHR:               "SHR",
	USHR:              "USHR",
	PLUS:              "PLUS",
	MINUS:             "MINUS",
	MULTIPLY:          "MULTIPLY",
	DIVIDE:            "DIVIDE",
	MODULO:            "MODULO",
	BIT_NOT:           "BIT_NOT",
	NOT:               "NOT",
	INSTANCEOF:        "INSTANCEOF",
	DOT:               "DOT",
	LPAREN:            "LPAREN",
	RPAREN:            "RPAREN",
	LBRACE:            "LBRACE",
	RBRACE:            "RBRACE",
	LBRACK:            "LBRACK",
	RBRACK:            "RBRACK",
	TRUE:              "TRUE",
	FALSE:             "FALSE",
	NULL:              "NULL",
	NEW:               "NEW",
	THIS:              "THIS",
	ROOT:              "ROOT",
	HASH:              "HASH",
	AT:                "AT",
	DOLLAR:            "DOLLAR",
	DYNAMIC_SUBSCRIPT: "DYNAMIC_SUBSCRIPT",
}

// String 返回Token的字符串表示
func (t Token) String() string {
	typeName, ok := TokenTypeNames[t.Type]
	if !ok {
		typeName = "UNKNOWN"
	}
	return fmt.Sprintf("Token{Type: %s, Value: %s, Line: %d, Column: %d}",
		typeName, t.Value, t.Line, t.Column)
}

// Keywords 关键字映射
var Keywords = map[string]TokenType{
	"and":        AND,
	"or":         OR,
	"bor":        BIT_OR,
	"xor":        XOR,
	"band":       BIT_AND,
	"eq":         EQ,
	"neq":        NOT_EQ,
	"lt":         LT,
	"gt":         GT,
	"lte":        LT_EQ,
	"gte":        GT_EQ,
	"in":         IN,
	"not":        NOT, // 单独的 not 是逻辑非运算符，not in 在 lexer 中特殊处理
	"shl":        SHL,
	"shr":        SHR,
	"ushr":       USHR,
	"instanceof": INSTANCEOF,
	"true":       TRUE,
	"false":      FALSE,
	"null":       NULL,
	"new":        NEW,
}

// LookupIdent 检查标识符是否为关键字
func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// DynamicSubscriptType 动态下标类型
type DynamicSubscriptType int

const (
	FIRST DynamicSubscriptType = iota // [^] - 第一个元素
	MID                               // [|] - 中间元素
	LAST                              // [$] - 最后一个元素
	ALL                               // [*] - 所有元素
)

// DynamicSubscriptNames 动态下标名称映射
var DynamicSubscriptNames = map[DynamicSubscriptType]string{
	FIRST: "FIRST",
	MID:   "MID",
	LAST:  "LAST",
	ALL:   "ALL",
}
