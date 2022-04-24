package bj15817

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

type pipe struct {
	l int
	c int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, x := scanInt(), scanInt()
	input := make([]pipe, n+1)
	for i := 1; i <= n; i++ {
		input[i] = pipe{scanInt(), scanInt()}
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, x+1)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, x+1)
		for j := 0; j <= x; j++ {
			dp[i][j] += dp[i-1][j]
			if dp[i][j] != 0 {
				for k := 1; k <= input[i].c; k++ {
					if j+k*input[i].l <= x {
						dp[i][j+k*input[i].l] += dp[i-1][j]
					}
				}
			}
		}
	}
	fmt.Fprintln(writer, dp[n][x])

	/*
		예제 입력:
		3 20
		4 3
		6 3
		9 2

		dp:
		[1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		[1 0 0 0 1 0 0 0 1 0 0 0 1 0 0 0 0 0 0 0 0]
		[1 0 0 0 1 0 1 0 1 0 1 0 2 0 1 0 1 0 2 0 1]
		[1 0 0 0 1 0 1 0 1 1 1 0 2 1 1 1 1 1 3 1 1]

		예제 출력:
		1
	*/
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
