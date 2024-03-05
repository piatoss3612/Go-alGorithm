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
	T, N    int
)

// 20529번: 가장 가까운 세 사람의 심리적 거리
// hhttps://www.acmicpc.net/problem/20529
// 난이도: 실버 1
// 메모리: 2972 KB
// 시간: 16 ms
// 분류: 브루트포스 알고리즘, 비둘기집 원리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	T = scanInt()
	for t := 0; t < T; t++ {
		Solve()
	}
}

func Solve() {
	N = scanInt()

	mbti := make([]string, N)
	for i := 0; i < N; i++ {
		mbti[i] = scanString()
	}

	if N > 32 {
		fmt.Fprintln(writer, 0)
		return
	}

	ans := 100

	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				cnt := 0
				for l := 0; l < 4; l++ {
					if mbti[i][l] != mbti[j][l] {
						cnt++
					}

					if mbti[j][l] != mbti[k][l] {
						cnt++
					}

					if mbti[k][l] != mbti[i][l] {
						cnt++
					}
				}

				if cnt == 0 {
					fmt.Fprintln(writer, 0)
					return
				}

				ans = min(ans, cnt)
			}
		}
	}

	fmt.Fprintln(writer, ans)
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
