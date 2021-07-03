package ast

import (
	"cql/token"
	"bytes"
)

type Program struct{
	Statements []Statement
}

type Node interface{
	TokenLiteral() string
	String() string
}

type Statement interface{
	Node
	statementNode()
}

type Expression interface{
	Node
	expressionNode()
}

type ExpressionStatement struct{
	Token token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Literal }

func (e *ExpressionStatement) String() string{
	if e.Expression != nil{
		return e.Expression.String()
	}
	return ""
}

type CreateStatement struct{
	Token token.Token
	Target *Target
	Name *Identifier
}
func (c *CreateStatement) statementNode() {}

func (c *CreateStatement) TokenLiteral() string {return c.Token.Literal}

func (c *CreateStatement) String() string{
	var out bytes.Buffer

	out.WriteString(c.TokenLiteral() + " " + c.Target.String() + " " + c.Name.String() + ";")

	return out.String()
}

type Target struct{
	Token token.Token
	Value string
}

func (t *Target) expressionNode() {}
func (t *Target) TokenLiteral() string { return t.Token.Literal }
func (t *Target) String() string {return t.Value}

type Identifier struct{
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {return i.Value}