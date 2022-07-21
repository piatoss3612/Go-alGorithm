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
	seg     [100001]int
)

// 메모리: 1792KB
// 시간: 20ms
// LIS를 구하는 이분 탐색 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	seg[0] = 0
	segLen := 1
	LIS := 0

	var l, r, m int
	var x int

	for i := 1; i <= n; i++ {
		x = scanInt()

		if x > seg[segLen-1] {
			seg[segLen] = x
			segLen += 1
			LIS += 1
		} else {
			l, r = 0, segLen
			for l < r {
				m = (l + r) / 2
				if x <= seg[m] {
					r = m
				} else {
					l = m + 1
				}
			}
			seg[r] = x
		}
	}

	fmt.Fprintln(writer, n-LIS) // 잘라내야 하는 전선의 수
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
