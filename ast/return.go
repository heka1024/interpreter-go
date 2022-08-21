package ast

import "interpreter-go/token"

type ReturnStatement struct {
	Token       token.Token // return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
