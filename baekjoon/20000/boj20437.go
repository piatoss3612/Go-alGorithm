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
	T       int
	W       string
	K       int
)

// 20437번: 문자열 게임 2
// hhttps://www.acmicpc.net/problem/
// 난이도: 골드 5
// 메모리: 6360 KB
// 시간: 44 ms
// 분류: 문자열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()
	for t := 0; t < T; t++ {
		Setup()
		Solve()
	}
}

func Setup() {
	W = scanString()
	K = scanInt()
}

func Solve() {
	alphabetIndices := make(map[byte][]int) // 키: 알파벳, 값: 해당 알파벳이 문자열 W에 등장하는 인덱스

	for i := 0; i < len(W); i++ {
		alphabetIndices[W[i]] = append(alphabetIndices[W[i]], i) // 해당 알파벳이 문자열 W에 등장하는 인덱스 추가
	}

	ans1 := 987654321
	ans2 := -1

	for _, indices := range alphabetIndices {
		// 해당 알파벳이 K번 이상 등장하지 않으면 패스
		if len(indices) < K {
			continue
		}

		// 문자가 정확히 K개 포함된 부분 연속 문자열의 길이 구하기
		for i := 0; i < len(indices)-K+1; i++ {
			length := indices[i+K-1] - indices[i] + 1
			ans1 = min(ans1, length) // 최소값 갱신 (문자열이 정확히 K개 포함된 부분 연속 문자열의 길이의 최솟값도 첫 번째 글자와 마지막 글자가 동일한 경우에만 가능)
			ans2 = max(ans2, length) // 최대값 갱신
		}
	}

	// 정답이 없는 경우
	if ans1 == 987654321 {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, ans1, ans2)
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
