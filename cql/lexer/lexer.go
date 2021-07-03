package lexer

import(
	"cql/token"
	"fmt"
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

func (lex *Lexer) Puts(){
	fmt.Printf("input: %v\n", lex.input)
	fmt.Printf("position: %v\n", lex.position)
	fmt.Printf("nextPosition: %v\n", lex.nextPosition)
	fmt.Printf("ch: %v\n", lex.ch)
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

	lex.skipWhitespace()

	switch lex.ch{
	case '(':
		tok = token.New(token.LPAREN, string(lex.ch))
	case ')':
		tok = token.New(token.RPAREN, string(lex.ch))
	case 0:
		tok = token.New(token.EOF, "")
	case ';':
		tok = token.New(token.SEMICOLON, string(lex.ch))
	default:

		if isLetter(lex.ch) {
			literal := lex.readIdentifiter()
			tokType := token.LookupIdent(literal)

			return token.New(tokType, literal)
		}


		tok = token.New(token.ILLEGAL, "")
	}

	lex.readChar()

	return tok
}

func isLetter(ch byte) bool{
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lex *Lexer) readIdentifiter() string{
	position := lex.position

	for isLetter(lex.ch){
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) skipWhitespace(){
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r'{
		lex.readChar()
	}
}