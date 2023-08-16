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

	N, M, X int
	rtree   [100001][]int // 방향이 뒤집힌 트리
	visited [100001]bool // 방문 여부
	ans     int // 방문해야 하는 상위 노드의 수
)

// 난이도: Silver 1
// 메모리: 16872KB
// 시간: 96ms
// 분류: 그래프 이론, 그래프 탐색, 트리, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 0; i < M; i++ {
		u, v := scanInt(), scanInt()
		// 중복된 정보가 존재하지 않고 전후 관계가 분명하므로 트리의 간선들의 방향을 뒤집어준다.
		rtree[v] = append(rtree[v], u)
	}
	X = scanInt()
}

func Solve() {
	// 노드 X에서 dfs를 수행하면 X의 상위 노드들을 모두 방문할 수 있다.
	dfs(X)
	fmt.Fprintln(writer, ans-1) // X 자신은 제외
}

func dfs(v int) {
	visited[v] = true
	ans += 1
	for _, u := range rtree[v] {
		if !visited[u] {
			dfs(u)
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
