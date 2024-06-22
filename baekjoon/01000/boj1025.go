package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M    int
	nums    [][]int
)

// 1025번: 제곱수 찾기
// hhttps://www.acmicpc.net/problem/1025
// 난이도: 골드 5
// 메모리: 868 KB
// 시간: 4 ms
// 분류: 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()

	nums = make([][]int, N)
	for i := 0; i < N; i++ {
		b := scanBytes()
		nums[i] = make([]int, M)
		for j := 0; j < M; j++ {
			nums[i][j] = int(b[j] - '0')
		}
	}
}

func Solve() {
	ans := -1

	for r := 0; r < N; r++ { // 행
		for c := 0; c < M; c++ { // 열
			for i := -N + 1; i < N; i++ { // 행의 등차
				for j := -M + 1; j < M; j++ { // 열의 등차
					// 길이가 1인 경우
					if i == 0 && j == 0 {
						if isSquare(nums[r][c]) {
							ans = max(ans, nums[r][c])
						}
						continue
					}

					num := check(i, j, r, c)
					ans = max(ans, num)
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func check(d1, d2, r, c int) int {
	num := 0
	ans := -1

	// r과 c가 각각 d1, d2만큼 이동하면서 제곱수인지 확인
	for r >= 0 && c >= 0 && r < N && c < M {
		num = num*10 + nums[r][c]
		if isSquare(num) {
			ans = max(ans, num)
		}
		r += d1
		c += d2
	}

	return ans
}

func isSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
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
