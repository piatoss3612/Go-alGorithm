package bj2748

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

	if n <= 2 {
		if n == 0 {
			fmt.Fprintln(writer, 0)
			return
		} else {
			fmt.Fprintln(writer, 1)
			return
		}
	}
	var a, b int64 = 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	fmt.Fprintln(writer, b)
}
