package parser

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func isDigit(exp Expression, i Index) bool {
	c, err := exp.Get(i)
	if err != nil {
		return false
	}
	ch, _ := utf8.DecodeRune([]byte{byte(c)})
	return unicode.IsDigit(ch)
}

func isSpace(exp Expression, i Index) bool {
	c, err := exp.Get(i)
	if err != nil {
		return false
	}
	ch, _ := utf8.DecodeRune([]byte{byte(c)})
	return unicode.IsSpace(ch)
}

// Create substring from expression from start (inclusive) to end (exclusive)
func substring(text string, start, end Index) string {
	// If end is equal to the last index, then just return the substring from start to end
	if int(end) >= len(text) {
		return text[start:]
	}

	return text[start:end]
}

func skipWhitespaces(exp Expression, i Index) Index {
	for {
		if !isSpace(exp, i) {
			return i
		}
		i++
	}
}

func trackNumber(exp Expression, i Index) Index {
	for {
		if !isDigit(exp, i) {
			break
		}
		i++
	}

	d, err := exp.Get(i)
	if err != nil {
		// Means this is the last digit already
		return i
	}

	if d == Symbol('.') {
		i++
	}

	for {
		if !isDigit(exp, i) {
			break
		}
		i++
	}

	return i
}

func getNumber(exp Expression, i Index) (Index, Value, error) {
	// p is a moving pointer that creates a span over the text with i to monitor where the number is
	i = skipWhitespaces(exp, i)

	// If the first character is not even a digit, return a ParseError automatically
	if !isDigit(exp, i) {
		return -1, -1.0, &ParseError{"Number expected but not found!"}
	}

	p := trackNumber(exp, i)

	num, err := strconv.ParseFloat(substring(string(exp), i, p), 32)
	if err != nil {
		return -1, -1.0, err
	}

	return p, Value(num), nil
}

func match(token *Token, exp Expression, i Index, expected Symbol) (*Token, Index, error) {
	if c, err := exp.Get(i - 1); c == expected {
		return nextToken(token, exp, i)
	} else if err != nil {
		return token, i, err
	} else {
		return token, i, &ParseError{fmt.Sprintf("Expected token %s did not match given token %s", string(expected), string(c))}
	}
}
