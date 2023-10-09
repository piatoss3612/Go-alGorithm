package main

import (
	"bufio"
	"fmt"
	"math/big"
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
	m := scanInt()

	dp := make([][]*big.Int, n+1)
	dp[0] = []*big.Int{big.NewInt(1)}
	dp[1] = []*big.Int{big.NewInt(1), big.NewInt(1)}
	for i := 2; i <= n; i++ {
		comb := make([]*big.Int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				comb[j] = big.NewInt(1)
			} else {
				comb[j] = big.NewInt(0).Add(dp[i-1][j-1], dp[i-1][j])
			}
		}
		dp[i] = comb
	}
	fmt.Fprintln(writer, dp[n][m])
	// fmt.Fprintln(writer, big.NewInt(0).Binomial(int64(n), int64(m)))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
