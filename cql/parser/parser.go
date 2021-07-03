package paser

import (
	"cql/token"
	"cql/lexer"
	"cql/ast"
	"fmt"
)

const (
	_int = iota     
	LOWSET
	CALL
)

type(
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

var precedences = map[token.TokenType]int{
	token.LPAREN: CALL,
}

type Parser struct{
	lex *lexer.Lexer
	curToken token.Token
	peekToken token.Token
	errors []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

func New(lex *lexer.Lexer) *Parser{
	p := &Parser{
		lex: lex,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	// 1回目: curToken =  null → curToken = 値１
	// 2回目: peekToken = 値１ → peekToken = 値2
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool{
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekToken(t token.TokenType) bool{
	if p.peekTokenIs(t){
		p.nextToken()
		return true
	}else{
		p.peekError(t)
		return false
	}
}

func (p *Parser) expectTargetToken() bool{
	if p.peekTokenIs(token.DIR) || p.peekTokenIs(token.CSV){
		p.nextToken()
		return true
	}else{
		p.peekTargetError()
		return false
	}
}


func (p *Parser) peekPrecedence() int{
	if p, ok := precedences[p.peekToken.Type]; ok{
		return p
	}

	return LOWSET
}

func (p *Parser) ParseProgram() *ast.Program{
	program := &ast.Program{}

	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF{

		stmt := p.parseStatement()
		if stmt != nil{
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement{
	switch p.curToken.Type{
	case token.CREATE:
		return p.pareseCreateStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) pareseCreateStatement() *ast.CreateStatement{
	stmt := &ast.CreateStatement{ Token: p.curToken }

	// DIR or CSV
	if !p.expectTargetToken(){
		return nil
	}

	stmt.Target = &ast.Target{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeekToken(token.IDENT){
		return nil
	}

	// name
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn){
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) parseIdentifier() ast.Expression{
	return &ast.Identifier{ Token: p.curToken, Value: p.curToken.Literal }
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement{

	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWSET)

	if p.peekTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression{

	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil{
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence(){

		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil{
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) peekError(t token.TokenType){
	msg := fmt.Sprintf("expected next token no be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekTargetError(){
	msg := fmt.Sprintf("expected target token no")
	p.errors = append(p.errors, msg)
}


func (p *Parser) noPrefixParseFnError(t token.TokenType){
	msg := fmt.Sprintf("no prefix parse fuction for %s found", t)
	p.errors = append(p.errors, msg)
}
