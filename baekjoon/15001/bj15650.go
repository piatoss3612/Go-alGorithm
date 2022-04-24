package bj15650

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n, m    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	for i := 1; i <= n-m+1; i++ {
		BackTracking(i, n, m, []int{})
	}
}

func BackTracking(s, e, l int, slice []int) {
	slice = append(slice, s)

	if len(slice) > m {
		return
	}

	if len(slice) == m {
		for _, v := range slice {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
		return
	}

	for i := s + 1; i <= e; i++ {
		BackTracking(i, e, l-1, slice)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
