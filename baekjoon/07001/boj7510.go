package bj7510

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
	a, b, c float64
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	for i := 1; i <= n; i++ {
		a, b, c = scanFloat64(), scanFloat64(), scanFloat64()
		max := getMax(getMax(a, b), c)
		fmt.Fprintf(writer, "Scenario #%d:\n", i)
		if max == a {
			checkPythagoras(b, c, a)
		} else if max == b {
			checkPythagoras(a, c, b)
		} else {
			checkPythagoras(a, b, c)
		}
		fmt.Fprintln(writer)
	}
}

func checkPythagoras(x, y, z float64) {
	if math.Sqrt(x*x+y*y) == z {
		fmt.Fprintln(writer, "yes")
	} else {
		fmt.Fprintln(writer, "no")
	}
}

func getMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return float64(n)
}
