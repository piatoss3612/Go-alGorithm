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
	D, N    int
	oven    []int
	pizza   []int
)

// 1756번: 피자 굽기
// hhttps://www.acmicpc.net/problem/1756
// 난이도: 골드 5
// 메모리: 5808 KB
// 시간: 68 ms
// 분류: 구현, 스택
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	D, N = scanInt(), scanInt()
	oven = make([]int, D)
	pizza = make([]int, N)
	for i := 0; i < D; i++ {
		oven[i] = scanInt()
	}
	for i := 0; i < N; i++ {
		pizza[i] = scanInt()
	}
}

func Solve() {
	// i번째 관의 지름이 i-1번째 관의 지름보다 작은 경우,
	// i-1번째 관에는 i번째 지름보다 큰 지름의 피자가 들어올 수 없다
	for i := 1; i < D; i++ {
		oven[i] = min(oven[i], oven[i-1])
	}

	for {
		// 오븐에 빈 자리가 없거나, 피자가 오븐에 다 들어간 경우
		if len(oven) == 0 || len(pizza) == 0 {
			break
		}

		ovenLen := len(oven) - 1

		// 오븐에 피자를 넣을 수 있는 경우
		if oven[ovenLen] >= pizza[0] {
			pizza = pizza[1:]
		}

		oven = oven[:ovenLen]
	}

	// 모든 피자가 오븐에 들어가지 못한 경우
	if len(pizza) > 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 오븐의 남은 자리 + 1이 마지막 피자가 들어간 위치
	fmt.Fprintln(writer, len(oven)+1)
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
