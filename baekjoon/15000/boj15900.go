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

	N       int
	tree    [500001][]int // 트리
	depth   [500001]int // 깊이
	visited [500001]bool // 방문 여부
	isLeaf  [500001]bool // 리프 노드 여부
)

// 난이도: Silver 1
// 메모리: 84064KB
// 시간: 376ms
// 분류: 트리, 그래프 이론, 그래프 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N-1; i++ {
		u, v := scanInt(), scanInt()
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}
}

func Solve() {
	dfs(1, 0) // 루트 노드인 1번 노드와 연결된 노드들의 깊이를 구한다.

	sum := 0 // 리프 노드의 깊이의 합

	for i := 2; i <= N; i++ {
		if isLeaf[i] {
			sum += depth[i]
		}
	}

	// 리프 노드의 깊이의 합이 짝수면 게임을 먼저 시작한 사람이 지고, 홀수면 이긴다.
	if sum%2 == 0 {
		fmt.Fprintln(writer, "No")
	} else {
		fmt.Fprintln(writer, "Yes")
	}
}

func dfs(v, d int) {
	visited[v] = true
	depth[v] = d

	cnt := 0 // 방문하지 않은 자식 노드의 개수

	for _, w := range tree[v] {
		if !visited[w] {
			cnt++
			dfs(w, d+1)
		}
	}

	// 방문하지 않은 자식 노드가 없으면 리프 노드이다.
	if cnt == 0 {
		isLeaf[v] = true
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
