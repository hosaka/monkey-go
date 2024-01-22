package parser

import (
	"testing"

	"github.com/hosaka/monkey-go/ast"
	"github.com/hosaka/monkey-go/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
  let y = 10;
  let foo = 42;`

	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("wrong number of statements, expected=%d, got=%d", 3, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		t.Errorf("parser error: %q", err)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("wrong token literal, expected=%s, got=%q", "let", s.TokenLiteral())
		return false
	}

	statement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatement, got=%T", s)
		return false
	}

	if statement.Name.Value != name {
		t.Errorf("wrong statement name, expected=%s, got=%s", name, statement.Name.Value)
		return false
	}

	if statement.Name.TokenLiteral() != name {
		t.Errorf("wrong statement token, expected=%s, got=%s", name, statement.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `return 5;
  return 10;
  return 42;`

	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatal("Parse() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("wrong number of statements, expected=%d, got=%d", 3, len(program.Statements))
	}

	for _, s := range program.Statements {
		statement, ok := s.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s is not *ast.ReturnStatement, got=%T", s)
			continue
		}
		if statement.TokenLiteral() != "return" {
			t.Errorf("wrong statement token, expected=%s, got=%s", "return", statement.TokenLiteral())
		}
	}
}
