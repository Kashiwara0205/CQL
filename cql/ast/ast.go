package ast

import (
	"cql/token"
	"bytes"
	"strings"
)

type Program struct{
	Statements []Statement
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}

func (p *Program) String() string{
	var out bytes.Buffer

	for _, s := range p.Statements{
		out.WriteString(s.String())
	}

	return out.String()
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

type CreateDirStatement struct{
	Token token.Token
	Name *Identifier
}
func (c *CreateDirStatement) statementNode() {}

func (c *CreateDirStatement) TokenLiteral() string {return c.Token.Literal}

func (c *CreateDirStatement) String() string{
	var out bytes.Buffer

	out.WriteString(c.TokenLiteral() + " " + "dir" + " " + c.Name.String() + ";")

	return out.String()
}

type CreateCsvStatement struct{
	Token token.Token
	Name *Identifier
	Columns []*Identifier
}
func (c *CreateCsvStatement) statementNode() {}

func (c *CreateCsvStatement) TokenLiteral() string {return c.Token.Literal}

func (c *CreateCsvStatement) String() string{
	var out bytes.Buffer

	columns := []string{}
	for _, p := range c.Columns{
		columns = append(columns, p.String())
	}

	out.WriteString(c.TokenLiteral() + " " + "csv" + " " + c.Name.String()  + " (" + strings.Join(columns, ", ") + ")" + ";")

	return out.String()
}

type Identifier struct{
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {return i.Value}