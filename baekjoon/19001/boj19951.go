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

	N, M int
	camp []int // 연병장의 높이
	psum []int // 연병장 각 칸의 높이 변화
)

// 난이도: Gold 5
// 메모리: 6228KB
// 시간: 76ms
// 분류: 누적 합
// 시간복잡도: O(N)
// 공간 복잡도: O(N)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	camp = make([]int, N+1)
	psum = make([]int, N+2)

	for i := 1; i <= N; i++ {
		camp[i] = scanInt()
	}

	for i := 1; i <= M; i++ {
		a, b, k := scanInt(), scanInt(), scanInt()
		// a~b 칸의 높이를 k만큼 증가시킨다
		psum[a] += k   // 작업을 시작하는 칸에 k만큼 더해준다
		psum[b+1] -= k // 작업을 끝나는 칸의 다음 칸에 k만큼 빼준다
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		psum[i] += psum[i-1] // 연병장 i번째 칸의 높이 변화량
		camp[i] += psum[i]   // 연병장 i번째 칸의 높이
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", camp[i])
	}
	fmt.Fprintln(writer)
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
