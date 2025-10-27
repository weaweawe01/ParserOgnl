package lexer

import (
	"strconv"
	"strings"

	"github.com/weaweawe01/ParserOgnl/token"
)

// Lexer 词法分析器
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
}

// NewLexer 创建新的词法分析器
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

// readChar 读取下一个字符并前进指针
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++
}

// peekChar 查看下一个字符但不移动指针
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// skipWhitespace 跳过空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		} else {
			l.column++
		}
		l.readChar()
	}
}

// readString 读取字符串字面量（支持转义字符）
func (l *Lexer) readString() string {
	var result []byte
	l.readChar() // 跳过开始的 "

	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				result = append(result, '\n')
			case 't':
				result = append(result, '\t')
			case 'r':
				result = append(result, '\r')
			case 'b':
				result = append(result, '\b')
			case 'f':
				result = append(result, '\f')
			case '\\':
				result = append(result, '\\')
			case '\'':
				result = append(result, '\'')
			case '"':
				result = append(result, '"')
			case '0', '1', '2', '3', '4', '5', '6', '7':
				// 八进制转义序列
				octal := string(l.ch)
				if isOctalDigit(l.peekChar()) {
					l.readChar()
					octal += string(l.ch)
					if isOctalDigit(l.peekChar()) {
						l.readChar()
						octal += string(l.ch)
					}
				}
				if val, err := strconv.ParseInt(octal, 8, 32); err == nil {
					result = append(result, byte(val))
				}
			default:
				result = append(result, l.ch)
			}
			l.readChar() // 移动到下一个字符（转义序列之后）
		} else {
			result = append(result, l.ch)
			l.readChar() // 移动到下一个字符（普通字符之后）
		}
	}

	// 此时 l.ch 应该是结束的 " 或 0
	return string(result)
} // readCharLiteral 读取字符字面量（支持转义字符）
// readCharLiteral 读取单引号字符或字符串字面量
// 如果内容只有一个字符，返回字符；否则返回字符串
func (l *Lexer) readCharLiteral() string {
	l.readChar() // 跳过开始的 '

	var result strings.Builder

	for l.ch != '\'' && l.ch != 0 {
		if l.ch == '\\' {
			// 处理转义字符
			l.readChar()
			switch l.ch {
			case 'n':
				result.WriteByte('\n')
			case 't':
				result.WriteByte('\t')
			case 'r':
				result.WriteByte('\r')
			case 'b':
				result.WriteByte('\b')
			case 'f':
				result.WriteByte('\f')
			case '\\':
				result.WriteByte('\\')
			case '\'':
				result.WriteByte('\'')
			case '"':
				result.WriteByte('"')
			case '0', '1', '2', '3', '4', '5', '6', '7':
				// 八进制转义序列
				octal := string(l.ch)
				if isOctalDigit(l.peekChar()) {
					l.readChar()
					octal += string(l.ch)
					if isOctalDigit(l.peekChar()) {
						l.readChar()
						octal += string(l.ch)
					}
				}
				if val, err := strconv.ParseInt(octal, 8, 32); err == nil {
					result.WriteByte(byte(val))
				}
			default:
				result.WriteByte(l.ch)
			}
			l.readChar()
		} else {
			result.WriteByte(l.ch)
			l.readChar()
		}
	}

	return result.String()
}

// readBackCharLiteral 读取反引号字符字面量（不支持转义）
func (l *Lexer) readBackCharLiteral() string {
	l.readChar() // 跳过开始的 `

	if l.ch == 0 {
		return ""
	}

	result := l.ch
	l.readChar() // 读取字符

	// 跳过结束的 `
	for l.ch != '`' && l.ch != 0 {
		l.readChar()
	}

	return string(result)
}

