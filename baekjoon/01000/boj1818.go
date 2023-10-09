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
	N       int
)

// 메모리: 3848KB
// 시간: 40ms
// 책 번호의 가장 긴 증가하는 부분 수열 lis를 구하고 전체 N개의 책 중에 lis만큼을 뺀 나머지 책을 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	lis := make([]int, 0, N)

	for i := 0; i < N; i++ {
		x := scanInt()

		// lis가 비어있거나 입력값 x가 lis의 마지막 값보다 큰 경우
		if len(lis) == 0 || x > lis[len(lis)-1] {
			lis = append(lis, x)
			continue
		}

		// x가 lis의 마지막 값보다 작거나 같은 경우
		// 이분 탐색으로 x의 lower bound를 찾는다
		l, r := 0, len(lis)-1
		for l < r {
			m := (l + r) / 2
			if x <= lis[m] {
				r = m
			} else {
				l = m + 1
			}
		}
		lis[r] = x
	}

	// 전체 N개의 책 중에 lis의 길이를 뺀 값이 책을 정렬해야 하는 횟수
	fmt.Fprintln(writer, N-len(lis))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
