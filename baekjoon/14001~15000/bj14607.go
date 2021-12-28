package bj14607

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
	// 다이나믹 프로그래밍으로 풀면 메모리 초과 발생
	// dp := make([]int, n + 1)
	// dp[1] = 0
	// for i := 2; i <= n; i++ {
	//   a := i / 2
	//   b := i - a
	//   dp[i] = a * b + dp[a] + dp[b]
	// }
	// fmt.Fprintln(writer, dp[n])
	fmt.Fprintln(writer, n*(n-1)/2)
}
func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
