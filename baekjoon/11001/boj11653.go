package bj11653

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	if n == 1 {
		return
	}

	pfs := []int{}

	for i := 2; n > 1 && i <= n; {
		if n%i == 0 {
			n /= i
			pfs = append(pfs, i)
			i = 2
		} else {
			i++
		}
	}
	sort.Ints(pfs)
	for _, v := range pfs {
		fmt.Fprintln(writer, v)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
