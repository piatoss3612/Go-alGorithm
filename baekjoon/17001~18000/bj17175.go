package bj17175

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

	if n == 0 || n == 1 {
		fmt.Fprintln(writer, 1)
		return
	}

	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(1)
	dp[1] = big.NewInt(1)
	for i := 2; i <= n; i++ {
		dp[i] = new(big.Int).Add(dp[i-1], dp[i-2])
		dp[i] = dp[i].Add(dp[i], big.NewInt(1))
	}
	fmt.Fprintln(writer, new(big.Int).Mod(dp[n], big.NewInt(1000000007)))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
