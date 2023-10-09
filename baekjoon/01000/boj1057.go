package bj1057

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
	_, cnt := scanInt(), 0
	a, b := scanInt(), scanInt()
	for a != b {
		a -= a / 2
		b -= b / 2
		cnt += 1
	}
	fmt.Fprintln(writer, cnt)
}
func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
