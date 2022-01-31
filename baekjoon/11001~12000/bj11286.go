package bj11286

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type Heap []int

// 절댓값 힙을 구현하기 위해 sort 인터페이스의
// Less() 메서드에서 절댓값을 비교하고 같은 경우에는
// 음수가 먼저 가도록 결과를 반환한다
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if math.Abs(float64(h[i])) == math.Abs(float64(h[j])) {
		return h[i] < h[j]
	}
	return math.Abs(float64(h[i])) < math.Abs(float64(h[j]))
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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
