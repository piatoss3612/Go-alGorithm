package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
)

// 난이도: Gold 5
// 메모리: 916KB
// 시간: 4ms
// 분류: 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	// 4와 7로만 이루어진 수를 차례대로 나열했을 때, N번째 수를 구하는 문제
	/*
		   4              7
	   44      47     74      77
	444 447 474 477 744 747 774 777
				...

	규칙 찾기:
	1. 4로 끝나는 수는 홀수번째, 7로 끝나는 수는 짝수번째
	2. i번째 수에 4를 붙이면 2i+1번째 수가 되고, 7을 붙이면 2i+2번째 수가 됨
	*/

	b := strings.Builder{} // 정수 타입을 사용하면 오버플로우가 발생하므로 문자열로 처리

	// 역순으로 N번째 수를 구함
	for N > 0 {
		// 4로 끝나는 수는 홀수번째, 7로 끝나는 수는 짝수번째
		if N%2 == 0 {
			b.WriteRune('7')
			N = (N - 2) / 2 // 7을 붙이면 2i+2번째 수가 되므로 2를 빼고 2로 나눔
		} else {
			b.WriteRune('4')
			N = (N - 1) / 2 // 4를 붙이면 2i+1번째 수가 되므로 1을 빼고 2로 나눔
		}
	}

	// 역순으로 구했으므로 뒤집어줌
	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}

		return string(r)
	}

	fmt.Fprintln(writer, reverse(b.String())) // 출력
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
