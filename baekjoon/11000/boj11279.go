package bj11279

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
)

type Heap []int

/// 1927번 문제 최소 힙에서 정렬 인터페이스의 비교 메서드만 변경
// h[i] < h[j] --> h[i] > h[j]
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
	*h = old[0 : n-1]
	return x
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	h := &Heap{}
	heap.Init(h)
	n := scanInt()
	for i := 1; i <= n; i++ {
		x := scanInt()
		if x == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, (*h)[0])
				heap.Pop(h)
			}
		} else {
			heap.Push(h, x)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
