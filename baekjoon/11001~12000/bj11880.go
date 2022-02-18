package bj11880

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
	t := scanInt()

	for i := 1; i <= t; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()

		max := getMax(getMax(a, b), c)

		if max == a {
			fmt.Fprintln(writer, (b+c)*(b+c)+a*a)
		} else if max == b {
			fmt.Fprintln(writer, (a+c)*(a+c)+b*b)
		} else {
			fmt.Fprintln(writer, (a+b)*(a+b)+c*c)
		}
	}
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
