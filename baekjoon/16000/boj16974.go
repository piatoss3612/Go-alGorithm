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
	N, X    int
	lengths [51]int
	patties [51]int
)

// 16974번: 레벨 햄버거
// hhttps://www.acmicpc.net/problem/16974
// 난이도: 골드 5
// 메모리: 868 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 분할 정복, 재귀
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, X = scanInt(), scanInt()
}

func Solve() {
	lengths[0] = 1
	patties[0] = 1
	for i := 1; i <= N; i++ {
		lengths[i] = 2*lengths[i-1] + 3 // 햄버거 번 + 레벨(i-1) 버거 + 패티 + 레벨(i-1) 버거 + 햄버거 번
		patties[i] = 2*patties[i-1] + 1 // 패티 + 레벨(i-1) 버거 + 패티
	}

	fmt.Fprintln(writer, dac(N, X))
}

func dac(n, x int) int {
	// 기저 사례: n == 0인 경우, 패티 하나만 먹는다.
	if n == 0 {
		return 1
	}

	// x가 1이면 왼쪽 끝 햄버거 번을 먹는다.
	if x == 1 {
		return 0
	}

	// x가 1+lengths[n-1]보다 작으면 왼쪽 레벨(n-1) 버거를 탐색한다.
	if x <= 1+lengths[n-1] {
		return dac(n-1, x-1)
	}

	// x가 2+lengths[n-1]와 같으면 레벨(n-1) 버거와 패티를 먹는다.
	if x == 2+lengths[n-1] {
		return patties[n-1] + 1
	}

	// x가 2+2*lengths[n-1]보다 작으면 레벨(n-1) 버거, 패티를 먹고, 레벨(n-1) 버거를 탐색한다.
	if x <= 2+2*lengths[n-1] {
		return patties[n-1] + 1 + dac(n-1, x-2-lengths[n-1])
	}

	return patties[n]
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
