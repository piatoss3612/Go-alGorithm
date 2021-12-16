package bj1788

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	fib     = [1000001]int{0, 1, 1}
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	isNegative := false
	if n < 0 {
		isNegative = true
		n *= -1
	} else if n == 0 {
		fmt.Fprintf(writer, "%d\n%d\n", 0, 0)
		return
	}
	for i := 3; i <= n; i++ {
		fib[i] = (fib[i-1] + fib[i-2]) % 1000000000
	}

	if isNegative {
		if n%2 == 0 {
			fmt.Fprintf(writer, "%d\n%d\n", -1, fib[n])
		} else {
			fmt.Fprintf(writer, "%d\n%d\n", 1, fib[n])
		}
	} else {
		fmt.Fprintf(writer, "%d\n%d\n", 1, fib[n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
