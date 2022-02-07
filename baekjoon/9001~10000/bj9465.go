package bj9465

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
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	n := scanInt()
	input := make([][]int, 2)
	dp := make([][]int, 2)
	for i := 0; i <= 1; i++ {
		input[i] = make([]int, n+1)
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			input[i][j] = scanInt()
		}
	}

	dp[0][1] = input[0][1]
	dp[1][1] = input[1][1]

	for i := 1; i <= n; i++ {

		if i+1 <= n {
			dp[0][i+1] = getMax(dp[0][i+1], dp[1][i]+input[0][i+1])
			dp[1][i+1] = getMax(dp[1][i+1], dp[0][i]+input[1][i+1])
		}

		if i+2 <= n {
			dp[0][i+2] = getMax(dp[0][i+2], dp[0][i]+input[0][i+2])
			dp[1][i+2] = getMax(dp[1][i+2], dp[0][i]+input[1][i+2])
			dp[0][i+2] = getMax(dp[0][i+2], dp[1][i]+input[0][i+2])
			dp[1][i+2] = getMax(dp[1][i+2], dp[1][i]+input[1][i+2])
		}
	}
	fmt.Fprintln(writer, getMax(dp[0][n], dp[1][n]))
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
