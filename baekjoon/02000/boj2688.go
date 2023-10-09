package bj2688

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
	dp := make([][10]int, 65)
	for i := 0; i <= 9; i++ {
		dp[1][i] = 1 // 자리 수가 1인 경우, 모두 1로  초기화
	}
	for i := 2; i <= 64; i++ {
		dp[i][9] = 1 // 9로 시작하는 수는 뒤에 9만 올 수 있으므로 경우의 수가 1개 밖에 없다
		for j := 8; j >= 0; j-- {
			// 3자리 수인 경우
			// 7 뒤에 올 수 있는 수: 77, 78, 79, 88, 89, 99
			// 2자리 수인 경우 7, 8, 9 뒤에 오는 경우의 수가 누적된 값
			// dp[3][7] = dp[2][7] + dp[2][8] + dp[2][9]
			// dp[3][9] = dp[2][9]
			// dp[3][8] = dp[2][8] + dp[2][9]
			// dp[3][7] = dp[2][7] + dp[3][8]
			// 따라서 dp[i][j] = dp[i - 1][j] + dp[i][j + 1]
			dp[i][j] = dp[i-1][j] + dp[i][j+1]
		}
	}

	t := scanInt()
	for i := 1; i <= t; i++ {
		n := scanInt()
		result := 0
		// 결과는 자리 수에 해당하는 행의 모든 값은 더한 것
		for _, v := range dp[n] {
			result += v
		}
		fmt.Fprintln(writer, result)
	}

}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
