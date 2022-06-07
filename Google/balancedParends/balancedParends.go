package balancedParends

import "fmt"

type Stack []rune

func (stack *Stack) isEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) push(r rune) {
	*stack = append(*stack, r) // Append to the end of the slice.
}

func (stack *Stack) peek() rune {
	last := len(*stack) - 1
	return (*stack)[last] // Peek at the last element in the slice.
}

func (stack *Stack) pop() rune {
	r := rune(0)
	if stack.isEmpty() {
		return r
	} else {
		last := len(*stack) - 1   // Get the index of the top (last) most element.
		element := (*stack)[last] // Index into the slice and obtain the element.
		*stack = (*stack)[:last]  // Remove last element.
		return element
	}
}

func isBalanced(str string) bool {
	runes := []rune(str)

	s := Stack{}
	for _, r := range runes {
		if r == '{' {
			s.push('}')
		} else if r == '[' {
			s.push(']')
		} else if r == '(' {
			s.push(')')
		} else if r == '}' {
			if s.peek() != '}' {
				return false
			}
			s.pop()
		} else if r == ']' {
			if s.peek() != ']' {
				return false
			}
			s.pop()
		} else if r == ')' {
			if s.peek() != ')' {
				return false
			}
			s.pop()
		}
	}

	if !s.isEmpty() {
		return false
	}
	return true
}

func Balanced() {
	fmt.Println(isBalanced("ab{cde)}"))
	fmt.Println(isBalanced("ab{cde}"))
	fmt.Println(isBalanced("ab{{cde}"))
	fmt.Println(isBalanced("ab{[]()cde}"))
	fmt.Println(isBalanced("ab{cde}"))
}
