package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, L, R, X int
	question   []int // 각 문제의 난이도
	ans        = 0
)

// 메모리: 920KB
// 시간: 8ms
// 브루트포스
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, L, R, X = scanInt(), scanInt(), scanInt(), scanInt()
	question = make([]int, N+1)
	for i := 1; i <= N; i++ {
		question[i] = scanInt()
	}

	// 브루트포스: 전수 조사 시작
	for i := 1; i <= N; i++ {
		rec(i, question[i], question[i], question[i])
	}

	fmt.Fprintln(writer, ans)
}

func rec(questNo, totalDiff, minDiff, maxDiff int) {
	if L <= totalDiff && totalDiff <= R && maxDiff-minDiff >= X {
		ans++
	}

	for i := questNo + 1; i <= N; i++ {
		rec(
			i,
			totalDiff+question[i],
			min(minDiff, question[i]),
			max(maxDiff, question[i]),
		)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
