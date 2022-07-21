package bj11050

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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	result := factorial(n) / (factorial(k) * factorial(n-k))
	fmt.Fprintln(writer, result)
}

func factorial(n int) int {
	result := 1
	if n == 1 {
		return result
	}
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
