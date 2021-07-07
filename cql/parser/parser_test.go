package parser

import (
	"testing"
	"cql/lexer"
	"cql/ast"
	"cql/token"
)

func assertTokenData(tok token.Token, literal string, tokenType token.TokenType) bool {
	if literal != tok.Literal{ return false }
	if tokenType != tok.Type{ return false }

	return true
}

func TestCreateDirStatement(t *testing.T){
	input := `
		create dir test;
	`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()

	if len(program.Statements) != 1{ t.Errorf("Failed Test") }

	stmt := program.Statements[0].(*ast.CreateDirStatement)

	if "create dir test;" != stmt.String(){ t.Errorf("Failed Test") }

	// test create token
	var tok = stmt.Token
	if !assertTokenData(tok, "create", token.CREATE){ t.Errorf("Failed Test") }

	// test name token
	var name = stmt.Name
	tok = name.Token
	if !assertTokenData(tok, "test", token.IDENT){ t.Errorf("Failed Test") }
}

func TestCreateCsvStatement(t *testing.T){
	input := `
		create csv test (id, name);
	`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()

	if len(program.Statements) != 1{ t.Errorf("Failed Test") }

	stmt := program.Statements[0].(*ast.CreateCsvStatement)

	if "create csv test (id, name);" != stmt.String(){ t.Errorf("Failed Test") }
}