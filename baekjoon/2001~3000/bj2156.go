package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	_ "math"
	_ "math/big"
	"os"
	_ "sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]int, 10001)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}
	dp := make([]int, 10001)
	dp[1] = input[1]         // 1번째 항에서의 최댓값
	dp[2] = dp[1] + input[2] // // 2번째 항에서의 최댓값

	for i := 3; i <= n; i++ {
		// i-2번째 항을 빼고 dp[i-3]까지의 최댓값과 i-1번째 항, 그리고 i번째 항을 더한 값과
		// i-1번째 항을 빼고 dp[i-2]까지의 최댓값과 i번째 항을 더한 값을 비교
		dp[i] = getMax(dp[i-3]+input[i-1]+input[i], dp[i-2]+input[i])
		// i-1번째 항까지의 최댓값과 비교
		dp[i] = getMax(dp[i-1], dp[i])
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
