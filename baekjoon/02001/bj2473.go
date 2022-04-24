package main

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
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}

	sort.Ints(input)

	min := 3000000001

	var a, b, c int

	// l을 0~n-3으로 고정시켜놓고
	// m과 r을 이용한 두 포인트 알고리즘

	for l := 0; l < n-2; l++ {

		m, r := l+1, n-1

		for m < r {
			tmp := input[l] + input[m] + input[r]

			if tmp == 0 {
				a, b, c = input[l], input[m], input[r]
				break
			}

			if abs(tmp) < min {
				min = abs(tmp)
				a, b, c = input[l], input[m], input[r]
			}

			if tmp > 0 {
				r -= 1
			} else {
				m += 1
			}
		}
	}

	fmt.Fprintln(writer, a, b, c)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
