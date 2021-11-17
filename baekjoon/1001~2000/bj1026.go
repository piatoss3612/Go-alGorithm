package bj1026

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num1, _ := strconv.Atoi(scanner.Text())
		a = append(a, num1)
	}
	sort.Ints(a)

	b := make([]int, 0, n)
	for j := 0; j < n; j++ {
		scanner.Scan()
		num2, _ := strconv.Atoi(scanner.Text())
		b = append(b, num2)
	}
	sort.Ints(b)

	getMin(a, b, n)
}

func getMin(a, b []int, n int) {
	result := 0
	for i := 0; i < n; i++ {
		result += a[i] * b[n-i-1]
	}
	fmt.Fprintln(writer, result)
}
