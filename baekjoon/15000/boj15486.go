package bj15486

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
	max := dp[0] // i번째 항 이전까지 누적합 중에 최댓값을 저장
	for i := 1; i <= n; i++ {
		// i번째 항에서 일을 맡은 경우에 끝나게 되는 날
		end := i + t[i] - 1
		if end <= n {
			tmp := max + p[i] // i번째 항의 수익과 이전의 최대 누적합을 더한 값
			if tmp > dp[end] {
				dp[end] = tmp
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Fprintln(writer, max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
