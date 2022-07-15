// This package is responsible for parsing the expression by the user into an equation
package parser

import (
	"fmt"
)

type (
	TokenType  int
	Index      int
	Value      float32
	Expression string
	Symbol     byte
)

func (exp Expression) Get(i Index) (Symbol, error) {
	if int(i) >= len(exp) {
		return 0, &ParseError{fmt.Sprintf("Index %d >= length of expression %d", int(i), len(exp))}
	}
	return Symbol(exp[i]), nil
}

const (
	Error TokenType = iota
	Plus
	Minus
	Mul
	Div
	EndOfText
	OpenParenthesis
	ClosedParenthesis
	Number
)

type Token struct {
	tokenType TokenType
	value     Value
	symbol    Symbol
}

type ParseError struct {
	e string
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("Parsing failed due to reason: %s", p.e)
}

func unexpectedTokenError(token *Token, i Index) *ParseError {
	return &ParseError{fmt.Sprintf("Invalid token %s found at %d", string(token.symbol), i)}
}

func nextToken(token *Token, exp Expression, i Index) (*Token, Index, error) {
	i = skipWhitespaces(exp, i)

	token.value = 0
	token.symbol = 0
	c, err := exp.Get(i)
	if err != nil {
		token.tokenType = EndOfText
		return token, i, nil
	}

	if isDigit(exp, i) {
		token.tokenType = Number
		i, num, err := getNumber(exp, i)
		if err != nil {
			return token, i, err
		}
		token.value = num
		return token, i, nil
	}

	token.tokenType = Error

	switch c {
	case '+':
		token.tokenType = Plus
	case '-':
		token.tokenType = Minus
	case '*':
		token.tokenType = Mul
	case '/':
		token.tokenType = Div
	case '(':
		token.tokenType = OpenParenthesis
	case ')':
		token.tokenType = ClosedParenthesis
	}

	if token.tokenType != Error {
		token.symbol = c
		return token, i + 1, nil
	} else {
		return token, i, unexpectedTokenError(token, i)
	}
}

func expression(token *Token, exp Expression, i Index) (*Token, Index, *ASTNode, error) {
	token, i, tNode, err := term(token, exp, i)
	if err != nil {
		return token, i, nil, err
	}

	token, i, e1Node, err := expression1(token, exp, i)
	if err != nil {
		return token, i, nil, err
	}

	return token, i, createASTNode(OperatorPlus, tNode, e1Node), nil
}

func expression1(token *Token, exp Expression, i Index) (*Token, Index, *ASTNode, error) {
	switch token.tokenType {
	case Plus, Minus:
		astNodeType := OperatorPlus
		if token.tokenType == Minus {
			astNodeType = OperatorMinus
		}

		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, tNode, err := term(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, e1Node, err := expression1(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		return token, i, createASTNode(astNodeType, e1Node, tNode), nil

	default:
		return token, i, createNumberASTNode(0), nil
	}
}

func term(token *Token, exp Expression, i Index) (*Token, Index, *ASTNode, error) {
	token, i, fNode, err := factor(token, exp, i)
	if err != nil {
		return token, i, nil, err
	}

	token, i, t1Node, err := term1(token, exp, i)
	if err != nil {
		return token, i, nil, err
	}

	return token, i, createASTNode(OperatorMul, fNode, t1Node), err
}

func term1(token *Token, exp Expression, i Index) (*Token, Index, *ASTNode, error) {
	switch token.tokenType {
	case Mul, Div:
		astNodeType := OperatorMul
		if token.tokenType == Div {
			astNodeType = OperatorDiv
		}

		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, fNode, err := factor(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, t1Node, err := term1(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		node := createASTNode(astNodeType, fNode, t1Node)

		return token, i, node, nil

	default:
		return token, i, createNumberASTNode(1), nil
	}
}

func factor(token *Token, exp Expression, i Index) (*Token, Index, *ASTNode, error) {
	switch token.tokenType {
	case OpenParenthesis:
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, eNode, err := expression(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, err = match(token, exp, i, ')')
		if err != nil {
			return token, i, nil, err
		}

		return token, i, eNode, err

	case Minus:
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		token, i, fNode, err := factor(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		return token, i, createUnaryASTNode(fNode), nil

	case Number:
		value := token.value
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, nil, err
		}

		return token, i, createNumberASTNode(value), nil

	default:
		return token, i, nil, unexpectedTokenError(token, i)
	}
}

func ParseExpression(exp Expression) (*ASTNode, error) {
	token, i, err := nextToken(&Token{Error, 0.0, '0'}, exp, 0)
	if err != nil {
		return nil, err
	}

	_, _, astNode, err := expression(token, exp, i)
	if err != nil {
		return nil, err
	}

	return astNode, nil
}
