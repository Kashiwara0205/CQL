package lexer

import (
	"testing"
	"cql/token"
)

func assertLexData(lex *Lexer, input string, position int, nextPosition int, ch byte) bool {
	if input != lex.input{ return false }
	if position != lex.position{ return false }
	if nextPosition != lex.nextPosition{ return false }
	if ch != lex.ch{ return false }

	return true
}

func assertTokenData(tok token.Token, literal string, tokenType token.TokenType) bool {
	if literal != tok.Literal{ return false }
	if tokenType != tok.Type{ return false }

	return true
}

func TestCreateToken(t *testing.T){
	lex := New("create")
	tok := lex.NextToken()
	if !assertLexData(lex, "create", 6, 7, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, "create", token.CREATE){ t.Errorf("Failed Test") }
}

func TestDirToken(t *testing.T){
	lex := New("dir")
	tok := lex.NextToken()
	if !assertLexData(lex, "dir", 3, 4, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, "dir", token.DIR){ t.Errorf("Failed Test") }
}

func TestCsvToken(t *testing.T){
	lex := New("csv")
	tok := lex.NextToken()
	if !assertLexData(lex, "csv", 3, 4, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, "csv", token.CSV){ t.Errorf("Failed Test") }
}

func TestSemicolonToken(t *testing.T){
	lex := New(";")
	tok := lex.NextToken()
	if !assertLexData(lex, ";", 1, 2, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, ";", token.SEMICOLON){ t.Errorf("Failed Test") }
}

func TestCommaToken(t *testing.T){
	lex := New(",")
	tok := lex.NextToken()
	if !assertLexData(lex, ",", 1, 2, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, ",", token.COMMA){ t.Errorf("Failed Test") }
}

func TestRparenToken(t *testing.T){
	lex := New(")")
	tok := lex.NextToken()
	if !assertLexData(lex, ")", 1, 2, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, ")", token.RPAREN){ t.Errorf("Failed Test") }
}

func TestLparenToken(t *testing.T){
	lex := New("(")
	tok := lex.NextToken()
	if !assertLexData(lex, "(", 1, 2, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, "(", token.LPAREN){ t.Errorf("Failed Test") }
}

func TestLiteralToken(t *testing.T){
	lex := New("apple")
	tok := lex.NextToken()
	if !assertLexData(lex, "apple", 5, 6, 0){ t.Errorf("Failed Test") }
	if !assertTokenData(tok, "apple", token.IDENT){ t.Errorf("Failed Test") }
}

func TestskipWhitespace(t *testing.T){
	lex := New("create dir test;")

	var tok = lex.NextToken()
	if !assertTokenData(tok, "create", token.CREATE){ t.Errorf("Failed Test") }

	tok = lex.NextToken()
	if !assertTokenData(tok, "dir", token.DIR){ t.Errorf("Failed Test") }

	tok = lex.NextToken()
	if !assertTokenData(tok, "test", token.IDENT){ t.Errorf("Failed Test") }

	tok = lex.NextToken()
	if !assertTokenData(tok, ";", token.SEMICOLON){ t.Errorf("Failed Test") }
}

func TestNewLexer(t *testing.T){
	lex := New("CREATE DIR test;")

	if !assertLexData(lex, "CREATE DIR test;", 0, 1, 'C'){ t.Errorf("Failed Test") }

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

	if !assertLexData(lex, "CREATE;", 0, 1, 'C'){ t.Errorf("Failed Test") }
	
	lex.readChar()
	if !assertLexData(lex, "CREATE;", 1, 2, 'R'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 2, 3, 'E'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 3, 4, 'A'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 4, 5, 'T'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 5, 6, 'E'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 6, 7, ';'){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 7, 8, 0){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 8, 9, 0){ t.Errorf("Failed Test") }

	lex.readChar()
	if !assertLexData(lex, "CREATE;", 9, 10, 0){ t.Errorf("Failed Test") }
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