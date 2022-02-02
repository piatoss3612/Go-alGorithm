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

	//가장 작은 생성자를 구해야 하므로 n - 100 (얼추 자릿수 * 9보다 큰 수)
	for i := n - 100; i < n; i++ {
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
