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
	T, N    int
	conn    []int
)

// 메모리: 3104KB
// 시간: 40ms
// 최장 증가 부분 수열 (LIS)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		TestCase()
	}
}

func TestCase() {
	N = scanInt()
	LIS := make([]int, 0, N)

	for i := 0; i < N; i++ {
		x := scanInt()

		// LIS의 길이가 0이거나 LIS의 최댓값이 x보다 작은 경우
		if len(LIS) == 0 || LIS[len(LIS)-1] < x {
			LIS = append(LIS, x)
			continue
		}

		// x가 LIS의 최댓값보다 작은 경우
		// 이분 탐색으로 lower bound를 찾는다
		l, r := 0, len(LIS)-1
		for l < r {
			m := (l + r) / 2
			if x < LIS[m] {
				r = m
			} else {
				l = m + 1
			}
		}
		LIS[r] = x
	}
	fmt.Fprintln(writer, len(LIS))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
