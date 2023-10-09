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
	T       int
	n, s    int
	seats   []int
)

// 난이도: Gold 5
// 메모리: 20280KB
// 시간: 788ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	n, s = scanInt(), scanInt()
	seats = make([]int, n+1)
	for i := 1; i <= n; i++ {
		seats[i] = scanInt()
	}
	sort.Ints(seats) // 오름차순 정렬 필요
}

func Solve() {
	fmt.Fprintln(writer, BinarySearch())
}

// 이분 탐색을 통해 가장 가까운 두 좌석의 거리의 최댓값(upper bound)를 찾는다
func BinarySearch() int {
	l, r := 0, 1000000000-1
	for l <= r {
		m := (l + r) / 2
		if Satisfy(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return r
}

// 가장 가까운 두 좌석의 거리가 expected일 때 좌석을 s개 이상 설치할 수 있는지 여부를 확인한다
func Satisfy(expected int) bool {
	prev := seats[1]
	cnt := 1

	for i := 2; i <= n; i++ {
		if seats[i]-prev >= expected {
			cnt++
			prev = seats[i]
		}

		if cnt >= s {
			return true
		}
	}

	return cnt >= s
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
