package parser

import (
	"github.com/pandulaDW/interpreter-and-compiler-in-go/ast"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/lexer"
	"github.com/stretchr/testify/require"
	"testing"
)

func initiateTest(t *testing.T, input string) *ast.Program {
	t.Helper()
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	return program
}

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
`
	program := initiateTest(t, input)
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
	program := initiateTest(t, input)
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

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	program := initiateTest(t, input)
	require.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	ident, ok := stmt.Expression.(*ast.Identifier)
	require.True(t, ok)
	require.Equal(t, "foobar", ident.Value)
	require.Equal(t, "foobar", ident.TokenLiteral())
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "25;"
	program := initiateTest(t, input)
	require.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	require.True(t, ok)
	require.Equal(t, 25, ident.Value)
	require.Equal(t, "25", ident.TokenLiteral())
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-25", "-", 25},
	}

	for _, tt := range prefixTests {
		program := initiateTest(t, tt.input)
		require.Equal(t, 1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		require.True(t, ok)

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		require.True(t, ok)
		require.Equal(t, tt.operator, exp.Operator)

		intExpr, ok := exp.Right.(*ast.IntegerLiteral)
		require.True(t, ok)
		require.Equal(t, tt.integerValue, intExpr.Value)
	}
}
