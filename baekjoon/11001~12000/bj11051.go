package bj11051

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
	dp := make([][]int, n+1)
	dp[1] = append(dp[1], 1, 1)

	// 조합 5C2의 결과는 4C1 + 4C2
	// 다이나믹 프로그래밍을 통해 팩토리얼을 계산하지 않고도 구할 수 있다
	for i := 2; i <= n; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = (dp[i-1][j-1] + dp[i-1][j]) % 10007
			}
		}
		dp[i] = row
	}
	fmt.Fprintln(writer, dp[n][k])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
