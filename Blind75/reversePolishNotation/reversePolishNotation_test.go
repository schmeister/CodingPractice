package reversePolishNotation

import (
	"fmt"
	"testing"
)

func TestRPN(t *testing.T) {

	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(RPN(tokens))
	tokens2 := []string{"4","13","5","/","+"}
	fmt.Println(RPN(tokens2))
	tokens3 := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	fmt.Println(RPN(tokens3))
}
