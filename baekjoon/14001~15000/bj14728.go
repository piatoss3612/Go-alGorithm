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
)

type Study struct {
	time  int
	score int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, t := scanInt(), scanInt()

	input := make([]Study, n+1)
	for i := 1; i <= n; i++ {
		input[i] = Study{scanInt(), scanInt()}
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, t+1)
	}

	for i := 1; i <= n; i++ {
		time, score := input[i].time, input[i].score
		for j := 0; j <= t; j++ {
			if j-time >= 0 {
				dp[i][j] = getMax(dp[i-1][j], dp[i-1][j-time]+score)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Fprintln(writer, dp[n][t])
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
