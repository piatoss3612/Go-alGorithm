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
	pq      *PQ
)

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// 28703번: Double it
// https://www.acmicpc.net/problem/28703
// 난이도: 골드 3
// 메모리: 17316 KB
// 시간: 1308 ms
// 분류: 자료 구조, 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	pq = new(PQ)
	heap.Init(pq)
}

func Solve() {
	mx := 0 // 입력받은 수 중 최대값
	for i := 0; i < N; i++ {
		x := scanInt()
		heap.Push(pq, x)
		mx = max(mx, x)
	}

	curMax := mx
	ans := 987654321

	// 모든 연산이 끝난 후 pq에 들어있는 수 중 최솟값을 v라고 하면
	// 최댓값과 최솟값의 차이의 최솟값을 구하려면 최댓값을 최소화해야 한다.
	// 따라서 pq에 들어있는 나머지 수들이 모두 v보다 크거나 같아질 때까지
	// 연산을 수행하고 그 과정에서 차이의 최솟값을 갱신해준다.
	// 이 때의 v를 가능한 작은 값으로 설정하기 위해 pq에 들어있는 수 중 최댓값 mx를 사용한다.
	for {
		e := heap.Pop(pq).(int)
		ans = min(ans, curMax-e)

		if e >= mx {
			heap.Push(pq, e)
			break
		}

		heap.Push(pq, e*2)
		curMax = max(curMax, e*2)
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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