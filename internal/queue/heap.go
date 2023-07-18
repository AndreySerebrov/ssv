package queue

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
