package bj11568

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if input[j] < input[i] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
	}

	sort.Ints(dp)
	fmt.Fprintln(writer, dp[n])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
