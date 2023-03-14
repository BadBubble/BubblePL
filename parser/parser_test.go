package parser

import (
	"BubblePL/ast"
	"BubblePL/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x  5;
let = 10;
let 838 383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseError(t, p)
	if program == nil {
		t.Fatalf("Parse Program Error\n")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("get wrong statements, expected=%d, got=%d", 3, len(program.Statements))
	}
	tests := []struct {
		expectedNameStr string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for idx, test := range tests {
		testLetStatement(t, program.Statements[idx], test.expectedNameStr)
	}
}

func testLetStatement(t *testing.T, st ast.Statement, name string) {
	if st.ToLiteral() != "let" {
		t.Fatalf("Statement.ToLiteral error, expected=%s, got=%s", "let", st.ToLiteral())
	}

	letSt, ok := st.(*ast.LetStatement)
	if !ok {
		t.Fatalf("got wrong type, expected=*ast.LetStatement, got=%T", letSt)
	}

	if letSt.Name.ToLiteral() != name {
		t.Fatalf("get wrong name, expected='%s', got='%s'", name, letSt.Name.ToLiteral())
	}

}

func checkParseError(t *testing.T, p *Parser) {
	errs := p.Errors()
	if len(errs) == 0 {
		return
	}
	for _, errMsg := range errs {
		t.Errorf(errMsg)
	}
	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return 993 233;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseError(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStatement, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStatement.ToLiteral() != "return" {
			t.Errorf("returnStatement.ToLiteral not 'return'. got=%s", returnStatement.ToLiteral())
		}
	}
}
