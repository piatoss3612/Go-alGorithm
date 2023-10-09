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
	S       []int // 매물번호를 저장하는 슬라이스
)

// 난이도: Gold 5
// 메모리: 4212KB
// 시간: 144ms
// 분류: 분할 정복, 유클리드 호제법
// 분할 정복을 통해 GCD 합의 최댓값을 구한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	S = make([]int, N+1)
	for i := 1; i <= N; i++ {
		S[i] = scanInt()
	}
}

func Solve() {
	ans := DAC(1, N)
	fmt.Fprintln(writer, ans)
}

// s: 시작 인덱스
// e: 마지막 인덱스
// 슬라이스 S의 인덱스 s부터 e까지의 범위를 분할 정복을 통해 자취방의 아름다움의 최댓값(GCD 합의 최댓값)을 구한다
func DAC(s, e int) int {
	// 범위 안의 원소가 하나뿐인 경우
	if e-s == 0 {
		return S[e]
	}

	mid := (s + e - 1) / 2 // (원소의 개수 / 2)

	// 왼쪽 s부터 mid까지의 원소를 선택한 경우
	// s부터 mid까지의 최대 공약수와 mid+1부터 e까지의 분할 정복 결과를 더한 값을 구한다
	left := CommonGCD(s, mid) + DAC(mid+1, e)

	// mid+1부터 오른쪽 e까지의 원소를 선택한 경우
	// s부터 mid까지의 분할 정복 결과와 mid+1부터 e까지의 최대 공약수를 더한 값을 구한다
	right := DAC(s, mid) + CommonGCD(mid+1, e)

	return max(left, right)
}

// s: 시작 인덱스
// e: 마지막 인덱스
// 슬라이스 S의 인덱스 s부터 e까지의 매물번호들의 최대 공약수를 구한다
func CommonGCD(s, e int) int {
	gcd := S[s]
	for i := s + 1; i <= e; i++ {
		gcd = GCD(gcd, S[i])
	}
	return gcd
}

// 유클리드 호제법: 최대 공약수 구하기
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
