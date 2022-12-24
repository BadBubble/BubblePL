package lexer

import (
	"BubblePL/token"
	"testing"
)

// TestNextToken 测试lexer.NextToken功能
func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);`

	tests := []struct {
		ExpectedTokenType token.TokenType
		ExpectedLiteral   string
	}{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// let add = fn(a, y) {
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		// x + y;
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		// };
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		// EOF
		{token.EOF, ""},
	}
	l := New(input)

	for idx, test := range tests {
		tk := l.NextToken()
		if tk.Type != test.ExpectedTokenType {
			t.Fatalf("tests[%d] - TokenType error. expected=%q, but got=%q",
				idx, test.ExpectedTokenType, tk.Type)
		}

		if tk.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - Literal error. expected=%q, but got=%q",
				idx, test.ExpectedLiteral, tk.Literal)
		}
	}
}
