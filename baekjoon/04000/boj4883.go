package bj4883

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
	k := 0
	for {
		n := scanInt()
		if n == 0 {
			break
		}
		k += 1
		input := make([][4]int64, n+1)
		for i := 1; i <= n; i++ {
			for j := 1; j <= 3; j++ {
				input[i][j] = scanInt64()
			}
		}
		dp := make([][4]int64, n+1)
		// dp[1][1]: 사용하지 않을 것이므로 큰 수로 초기화
		// dp[1][2]: 첫 번째 행의 가운데 정점이 이동할 수 있는 점이므로 이동한 결과로 초기화
		dp[1] = [4]int64{0, 1000000, input[1][2], input[1][2] + input[1][3]}
		for i := 2; i <= n; i++ {
			dp[i][1] = input[i][1] + getMin(dp[i-1][1], dp[i-1][2])
			dp[i][2] = input[i][2] + getMin(getMin(dp[i][1], dp[i-1][1]), getMin(dp[i-1][2], dp[i-1][3]))
			dp[i][3] = input[i][3] + getMin(getMin(dp[i][2], dp[i-1][2]), dp[i-1][3])
		}
		fmt.Fprintf(writer, "%d. %d\n", k, dp[n][2])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanInt64() int64 {
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return n
}

func getMin(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
