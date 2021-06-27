package token 

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const(
	// QUERY OPERATOR
	CREATE = "CREATE"
	CSV = "CSV"
	DIR = "DIR"

	IDENT = "IDENT"

	LPAREN = "("
	RPAREN = ")"

	SEMICOLON = ";"
)