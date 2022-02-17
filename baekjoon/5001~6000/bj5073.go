package bj5073

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	a, b, c int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	for {
		a, b, c = scanInt(), scanInt(), scanInt()

		if isFinished() {
			return
		}

		if !isValid() {
			fmt.Fprintln(writer, "Invalid")
			continue
		}

		if isEqual() {
			fmt.Fprintln(writer, "Equilateral")
		} else if isIsos() {
			fmt.Fprintln(writer, "Isosceles")
		} else {
			fmt.Fprintln(writer, "Scalene")
		}
	}
}

func isFinished() bool {
	if a == 0 && b == 0 && c == 0 {
		return true
	}
	return false
}

func isValid() bool {
	max := getMax(getMax(a, b), c)
	if max == a {
		return a < b+c
	} else if max == b {
		return b < a+c
	}
	return c < a+b
}

func isEqual() bool {
	return a == b && b == c
}

func isIsos() bool {
	return a == b || b == c || c == a
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
