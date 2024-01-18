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
	N       int
	arr     []int
)

// 31091번: 거짓말
// https://www.acmicpc.net/problem/31091
// 난이도: 골드 5
// 메모리: 27872 KB
// 시간: 164 ms
// 분류: 애드 혹, 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	// 0명이 거짓말 -> arr[i]가 1이상이면 안됨
	// 1명이 거짓말 -> arr[i]가 0이거나 2이상이면 안됨
	// 2명이 거짓말 -> arr[i]가 -1~0이거나 3이상이면 안됨
	// 3명이 거짓말 -> arr[i]가 -2~0이거나 4이상이면 안됨

	/*
		완전탐색 시간초과!

		ans := []int{}

		for i := 0; i <= N; i++ {
			cnt := 0
			for j := 0; j < N; j++ {
				if arr[j] > i || (arr[j] > -i && arr[j] <= 0) {
					cnt++
				}
			}

			if cnt == i {
				ans = append(ans, i)
			}
		}

		sort.Ints(ans)

		fmt.Fprintln(writer, len(ans))
		for _, v := range ans {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
	*/

	/* 누적합을 이용한 풀이 */

	sum := make([]int, N+2) // sum[i]는 i명이 거짓말을 할 수 있는지 여부 (범위가 N을 포함하므로 N+2로 초기화)

	// i의 주장이 맞다면 거짓말을 하는 사람의 경우의 수를 누적합으로 계산
	for i := 0; i < N; i++ {
		if arr[i] <= 0 { // 'arr[i]명 이하가 거짓말을 하고 있다'
			sum[0] += 1
			sum[-arr[i]+1] -= 1
		} else { // 'arr[i]명 이상이 거짓말을 하고 있다'
			sum[arr[i]] += 1
		}
	}

	for i := 1; i <= N; i++ {
		sum[i] += sum[i-1]
	}

	ans := []int{}

	for i := 0; i <= N; i++ {
		if sum[i] == N-i { // i명이 거짓말을 할 수 있으려면 거짓말을 안하는 사람이 sum[i]명이어야 함
			ans = append(ans, i)
		}
	}

	fmt.Fprintln(writer, len(ans))
	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
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
