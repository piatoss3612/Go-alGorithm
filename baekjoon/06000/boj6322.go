package bj6322

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
	cnt := 1
	for {
		a, b, c = scanFloat64(), scanFloat64(), scanFloat64()

		if isFinished() {
			return
		}

		fmt.Fprintf(writer, "Triangle #%d\n", cnt)

		if a == -1 {
			a = math.Sqrt(c*c - b*b)
			if a > 0 {
				fmt.Fprintf(writer, "a = %0.3f\n", a)
			} else {
				fmt.Fprintln(writer, "Impossible.")
			}
		} else if b == -1 {
			b = math.Sqrt(c*c - a*a)
			if b > 0 {
				fmt.Fprintf(writer, "b = %0.3f\n", b)
			} else {
				fmt.Fprintln(writer, "Impossible.")
			}
		} else {
			c = math.Sqrt(a*a + b*b)
			fmt.Fprintf(writer, "c = %0.3f\n", c)
		}
		cnt += 1
		fmt.Fprintln(writer)
	}
}

func isFinished() bool {
	return a == 0 && b == 0 && c == 0
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return float64(n)
}
