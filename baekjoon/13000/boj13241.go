package bj13241

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
	a, b := scanInt64(), scanInt64()
	fmt.Fprintln(writer, lcm(a, b))
}

func scanInt64() int64 {
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return n
}

func gcd(a, b int64) int64 {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	d := gcd(a, b)
	return a * b / d
}
