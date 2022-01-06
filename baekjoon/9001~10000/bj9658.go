package bj9658

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
	// SK win: 1
	// CY win: 0
	dp[1] = 0
	dp[2] = 1
	dp[3] = 0
	dp[4] = 1
	for i := 5; i <= n; i++ {
		// SK가 먼저 게임을 시작하는데
		// 1, 3, 4 중 하나라도 자신이 이기는 경우를 선택하면 이길 수 있다
		// 따라서 값이 0인 경우가 하나라도 있다면 SK의 승리
		if dp[i-1] == 0 || dp[i-3] == 0 || dp[i-4] == 0 {
			dp[i] = 1
		} else {
			dp[i] = 0
		}
	}
	if dp[n] == 1 {
		fmt.Fprintln(writer, "SK")
	} else {
		fmt.Fprintln(writer, "CY")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
