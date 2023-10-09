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
	N, L    int
	puddles []Puddle
)

type Puddle struct {
	start, end int
}

// 난이도: Silver 1
// 메모리: 1340KB
// 시간: 12ms
// 분류: 정렬, 스위핑
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, L = scanInt(), scanInt()
	puddles = make([]Puddle, N)
	for i := 0; i < N; i++ {
		puddles[i] = Puddle{scanInt(), scanInt()}
	}

	sort.Slice(puddles, func(i, j int) bool {
		if puddles[i].start == puddles[j].start {
			return puddles[i].end < puddles[j].end
		}
		return puddles[i].start < puddles[j].start
	})
}

func Solve() {
	r := 0
	ans := 0

	for len(puddles) > 0 {
		puddle := puddles[0]
		puddles = puddles[1:]

		r = max(r, puddle.start)

		if r < puddle.end {
			planks := numOfPlanks(r, puddle.end)
			ans += planks
			r += planks * L
		}
	}

	fmt.Fprintln(writer, ans)
}

func numOfPlanks(x, y int) int {
	res := (y - x) / L
	if (y-x)%L > 0 {
		res++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
