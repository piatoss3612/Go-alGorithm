package bj11048

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
	n, m := scanInt(), scanInt()

	input := make([][]int, n+1)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		input[i] = make([]int, m+1)
		dp[i] = make([]int, m+1)
		if i != 0 {
			for j := 1; j <= m; j++ {
				input[i][j] = scanInt()
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			dp[i][j] = getMax(dp[i-1][j-1], getMax(dp[i-1][j], dp[i][j-1]))

			dp[i][j] += input[i][j]
		}
	}

	fmt.Fprintln(writer, dp[n][m])
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
