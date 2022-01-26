package bj11058

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
	dp := make([]int, 101)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 4
	dp[5] = 5
	dp[6] = 6
	// 6번째 항까지는 단순히 A를 출력함으로써 최댓값에 도달
	for i := 7; i <= 100; i++ {
		/*
			ex) 7번째 항
			a = dp[3] * 3 = AAA - Ctrl-A - Ctrl-C - Ctrl-V - Ctrl-V
			b = dp[2] * 4 = AA - Ctrl-A - Ctrl-C - Ctrl-V - Ctrl-V - Ctrl-V
			c = dp[1] * 5 = A - Ctrl-A - Ctrl-C - Ctrl-V - Ctrl-V - Ctrl-V - Ctrl-V

			단순히 dp[i-4] * 3로 계산해도 될 것 같지만, 혹시 모를 예외 상황에 대비해
			a, b, c를 구해서 최댓값을 비교한다
		*/
		a := dp[i-4] * 3
		b := dp[i-5] * 4
		c := dp[i-6] * 5
		dp[i] = getMax(getMax(a, b), c)
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
