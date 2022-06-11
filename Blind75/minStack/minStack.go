package minStack

type MinStack struct {
	stack []int
	minS  []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minS:  []int{},
		min:   1<<63 - 1,
	}
}

func (stack *MinStack) Push(val int) {
	if val < stack.min {
		stack.min = val
		stack.minS = append(stack.minS, val)
	}
	stack.stack = append(stack.stack, val)
}

func (stack *MinStack) Pop() {
	if stack.stack[len(stack.stack)-1] == stack.min {
		stack.minS = stack.minS[:len(stack.minS)-1]
		if len(stack.minS) > 0 {
			stack.min = stack.minS[len(stack.minS)-1]
		} else {
			stack.min = 1<<63 - 1
		}
	}
	stack.stack = stack.stack[:len(stack.stack)-1]
}

func (stack *MinStack) Top() int {
	return stack.stack[len(stack.stack)-1]
}

func (stack *MinStack) GetMin() int {
	return stack.min
}
