package bj14430

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
	input := make([][]int, n+1)
	input[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		row := make([]int, m+1)
		for j := 1; j <= m; j++ {
			row[j] = scanInt()
		}
		input[i] = row
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = input[i]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] += dp[i][j-1]
			} else {
				dp[i][j] += dp[i-1][j]
			}
		}
	}
	getMax(dp)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMax(slice [][]int) {
	max := 0
	for _, v1 := range slice {
		for _, v2 := range v1 {
			if v2 > max {
				max = v2
			}
		}
	}
	fmt.Fprintln(writer, max)
}
