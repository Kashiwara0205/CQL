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

type UseStatement struct{
	Token token.Token
	DirName *Identifier
}
func (u *UseStatement) statementNode() {}

func (u *UseStatement) TokenLiteral() string {return u.Token.Literal}

func (u *UseStatement) String() string{
	var out bytes.Buffer

	out.WriteString(u.TokenLiteral() + " " + u.DirName.String() + ";")

	return out.String()
}

type FromStatement struct{
	Token token.Token
	Name *Identifier
}
func (c *FromStatement) statementNode() {}

func (c *FromStatement) TokenLiteral() string {return c.Token.Literal}

func (c *FromStatement) String() string{
	var out bytes.Buffer

	out.WriteString(c.TokenLiteral() + " " + c.Name.String())

	return out.String()
}

type SelectStatement struct{
	Token token.Token
	Columns []Expression // COUNTやMAXなども入れる予定なのでIdentifierでなく、Expressionにしている
}
func (s *SelectStatement) statementNode() {}

func (s *SelectStatement) TokenLiteral() string {return s.Token.Literal}

func (s *SelectStatement) String() string{
	var out bytes.Buffer

	args := []string{}
	for _, p := range s.Columns{
		args = append(args, p.String())
	}

	out.WriteString(s.TokenLiteral() + " "  + strings.Join(args, ", "))

	return out.String()
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

type BuitInFunction struct{
	Token token.Token
	Args []*Identifier
}
func (b *BuitInFunction) expressionNode() {}

func (b *BuitInFunction) TokenLiteral() string { return b.Token.Literal }

func (b *BuitInFunction) String() string{
	var out bytes.Buffer

	args := []string{}
	for _, p := range b.Args{
		args = append(args, p.String())
	}

	out.WriteString(b.TokenLiteral() + " ("  + strings.Join(args, ", ") + " )")

	return out.String()
}

type Identifier struct{
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {return i.Value}