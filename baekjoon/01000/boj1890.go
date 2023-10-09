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

	N     int
	board [101][101]int
	dp    [101][101]int
)

// 난이도: Silver 1
// 메모리: 960KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			board[i][j] = scanInt()
			dp[i][j] = -1
		}
	}
}

func Solve() {
	fmt.Fprintln(writer, rec(1, 1))
}

func rec(r, c int) int {
	if r == N && c == N {
		return 1
	}

	ret := &dp[r][c]
	// -1이 아니라 0으로 체크할 경우, 스택 호출이 너무 많아져서 메모리 초과가 발생한다.
	if *ret != -1 {
		return *ret
	}

	*ret = 0

	if r+board[r][c] <= N {
		*ret += rec(r+board[r][c], c)
	}

	if c+board[r][c] <= N {
		*ret += rec(r, c+board[r][c])
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
