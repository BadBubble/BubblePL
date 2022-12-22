package lexer

import (
	"BubblePL/token"
	"testing"
)

// TestNextToken 测试lexer.NextToken功能
func TestNextToken(t *testing.T) {
	input := `+={}(),;`

	tests := []struct {
		ExpectedTokenType token.TokenType
		ExpectedLiteral   string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
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
