package token

type TokenType string

const (
	/*特殊类*/
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	/*标识符*/
	INT   = "INT"
	IDENT = "IDENT"
	/*运算符*/
	ASSIGN = "="
	PLUS   = "+"
	/*符号*/
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	/*关键字*/
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// KeywordsMap 关键字的Literal到TokenType的映射
var KeywordsMap = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Token 通过lexer将代码转换成一个一个的Token
type Token struct {
	Type    TokenType // Token的类别
	Literal string    // Token的字面量
}

// LookupIdentifier 根据字面量查找是否是关键字，如果是关键字就返回关键字的TokenType否则就是IDENT
func LookupIdentifier(literal string) TokenType {
	if tokenType, ok := KeywordsMap[literal]; ok {
		return tokenType
	}
	return IDENT
}

// New 创建Token
func New(tokenType TokenType, literal byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}