// NextToken 扫描输入并返回下一个token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case 0:
		tok = token.Token{Type: token.EOF, Value: "", Line: l.line, Column: l.column}
		return tok // 不调用 readChar，因为已经在 EOF
	case '+':
		tok = token.Token{Type: token.PLUS, Value: "+", Line: l.line, Column: l.column}
	case '-':
		tok = token.Token{Type: token.MINUS, Value: "-", Line: l.line, Column: l.column}
	case '*':
		tok = token.Token{Type: token.MULTIPLY, Value: "*", Line: l.line, Column: l.column}
	case '/':
		tok = token.Token{Type: token.DIVIDE, Value: "/", Line: l.line, Column: l.column}
	case '%':
		tok = token.Token{Type: token.MODULO, Value: "%", Line: l.line, Column: l.column}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else {
			tok = token.Token{Type: token.ASSIGN, Value: "=", Line: l.line, Column: l.column}
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else {
			tok = token.Token{Type: token.NOT, Value: "!", Line: l.line, Column: l.column}
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LT_EQ, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.SHL, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else {
			tok = token.Token{Type: token.LT, Value: "<", Line: l.line, Column: l.column}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GT_EQ, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else if l.peekChar() == '>' {
			l.readChar()
			if l.peekChar() == '>' {
				l.readChar()
				tok = token.Token{Type: token.USHR, Value: ">>>", Line: l.line, Column: l.column}
			} else {
				tok = token.Token{Type: token.SHR, Value: ">>", Line: l.line, Column: l.column}
			}
		} else {
			tok = token.Token{Type: token.GT, Value: ">", Line: l.line, Column: l.column}
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else {
			tok = token.Token{Type: token.BIT_AND, Value: "&", Line: l.line, Column: l.column}
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Value: string(ch) + string(l.ch), Line: l.line, Column: l.column}
		} else {
			tok = token.Token{Type: token.BIT_OR, Value: "|", Line: l.line, Column: l.column}
		}
	case '^':
		tok = token.Token{Type: token.XOR, Value: "^", Line: l.line, Column: l.column}
	case '~':
		tok = token.Token{Type: token.BIT_NOT, Value: "~", Line: l.line, Column: l.column}
	case '(':
		tok = token.Token{Type: token.LPAREN, Value: "(", Line: l.line, Column: l.column}
	case ')':
		tok = token.Token{Type: token.RPAREN, Value: ")", Line: l.line, Column: l.column}
	case '[':
		tok = token.Token{Type: token.LBRACK, Value: "[", Line: l.line, Column: l.column}
	case ']':
		tok = token.Token{Type: token.RBRACK, Value: "]", Line: l.line, Column: l.column}
	case '{':
		tok = token.Token{Type: token.LBRACE, Value: "{", Line: l.line, Column: l.column}
	case '}':
		tok = token.Token{Type: token.RBRACE, Value: "}", Line: l.line, Column: l.column}
	case ',':
		tok = token.Token{Type: token.COMMA, Value: ",", Line: l.line, Column: l.column}
	case '.':
		// 检查 . 后面是否跟着数字，如果是则识别为浮点数
		if l.position+1 < len(l.input) && isDigit(l.input[l.position+1]) {
			// 这是一个以 . 开头的浮点数，如 .1234
			value, tokenType := l.readNumber()
			tok = token.Token{Type: tokenType, Value: value, Line: l.line, Column: l.column}
			return tok // 不调用 readChar，因为 readNumber 已经移动了指针
		}
		tok = token.Token{Type: token.DOT, Value: ".", Line: l.line, Column: l.column}
	case ':':
		tok = token.Token{Type: token.COLON, Value: ":", Line: l.line, Column: l.column}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Value: ";", Line: l.line, Column: l.column}
	case '?':
		tok = token.Token{Type: token.QUESTION, Value: "?", Line: l.line, Column: l.column}
	case '#':
		tok = token.Token{Type: token.HASH, Value: "#", Line: l.line, Column: l.column}
	case '$':
		tok = token.Token{Type: token.DOLLAR, Value: "$", Line: l.line, Column: l.column}
	case '@':
		tok = token.Token{Type: token.AT, Value: "@", Line: l.line, Column: l.column}
	case '"':
		tok.Type = token.STR_LITERAL
		tok.Value = l.readString()
		tok.Line = l.line
		tok.Column = l.column
		// readString 返回后，l.ch 应该指向结束的 "
		// 我们需要跳过它
		if l.ch == '"' {
			l.readChar() // 跳过结束的 "
		}
		return tok // 直接返回，不执行末尾的 readChar()
	case '\'':
		value := l.readCharLiteral()
		tok.Line = l.line
		tok.Column = l.column

		// 根据内容长度判断是字符还是字符串
		// 单个字符（包括转义字符）→ CHAR_LITERAL
		// 多个字符或空字符串 → STR_LITERAL (视为字符串)
		if len(value) == 1 {
			tok.Type = token.CHAR_LITERAL
			tok.Value = value
		} else {
			// 多字符或空字符串，视为字符串字面量
			tok.Type = token.STR_LITERAL
			tok.Value = value
		}

		l.readChar() // 跳过结束的 '
		return tok   // 直接返回，不执行末尾的 readChar()
	case '`':
		tok = token.Token{Type: token.BACK_CHAR_LITERAL, Value: l.readBackCharLiteral(), Line: l.line, Column: l.column}
		l.readChar() // 跳过结束的 `
		return tok   // 直接返回，不执行末尾的 readChar()
	default:
		if isLetter(l.ch) {
			value := l.readIdentifier()
			tokenType := token.LookupIdent(value)

			// 特殊处理 "not" 关键字
			// 如果是 "not" 并且后面跟着 "in"，则识别为 NOT_IN
			// 否则识别为 NOT（逻辑非运算符）
			if value == "not" {
				// 跳过空白字符
				l.skipWhitespace()

				// 保存当前位置（跳过空白字符后的位置）
				savedPos := l.position
				savedCh := l.ch
				savedLine := l.line
				savedCol := l.column

				// 检查是否后面跟着 "in"
				if isLetter(l.ch) {
					nextValue := l.readIdentifier()
					if nextValue == "in" {
						// 这是 "not in" 运算符
						tok = token.Token{Type: token.NOT_IN, Value: "not in", Line: l.line, Column: l.column}
						return tok
					}
					// 不是 "in"，需要回退到读取nextValue之前的位置
					l.position = savedPos
					l.ch = savedCh
					l.line = savedLine
					l.column = savedCol
				}
				// 如果后面不是字母，不需要回退

				// 单独的 "not" 作为逻辑非运算符
				tokenType = token.NOT
			}

			tok = token.Token{Type: tokenType, Value: value, Line: l.line, Column: l.column}
			return tok // 不调用 readChar，因为 readIdentifier 已经移动了指针
		} else if isDigit(l.ch) {
			value, tokenType := l.readNumber()
			tok = token.Token{Type: tokenType, Value: value, Line: l.line, Column: l.column}
			return tok // 不调用 readChar，因为 readNumber 已经移动了指针
		} else {
			tok = token.Token{Type: token.ILLEGAL, Value: string(l.ch), Line: l.line, Column: l.column}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	// 添加边界检查
	if position > len(l.input) {
		return ""
	}
	if l.position > len(l.input) {
		return l.input[position:]
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (string, token.TokenType) {
	position := l.position
	base := 10
	isFloat := false

	// 检查十六进制和八进制
	if l.ch == '0' {
		l.readChar()
		if l.ch == 'x' || l.ch == 'X' {
			// 十六进制
			base = 16
			l.readChar()
			for isHexDigit(l.ch) {
				l.readChar()
			}
		} else if isDigit(l.ch) {
			// 八进制
			base = 8
			for isOctalDigit(l.ch) {
				l.readChar()
			}
		}
	} else {
		// 十进制数字
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	// 检查小数部分（仅十进制）
	// 在 OGNL 中，5. 也是合法的浮点数
	if base == 10 && l.ch == '.' {
		nextCh := l.peekChar()
		// 如果小数点后面是数字，或者小数点后面不是字母（避免 5.toString() 被误识别）
		// 并且后面不是另一个小数点（避免范围运算符 .. 被误识别）
		if isDigit(nextCh) {
			// 标准小数：5.0, 5.123
			isFloat = true
			l.readChar() // 消费小数点
			for isDigit(l.ch) {
				l.readChar()
			}
		} else if isNumberSuffix(nextCh) {
			// 小数点后直接跟数字后缀：2.B, 5.d 等
			isFloat = true
			l.readChar() // 消费小数点
		} else if !isLetter(nextCh) && nextCh != '.' {
			// 尾随小数点：5. (等同于 5.0)
			// 但要确保后面不是标识符的开始，也不是另一个点（范围运算符）
			isFloat = true
			l.readChar() // 消费小数点
		}
	}

	// 检查指数部分（仅十进制）
	if base == 10 && (l.ch == 'e' || l.ch == 'E') {
		isFloat = true
		l.readChar()
		if l.ch == '+' || l.ch == '-' {
			l.readChar()
		}
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	// 检查后缀
	if base == 10 {
		if l.ch == 'd' || l.ch == 'D' || l.ch == 'f' || l.ch == 'F' || l.ch == 'b' || l.ch == 'B' {
			isFloat = true
			l.readChar()
		} else if l.ch == 'l' || l.ch == 'L' || l.ch == 'h' || l.ch == 'H' {
			l.readChar()
			return l.input[position:l.position], token.INT_LITERAL
		}
	}

	if isFloat {
		return l.input[position:l.position], token.FLT_LITERAL
	}
	return l.input[position:l.position], token.INT_LITERAL
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isOctalDigit(ch byte) bool {
	return '0' <= ch && ch <= '7'
}

func isHexDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

func isNumberSuffix(ch byte) bool {
	// 检查是否是数字后缀：d, D, f, F, b, B, l, L, h, H
	return ch == 'd' || ch == 'D' || ch == 'f' || ch == 'F' ||
		ch == 'b' || ch == 'B' || ch == 'l' || ch == 'L' ||
		ch == 'h' || ch == 'H'
}
