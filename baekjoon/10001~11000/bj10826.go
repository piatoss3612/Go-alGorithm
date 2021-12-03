package bj10826

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
	dp      = map[int]*big.Int{0: big.NewInt(0), 1: big.NewInt(1)}
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	fmt.Fprintln(writer, fib(n))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func fib(n int) *big.Int {
	_, ok := dp[n]
	if !ok {
		dp[n] = big.NewInt(0).Add(fib(n-1), fib(n-2))
	}
	return dp[n]
}
