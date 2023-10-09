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

	N int
	cols []int
)

// 난이도: Platinum 5
// 메모리: 14984KB
// 시간: 196ms
// 분류: 분할 정복
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	cols = make([]int, N+1)

	for i := 1; i <= N; i++ {
		cols[i] = scanInt()
	}
}

func Solve() {
	fmt.Fprintln(writer, DivideAndConquer(1, N))
}

func DivideAndConquer(left, right int) int {
	if left == right {
		return cols[left]
	}

	mid := (left + right) / 2
	ret := max(DivideAndConquer(left, mid), DivideAndConquer(mid+1, right))

	sl, sr := mid, mid+1
	height := min(cols[sl], cols[sr])

	for left <= sl && sr <= right {
		ret = max(ret, height*(sr-sl+1))

		if sr < right && (sl == left || cols[sl-1] < cols[sr+1]) {
			sr += 1
			height = min(height, cols[sr])
		} else {
			sl -= 1
			height = min(height, cols[sl])
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
