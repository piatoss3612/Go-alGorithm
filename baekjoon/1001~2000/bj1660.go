package bj1660

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
	s := make([]int, 121) // 30만 개 이하로 만들 수 있는 모든 사면체
	for i := 1; i <= 120; i++ {
		s[i] = i * (i + 1) * (i + 2) / 6
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for _, v := range s {
			if v <= i {
				// 사면체를 구성하는 대포의 개수와 i가 같으면 dp[i] = 1
				if i == v {
					dp[i] = 1
					break
				}
				// 사면체 값에 해당하는 dp값 + 현재 i에서 사면체 값을 뺀 dp값이 최소가 되는 tmp를 구하면 된다
				tmp := dp[v] + dp[i-v]
				if tmp < dp[i] {
					dp[i] = tmp
				}
			} else {
				break
			}
		}
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
