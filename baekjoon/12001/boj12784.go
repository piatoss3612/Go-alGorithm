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
	T, N, M int
	conn    [1001][]Conn
	visited [1001]bool
	dp      [1001]int
)

type Conn struct {
	y    int
	cost int
}

// 난이도: Gold 3
// 메모리: 4348KB
// 시간: 28ms
// 분류: 다이나믹 프로그래밍, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

// 1번 노드를 루트로 갖는 트리 구조에서 다른 모든 리프 노드로부터
// 루트 노드로 이어지는 경로를 끊는 최소 비용을 구한다
func Setup() {
	N, M = scanInt(), scanInt()
	conn = [1001][]Conn{}
	visited = [1001]bool{}
	dp = [1001]int{}
	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		conn[a] = append(conn[a], Conn{b, c})
		conn[b] = append(conn[b], Conn{a, c})
	}
}

func Solve() {
	DFS(1) // 1번 노드에서 깊이 우선 탐색 시작
	fmt.Fprintln(writer, dp[1])
}

func DFS(x int) {
	visited[x] = true // x번 노드 방문 처리
	for _, next := range conn[x] {
		// 아직 방문하지 않은 노드인 경우
		if !visited[next.y] {
			DFS(next.y) // 깊이 우선 탐색 재귀 호출

			// 리프 노드인 경우(dp[next.y] = 0)와 아닌 경우
			if dp[next.y] == 0 {
				dp[x] += next.cost // 리프 노드와 연결을 끊는다
			} else {
				// 1. x와 자식 노드(next.y)의 연결을 끊거나 (next.cost)
				// 2. 자식 노드(next.y)가 자신의 자식 노드들과의 연결을 끊거나 (dp[next.y])
				dp[x] += min(next.cost, dp[next.y])
			}
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
