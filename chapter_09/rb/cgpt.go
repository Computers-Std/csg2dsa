package main

// Stack is a simple stack of runes
type Stack struct {
	items []rune
}

func NewStack() *Stack {
	return &Stack{items: []rune{}}
}
