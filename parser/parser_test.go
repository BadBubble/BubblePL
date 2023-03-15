package parser

import (
	"BubblePL/ast"
	"BubblePL/lexer"
	"fmt"
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

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`
	l := lexer.New(input)
	parser := New(l)
	program := parser.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements got wrong length got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("Statement is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Errorf("Statement.Expression is not ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("Identifier.Value is not 'foobar', got=%s", ident.Value)
	}

	if ident.ToLiteral() != "foobar" {
		t.Errorf("Identifier.ToLiteral is not 'foobar', got=%s", ident.Value)
	}
}

func TestIntegerExpression(t *testing.T) {
	input := `5;`
	l := lexer.New(input)
	parser := New(l)
	program := parser.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements got wrong length got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("Statement is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("Statement.Expression is not ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if ident.Value != 5 {
		t.Errorf("IntegerLiteral.Value is not 5, got=%d", ident.Value)
	}

	if ident.ToLiteral() != "5" {
		t.Errorf("Identifier.ToLiteral is not '5', got=%d", ident.Value)
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	exp, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("Experssion is not ast.IntegerLiteral. got=%T", il)
		return false
	}

	if exp.Value != value {
		t.Errorf("ast.IntegerLiteral.Value is not %d. got=%d", value, exp.Value)
		return false
	}
	if exp.ToLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("ast.IntegerLiteral.ToLiteral is not '%d'. got=%s", value, exp.ToLiteral())
		return false
	}
	return true
}

func TestPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		parser := New(l)
		program := parser.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements got wrong length got=%d", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Errorf("stmt.Expression is not *ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Errorf("exp.Operator is not '%s' *ast.PrefixExpression. got='%s'", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}
