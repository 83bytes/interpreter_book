package ast

import (
	"bytes"
	"monkey/token"
)

// We are building an AST for the parser first
// The AST is a tree which is made up of nodes.
// This is a logical representation of the code.
// We are trying to parse the simplest statemets in Monkey first
// the let statement
// let x = 5 // here x is an identifier and 5 is an expression and the entire things is a statement

// Every node in our AST has to impement the Node interface
// it will provide the TokenLiteral method which returns the token that this node is associated with.
// this will only be used for debugging purposes
// The AST itself will be made by connecting nodes to nodes. (its a tree)
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

// This is the root node of every AST for MonkeyLang
// Every valid moneky program is a series of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// For Printing AST Nodes
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Now we will define how a node for variable binding looks like
// the node should have the following info
// the identifier
// the expression on the right hand side of the operator
// the token (let) in this case
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // pointer to an Identifier Node
	Value Expression  // Value of an Expression Node, as they eval into values
}

func (ls *LetStatement) statementNode()       {} // empty method only needed to satisfy the interface and make the compiler happy
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Now we define the Identifier Node
// Why is the Identifier satisfy the expressionNode interface ?
// to keep things and code simple
// we will use the identifier to represent the name in a variable binding
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}

// The AST for a simple Monkey program like
// let x = 5 would look like
// *ast.Program
//   - Statements
// 		- *ast.LetStatement
//        Name:
//			- *ast.Identifier
//        Value:
//          - *ast.Expression

// Parsing reurn statements

// the return statement is
// return <expression>
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// The Expression Statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Integer Literal

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
