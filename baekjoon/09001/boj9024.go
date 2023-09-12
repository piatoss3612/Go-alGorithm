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
	T, N, K int
	arr     []int
)

// 난이도: Gold 5
// 메모리: 30296KB
// 시간: 724ms
// 분류: 정렬, 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
}

func Setup() {
	T = scanInt()
	for i := 0; i < T; i++ {
		N, K = scanInt(), scanInt()
		arr = make([]int, N)
		for j := 0; j < N; j++ {
			arr[j] = scanInt()
		}
		sort.Ints(arr) // 오름차순 정렬
		Solve()
	}
}

func Solve() {
	cnt := 0
	minDiff := 987654321 // 두 수의 합과 K의 차이의 절댓값의 최솟값
	l, r := 0, N-1 // 두 포인터

	for l < r {
		absDiff := abs(arr[l] + arr[r] - K) // 두 수의 합과 K의 차이의 절댓값

		// 두 수의 합과 K의 차이의 절댓값이 최솟값보다 작으면 최솟값 갱신
		if absDiff < minDiff {
			minDiff = absDiff
			cnt = 1 // 최솟값이 갱신되면 카운트 초기화
		} else if absDiff == minDiff { // 최솟값과 같으면 카운트 증가
			cnt++
		}

		// 두 수의 합이 K보다 작으면 왼쪽 포인터 증가, 크면 오른쪽 포인터 감소
		// 같으면 두 포인터 모두 이동
		if arr[l]+arr[r] == K {
			l++
			r--
		} else if arr[l]+arr[r] < K {
			l++
		} else {
			r--
		}
	}

	fmt.Fprintf(writer, "%d\n", cnt)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
