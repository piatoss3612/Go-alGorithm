package bj1309

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
	dp[0] = 1
	dp[1] = 3
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1]*2 + dp[i-2]) % 9901
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
