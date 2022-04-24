package bj1931

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
	input := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, 2)
		row[0], row[1] = scanInt(), scanInt()
		input[i] = row
	}
	// 끝나는 시간을 오름차순으로 정렬
	sort.Slice(input, func(i, j int) bool {
		// 끝나는 시간이 같다면 시작 시간이 빠른 순으로 정렬
		if input[i][1] == input[j][1] {
			return input[i][0] < input[j][0]
		}
		return input[i][1] < input[j][1]
	})

	dp := make([]int, n)
	dp[0] = 1
	idx := 0 // 끝나는 시간이 가장 작은 값의 인덱스
	for i := 1; i < n; i++ {
		dp[i] = 1
		// idx번째 항의 끝나는 시간이 i번째 항의 시작 시간보다 작거나 같다면
		if input[i][0] >= input[idx][1] {
			// idx번째 항에서의 최대 연속된 회의의 개수에 1을 더한 값을 dp[i]에 덧씌운다
			dp[i] = dp[idx] + 1
			idx = i
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
