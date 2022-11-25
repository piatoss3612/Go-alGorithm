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
	N       int
	h       *MinHeap
)

// 최소 힙 정의
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 2320KB
// 시간: 20ms
// 분류: 그리디 알고리즘, 우선순위 큐
// 널빤지를 톱질하여 나눌 때마다 나눠진 크기 a, b만큼의 비용이 청구되고
// 총 N개의 널빤지를 만들기 위해 N-1번 톱질이 필요한 경우의 총 비용의 최솟값 구하기
// = 반대로 생각해보면 널빤지를 합치는 비용의 최솟값을 구하면 되는 것
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	h = new(MinHeap)
	heap.Init(h)
	// N개의 널빤지가 필요하고 필요한 널빤지의 크기가 주어진다
	for i := 1; i <= N; i++ {
		heap.Push(h, scanInt())
	}
}

func Solve() {
	totalCost := 0
	// 주어진 널빤지의 크기를 가장 작은 값을 기준으로
	// 2개씩 꺼내어 합치고 그 비용을 전체 비용에 추가
	// 합친 널빤지 크기는 최소 힙에 집어넣는다
	// 널빤지를 더 이상 합칠 수 없을 때까지 반복
	for len(*h) > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		totalCost += a + b
		heap.Push(h, a+b)
	}
	fmt.Fprintln(writer, totalCost)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
