package bj2293

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
	n, k := scanInt(), scanInt()
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}
	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1 // j에서 input[i] 자기 자신을 빼는 경우
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j] // j(1~k)가 input[i]보다 작은 경우
			if j >= input[i] {
				// 테이블을 작성하다 보니 점화식이 이렇게 나옴
				dp[i][j] = dp[i][j-input[i]] + dp[i-1][j]
			}
		}
	}
	fmt.Fprintln(writer, dp[n][k])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
