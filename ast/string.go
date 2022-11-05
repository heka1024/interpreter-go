package ast

import "interpreter-go/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (s *StringLiteral) TokenLiteral() string {
	return s.Token.Literal
}

func (s *StringLiteral) String() string {
	return s.Value
}

func (s *StringLiteral) expressionNode() {}
