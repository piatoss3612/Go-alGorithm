package bj11047

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
	n, k := scanInt(), scanInt()

	vs := make([]int, 0, n)

	for i := 0; i < n; i++ {
		v := scanInt()
		vs = append(vs, v)
	}

	cnt := 0
	for i := n - 1; i >= 0; i-- {
		r := k / vs[i]
		if r > 0 {
			cnt += r
			k = k % vs[i]
		}
	}
	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
