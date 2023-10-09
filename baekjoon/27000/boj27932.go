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
	height  []int
)

// 난이도: Silver 2
// 메모리: 16372KB
// 시간: 348ms
// 분류: 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	height = make([]int, N+2)
	for i := 1; i <= N; i++ {
		height[i] = scanInt()
	}
	height[0] = height[1]
	height[N+1] = height[N]
}

func Solve() {
	l, r := 0, 10000000000
	for l <= r {
		m := (l + r) / 2 // 점수 차이

		cnt := 0 // 지치는 사람의 수
		for i := 1; i <= N; i++ {
			dl, dr := abs(height[i]-height[i-1]), abs(height[i]-height[i+1])
			if dl > m || dr > m {
				cnt++
			}
		}

		// 지치는 사람의 수가 K명 이하면 점수 차이를 줄여본다
		if cnt <= K {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	fmt.Fprintln(writer, l) // 점수 차이의 lower bound
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
