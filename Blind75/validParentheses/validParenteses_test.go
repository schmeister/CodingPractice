package validParentheses

import (
	"fmt"
	"testing"
)

func TestValidParenteses(t *testing.T) {
	fmt.Println(ValidParenteses("()"))
	fmt.Println(ValidParenteses("()({})"))
	fmt.Println(ValidParenteses("()({}){}[]{}"))
}
