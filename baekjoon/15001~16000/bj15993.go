package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	_ "math"
	_ "math/big"
	"os"
	_ "sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	dp := make([][]int, 10001)
	for i := 0; i <= 10000; i++ {
		dp[i] = make([]int, 10001)
	}

	dp[1][1] = 1
	dp[2][1] = 1
	dp[2][2] = 1
	dp[3][1] = 1
	dp[3][2] = 2
	dp[3][3] = 1

	for i := 4; i <= 10000; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = (dp[i-3][j-1] + dp[i-2][j-1] + dp[i-1][j-1]) % 1000000009
		}
	}

	t := scanInt()

	for i := 1; i <= t; i++ {
		n := scanInt()
		odd := 0
		even := 0
		for j := 1; j <= n; j++ {
			if j%2 == 0 {
				even = (even + dp[n][j]) % 1000000009
			} else {
				odd = (odd + dp[n][j]) % 1000000009
			}
		}
		fmt.Fprintf(writer, "%d %d\n", odd, even)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
