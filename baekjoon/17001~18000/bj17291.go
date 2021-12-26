package bj17291

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
	dp := make([]int, 21)
	dp[1] = 1
	dp[4] -= dp[1] // 1은 홀수이므로 1+3번째 값에서 1번째에서 생성된 세포의 수 1을 빼준다
	for i := 2; i <= 20; i++ {
		dp[i] += dp[i-1] * 2 // i번째 항의 값에 이전 항을 2배한 수를 더한다
		if i%2 == 0 {
			if i+4 <= 20 {
				dp[i+4] -= dp[i-1] // i가 짝수면 i+4번째에서 i번째에 생성된 세포의 수를 빼준다
			}
		} else {
			if i+3 <= 20 {
				dp[i+3] -= dp[i-1] // i가 홀수면 i+3번째에서 i번째에 생성된 세포의 수를 빼준다
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
