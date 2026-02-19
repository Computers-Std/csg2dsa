package main

import (
	"fmt"
	"ukiran/linter/stack"
)

func Lint(text string) (bool, error) {
	s := stack.Stack{}
	for _, r := range text {
		switch {
		case isOpeningBrace(r):
			s.Push(r)
		case isClosingBrace(r):
			top, err := s.Top()
			if err != nil { // empty stack
				return false, fmt.Errorf("%c doesn't have opening brace", r)
			}
			if isNotAMatch(top, r) {
				return false, fmt.Errorf("%c has mismatched opening brace", r)
			} else {
				s.Pop()
			}
		}
	}
	if s.IsEmpty() {
		return true, nil
	}
	return false, fmt.Errorf("mismatching braces")
}

// helpers
func isOpeningBrace(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func isClosingBrace(r rune) bool {
	return r == ')' || r == ']' || r == '}'
}

func isNotAMatch(opening, closing rune) bool {
	matches := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	return matches[opening] != closing
}

func main() {
	fmt.Println(Lint("{ [ ( ) ] }")) // true
	fmt.Println(Lint("{ [ ( ] ) }")) // false
	fmt.Println(Lint("{ [ ( ) }"))   // false
}
