package bj14494

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
	var dp [1001][1001]int
	n, m := scanInt(), scanInt()
	dp[1][1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i*j != 1 {
				dp[i][j] = (dp[i][j-1] + dp[i-1][j] + dp[i-1][j-1]) % 1000000007
				// dp[i][j-1]: 오른쪽으로 이동한 경우
				// dp[i-1][j]: 아래로 이동한 경우
				// dp[i-1][j-1]: 아래오른쪽으로 이동한 경우
			}
		}
	}
	fmt.Fprintln(writer, dp[n][m])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
