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

	R, C, K int
	grid    [101][101]int
	dp      [101][101]int
)

const MOD = 1000000007

// 난이도: Silver 1
// 메모리: 1124KB
// 시간: 168ms
// 분류: 다이나믹 프로그래밍, 그래프 이론, 그래프 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	R, C, K = scanInt(), scanInt(), scanInt()
	for i := 1; i <= R; i++ {
		for j := 1; j <= C; j++ {
			grid[i][j] = scanInt()
			dp[i][j] = -1
		}
	}
}

func Solve() {
	fmt.Fprintln(writer, dfs(1, 1))
}

func dfs(r, c int) int {
	// 기저 사례: R, C에 도달하면 1 반환
	if r == R && c == C {
		return 1
	}

	ret := &dp[r][c]

	// 방문 여부를 확인하고 방문했으면 그 값을 반환
	if *ret != -1 {
		return *ret
	}

	*ret = 0

	// 최소 하나의 행과 하나의 열을 이동해야 하므로 r+1, c+1부터 시작
	for i := r + 1; i <= R; i++ {
		for j := c + 1; j <= C; j++ {
			// 격자의 값이 다르면 이동
			if grid[i][j] != grid[r][c] {
				*ret = (*ret + dfs(i, j)) % MOD
			}
		}
	}

	return *ret
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
