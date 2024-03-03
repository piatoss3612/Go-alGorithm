package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	to       [][2]int
	inDegree []int
)

// 1889번: 선물 교환
// hhttps://www.acmicpc.net/problem/1889
// 난이도: 골드 4
// 메모리: 11660 KB
// 시간: 80 ms
// 분류: 위상 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	to = make([][2]int, N+1)
	inDegree = make([]int, N+1)

	for i := 1; i <= N; i++ {
		a, b := scanInt(), scanInt()
		to[i] = [2]int{a, b}
		inDegree[a]++
		inDegree[b]++
	}
}

func Solve() {
	exclude := make([]bool, N+1)
	candidate := make([]int, 0, N)

	// 진입 차수가 2 미만인 노드들을 제거하다보면 남는 모든 노드들의 진입 차수는 2가 된다.

	for i := 1; i <= N; i++ {
		if inDegree[i] < 2 {
			candidate = append(candidate, i)
		}
	}

	for len(candidate) > 0 {
		c := candidate[0]
		candidate = candidate[1:]

		if exclude[c] {
			continue
		}

		exclude[c] = true
		inDegree[to[c][0]]--
		inDegree[to[c][1]]--

		if inDegree[to[c][0]] < 2 {
			candidate = append(candidate, to[c][0])
		}

		if inDegree[to[c][1]] < 2 {
			candidate = append(candidate, to[c][1])
		}
	}

	result := make([]int, 0, N)

	for i := 1; i <= N; i++ {
		if !exclude[i] {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	fmt.Fprintln(writer, len(result))
	for _, r := range result {
		fmt.Fprint(writer, r, " ")
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
