package parser

import (
	"fmt"
	"testing"
)

func TestGetNumberPass(t *testing.T) {
	passVariations := map[Expression]float32{
		"15":             15,
		" 15":            15,
		"  15":           15,
		"  15  ":         15,
		"124.15":         124.15,
		"15879.65498498": 15879.65498498,
		" 124.15":        124.15,
		" 124.15  ":      124.15,
	}

	for k, v := range passVariations {
		fmt.Printf("Testing getNumber on '%s'\n", k)
		_, n, _ := getNumber(k, 0)
		fmt.Printf("n from %s is %f\n", k, n)
		if float32(n) != v {
			t.Errorf("Expected  %f to be parsed from %s but got %f instead\n", v, k, n)
		}
	}
}

func TestGetNumberFail(t *testing.T) {
	failVariations := []Expression{
		"",
		"  ",
		"foo",
	}

	for _, v := range failVariations {
		_, _, err := getNumber(v, 0)
		if _, ok := err.(*ParseError); !ok {
			t.Errorf("Expected ParseError, got %v instead\n", err)
		}
	}
}

func TestParseExpressionPass(t *testing.T) {
	variations := []Expression{
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

func TestParseExpressionFail(t *testing.T) {
	variations := []Expression{
		"   1*2,5",
		"   1*2.5e2",
		"M1 + 2.5",
		"1 + 2&5",
		"1 * 2.5.6",
		"1 ** 2.5",
		"*1 / 2.5",
	}

	for _, v := range variations {
		_, err := ParseExpression(v)
		if _, ok := err.(*ParseError); !ok {
			t.Errorf("Expected ParseError, got %v instead\n", err)
		}
	}
}
