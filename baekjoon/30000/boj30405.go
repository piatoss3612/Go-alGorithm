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
	N, M    int
	points  []int
)

// 30405번: 박물관 견학
// hhttps://www.acmicpc.net/problem/30405
// 난이도: 골드 5
// 메모리: 4112 KB
// 시간: 348 ms
// 분류: 수학, 그리디 알고리즘, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	// N: 고양이 수, M: 전시장 수
	N, M = scanInt(), scanInt()
	// points: 고양이들이 방문할 전시장 순서의 시작과 끝의 모음
	points = make([]int, 0, 2*N)
	for i := 1; i <= N; i++ {
		k := scanInt()
		for j := 1; j <= k; j++ {
			x := scanInt()
			if j == 1 || j == k {
				points = append(points, x)
			}
		}
	}
}

func Solve() {
	// 전시장 번호를 오름차순으로 정렬
	sort.Ints(points)

	// 2N개의 전시장 중에 N번째 전시장을 선택
	fmt.Fprintln(writer, points[N-1])
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
