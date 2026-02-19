package main

import (
	"fmt"
)

// Stack is a simple stack of runes
type Stack struct {
	items []rune
}

func NewStack() *Stack {
	return &Stack{items: []rune{}}
}

func (s *Stack) Push(r rune) {
	s.items = append(s.items, r)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack) Read() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Linter checks matching braces
type Linter struct {
	stack *Stack
}

func NewLinter() *Linter {
	return &Linter{stack: NewStack()}
}

func (l *Linter) Lint(text string) interface{} {
	for _, char := range text {
		if isOpeningBrace(char) {
			l.stack.Push(char)
		} else if isClosingBrace(char) {
			popped, ok := l.stack.Pop()
			if !ok {
				return fmt.Sprintf("%c doesn't have opening brace", char)
			}
			if isNotAMatch(popped, char) {
				return fmt.Sprintf("%c has mismatched opening brace", char)
			}
		}
	}

	if top, ok := l.stack.Read(); ok {
		return fmt.Sprintf("%c does not have closing brace", top)
	}

	return true
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

// demo
func main() {
	l := NewLinter()
	fmt.Println(l.Lint("{ [ ( ) ] }")) // true
	fmt.Println(l.Lint("{ [ ( ] ) }")) // error
	fmt.Println(l.Lint("{ [ ( ) }"))   // error
}
