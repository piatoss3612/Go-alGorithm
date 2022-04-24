package bj1929

import (
	"bufio"
	"fmt"
	"math"
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
	m, _ := strconv.Atoi(scanner.Text())
	if m == 1 {
		m = 2
	}

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := m; i <= n; i++ {
		if isPrime(i) {
			fmt.Fprintln(writer, i)
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
