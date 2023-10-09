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

	N   int
	arr []int
)

// 난이도: Platinum 5
// 메모리: 2568KB
// 시간: 24ms
// 분류: 분할 정복, 세그먼트 트리
// 비고: 2104번 문제에서 구간의 인덱스만 추가된 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N+1)
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	v, l, r := DAC(1, N)
	fmt.Fprintf(writer, "%d\n%d %d\n", v, l, r)
}

func DAC(l, r int) (int, int, int) {
	if l == r {
		return arr[l] * arr[l], l, l
	}

	var ret, maxl, maxr int

	m := (l + r) / 2
	v1, maxl1, maxr1 := DAC(l, m)
	v2, maxl2, maxr2 := DAC(m+1, r)
	if v1 > v2 {
		ret, maxl, maxr = v1, maxl1, maxr1
	} else {
		ret, maxl, maxr = v2, maxl2, maxr2
	}

	sl, sr := m, m+1
	sum := arr[sl] + arr[sr]
	minNum := min(arr[sl], arr[sr])

	for l <= sl && sr <= r {
		temp := sum * minNum
		if temp > ret {
			ret = temp
			maxl, maxr = sl, sr
		}

		if sr < r && (sl == l || arr[sl-1] < arr[sr+1]) {
			sr += 1
			sum += arr[sr]
			minNum = min(minNum, arr[sr])
		} else {
			sl -= 1
			sum += arr[sl]
			minNum = min(minNum, arr[sl])
		}
	}

	return ret, maxl, maxr
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
