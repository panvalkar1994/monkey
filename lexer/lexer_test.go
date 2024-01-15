package lexer

import (
	"testing"

	"github.com/panvalkar1994/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	// Create a new lexer
	l := New(input)

	for i, tt := range tests {
		// Get the next token
		tok := l.NextToken()

		// Check the type of the token
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", 
				i, tt.expectedType, tok.Type)
		}

		// Check the literal value of the token
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", 
				i, tt.expectedLiteral, tok.Literal)
		}
	}
	
}