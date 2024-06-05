package main

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
	N, M    int
	code    []int
	roads   [][]Road
	isPrime [10000001]bool
)

type Road struct {
	to, cost int
}

type Node struct {
	idx, cost int
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

const INF = math.MaxInt64 // INF 값을 크게 설정!

// 30974번: What's your ETA?
// hhttps://www.acmicpc.net/problem/30974
// 난이도: Gold 4
// 메모리: 124236 KB
// 시간: 988 ms
// 분류: 데이크스트라, 최단 경로, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	// N개의 버스 정류장, M개의 양방향 도로
	N, M = scanInt(), scanInt()
	// 각 정류장의 재난코드 값
	code = make([]int, N+1)
	for i := 1; i <= N; i++ {
		code[i] = scanInt()
	}

	// 소수 판별
	for i := 2; i <= 10000000; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= 10000000; i++ {
		if isPrime[i] {
			for j := i * i; j <= 10000000; j += i {
				isPrime[j] = false
			}
		}
	}

	// 도로 정보
	roads = make([][]Road, N+1)

	for i := 0; i < M; i++ {
		// 양방향 도로
		a, b, c := scanInt(), scanInt(), scanInt()

		codesum := code[a] + code[b]
		// 두 정류장의 재난코드 합이 소수인 경우에만 사용 가능한 도로로 판단
		if isPrime[codesum] {
			roads[a] = append(roads[a], Road{to: b, cost: c})
			roads[b] = append(roads[b], Road{to: a, cost: c})
		}
	}
}

func Solve() {
	pq := new(PQ)
	heap.Init(pq)

	costs := make([]int, N+1)
	for i := 1; i <= N; i++ {
		costs[i] = INF
	}

	costs[1] = 0
	heap.Push(pq, &Node{idx: 1, cost: 0})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*Node)

		if node.cost > costs[node.idx] {
			continue
		}

		for _, road := range roads[node.idx] {
			nextCost := node.cost + road.cost
			if nextCost < costs[road.to] {
				costs[road.to] = nextCost
				heap.Push(pq, &Node{idx: road.to, cost: nextCost})
			}
		}
	}

	if costs[N] == INF {
		fmt.Fprintln(writer, "Now where are you?")
	} else {
		fmt.Fprintln(writer, costs[N])
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
