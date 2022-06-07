package lruCache

import "fmt"

type Node struct {
	next      *Node
	frameNum  interface{}
	frameData interface{}
}

type List struct {
	head *Node
}

var caches = make(map[interface{}]interface{})

func (L *List) Insert(frameNum interface{}) {

	// Check to see if cacheMap contains key, if so, move to front
	if _, ok := caches[frameNum]; ok {
		curr := L.head

		var prev *Node
		for curr != nil {
			if curr.frameNum == frameNum {
				break
			}
			prev = curr
			curr = curr.next
		}
		prev.next = curr.next
		curr.next = L.head
		L.head = curr
		return
	}
	// Add key to map and front of list
	caches[frameNum] = frameNum
	list := &Node{
		next:     L.head,
		frameNum: frameNum,
	}
	L.head = list

	// Make sure number of frames is less than cache size
	if len(caches) > 5 {
		// find the last item in the cache - remove it from cache and map
		curr := L.head
		var prev *Node
		for curr != nil {
			if curr.next == nil {
				prev.next = nil
				delete(caches, curr.frameNum)
			}
			prev = curr
			curr = curr.next
		}
	}
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v", list.frameNum)
		if list.next != nil {
			fmt.Printf(" -> ")
		}
		list = list.next
	}
	fmt.Println()
}

func LRUCache() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(5)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)
	link.Insert(37)
	link.Insert(38)

	fmt.Println("\n==============================")
	link.Display()
	fmt.Println("==============================")
	link.Insert(5)
	link.Display()
	fmt.Println("==============================")
	link.Insert(37)
	link.Display()
	fmt.Println("==============================")
}
