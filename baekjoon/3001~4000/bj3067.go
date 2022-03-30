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
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	n := scanInt()
	coins := make([]int, n+1)
	for i := 1; i <= n; i++ {
		coins[i] = scanInt()
	}
	m := scanInt()

	dp := make([]int, m+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := coins[i]; j <= m; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	fmt.Fprintln(writer, dp[m])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
