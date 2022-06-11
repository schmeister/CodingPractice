package minStack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {

	fmt.Println(1<<63 - 1)

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
