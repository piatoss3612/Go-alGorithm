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

	N  int
	A  [5001]int
	dp [5001]int
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 1044KB
// 시간: 48ms
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
		A[i] = scanInt()
		dp[i] = INF
	}
}

func Solve() {
	dp[1] = 0
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			temp := (i - j) * (1 + abs(A[i]-A[j]))
			temp = max(temp, dp[j]) // j에 도달할 때 필요한 힘 dp[j]와 j에서 i까지 가는데 필요한 힘을 비교하여 더 큰 값을 temp에 저장 (최소값의 최대값)
			dp[i] = min(dp[i], temp) // dp[i]를 최소값으로 갱신
		}
	}
	fmt.Fprintln(writer, dp[N])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
