package bj2502

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
	d, k := scanInt(), scanInt()
	dp := make([][2]int, d+1)
	// 피보나치 수열이므로 모든 항은 a와 b의 갯수로 나타낼 수 있다
	dp[1] = [2]int{1, 0}
	dp[2] = [2]int{0, 1}
	for i := 3; i <= d; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-2][0]
		dp[i][1] = dp[i-1][1] + dp[i-2][1]
	}
	n, m := dp[d][0], dp[d][1]

	// k = a*n + b*m 인데, a와 b는 정수이므로
	// (a*n - k) / m = b, (a*n - k) % m = 0 이 성립된다
	a := 1
	for {
		tmp1 := a * n
		tmp2 := k - tmp1
		if tmp2%m == 0 {
			fmt.Fprintf(writer, "%d\n%d\n", a, tmp2/m)
			break
		}
		a += 1
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
