package bj17845

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

type Study struct {
	prior int
	time  int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k := scanInt(), scanInt()
	input := make([]Study, k+1)
	for i := 1; i <= k; i++ {
		input[i] = Study{scanInt(), scanInt()}
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= k; i++ {
		prior, time := input[i].prior, input[i].time
		for j := 0; j <= n; j++ {
			if j-time >= 0 {
				dp[i][j] = getMax(dp[i-1][j], dp[i-1][j-time]+prior)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Fprintln(writer, dp[k][n])
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
