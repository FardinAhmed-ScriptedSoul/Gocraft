package token

// TokenType represents the category of a token.
type TokenType string

// Token represents a single lexical token.
type Token struct {
	Type    TokenType
	Literal string
}
