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

// 24023번: 아기 홍윤
// hhttps://www.acmicpc.net/problem/24023
// 난이도: 골드 5
// 메모리: 2460 KB
// 시간: 36 ms
// 분류: 그리디 알고리즘, 비트마스킹
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
	l, r := 0, 0 // 왼쪽, 오른쪽 포인터
	or := 0      // or 연산 결과

	for r < N {
		// arr[r]과 K를 or 연산한 결과가 arr[r]과 같다면
		// arr[r]을 or에 | 연산을 하면 K를 만들 수 있다.
		if arr[r]&K == arr[r] {
			or |= arr[r]

			// or 연산 결과가 K와 같다면
			if or == K {
				break
			}

			// or 연산 결과가 K보다 작다면 r을 증가시킨다.
			r++
			continue
		}

		// arr[r]과 K를 or 연산한 결과가 arr[r]과 다르다면
		// arr[r]을 or에 | 연산을 하면 K를 만들 수 없다.
		// 따라서 l을 증가시키고 or을 초기화한다.
		or = 0
		l = r + 1
		r = l
	}

	// 결과 출력
	if or == K {
		fmt.Fprintln(writer, l+1, r+1)
	} else {
		fmt.Fprintln(writer, -1)
	}
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
