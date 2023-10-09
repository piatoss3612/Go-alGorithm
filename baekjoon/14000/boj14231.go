package bj14231

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		max := 0
		// i번째 항의 값보다 작은 이전 항들 중에서
		// 최대 박스의 갯수를 구해서 1은 더한 값을 할당
		for j := 1; j < i; j++ {
			if input[i] > input[j] {
				if dp[j] > max {
					max = dp[j]
				}
			}
		}
		dp[i] = max + 1
	}
	sort.Ints(dp)
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
