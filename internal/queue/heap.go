package queue

// type TaskHeap []Task

// func (h TaskHeap) Len() int           { return len(h) }
// func (h TaskHeap) Less(i, j int) bool { return h[i].height < h[j].height }
// func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *TaskHeap) Push(x Task) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x)
// }

// func (h *TaskHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// type Item struct {
// 	value    string // The value of the item; arbitrary.
// 	priority int    // The priority of the item in the queue.
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// A TaskHeap implements heap.Interface and holds Items.
type TaskHeap []*Task

func (pq TaskHeap) Len() int { return len(pq) }

func (pq TaskHeap) Less(i, j int) bool {
	return pq[i].Height > pq[j].Height
}

func (pq TaskHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TaskHeap) Push(x any) {
	item := x.(Task)
	*pq = append(*pq, &item)
}

func (pq *TaskHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
