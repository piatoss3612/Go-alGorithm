package bj12865

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
	input := make([][3]int, n+1)
	for i := 1; i <= n; i++ {
		input[i][1], input[i][2] = scanInt(), scanInt()
	}
	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			w := input[i][1]
			v := input[i][2]
			dp[i][j] = dp[i-1][j]
			if j >= w {
				tmp := dp[i-1][j-w] + v
				if tmp > dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	fmt.Fprintln(writer, dp[n][k])

	/*
		입력:
		4 7
		6 13
		4 8
		3 6
		5 12

		dp:
		[0 0 0 0 0 0 0 0]
		[0 0 0 0 0 0 13 13] 첫 번째 항을 선택한 경우, 1 ~ 최대 무게 k까지의 최댓값
		[0 0 0 0 8 8 13 13]
		[0 0 0 6 8 8 13 14]
		[0 0 0 6 8 12 13 14] 첫 번째 ~ n 번째 항을 조합한 값의 누적된 최댓값

		출력 - dp[n][k]:
		14
	*/

}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
