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

	N, K int
	arr  []int
)

// 난이도: Gold 4
// 메모리: 3384KB
// 시간: 56ms
// 분류: 이분 탐색, 매개 변수 탐색, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
	sort.Ints(arr) // 오름차순 정렬
}

func Solve() {
	l, r := 0, 1000000000

	for l <= r {
		m := (l + r) / 2

		//  원소간의 차이가 m보다 크거나 같은 경우, 원소의 개수가 K보다 크거나 같은 멀티셋을 만들 수 있는지 확인
		if isPossible(m) {
			l = m + 1 // 가능한 경우, m을 늘려서 더 큰 값으로도 가능한지 확인
		} else {
			r = m - 1 // 불가능한 경우, m을 줄여서 더 작은 값으로 가능한지 확인
		}
	}

	fmt.Fprintln(writer, r) // r이 가능한 최대값(upper bound)이므로 r 출력
}

func isPossible(param int) bool {
	cnt := 1
	prev := arr[0]

	for i := 1; i < N; i++ {
		if arr[i]-prev >= param {
			cnt++
			prev = arr[i]
		}

		if cnt >= K {
			return true
		}
	}

	return cnt >= K
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
