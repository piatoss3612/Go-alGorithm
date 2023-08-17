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

	N     int
	belts [101][]int
	visit [101]int
)

// 난이도: Silver 1
// 메모리: 936KB
// 시간: 4ms
// 분류: 그래프 이론, 그래프 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N-1; i++ {
		a, b := scanInt(), scanInt()
		belts[a] = append(belts[a], b)
	}
}

func Solve() {
	// 방향이 없는 스패닝 트리에서 방향 트리로 변경됨
	// 사이클이 존재하지 않으므로 방문 처리는 불필요
	// 모든 노드에서 dfs를 돌려서 다른 노드를 방문할 때마다 visit을 증가시킴
	// visit이 N이 되는 노드가 있다면 그 노드는 다른 어떤 노드에서 시작해도 방문할 수 있는 노드
	// visit이 N이 되는 노드가 없다면 -1 출력
	for i := 1; i <= N; i++ {
		dfs(i)
	}

	for i := 1; i <= N; i++ {
		if visit[i] == N {
			fmt.Fprintln(writer, i)
			return
		}
	}

	fmt.Fprintln(writer, -1)
}

func dfs(x int) {
	visit[x] += 1
	for _, y := range belts[x] {
		dfs(y)
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
