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
	fenwick [MAX + 1]int
)

const MAX = 200000

// 난이도: Platinum 5
// 메모리: 5396KB
// 시간: 56ms
// 분류: 자료 구조, 펜윅 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	swap := 0
	for i := 1; i <= N; i++ {
		x := scanInt()
		swap += sum(N) - sum(x)
		add(x, 1)
	}
	fmt.Fprintln(writer, swap)
}

func sum(pos int) (ret int) {
	for pos > 0 {
		ret += fenwick[pos]
		pos &= pos - 1
	}
	return
}

func add(pos, val int) {
	for pos <= N {
		fenwick[pos] += val
		pos += pos & -pos
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
