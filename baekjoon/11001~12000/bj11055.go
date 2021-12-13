package bj11055

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
	dp[1] = input[1]
	for i := 2; i <= n; i++ {
		dp[i] = input[i]
		max := 0 // i번째 항보다 작은 값들 중 dp[j]의 최댓값을 저장
		for j := i - 1; j > 0; j-- {
			if input[j] < input[i] {
				if dp[j] > max {
					max = dp[j]
				}
			}
		}
		dp[i] += max
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
