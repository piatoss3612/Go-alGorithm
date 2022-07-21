package bj1149

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
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, 3)
		row[0] = scanInt()
		row[1] = scanInt()
		row[2] = scanInt()
		dp[i] = row
	}

	for i := 1; i < n; i++ {
		// 색깔 중복을 방지하기 위해 이전 행의 다른 익데스에서 최솟값을 찾아서 더한다
		dp[i][0] += getMin(dp[i-1][1], dp[i-1][2])
		dp[i][1] += getMin(dp[i-1][0], dp[i-1][2])
		dp[i][2] += getMin(dp[i-1][0], dp[i-1][1])
	}
	sort.Ints(dp[n-1])
	fmt.Fprintln(writer, dp[n-1][0])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
