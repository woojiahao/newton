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

func (exp Expression) Get(i Index) (Symbol, error) {
	if int(i) >= len(exp) {
		return 0, &ParseError{fmt.Sprintf("Index %d >= length of expression %d", int(i), len(exp))}
	}
	return Symbol(exp[i]), nil
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

func expression(token *Token, exp Expression, i Index) (*Token, Index, error) {
	token, i, err := term(token, exp, i)
	if err != nil {
		return token, i, err
	}

	return expression1(token, exp, i)
}

func expression1(token *Token, exp Expression, i Index) (*Token, Index, error) {
	switch token.tokenType {
	case Plus, Minus:
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, err
		}
		token, i, err = term(token, exp, i)
		if err != nil {
			return token, i, err
		}
		return expression1(token, exp, i)

	default:
		return token, i, nil
	}
}

func term(token *Token, exp Expression, i Index) (*Token, Index, error) {
	token, i, err := factor(token, exp, i)
	if err != nil {
		return token, i, err
	}

	return term1(token, exp, i)
}

func term1(token *Token, exp Expression, i Index) (*Token, Index, error) {
	switch token.tokenType {
	case Mul, Div:
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, err
		}
		token, i, err = factor(token, exp, i)
		if err != nil {
			return token, i, err
		}
		return term1(token, exp, i)

	default:
		return token, i, nil
	}
}

func factor(token *Token, exp Expression, i Index) (*Token, Index, error) {
	switch token.tokenType {
	case OpenParenthesis:
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, err
		}
		token, i, err = expression(token, exp, i)
		if err != nil {
			return token, i, err
		}
		return match(token, exp, i, ')')

	case Minus:
		token, i, err := nextToken(token, exp, i)
		if err != nil {
			return token, i, err
		}
		return factor(token, exp, i)

	case Number:
		return nextToken(token, exp, i)

	default:
		return token, i, unexpectedTokenError(token, i)
	}
}

func ParseExpression(exp Expression) error {
	token, i, err := nextToken(&Token{Error, 0.0, '0'}, exp, 0)
	if err != nil {
		return err
	}

	_, _, err = expression(token, exp, i)
	if err != nil {
		return err
	}

	return nil
}

/*
(1+2)*(3+4)
nextToken() - ( 1
expression()
	term()
		factor()
			nextToken() - 1 2
			expression()
				term()
					factor()
						nextToken() - + 3
					term1()
				expression1()
					nextToken() - 2 4
					term()
						factor()
							nextToken() -
*/
