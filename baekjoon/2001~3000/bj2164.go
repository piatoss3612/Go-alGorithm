package bj2164

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

	if n == 1 {
		fmt.Fprintln(writer, 1)
		return
	}

	for i := 2; i <= n; i *= 2 {
		if i == n {
			fmt.Fprintln(writer, n)
			break
		}
		if i*2 > n {
			p := n - i
			result := 2 * p
			fmt.Fprintln(writer, result)
			break
		}
	}
}
