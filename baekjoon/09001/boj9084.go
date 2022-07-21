package bj9084

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
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	m := scanInt()

	dp := make([]int, m+1)

	dp[0] = 1 // j가 input[i]일 때의 경우의 수

	for i := 1; i <= n; i++ {
		for j := input[i]; j <= m; j++ {
			// j - input[i]가 0보다 크거나 같은 경우
			// ex) input[i] = 2, dp[2] += dp[2 - 2]
			dp[j] += dp[j-input[i]]
		}
	}
	fmt.Fprintln(writer, dp[m])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
