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
	N       int
	puzzle  [100001]int
)

// 난이도: Platinum 5
// 메모리: 2500KB
// 시간: 24ms
// 분류: 분할 정복, 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		puzzle[i] = scanInt()
	}
}

func Solve() {
	fmt.Fprintln(writer, DAC(1, N))
}

func DAC(l, r int) int {
	if l == r {
		return puzzle[l]
	}

	m := (l + r) / 2

	ret := max(DAC(l, m), DAC(m+1, r))

	sl, sr := m, m+1
	squareHeight := min(puzzle[sl], puzzle[sr])

	for l <= sl && sr <= r {
		ret = max(ret, (sr-sl+1)*squareHeight)

		if sr < r && (sl == l || puzzle[sl-1] < puzzle[sr+1]) {
			sr += 1
			squareHeight = min(squareHeight, puzzle[sr])
		} else {
			sl -= 1
			squareHeight = min(squareHeight, puzzle[sl])
		}
	}

	return ret
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

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
