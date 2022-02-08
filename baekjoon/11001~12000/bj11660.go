package bj11660

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
	table := make([][]int, n+1)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]int, n+1)
		dp[i] = make([]int, n+1)
	}

	// 전체 누적 합 구하기
	for i := 1; i <= n; i++ {
		tmp := 0
		for j := 1; j <= n; j++ {
			table[i][j] = scanInt()
			tmp += table[i][j-1]
			dp[i][j] = table[i][j] + dp[i-1][j] + tmp
		}
	}

	for i := 1; i <= m; i++ {
		x1, y1 := scanInt(), scanInt()
		x2, y2 := scanInt(), scanInt()

		// x2, y2까지의 전체 누적 합에서 x2, y1-1까지의 누적 합과 x1-1,y2까지의 누적합을 뺀 후
		// 중복해서 뺀 x1-1, y1-1까지의 합을 더해주면 x1,y1부터 x2,y2까지의 누적합을 구할 수 있다
		result := dp[x2][y2] - dp[x2][y1-1] - dp[x1-1][y2] + dp[x1-1][y1-1]
		fmt.Fprintln(writer, result)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
