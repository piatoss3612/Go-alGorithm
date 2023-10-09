package bj13699

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	dp      = map[int]int{0: 1, 1: 1}
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	for i := 2; i <= n; i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			tmp += dp[j] * dp[i-j-1]
		}
		dp[i] = tmp
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
