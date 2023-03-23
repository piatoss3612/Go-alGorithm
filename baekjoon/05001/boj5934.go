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
	road    [50001][]int  // 인접 리스트
	visited [50001]bool   // 방문 여부
	dp      [50001][2]int // dp[i][0]: i를 방문하지 않았을 때, 만날 수 있는 소의 수의 최댓값, dp[i][1]: i를 방문했을 때, 만날 수 있는 소의 수의 최댓값
)

// 난이도: Gold 2
// 메모리: 6380KB
// 시간: 28ms
// 분류: 트리, 다이나믹 프로그래밍
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
		road[a] = append(road[a], b)
		road[b] = append(road[b], a)
	}
}

func Solve() {
	ans := max(DFS(1)) // 아무 노드나 시작점으로 잡아도 된다.
	fmt.Fprintln(writer, ans)
}

// 깊이 우선 탐색을 통해 v를 방문했을 때, 방문하지 않았을 때의 만날 수 있는 소의 수의 최댓값을 구한다.
func DFS(v int) (int, int) {
	visited[v] = true // 방문 표시
	dp[v][1] = 1      // v를 방문했을 때, 만날 수 있는 소의 수는 1

	for _, next := range road[v] {
		if !visited[next] {
			a, b := DFS(next)     // a: next를 방문하지 않았을 때 만날 수 있는 소의 수의 최댓값, b: next를 방문했을 때 만날 수 있는 소의 수의 최댓값
			dp[v][0] += max(a, b) // v를 방문하지 않은 경우 next를 방문하든 하지 않든 상관 없으므로 a와 b 중 큰 값을 더해준다.
			dp[v][1] += a         // v를 방문한 경우 next를 방문하지 않아야 하므로 a를 더해준다.
		}
	}

	return dp[v][0], dp[v][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
