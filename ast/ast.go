package ast

import "github.com/hosaka/monkey-go/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// the root of the Abstract Syntax Tree
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Identifier struct {
	Value string
	Token token.Token // token.IDENT
}

func (s *Identifier) expressionNode()      {}
func (s *Identifier) TokenLiteral() string { return s.Token.Literal }

type LetStatement struct {
	Value Expression
	Name  *Identifier
	Token token.Token // token.LET
}

func (s *LetStatement) statementNode()       {}
func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }

type ReturnStatement struct {
	Value Expression
	Token token.Token // token.RETURN
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
