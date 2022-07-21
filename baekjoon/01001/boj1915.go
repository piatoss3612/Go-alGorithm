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
	n, m := scanInt(), scanInt()
	dp := make([][]int, n+1)

	dp[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		dp[i] = append([]int{0}, scanByteToInt()...)
	}

	max := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if dp[i][j] > 0 && dp[i-1][j-1] > 0 && dp[i-1][j] > 0 && dp[i][j-1] > 0 {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	fmt.Fprintln(writer, max*max)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanByteToInt() []int {
	scanner.Scan()
	bts := scanner.Bytes()
	nums := make([]int, len(bts))
	for i := 0; i < len(bts); i++ {
		if bts[i] == '1' {
			nums[i] = 1
		}
	}
	return nums
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
