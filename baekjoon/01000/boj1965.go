package bj1965

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
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if input[j] < input[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	getMax(dp)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func getMax(slice []int) {
	max := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	fmt.Fprintln(writer, max)
}
