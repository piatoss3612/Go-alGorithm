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
	N, K, M int
	trimmed []int // 손질된 김밥들
	longest int   // 손질된 김밥 중에 길이가 가장 긴 김밥
)

// 난이도: Silver 2
// 메모리: 25888KB
// 시간: 432ms
// 분류: 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, K, M = scanInt(), scanInt(), scanInt()
	for i := 0; i < N; i++ {
		kimbob := scanInt()

		// 폐기
		if kimbob <= K {
			continue
		}

		var trimmedKimbob int

		if kimbob < K*2 {
			trimmedKimbob = kimbob - K // 한쪽 꼬다리만 자르기
		} else {
			trimmedKimbob = kimbob - 2*K // 양쪽 꼬다리 자르기
		}

		longest = max(longest, trimmedKimbob)
		trimmed = append(trimmed, trimmedKimbob)
	}
}

func Solve() {
	l, r := 1, longest
	ans := -1

	// 이분 탐색으로 김밥 조각의 길이 P 구하기
	for l <= r {
		P := (l + r) / 2 // 김밥 조각의 길이
		cnt := 0         // 김밥 조각의 개수

		for _, kb := range trimmed {
			cnt += kb / P
		}

		if cnt >= M {
			// 김밥 조각의 개수가 M 이상인 경우
			// P의 길이를 더 늘려본다
			ans = max(ans, P)
			l = P + 1
		} else {
			// 김밥 조각의 개수가 M 미만인 경우
			// P의 길이를 줄여야 한다
			r = P - 1
		}
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
