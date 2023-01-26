package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	h, c      int
	coworkers *Coworkers
)

// 우선순위 큐 정의 및 구현
type Coworkers []*Coworker

func (c Coworkers) Len() int { return len(c) }
func (c Coworkers) Less(i, j int) bool {
	// 가장 분노한 동료의 분노 지수를 최소화하려면
	// 도움을 요청했을 때 변화한 분노 지수가 가장 작은 동료에게 먼저 도움을 요청해야 한다
	return c[i].annoyed+c[i].increase < c[j].annoyed+c[j].increase
}
func (c Coworkers) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Coworkers) Push(x interface{}) {
	*c = append(*c, x.(*Coworker))
}
func (c *Coworkers) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

type Coworker struct {
	annoyed  int // 분노 지수
	increase int // 증가량
}

// 난이도: Gold 5
// 메모리: 8468KB
// 시간: 148ms
// 분류: 우선순위 큐, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	h, c = scanInt(), scanInt()
	coworkers = new(Coworkers)
	heap.Init(coworkers)
	for i := 1; i <= c; i++ {
		heap.Push(coworkers, &Coworker{scanInt(), scanInt()})
	}
}

func Solve() {
	// h번 도움 요청하기
	for i := 1; i <= h; i++ {
		c := heap.Pop(coworkers).(*Coworker)
		c.annoyed += c.increase
		heap.Push(coworkers, c)
	}

	// 가장 분노한 동료의 분노 지수 찾기
	ans := 0
	for len(*coworkers) > 0 {
		c := heap.Pop(coworkers).(*Coworker)
		ans = max(ans, c.annoyed)
	}
	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
