package ast

import "interpreter-go/token"

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (p PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p PrefixExpression) String() string {
	return "(" + p.Operator + p.Right.String() + ")"
}

func (p PrefixExpression) expressionNode() {}
