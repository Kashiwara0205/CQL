package token 

import(
	"fmt"
)

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const(
	EOF = "EOF"
	ILLEGAL = "ILLEGAL"
	
	// QUERY OPERATOR
	CREATE = "CREATE"
	CSV = "CSV"
	DIR = "DIR"

	IDENT = "IDENT"

	LPAREN = "("
	RPAREN = ")"
	
	COMMA = ","
	SEMICOLON = ";"
)

var keywords = map[string] TokenType{
	"create": CREATE,
	"csv": CSV,
	"dir": DIR,
}

func New(tokenType TokenType, ch string) Token{
	return Token{ Type: tokenType, Literal: ch }
}

func (tok *Token) Puts(){
	fmt.Printf("type: %v\n", tok.Type)
	fmt.Printf("literal: %v\n", tok.Literal)
} 

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}

	return IDENT
}