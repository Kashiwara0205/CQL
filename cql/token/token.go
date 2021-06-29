package token 

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

	SEMICOLON = ";"
)

func New(tokenType TokenType, ch string) Token{
	return Token{ Type: tokenType, Literal: ch }
}