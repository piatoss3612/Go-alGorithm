package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	n, k    int
	friends []int
	cnt     []int
)

type Car struct {
	driver int
	moved  int
}

type PQ []*Car

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	if pq[i].moved == pq[j].moved {
		return pq[i].driver < pq[j].driver
	}
	return pq[i].moved < pq[j].moved
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Car))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// 난이도: Gold 4
// 메모리: 2240KB
// 시간: 12ms
// 분류: 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	n, k = scanInt(), scanInt()
	friends = make([]int, n)
	cnt = make([]int, n)
	for i := 0; i < n; i++ {
		friends[i] = scanInt()
	}
	sort.Ints(friends)
}

func Solve() {
	pq := new(PQ)
	// 운전 속도가 빠른 친구부터 k명을 뽑아서 운전시킨다.
	for i := 0; i < n && i < k; i++ {
		pq.Push(&Car{driver: i, moved: friends[i]}) // i번째 운전자가 최초로 스타디움까지 운전한 경우의 이동거리로 초기화
	}
	heap.Init(pq)

	ans := 0

	for n > 0 {
		car := heap.Pop(pq).(*Car)
		driver, moved := car.driver, car.moved

		ans = moved
		heap.Push(pq, &Car{driver: driver, moved: moved + friends[driver]*2}) // 다음에 이동할 경우의 이동거리로 갱신

		// 운전자가 아직 운전을 하지 않았다면 운전자를 포함하여 5명이 스타디움에 도착
		if cnt[driver] == 0 {
			n -= 5
		} else {
			n -= 4 // 운전자는 운전을 처음 시작할 때 카운팅되었으므로 운전자를 제외한 4명이 스타디움에 도착
		}
		cnt[driver] += 1
	}
	fmt.Fprintln(writer, ans)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
