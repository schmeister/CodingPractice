package queues

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	q.Push(6)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}