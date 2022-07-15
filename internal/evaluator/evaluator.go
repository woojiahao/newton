package evaluator

import (
	"fmt"
	"woojiahao.com/newton/internal/parser"
)

type EvaluateError struct {
	e string
}

func (e *EvaluateError) Error() string {
	return fmt.Sprintf("Evaluation failed due to reason: %s", e.e)
}

func invalidASTError() *EvaluateError {
	return &EvaluateError{"Invalid AST provided"}
}

func evaluateSubtree(ast *parser.ASTNode) (parser.Value, error) {
	if ast == nil {
		return -1, invalidASTError()
	}

	if ast.ASTNodeType == parser.NumberValue {
		return ast.Value, nil
	}

	if ast.ASTNodeType == parser.UnaryMinus {
		sub, err := evaluateSubtree(ast.Left)
		if err != nil {
			return -1, err
		}

		return -sub, nil
	}

	left, err := evaluateSubtree(ast.Left)
	if err != nil {
		return -1, err
	}

	right, err := evaluateSubtree(ast.Right)
	if err != nil {
		return -1, err
	}

	switch ast.ASTNodeType {
	case parser.OperatorPlus:
		return left + right, nil
	case parser.OperatorMinus:
		return left - right, nil
	case parser.OperatorMul:
		return left * right, nil
	case parser.OperatorDiv:
		return left / right, nil
	}

	return -1, invalidASTError()
}

func EvaluateAST(ast *parser.ASTNode) (parser.Value, error) {
	if ast == nil {
		return -1, invalidASTError()
	}

	return 1, nil
}
