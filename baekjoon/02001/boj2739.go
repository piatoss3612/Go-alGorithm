package bj2739

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
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	mulTable(n)
	writer.Flush()
}

func mulTable(n int) {
	for i := 1; i <= 9; i++ {
		fmt.Fprintf(writer, "%d * %d = %d\n", n, i, n*i)
	}
}
