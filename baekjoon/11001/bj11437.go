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
	n, m    int
	graph   [][]int // 연결 요소
	depth   []int   // 노드가 속한 트리의 깊이
	parent  []int   // 노드 i의 부모 노드
	visited []bool  // 노드 방문 여부
)

// 메모리: 13404KB
// 시간: 1048ms
// 트리를 구성하는 노드 a와 b의 공통된 조상 노드를 찾는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	n = scanInt()
	graph = make([][]int, n+1)

	// 처음에 주어지는 정보는 루트 노드가 1이라는 것을 제외하고
	// 어떤 노드가 부모고 자식인지 모르므로 모든 입력을 먼저 받아야 한다
	var a, b int
	for i := 1; i < n; i++ {
		a, b = scanInt(), scanInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	depth = make([]int, n+1)
	parent = make([]int, n+1)
	visited = make([]bool, n+1)

	// 모든 입력을 받은 후에, 각각의 노드가 트리에서 속한 깊이와 부모 노드를 찾는다
	Set(1, 0)

	m = scanInt()

	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		fmt.Fprintln(writer, LCA(a, b)) // 노드 a와 b의 최소 공통 조상 출력
	}
}

// 노드가 속한 트리의 깊이와 부모 노드를 구하는 함수
func Set(node, nodeDepth int) {
	depth[node] = nodeDepth // 현재 노드가 속한 깊이 초기화
	visited[node] = true    // 방문 여부를 체크함으로써 자식에서 부모로 루프가 발생하는 것을 방지

	// 연결된 요소 검사
	for _, n := range graph[node] {
		// 아직 방문하지 않은 노드, 즉 자식 노드인 경우
		if !visited[n] {
			parent[n] = node    // 자식 노드의 부모 노드 초기화
			Set(n, nodeDepth+1) // 자식 노드 탐색
		}
	}
}

// 최소 공통 조상을 구하는 함수
func LCA(a, b int) int {
	// a 노드와 b 노드의 깊이가 다른 경우
	for depth[a] != depth[b] {
		// 깊이가 더 깊은 노드를 부모 노드로 이동
		if depth[a] > depth[b] {
			a = parent[a]
		} else {
			b = parent[b]
		}
	}

	// 공통 조상을 아직 찾지 못한 경우
	// 동일한 깊이에서 동시에 부모 노드로 이동하며 공통 조상을 찾는다
	for a != b {
		a = parent[a]
		b = parent[b]
	}

	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
