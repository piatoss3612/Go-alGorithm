package bj9237

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

	days := make([]int, n)

	for i := 0; i < n; i++ {
		days[i] = scanInt()
	}

	sort.Ints(days)

	max := 0
	cnt := 2
	for j := n - 1; j >= 0; j-- {
		days[j] += cnt
		if days[j] > max {
			max = days[j]
		}
		cnt += 1
	}
	fmt.Fprintln(writer, max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
