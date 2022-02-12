package ast

import (
	"bytes"
	"fmt"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of every AST the parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for i, stmt := range p.Statements {
		out.WriteString(stmt.String())
		if len(p.Statements) > 1 && i != len(p.Statements)-1 {
			out.WriteByte('\n')
		}
	}
	return out.String()
}

// Identifier ---------------------------------
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral ------------------------------
type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int
}

func (in *IntegerLiteral) expressionNode() {}
func (in *IntegerLiteral) TokenLiteral() string {
	return in.Token.Literal
}
func (in *IntegerLiteral) String() string {
	return in.TokenLiteral()
}

// LetStatement ---------------------------------
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal // let
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(
		fmt.Sprintf("%s %s = ", ls.TokenLiteral(), ls.Name.String()))

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteByte(';')

	return out.String()
}

// ReturnStatement ---------------------------------
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal // return
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteByte(';')
	return out.String()
}

// ExpressionStatement -------------------------------
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
