package bj17953

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

	var input [11][100001]int
	var dp [11][1000001]int

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			input[i][j] = scanInt()
		}
	}

	for i := 1; i <= m; i++ {
		dp[i][1] = input[i][1]
	}

	// i번째 열의 j번째 값을 i-1번째 열의 값들 중
	// 최댓값(j == k인 경우는 만족도 반감)과 누적하여 더한다

	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 1; k <= m; k++ {
				if j == k {
					dp[j][i] = getMax(dp[j][i], dp[k][i-1]+input[j][i]/2)
				} else {
					dp[j][i] = getMax(dp[j][i], dp[k][i-1]+input[j][i])
				}
			}
		}
	}

	ans := 0

	for i := 0; i <= m; i++ {
		ans = getMax(ans, dp[i][n])
	}

	fmt.Fprintln(writer, ans)
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
