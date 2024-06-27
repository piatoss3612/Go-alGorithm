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
	N, K    int
	arr     []int
)

// 22862번: 가장 긴 짝수 연속한 부분 수열 (large)
// hhttps://www.acmicpc.net/problem/22862
// 난이도: 골드 5
// 메모리: 9160 KB
// 시간: 140 ms
// 분류: 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	// 원소가 1개인 경우
	if N == 1 {
		if arr[0]%2 == 0 {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
		return
	}

	l, r := 0, 0 // left, right
	odds := 0    // 홀수의 개수
	maxLen := 0  // 최대 길이

	for r < N {
		// 홀수인 경우
		if arr[r]%2 != 0 {
			odds++
		}

		// 연속한 부분 수열에서 홀수의 개수가 K개 이하인 경우: 최대 길이 갱신
		if odds <= K {
			maxLen = max(maxLen, r-l+1-odds)
		} else {
			// 연속한 부분 수열에서 홀수의 개수가 K개 초과인 경우: 왼쪽 포인터 이동
			for odds > K {
				if arr[l]%2 != 0 {
					odds--
				}
				l++
			}
		}
		// 오른쪽 포인터 이동
		r++
	}

	fmt.Fprintln(writer, maxLen)
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
