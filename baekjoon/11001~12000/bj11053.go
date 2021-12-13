package bj11053

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
	dp[1] = 1 // 첫번째 항은 반드시 1
	for i := 2; i <= n; i++ {
		dp[i] = 1 // i번째 항보다 작은 값이 없을 경우 1
		// i번째 항보다 작은 값들과 비교하여 dp값이 최대인 항 + 1을 dp[i]에 저장
		for j := i - 1; j > 0; j-- {
			if input[j] < input[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	getMax(dp)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
func getMax(slice []int) {
	max := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	fmt.Fprintln(writer, max)
}
