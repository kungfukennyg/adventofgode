package aoc

import "container/heap"

// QueueItem is a generic item in our PriorityQueue.
type QueueItem[T any] struct {
	wrapped *queueItem
}

func (qi QueueItem[T]) Value() T {
	return qi.wrapped.value.(T)
}

func (qi QueueItem[T]) Priority() int {
	return qi.wrapped.priority
}

func (qi QueueItem[T]) Index() int {
	return qi.wrapped.index
}

// PriorityQueue is a generic priority queue, backed internally by the heap.Interface interface
type PriorityQueue[T any] struct {
	wrapped *priorityQueue
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	wrapped := make(priorityQueue, 0)
	return &PriorityQueue[T]{wrapped: &wrapped}
}

func (pq *PriorityQueue[T]) Init() {
	heap.Init(pq.wrapped)
}

func (pq *PriorityQueue[T]) Push(t T, priority int) int {
	qi := queueItem{
		value:    t,
		priority: priority,
	}
	heap.Push(pq.wrapped, &qi)
	return qi.index
}

func (pq *PriorityQueue[T]) Get(index int) QueueItem[T] {
	internal := *(pq.wrapped)
	qi := internal[index]
	return QueueItem[T]{wrapped: qi}
}

// Consume empties the entire queue and stores each element in order in a slice of its values
func (pq *PriorityQueue[T]) Consume() []T {
	out := make([]T, 0, pq.Len())
	for pq.Len() > 0 {
		qi := pq.Pop()
		out = append(out, qi.Value())
	}
	return out
}

func (pq *PriorityQueue[T]) Pop() QueueItem[T] {
	t := heap.Pop(pq.wrapped)
	qi := t.(*queueItem)
	return QueueItem[T]{wrapped: qi}
}

func (pq *PriorityQueue[T]) Update(index int, value T, priority int) {
	internal := *(pq.wrapped)
	qi := internal[index]
	pq.wrapped.update(qi, value, priority)
	heap.Fix(pq.wrapped, qi.index)
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.wrapped.Len()
}

// the indirected heap interfaces are implemented below. We wrap these interfaces
// because the heap package does not support generics.
type queueItem struct {
	value    any
	priority int
	index    int
}

// priorityQueue implements the heap.Interface interface.
type priorityQueue []*queueItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*queueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	qi := old[n-1]
	old[n-1] = nil
	qi.index = -1
	*pq = old[0 : n-1]
	return qi
}

func (pq *priorityQueue) update(qi *queueItem, value any, priority int) {
	qi.value = value
	qi.priority = priority
}
