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
	tree    [100000][]int // 인접한 노드들을 저장
	visited [100000]bool  // 방문 여부를 저장
	dp      [100000]int   // 각 노드에서의 최대값을 저장
)

// 난이도: Gold 4
// 메모리: 25128KB
// 시간: 84ms
// 분류: 깊이 우선 탐색, 다이나믹 프로그래밍, 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	n = scanInt()
	for i := 0; i < n-1; i++ {
		a, b := scanInt(), scanInt()
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}
	for i := 0; i < n; i++ {
		dp[i] = scanInt()
	}
}

func Solve() {
	fmt.Fprintln(writer, DFS(0)) // 정점 0에서 시작
}

// DFS를 이용하여 각 노드에서의 최대값을 구한다.
func DFS(x int) int {
	visited[x] = true // 방문 표시

	// 인접한 노드들을 방문하지 않았다면, DFS를 이용하여 최대값을 구한다.
	for _, next := range tree[x] {
		if !visited[next] {
			dp[x] = max(dp[x], dp[x]+DFS(next)) // 현재 노드의 최댓값 갱신
		}
	}

	return dp[x]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
