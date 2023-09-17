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

	X int
)

// 난이도: Silver 5
// 메모리: 912KB
// 시간: 4ms
// 분류: 수학, 구현, 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	X = scanInt()
}

func Solve() {
	l, r := 1, 5000
	for l <= r {
		mid := (l + r) / 2
		if (mid * (mid + 1) / 2) >= X {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	var x, y int

	if l%2 == 0 {
		x, y = X-(l*(l-1)/2), l+1-(X-(l*(l-1)/2))
	} else {
		x, y = l+1-(X-(l*(l-1)/2)), X-(l*(l-1)/2)
	}

	fmt.Fprintf(writer, "%d/%d\n", x, y)
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
