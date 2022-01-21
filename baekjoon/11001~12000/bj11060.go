package bj11060

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
	input := make([]int, n+1)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
		dp[i] = n // dp의 모든 값을 최댓값 n으로 초기화
	}

	dp[1] = 0 // n = 1인 경우, 움직이지 않아도 되므로 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= input[i]; j++ {
			if i+j > n {
				break
			}
			// dp[i]에서 i + input[i]보다 작거나 같은 인덱스로 이동하면서
			// 최솟값을 찾는다
			tmp := dp[i] + 1
			if tmp < dp[i+j] {
				dp[i+j] = tmp
			}
		}
	}
	if dp[n] == n { // dp[n]이 n이면 dp[n]에 도달할 수 없는 경우이므로 -1 출력
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, dp[n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
