package bj19947

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
	h := scanInt()
	y := scanInt()

	dp := make([]int, y+1)
	dp[0] = h

	for i := 1; i <= y; i++ {
		dp[i] = int(float64(dp[i-1]) * 1.05)
		if i >= 3 {
			if int(float64(dp[i-3])*1.2) > dp[i] {
				dp[i] = int(float64(dp[i-3]) * 1.2)
			}
		}
		if i >= 5 {
			if int(float64(dp[i-5])*1.35) > dp[i] {
				dp[i] = int(float64(dp[i-5]) * 1.35)
			}
		}
	}
	fmt.Fprintln(writer, dp[y])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
