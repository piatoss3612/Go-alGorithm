package bj14606

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

	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i/2] + dp[i-(i/2)] + (i/2)*(i-(i/2))
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
