package bj2294

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
	s := make([]int, n+1)
	dp := make([]int, k+1)
	// 최솟값 비교를 위해 dp의 모든 값을 최댓값으로 할당
	for i := 0; i <= k; i++ {
		dp[i] = 10001
	}
	for i := 1; i <= n; i++ {
		s[i] = scanInt()
		// k보다 작거나 같은 입력값은 1을 할당
		if s[i] <= k {
			dp[s[i]] = 1
		}
	}

	for i := 1; i <= k; i++ {
		for _, v := range s {
			if v <= i {
				if v == i {
					break
				}
				// v에 해당하는 dp값과 i-v에 해당하는 dp값의 합이 최소가 되는 경우를 찾는다
				tmp := dp[v] + dp[i-v]
				if tmp < dp[i] {
					dp[i] = tmp
				}
			}
		}
	}
	// 최댓값 그대로인 경우 -1을 출력
	if dp[k] == 10001 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, dp[k])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
