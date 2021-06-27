package lexer

type Lexer struct{
	input string
	position int
	nextPosition int
	ch byte
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

func New(input string) *Lexer{
	lex := &Lexer{input: input}
	lex.readChar()

	return lex
}