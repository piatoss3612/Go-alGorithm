package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, M int
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 난이도: Gold 5
// 메모리: 9352KB
// 시간: 88ms
// 분류: 자료 구조, 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
}

func Solve() {
	cnt := 0
	sum := 0

	h := &MaxHeap{}
	heap.Init(h)

	for i := 1; i <= N; i++ {
		x := scanInt()
		sum += x
		heap.Push(h, x)

		if sum >= M {
			y := heap.Pop(h).(int)
			sum -= y * 2
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
