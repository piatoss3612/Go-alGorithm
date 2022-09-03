package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	tree     [100001][]int // 양방향 트리 연결 정보
	visited  [100001]bool  // 노드 방문 여부
	order    []int         // 정해진 DFS 방문 순서
	expected [100001]int   // 특정 노드의 예상 방문 순서
	dfs      []int         // 실제 DFS 방문 순서
)

// 메모리: 15996KB
// 시간: 132ms
// 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i < N; i++ {
		a, b := scanInt(), scanInt()
		// 양방향 그래프임에 주의!
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	order = make([]int, N+1)
	dfs = make([]int, 0, N+1)

	for i := 1; i <= N; i++ {
		order[i] = scanInt()
		expected[order[i]] = i
	}

	// 트리의 DFS 방문 순서는 항상 1부터 시작되어야 한다
	if order[1] != 1 {
		fmt.Fprintln(writer, 0)
		return
	}

	for k := 1; k <= N; k++ {
		// 각 노드의 예상되는 탐색 순서에 따라 트리의 연결 정보를 정렬
		sort.Slice(tree[k], func(i, j int) bool {
			return expected[tree[k][i]] < expected[tree[k][j]]
		})
	}

	dfs = append(dfs, 0)
	DFS(1)

	// 정해진 DFS 방문 순서와 실제 DFS 방문 순서를 비교
	for i := 1; i <= N; i++ {
		if order[i] != dfs[i] {
			fmt.Fprintln(writer, 0)
			return
		}
	}

	fmt.Fprintln(writer, 1)
}

func DFS(node int) {
	if visited[node] {
		return
	}

	visited[node] = true
	dfs = append(dfs, node)

	for _, child := range tree[node] {
		if !visited[child] {
			DFS(child)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
