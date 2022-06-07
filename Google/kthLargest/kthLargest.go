package kthLargest

/*  Heap, PriorityQueue */

import (
	"container/heap"
	"fmt"
)

// KthLargest object will be instantiated and called as such:
// obj := Constructor(k, nums);
// param_1 := obj.Add(val);
type KthLargest struct {
	k    int
	Heap intHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	//Convert to heap
	heap.Init(&h)

	return KthLargest{
		k:    k,
		Heap: h,
	}
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.Heap, val)

	//Pop minimum element until len(h) = k and kthlargest = h[0]
	for len(kl.Heap) > kl.k {
		heap.Pop(&kl.Heap)
	}

	return kl.Heap[0]
}

func KthLargestElement() {
	k := 3
	nums := []int{4, 5, 8, 2}

	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3), obj.Heap)
	fmt.Println(obj.Add(5), obj.Heap)
	fmt.Println(obj.Add(10), obj.Heap)
	fmt.Println(obj.Add(9), obj.Heap)
	fmt.Println(obj.Add(4), obj.Heap)

}

/////////////

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
