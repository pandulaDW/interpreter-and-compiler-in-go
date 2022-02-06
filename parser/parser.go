package parser

import (
	"github.com/pandulaDW/interpreter-and-compiler-in-go/ast"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/lexer"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken sets the current token and advances the tokenizer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram builds the ast and returns the root
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
