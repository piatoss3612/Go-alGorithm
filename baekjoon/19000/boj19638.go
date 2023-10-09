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
	N, H, T int
	giants  *Heap
)

// 최대 힙 정의 및 구현
type Heap []int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
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

// 난이도: Silver 1
// 메모리: 7084KB
// 시간: 68ms
// 분류: 구현, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, H, T = scanInt(), scanInt(), scanInt()
	giants = new(Heap)
	heap.Init(giants)
	for i := 1; i <= N; i++ {
		heap.Push(giants, scanInt())
	}
}

func Solve() {
	cnt := 0 // 뿅망치 사용 횟수
	for cnt < T {
		giant := heap.Pop(giants).(int)
		// 기저 사례: 거인의 키가 센티보다 작거나 1인경우 탐색 종료
		if giant < H || giant == 1 {
			heap.Push(giants, giant)
			break
		}

		giant /= 2               // 거인의 키 반땅
		cnt++                    // 뿅망치 사용 횟수 증가
		heap.Push(giants, giant) // 최대 힙으로 거인 되돌리기
	}

	biggest := heap.Pop(giants).(int)
	if biggest < H {
		fmt.Fprintln(writer, "YES")
		fmt.Fprintln(writer, cnt)
	} else {
		fmt.Fprintln(writer, "NO")
		fmt.Fprintln(writer, biggest)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
