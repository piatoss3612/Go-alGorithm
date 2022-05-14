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
	seg     [40001]int
)

// 메모리: 1440KB
// 시간: 12ms
// n to n개의 포트들을 연결선이 꼬이지 않고 최대로 연결할 수 있는 방법은
// 곧 LIS(최장 증가 부분수열)을 구하는 것이다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	segLen := 1 // seg의 현재 길이, 비교를 위해 0번 인덱스에는 입력될 수 있는 값의 최솟값보다 작은 값 0
	ans := 0    // LIS

	var temp int
	var l, r, m int
	for i := 1; i <= n; i++ {
		temp = scanInt()

		// 입력값이 segLen-1번째 값보다 큰 경우
		if temp > seg[segLen-1] {
			seg[segLen] = temp
			segLen += 1
			ans += 1
		} else {
			// 이분 탐색: lower bound를 찾아 temp 삽입
			l, r = 0, segLen
			for l < r {
				m = (l + r) / 2
				if seg[m] >= temp {
					r = m
				} else {
					l = m + 1
				}
			}
			seg[r] = temp
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
