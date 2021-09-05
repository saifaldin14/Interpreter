package ast

// import (
// 	"bytes"
// 	"Interpreter/src/saif/token"
// 	"strings"
// )

//The Base Node Interface
type Node interface {
	TokenLiteral() string
	String() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}
