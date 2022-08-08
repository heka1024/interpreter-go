package token

import (
	"fmt"
	"testing"
)

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input string
		want  Type
	}{
		{"fn", FUNCTION},
		{"let", LET},
		{"five", IDENT},
		{"ten", IDENT},
		{"true", TRUE},
		{"false", FALSE},
		{"if", IF},
		{"else", ELSE},
		{"return", RETURN},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("input: %s, want: %s", tt.input, tt.want), func(t *testing.T) {
			if got := LookupIdent(tt.input); got != tt.want {
				t.Errorf("LookupIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}
