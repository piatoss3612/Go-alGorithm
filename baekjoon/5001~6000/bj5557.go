package bj5557

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
	dp := make([][]int, n+1) // 숫자의 개수 + 1 크기의 2차원 슬라이스
	for i := 0; i <= n; i++ {
		// 연산의 결과는 0 ~ 20이므로 인덱스가 0~20이 되도록 크기를 21로 설정
		dp[i] = make([]int, 21)
	}

	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	dp[1][input[1]] = 1 // dp[1][첫번째 입력값] = 1
	for i := 2; i <= n; i++ {
		// i - 1 번째까지의 입력값을 통해 만들 수 있는 등식의 누적 수를 사용해
		// i 번째 입력값과의 연산을 통해 만들 수 있는 등식의 수를 구한다
		for j := 0; j <= 20; j++ {
			if dp[i-1][j] != 0 {
				if valid(j - input[i]) {
					dp[i][j-input[i]] += dp[i-1][j]
				}
				if valid(j + input[i]) {
					dp[i][j+input[i]] += dp[i-1][j]
				}
			}
		}
	}
	fmt.Fprintln(writer, dp[n-1][input[n]])
}

func valid(x int) bool {
	return x <= 20 && x >= 0
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
