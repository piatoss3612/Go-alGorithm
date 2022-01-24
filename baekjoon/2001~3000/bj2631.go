package bj2631

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
	line := make([]int, n+1)
	for i := 1; i <= n; i++ {
		line[i] = scanInt()
	}
	dp := make([]int, n+1)
	dp[1] = 1
	// 가장 긴 부분 증가 수열을 찾아 n에서 그 길이 만큼을 차감하면
	// 위치를 옮기는 아이들의 최솟값을 구할 수 있다
	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if line[j] < line[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	sort.Ints(dp)
	fmt.Fprintln(writer, n-dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
