package compiler

import (
	"woojiahao.com/newton/internal/evaluator"
	"woojiahao.com/newton/internal/parser"
)

func Compile(expression parser.Expression) parser.Value {
	ast, err := parser.ParseExpression(expression)
	if err != nil {
		panic(err)
	}

	value, err := evaluator.EvaluateAST(ast)
	if err != nil {
		panic(err)
	}

	return value
}