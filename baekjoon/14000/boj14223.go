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

	N    int
	work []int
)

// 난이도: Silver 1
// 메모리: 2360KB
// 시간: 20ms
// 분류: 이분 탐색, 매개 변수 탐색, 정렬
func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, 300000), 300000)
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	work = make([]int, N)
	for i := 0; i < N; i++ {
		work[i] = scanInt()
	}
	sort.Ints(work)
}

func Solve() {
	// N일 동안 모든 일을 k시간 씩해서 N일 안에 모든 일을 끝낼 수 있는 가장 큰 k를 구하면 된다 
	l, r := 1, 1000000000
	for l <= r {
		m := (l + r) / 2

		if Possible(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	fmt.Fprintln(writer, r)
}

func Possible(m int) bool {
	for i := 0; i < N; i++ {
		if m*(i+1) > work[i] {
			return false
		}
	}

	return true
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
