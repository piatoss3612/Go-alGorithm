package bj10844

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
	dp[1] = [10]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][1] // 0은 앞에 오는 수가 1 밖에 없음
		dp[i][9] = dp[i-1][8] // 9는 앞에 오는 수가 8 밖에 없음
		for j := 1; j <= 8; j++ {
			// j 앞에 올 수 있는 수들은 j - 1과 j + 1이므로
			// 이것들의 누적된 갯수를 더하면 된다
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]%1000000000
		}
	}
	result := 0
	for _, v := range dp[n] {
		result = (result + v) % 1000000000
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
