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

	N, K int
	arr  []int
)

// 난이도: Silver 2
// 메모리: 11288KB
// 시간: 124ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	arr = make([]int, N+1)
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	l, r := 0, 1000000000000
	for l <= r {
		m := (l + r) / 2
		// 점수가 m점이 넘는 학생들에게 K개 이하의 사탕을 줄 수 있는가?
		if possible(m) {
			r = m - 1 // 가능하면 점수를 낮춰본다.
		} else {
			l = m + 1 // 불가능하면 점수를 높여본다.
		}
	}
	fmt.Fprintln(writer, l) // 가능한 최소 점수 출력
}

func possible(m int) bool {
	cnt := 0
	for i := 1; i <= N; i++ {
		if arr[i] > m {
			cnt += (arr[i] - m)
		}
		if cnt > K {
			return false
		}
	}
	return true
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
