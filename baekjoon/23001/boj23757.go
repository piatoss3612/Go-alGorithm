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

	N, M     int
	gifts    *MaxHeap
	children []int
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

// 난이도: Silver 2
// 메모리: 7696KB
// 시간: 68ms
// 분류: 자료 구조, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	gifts = new(MaxHeap)
	children = make([]int, M)
	for i := 0; i < N; i++ {
		heap.Push(gifts, scanInt())
	}
	for i := 0; i < M; i++ {
		children[i] = scanInt()
	}
}

func Solve() {
	for i := 0; i < M; i++ {
		max := heap.Pop(gifts).(int) // 선물이 가장 많이 들어있는 상자를 꺼낸다.

		// 선물이 가장 많이 들어있는 상자에 아이가 원하는 만큼의 선물이 들어있지 않으면 0을 출력하고 종료한다.
		if max < children[i] {
			fmt.Fprintln(writer, 0)
			return
		}

		// 아이가 선물을 가져가고도 선물이 남아있으면 다시 큐에 넣는다.
		if max > children[i] {
			heap.Push(gifts, max-children[i])
		}
	}

	fmt.Fprintln(writer, 1)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
