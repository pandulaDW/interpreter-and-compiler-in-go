package parser

import (
	"github.com/pandulaDW/interpreter-and-compiler-in-go/ast"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/lexer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	require.NotNil(t, program)
	checkParseErrors(t, p)
	require.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
   return 5;
   return 10;
   return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	require.NotNil(t, program)
	checkParseErrors(t, p)
	require.Equal(t, 3, len(program.Statements))

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		require.True(t, ok)
		require.Equal(t, "return", returnStatement.TokenLiteral())
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	t.Helper()
	require.Equal(t, s.TokenLiteral(), "let")

	letStmt, ok := s.(*ast.LetStatement)
	require.True(t, ok)

	require.Equal(t, name, letStmt.Name.Value)
	require.Equal(t, name, letStmt.Name.TokenLiteral())
}

func checkParseErrors(t *testing.T, p *Parser) {
	t.Helper()

	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d error(s)", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
