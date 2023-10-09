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

	dp [51][51][51]int
)

// 난이도: Silver 2
// 메모리: 1464KB
// 시간: 16ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 20; j++ {
			for k := 0; k <= 20; k++ {
				dp[i][j][k] = calc(i, j, k)
			}
		}
	}
}

func Solve() {
	for {
		a, b, c := scanInt(), scanInt(), scanInt()
		if a == -1 && b == -1 && c == -1 {
			break
		}

		fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, calc(a, b, c))
	}
}

func calc(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}

	if a > 20 || b > 20 || c > 20 {
		return calc(20, 20, 20)
	}

	if dp[a][b][c] != 0 {
		return dp[a][b][c]
	}

	if a < b && b < c {
		dp[a][b][c] = calc(a, b, c-1) + calc(a, b-1, c-1) - calc(a, b-1, c)
		return dp[a][b][c]
	}

	dp[a][b][c] = calc(a-1, b, c) + calc(a-1, b-1, c) + calc(a-1, b, c-1) - calc(a-1, b-1, c-1)
	return dp[a][b][c]
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
