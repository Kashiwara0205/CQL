package token 

import (
	"testing"
)

func TestNewToken(t *testing.T){
	var token = New(SEMICOLON, ";")

	if token.Type != SEMICOLON{
		t.Errorf("Failed Test")
	}

	if token.Literal != ";"{
		t.Errorf("Failed Test")
	}

	token = New(IDENT, "test")

	if token.Type != IDENT{
		t.Errorf("Failed Test")
	}

	if token.Literal != "test"{
		t.Errorf("Failed Test")
	}
}