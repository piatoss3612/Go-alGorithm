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
	N, Q    int
	S       string
	idxR    []int
	idxB    []int
)

// 28140번: 빨강~ 빨강~ 파랑! 파랑! 달콤한 솜사탕!
// https://www.acmicpc.net/problem/28140
// 난이도: 골드 5
// 메모리: 32332 KB
// 시간: 1632 ms
// 분류: 이분 탐색
// 주의: 버퍼 사이즈를 크게 잡아야 함
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 2000000), 2000000)

	Setup()
	Solve()
}

func Setup() {
	N, Q = scanInt(), scanInt()
	S = scanString()

	idxR = make([]int, 0, N)
	idxB = make([]int, 0, N)

	for i := 0; i < N; i++ {
		if S[i] == 'R' {
			idxR = append(idxR, i) // R의 인덱스를 저장
		} else if S[i] == 'B' {
			idxB = append(idxB, i) // B의 인덱스를 저장
		}
	}
}

func Solve() {
	for i := 1; i <= Q; i++ {
		l, r := scanInt(), scanInt()

		if len(idxR) < 2 || len(idxB) < 2 {
			fmt.Fprintln(writer, -1)
			continue
		}

		// 문자열 S에서 RRBB를 만들 수 있는 a, b, c, d를 찾아야 함 (a < b < c < d)

		startR := lowerBound(idxR, l) // l, r 구간에서 가장 먼저 나오는 R의 인덱스
		endR := upperBound(idxR, r)   // l, r 구간에서 가장 나중에 나오는 R의 인덱스
		startB := lowerBound(idxB, l) // l, r 구간에서 가장 먼저 나오는 B의 인덱스
		endB := upperBound(idxB, r)   // l, r 구간에서 가장 나중에 나오는 B의 인덱스

		// R, B가 모두 2개 이상이어야 함
		if endR-startR+1 < 2 || endB-startB+1 < 2 {
			fmt.Fprintln(writer, -1)
			continue
		}

		// 겹치는 구간으로 인해 각각의 구간이 2개 이상의 원소를 가지지 않는 경우
		if idxR[startR+1] >= idxB[endB-1] {
			fmt.Fprintln(writer, -1)
			continue
		}

		// 가능한 경우를 출력
		fmt.Fprintf(writer, "%d %d %d %d\n", idxR[startR], idxR[startR+1], idxB[endB-1], idxB[endB])
	}
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

func upperBound(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
