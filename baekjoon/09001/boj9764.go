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
	dp := make([][]int, 2001)
	for i := 0; i <= 2000; i++ {
		dp[i] = make([]int, 2001)

		if i == 0 {
			for j := 0; j <= 2000; j++ {
				dp[0][j] = 1
			}
		}
	}

	for i := 1; i <= 2000; i++ {
		for j := 1; j <= 2000; j++ {
			dp[i][j] = dp[i][j-1]
			if i >= j {
				dp[i][j] = (dp[i][j] + dp[i-j][j-1]) % 100999
			}
		}
	}

	t := scanInt()
	for i := 1; i <= t; i++ {
		n := scanInt()
		fmt.Fprintln(writer, dp[n][n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
