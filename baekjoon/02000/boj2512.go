package bj2512

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
	budgets := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		budgets[i] = scanInt()
		sum += budgets[i]
	}
	sort.Ints(budgets)
	m := scanInt()
	if sum <= m {
		fmt.Fprintln(writer, budgets[len(budgets)-1])
		return
	}

	start := 0
	end := budgets[len(budgets)-1]
	result := 0
	for start <= end {
		mid := (start + end) / 2
		sum := 0
		for _, v := range budgets {
			if v > mid {
				sum += mid
			} else {
				sum += v
			}
		}
		if sum > m {
			end = mid - 1
		} else {
			if mid > result {
				result = mid
			}
			start = mid + 1
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
