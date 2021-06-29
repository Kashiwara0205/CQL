package lexer

import (
	"testing"
	"cql/token"
)

func checkLexData(lex *Lexer, input string, position int, nextPosition int, ch byte) bool {
	if input != lex.input{ return false }

	if position != lex.position{ return false }

	if nextPosition != lex.nextPosition{ return false }

	if ch != lex.ch{ return false }

	return true
}

func TestNewLexer(t *testing.T){
	lex := New("CREATE DIR test;")

	if !checkLexData(lex, "CREATE DIR test;", 0, 1, 'C'){ t.Errorf("Failed Test") }
}

func TestisEnd(t *testing.T){
	// 読みとり文字が存在したらfalse
	var lex = New("C")

	if !lex.isEnd(){
		t.Errorf("Failed Test")
	}

	// 読み取り文字がなかったらtrue
	lex = New("")

	if lex.isEnd(){
		t.Errorf("Failed Test")
	}
}

func TestGetNextChar(t *testing.T){
	lex := New("CREATE DIR test;")

	if 'C' != lex.ch {
		t.Errorf("Failed Test")
	}

	if 'R' != lex.getNextChar(){
		t.Errorf("Failed Test")
	}
}

func TestUpdatePosition(t *testing.T){
	lex := New("CREATE DIR test;")

	if 0 != lex.position{
		t.Errorf("Failed Test")
	}

	if 1 != lex.nextPosition{
		t.Errorf("Failed Test")
	}

	lex.updatePosition()

	if 1 != lex.position{
		t.Errorf("Failed Test")
	}

	if 2 != lex.nextPosition{
		t.Errorf("Failed Test")
	}
}
func TestReadChar(t *testing.T){
	lex := New("CREATE;")

	if !checkLexData(lex, "CREATE;", 0, 1, 'C'){ t.Errorf("Failed Test") }
	
	lex.readChar()
	if !checkLexData(lex, "CREATE;", 1, 2, 'R'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 2, 3, 'E'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 3, 4, 'A'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 4, 5, 'T'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 5, 6, 'E'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 6, 7, ';'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 7, 8, 0){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 8, 9, 0){ t.Errorf("Failed Test") }

	lex.readChar()
	if !checkLexData(lex, "CREATE;", 9, 10, 0){ t.Errorf("Failed Test") }
}

func TestNextToken(t *testing.T){
	var lex = New("()")

	var tok = lex.NextToken()

	if tok.Type != token.LPAREN{
		t.Errorf("Failed Test")
	}

	if tok.Literal != "(" {
		t.Errorf("Failed Test")
	}

	tok = lex.NextToken()

	if tok.Type != token.RPAREN{
		t.Errorf("Failed Test")
	}

	if tok.Literal != ")" {
		t.Errorf("Failed Test")
	}

	tok = lex.NextToken()

	if tok.Type != token.EOF{
		t.Errorf("Failed Test")
	}

	if tok.Literal != "" {
		t.Errorf("Failed Test")
	}

	tok = lex.NextToken()

	if tok.Type != token.EOF{
		t.Errorf("Failed Test")
	}

	if tok.Literal != "" {
		t.Errorf("Failed Test")
	}
}