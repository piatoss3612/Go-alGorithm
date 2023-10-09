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
	village [10001]int
	path    [10001][]int
	visited [10001]bool
	dp      [10001][2]int
)

// 난이도: Gold 2
// 메모리: 3564KB
// 시간: 8ms
// 분류: 다이나믹 프로그래밍, 트리, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	// N개의 마을의 주민 수
	for i := 1; i <= N; i++ {
		village[i] = scanInt()
	}
	// N-1개의 길
	for i := 1; i < N; i++ {
		a, b := scanInt(), scanInt()
		path[a] = append(path[a], b)
		path[b] = append(path[b], a)
	}
}

func Solve() {
	// 문제 조건1: '우수 마을'로 선정된 마을 주민 수의 총 합의 최댓값을 구한다
	// 사이클이 존재하지 않으므로 임의의 마을(1)을 선택하여 깊이 우선 탐색 실행
	fmt.Fprintln(writer, max(DFS(1)))
}

func DFS(x int) (int, int) {
	visited[x] = true     // x번 마을 방문 처리
	dp[x][1] = village[x] // x번 마을을 우수 마을로 선정하는 경우

	for _, next := range path[x] {
		// next번 마을을 아직 방문하지 않은 경우
		if !visited[next] {
			a, b := DFS(next) // next번 마을을 방문하여 깊이 우선 탐색 실행

			// 문제 조건2: '우수 마을'끼리는 서로 인접해 있을 수 없다
			// 따라서 x번 마을을 우수 마을로 선정한 경우는 인접한 마을이 우수 마을이 아닐 때의 우수 마을 주민 수의 총 합의 최댓값을 누적해준다
			dp[x][1] += a

			// 문제 조건3: '우수 마을'로 선정되지 못한 마을은 적어도 하나의 '우수 마을'과는 인접해 있어야 한다
			// 이 문제 조건을 어떻게 충족시켜야 할지 고민을 해보았는데
			// 주민 수가 음수가 될 수 없으므로 최댓값을 찾는 과정에서 이 조건은 항상 충족된다는 것을 알 수 있었다
			dp[x][0] += max(a, b)
		}
	}

	return dp[x][0], dp[x][1]
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
