package ast

import (
	"interpreter-go/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	testCases := []struct {
		expected string
		input    *Program
	}{
		{
			expected: "let myVar = anotherVar;",
			input:    program,
		},
	}

	for _, tc := range testCases {
		assertProgramString(t, tc.expected, tc.input)
	}
}

func assertProgramString(t *testing.T, expected string, prog *Program) {
	if prog.String() != expected {
		t.Errorf("program.String() wrong. want %q, got=%q", expected, prog.String())
	}
}
