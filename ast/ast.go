package ast

import "BubblePL/token"

type Node interface {
	ToLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) ToLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].ToLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) ToLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
}

type Identifier struct {
	token.Token
	Value string
}

func (i *Identifier) ToLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {

}
