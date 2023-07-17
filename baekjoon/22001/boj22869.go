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

	N, K int
	A    [5001]int
	dp   [5001]int
)

// 난이도: Silver 1
// 메모리: 1016KB
// 시간: 32ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		A[i] = scanInt()
	}
}

func Solve() {
	dp[1] = 1
	for i := 1; i <= N; i++ {
		// i번째 돌을 밟을 수 있는 경우가 없으면 다음 돌로 넘어간다.
		if dp[i] == 0 {
			continue
		}

		// i번째 돌을 밟을 수 있는 경우, i번째 돌에서 갈 수 있는 모든 돌을 체크한다.
		for j := i + 1; j <= N; j++ {
			temp := (j - i) * (1 + abs(A[i]-A[j]))
			if temp <= K {
				dp[j] = 1
			}
		}
	}

	if dp[N] > 0 {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
