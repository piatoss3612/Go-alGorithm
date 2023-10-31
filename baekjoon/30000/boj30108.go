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
	tree    [][]int
	val     []int
)

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return val[pq[i]] > val[pq[j]]
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

// 30108번: 교육적인 트리 문제
// https://www.acmicpc.net/problem/30108
// 난이도: 골드 4
// 메모리: 36780 KB
// 시간: 332 ms
// 분류: 자료 구조, 그리디 알고리즘, 우선순위 큐, 정렬, 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	tree = make([][]int, N+1)
	for i := 2; i <= N; i++ {
		p := scanInt()
		tree[p] = append(tree[p], i)
	}
	val = make([]int, N+1)
	for i := 1; i <= N; i++ {
		val[i] = scanInt()
	}
}

func Solve() {
	sum := 0
	pq := new(PQ)
	heap.Init(pq)

	heap.Push(pq, 1)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(int)
		sum += val[cur]
		fmt.Fprintln(writer, sum)

		for _, next := range tree[cur] {
			heap.Push(pq, next)
		}
	}
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
