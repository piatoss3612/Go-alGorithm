package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M    int
	edges   [][]int
	visited []bool
)

// 24762번: Ticket Completed?
// hhttps://www.acmicpc.net/problem/24762
// 난이도: 골드 5
// 메모리: 10988 KB
// 시간: 48 ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색, 조합론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	edges = make([][]int, N+1)  // 연결된 간선 정보
	visited = make([]bool, N+1) // 방문 여부
	for i := 0; i < M; i++ {
		u, v := scanInt(), scanInt()
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}
}

func Solve() {
	if M == 0 { // 연결 가능한 간선이 없는 경우
		fmt.Fprintln(writer, 0)
		return
	}

	componentSizes := make([]int, 0) // 각 컴포넌트의 크기

	for i := 1; i <= N; i++ {
		if !visited[i] {
			size := bfs(i)
			componentSizes = append(componentSizes, size)
		}
	}

	connectablePairs := 0
	for _, size := range componentSizes {
		connectablePairs += size * (size - 1) / 2 // size C 2 -> 컴포넌트 내에서 연결 가능한 노드 쌍의 개수
	}

	totalTickets := N * (N - 1) / 2                                  // 전체 노드에서 연결 가능한 노드 쌍의 개수
	probability := float64(connectablePairs) / float64(totalTickets) // 연결 가능한 노드 쌍의 개수 / 전체 노드에서 연결 가능한 노드 쌍의 개수

	fmt.Fprintln(writer, probability)

}

func bfs(start int) int {
	queue := []int{start}
	visited[start] = true
	count := 0 // start 노드를 포함한 연결된 노드의 개수

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		count++

		for _, v := range edges[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return count
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
