package bj18353

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
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = 1
		// 감소하는 부분 수열을 찾기 위해
		// i번째 항 이전에 i번째 항보다 큰 값이 나오는 경우
		// dp[j] + 1 값과 비교하여 더 큰 값이 가장 긴 감소하는 부분 수열이다
		for j := 1; j < i; j++ {
			if input[j] > input[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	fmt.Fprintln(writer, n-getMax(dp))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMax(slice []int) int {
	max := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
