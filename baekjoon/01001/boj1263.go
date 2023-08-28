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
	N       int
	works   []Work
)

type Work struct {
	T, S int
}

// 난이도: Gold 5
// 메모리: 944KB
// 시간: 4ms
// 분류: 그리디 알고리즘, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	works = make([]Work, N)
	for i := 0; i < N; i++ {
		works[i] = Work{
			T: scanInt(),
			S: scanInt(),
		}
	}

	// 작업을 마무리해야 하는 시간을 기준으로 내림차순 정렬
	sort.Slice(works, func(i, j int) bool {
		return works[i].S > works[j].S
	})
}

func Solve() {
	// 작업을 시작하는 시간을 최대한 늦춰야 함 -> 그리디 알고리즘

	start := works[0].S - works[0].T // 마지막 작업을 마무리하기 위해 작업을 시작해야 하는 시간(=이전 작업을 마무리한 시간)계산

	// 역순으로 작업을 진행하면서 작업을 시작해야 하는 시간을 계산
	for i := 1; i < N; i++ {
		// 작업을 시작해야 하는 시간이 음수가 되면 불가능
		if start < 0 {
			fmt.Fprintln(writer, -1)
			return
		}

		// 다음 작업을 시작하는 시간보다 이전 작업을 마무리해야 하는 시간이 빠르면 이전 작업을 마무리해야 하는 시간에 맞춰 작업 시작 시간을 조정
		if start > works[i].S {
			start = works[i].S - works[i].T
		} else {
			// 다음 작업을 시작하는 시간보다 이전 작업을 마무리해야 하는 시간이 늦으면 다음 작업을 시작하는 시간에 맞춰 작업 시작 시간을 조정
			start -= works[i].T
		}
	}

	if start < 0 {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, start)
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
