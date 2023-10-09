package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, M   int
	member []int
)

// 난이도: Silver 1
// 메모리: 3300KB
// 시간: 48ms
// 분류: 정렬, 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	member = make([]int, N+1)
	for i := 1; i <= N; i++ {
		member[i] = scanInt()
	}
	sort.Ints(member) // 오름차순 정렬
}

func Solve() {
	cnt := 0
	l, r := 1, N // 왼쪽, 오른쪽 포인터 (가장 작은 값, 가장 큰 값)
	for l < r {
		temp := member[l] + member[r] // 두 포인터의 합
		// 두 포인터의 합이 M보다 크거나 같으면
		if temp >= M {
			cnt += 1 // 카운트 증가
			l += 1  // 왼쪽 포인터 증가
			r -= 1 // 오른쪽 포인터 감소
		} else {
			// 두 포인터의 합이 M보다 작으면
			l += 1 // 왼쪽 포인터 증가
		}
	}

	fmt.Fprintln(writer, cnt)
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
