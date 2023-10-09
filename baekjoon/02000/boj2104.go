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

	N   int
	arr []int
)

// 난이도: Platinum 5
// 메모리: 2556KB
// 시간: 28ms
// 분류: 분할 정복, 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N+1)
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	fmt.Fprintln(writer, DAC(1, N))
}

// 분할 정복
// 시간 복잡도: O(NlogN)
// 공간 복잡도: O(1)
func DAC(l, r int) int {
	if l == r {
		return arr[l] * arr[l]
	}

	m := (l + r) / 2
	ret := max(DAC(l, m), DAC(m+1, r))

	sl, sr := m, m+1
	sum := arr[sl] + arr[sr]
	minNum := min(arr[sl], arr[sr])

	for l <= sl && sr <= r {
		ret = max(ret, sum*minNum)

		if sr < r && (sl == l || arr[sl-1] < arr[sr+1]) {
			sr += 1
			sum += arr[sr]
			minNum = min(minNum, arr[sr])
		} else {
			sl -= 1
			sum += arr[sl]
			minNum = min(minNum, arr[sl])
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
