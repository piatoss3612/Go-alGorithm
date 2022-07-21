package bj9657

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
	dp[1] = 1
	dp[2] = 0
	dp[3] = 1
	dp[4] = 1
	for i := 5; i <= n; i++ {
		// SK가 먼저 1,3,4 중 선택하고
		// 남은 돌의 개수에 해당하는 결과가 모두 같은 경우(SK가 이긴 경우인 1과 같음)
		// SK가 게임을 이길 수 있는 경우가 없으므로 CY의 승리
		if dp[i-1] == 1 && dp[i-3] == 1 && dp[i-4] == 1 {
			dp[i] = 0
		} else {
			dp[i] = 1
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
