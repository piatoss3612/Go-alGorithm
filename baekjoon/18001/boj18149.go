package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	T, n     int
	packages *Packages
)

// 우선순위 큐(최소 힙) 정의 및 구현
type Packages []int

func (p Packages) Len() int           { return len(p) }
func (p Packages) Less(i, j int) bool { return p[i] < p[j] }
func (p Packages) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *Packages) Push(x interface{}) {
	*p = append(*p, x.(int))
}
func (p *Packages) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

// 난이도: Gold 5
// 메모리: 1200KB
// 시간: 8ms
// 분류: 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Input()
		Solve()
	}
}

// 우선순위 큐 초기화 및 n개의 패키지 크기 입력
func Input() {
	n = scanInt()
	packages = new(Packages)
	heap.Init(packages)

	for i := 1; i <= n; i++ {
		heap.Push(packages, scanInt())
	}
}

func Solve() {
	// 그리디 알고리즘:
	// 한 번에 2개의 패키지만 로프로 묶어 패키징할 수 있을 때,
	// 크기가 가장 작은 2개의 패키지를 선택하여 로프로 묶어 하나의 패키지로 변형하는 과정을 반복하여
	// 마지막에는 하나의 커다란 패키지로 만듬으로써 필요한 로프의 길이를 최소화할 수 있다
	totalRope := 0

	for len(*packages) > 1 {
		a, b := heap.Pop(packages).(int), heap.Pop(packages).(int)
		c := a + b
		totalRope += c
		heap.Push(packages, c)
	}

	fmt.Fprintln(writer, totalRope)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
