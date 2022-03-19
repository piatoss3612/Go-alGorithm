package main

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
	g := scanInt()

	a, b := 1, 2

	ans := []int{}

	for a < b {
		tmp := b*b - a*a
		if tmp > g {
			a += 1
		} else if tmp < g {
			b += 1
		} else {
			ans = append(ans, b)
			b += 1
		}
	}

	if len(ans) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		for _, v := range ans {
			fmt.Fprintln(writer, v)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
