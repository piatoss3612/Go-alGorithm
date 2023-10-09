package bj1932

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
	input := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		row := make([]int, i+1)
		for j := 1; j <= i; j++ {
			row[j] = scanInt()
		}
		input[i] = row
	}
	dp := make([][]int, n+1)
	dp[1] = []int{0, input[1][1]}
	for i := 2; i <= n; i++ {
		row := make([]int, i+1)
		for j := 1; j <= i; j++ {
			if j == i {
				row[j] = input[i][j] + dp[i-1][j-1]
			} else {
				a := input[i][j] + dp[i-1][j]
				b := input[i][j] + dp[i-1][j-1]
				if a > b {
					row[j] = a
				} else {
					row[j] = b
				}
			}
		}
		dp[i] = row
	}
	getMax(dp[n])
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
