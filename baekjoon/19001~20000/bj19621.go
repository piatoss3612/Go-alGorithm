package bj19621

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
	input := make([][4]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = [4]int{0, scanInt(), scanInt(), scanInt()}
	}
	dp := make([]int, n+1)
	dp[1] = input[1][3]
	for i := 2; i <= n; i++ {
		max := 0
		// i번째 항과 1-1번째 항은 시간이 겹치므로
		// 1번째 항부터 i-2번째 항까지 비교하여
		// 누적된 이용자 수가 가장 큰 값과 i번째 항의 이용자 수를 더한다
		for j := 1; j <= i-2; j++ {
			if dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + input[i][3]
	}
	sort.Ints(dp)
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
