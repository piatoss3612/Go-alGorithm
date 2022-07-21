package bj2705

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
	dp := make([]int, 1001)
	dp[1] = 1
	for i := 2; i <= 1000; i++ {
		dp[i] = 1 // i번째 항은 i를 기본 팰린드롬 수로 가지고 있다
		// 1부터 i/2번째 항 까지의 팰린드롬 수의 개수를 더하면 i번째 항의 값을 구할 수 있다
		for j := i / 2; j >= 1; j-- {
			dp[i] += dp[j]
		}
	}
	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		fmt.Fprintln(writer, dp[n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
