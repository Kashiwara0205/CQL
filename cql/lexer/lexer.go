package lexer

import(
	"cql/token"
)

type Lexer struct{
	input string
	position int
	nextPosition int
	ch byte
}

func New(input string) *Lexer{
	lex := &Lexer{input: input}
	lex.readChar()

	return lex
}

func (lex *Lexer) isEnd() bool{
	return lex.nextPosition >= len(lex.input)
}

func (lex *Lexer) getNextChar() byte{
	return lex.input[lex.nextPosition]
}

func(lex *Lexer) updatePosition(){
	lex.position = lex.nextPosition
	lex.nextPosition += 1
}

func (lex *Lexer) readChar(){
	if lex.isEnd() {
		lex.ch = 0
	} else {
		lex.ch = lex.getNextChar()
	}

	lex.updatePosition()
}

func (lex *Lexer) NextToken() token.Token{
	var tok token.Token

	switch lex.ch{
	case '(':
		tok = token.New(token.LPAREN, string(lex.ch))
	case ')':
		tok = token.New(token.RPAREN, string(lex.ch))
	case 0:
		tok = token.New(token.EOF, "")
	default:
		tok = token.New(token.ILLEGAL, "")
	}

	lex.readChar()

	return tok
}
