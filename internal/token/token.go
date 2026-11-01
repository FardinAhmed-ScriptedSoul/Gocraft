package token

// TokenType represents the category of a token.
type TokenType string

// Token represents a single lexical token.
type Token struct {
	Type    TokenType
	Literal string
}

// Special tokens
const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
)

// Identifiers and literals
const (
	IDENT  TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"
)

// Operators
const (
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
	LTE    TokenType = "<="
	GTE    TokenType = ">="

	AND TokenType = "&&"
	OR  TokenType = "||"
)

// Delimiters
const (
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"

	LBRACE TokenType = "{"
	RBRACE TokenType = "}"
)

// Keywords
const (
	LET    TokenType = "LET"
	FN     TokenType = "FN"
	IF     TokenType = "IF"
	ELSE   TokenType = "ELSE"
	RETURN TokenType = "RETURN"
	TRUE   TokenType = "TRUE"
	FALSE  TokenType = "FALSE"
)

// keywords maps reserved words to their token types.
var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FN,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// LookupIdent checks whether an identifier is a reserved keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

/*
TokenType = What kind of lexical unit is this?

Literal = What exact source text did we encounter?
*/
