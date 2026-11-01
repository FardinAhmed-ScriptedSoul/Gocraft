package lexer

import (
	"testing"

	"github.com/FardinAhmed-ScriptedSoul/Gocraft/internal/token"
)

func TestReadChar(t *testing.T) {
	l := New("let")

	expected := []byte{'l', 'e', 't', 0}

	for i, want := range expected {
		if l.ch != want {
			t.Fatalf("step %d: got %q, want %q", i, l.ch, want)
		}

		if i < len(expected)-1 {
			l.readChar()
		}
	}
}

func TestPeekChar(t *testing.T) {
	l := New("ab")

	if got := l.peekChar(); got != 'b' {
		t.Fatalf("peekChar() = %q, want 'b'", got)
	}

	if l.ch != 'a' || l.position != 0 || l.readPosition != 1 {
		t.Fatal("peekChar() unexpectedly changed lexer state")
	}

	l.readChar()

	if got := l.peekChar(); got != 0 {
		t.Fatalf("peekChar() = %q, want EOF (0)", got)
	}
}

func TestNextToken(t *testing.T) {
	input := "+-*/(){};,"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d]: got type %q, want %q",
				i, tok.Type, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]: got literal %q, want %q",
				i, tok.Literal, tt.expectedLiteral)
		}
	}
}

func TestIdentifiersAndKeywords(t *testing.T) {
	input := "let age calculate _name age25"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "age"},
		{token.IDENT, "calculate"},
		{token.IDENT, "_name"},
		{token.IDENT, "age25"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d]: got type %q, want %q",
				i, tok.Type, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]: got literal %q, want %q",
				i, tok.Literal, tt.expectedLiteral)
		}
	}
}

func TestIntegerLiterals(t *testing.T) {
	input := "5 25 250 12345"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "5"},
		{token.INT, "25"},
		{token.INT, "250"},
		{token.INT, "12345"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d]: got type %q, want %q",
				i, tok.Type, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]: got literal %q, want %q",
				i, tok.Literal, tt.expectedLiteral)
		}
	}
}
