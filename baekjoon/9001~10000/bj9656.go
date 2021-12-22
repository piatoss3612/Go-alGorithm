package bj9656

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
	dp := make([]int, 1001)
	// 1 = SK loose
	// 0 = CY loose
	dp[1] = 1
	dp[2] = 0
	dp[3] = 1
	for i := 4; i <= n; i++ {
		// 돌 하나를 가져가거나 3개를 가져갔을 때 모두 패배하는 경우
		if dp[i-3] == 1 && dp[i-1] == 1 {
			dp[i] = 0
		} else {
			dp[i] = 1
		}
	}
	if dp[n] == 1 {
		fmt.Fprintln(writer, "CY")
	} else {
		fmt.Fprintln(writer, "SK")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
