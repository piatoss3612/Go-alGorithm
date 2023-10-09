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
	conn    [1000001][]int // 연결 정보
	visited [1000001]bool  // 방문 여부
	// dp[i][0]: i번 노드가 얼리 어답터가 아닐 때 모든 노드가 새로운 아이디어를 수용하기 위해 필요한 얼리 어답터 수의 최솟값
	// dp[i][1]: i번 노드가 얼리 어답터일 때 ""
	dp [1000001][2]int
)

// 난이도: Gold 3
// 메모리: 119452KB
// 시간: 828ms
// 분류: 다이나믹 프로그래밍, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i < N; i++ {
		a, b := scanInt(), scanInt()
		conn[a] = append(conn[a], b)
		conn[b] = append(conn[b], a)
	}
}

func Solve() {
	// 사이클이 존재하지 않는 트리이므로 어디서 탐색을 시작하는지는 상관없다
	// 깊이 우선 탐색으로 반환되는 값 두 개 중에 최솟값을 출력한다
	fmt.Fprintln(writer, min(DFS(1)))
}

// 깊이 우선 탐색을 통해 node의 자식 노드들을 탐색하고
// node가 얼리 어답터가 아닌 경우와 얼리 어답터인 경우 각각의 경우에 대해
// 모든 노드가 새로운 아이디어를 받아들이기 위해 node의 자식 노드 중에 얼리 어답터가 되어야 하는 노드 수의 최솟값을 구한다
func DFS(node int) (int, int) {
	visited[node] = true
	dp[node][0] = 0
	dp[node][1] = 1 // node가 얼리 어답터인 경우, 자기 자신을 얼리 어답터 수에 포함

	for _, next := range conn[node] {
		// 아직 방문하지 않은 노드에 대해 깊이 우선 탐색
		if !visited[next] {
			// a: next가 얼리 어답터가 아닌 경우의 얼리 어답터 수의 최솟값
			// b: next가 얼리 어답터인 경우의 얼리 어답터 수의 최솟값
			a, b := DFS(next)
			dp[node][0] += b         // node가 얼리 어답터가 아니라면 next는 항상 얼리 어답터여야 하므로 b를 더해준다
			dp[node][1] += min(a, b) // node가 얼리 어답터라면 인접한 노드가 얼리 어답터이든 아니든 상관이 없으므로 a와 b 중 최솟값을 더해준다
		}
	}

	return dp[node][0], dp[node][1]
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
