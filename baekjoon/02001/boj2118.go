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

	N    int
	psum []int
)

// 난이도: Gold 5
// 분류: 누적합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	// Solve1()
	Solve2()
}

func Setup() {
	N = scanInt()
	psum = make([]int, N+1)
	for i := 1; i <= N; i++ {
		psum[i] = scanInt() + psum[i-1]
	}
}

// 브루트 포스
// 메모리: 1616KB
// 시간: 2460ms
// 시간 복잡도: O(N^2)
// 공간 복잡도: O(N)
func Solve1() {
	dist := 0

	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			cw := psum[j-1] - psum[i-1]            // 시계 방향
			ccw := psum[N] - psum[j-1] + psum[i-1] // 반시계 방향
			dist = max(dist, min(cw, ccw))
		}
	}

	fmt.Fprintln(writer, dist)
}

// 두 포인터
// 메모리: 1600KB
// 시간: 12ms
// 시간 복잡도: O(N)
// 공간 복잡도: O(N)
func Solve2() {
	l, r := 1, 2
	dist := 0

	for l <= r && r <= N {
		cw := psum[r-1] - psum[l-1]            // 시계 방향
		ccw := psum[N] - psum[r-1] + psum[l-1] // 반시계 방향
		dist = max(dist, min(cw, ccw))

		// l~r까지의 시계 방향의 거리가 반시계 방향의 거리보다 크거나 같아지면
		// r을 계속 늘리더라도 시계 방향의 거리가 항상 반시계 방향의 거리보다 커지면서 최대값을 찾을 수 없다.
		// 따라서 l을 늘려서 시계 방향의 거리를 줄여야 한다.
		if cw >= ccw {
			l += 1
		} else {
			r += 1
		}
	}

	fmt.Fprintln(writer, dist)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
