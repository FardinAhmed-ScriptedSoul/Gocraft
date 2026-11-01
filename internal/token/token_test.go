//table-driven test
package token

import "testing"

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"let", LET},
		{"fn", FN},
		{"if", IF},
		{"else", ELSE},
		{"return", RETURN},
		{"true", TRUE},
		{"false", FALSE},
		{"age", IDENT},
		{"calculate", IDENT},
	}

	for _, tt := range tests {
		actual := LookupIdent(tt.input)

		if actual != tt.expected {
			t.Errorf(
				"LookupIdent(%q) = %q, expected %q",
				tt.input,
				actual,
				tt.expected,
			)
		}
	}
}
