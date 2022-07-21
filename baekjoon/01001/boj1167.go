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
	n       int
	edges   []Edges
	visited []bool // 방문 여부
	res     = 0    // 지름의 최댓값
	idx     int    // 임의의 노드 v에서 시작했을 때 지름의 길이가 최댓값이 되는 노드의 번호
)

type Edge struct { // 간선 정보: 연결된 노드, 거리
	node int
	dist int
}

type Edges []Edge // i번째 노드에 연결된 간선들을 저장

// 메모리: 14788KB
// 시간: 120ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	edges = make([]Edges, n+1)

	var a, b, c int
	for i := 1; i <= n; i++ {
		a = scanInt()
		b = scanInt()
		for b != -1 {
			c = scanInt()
			edges[a] = append(edges[a], Edge{b, c})
			b = scanInt()
		}
	}

	/*
		임의의 노드 v에서 시작했을 때 지름이 최댓값이 되는 노드 v1을 먼저 찾고
		다시 v1에서 깊이 우선 탐색을 실행하면 지름의 최댓값을 찾을 수 있다

		이것이 어떻게 가능한가?

		임의의 노드 v에서 시작하는 경로는 반드시 루트 노드를 거쳐 다른 리프 노드에 도달하게 된다
		v에서 루트 노드까지의 거리를 제외하면 결국 노드 v1까지의 거리는 루트 노드에서 리프 노드들까지의 거리 중 최댓값이 된다

		따라서 v1에서 깊이 우선 탐색을 재실행하면 리프 노드~루트 노드 사이의 거리중 최댓값 + 그 다음 최댓값이 된다
	*/
	visited = make([]bool, n+1)
	visited[1] = true
	dfs(1, 0)

	visited = make([]bool, n+1)
	visited[idx] = true
	dfs(idx, 0)

	fmt.Fprintln(writer, res)
}

// 깊이 우선 탐색
func dfs(node, dist int) {
	if dist > res {
		res = dist
		idx = node
	}

	for _, edge := range edges[node] {
		if !visited[edge.node] {
			visited[node] = true
			dfs(edge.node, dist+edge.dist)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
