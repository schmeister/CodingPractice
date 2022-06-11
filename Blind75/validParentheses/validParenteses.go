package validParentheses

// create stack
type Stack []rune

var openers = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}
var closers = map[rune]bool{
	'}': true,
	')': true,
	']': true,
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(r rune) {
	*stack = append(*stack, r)
}

func (stack *Stack) Peek() rune {
	if stack.IsEmpty() {
		return rune(0)
	}
	return (*stack)[len(*stack)-1]
}

func (stack *Stack) Pop() rune {
	if stack.IsEmpty() {
		return rune(0)
	}
	last := len(*stack) - 1
	r := (*stack)[last]
	*stack = (*stack)[:last]
	return r
}

func ValidParenteses(str string) bool {

	runes := []rune(str)

	stack := Stack{}

	for _, r := range runes {
		v, isOpener := openers[r]
		if isOpener {
			stack.Push(v)
			continue
		}

		_, isCloser := closers[r]
		if isCloser {
			popd := stack.Pop()
			if popd != r {
				return false
			}
		}
	}

	return stack.IsEmpty()
}
