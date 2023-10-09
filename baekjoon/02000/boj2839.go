package bj2839

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

	result := 0
	for n >= 0 {
		if n%5 == 0 {
			fmt.Fprintln(writer, result+(n/5))
			return
		}
		n -= 3
		result++
	}

	fmt.Fprintln(writer, -1)
}
