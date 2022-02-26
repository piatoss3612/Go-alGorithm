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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	dp := make([][]int, 100001)
	for i := 0; i <= 100000; i++ {
		dp[i] = make([]int, 2)
	}

	dp[1][1] = 1
	dp[2][0] = 1
	dp[2][1] = 1
	dp[3][0] = 2
	dp[3][1] = 2

	for i := 4; i <= 100000; i++ {
		dp[i][0] = (dp[i-3][1] + dp[i-2][1] + dp[i-1][1]) % 1000000009
		dp[i][1] = (dp[i-3][0] + dp[i-2][0] + dp[i-1][0]) % 1000000009
	}

	t := scanInt()

	for i := 0; i < t; i++ {
		n := scanInt()
		fmt.Fprintln(writer, dp[n][1], dp[n][0])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
