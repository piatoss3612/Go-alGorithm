package bj1629

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
	a, b, c := scanInt(), scanInt(), scanInt()
	fmt.Fprintln(writer, dcMul(a, b, c))
}

func dcMul(a, b, c int) int {
	if b == 1 {
		return a % c
	}
	tmp := dcMul(a, b/2, c)

	if b%2 == 0 {
		return (tmp * tmp) % c
	}

	return (((tmp * tmp) % c) * a) % c
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
