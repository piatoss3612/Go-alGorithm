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

	N, d, k, c int
	plates     []int
	slide      []int
	kind       [3001]int
)

// 난이도: Silver 1
// 메모리: 1348KB
// 시간: 12ms
// 분류: 구현, 슬라이딩 윈도우
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, d, k, c = scanInt(), scanInt(), scanInt(), scanInt()
	plates = make([]int, N)
	slide = make([]int, k)
}

func Solve() {
	dups := 0
	ans := 0

	for i := 0; i < N+k-1; i++ {
		pos := i % k

		if slide[pos] != 0 {
			kind[slide[pos]]--
			if kind[slide[pos]] > 0 {
				dups--
			}
		}

		if i < N {
			plates[i] = scanInt()
			slide[pos] = plates[i]
		} else {
			slide[pos] = plates[i%N]
		}

		kind[slide[pos]]++
		if kind[slide[pos]] > 1 {
			dups++
		}

		if i >= k-1 {
			if kind[c] == 0 {
				ans = max(ans, k-dups+1)
			} else {
				ans = max(ans, k-dups)
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
