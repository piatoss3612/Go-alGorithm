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

	N, Q    int
	tree    [][]int
	visited [1001]bool
	dp      []int
)

// 난이도: Gold 4
// 메모리: 1044KB
// 시간: 12ms
// 분류: 깊이 우선 탐색, 트리
// 시간복잡도: O(N*Q) -> Q: 쿼리의 개수, N: 간선의 개수 (N-1)
// 공간복잡도: O(N*N) -> 양방향 그래프이므로 N*N
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	tree = make([][]int, N+1)
	dp = make([]int, N+1)

	// N-1개의 간선을 입력받아 트리를 구성, 양방향 그래프이므로 양쪽에 모두 추가
	for i := 1; i <= N-1; i++ {
		a, b := scanInt(), scanInt()
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	Q = scanInt()
}

func Solve() {
	for i := 1; i <= Q; i++ {
		q := scanInt()
		switch q {
		// 1번 쿼리
		case 1:
			u, v := scanInt(), scanInt()
			visited = [1001]bool{}
			_ = dfs(u, v, 0)
		// 2번 쿼리
		case 2:
			x := scanInt()
			fmt.Fprintln(writer, dp[x])
		}
	}
}

func dfs(here, target, turn int) bool {
	visited[here] = true

	if here == target {
		return true
	}

	result := false

	for _, next := range tree[here] {
		// 방문하지 않은 노드만 탐색
		if !visited[next] {
			// 목표 노드를 찾을 수 있는 경로인 경우에만 dp 갱신
			if dfs(next, target, turn+1) {
				result = true
				dp[next] += turn + 1
			}
		}
	}

	return result
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
