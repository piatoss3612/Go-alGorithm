package bj15489

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
	r, c, w := scanInt(), scanInt(), scanInt()

	n := r + w - 1
	dp := make([][]int, n)

	dp[0] = []int{1}

	for i := 1; i < n; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				row[j] = 1
			} else if j == i {
				row[j] = 1
			} else {
				row[j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
		dp[i] = row
	}
	result := 0

	end := c
	for k := r - 1; k < n; k++ {
		result += sum(dp[k][c-1 : end])
		end += 1
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func sum(slice []int) int {
	tmp := 0
	for _, v := range slice {
		tmp += v
	}
	return tmp
}
