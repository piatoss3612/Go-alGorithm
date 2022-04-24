package bj2776

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
	t := scanInt()

	for i := 0; i < t; i++ {
		testCase()
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func testCase() {
	n := scanInt()
	arrN := make([]int, n)
	for i := 0; i < n; i++ {
		arrN[i] = scanInt()
	}
	sort.Ints(arrN)

	m := scanInt()
	arrM := make([]int, m)
	for i := 0; i < m; i++ {
		tmp := scanInt()
		left := 0
		right := n - 1
		isIn := false
		for left <= right {
			mid := (left + right) / 2
			if tmp <= arrN[mid] {
				if tmp == arrN[mid] {
					isIn = true
				}
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if isIn {
			arrM[i] = 1
		}
	}
	for _, v := range arrM {
		fmt.Fprintln(writer, v)
	}
}
