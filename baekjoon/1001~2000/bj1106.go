package bj1106

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type plan struct {
	cost   int
	people int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	c, n := scanInt(), scanInt()

	input := make([]plan, n+1)
	for i := 1; i <= n; i++ {
		input[i] = plan{scanInt(), scanInt()}
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, c+1)
	}

	for i := 1; i <= n; i++ {
		cost := input[i].cost
		people := input[i].people

		for j := 0; j <= c; j += people {
			if j+people >= c {
				// 고객의 수가 초과되는 경우에도 최솟값이 될 수 있다
				dp[i][c] = getMin(dp[i][c], dp[i][j]+cost)
			} else {
				dp[i][j+people] = getMin(dp[i][j+people], dp[i][j]+cost)
			}
		}

		for j := 1; j <= c; j++ {
			dp[i][j] = getMin(dp[i][j], dp[i-1][j])
			if dp[i][j] > 0 {
				tmp := j + people
				if tmp <= c {
					dp[i][tmp] = getMin(dp[i][tmp], dp[i][j]+cost)
				} else {
					// 고객의 수가 c를 초과하는 경우에도 최솟값인지 검사
					dp[i][c] = getMin(dp[i][c], dp[i][j]+cost)
				}
			}
		}
	}

	fmt.Fprintln(writer, dp[n][c])
}

func getMin(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
