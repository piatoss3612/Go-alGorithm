package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	A, B, C, N int
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

// 난이도: Gold 5
// 메모리: 1036KB
// 시간: 4ms
// 분류: 정수론, 최소 힙
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	A, B, C, N = scanInt(), scanInt(), scanInt(), scanInt()
}

func Solve() {
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, A)
	heap.Push(h, B)
	heap.Push(h, C)

	for i := 1; i <= N; i++ {
		v := heap.Pop(h).(int)
		if i == N {
			fmt.Fprintln(writer, v)
			return
		}

		heap.Push(h, v*A)
		heap.Push(h, v*B)
		heap.Push(h, v*C)

		for h.Len() > 0 {
			w := heap.Pop(h).(int)
			if v != w {
				heap.Push(h, w)
				break
			}
		}
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
