package bj15988

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
	dp := make([]int, 1000001)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i <= 1000000; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000009
	}
	for i := 0; i < t; i++ {
		n := scanInt()
		fmt.Fprintln(writer, dp[n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
