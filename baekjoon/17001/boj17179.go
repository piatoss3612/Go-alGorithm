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

	N, M, L int
	points  []int
)

// 난이도: Gold 5
// 메모리: 900KB
// 시간: 64ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, L = scanInt(), scanInt(), scanInt()
	points = make([]int, M)
	for i := 0; i < M; i++ {
		points[i] = scanInt()
	}
	sort.Ints(points)
}

func Solve() {
	for i := 1; i <= N; i++ {
		x := scanInt()

		l, r := 1, 4000000

		for l <= r {
			m := (l + r) / 2
			// m이 최소 길이라고 가정했을 때, x번 이상 커팅할 수 있는지 확인한다.
			if isPossible(m, x) {
				l = m + 1
			} else {
				r = m - 1
			}
		}

		fmt.Fprintln(writer, r) // 최솟값의 최댓값을 구하는 것이므로 r(upper bound)을 출력한다.
	}
}

func isPossible(m, target int) bool {
	from := 0
	cnt := 0 // 조각의 개수
	for i := 0; i < M; i++ {
		if points[i]-from >= m {
			from = points[i]
			cnt++
		}
		
		// 조각의 개수가 커팅횟수보다 많은지 확인한다.
		if cnt > target {
			return true
		}
	}

	// 마지막 조각을 확인한다.
	if L-from >= m {
		cnt++
	}

	// 조각의 개수가 커팅횟수보다 많은지 확인한다.
	return cnt > target
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
