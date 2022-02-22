package bj1495

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
	n, s, m := scanInt(), scanInt(), scanInt()

	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][s] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if dp[i-1][j] == true {
				if j+input[i] <= m {
					dp[i][j+input[i]] = true
				}
				if j-input[i] >= 0 {
					dp[i][j-input[i]] = true
				}
			}
		}
	}

	isOk := false
	ans := 0
	for i, v := range dp[n] {
		if v {
			isOk = true
			ans = i
		}
	}

	if isOk {
		fmt.Fprintln(writer, ans)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
