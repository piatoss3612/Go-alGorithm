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
	parent  []int
)

// 28118번: 안전한 건설 계획
// https://www.acmicpc.net/problem/28118
// 난이도: 골드 4
// 메모리: 856 KB
// 시간: 8 ms
// 분류: 자료 구조, 그래프 이론, 그래프 탐색, 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	for i := 0; i < M; i++ {
		a, b := scanInt(), scanInt()
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pb] = pa
		}
	}
}

func Solve() {
	count := 0
	for i := 1; i <= N; i++ {
		if parent[i] == i {
			count++
		}
	}
	// 분리 집합의 개수 - 1 = 최소 비용
	// 비용이 1인 보강 작업을 하면 연결 요소의 개수를 최대 1개 줄일 수 있음
	// 마지막 연결은 비용이 0이므로 제외
	fmt.Fprintln(writer, count-1)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}

	parent[x] = find(parent[x])
	return parent[x]
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
