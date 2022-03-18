package main

import (
	"bufio"
	"fmt"
	"math"
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
	n, m := scanInt(), scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Ints(input)

	min := math.MaxInt64
	s, e := 0, 1

	for s <= e && e < n {
		tmp := input[e] - input[s]
		if tmp >= m {
			min = getMin(min, tmp)
			if s == e {
				e += 1
				if e == n {
					break
				}
			} else {
				s += 1
			}
		} else {
			e += 1
			if e == n {
				break
			}
		}
	}

	fmt.Fprintln(writer, min)
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
