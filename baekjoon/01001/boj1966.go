package bj1966

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
	t := scanInt()

	for i := 0; i < t; i++ {
		n, m := scanInt(), scanInt()
		slice := make([]int, 0, n)
		for j := 0; j < n; j++ {
			slice = append(slice, scanInt())
		}

		idx := m
		order := 0

		if len(slice) == 1 {
			fmt.Fprintln(writer, order+1)
			continue
		}
		for len(slice) >= 1 {
			if slice[0] < max(slice[1:]) {
				slice = append(slice[1:], slice[0])
				if idx == 0 {
					idx = len(slice) - 1
				} else {
					idx -= 1
				}
			} else {
				if idx == 0 {
					order += 1
					break
				} else {
					slice = slice[1:]
					idx -= 1
					order += 1
				}
			}
		}

		fmt.Fprintln(writer, order)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func max(slice []int) int {
	max := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
