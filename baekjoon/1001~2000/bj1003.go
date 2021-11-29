package bj1003

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
		n := scanInt()
		if n == 0 {
			fmt.Fprintln(writer, 1, 0)
		} else if n == 1 {
			fmt.Fprintln(writer, 0, 1)
		} else {
			testFib(n)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func testFib(n int) {
	dp := make([][]int, n+1)
	dp[0] = []int{1, 0}
	dp[1] = []int{0, 1}

	for i := 2; i <= n; i++ {
		dp[i] = []int{dp[i-1][0] + dp[i-2][0], dp[i-1][1] + dp[i-2][1]}
	}

	fmt.Fprintln(writer, dp[n][0], dp[n][1])
}
