package bj2805

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
	n, m := scanInt(), scanInt()
	trees := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		trees[i] = scanInt()
		if trees[i] > max {
			max = trees[i]
		}
	}
	start, end := 0, max
	result := 0
	for start <= end {
		mid := (start + end) / 2
		sum := 0
		for i := 0; i < n; i++ {
			if trees[i] > mid {
				sum += trees[i] - mid
			}
		}
		if sum >= m {
			if result < mid {
				result = mid
			}
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
