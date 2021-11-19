package bj2231

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	for i := n - 50; i < n; i++ {
		if isGenerator(i, n) {
			fmt.Fprintln(writer, i)
			return
		}
	}
	fmt.Fprintln(writer, 0)
}

func isGenerator(m, n int) bool {
	sum := m
	str := strconv.Itoa(m)
	slice := strings.Split(str, "")
	for _, v := range slice {
		digit, _ := strconv.Atoi(v)
		sum += digit
	}
	if sum == n {
		return true
	}
	return false
}
