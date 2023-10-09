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
	path    [100001][]int
	visited [100001]bool
	dp      [100001][2]int
)

// 난이도: Gold 3
// 메모리: 12096KB
// 시간: 88ms
// 분류: 다이나믹 프로그래밍, 트리, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	// N-1개의 양방향 도로
	for i := 1; i < N; i++ {
		u, v := scanInt(), scanInt()
		path[u] = append(path[u], v)
		path[v] = append(path[v], u)
	}
}

func Solve() {
	// 모든 도시들 사이에는 단 한개의 경로만이 존재 -> 스패닝 트리
	// 사이클이 존재하지 않으므로 임의의 위치(1)에서 깊이 우선 탐색 시작
	// 반환된 2개의 값 중 최솟값을 출력
	fmt.Fprintln(writer, min(DFS(1)))
}

func DFS(x int) (int, int) {
	visited[x] = true // 방문 처리
	dp[x][1] = 1      // x번 도시에 경찰서를 설치하는 경우

	// x번 도시와 연결된 도시들에 대한 탐색
	for _, next := range path[x] {
		// next번 도시를 아직 방문하지 않은 경우
		if !visited[next] {
			// next번 도시를 방문하여 깊이 우선 탐색 실행
			a, b := DFS(next)
			dp[x][0] += b         // x번 도시에 경찰서가 없다면 연결된 도시에는 반드시 경찰서가 설치되어 있어야 한다
			dp[x][1] += min(a, b) // x번 도시에 경찰서가 있다면 연결된 도시에는 경찰서가 있던 없던 상관없으므로 최솟값으로 가져온다
		}
	}
	return dp[x][0], dp[x][1] // x번 도시에 경찰서가 설치되어 있지 않은 경우 경찰서의 수의 최솟값, x번 도시에 경찰서가 설치되어 있는 경우 경찰서의 수의 최솟값
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
