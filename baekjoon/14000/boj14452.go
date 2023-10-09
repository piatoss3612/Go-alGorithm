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
	N, T    int
	cows    []int
)

type PQ []int

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 3556KB
// 시간: 32ms
// 분류: 우선순위 큐, 이분 탐색

// 틀린 이유 - 놓친 부분:
// 1. 소들이 춤을 추는 순서는 1~N번까지 순서대로 -> 정렬하면 안된다
// 2. 가능한 무대의 크기 K의 최솟값 구하기 -> 이분 탐색 필요
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, T = scanInt(), scanInt()
	cows = make([]int, N)
	for i := 0; i < N; i++ {
		cows[i] = scanInt()
	}
}

func Solve() {
	// 무대의 크기 K의 최댓값은 소들의 수 N과 같으므로
	// 이분탐색 r의 값을 N으로 초기화
	l, r := 0, N
	var ans int

	for l <= r {
		m := (l + r) / 2
		// 무대의 크기가 m일 때 모든 소들이 T시간 안에 춤을 출 수 있는지 체크
		if check(m) {
			ans = m
			r = m - 1 // 가능하다면 무대의 크기를 더 줄여본다
		} else {
			l = m + 1 // 불가능하다면 무대의 크기를 늘려본다
		}
	}
	fmt.Fprintln(writer, ans)
}

func check(k int) bool {
	pq := new(PQ)
	heap.Init(pq)

	// 무대(우선순위 큐)에 k마리의 소가 먼저 올라간다
	for i := 0; i < k; i++ {
		heap.Push(pq, cows[i])
	}

	// 미리 올라간 k마리의 소를 제외하고
	// 먼저 무대에서 내려오는 소의 시간에 맞춰 k+1번째 소부터 차례대로 무대에 오른다
	for i := k; i < N; i++ {
		finishTime := heap.Pop(pq).(int) // 가장 먼저 끝난 소가 무대에서 내려온다

		// 다음 차례의 소가 무대에서 내려오는 시간이 T 시간보다 작거나 같은 경우
		if next := finishTime + cows[i]; next <= T {
			heap.Push(pq, next) // 다음 차례의 소가 무대에 오른다
		} else {
			return false // 크기가 k인 무대로는 공연을 성공적으로 진행할 수 없다
		}
	}

	// 모든 소들이 성공적으로 공연을 마친 경우
	return true
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
