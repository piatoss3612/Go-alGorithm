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

	S, T int
	a    [MAX + 1]int
)

const (
	MAX = 10000000
)

// 난이도: Gold 5
// 메모리: 79044KB
// 시간: 1864ms
// 분류: 수학, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S, T = scanInt(), scanInt()

	for i := 1; i <= MAX; i += 2 {
		for j := i; j <= MAX; j += i {
			a[j] -= 1
		}
	}
	for i := 2; i <= MAX; i += 2 {
		for j := i; j <= MAX; j += i {
			a[j] += 1
		}
	}
}

func Solve() {
	ans := 0
	for i := S; i <= T; i++ {
		ans += a[i]
	}
	fmt.Fprintln(writer, ans)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
