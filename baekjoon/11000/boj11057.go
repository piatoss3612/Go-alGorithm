package bj11057

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
	n := scanInt()
	dp := make([][10]int, n+1)
	dp[1] = [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		dp[i][9] = dp[i-1][9] // 오르막 수에서 9 뒤에는 9만 가능
		for j := 8; j >= 0; j-- {
			// 오르막 수에서 j 뒤에는 j~9까지 올 수 있으므로 누적된 값을 초기화한다
			dp[i][j] = (dp[i][j+1] + dp[i-1][j]) % 10007
		}
	}
	result := 0
	for _, v := range dp[n] {
		result = (result + v) % 10007
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
