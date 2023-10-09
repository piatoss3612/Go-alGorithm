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

	K          int
	candidates [14]int
	lotto      [7]int
)

// 난이도: Silver 2
// 메모리: 900KB
// 시간: 4ms
// 분류: 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	for {
		if !Setup() {
			break
		}
	}
}

func Setup() bool {
	K = scanInt()
	if K == 0 {
		return false
	}

	for i := 1; i <= K; i++ {
		candidates[i] = scanInt()
	}

	Solve()

	return true
}

func Solve() {
	backTracking(1, 0)
	fmt.Fprintln(writer)
}

func backTracking(turn, lastIdx int) {
	// 6개를 모두 뽑았을 때
	if turn == 7 {
		for i := 1; i <= 6; i++ {
			fmt.Fprintf(writer, "%d ", lotto[i])
		}
		fmt.Fprintln(writer)
		return
	}

	// 백트래킹으로 다음에 뽑을 수 있는 숫자들을 모두 뽑는다.
	for i := lastIdx + 1; i <= K; i++ {
		lotto[turn] = candidates[i]
		backTracking(turn+1, i)
		lotto[turn] = 0
	}
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
