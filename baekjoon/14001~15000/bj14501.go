package bj14501

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
	t := make([]int, n+1)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t[i], p[i] = scanInt(), scanInt()
	}
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		idx := i + t[i] - 1 // i항에서 시작한 상담을 완료하는 날
		if idx > n {
			continue
		} else {
			// i항에서 시작하므로 이전의 항들 중 최댓값과 p[i]를 더한다
			tmp := getMax(dp[:i]) + p[i]
			if tmp > dp[idx] {
				dp[idx] = tmp
			}
		}
	}
	fmt.Fprintln(writer, getMax(dp))
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
