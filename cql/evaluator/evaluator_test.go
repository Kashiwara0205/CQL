package evaluator

import (
	"testing"
	"cql/lexer"
	"cql/parser"
)

func TestEvalCreateDir(t *testing.T){
	input := `
		create dir test;
	`

	lex := lexer.New(input)
	p := parser.New(lex)
	program := p.ParseProgram()
	Eval(program)
}

func TestEvalCreateCsv(t *testing.T){
	input := `
		create csv test;
	`

	lex := lexer.New(input)
	p := parser.New(lex)
	program := p.ParseProgram()
	Eval(program)
}
