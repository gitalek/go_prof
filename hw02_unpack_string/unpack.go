package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runes := []rune(s)
	runesLength := len(runes)
	switch runesLength {
	case 0:
		return "", nil
	case 1:
		t := NewToken(runes[0])
		if t.state != suspended {
			return "", nil
		}
		return t.String(), nil
	}

	prev := runes[0]
	rest := runes[1:]

	tokens := make([]*Token, 0, runesLength)
	tokenCurrent := NewToken(prev)

	for _, cur := range rest {
		switch tokenCurrent.state {
		case invalid:
			return "", ErrInvalidString
		case done:
			tokenCurrent = NewToken(cur)
		case escaped:
			tokenCurrent.Set(cur)
		case suspended:
			switch {
			case unicode.IsDigit(cur):
				times, err := strconv.Atoi(string(cur))
				if err != nil {
					return "", err
				}
				tokenCurrent.Repeat(times)
				tokenCurrent.Freeze()
				tokens = append(tokens, tokenCurrent)
			default:
				tokenCurrent.Freeze()
				tokens = append(tokens, tokenCurrent)
				tokenCurrent = NewToken(cur)
			}
		default:
			continue
		}
	}

	if tokenCurrent.state != done {
		tokenCurrent.Freeze()
		tokens = append(tokens, tokenCurrent)
	}

	return convertTokensToString(tokens), nil
}

func convertTokensToString(tokens []*Token) string {
	var b strings.Builder
	for _, t := range tokens {
		b.WriteString(t.String())
	}
	return b.String()
}
