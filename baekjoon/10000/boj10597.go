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

	S      string
	used   [51]bool
	result []int
)

// 10597번: 순열장난
// hhttps://www.acmicpc.net/problem/10597
// 난이도: 골드 5
// 메모리: 904 KB
// 시간: 4 ms
// 분류: 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
}

func Solve() {
	n := findN(len(S)) // N 찾기

	// 백트래킹
	if backtrack(n, 0) {
		// 결과 출력
		for _, num := range result {
			fmt.Fprintf(writer, "%d ", num)
		}
		fmt.Fprintln(writer)
	}
}

func backtrack(n int, idx int) bool {
	// 문자열의 마지막 원소까지 모두 사용한 경우
	if idx == len(S) {
		return true
	}

	// 다음 숫자 찾기: 최대 두자릿수, 문자열 S의 길이를 넘지 않기
	for i := 1; i <= 2 && idx+i <= len(S); i++ {
		num, _ := strconv.Atoi(S[idx : idx+i]) // 숫자 변환

		// 숫자가 1보다 크거나 같고 n보다 작거나 같으며 아직 사용하지 않은 경우
		if num >= 1 && num <= n && !used[num] {
			// 상태 변경
			used[num] = true
			result = append(result, num)

			// idx+i부터 재탐색
			if backtrack(n, idx+i) {
				return true
			}

			// 상태 되돌리기
			result = result[:len(result)-1]
			used[num] = false
		}
	}

	// 적절한 수를 찾기 못한 경우
	return false
}

func findN(length int) int {
	if length <= 9 {
		return length // 1~9까지는 한자리
	}
	return (length-9)/2 + 9 // 10부터는 두자리
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
