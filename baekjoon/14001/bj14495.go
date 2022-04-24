package bj14495

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	fib     = map[int]int{1: 1, 2: 1, 3: 1}
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	for i := 4; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-3]
	}
	fmt.Fprintln(writer, fib[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
