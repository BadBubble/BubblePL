package lexer

import (
	"BubblePL/token"
)

// Lexer 负责将源代码转换成Tokens
type Lexer struct {
	input        string // 程序字符串
	position     int    // 当前处理的字符在字符串中的位置
	readPosition int    // 下一个要读取的字符串的位置 position + 1
	ch           byte   // 当前处理的字符串
}

// readChar 读取下一个字符
func (l *Lexer) readChar() {
	// 如果读取完成，就把字符串赋值为\0代表字符串结束
	if l.readPosition >= len(l.input) {
		// ASCII 0 = NUL
		l.ch = 0
	} else {
		// 处理下一个字符
		l.ch = l.input[l.readPosition]
	}
	// 更新位置信息
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken 生成源代码的下一个Token
func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()
	var tk token.Token
	switch l.ch {
	case '+':
		tk = token.New(token.PLUS, l.ch)
	case '=':
		if l.peekChar() == '=' {
			tk = token.Token{Type: token.EQ, Literal: string(l.ch) + string(l.peekChar())}
			l.readChar()
		} else {
			tk = token.New(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tk = token.Token{Type: token.NOT_EQ, Literal: string(l.ch) + string(l.peekChar())}
			l.readChar()
		} else {
			tk = token.New(token.BAND, l.ch)
		}
	case ';':
		tk = token.New(token.SEMICOLON, l.ch)
	case ',':
		tk = token.New(token.COMMA, l.ch)
	case '{':
		tk = token.New(token.LBRACE, l.ch)
	case '}':
		tk = token.New(token.RBRACE, l.ch)
	case '(':
		tk = token.New(token.LPAREN, l.ch)
	case ')':
		tk = token.New(token.RPAREN, l.ch)
	case '/':
		tk = token.New(token.SLASH, l.ch)
	case '*':
		tk = token.New(token.ASTERISK, l.ch)
	case '-':
		tk = token.New(token.MINUS, l.ch)
	case '<':
		tk = token.New(token.LT, l.ch)
	case '>':
		tk = token.New(token.GT, l.ch)
	case '[':
		tk = token.New(token.LBRACKET, l.ch)
	case ']':
		tk = token.New(token.RBRACKET, l.ch)
	case 0:
		tk.Type = token.EOF
		tk.Literal = ""
	case '"':
		tk.Type = token.STRING
		tk.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookupIdentifier(tk.Literal)
			// 因为在Lexer.readIdentifier中已经调用Lexer.readChar将position的位置移动到了当前identifier后第一个位置，这里直接返回
			return tk
		} else if isNumber(l.ch) {
			tk.Literal = l.readNumber()
			tk.Type = token.INT
			return tk
		} else {
			tk = token.New(token.ILLEGAL, l.ch)
		}
	}
	// 读取
	l.readChar()
	return tk
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}

// isLetter 检测byte是否是字母
func isLetter(ch byte) bool {
	if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' {
		return true
	}
	return false
}

// isNumber 检查byte是否为数字
func isNumber(ch byte) bool {
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}

// readIdentifier 读取标识符的字面量
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber 读取数字的字面量
func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// eatWhitespace 去掉无意义的符号
func (l *Lexer) eatWhitespace() {
	for l.ch == '\t' || l.ch == ' ' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// peekChar获取readPosition的字符
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}
	// 初始化position=0, readPosition=1, ch=input的第一个字符
	l.readChar()
	return l
}
