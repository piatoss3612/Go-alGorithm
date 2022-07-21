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

// 메모리: 16060KB
// 시간: 268ms
// 12015번과 같은 문제
// 다른 부분은 입력값의 최솟값이 -10억이라는 점
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	seg := make([]int, 1000001)
	seg[0] = -9876543210
	segLen := 1

	ans := 0

	var t int
	var l, r, m int
	for i := 0; i < n; i++ {
		t = scanInt()
		if t > seg[segLen-1] {
			seg[segLen] = t
			segLen += 1
			ans += 1
		} else {
			l, r = 0, segLen
			for l < r {
				m = (l + r) / 2
				if seg[m] >= t {
					r = m
				} else {
					l = m + 1
				}
			}
			seg[r] = t
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
