
package lexer

import (
	"testing"

	"github.com/FardinAhmed-ScriptedSoul/Gocraft/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `let age = 25;
let salary = 50000;

fn add(a, b) {
    return a + b;
}

if (age >= 18) {
    print(age);
}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "age"},
		{token.ASSIGN, "="},
		{token.INT, "25"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "salary"},
		{token.ASSIGN, "="},
		{token.INT, "50000"},
		{token.SEMICOLON, ";"},

		{token.FN, "fn"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "age"},
		{token.GTE, ">="},
		{token.INT, "18"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.IDENT, "age"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d]: wrong token type. got=%q, want=%q",
				i, tok.Type, tt.expectedType,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d]: wrong literal. got=%q, want=%q",
				i, tok.Literal, tt.expectedLiteral,
			)
		}
	}
}

func TestStringLiterals(t *testing.T) {
	input := `"hello" "Gocraft compiler" "123"`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "hello"},
		{token.STRING, "Gocraft compiler"},
		{token.STRING, "123"},
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

func TestUnterminatedString(t *testing.T) {
	l := New(`"hello`)

	tok := l.NextToken()

	if tok.Type != token.ILLEGAL {
		t.Fatalf("got type %q, want %q", tok.Type, token.ILLEGAL)
	}

	if tok.Literal != "hello" {
		t.Fatalf("got literal %q, want %q", tok.Literal, "hello")
	}
}

func TestSingleLineComments(t *testing.T) {
	input := `// This is a comment
let age = 25; // Another comment
// Final comment`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "age"},
		{token.ASSIGN, "="},
		{token.INT, "25"},
		{token.SEMICOLON, ";"},
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

func TestRealisticGocraftProgram(t *testing.T) {
	input := `
	// Calculate a user's salary bonus.
	let name = "Fardin";
	let salary = 50000;
	let bonus = 1000;

	fn calculate(total, extra) {
		return total + extra;
	}

	if (salary >= 50000 && true) {
		print(name);
		print(calculate(salary, bonus));
	}
	`

	expected := []token.TokenType{
		token.LET, token.IDENT, token.ASSIGN, token.STRING, token.SEMICOLON,
		token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON,
		token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON,

		token.FN, token.IDENT, token.LPAREN, token.IDENT, token.COMMA,
		token.IDENT, token.RPAREN, token.LBRACE,

		token.RETURN, token.IDENT, token.PLUS, token.IDENT, token.SEMICOLON,
		token.RBRACE,

		token.IF, token.LPAREN, token.IDENT, token.GTE, token.INT,
		token.AND, token.TRUE, token.RPAREN, token.LBRACE,

		token.IDENT, token.LPAREN, token.IDENT, token.RPAREN, token.SEMICOLON,

		token.IDENT, token.LPAREN, token.IDENT, token.LPAREN,
		token.IDENT, token.COMMA, token.IDENT, token.RPAREN,
		token.RPAREN, token.SEMICOLON,

		token.RBRACE,
		token.EOF,
	}

	l := New(input)

	for i, expectedType := range expected {
		tok := l.NextToken()

		if tok.Type != expectedType {
			t.Fatalf(
				"token %d: got %q (%q), want type %q",
				i, tok.Type, tok.Literal, expectedType,
			)
		}
	}
}
