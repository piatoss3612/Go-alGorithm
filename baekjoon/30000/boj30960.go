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
	arr     []int
)

// 30960번: 조별 과제
// https://www.acmicpc.net/problem/30960
// 난이도: 골드 5
// 메모리: 13312 KB
// 시간: 204 ms
// 분류: 누적합
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
	sort.Ints(arr) // O(NlogN)
}

func Solve() {
	// N명 중에 2명씩 (N-3)/2개의 조를 만들고 3명이 하나의 조를 만들어야 한다.
	// 각 조의 어색함은 각 조원의 학번의 최댓값과 최솟값의 차이이다.
	// 어색함을 최소화하기 위해서는 가능하면 연속된 학번을 가진 학생들이 한 조가 되어야 한다.
	// 따라서 우선 학생들을 학번 순으로 정렬한 뒤, 문제에서 요구하는 조의 앞에서부터 둘씩 학생들을 묶어주면 된다.
	// 문제는 어디에서 3명이 한 조가 되어야 하는지를 찾는 것이다.
	tw, tr := make([]int, N), make([]int, N)

	tw[1] = arr[1] - arr[0]
	tr[2] = arr[2] - arr[0]

	// O(N)
	for i := 3; i < N; i++ {
		if i%2 == 1 { // 학생이 짝수명인 경우 (인덱스가 0부터 시작하므로 홀수로 판별)
			tw[i] = tw[i-2] + arr[i] - arr[i-1] // 두 명의 학생으로 새로운 조를 만들어 어색함의 누적합을 계산
		} else { // 3명으로 새로운 조를 만들 수 있는 경우
			tr[i] = min(tr[i-2]+arr[i]-arr[i-1], tw[i-3]+arr[i]-arr[i-2]) // 어색함이 최소가 되는 경우를 선택 (3명으로 새로운 조를 만들거나, 2명으로 새로운 조를 만들거나-이 경우는 이미 3명으로 구성된 조가 있음)
		}
	}

	fmt.Fprintln(writer, tr[N-1]) // 3명으로 만든 조가 포함된 경우의 어색함의 누접합의 최솟값을 출력
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
