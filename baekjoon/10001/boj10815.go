package bj10815

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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	m := scanInt()
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i] = scanInt()
	}
	sort.Ints(a)
	result := make([]int, m)

	for i := 0; i < m; i++ {
		start := 0
		end := len(a) - 1
		isIn := false
		for start <= end {
			mid := (start + end) / 2
			if b[i] == a[mid] {
				isIn = true
				break
			} else if b[i] < a[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		if isIn {
			result[i] = 1
		}
	}
	for _, v := range result {
		fmt.Fprint(writer, v, " ")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
