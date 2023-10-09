package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M, K int
	score   []Score
)

type Score struct {
	a, b int
}

// 난이도: Gold 5
// 메모리: 8652KB
// 시간: 180ms
// 분류: 정렬, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	score = make([]Score, N)
	for i := 0; i < N; i++ {
		score[i] = Score{scanInt(), scanInt()}
	}

	// 심판은 총 K명을 선택해야 하는데 심판이 매긴 점수가 높은 순으로 K명을 선택해야 한다.
	// 주최자가 먼저 M명을 선택하고 그 다음 K명을 선택하게 되면 주최자의 선택에 따라 심판의 선택이 달라진다.
	// 심판은 어찌되었든 자신이 매긴 점수가 높은 순으로 K명을 선택하므로
	// 심판이 먼저 K명을 선택하고 그 다음 주최자가 자신이 매긴 점수가 높은 순으로 M명을 선택하면 최적의 해를 구할 수 있다.

	// 심판 점수가 높은 순으로 내림차순 정렬
	sort.Slice(score, func(i, j int) bool {
		if score[i].b == score[j].b {
			return score[i].a > score[j].a
		}
		return score[i].b > score[j].b
	})
}

func Solve() {
	sum := 0

	for i := 0; i < K; i++ {
		sum += score[i].a
	}

	score = score[K:]

	sort.Slice(score, func(i, j int) bool {
		return score[i].a > score[j].a
	})

	for i := 0; i < M; i++ {
		sum += score[i].a
	}

	fmt.Fprintln(writer, sum)
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
