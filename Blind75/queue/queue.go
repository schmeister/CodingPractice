package queues

type Queue []int

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Peek() int {
	val := (*q)[0]

	return val
}

func (q *Queue) Pop() int {

	val := (*q)[0]
	*q = (*q)[1:]

	return val
}
