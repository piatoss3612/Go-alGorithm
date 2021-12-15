package bj2579

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
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}
	if n == 1 {
		fmt.Fprintln(writer, input[n])
		return
	}
	dp[1] = input[1]
	dp[2] = input[2] + input[1]
	for i := 3; i <= n; i++ {
		dp[i] = getMax(dp[i-3]+input[i-1]+input[i], dp[i-2]+input[i])
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
