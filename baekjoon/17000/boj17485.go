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
	N, M    int
	matrix  [1001][1001]int
	dp      [4][1001][1001]int // dp[k][i][j]: k 방향으로 움직여 (i, j)에 도달했을 때 달에 도착하기 위해 필요한 연료의 최솟값
)

const INF = 987654321

// 난이도: Gold 5
// 메모리: 57740KB
// 시간: 256ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			matrix[i][j] = scanInt()
		}
	}
}

func Solve() {
	ans := INF
	for i := 1; i <= M; i++ {
		ans = min(ans, rec(0, 0, i))
	}
	fmt.Fprintln(writer, ans)
}

func rec(moved, y, x int) int {
	// 기저 사례: 달에 도달한 경우
	if y == N {
		return 0
	}

	ret := &dp[moved][y][x]
	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최솟값 비교를 위해 INF로 초기화

	// 이전에 1번 방향(왼쪽)으로 움직이지 않은 경우
	if moved != 1 {
		ny, nx := y+1, x-1
		if valid(ny, nx) {
			*ret = min(*ret, rec(1, ny, nx)+matrix[ny][nx])
		}
	}

	// 이전에 2번 방향(중앙)으로 움직이지 않은 경우
	if moved != 2 {
		ny, nx := y+1, x
		if valid(ny, nx) {
			*ret = min(*ret, rec(2, ny, nx)+matrix[ny][nx])
		}
	}

	// 이전에 3번 방향(오른쪽)으로 움직이지 않은 경우
	if moved != 3 {
		ny, nx := y+1, x+1
		if valid(ny, nx) {
			*ret = min(*ret, rec(3, ny, nx)+matrix[ny][nx])
		}
	}

	return *ret
}

func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
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
