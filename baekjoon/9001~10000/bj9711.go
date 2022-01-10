package bj9711

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
	var dp [10001]*big.Int
	dp[1], dp[2] = big.NewInt(1), big.NewInt(1)
	for i := 3; i <= 10000; i++ {
		dp[i] = big.NewInt(0).Add(dp[i-1], dp[i-2])
	}

	t := scanInt()
	for i := 1; i <= t; i++ {
		p, q := scanInt(), scanInt()
		fmt.Fprintf(writer, "Case #%d: %d\n", i, big.NewInt(0).Mod(dp[p], big.NewInt(int64(q))))
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
