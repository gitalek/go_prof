package hw02unpackstring

import (
	"strings"
	"unicode"
)

type tokenState int

const (
	done      tokenState = 0
	escaped   tokenState = 1
	suspended tokenState = 2
	invalid   tokenState = 3
)

type Token struct {
	char   string
	repeat int
	state  tokenState
}

func NewToken(r rune) *Token {
	state := suspended
	switch {
	case unicode.IsDigit(r):
		state = invalid
	case r == '\\':
		state = escaped
	}
	return &Token{char: string(r), repeat: 1, state: state}
}

func (t *Token) Set(r rune) {
	t.char = string(r)
	t.state = suspended
}

func (t *Token) Repeat(times int) {
	t.repeat = times
}

func (t *Token) Freeze() {
	t.state = done
}

func (t *Token) String() string {
	return strings.Repeat(t.char, t.repeat)
}
