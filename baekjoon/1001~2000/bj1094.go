package bj1094

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
	x, _ := strconv.Atoi(scanner.Text())
	cnt := 0
	for i := 64; i >= 1; i /= 2 {
		if x/i >= 1 {
			cnt += x / i
			x %= i
		}
	}
	fmt.Fprintln(writer, cnt)
}
