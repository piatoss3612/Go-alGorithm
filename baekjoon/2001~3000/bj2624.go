package bj2624

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type coin struct {
	p int
	n int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	t, k := scanInt(), scanInt()
	input := make([]coin, k+1)

	for i := 1; i <= k; i++ {
		input[i] = coin{scanInt(), scanInt()}
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i].p < input[j].p
	})

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, t+1)
	}

	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		for j := 0; j <= input[i].n; j++ {
			for k := 0; k <= t; k++ {
				tmp := input[i].p*j + k
				if tmp > t {
					break
				}
				dp[i][tmp] += dp[i-1][k]
			}
		}
	}

	fmt.Fprintln(writer, dp[k][t])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
