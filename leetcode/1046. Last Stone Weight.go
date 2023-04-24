package main

import "container/heap"

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Solution - Heap
// Runtime: 0ms
// Memory Usage: 2MB
// Time complexity: O(nlogn)
// Space complexity: O(n)
func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	h := Heap(stones)
	heap.Init(&h)

	for h.Len() >= 2 {
		a, b := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, sub(a, b))
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return heap.Pop(&h).(int)
}

func sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
