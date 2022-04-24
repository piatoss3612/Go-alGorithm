package bj16401

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
	m, n := scanInt(), scanInt()
	left, right := 1, 0
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
		if input[i] > right {
			right = input[i]
		}
	}
	result := 0
	for left <= right {
		mid := (left + right) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			cnt += input[i] / mid
		}
		if cnt >= m {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Fprintln(writer, result)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
