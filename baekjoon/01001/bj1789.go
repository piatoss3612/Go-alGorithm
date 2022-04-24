package bj1789

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
	n := scanInt()

	sub := 1
	cnt := 0
	for n > 0 {
		if n-sub <= sub {
			cnt += 1
			break
		} else {
			n -= sub
			sub += 1
			cnt += 1
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
