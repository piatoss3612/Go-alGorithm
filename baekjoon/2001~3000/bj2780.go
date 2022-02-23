package bj2780

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]int
	dp      [][]int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	graph = [][]int{
		{7},
		{2, 4},
		{1, 3, 5},
		{2, 6},
		{1, 5, 7},
		{2, 4, 6, 8},
		{3, 5, 9},
		{0, 4, 8},
		{5, 7, 9},
		{6, 8},
	}

	dp = make([][]int, 1001)
	for i := 1; i <= 1000; i++ {
		dp[i] = make([]int, 10)

		if i == 1 {
			for j := 0; j <= 9; j++ {
				dp[i][j] = 1
			}
			continue
		}

		for j := 0; j <= 9; j++ {
			for k := 0; k < len(graph[j]); k++ {
				dp[i][j] += (dp[i-1][graph[j][k]]) % 1234567
			}
		}
	}

	t := scanInt()

	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	n := scanInt()

	ans := 0

	for i := 0; i <= 9; i++ {
		ans += dp[n][i]
	}

	fmt.Fprintln(writer, ans%1234567)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
