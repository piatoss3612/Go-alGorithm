package bj13398

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	// dp[n][0]: 1~n까지의 입력값들의 연속합의 최댓값
	// dp[n][1]: 1~n까지의 입력값들 중 하나를 제외한 연속합의 최댓값
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[1] = []int{input[1], input[1]}
	ans := input[1]

	for i := 2; i <= n; i++ {
		dp[i][0] = getMax(input[i], dp[i-1][0]+input[i])
		dp[i][1] = getMax(dp[i-1][0], dp[i-1][1]+input[i])
		ans = getMax(ans, getMax(dp[i][0], dp[i][1]))
	}
	fmt.Fprintln(writer, ans)
}

func getMax(a, b int) int {
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
