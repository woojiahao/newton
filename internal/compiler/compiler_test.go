package compiler

import (
	"testing"
	"woojiahao.com/newton/internal/parser"
)

func TestEvaluateAST(t *testing.T) {
	variations := map[parser.Expression]parser.Value{
		"1+2+3+4":       10,
		"1*2*3*4":       24,
		"1-2-3-4":       -8,
		"1/2/3/4":       0.0416667,
		"1*2+3*4":       14,
		"1+2*3+4":       11,
		"(1+2)*(3+4)":   21,
		"1+(2*3)*(4+5)": 55,
		"1+(2*3)/4+5":   7.5,
		"5/(4+3)/2":     0.357143,
		"1 + 2.5":       3.5,
		"125":           125,
		"-1":            -1,
		"-1+(-2)":       -3,
		"-1+(-2.0)":     -3,
	}

	for exp, expected := range variations {
		result := Compile(exp)
		if result != expected {
			t.Errorf("Expression %s evaluated to %f, not %f", exp, result, expected)
		}
	}
}
