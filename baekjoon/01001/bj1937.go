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
	graph   [501][501]int
	dp      [501][501]int
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
)

// 메모리: 36764KB
// 시간: 88ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			graph[i][j] = scanInt()
			dp[i][j] = -1 // 방문 여부를 확인하기 위해 -1로 초기화
		}
	}

	ans := 0

	// 어느 위치에서 대나무를 먹기 시작해야 최솟값을 얻을 수 있을지
	// 모르는 상태이므로 완전 탐색 수행
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ans = max(ans, rec(i, j))
		}
	}
	fmt.Fprintln(writer, ans)
}

// 깊이 우선 탐색
func rec(x, y int) int {
	ret := &dp[x][y]

	if *ret != -1 { // 방문 여부 확인
		return *ret
	}

	*ret = 1 // 현재 위치의 대나무를 먹는다
	var nx, ny int
	// 상, 하, 좌, 우 탐색
	for i := 0; i < 4; i++ {
		nx = x + dx[i]
		ny = y + dy[i]
		if valid(nx, ny) && graph[nx][ny] > graph[x][y] {
			*ret = max(*ret, rec(nx, ny)+1) // 최댓값 갱신
		}
	}
	return *ret
}

func valid(x, y int) bool {
	if x >= 1 && x <= n && y >= 1 && y <= n {
		return true
	}
	return false
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
