package evaluator

import (
	"testing"
	"woojiahao.com/newton/internal/parser"
)

func TestEvaluateAST(t *testing.T) {
	variations := []parser.Expression{
		"1+2+3+4",
		"1*2*3*4",
		"1-2-3-4",
		"1/2/3/4",
		"1*2+3*4",
		"1+2*3+4",
		"(1+2)*(3+4)",
		"1+(2*3)*(4+5)",
		"1+(2*3)/4+5",
		"5/(4+3)/2",
		"1 + 2.5",
		"125",
		"-1",
		"-1+(-2)",
		"-1+(-2.0)",
	}

	for _, v := range variations {
		_, err := ParseExpression(v)
		if err != nil {
			t.Errorf("Expression %s should be successful, got %v instead\n", v, err)
		}
	}
}
