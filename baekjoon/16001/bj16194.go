package bj16194

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
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = scanInt()
	}
	// 1 ~ n번째 항까지 각각의 최솟값을 저장하는 배열
	dp := make([]int, n+1)
	dp[1] = p[1]
	for i := 2; i <= n; i++ {
		dp[i] = p[i] // i번째 항의 초깃값은 p[i]
		// i 번째 항의 최솟값을 구하기 위해
		// i - 1부터 시작하여 1까지 역순회
		for j := i - 1; j >= 1; j-- {
			// tmp는 dp[i]의 최솟값일 수도 있는 값
			tmp := dp[j] + dp[i-j]
			if tmp < dp[i] {
				dp[i] = tmp
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
