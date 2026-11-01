
package lexer

import "github.com/FardinAhmed-ScriptedSoul/Gocraft/internal/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	if isLetter(l.ch) {
		literal := l.readIdentifier()

		return token.Token{
			Type:    token.LookupIdent(literal),
			Literal: literal,
		}
	}

	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: "-"}
	case '!':
		tok = token.Token{Type: token.BANG, Literal: "!"}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: "*"}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: "/"}
	case '<':
		tok = token.Token{Type: token.LT, Literal: "<"}
	case '>':
		tok = token.Token{Type: token.GT, Literal: ">"}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: ","}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: ";"}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: "("}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: ")"}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: "}"}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		tok = token.Token{
			Type:    token.ILLEGAL,
			Literal: string(l.ch),
		}
	}

	l.readChar()

	return tok
}
