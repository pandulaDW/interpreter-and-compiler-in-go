package parser

import (
	"fmt"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/token"
)

// nextToken sets the current token and advances the tokenizer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenIs returns true if t is the current token, false otherwise
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs returns true if t is the next token, false otherwise
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek returns true if t is the expected token and advances the token.
// Returns false otherwise without advancing.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// Errors returns the errors encountered during parsing
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError sets an error if peek token is not the expected token
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t,
		p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// registerPrefix registers prefix functions for a given token type
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix registers prefix functions for a given token type
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
