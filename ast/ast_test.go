package ast

import (
	"github.com/pandulaDW/interpreter-and-compiler-in-go/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{Statements: []Statement{
		&LetStatement{
			Token: token.Token{Type: token.LET, Literal: "let"},
			Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
			Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
		},
	}}

	assert.Equal(t, "let myVar = anotherVar;", program.String())
}
