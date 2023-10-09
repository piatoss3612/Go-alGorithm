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
	LIS     []int
	N, K    int
)

// 메모리: 2776KB
// 시간: 40ms
// 최장 증가 부분 수열(LIS)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T := scanInt()
	for i := 1; i <= T; i++ {
		TestTrading(i)
	}
}

func TestTrading(caseNum int) {
	N, K = scanInt(), scanInt()
	LIS = make([]int, 0, N)

	// 주가가 증가하는 가장 긴 부분 수열을 구하는 문제
	for i := 1; i <= N; i++ {
		stock := scanInt()

		if len(LIS) == 0 || stock > LIS[len(LIS)-1] {
			LIS = append(LIS, stock)
			continue
		}

		l, r, m := 0, len(LIS), 0
		for l < r {
			m = (l + r) / 2
			if stock <= LIS[m] {
				r = m
			} else {
				l = m + 1
			}
		}
		LIS[r] = stock
	}

	fmt.Fprintf(writer, "Case #%d\n", caseNum)

	// K일 동안 주가 증가 추세로 주식을 구매할 수 있다면
	if len(LIS) >= K {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
