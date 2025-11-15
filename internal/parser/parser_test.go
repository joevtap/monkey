package parser

import (
	"testing"

	"github.com/joevtap/monkey/internal/ast"
	"github.com/joevtap/monkey/internal/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 10;
let y = 5;
let z = 100;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	statements := program.Statements
	if len(statements) != 3 {
		t.Fatalf("program.Statements: want %v got %v statements", 3, len(statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `return 5;
return 3;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements: want %v statements got %v", 2, len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not *ast.ReturnsTatement, got: %T", statement)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral: expected \"return\" token literal got %v", returnStatement.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral(): want %v got %v", "let", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("want %T got %T", ast.LetStatement{}, s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("s.Name.Value: want %v got %v", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name.TokenLiteral(): want %v got %v", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := len(p.Errors())
	if errors > 0 {
		t.Errorf("parser has %v errors", errors)
		for _, err := range p.Errors() {
			t.Log(err)
		}

		t.FailNow()
	}
}
