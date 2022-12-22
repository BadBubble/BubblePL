package lexer

import "BubblePL/token"

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
	var tk token.Token
	switch l.ch {
	case '+':
		tk = token.New(token.PLUS, l.ch)
	case '=':
		tk = token.New(token.ASSIGN, l.ch)
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
	case '0':
		tk = token.New(token.EOF, l.ch)
	default:
		tk = token.New(token.ILLEGAL, l.ch)
	}
	// 读取
	l.readChar()
	return tk
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
